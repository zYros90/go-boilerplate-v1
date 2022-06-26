package service

import (
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"go.uber.org/zap"
)

// Service struct.
type Service struct {
	pb.UserSvcServer
	pb.LoginSvcServer
	logger   *zap.Logger
	conf     *config.Config
	bizUsr   *biz.UserBiz
	bizLogin *biz.LoginBiz
}

// New create new service and registers routes.
func New(
	conf *config.Config,
	logger *zap.Logger,
	usrBiz *biz.UserBiz,
	bizLogin *biz.LoginBiz,
) *Service {
	svc := &Service{
		logger:   logger,
		conf:     conf,
		bizUsr:   usrBiz,
		bizLogin: bizLogin,
	}
	return svc
}
