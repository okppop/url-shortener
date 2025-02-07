package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/okppop/url-shortener/models"
	"github.com/redis/go-redis/v9"
)

type cacher interface {
	SetURL(context.Context, *models.URLCreateResponse) error
	GetURL(context.Context, string) (*models.URLCreateResponse, error)
}

type rds struct {
	client *redis.Client
}

func Newrds(client *redis.Client) *rds {
	return &rds{
		client: client,
	}
}

func (r *rds) SetURL(ctx context.Context, responseData *models.URLCreateResponse) error {
	data, err := json.Marshal(responseData)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, responseData.ShortPath, data, 3*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *rds) GetURL(ctx context.Context, shortPath string) (*models.URLCreateResponse, error) {
	data, err := r.client.Get(ctx, shortPath).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	var responseData models.URLCreateResponse
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
