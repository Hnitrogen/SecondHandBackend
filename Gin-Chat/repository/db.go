package repository

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"SecondHandBackend-master/Gin-Chat/config"
	"SecondHandBackend-master/Gin-Chat/model"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

// InitDB 初始化数据库连接
func InitDB() error {
	dbConfig := config.GlobalConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	// 自动迁移表结构
	err = DB.AutoMigrate(&model.Message{}, &model.Conversation{})
	if err != nil {
		return err
	}

	log.Println("数据库连接初始化成功")
	return nil
}

// InitRedis 初始化Redis连接
func InitRedis() error {
	redisConfig := config.GlobalConfig.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	log.Println("Redis连接初始化成功")
	return nil
}
