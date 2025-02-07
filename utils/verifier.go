package utils

import (
	"errors"
	"net/url"
)

func IsURL(s string) error {
	if len(s) < 8 {
		return errors.New("url too short")
	}

	if len(s) > 8182 {
		return errors.New("url too long")
	}

	u, err := url.ParseRequestURI(s)
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("url neither http or https")
	}

	return nil
}
