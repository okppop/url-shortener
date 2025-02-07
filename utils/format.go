package utils

import (
	"net/http"

	"github.com/labstack/gommon/log"
)

func ToLogJ(r *http.Request, msg string) log.JSON {
	return log.JSON{
		"at":  r.Method + " " + r.URL.Path,
		"msg": msg,
	}
}
