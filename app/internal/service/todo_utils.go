package service

import (
	"time"

	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func parseTodoResp(todo *biz.Todo) *pb.TodoResp {
	return &pb.TodoResp{
		TodoId:    todo.TodoID,
		Todo:      todo.Todo,
		DueAt:     timestamppb.New(todo.DueAt),
		NotifyAt:  timestamppb.New(todo.NotifyAt),
		CreatedAt: timestamppb.New(todo.CreatedAt),
		UpdatedAt: timestamppb.New(todo.UpdatedAt),
	}
}

func parseTodoCreateReq(todoReq *pb.CreateTodoReq) *biz.Todo {
	return &biz.Todo{
		TodoID:    "",
		Todo:      todoReq.Todo,
		DueAt:     todoReq.DueAt.AsTime(),
		NotifyAt:  todoReq.NotifyAt.AsTime(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
