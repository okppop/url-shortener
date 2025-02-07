package models

import "time"

type URLCreateRequest struct {
	OriginalURL   string `json:"original_url"`
	DurationHours int    `json:"duration_hours"`
}

type URLCreateResponse struct {
	OriginalURL string     `json:"original_url"`
	ShortPath   string     `json:"short_path"`
	ExpiredAt   *time.Time `json:"expired_at"`
}
