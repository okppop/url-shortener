package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/okppop/url-shortener/api"
	"github.com/okppop/url-shortener/conf"
	"github.com/okppop/url-shortener/services"
	"github.com/redis/go-redis/v9"
)

type App struct {
	config *conf.Config
	url    api.URL
	e      *echo.Echo
}

func (a *App) Init(configFilePath string) error {
	cfg, err := conf.Load(configFilePath)
	if err != nil {
		return err
	}
	a.config = cfg

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rdsClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", a.config.Redis.Host, a.config.Redis.Port),
		Password: a.config.Redis.Password,
		DB:       a.config.Redis.Database,
	})

	err = rdsClient.Ping(ctx).Err()
	if err != nil {
		causeErr := context.Cause(ctx)
		if causeErr != nil {
			return causeErr
		}

		return err
	}

	db, err := sql.Open("postgres", a.config.Postgresql.GetDSN())
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	if err != nil {
		causeErr := context.Cause(ctx)
		if causeErr != nil {
			return causeErr
		}

		return err
	}

	a.url = api.NewURLHandler(services.NewService(db, services.Newrds(rdsClient)))

	a.e = echo.New()
	a.e.POST("/api/url", a.url.CreateURL)
	a.e.GET("/:short_path", a.url.GetURL)
	a.e.HideBanner = true
	a.e.Logger.SetLevel(log.INFO)

	return nil
}

func (a *App) Start() {
	err := a.e.Start(a.config.HttpServer.GetListenAddress())
	if err != nil {
		panic(err)
	}
}
