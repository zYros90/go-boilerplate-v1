package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"go.uber.org/zap"
)

// New creates new echo server.
func New(
	conf *config.Config,
	logger *zap.Logger,
) (*echo.Echo, error) {
	echoSrv := echo.New()
	echoSrv.Use(middleware.Logger())
	echoSrv.Use(middleware.Recover())

	echoSrv.Use(middleware.CORSWithConfig(middleware.CORSConfig{ // nolint:exhaustruct
		AllowOrigins: conf.Server.AllowOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))

	return echoSrv, nil
}
