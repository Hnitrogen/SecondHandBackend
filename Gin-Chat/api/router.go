package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SecondHandBackend-master/Gin-Chat/handler"
)

// SetupRouter 注册路由
func SetupRouter(r *gin.Engine) {
	// API 接口
	v1 := r.Group("/v1/chat")
	{
		// 发送消息
		v1.POST("/message", handler.SendMessageHandler)

		// 获取消息历史
		v1.GET("/messages", handler.GetMessagesHandler)

		// 获取会话列表
		v1.GET("/conversations", handler.GetConversationsHandler)

		// 获取用户在线状态
		v1.GET("/status", handler.GetUserOnlineStatusHandler)
	}

	// WebSocket 接口
	r.GET("/ws", handler.WebSocketHandler)

	// Consul健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

}
