package main

import (
	"log"
	"path/filepath"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"SecondHandBackend-master/Gin-Chat/api"
	"SecondHandBackend-master/Gin-Chat/config"
	"SecondHandBackend-master/Gin-Chat/consul"
	"SecondHandBackend-master/Gin-Chat/repository"
)

func main() {
	// 加载配置
	configPath := filepath.Join("config", "config.json")
	err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接
	err = repository.InitDB()
	if err != nil {
		log.Fatalf("初始化数据库连接失败: %v", err)
	}

	// 初始化Redis连接
	err = repository.InitRedis()
	if err != nil {
		log.Fatalf("初始化Redis连接失败: %v", err)
	}

	// 初始化Consul连接
	consul.Register(config.ConsulAddr, 8084, "Gin-Chat", []string{"gin-chat"}, "Gin-Chat")

	// 创建Gin实例
	r := gin.Default()
	corsCFG := cors.DefaultConfig()
	corsCFG.AllowAllOrigins = true
	// 或者指定域名
	// config.AllowOrigins = []string{"http://localhost:8080"}
	r.Use(cors.New(corsCFG))

	// 设置路由
	api.SetupRouter(r)

	// 启动服务
	serverAddr := ":" + config.GlobalConfig.Server.Port
	log.Printf("聊天服务启动在 %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
