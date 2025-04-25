package data

import (
	"context"
	"math"
	v1 "stuff/api/stuff/v1"
	"stuff/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stuff struct {
	gorm.Model
	ID          int64           `gorm:"primarykey" autoIncrement:"true"`
	Name        string          `gorm:"column:name;type:varchar(255)"`
	Category    int64           `gorm:"column:category;type:int"`
	Price       decimal.Decimal `gorm:"column:price;type:decimal(10,2)"`
	Photos      string          `gorm:"column:photos;"`
	Publisher   int64           `gorm:"column:publisher;type:int"`
	Status      string          `gorm:"column:status;type:varchar(255)"`
	Condition   string          `gorm:"column:condition;type:varchar(255)"`
	Description string          `gorm:"column:description;type:text"`
}

// StuffRepo 实现 biz.StuffRepo 接口
type stuffRepo struct {
	data *Data
	log  *log.Helper
}

// NewStuffRepo .
func NewStuffRepo(data *Data, logger log.Logger) biz.StuffRepo {
	return &stuffRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 实现 biz.StuffRepo 接口
func (r *stuffRepo) Create(ctx context.Context, g *biz.Stuff) error {
	stuff := &Stuff{
		ID:          g.ID,
		Name:        g.Name,
		Category:    g.Category,
		Price:       g.Price,
		Photos:      g.Photos,
		Publisher:   g.Publisher,
		Status:      g.Status,
		Condition:   g.Condition,
		Description: g.Description,
	}

	return r.data.db.WithContext(ctx).Create(stuff).Error
}

// Update 实现 biz.StuffRepo 接口
func (r *stuffRepo) Update(ctx context.Context, g *biz.Stuff) error {
	stuff := &Stuff{
		ID:   g.ID,
		Name: g.Name,
	}
	return r.data.db.WithContext(ctx).Where("id = ?", stuff.ID).Updates(stuff).Error
}

// Delete 实现 biz.StuffRepo 接口
func (r *stuffRepo) Delete(ctx context.Context, id int64) error {
	return r.data.db.WithContext(ctx).Where("id = ?", id).Delete(&Stuff{}).Error
}

// Get 实现 biz.StuffRepo 接口
func (r *stuffRepo) Get(ctx context.Context, id int64) (*biz.Stuff, error) {
	var stuff Stuff
	if err := r.data.db.WithContext(ctx).Where("id = ?", id).First(&stuff).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &biz.Stuff{
		ID:          stuff.ID,
		Name:        stuff.Name,
		Category:    stuff.Category,
		Price:       stuff.Price,
		Photos:      stuff.Photos,
		Publisher:   stuff.Publisher,
		Status:      stuff.Status,
		Condition:   stuff.Condition,
		Description: stuff.Description,
	}, nil
}

// List 实现 biz.StuffRepo 接口
func (r *stuffRepo) List(ctx context.Context) ([]*biz.Stuff, error) {
	var stuffs []*Stuff
	if err := r.data.db.WithContext(ctx).Find(&stuffs).Error; err != nil {
		return nil, err
	}

	result := make([]*biz.Stuff, 0, len(stuffs))
	for _, s := range stuffs {
		result = append(result, &biz.Stuff{
			ID:   s.ID,
			Name: s.Name,
		})
	}
	return result, nil
}

// ConvertToProto 将数据库模型转换为proto消息
func (s *Stuff) ConvertToProto() *v1.GetStuffReply {
	return &v1.GetStuffReply{
		Id:   s.ID,
		Name: s.Name,
	}
}

// ListByCategory 实现 biz.StuffRepo 接口
func (r *stuffRepo) ListByCategory(ctx context.Context, category int64, page int64, pageSize int64) ([]*biz.Stuff, error) {
	var stuffs []*Stuff

	offset := (page - 1) * pageSize // 先计算偏移量，避免类型转换问题
	r.log.WithContext(ctx).Infof("ListByCategory - offset: %v, page: %v, pageSize: %v", offset, page, pageSize)

	if err := r.data.db.WithContext(ctx).
		Where("category = ?", category).
		Offset(int(offset)). // 使用计算好的 offset
		Limit(int(pageSize)).
		Find(&stuffs).Error; err != nil {
		r.log.WithContext(ctx).Errorf("ListByCategory: %v", err)
		return nil, err
	}

	result := make([]*biz.Stuff, 0, len(stuffs))
	for _, s := range stuffs {
		result = append(result, &biz.Stuff{
			ID:          s.ID,
			Name:        s.Name,
			Category:    s.Category,
			Price:       s.Price,
			Photos:      s.Photos,
			Publisher:   s.Publisher,
			Status:      s.Status,
			Condition:   s.Condition,
			Description: s.Description,
		})
	}
	return result, nil
}

func (r *stuffRepo) GetTotalByCategory(ctx context.Context, category int64) int64 {
	var total int64
	r.data.db.WithContext(ctx).Model(&Stuff{}).Where("category = ?", category).Count(&total)
	return total
}

func (r *stuffRepo) GetPageCountByCategory(ctx context.Context, category int64, pageSize int64) int64 {
	var total int64
	r.data.db.WithContext(ctx).Model(&Stuff{}).Where("category = ?", category).Count(&total)

	return int64(math.Ceil(float64(total) / float64(pageSize)))
}

func (r *stuffRepo) ListAllByPage(ctx context.Context, page int64, pageSize int64) ([]*biz.Stuff, error) {
	var stuffs []*Stuff
	offset := (page - 1) * pageSize
	if err := r.data.db.WithContext(ctx).
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&stuffs).Error; err != nil {
		return nil, err
	}

	result := make([]*biz.Stuff, 0, len(stuffs))
	for _, s := range stuffs {
		result = append(result, &biz.Stuff{
			ID:          s.ID,
			Name:        s.Name,
			Category:    s.Category,
			Price:       s.Price,
			Photos:      s.Photos,
			Publisher:   s.Publisher,
			Status:      s.Status,
			Condition:   s.Condition,
			Description: s.Description,
		})
	}
	return result, nil
}

func (r *stuffRepo) GetTotal(ctx context.Context) int64 {
	var total int64
	r.data.db.WithContext(ctx).Model(&Stuff{}).Count(&total)
	return total
}

func (r *stuffRepo) ListByUser(ctx context.Context, userID int64, page int64, pageSize int64) ([]*biz.Stuff, error) {
	var stuffs []*Stuff
	offset := (page - 1) * pageSize

	if err := r.data.db.WithContext(ctx).
		Where("publisher = ?", userID).
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&stuffs).Error; err != nil {
		r.log.WithContext(ctx).Errorf("ListByUser: %v", err)
		return nil, err
	}

	result := make([]*biz.Stuff, 0, len(stuffs))
	for _, s := range stuffs {
		result = append(result, &biz.Stuff{
			ID:          s.ID,
			Name:        s.Name,
			Category:    s.Category,
			Price:       s.Price,
			Photos:      s.Photos,
			Publisher:   s.Publisher,
			Status:      s.Status,
			Condition:   s.Condition,
			Description: s.Description,
		})
	}
	return result, nil
}

func (r *stuffRepo) GetTotalByUser(ctx context.Context, userID int64) int64 {
	var total int64
	r.data.db.WithContext(ctx).Model(&Stuff{}).Where("publisher = ?", userID).Count(&total)
	return total
}
