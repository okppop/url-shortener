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
		c.Logger().Errorj(utils.EncodeLogJ(c.Request(), err.Error()))
		return echo.ErrBadRequest
	}

	// verify data.OriginalUrl
	err = utils.IsURL(requestData.OriginalUrl)
	if err != nil {
		c.Logger().Warnj(utils.EncodeLogJ(c.Request(), err.Error()))
		return echo.ErrBadRequest
	}

	// verify data.DurationDays
	if requestData.ExpireHours == 0 {
		c.Logger().Warnj(utils.EncodeLogJ(c.Request(), "expired_hours can't be zero"))
		return echo.ErrBadRequest
	}

	var responseData *models.UrlApiPOSTResponse
	responseData, err = h.db.CreateUrl(c.Request().Context(), requestData)
	if err != nil {
		c.Logger().Errorj(utils.EncodeLogJ(c.Request(), err.Error()))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, responseData)
}
