package data

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
)

func usrKey(usrName string) string {
	return fmt.Sprintf("%s:%s", cacheKeyPrefix, usrName)
}

func (repo *User) cacheUsr(ctx context.Context, usrBiz *biz.User) {
	// set in redis
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(usrBiz)
	if err != nil {
		repo.logger.Sugar().Error("error encoding usrBiz. ", err)
		return
	}

	err = repo.data.redis.Set(ctx, usrKey(usrBiz.Username), buf.Bytes(), cacheExpireTime).Err()
	if err != nil {
		repo.logger.Sugar().Error("error setting key", err)
	}
	repo.logger.Sugar().Debug("cached user successfully")
}

func (repo *User) getCachedUsr(ctx context.Context, username string) (*biz.User, error) {
	cachedUsr, err := repo.data.redis.Get(ctx, usrKey(username)).Bytes()
	if err != nil {
		return nil, err
	}

	// set in redis
	usrBuf := bytes.NewReader(cachedUsr)
	dec := gob.NewDecoder(usrBuf)

	var usrBiz biz.User
	err = dec.Decode(&usrBiz)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding cached usr")
	}
	return &usrBiz, nil
}

func (repo *User) delCachedUsr(ctx context.Context, username string) error {
	err := repo.data.redis.Del(ctx, usrKey(username)).Err()
	if err != nil {
		if !strings.Contains(err.Error(), "redis: nil") {
			return errors.Wrap(err, "error deleting key")
		}
	}
	return nil
}
