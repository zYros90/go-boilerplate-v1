package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/pkg/srvctx"
)

// CreateTodo endpoint.
func (srv *Server) CreateTodo(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		srv.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	// bind user request
	req := new(pb.CreateTodoReq)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	// validate proto
	err = req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// create server context
	srvCtx := srvctx.New(ctx, username)

	// call svc
	resp, err := srv.svc.CreateTodo(srvCtx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &SrvError{err.Error()})
	}

	// validate proto
	err = resp.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// UpdateTodo endpoint.
func (srv *Server) UpdateTodo(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		srv.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	// bind user request
	req := new(pb.CreateTodoReq)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	// validate proto
	err = req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// create server context
	srvCtx := srvctx.New(ctx, username)

	// call svc
	resp, err := srv.svc.UpdateTodo(srvCtx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &SrvError{err.Error()})
	}

	// validate proto
	err = resp.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// GetTodo endpoint.
func (srv *Server) GetTodo(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		srv.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	todoID := c.Param("todo-id")

	req := &pb.GetTodoReq{
		TodoId: todoID,
	}

	// validate proto
	err := req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// create server context
	srvCtx := srvctx.New(ctx, username)

	// call svc
	resp, err := srv.svc.GetTodo(srvCtx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &SrvError{err.Error()})
	}

	// validate proto
	err = resp.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteTodo Todo endpoint.
func (srv *Server) DeleteTodo(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		srv.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}
	req := &pb.DeleteTodoReq{}

	// validate proto
	err := req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// create server context
	srvCtx := srvctx.New(ctx, username)

	// call svc
	resp, err := srv.svc.DeleteTodo(srvCtx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &SrvError{err.Error()})
	}

	// // validate proto // info: no validation for emptypb.Empty{}
	// err = resp.Validate()
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	// }

	return c.JSON(http.StatusOK, resp)
}
