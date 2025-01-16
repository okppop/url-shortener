package database

import "github.com/okppop/url-shortener/models"

type Database interface {
	CreateUrl(models.UrlApiPOSTRequest) (models.UrlApiPOSTResponse, error)
}
