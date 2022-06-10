package service

import (
	"github.com/labstack/echo/v4"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/pkg/echomw"
	"go.uber.org/zap"
)

type Service struct {
	logger   *zap.Logger
	conf     *config.Config
	srv      *echo.Echo
	bizUsr   *biz.UserBiz
	bizLogin *biz.LoginBiz
}

func Init(
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
	svc.Register()
}

func (svc *Service) Register() {
	e := svc.srv

	jwtMW := echomw.JWT(svc.conf.Server.JWTSecret)

	// USER
	user := e.Group("/user/v1")
	user.POST("", svc.CreateUser)
	user.PUT("", svc.UpdateUser, jwtMW)
	user.GET("", svc.GetUser, jwtMW)
	user.DELETE("", svc.DeleteUser, jwtMW)

	// login
	login := e.Group("/login/v1")
	login.POST("", svc.Login)
}
