package models

import "time"

type UrlApiPOSTRequest struct {
	OriginalUrl string `json:"original_url"`
	ExpireHours int    `json:"expire_hours"`
}

type UrlApiPOSTResponse struct {
	ShortPath string    `json:"short_path"`
	ExpiredAt time.Time `json:"expired_at"`
}
