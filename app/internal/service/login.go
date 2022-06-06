package service

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
)

func (h *Service) Login(c echo.Context) error {
	// bind login request
	loginReq := new(pb.LoginReq)
	err := c.Bind(loginReq)
	if err != nil {
		h.logger.Sugar().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	// check password and get token
	token, err := h.bizLogin.Login(ctx, loginReq.GetUsername(), loginReq.GetPassword())
	if err != nil {
		h.logger.Sugar().Error(err)
		return c.JSON(http.StatusMethodNotAllowed, err)
	}
	resp := &pb.LoginResp{
		Token: token,
	}
	return c.JSON(http.StatusOK, resp)
}
