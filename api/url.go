package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/okppop/url-shortener/models"
	"github.com/okppop/url-shortener/services"
	"github.com/okppop/url-shortener/utils"
)

type URL interface {
	CreateURL(echo.Context) error
	GetURL(echo.Context) error
}

type URLHandler struct {
	service services.Servicer
}

func NewURLHandler(service services.Servicer) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

func (u *URLHandler) CreateURL(c echo.Context) error {
	var requestData models.URLCreateRequest

	err := c.Bind(&requestData)
	if err != nil {
		c.Logger().Warnj(utils.ToLogJ(c.Request(), err.Error()))
		return echo.ErrBadRequest
	}

	// verify original_url
	err = utils.IsURL(requestData.OriginalURL)
	if err != nil {
		c.Logger().Warnj(utils.ToLogJ(c.Request(), err.Error()))
		return echo.ErrBadRequest
	}

	// verify duration_hours
	// 0 => no expire
	// 43800 => 5 years
	if requestData.DurationHours < 0 || requestData.DurationHours > 43800 {
		c.Logger().Warnj(utils.ToLogJ(c.Request(), "duration_hours is invalid"))
		return echo.ErrBadRequest
	}

	responseData, err := u.service.CreateURL(c.Request().Context(), requestData)
	if err != nil {
		c.Logger().Errorj(utils.ToLogJ(c.Request(), err.Error()))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, responseData)
}

func (u *URLHandler) GetURL(c echo.Context) error {
	shortPath := c.Param("short_path")

	originalURL, err := u.service.GetURL(c.Request().Context(), shortPath)
	if err != nil {
		c.Logger().Errorj(utils.ToLogJ(c.Request(), err.Error()))
		return echo.ErrInternalServerError
	}

	if originalURL == "" {
		return echo.ErrNotFound
	}

	return c.Redirect(http.StatusTemporaryRedirect, originalURL)
}
