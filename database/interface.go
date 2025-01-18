package database

import (
	"context"

	"github.com/okppop/url-shortener/models"
)

type Database interface {
	CreateUrl(context.Context, models.UrlApiPOSTRequest) (*models.UrlApiPOSTResponse, error)
}
