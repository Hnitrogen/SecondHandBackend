package service

import (
	"context"
	"strconv"

	pb "stuff/api/category/v1"
	"stuff/internal/biz"
)

type CategoryService struct {
	pb.UnimplementedCategoryServer
	uc *biz.CategoryUsecase
}

func NewCategoryService(uc *biz.CategoryUsecase) *CategoryService {
	return &CategoryService{uc: uc}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	return &pb.CreateCategoryReply{}, nil
}
func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryReply, error) {
	return &pb.UpdateCategoryReply{}, nil
}
func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	return &pb.DeleteCategoryReply{}, nil
}
func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryReply, error) {
	return &pb.GetCategoryReply{}, nil
}
func (s *CategoryService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	category, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}

	// 手动映射 biz -> proto
	categoryList := make([]*pb.CategoryWrapper, 0)
	for _, c := range category {
		categoryList = append(categoryList, &pb.CategoryWrapper{
			Id:   strconv.FormatInt(c.ID, 10),
			Name: c.Name,
		})
	}

	return &pb.ListCategoryReply{
		Category: categoryList,
	}, nil
}
