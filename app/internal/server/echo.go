package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/service"
	"github.com/zYros90/pkg/echomw"
	"go.uber.org/zap"
)

const (
	timeOut = 5 * time.Second

	todoIDParam = "todo_id"

	usernameKeyEchoCtx = "username"
)

type SrvError struct {
	Error string `json:"error"`
}

type Server struct {
	echoSrv *echo.Echo
	svc     *service.Service
	conf    *config.Config
	logger  *zap.Logger
}

// NewEcho creates new echo server.
func NewEcho(
	conf *config.Config,
	logger *zap.Logger,
	svc *service.Service,
) (*Server, error) {
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

	srv := &Server{
		echoSrv: echoSrv,
		svc:     svc,
		conf:    conf,
		logger:  logger,
	}
	srv.register()
	return srv, nil
}

func (srv *Server) Start(address string) error {
	return srv.echoSrv.Start(address)
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.echoSrv.Shutdown(ctx)
}

func (srv *Server) register() {

	jwtMW := echomw.JWT(srv.conf.Server.JWTSecret)

	// USER
	user := srv.echoSrv.Group("/user/v1")
	user.POST("", srv.CreateUser)
	user.PUT("", srv.UpdateUser, jwtMW)
	user.GET("", srv.GetUser, jwtMW)
	user.DELETE("", srv.DeleteUser, jwtMW)

	// LOGIN
	login := srv.echoSrv.Group("/login/v1")
	login.POST("", srv.Login)

	// TODO
	todo := srv.echoSrv.Group("/todo/v1")
	todo.POST("", srv.CreateTodo, jwtMW)
	todo.PUT("", srv.UpdateTodo, jwtMW)
	todo.GET(fmt.Sprintf("/:%s", todoIDParam), srv.GetTodo, jwtMW)
	todo.DELETE(fmt.Sprintf("/:%s", todoIDParam), srv.DeleteTodo, jwtMW)
}
