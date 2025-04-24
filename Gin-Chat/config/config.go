package config

import (
	"encoding/json"
	"os"
)

// Config 配置结构
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Redis    RedisConfig    `json:"redis"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port              string `json:"port"`
	WebsocketEndpoint string `json:"websocket_endpoint"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// GlobalConfig 全局配置实例
var GlobalConfig Config

// LoadConfig 从文件加载配置
func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&GlobalConfig)
}

const (
	ConsulAddr   = "127.0.0.1:8500"
	ConsulHealth = "http://127.0.0.1:8084/health"
)
