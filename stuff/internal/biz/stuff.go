package biz

import (
	"context"
	"math"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"

	userpb "stuff/api/other/user/v1"
	pb "stuff/api/stuff/v1"
)

// Stuff 是业务实体
type Stuff struct {
	ID          int64
	Name        string
	Category    int64
	Price       decimal.Decimal
	Photos      string
	Publisher   int64
	Status      string
	Condition   string
	Description string
}

// StuffRepo 定义存储接口
type StuffRepo interface {
	Create(context.Context, *Stuff) error
	Update(context.Context, *Stuff) error
	Delete(context.Context, int64) error
	Get(context.Context, int64) (*Stuff, error)
	List(context.Context) ([]*Stuff, error)
	ListByCategory(context.Context, int64, int64, int64) ([]*Stuff, error)
	GetTotalByCategory(context.Context, int64) int64
	GetPageCountByCategory(context.Context, int64, int64) int64
	ListAllByPage(context.Context, int64, int64) ([]*Stuff, error)
	GetTotal(context.Context) int64
	ListByUser(context.Context, int64, int64, int64) ([]*Stuff, error)
	GetTotalByUser(context.Context, int64) int64
}

// StuffUsecase 是 Stuff 的业务用例
type StuffUsecase struct {
	repo          StuffRepo
	log           *log.Helper
	UserRpcClient userpb.UserClient
}

// NewStuffUsecase 创建一个新的 Stuff 用例
func NewStuffUsecase(repo StuffRepo, logger log.Logger, rpcClient userpb.UserClient) *StuffUsecase {
	return &StuffUsecase{
		repo:          repo,
		log:           log.NewHelper(logger),
		UserRpcClient: rpcClient,
	}
}

// rpc 获取用户信息
func (uc *StuffUsecase) GetUserInfoRPC(ctx context.Context, id int64) (*userpb.GetUserReply, error) {
	userResp, err := uc.UserRpcClient.GetUser(ctx, &userpb.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	uc.log.WithContext(ctx).Infof("GetUserInfoRPC: %v", userResp)
	return userResp, nil
}

// Create 创建新的 Stuff
func (uc *StuffUsecase) Create(ctx context.Context, s *Stuff) error {
	uc.log.WithContext(ctx).Infof("Create Stuff: %v", s.Name)
	// set status = '上架'
	s.Status = "上架"
	return uc.repo.Create(ctx, s)
}

// Update 更新 Stuff
func (uc *StuffUsecase) Update(ctx context.Context, s *Stuff) error {
	uc.log.WithContext(ctx).Infof("Update Stuff: %v", s.ID)
	return uc.repo.Update(ctx, s)
}

// Delete 删除 Stuff
func (uc *StuffUsecase) Delete(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("Delete Stuff: %v", id)
	return uc.repo.Delete(ctx, id)
}

// Get 获取单个 Stuff
func (uc *StuffUsecase) Get(ctx context.Context, id int64) (*Stuff, error) {
	uc.log.WithContext(ctx).Infof("Get Stuff: %v", id)
	return uc.repo.Get(ctx, id)
}

// List 获取 Stuff 列表
func (uc *StuffUsecase) List(ctx context.Context) ([]*Stuff, error) {
	uc.log.WithContext(ctx).Info("List Stuff")
	return uc.repo.List(ctx)
}

// ListByCategory 获取指定分类的 Stuff 列表
func (uc *StuffUsecase) ListByCategory(ctx context.Context, category int64, page int64, pageSize int64) (pb.ListStuffByCategoryReply, error) {
	uc.log.WithContext(ctx).Infof("List Stuff By Category: %v", category)
	stuffs, err := uc.repo.ListByCategory(ctx, category, page, pageSize)
	if err != nil {
		return pb.ListStuffByCategoryReply{}, err
	}

	result := make([]*pb.StuffWrapper, 0, len(stuffs))
	// Rpc 获取用户信息
	for _, s := range stuffs {
		userResp, err := uc.GetUserInfoRPC(ctx, s.Publisher)
		if err != nil {
			uc.log.WithContext(ctx).Errorf("GetUserInfoRPC: %v", err)
			return pb.ListStuffByCategoryReply{}, err
		}
		decPrice, _ := s.Price.Float64()
		result = append(result, &pb.StuffWrapper{
			Id:     s.ID,
			Name:   s.Name,
			Price:  float32(decPrice),
			Photos: s.Photos,
			Publisher: &pb.UserInfo{
				Name:   userResp.Name,
				Avatar: userResp.Avatar,
			},
			Condition: s.Condition,
		})
	}

	response := &pb.ListStuffByCategoryReply{
		Stuffs:    result,
		Page:      page,
		PageSize:  pageSize,
		Total:     uc.repo.GetTotalByCategory(ctx, category),
		TotalPage: uc.repo.GetPageCountByCategory(ctx, category, pageSize),
	}

	return *response, nil
}

func (uc *StuffUsecase) ListAllByPage(ctx context.Context, page int64, pageSize int64) (pb.ListAllStuffReply, error) {
	uc.log.WithContext(ctx).Info("List All Stuff By Page")
	stuffs, err := uc.repo.ListAllByPage(ctx, page, pageSize)
	if err != nil {
		return pb.ListAllStuffReply{}, err
	}

	result := make([]*pb.StuffWrapper, 0, len(stuffs))
	for _, s := range stuffs {
		userResp, err := uc.GetUserInfoRPC(ctx, s.Publisher)
		if err != nil {
			uc.log.WithContext(ctx).Errorf("GetUserInfoRPC: %v", err)
			return pb.ListAllStuffReply{}, err
		}

		decPrice, _ := s.Price.Float64()
		result = append(result, &pb.StuffWrapper{
			Id:     s.ID,
			Name:   s.Name,
			Price:  float32(decPrice),
			Photos: s.Photos,
			Publisher: &pb.UserInfo{
				Name:   userResp.Name,
				Avatar: userResp.Avatar,
			},
			Condition: s.Condition,
		})
	}

	total := uc.repo.GetTotal(ctx)
	totalPage := math.Ceil(float64(total) / float64(pageSize))
	return pb.ListAllStuffReply{
		Stuffs:    result,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: int64(totalPage),
	}, nil
}

func (uc *StuffUsecase) ListByUser(ctx context.Context, userID int64, page int64, pageSize int64) (pb.ListStuffByUserReply, error) {
	uc.log.WithContext(ctx).Infof("List Stuff By User: %v", userID)
	stuffs, err := uc.repo.ListByUser(ctx, userID, page, pageSize)
	if err != nil {
		return pb.ListStuffByUserReply{}, err
	}

	result := make([]*pb.StuffWrapper, 0, len(stuffs))
	// Get user info once since all items are from same user
	userResp, err := uc.GetUserInfoRPC(ctx, userID)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("GetUserInfoRPC: %v", err)
		return pb.ListStuffByUserReply{}, err
	}

	for _, s := range stuffs {
		decPrice, _ := s.Price.Float64()
		result = append(result, &pb.StuffWrapper{
			Id:     s.ID,
			Name:   s.Name,
			Price:  float32(decPrice),
			Photos: s.Photos,
			Publisher: &pb.UserInfo{
				Name:   userResp.Name,
				Avatar: userResp.Avatar,
			},
			Condition: s.Condition,
		})
	}

	total := uc.repo.GetTotalByUser(ctx, userID)
	totalPage := math.Ceil(float64(total) / float64(pageSize))
	return pb.ListStuffByUserReply{
		Stuffs:    result,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: int64(totalPage),
	}, nil
}
