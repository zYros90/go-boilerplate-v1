package service

import (
	"github.com/labstack/echo/v4"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/pkg/echomw"
	"go.uber.org/zap"
)

// Service struct.
type Service struct {
	logger   *zap.Logger
	conf     *config.Config
	srv      *echo.Echo
	bizUsr   *biz.UserBiz
	bizLogin *biz.LoginBiz
}

// New create new service and registers routes.
func New(
	conf *config.Config,
	logger *zap.Logger,
	srv *echo.Echo,
	usrBiz *biz.UserBiz,
	bizLogin *biz.LoginBiz,
) {
	svc := &Service{
		logger:   logger,
		conf:     conf,
		srv:      srv,
		bizUsr:   usrBiz,
		bizLogin: bizLogin,
	}
	svc.register()
}

func (svc *Service) register() {
	echoSrv := svc.srv

	jwtMW := echomw.JWT(svc.conf.Server.JWTSecret)

	// USER
	user := echoSrv.Group("/user/v1")
	user.POST("", svc.CreateUser)
	user.PUT("", svc.UpdateUser, jwtMW)
	user.GET("", svc.GetUser, jwtMW)
	user.DELETE("", svc.DeleteUser, jwtMW)

	// login
	login := echoSrv.Group("/login/v1")
	login.POST("", svc.Login)
}
