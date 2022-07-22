package service

import (
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/pkg/srvctx"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateTodo endpoint.
func (svc *Service) CreateTodo(c *srvctx.Ctx, req *pb.CreateTodoReq) (*pb.TodoResp, error) {
	todo := parseTodoCreateReq(req)

	// create todo
	todo, err := svc.bizTodo.Create(c, todo)
	if err != nil {
		return nil, err
	}

	// parse to response model
	todoResp := parseTodoResp(todo)

	return todoResp, nil
}

// UpdateTodo endpoint.
func (svc *Service) UpdateTodo(c *srvctx.Ctx, req *pb.CreateTodoReq) (*pb.TodoResp, error) {
	// parse to biz model
	todo := parseTodoCreateReq(req)

	// create todo
	todo, err := svc.bizTodo.Update(c, todo)
	if err != nil {
		return nil, err
	}

	// parse to response model
	todoResp := parseTodoResp(todo)

	return todoResp, nil
}

// GetTodo endpoint.
func (svc *Service) GetTodo(c *srvctx.Ctx, req *pb.GetTodoReq) (*pb.TodoResp, error) {
	todo, err := svc.bizTodo.Get(c, c.Username)
	if err != nil {
		return nil, err
	}

	// parse to response model
	todoResp := parseTodoResp(todo)

	return todoResp, nil
}

// DeleteTodo User endpoint.
func (svc *Service) DeleteTodo(c *srvctx.Ctx, req *pb.DeleteTodoReq) (*emptypb.Empty, error) {
	err := svc.bizTodo.Delete(c, c.Username)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
