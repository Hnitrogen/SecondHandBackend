package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"SecondHandBackend-master/Gin-Chat/model"
)

const (
	// Redis key patterns
	UserOnlineKeyPattern = "user:online:%d"  // 用户在线状态
	UserWebSocketKey     = "user:ws:conn:%d" // 用户WebSocket连接
)

// SaveMessage 保存消息到数据库
func SaveMessage(message *model.Message) error {
	return DB.Create(message).Error
}

// GetMessages 获取两个用户之间的消息
func GetMessages(userId, targetId, itemId uint, page, pageSize int) ([]model.Message, int64, error) {
	var messages []model.Message
	var total int64

	query := DB.Model(&model.Message{})

	// 构建查询条件：消息在两个用户之间发送，且关联特定商品
	if itemId > 0 {
		query = query.Where("item_id = ?", itemId)
	}

	query = query.Where(
		DB.Where("(sender_id = ? AND receiver_id = ?)", userId, targetId).
			Or("(sender_id = ? AND receiver_id = ?)", targetId, userId),
	)

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页获取消息
	offset := (page - 1) * pageSize
	err = query.Order("created_at ASC").
		Limit(pageSize).
		Offset(offset).
		Find(&messages).Error

	return messages, total, err
}

// MarkMessagesAsRead 将消息标记为已读
func MarkMessagesAsRead(senderId, receiverId uint) error {
	return DB.Model(&model.Message{}).
		Where("sender_id = ? AND receiver_id = ? AND status = ?",
			senderId, receiverId, model.MessageStatusUnread).
		Update("status", model.MessageStatusRead).Error
}

// GetUnreadMessageCount 获取未读消息数量
func GetUnreadMessageCount(userId uint) (int64, error) {
	var count int64
	err := DB.Model(&model.Message{}).
		Where("receiver_id = ? AND status = ?", userId, model.MessageStatusUnread).
		Count(&count).Error
	return count, err
}

// SaveOrUpdateConversation 保存或更新会话
func SaveOrUpdateConversation(userId, targetUserId, itemId, lastMessageId uint, unreadCount int) error {
	var conversation model.Conversation

	// 检查会话是否存在
	result := DB.Where("user_id = ? AND target_user_id = ? AND item_id = ?",
		userId, targetUserId, itemId).First(&conversation)

	if result.Error == nil {
		// 会话存在，更新
		conversation.LastMessageId = lastMessageId
		conversation.UnreadCount = unreadCount
		conversation.UpdatedAt = time.Now()
		return DB.Save(&conversation).Error
	} else {
		// 会话不存在，创建新会话
		conversation = model.Conversation{
			UserId:        userId,
			TargetUserId:  targetUserId,
			ItemId:        itemId,
			LastMessageId: lastMessageId,
			UnreadCount:   unreadCount,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		return DB.Create(&conversation).Error
	}
}

// GetConversations 获取用户的所有会话列表
func GetConversations(userId uint) ([]model.Conversation, error) {
	var conversations []model.Conversation
	err := DB.Where("user_id = ?", userId).
		Order("updated_at DESC").
		Find(&conversations).Error
	return conversations, err
}

// SetUserOnlineStatus 设置用户在线状态
func SetUserOnlineStatus(userId uint, isOnline bool) error {
	ctx := context.Background()
	key := fmt.Sprintf(UserOnlineKeyPattern, userId)

	status := model.OnlineStatus{
		UserId:     userId,
		IsOnline:   isOnline,
		LastSeenAt: time.Now().Unix(),
	}

	jsonData, err := json.Marshal(status)
	if err != nil {
		return err
	}

	// 设置用户状态，有效期24小时
	return Redis.Set(ctx, key, string(jsonData), 24*time.Hour).Err()
}

// GetUserOnlineStatus 获取用户在线状态
func GetUserOnlineStatus(userId uint) (model.OnlineStatus, error) {
	ctx := context.Background()
	key := fmt.Sprintf(UserOnlineKeyPattern, userId)

	var status model.OnlineStatus

	// 从Redis获取用户状态
	jsonData, err := Redis.Get(ctx, key).Result()
	if err != nil {
		// 设置默认离线状态
		status = model.OnlineStatus{
			UserId:     userId,
			IsOnline:   false,
			LastSeenAt: 0,
		}
		return status, nil
	}

	// 解析JSON数据
	if err := json.Unmarshal([]byte(jsonData), &status); err != nil {
		return status, err
	}

	return status, nil
}

// SetUserWebsocketConn 设置用户WebSocket连接标识
func SetUserWebsocketConn(userId uint, connId string) error {
	ctx := context.Background()
	key := fmt.Sprintf(UserWebSocketKey, userId)

	// 设置用户WebSocket连接标识，连接关闭时自动删除
	return Redis.Set(ctx, key, connId, 0).Err()
}

// RemoveUserWebsocketConn 移除用户WebSocket连接标识
func RemoveUserWebsocketConn(userId uint) error {
	ctx := context.Background()
	key := fmt.Sprintf(UserWebSocketKey, userId)

	return Redis.Del(ctx, key).Err()
}

// GetUserIdByWebsocketConn 通过WebSocket连接标识获取用户ID
func GetUserIdByWebsocketConn(connId string) (uint, error) {
	ctx := context.Background()

	// 搜索所有匹配的键
	pattern := fmt.Sprintf(UserWebSocketKey, "*")
	keys, err := Redis.Keys(ctx, pattern).Result()
	if err != nil {
		return 0, err
	}

	// 检查每个键的值
	for _, key := range keys {
		val, err := Redis.Get(ctx, key).Result()
		if err == nil && val == connId {
			// 从键中提取用户ID
			// 格式: user:ws:conn:{userId}
			var userId uint64
			fmt.Sscanf(key, UserWebSocketKey, &userId)
			return uint(userId), nil
		}
	}

	return 0, fmt.Errorf("未找到用户")
}

// GetLastMessage 获取最后一条消息
func GetLastMessage(userId, targetId, itemId uint) (model.Message, error) {
	var message model.Message
	err := DB.Where("sender_id = ? AND receiver_id = ? AND item_id = ?", userId, targetId, itemId).Order("created_at DESC").First(&message).Error
	if err != nil {
		// 存在无消息的情况
		return model.Message{}, err
	}

	return message, nil
}
