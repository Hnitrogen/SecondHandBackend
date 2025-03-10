package service

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"SecondHandBackend-master/Gin-Chat/model"
	"SecondHandBackend-master/Gin-Chat/repository"
)

// WebSocketConnection 用于存储和管理WebSocket连接
type WebSocketConnection struct {
	Conn     *websocket.Conn
	UserId   uint
	ConnID   string
	ItemId   uint
	SendMu   sync.Mutex
	IsActive bool
}

var (
	// WebSocketConnections 存储所有活跃的WebSocket连接
	WebSocketConnections = make(map[string]*WebSocketConnection)
	// WebSocketMutex 用于同步对WebSocketConnections的访问
	WebSocketMutex = sync.RWMutex{}
)

// 注册新的WebSocket连接
func RegisterConnection(conn *websocket.Conn, userId, itemId uint) *WebSocketConnection {
	// 生成唯一连接ID
	connID := fmt.Sprintf("%d-%d-%d", userId, itemId, time.Now().UnixNano())

	wsConn := &WebSocketConnection{
		Conn:     conn,
		UserId:   userId,
		ConnID:   connID,
		ItemId:   itemId,
		IsActive: true,
	}

	// 存储连接
	WebSocketMutex.Lock()
	WebSocketConnections[connID] = wsConn
	WebSocketMutex.Unlock()

	// 在Redis中设置用户的WebSocket连接
	err := repository.SetUserWebsocketConn(userId, connID)
	if err != nil {
		log.Printf("设置用户WebSocket连接失败: %v", err)
	}

	// 设置用户在线状态
	err = repository.SetUserOnlineStatus(userId, true)
	if err != nil {
		log.Printf("设置用户在线状态失败: %v", err)
	}

	return wsConn
}

// 注销WebSocket连接
func UnregisterConnection(connID string) {
	WebSocketMutex.Lock()
	defer WebSocketMutex.Unlock()

	if wsConn, exists := WebSocketConnections[connID]; exists {
		// 从Redis中移除WebSocket连接
		err := repository.RemoveUserWebsocketConn(wsConn.UserId)
		if err != nil {
			log.Printf("移除用户WebSocket连接失败: %v", err)
		}

		// 设置用户离线状态
		err = repository.SetUserOnlineStatus(wsConn.UserId, false)
		if err != nil {
			log.Printf("设置用户离线状态失败: %v", err)
		}

		// 关闭连接
		wsConn.IsActive = false
		wsConn.Conn.Close()

		// 从连接池中删除
		delete(WebSocketConnections, connID)
	}
}

// 发送消息到WebSocket连接
func (wsConn *WebSocketConnection) SendMessage(msg interface{}) error {
	wsConn.SendMu.Lock()
	defer wsConn.SendMu.Unlock()

	if !wsConn.IsActive {
		return fmt.Errorf("连接已关闭")
	}

	return wsConn.Conn.WriteJSON(msg)
}

// 获取用户的WebSocket连接
func GetUserConnection(userId uint) *WebSocketConnection {
	WebSocketMutex.RLock()
	defer WebSocketMutex.RUnlock()

	for _, conn := range WebSocketConnections {
		if conn.UserId == userId && conn.IsActive {
			return conn
		}
	}

	return nil
}

// SendChatMessage 发送聊天消息
func SendChatMessage(senderId, receiverId, itemId uint, content string) error {
	// 创建消息
	message := &model.Message{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Content:    content,
		ItemId:     itemId,
		Status:     model.MessageStatusUnread,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// 保存消息到数据库
	err := repository.SaveMessage(message)
	if err != nil {
		return err
	}

	// 获取接收者在线状态
	receiverStatus, err := repository.GetUserOnlineStatus(receiverId)
	if err != nil {
		log.Printf("获取接收者状态失败: %v", err)
	}

	// 构建WebSocket消息
	wsMessage := model.WebSocketMessage{
		Type:   "text",
		Data:   content,
		From:   senderId,
		To:     receiverId,
		ItemId: itemId,
		Time:   time.Now().Unix(),
	}

	// 如果接收者在线，通过WebSocket发送
	if receiverStatus.IsOnline {
		// 获取接收者的WebSocket连接
		receiverConn := GetUserConnection(receiverId)
		if receiverConn != nil {
			// 通过WebSocket发送消息
			err = receiverConn.SendMessage(wsMessage)
			if err != nil {
				log.Printf("通过WebSocket发送消息失败: %v", err)
				// 发送失败不影响后续逻辑
			} else {
				// 如果发送成功，标记消息为已读
				err = repository.MarkMessagesAsRead(senderId, receiverId)
				if err != nil {
					log.Printf("标记消息为已读失败: %v", err)
				}
			}
		}
	}

	// 更新发送者的会话
	err = repository.SaveOrUpdateConversation(senderId, receiverId, itemId, message.ID, 0)
	if err != nil {
		log.Printf("更新发送者会话失败: %v", err)
	}

	// 更新接收者的会话
	// 获取接收者的未读消息数量
	unreadCount, err := repository.GetUnreadMessageCount(receiverId)
	if err != nil {
		log.Printf("获取未读消息数量失败: %v", err)
		unreadCount = 0
	}

	err = repository.SaveOrUpdateConversation(receiverId, senderId, itemId, message.ID, int(unreadCount))
	if err != nil {
		log.Printf("更新接收者会话失败: %v", err)
	}

	return nil
}

// GetMessages 获取消息历史
func GetMessages(userId, targetId, itemId uint, page, pageSize int) ([]model.Message, int64, error) {
	return repository.GetMessages(userId, targetId, itemId, page, pageSize)
}

// GetConversations 获取会话列表
func GetConversations(userId uint) ([]model.Conversation, error) {
	return repository.GetConversations(userId)
}

// MarkMessagesAsRead 标记消息为已读
func MarkMessagesAsRead(senderId, receiverId uint) error {
	return repository.MarkMessagesAsRead(senderId, receiverId)
}

// GetUserOnlineStatus 获取用户在线状态
func GetUserOnlineStatus(userId uint) (bool, error) {
	status, err := repository.GetUserOnlineStatus(userId)
	if err != nil {
		return false, err
	}
	return status.IsOnline, nil
}

// HandleWebSocketMessage 处理从WebSocket接收到的消息
func HandleWebSocketMessage(connID string, messageType int, data []byte) error {
	WebSocketMutex.RLock()
	wsConn, exists := WebSocketConnections[connID]
	WebSocketMutex.RUnlock()

	if !exists || !wsConn.IsActive {
		return fmt.Errorf("连接不存在或已关闭")
	}

	// 解析收到的消息
	var wsMessage model.WebSocketMessage
	err := json.Unmarshal(data, &wsMessage)
	if err != nil {
		return err
	}

	// 根据消息类型处理
	switch wsMessage.Type {
	case "text":
		// 发送文本消息
		content, ok := wsMessage.Data.(string)
		if !ok {
			return fmt.Errorf("无效的消息内容")
		}

		// 发送消息
		return SendChatMessage(wsConn.UserId, wsMessage.To, wsMessage.ItemId, content)

	case "ping":
		// 心跳消息，回复pong
		return wsConn.SendMessage(map[string]string{"type": "pong"})

	default:
		return fmt.Errorf("不支持的消息类型")
	}
}
