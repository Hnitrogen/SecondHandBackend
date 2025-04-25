package service

import (
	"context"
	"strconv"

	pb "stuff/api/stuff/v1"
	"stuff/internal/biz"
	"stuff/internal/conf"

	"github.com/shopspring/decimal"
)

type StuffService struct {
	pb.UnimplementedStuffServer
	uc    *biz.StuffUsecase
	media *conf.Media
}

func NewStuffService(uc *biz.StuffUsecase, conf *conf.Media) *StuffService {

	return &StuffService{
		uc:    uc,
		media: conf,
	}
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

	userResp, err := s.uc.GetUserInfoRPC(ctx, stuff.Publisher)
	if err != nil {
		return nil, err
	}

	price, _ := stuff.Price.Float64()
	return &pb.GetStuffReply{
		Id:          stuff.ID,
		Name:        stuff.Name,
		Description: stuff.Description,
		Publisher: &pb.UserInfo{
			Name:    userResp.Name,
			Avatar:  userResp.Avatar,
			Address: userResp.Address,
			UserId:  userResp.UserId,
		},
		Category:  stuff.Category,
		Price:     float32(price),
		Photos:    s.media.ImageUrl + "stuff/" + stuff.Photos,
		Condition: stuff.Condition,
	}, nil
}

func (s *StuffService) ListStuff(ctx context.Context, req *pb.ListStuffRequest) (*pb.ListStuffReply, error) {
	return &pb.ListStuffReply{}, nil
}

func (s *StuffService) ListStuffByCategory(ctx context.Context, req *pb.ListStuffByCategoryRequest) (*pb.ListStuffByCategoryReply, error) {
	page := req.Page
	pageSize := req.PageSize
	category, _ := strconv.ParseInt(req.Category, 10, 64)

	stuffs, err := s.uc.ListByCategory(ctx, category, page, pageSize)

	for _, stuff := range stuffs.Stuffs {
		stuff.Photos = s.media.ImageUrl + "stuff/" + stuff.Photos
	}

	if err != nil {
		return nil, err
	}

	return &stuffs, nil
}

func (s *StuffService) ListAllStuff(ctx context.Context, req *pb.ListAllStuffRequest) (*pb.ListAllStuffReply, error) {
	stuffs, err := s.uc.ListAllByPage(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	for _, stuff := range stuffs.Stuffs {
		stuff.Photos = s.media.ImageUrl + "stuff/" + stuff.Photos
	}

	return &stuffs, nil
}

func (s *StuffService) ListStuffByUser(ctx context.Context, req *pb.ListStuffByUserRequest) (*pb.ListStuffByUserReply, error) {
	// Set default values for pagination
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	stuffs, err := s.uc.ListByUser(ctx, req.UserId, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// Update photo URLs
	for _, stuff := range stuffs.Stuffs {
		if stuff.Photos != "" {
			stuff.Photos = s.media.ImageUrl + "stuff/" + stuff.Photos
		}
	}

	return &stuffs, nil
}
