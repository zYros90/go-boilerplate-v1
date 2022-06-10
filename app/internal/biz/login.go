package biz

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/pkg/hash"
	"github.com/zYros90/pkg/jwt"
	"go.uber.org/zap"
)

type LoginBiz struct {
	logger *zap.Logger
	conf   *config.Config
	usrBiz *UserBiz
}

// NewLoginUsecase create new login usecase.
func NewLoginUsecase(logger *zap.Logger, conf *config.Config, usrBiz *UserBiz) *LoginBiz {
	return &LoginBiz{
		logger: logger,
		conf:   conf,
		usrBiz: usrBiz,
	}
}

// Login handle business logic for user login.
func (biz *LoginBiz) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := biz.usrBiz.Get(ctx, username)
	if err != nil {
		return "", errors.Wrap(err, "error getting user")
	}

	// compare password
	ok := hash.CheckPasswordHash(password, user.Password)
	if !ok {
		return "", errors.New("wrong password")
	}

	// create token
	return jwt.GenerateJWTHS256(biz.conf.Server.JWTSecret, map[string]string{"username": user.Username}, 1*time.Hour)
}
