package data

import (
	"context"
	"stuff/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   int64  `gorm:"primarykey" autoIncrement:"true"`
	Name string `gorm:"column:name;type:varchar(255)"`
}

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// List 实现 biz.CategoryRepo 接口
func (r *categoryRepo) List(ctx context.Context) ([]*biz.Category, error) {
	var categories []*Category
	if err := r.data.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}

	result := make([]*biz.Category, 0, len(categories))
	for _, c := range categories {
		result = append(result, &biz.Category{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return result, nil
}

// ConvertToProto 将数据库模型转换为proto消息
// func (c *Category) ConvertToProto() *v1_category.GetCategoryReply {
// 	return &v1_category.GetCategoryReply{
// 		Id:          strconv.FormatInt(c.ID, 10),
// 		Name:        c.Name,
// 		Description: c.Description,
// 	}
// }
