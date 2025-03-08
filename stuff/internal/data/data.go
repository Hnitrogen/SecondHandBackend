package data

import (
	"context"
	"stuff/internal/conf"
	"time"

	userpb "stuff/api/other/user/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewStuffRepo,
	NewCategoryRepo,
)

// Data .
type Data struct {
	uc userpb.UserClient
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userpb.UserClient) (*Data, func(), error) {
	dsn := c.Database.Source

	// 创建一个适配器，将 Kratos 的 log.Logger 适配到 GORM 的日志接口
	gormLog := NewGormLogger(logger)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLog,
	})

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

// GormLogger 是一个适配器，将 Kratos 的 log.Logger 适配到 GORM 的日志接口
type GormLogger struct {
	logger *log.Helper
}

// NewGormLogger 创建一个新的 GormLogger 实例
func NewGormLogger(logger log.Logger) gormLogger.Interface {
	return &GormLogger{
		logger: log.NewHelper(logger),
	}
}

// LogMode 设置日志模式
func (l *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

// Info 打印信息
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Infof(msg, data...)
}

// Warn 打印警告
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warnf(msg, data...)
}

// Error 打印错误
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Errorf(msg, data...)
}

// Trace 打印 SQL 语句
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	l.logger.Infof("SQL: %s, Rows: %d, Error: %v", sql, rows, err)
}
