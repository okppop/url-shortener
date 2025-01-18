package utils

import (
	"errors"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/labstack/gommon/log"
)

func IsURL(s string) error {
	if len(s) < 8 {
		return errors.New("too short to be a http/https URL")
	}

	if len(s) > 8182 {
		return errors.New("URL over maximum length limit, current limit is 8182 characters")
	}

	u, err := url.ParseRequestURI(s)
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("URL scheme is neither http nor https")
	}

	return nil
}

const letters string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateShortPath() string {
	var shortPath []byte
	for i := 0; i < 10; i++ {
		shortPath = append(shortPath, letters[rand.IntN(len(letters))])
	}

	return string(shortPath)
}

func EncodeLogJ(r *http.Request, message string) log.JSON {
	return log.JSON{
		"at":  r.Method + " " + r.URL.Path,
		"msg": message,
	}
}
