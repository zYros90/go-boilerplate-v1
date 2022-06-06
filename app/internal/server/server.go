package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"go.uber.org/zap"
)

func New(
	conf *config.Config,
	logger *zap.Logger,
) (*echo.Echo, error) {
	echoServer, err := newEcho(conf, logger)
	if err != nil {
		return nil, err
	}
	return echoServer, nil
}

func newEcho(
	config *config.Config,
	logger *zap.Logger,
) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.Server.AllowOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))
	return e, nil
}
