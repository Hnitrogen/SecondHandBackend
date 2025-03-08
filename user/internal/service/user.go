package service

import (
	"context"
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

// 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	user := &data.User{
		ID:      uint(req.Id),
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
	err := s.uc.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{Response: "success"}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	userId := uint(req.Id)
	user, err := s.uc.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Name:    user.Name,
		Avatar:  s.conf.ImageUrl + user.Avatar,
		Address: user.Address,
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
		Address:  user.Address,
		Phone:    user.Phone,
		Id:       int64(user.ID),
	}, nil
}

func (s *UserService) UpdateUserAvatar(ctx context.Context, req *pb.UpdateUserAvatarRequest) (*pb.UpdateUserAvatarReply, error) {
	user := &data.User{
		ID:     uint(req.Id),
		Avatar: req.Avatar,
	}
	err := s.uc.UpdateUserAvatar(ctx, user)

	if err != nil {
		return &pb.UpdateUserAvatarReply{Response: "failed"}, nil
	}
	return &pb.UpdateUserAvatarReply{Response: "success"}, nil
}
