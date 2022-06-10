package service

import (
	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func parseUserResp(user *biz.User) *pb.UserResp {
	return &pb.UserResp{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func parseUserCreateReq(userReq *pb.CreateUserReq) *biz.User {
	return &biz.User{ // nolint:exhaustruct
		Username:  userReq.GetUsername(),
		Password:  userReq.GetPassword(),
		Email:     userReq.GetEmail(),
		FirstName: userReq.GetFirstName(),
		LastName:  userReq.GetLastName(),
	}
}
