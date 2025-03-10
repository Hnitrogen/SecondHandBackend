package model

import (
	"time"
)

// 消息状态
const (
	MessageStatusUnread = 0 // 未读
	MessageStatusRead   = 1 // 已读
)

// Message 聊天消息模型
type Message struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	SenderId   uint      `json:"sender_id" gorm:"not null;index"`   // 发送者ID
	ReceiverId uint      `json:"receiver_id" gorm:"not null;index"` // 接收者ID
	Content    string    `json:"content" gorm:"size:1000;not null"` // 消息内容
	ItemId     uint      `json:"item_id" gorm:"index"`              // 关联的商品ID
	Status     int       `json:"status" gorm:"default:0"`           // 消息状态: 0-未读, 1-已读
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Message) TableName() string {
	return "chat_messages"
}

// Conversation 会话模型 - 用于记录用户之间的会话
type Conversation struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserId        uint      `json:"user_id" gorm:"not null;index"`
	TargetUserId  uint      `json:"target_user_id" gorm:"not null;index"`
	ItemId        uint      `json:"item_id" gorm:"index"`
	LastMessageId uint      `json:"last_message_id"`
	UnreadCount   int       `json:"unread_count" gorm:"default:0"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

// TableName 指定表名
func (Conversation) TableName() string {
	return "chat_conversations"
}

// OnlineStatus 在线状态结构
type OnlineStatus struct {
	UserId     uint  `json:"user_id"`
	IsOnline   bool  `json:"is_online"`
	LastSeenAt int64 `json:"last_seen_at"`
}

// WebSocketMessage WebSocket消息格式
type WebSocketMessage struct {
	Type   string      `json:"type"`    // 消息类型: text, image, system
	Data   interface{} `json:"data"`    // 消息内容
	From   uint        `json:"from"`    // 发送者ID
	To     uint        `json:"to"`      // 接收者ID
	ItemId uint        `json:"item_id"` // 商品ID
	Time   int64       `json:"time"`    // 时间戳
}
