package service

import (
	"context"
	"strconv"
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/data"

	pb "user/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc   *biz.UserUseCase
	conf *conf.Media
}

func NewUserService(uc *biz.UserUseCase, conf *conf.Media) *UserService {
	return &UserService{uc: uc, conf: conf}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user := &data.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	err := s.uc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{Response: "success"}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	userId, _ := strconv.ParseUint(req.Id, 10, 64)
	user, err := s.uc.GetUser(ctx, uint(userId))
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Name:   user.Name,
		Avatar: s.conf.ImageUrl + user.Avatar,
	}, nil
}

func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	resp, err := s.uc.UserLogin(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	user := resp["user"].(*data.User)

	return &pb.UserLoginReply{
		Token:    resp["token"].(string),
		Username: user.Name,
		Email:    user.Email,
		Avatar:   s.conf.ImageUrl + user.Avatar,
		Role:     "",
	}, nil
}
