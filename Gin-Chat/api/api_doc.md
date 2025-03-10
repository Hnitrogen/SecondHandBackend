# 聊天服务 API 文档

## 基础地址

所有API的基础地址为: `http://your-domain/api/v1/chat`

## WebSocket 接口

### 建立WebSocket连接

- **URL**: `/ws`
- **方法**: GET
- **参数**:
  - `userId`: 用户ID
  - `itemId`: 商品ID

- **说明**: 
  建立WebSocket连接后，服务器会发送一条欢迎消息。
  连接建立后，用户状态会被设置为在线。
  
- **WebSocket消息格式**:
  ```json
  {
    "type": "text|image|system|error|ping|pong",
    "data": "消息内容",
    "from": 123,
    "to": 456,
    "itemId": 789,
    "time": 1615423694
  }
  ```

- **WebSocket支持的消息类型**:
  - `text`: 文本消息
  - `system`: 系统消息
  - `error`: 错误消息
  - `ping`: 心跳检测请求
  - `pong`: 心跳检测响应

- **示例请求**:
  ```
  ws://your-domain/ws?userId=123&itemId=456
  ```

## HTTP 接口

### 发送消息

- **URL**: `/message`
- **方法**: POST
- **参数**:
  ```json
  {
    "sender_id": 123,
    "receiver_id": 456,
    "item_id": 789,
    "content": "你好，这个商品还有吗？"
  }
  ```

- **响应**:
  ```json
  {
    "message": "消息发送成功"
  }
  ```

### 获取消息历史

- **URL**: `/messages`
- **方法**: GET
- **参数**:
  - `userId`: 用户ID
  - `targetId`: 对话对象ID
  - `itemId`: 商品ID (可选，如不提供则获取所有会话消息)
  - `page`: 页码 (默认为1)
  - `pageSize`: 每页条数 (默认为20)

- **响应**:
  ```json
  {
    "messages": [
      {
        "id": 1,
        "sender_id": 123,
        "receiver_id": 456,
        "content": "你好，这个商品还有吗？",
        "item_id": 789,
        "status": 1,
        "created_at": "2023-03-10T10:30:45Z",
        "updated_at": "2023-03-10T10:30:45Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 20
  }
  ```

### 获取会话列表

- **URL**: `/conversations`
- **方法**: GET
- **参数**:
  - `userId`: 用户ID

- **响应**:
  ```json
  {
    "conversations": [
      {
        "id": 1,
        "user_id": 123,
        "target_user_id": 456,
        "item_id": 789,
        "last_message_id": 10,
        "unread_count": 0,
        "updated_at": "2023-03-10T10:30:45Z",
        "created_at": "2023-03-10T10:30:45Z"
      }
    ]
  }
  ```

### 获取用户在线状态

- **URL**: `/status`
- **方法**: GET
- **参数**:
  - `userId`: 用户ID

- **响应**:
  ```json
  {
    "userId": 123,
    "isOnline": true
  }
  ```

## 前端集成

### WebSocket 客户端示例 (Vue.js)

```typescript
// 创建 WebSocket 连接
function createWebSocket(userId: number, itemId: number) {
  const ws = new WebSocket(`ws://your-domain/ws?userId=${userId}&itemId=${itemId}`);
  
  ws.onopen = () => {
    console.log('WebSocket 连接已建立');
  };
  
  ws.onmessage = (event) => {
    const message = JSON.parse(event.data);
    console.log('收到消息', message);
    
    // 处理不同类型的消息
    switch (message.type) {
      case 'text':
        // 处理文本消息
        break;
      case 'system':
        // 处理系统消息
        break;
      case 'error':
        // 处理错误消息
        break;
      case 'pong':
        // 处理心跳响应
        break;
    }
  };
  
  ws.onclose = () => {
    console.log('WebSocket 连接已关闭');
  };
  
  // 发送心跳包
  setInterval(() => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: 'ping' }));
    }
  }, 30000); // 每30秒发送一次
  
  return ws;
}

// 发送消息
function sendMessage(ws: WebSocket, receiverId: number, itemId: number, content: string) {
  if (ws.readyState === WebSocket.OPEN) {
    ws.send(JSON.stringify({
      type: 'text',
      to: receiverId,
      itemId: itemId,
      data: content
    }));
  }
}
``` 