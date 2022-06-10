package service

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
)

// Login endpoint.
func (svc *Service) Login(c echo.Context) error {
	// bind login request
	loginReq := new(pb.LoginReq)

	err := c.Bind(loginReq)
	if err != nil {
		svc.logger.Sugar().Error(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// check password and get token
	token, err := svc.bizLogin.Login(ctx, loginReq.GetUsername(), loginReq.GetPassword())
	if err != nil {
		svc.logger.Sugar().Error(err)

		return c.JSON(http.StatusMethodNotAllowed, err)
	}

	resp := &pb.LoginResp{
		Token: token,
	}

	return c.JSON(http.StatusOK, resp)
}
