package service

import (
	"context"
	"strconv"

	pb "stuff/api/stuff/v1"
	"stuff/internal/biz"

	"github.com/shopspring/decimal"
)

type StuffService struct {
	pb.UnimplementedStuffServer
	uc *biz.StuffUsecase
}

func NewStuffService(uc *biz.StuffUsecase) *StuffService {
	return &StuffService{uc: uc}
}

func (s *StuffService) CreateStuff(ctx context.Context, req *pb.CreateStuffRequest) (*pb.CreateStuffReply, error) {
	err := s.uc.Create(ctx, &biz.Stuff{
		Name:        req.Name,
		Category:    req.Category,
		Price:       decimal.NewFromFloat(float64(req.Price)),
		Photos:      req.Photos,
		Publisher:   req.Publisher,
		Condition:   req.Condition,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateStuffReply{}, nil
}

func (s *StuffService) UpdateStuff(ctx context.Context, req *pb.UpdateStuffRequest) (*pb.UpdateStuffReply, error) {
	return &pb.UpdateStuffReply{}, nil
}
func (s *StuffService) DeleteStuff(ctx context.Context, req *pb.DeleteStuffRequest) (*pb.DeleteStuffReply, error) {
	return &pb.DeleteStuffReply{}, nil
}

func (s *StuffService) GetStuff(ctx context.Context, req *pb.GetStuffRequest) (*pb.GetStuffReply, error) {
	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	stuff, err := s.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStuffReply{
		Id:   strconv.FormatInt(stuff.ID, 10),
		Name: stuff.Name,
	}, nil
}

func (s *StuffService) ListStuff(ctx context.Context, req *pb.ListStuffRequest) (*pb.ListStuffReply, error) {
	return &pb.ListStuffReply{}, nil
}
