package service

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
)

// *******************
// *** create user ***
// *******************.
func (h *Service) CreateUser(c echo.Context) error {
	// bind user request
	userReq := new(pb.CreateUserReq)
	err := c.Bind(userReq)
	if err != nil {
		h.logger.Sugar().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to biz model
	user := parseUserCreateReq(userReq)

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	// create user
	user, err = h.bizUsr.Create(ctx, user)
	if err != nil {
		h.logger.Sugar().Warn(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to response model
	userResp := parseUserResp(user)

	return c.JSON(http.StatusOK, userResp)
}

// *******************
// *** update user ***
// *******************.
func (h *Service) UpdateUser(c echo.Context) error {
	// bind user request
	userReq := new(pb.CreateUserReq)
	err := c.Bind(userReq)
	if err != nil {
		h.logger.Sugar().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to biz model
	user := parseUserCreateReq(userReq)

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	// create user
	user, err = h.bizUsr.Update(ctx, user)
	if err != nil {
		h.logger.Sugar().Warn(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to response model
	userResp := parseUserResp(user)

	return c.JSON(http.StatusOK, userResp)
}

// ****************
// *** get user ***
// ****************.
func (h *Service) GetUser(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		h.logger.Sugar().Warn("no key <username> found in context")
		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	user, err := h.bizUsr.Get(ctx, username)
	if err != nil {
		h.logger.Sugar().Warn(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// parse to response model
	userResp := parseUserResp(user)
	return c.JSON(http.StatusOK, userResp)
}

// *******************
// *** delete user ***
// *******************.
func (h *Service) DeleteUser(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		h.logger.Sugar().Warn("no key <username> found in context")
		return c.JSON(http.StatusMethodNotAllowed, "invalid token")
	}

	// add timeout to context
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.bizUsr.Delete(ctx, username)
	if err != nil {
		h.logger.Sugar().Warn(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "ok")
}
