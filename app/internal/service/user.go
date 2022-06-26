package service

import (
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/pkg/srvctx"
)

// CreateUser endpoint.
func (svc *Service) CreateUser(c *srvctx.Ctx, req *pb.CreateUserReq) (*pb.UserResp, error) {
	user := parseUserCreateReq(req)

	// create user
	user, err := svc.bizUsr.Create(c, user)
	if err != nil {
		return nil, err
	}

	// parse to response model
	userResp := parseUserResp(user)

	return userResp, nil
}

// UpdateUser endpoint.
func (svc *Service) UpdateUser(c *srvctx.Ctx, req *pb.CreateUserReq) (*pb.UserResp, error) {
	// parse to biz model
	user := parseUserCreateReq(req)

	// create user
	user, err := svc.bizUsr.Update(c, user)
	if err != nil {
		return nil, err
	}

	// parse to response model
	userResp := parseUserResp(user)

	return userResp, nil
}

// GetUser endpoint.
func (svc *Service) GetUser(c *srvctx.Ctx, req *pb.GetUserReq) (*pb.UserResp, error) {
	user, err := svc.bizUsr.Get(c, c.Username)
	if err != nil {
		return nil, err
	}

	// parse to response model
	userResp := parseUserResp(user)

	return userResp, nil
}

// DeleteUser User endpoint.
func (svc *Service) DeleteUser(c *srvctx.Ctx, req *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	err := svc.bizUsr.Delete(c, c.Username)
	if err != nil {
		return nil, err
	}
	resp := &pb.DeleteUserResp{}
	return resp, nil
}
