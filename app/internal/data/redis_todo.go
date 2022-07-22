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

func todoKey(usrName string) string {
	return fmt.Sprintf("%s:%s", cacheKeyPrefix, usrName)
}

func (repo *Todo) cacheTodo(ctx context.Context, todoBiz *biz.Todo) {
	// set in redis
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(todoBiz)
	if err != nil {
		repo.logger.Sugar().Error("error encoding todoBiz. ", err)
		return
	}

	err = repo.data.redis.Set(ctx, todoKey(todoBiz.TodoID), buf.Bytes(), cacheExpireTime).Err()
	if err != nil {
		repo.logger.Sugar().Error("error setting key", err)
	}
	repo.logger.Sugar().Debug("cached todo successfully")
}

func (repo *Todo) getCachedTodo(ctx context.Context, todoID string) (*biz.Todo, error) {
	cachedUsr, err := repo.data.redis.Get(ctx, todoKey(todoID)).Bytes()
	if err != nil {
		return nil, err
	}

	// set in redis
	usrBuf := bytes.NewReader(cachedUsr)
	dec := gob.NewDecoder(usrBuf)

	var todoBiz biz.Todo
	err = dec.Decode(&todoBiz)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding cached usr")
	}
	return &todoBiz, nil
}

func (repo *Todo) delCachedTodo(ctx context.Context, todoID string) error {
	err := repo.data.redis.Del(ctx, todoKey(todoID)).Err()
	if err != nil {
		if !strings.Contains(err.Error(), "redis: nil") {
			return errors.Wrap(err, "error deleting key")
		}
	}
	return nil
}
