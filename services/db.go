package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/okppop/url-shortener/models"
	"github.com/okppop/url-shortener/utils"
)

var (
	ErrTimeout           = errors.New("database timeout")
	ErrRetryTooManyTimes = errors.New("retry too many times")
)

const (
	sqlIsShortPathAvailable = "SELECT NOT EXISTS (SELECT 1 FROM url WHERE short_path = $1);"
	sqlCreateURL            = "INSERT INTO url (original_url, short_path, expired_at) VALUES ($1, $2, $3);"
	sqlGetURL               = "SELECT original_url, expired_at FROM url WHERE short_path = $1 AND (expired_at > CURRENT_TIMESTAMP OR expired_at IS NULL);"
)

type Servicer interface {
	CreateURL(context.Context, models.URLCreateRequest) (*models.URLCreateResponse, error)
	GetURL(context.Context, string) (string, error)
}

type Service struct {
	postgresql *sql.DB
	cache      cacher
}

func NewService(db *sql.DB, cache cacher) *Service {
	return &Service{
		postgresql: db,
		cache:      cache,
	}
}

func (s *Service) CreateURL(ctx context.Context, requestData models.URLCreateRequest) (*models.URLCreateResponse, error) {
	subCtx, cancel := context.WithTimeoutCause(ctx, time.Second, ErrTimeout)
	defer cancel()

	shortPath, err := s.getShortPath(subCtx, 5)
	if err != nil {
		causeErr := context.Cause(subCtx)
		if causeErr != nil {
			return nil, causeErr
		}

		return nil, err
	}

	var expiredAt *time.Time
	if requestData.DurationHours != 0 {
		expiredAt = new(time.Time)
		*expiredAt = time.Now().Add(time.Duration(requestData.DurationHours) * time.Hour)
	}

	_, err = s.postgresql.ExecContext(subCtx, sqlCreateURL, requestData.OriginalURL, shortPath, expiredAt)
	if err != nil {
		causeErr := context.Cause(subCtx)
		if causeErr != nil {
			return nil, causeErr
		}

		return nil, err
	}

	responseData := &models.URLCreateResponse{
		OriginalURL: requestData.OriginalURL,
		ShortPath:   shortPath,
		ExpiredAt:   expiredAt,
	}

	err = s.cache.SetURL(subCtx, responseData)
	if err != nil {
		causeErr := context.Cause(subCtx)
		if causeErr != nil {
			return nil, causeErr
		}

		return nil, err
	}

	return responseData, nil
}

func (s *Service) GetURL(ctx context.Context, shortPath string) (string, error) {
	subCtx, cancel := context.WithTimeoutCause(ctx, time.Second, ErrTimeout)
	defer cancel()

	responseData, err := s.cache.GetURL(subCtx, shortPath)
	if err != nil {
		causeErr := context.Cause(subCtx)
		if causeErr != nil {
			return "", causeErr
		}

		return "", err
	}

	if responseData != nil {
		if responseData.ExpiredAt == nil {
			return responseData.OriginalURL, nil
		}

		if responseData.ExpiredAt.After(time.Now()) {
			return responseData.OriginalURL, nil
		}

		return "", nil
	}

	data := new(models.URLCreateResponse)
	err = s.postgresql.QueryRowContext(subCtx, sqlGetURL, shortPath).Scan(
		&data.OriginalURL,
		&data.ExpiredAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}

		causeErr := context.Cause(subCtx)
		if causeErr != nil {
			return "", causeErr
		}

		return "", err
	}
	data.ShortPath = shortPath

	err = s.cache.SetURL(subCtx, data)
	if err != nil {
		causeErr := context.Cause(subCtx)
		if causeErr != nil {
			return "", causeErr
		}

		return "", err
	}

	return data.OriginalURL, nil
}

// n is the max retry times to get availble short_path
func (s *Service) getShortPath(ctx context.Context, n int) (string, error) {
	if n <= 0 {
		return "", ErrRetryTooManyTimes
	}

	shortPath := utils.GenShortPath(10)
	var isAvailable bool

	err := s.postgresql.QueryRowContext(ctx, sqlIsShortPathAvailable, shortPath).Scan(&isAvailable)
	if err != nil {
		return "", err
	}

	if isAvailable {
		return shortPath, nil
	}

	return s.getShortPath(ctx, n-1)
}
