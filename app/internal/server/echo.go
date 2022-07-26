package server

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	pbv1 "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
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

	jwtMW := echomw.JWT(conf.Server.JWTSecret, []string{pbv1.UserSvc_Create_Path})
	echoSrv.Use(jwtMW)

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

	// for _, apiAnnotation := range apiAnnotations{ // For later use -> generate echo handler funcs with protoc
	// 	srv.echoSrv.Add(pbv1.UserSvc_Create_Method, pbv1.UserSvc_Create_Path, srv.CreateUser)
	// }

	// User
	srv.echoSrv.POST(pbv1.UserSvc_Create_Path, srv.CreateUser)
	srv.echoSrv.PUT(pbv1.UserSvc_Update_Path, srv.UpdateUser)
	srv.echoSrv.GET(pbv1.UserSvc_Get_Path, srv.GetUser)
	srv.echoSrv.DELETE(pbv1.UserSvc_Delete_Path, srv.DeleteUser)

	// Login
	srv.echoSrv.POST(pbv1.LoginSvc_Login_Path, srv.Login)

	// Todos
	srv.echoSrv.POST(pbv1.TodoSvc_Create_Path, srv.CreateTodo)
	srv.echoSrv.PUT(pbv1.TodoSvc_Update_Path, srv.UpdateTodo)
	srv.echoSrv.GET(pbv1.TodoSvc_Get_Path, srv.GetTodo)
	srv.echoSrv.DELETE(pbv1.TodoSvc_Delete_Path, srv.DeleteTodo)
}
