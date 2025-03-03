package biz

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"

	userpb "stuff/api/other/user/v1"
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
		Id: strconv.FormatInt(id, 10),
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
