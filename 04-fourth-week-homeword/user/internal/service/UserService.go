package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"

	pb "user/api/user/v1"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServiceService(user *biz.UserUsecase, logger log.Logger) *UserServiceService {
	return &UserServiceService{}
}

func (s *UserServiceService) Getuser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
func (s *UserServiceService) Createuser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
