package service

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
)

const timeOut = 5 * time.Second

// CreateUser endpoint.
func (svc *Service) CreateUser(c echo.Context) error {
	// bind user request
	userReq := new(pb.CreateUserReq)

	err := c.Bind(userReq)
	if err != nil {
		svc.logger.Sugar().Error(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to biz model
	user := parseUserCreateReq(userReq)

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// create user
	user, err = svc.bizUsr.Create(ctx, user)
	if err != nil {
		svc.logger.Sugar().Warn(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to response model
	userResp := parseUserResp(user)

	return c.JSON(http.StatusOK, userResp)
}

// UpdateUser endpoint.
func (svc *Service) UpdateUser(c echo.Context) error {
	// bind user request
	userReq := new(pb.CreateUserReq)

	err := c.Bind(userReq)
	if err != nil {
		svc.logger.Sugar().Error(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to biz model
	user := parseUserCreateReq(userReq)

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	// create user
	user, err = svc.bizUsr.Update(ctx, user)
	if err != nil {
		svc.logger.Sugar().Warn(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to response model
	userResp := parseUserResp(user)

	return c.JSON(http.StatusOK, userResp)
}

// GetUser endpoint.
func (svc *Service) GetUser(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		svc.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	user, err := svc.bizUsr.Get(ctx, username)
	if err != nil {
		svc.logger.Sugar().Warn(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to response model
	userResp := parseUserResp(user)

	return c.JSON(http.StatusOK, userResp)
}

// DeleteUser User endpoint.
func (svc *Service) DeleteUser(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		svc.logger.Sugar().Warn("no key <username> found in context")

		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), timeOut)
	defer cancel()

	err := svc.bizUsr.Delete(ctx, username)
	if err != nil {
		svc.logger.Sugar().Warn(err)

		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "ok")
}
