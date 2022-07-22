package data

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data/ent/todo"
	"github.com/zYros90/pkg/logger"
	"go.uber.org/zap"
)

type Todo struct {
	data   *Data
	logger *zap.Logger
}

func Newtodo(data *Data, logger *logger.Log) *Todo {
	return &Todo{
		data:   data,
		logger: logger.Logger,
	}
}

func (repo *Todo) Create(ctx context.Context, todoBiz *biz.Todo) (*biz.Todo, error) {
	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transaction client")
	}

	item, err := tx.Todo.Create().
		SetTodoID(todoBiz.TodoID).
		SetTodo(todoBiz.Todo).
		SetNillableDueAt(&todoBiz.DueAt).
		SetNillableNotifyAt(&todoBiz.NotifyAt).
		Save(ctx)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return nil, err
	}

	todoBiz.CreatedAt = item.CreatedAt
	todoBiz.UpdatedAt = item.UpdatedAt

	repo.cacheTodo(ctx, todoBiz)
	return todoBiz, tx.Commit()
}

func (repo *Todo) Update(ctx context.Context, todoBiz *biz.Todo) (*biz.Todo, error) {
	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting transaction client")
	}

	_, err = tx.Todo.Update().
		Where(todo.TodoID(todoBiz.TodoID)).
		SetTodo(todoBiz.Todo).
		SetNillableDueAt(&todoBiz.DueAt).
		SetNillableNotifyAt(&todoBiz.NotifyAt).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return nil, err
	}

	repo.cacheTodo(ctx, todoBiz)
	return todoBiz, tx.Commit()
}

func (repo *Todo) Get(ctx context.Context, todoID string) (*biz.Todo, error) {
	cachedUsrBiz, err := repo.getCachedTodo(ctx, todoID)
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

	item, err := tx.Todo.Query().
		Where(
			todo.TodoIDEQ(todoID),
		).
		Only(ctx)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return nil, err
	}

	todoBiz := &biz.Todo{
		TodoID:    item.TodoID,
		Todo:      item.Todo,
		DueAt:     item.DueAt,
		NotifyAt:  item.NotifyAt,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	repo.cacheTodo(ctx, todoBiz)
	return todoBiz, tx.Commit()
}

func (repo *Todo) Delete(ctx context.Context, todoID string) error {
	tx, err := repo.data.ent.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "error getting transaction client")
	}

	_, err = tx.Todo.Delete().
		Where(
			todo.TodoIDEQ(todoID),
		).
		Exec(ctx)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			repo.logger.Sugar().Error(rErr)
		}

		return err
	}

	err = repo.delCachedTodo(ctx, todoID)
	if err != nil {
		repo.logger.Sugar().Error(err)
	}
	return tx.Commit()
}
