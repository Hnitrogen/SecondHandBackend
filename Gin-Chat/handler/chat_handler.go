package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"SecondHandBackend-master/Gin-Chat/model"
	"SecondHandBackend-master/Gin-Chat/service"
)

// 定义WebSocket升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有跨域请求，生产环境应当严格检查来源
		return true
	},
}

// WebSocketHandler WebSocket处理器
func WebSocketHandler(c *gin.Context) {
	// 获取参数
	userIdStr := c.Query("userId")
	itemIdStr := c.Query("itemId")

	// 参数校验
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil || userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品ID"})
		return
	}

	// 升级HTTP连接为WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("升级WebSocket连接失败: %v", err)
		return
	}

	// 注册连接
	wsConn := service.RegisterConnection(conn, uint(userId), uint(itemId))

	// 设置关闭处理
	defer service.UnregisterConnection(wsConn.ConnID)

	// 发送欢迎消息
	welcome := model.WebSocketMessage{
		Type: "system",
		Data: "连接成功",
		Time: time.Now().Unix(),
	}

	err = wsConn.SendMessage(welcome)
	if err != nil {
		log.Printf("发送欢迎消息失败: %v", err)
	}

	// 循环读取消息
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket错误: %v", err)
			}
			break
		}

		// 处理接收到的消息
		err = service.HandleWebSocketMessage(wsConn.ConnID, messageType, data)
		if err != nil {
			log.Printf("处理WebSocket消息失败: %v", err)
			// 发送错误信息给客户端
			errMsg := model.WebSocketMessage{
				Type: "error",
				Data: err.Error(),
				Time: time.Now().Unix(),
			}
			wsConn.SendMessage(errMsg)
		}
	}
}

// SendMessageHandler 发送消息处理器
func SendMessageHandler(c *gin.Context) {
	var req struct {
		SenderId   uint   `json:"sender_id" binding:"required"`
		ReceiverId uint   `json:"receiver_id" binding:"required"`
		ItemId     uint   `json:"item_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 发送消息
	err := service.SendChatMessage(req.SenderId, req.ReceiverId, req.ItemId, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "消息发送成功"})
}

// GetMessagesHandler 获取消息历史
func GetMessagesHandler(c *gin.Context) {
	// 获取参数
	userIdStr := c.Query("userId")
	targetIdStr := c.Query("targetId")
	itemIdStr := c.Query("itemId")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "20")

	// 参数校验
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil || userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	targetId, err := strconv.ParseUint(targetIdStr, 10, 64)
	if err != nil || targetId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的目标用户ID"})
		return
	}

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		itemId = 0 // 默认为0，表示获取所有会话消息
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 获取消息
	messages, total, err := service.GetMessages(uint(userId), uint(targetId), uint(itemId), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 标记消息为已读
	err = service.MarkMessagesAsRead(uint(targetId), uint(userId))
	if err != nil {
		log.Printf("标记消息为已读失败: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetConversationsHandler 获取会话列表
func GetConversationsHandler(c *gin.Context) {
	// 获取参数
	userIdStr := c.Query("userId")

	// 参数校验
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil || userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 获取会话列表
	conversations, err := service.GetConversations(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"conversations": conversations,
	})
}

// GetUserOnlineStatusHandler 获取用户在线状态
func GetUserOnlineStatusHandler(c *gin.Context) {
	// 获取参数
	userIdStr := c.Query("userId")

	// 参数校验
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil || userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 获取在线状态
	isOnline, err := service.GetUserOnlineStatus(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":   userId,
		"isOnline": isOnline,
	})
}
