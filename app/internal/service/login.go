package service

import (
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/pkg/srvctx"
)

// Login endpoint.
func (svc *Service) Login(ctx *srvctx.Ctx, req *pb.LoginReq) (*pb.LoginResp, error) {
	// check password and get token
	token, err := svc.bizLogin.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	resp := &pb.LoginResp{
		Token: token,
	}

	return resp, nil
}
