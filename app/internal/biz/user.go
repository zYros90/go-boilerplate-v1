package biz

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/pkg/hash"
	"go.uber.org/zap"
)

type User struct {
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string

	// from data layer
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBiz struct {
	logger *zap.Logger
	repo   UserRepo
	conf   *config.Config
}

// UserRepo with method signatures from data layer.
type UserRepo interface {
	Create(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	Get(context.Context, string) (*User, error)
	Delete(context.Context, string) error
}

// NewUserUsecase creates new user usercase.
func NewUserUsecase(repo UserRepo, logger *zap.Logger, conf *config.Config) *UserBiz {
	return &UserBiz{
		repo:   repo,
		logger: logger.Named("usrBiz"),
		conf:   conf,
	}
}

func (biz *UserBiz) Create(ctx context.Context, usr *User) (*User, error) {
	// hash password
	hashedPswd, err := hash.HashPassword(usr.Password)
	if err != nil {
		return nil, errors.Wrap(err, "error hashing user password")
	}

	usr.Password = hashedPswd

	// create in data layer
	user, err := biz.repo.Create(ctx, usr)
	if err != nil {
		return nil, errors.Wrap(err, "error creating new user")
	}

	return user, nil
}

func (biz *UserBiz) Update(ctx context.Context, usr *User) (*User, error) {
	// hash password
	if usr.Password != "" {
		hashedPswd, err := hash.HashPassword(usr.Password)
		if err != nil {
			return nil, errors.Wrap(err, "error hashing user password")
		}

		usr.Password = hashedPswd
	}

	// update in data layer
	user, err := biz.repo.Update(ctx, usr)
	if err != nil {
		return nil, errors.Wrap(err, "error updating new user")
	}

	return user, nil
}

func (biz *UserBiz) Get(ctx context.Context, username string) (*User, error) {
	// get from data layer
	user, err := biz.repo.Get(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "error get user: "+username)
	}

	return user, nil
}

func (biz *UserBiz) Delete(ctx context.Context, username string) error {
	err := biz.repo.Delete(ctx, username)
	if err != nil {
		return errors.Wrap(err, "error deleting user: "+username)
	}

	return nil
}
