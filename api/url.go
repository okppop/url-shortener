package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/okppop/url-shortener/database"
	"github.com/okppop/url-shortener/models"
	"github.com/okppop/url-shortener/utils"
)

type ApiUrlHandler struct {
	db database.Database
}

func (h *ApiUrlHandler) POST(c echo.Context) error {
	var requestData models.UrlApiPOSTRequest

	err := c.Bind(&requestData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// verify data.OriginalUrl
	err = utils.IsURL(requestData.OriginalUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// verify data.DurationDays
	if requestData.ExpireHours == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "expire time can't be zero")
	}

	var responseData models.UrlApiPOSTResponse
	responseData, err = h.db.CreateUrl(requestData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, responseData)
}
