package biz

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"go.uber.org/zap"
)

type Todo struct {
	TodoID   string
	Todo     string
	DueAt    time.Time
	NotifyAt time.Time

	// from data layer
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoBiz struct {
	logger *zap.Logger
	repo   TodoRepo
	conf   *config.Config
}

// TodoRepo with method signatures from data layer.
type TodoRepo interface {
	Create(context.Context, *Todo) (*Todo, error)
	Update(context.Context, *Todo) (*Todo, error)
	Get(context.Context, string) (*Todo, error)
	Delete(context.Context, string) error
}

// NewUserUsecase creates new todo usercase.
func NewTodoUsecase(repo TodoRepo, logger *zap.Logger, conf *config.Config) *TodoBiz {
	return &TodoBiz{
		repo:   repo,
		logger: logger.Named("todoBiz"),
		conf:   conf,
	}
}

func (biz *TodoBiz) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	// create in data layer
	todo, err := biz.repo.Create(ctx, todo)
	if err != nil {
		return nil, errors.Wrap(err, "error creating new todo")
	}

	return todo, nil
}

func (biz *TodoBiz) Update(ctx context.Context, todo *Todo) (*Todo, error) {
	// update in data layer
	todo, err := biz.repo.Update(ctx, todo)
	if err != nil {
		return nil, errors.Wrap(err, "error updating todo")
	}

	return todo, nil
}

func (biz *TodoBiz) Get(ctx context.Context, todoID string) (*Todo, error) {
	// get from data layer
	todo, err := biz.repo.Get(ctx, todoID)
	if err != nil {
		return nil, errors.Wrap(err, "error get todo: "+todoID)
	}

	return todo, nil
}

func (biz *TodoBiz) Delete(ctx context.Context, username string) error {
	err := biz.repo.Delete(ctx, username)
	if err != nil {
		return errors.Wrap(err, "error deleting todo: "+username)
	}

	return nil
}
