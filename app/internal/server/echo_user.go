package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/pkg/srvctx"
)

// CreateUser endpoint.
func (srv *Server) CreateUser(c echo.Context) error {
	// bind user request
	req := new(pb.CreateUserReq)

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
	srvCtx := srvctx.New(ctx, req.Username)

	// call svc
	resp, err := srv.svc.CreateUser(srvCtx, req)
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

// UpdateUser endpoint.
func (srv *Server) UpdateUser(c echo.Context) error {
	// bind user request
	req := new(pb.CreateUserReq)

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
	srvCtx := srvctx.New(ctx, req.Username)

	// call svc
	resp, err := srv.svc.UpdateUser(srvCtx, req)
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

// GetUser endpoint.
func (srv *Server) GetUser(c echo.Context) error {
	username, ok := c.Get(usernameKeyEchoCtx).(string)
	if !ok {
		srv.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	req := &pb.GetUserReq{}

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
	resp, err := srv.svc.GetUser(srvCtx, req)
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

// DeleteUser User endpoint.
func (srv *Server) DeleteUser(c echo.Context) error {
	username, ok := c.Get(usernameKeyEchoCtx).(string)
	if !ok {
		srv.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}
	req := &pb.DeleteUserReq{}

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
	resp, err := srv.svc.DeleteUser(srvCtx, req)
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
