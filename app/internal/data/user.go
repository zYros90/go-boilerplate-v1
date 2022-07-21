package data

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data/ent/user"
	"github.com/zYros90/pkg/logger"
	"go.uber.org/zap"
)

type User struct {
	data   *Data
	logger *zap.Logger
}

func NewUser(data *Data, logger *logger.Log) *User {
	return &User{
		data:   data,
		logger: logger.Logger,
	}
}

func (repo *User) Create(ctx context.Context, usrBiz *biz.User) (*biz.User, error) {
	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transaction client")
	}

	item, err := tx.User.Create().
		SetUsername(usrBiz.Username).
		SetPassword(usrBiz.Password).
		SetEmail(usrBiz.Email).
		SetFirstName(usrBiz.FirstName).
		SetLastName(usrBiz.LastName).
		Save(ctx)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return nil, err
	}

	usrBiz.CreatedAt = item.CreatedAt
	usrBiz.UpdatedAt = item.UpdatedAt

	repo.cacheUsr(ctx, usrBiz)
	return usrBiz, tx.Commit()
}

func (repo *User) Update(ctx context.Context, usrBiz *biz.User) (*biz.User, error) {
	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transaction client")
	}

	_, err = tx.User.Update().
		Where(user.UsernameEQ(usrBiz.Username)).
		SetPassword(usrBiz.Password).
		SetEmail(usrBiz.Email).
		SetFirstName(usrBiz.FirstName).
		SetLastName(usrBiz.LastName).
		Save(ctx)

	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return nil, err
	}

	repo.cacheUsr(ctx, usrBiz)
	return usrBiz, tx.Commit()
}

func (repo *User) Get(ctx context.Context, username string) (*biz.User, error) {
	cachedUsrBiz, err := repo.getCachedUsr(ctx, username)
	if err != nil {
		if !strings.Contains(err.Error(), "redis: nil") {
			repo.logger.Sugar().Error(err)
		} else {
			repo.logger.Sugar().Debug("redis nil")
		}
	} else {
		repo.logger.Sugar().Debug("return from cache")
		return cachedUsrBiz, nil
	}

	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transaction client")
	}

	item, err := tx.User.Query().
		Where(
			user.UsernameEQ(username),
		).
		Only(ctx)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return nil, err
	}

	usrBiz := &biz.User{
		Username:  username,
		Password:  item.Password,
		Email:     item.Email,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	repo.cacheUsr(ctx, usrBiz)
	return usrBiz, tx.Commit()
}

func (repo *User) Delete(ctx context.Context, username string) error {
	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "error getting transaction client")
	}

	_, err = tx.User.Delete().
		Where(
			user.UsernameEQ(username),
		).
		Exec(ctx)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return err
	}
	err = repo.delCachedUsr(ctx, username)
	if err != nil {
		repo.logger.Sugar().Error(err)
	}
	return tx.Commit()
}
