package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/okppop/url-shortener/models"
	"github.com/okppop/url-shortener/utils"
)

var (
	ErrGenShortPathTooManyTimes error = errors.New("generate short path too many times")
	ErrSQLTimeout               error = errors.New("SQL timeout")
)

const (
	// sqlIsShortPathDuplicate string = "SELECT true FROM url WHERE short_path = $1;"
	sqlIsShortPathDuplicate string = "SELECT EXISTS (SELECT 1 FROM url WHERE short_path = $1);"
	sqlInsert               string = "INSERT INTO url (original_url, short_path, expired_at) VALUES ($1, $2, $3);"
	// sqlInsert            string = "insert into url (original_url, short_path, expired_at) values ($1, $2, current_timestamp + interval '1 hour' * $3);"
)

type Postgresql struct {
	db *sql.DB
}

func (p *Postgresql) CreateUrl(ctx context.Context, requestData models.UrlApiPOSTRequest) (*models.UrlApiPOSTResponse, error) {
	var shortPath string
	var isDuplicate bool = true // assume true at first

	ctx1, cancel1 := context.WithTimeoutCause(ctx, 2*time.Second, ErrSQLTimeout)
	defer cancel1()

	for i := 0; i < 5; i++ {
		shortPath = utils.GenerateShortPath()

		err := p.db.QueryRowContext(ctx1, sqlIsShortPathDuplicate, shortPath).Scan(&isDuplicate)
		if err != nil {
			causeErr := context.Cause(ctx1)
			if causeErr != nil {
				return nil, causeErr
			}

			return nil, err
		}

		if !isDuplicate {
			break
		}
	}

	if isDuplicate {
		return nil, ErrGenShortPathTooManyTimes
	}

	expiredAt := time.Now().Add(time.Duration(requestData.ExpireHours) * time.Hour)

	ctx2, cancel2 := context.WithTimeoutCause(ctx, 1*time.Second, ErrSQLTimeout)
	defer cancel2()

	_, err := p.db.ExecContext(ctx2, sqlInsert, requestData.OriginalUrl, shortPath, expiredAt)
	if err != nil {
		causeErr := context.Cause(ctx2)
		if causeErr != nil {
			return nil, causeErr
		}

		return nil, err
	}

	return &models.UrlApiPOSTResponse{
		ShortPath: shortPath,
		ExpiredAt: expiredAt,
	}, nil
}
