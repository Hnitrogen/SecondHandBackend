package data

import (
	"stuff/internal/conf"

	userpb "stuff/api/other/user/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewStuffRepo,
	NewCategoryRepo,
	// 明确指定 StuffRepo 的实现
	// wire.Bind(new(biz.StuffRepo), new(*stuffRepo)),
)

// Data .
type Data struct {
	uc userpb.UserClient
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userpb.UserClient) (*Data, func(), error) {
	// dsn := "user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	// 自动迁移表结构
	db.AutoMigrate(&Stuff{})
	db.AutoMigrate(&Category{})
	return &Data{
			uc: uc,
			db: db,
		}, func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}, nil
}
