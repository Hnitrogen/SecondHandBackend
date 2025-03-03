package data

import (
	"context"
	"strconv"
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
		Id:   strconv.FormatInt(s.ID, 10),
		Name: s.Name,
	}
}
