package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/pkg/srvctx"
)

// CreateUser endpoint.
func (srv *Server) Login(c echo.Context) error {
	// bind login request
	req := new(pb.LoginReq)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
	resp, err := srv.svc.Login(srvCtx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// validate proto
	err = resp.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &SrvError{err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
