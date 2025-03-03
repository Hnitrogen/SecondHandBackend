package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Category struct {
	ID   int64
	Name string
}

type CategoryRepo interface {
	List(context.Context) ([]*Category, error)
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CategoryUsecase) List(ctx context.Context) ([]*Category, error) {
	return uc.repo.List(ctx)
}
