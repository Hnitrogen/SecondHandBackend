<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ElInput, ElButton, ElAvatar, ElMessage } from 'element-plus'
import { format } from 'date-fns'
import { zhCN } from 'date-fns/locale'
import { useUserStore } from '@/store/modules/user'
import { chatApi } from '../api/chat'

import { useRoute } from 'vue-router'

const currentUserId = '1'
const route = useRoute()
const userStore = useUserStore()
const chatInfo = ref({})

const ws = ref<WebSocket | null>(null)

const itemId = ref<number>()
const userId = ref<number>() 

onMounted(() => {
    if (route.query.chatInfo) {
        chatInfo.value = JSON.parse(decodeURIComponent(route.query.chatInfo as string))

        ElMessage.success("商品ID： " + chatInfo.value.productId)
        ElMessage.success("用户ID： " + parseInt(userStore.userInfo.id))

        itemId.value = chatInfo.value.productId
        userId.value = userStore.userInfo.id

        // 建立WebSocket连接
        connectWebSocket(userId,itemId)
    }

    chatApi.getMessageList({
        userId: parseInt(userId.value || '1'),
        targetId: parseInt(userId.value || '1'),
        itemId: parseInt(itemId.value || '1')
    }).then((res: any) => {
        console.log(res)
    })
})


onUnmounted(() => {
    // 组件卸载时关闭WebSocket连接
    if (ws.value) {
        ws.value.close()
    }
})

interface User {
    id: string
    name: string
    avatar: string
    description?: string
    isOnline?: boolean
}

interface ChatMessage {
    id: string
    senderId: string
    content: string
    timestamp: number
    isAutoReply?: boolean
}

// 模拟用户列表
const users = ref<User[]>([
    {
        id: '2',
        name: '打野艾翁',
        avatar: '/avatars/1.jpg',
        description: '随参直播，需要代练院玩以及复盘可以私信我，蟹蟹~',
        isOnline: true
    }
])

// 当前选中的聊天对象
const selectedUser = ref<User>(users.value[0])

const messages = ref<ChatMessage[]>([
    {
        id: '1',
        senderId: '2',
        content: '欢迎关注比亚迪汽车海洋~喜欢我们的视频请一键三连！有任何问题的内容都可以私信小迪~',
        timestamp: new Date('2024-10-19 20:57').getTime(),
        isAutoReply: true
    },
    {
        id: '2',
        senderId: '1',
        content: '1111',
        timestamp: new Date('2024-10-19 14:51').getTime()
    }
])

const newMessage = ref('')
const messageRef = ref<HTMLDivElement>()

function formatTime(timestamp: number): string {
    return format(timestamp, 'yyyy年MM月dd日 HH:mm', { locale: zhCN })
}

// WebSocket连接函数
function connectWebSocket(userId: any, itemId: any) {
    const wsUrl = `ws://localhost:8084/ws?userId=${userId.value}&itemId=${itemId.value}`
    ws.value = new WebSocket(wsUrl)

    ws.value.onopen = () => {
        console.log('WebSocket连接已建立')
    }

    ws.value.onmessage = (event) => {
        const message: WSMessage = JSON.parse(event.data)
        // 将收到的消息添加到消息列表
        messages.value.push({
            id: Date.now().toString(),
            senderId: message.from.toString(),
            content: message.data,
            timestamp: message.time * 1000 // 转换为毫秒
        })

        // 滚动到最新消息
        setTimeout(() => {
            if (messageRef.value) {
                messageRef.value.scrollTop = messageRef.value.scrollHeight
            }
        }, 0)
    }

    ws.value.onerror = (error) => {
        console.error('WebSocket错误:', error)
    }

    ws.value.onclose = () => {
        console.log('WebSocket连接已关闭')
    }
}

// 修改发送消息函数
function sendMessage() {
    if (!newMessage.value.trim() || !ws.value) return

    chatApi.sendMessage({
        sender_id: parseInt(userId.value || '1'),
        receiver_id: parseInt(selectedUser.value.id),
        item_id: parseInt(itemId.value || '1'),
        content: newMessage.value
    }).then((res: any) => {
        newMessage.value = ''
        ElMessage.success(res.message)
    })
}

function selectUser(user: User) {
    selectedUser.value = user
    // 这里可以加载与该用户的聊天记录
}
</script>

<template>
    <!-- <Navbar /> -->
    <div class="chat-container">
        <!-- 左侧用户列表 -->
        <div class="user-list">
            <div class="user-list-header">
                <h2>近期消息</h2>
            </div>
            <div class="user-items">
                <div v-for="user in users" :key="user.id" class="user-item"
                    :class="{ 'user-item-active': selectedUser.id === user.id }" @click="selectUser(user)">
                    <el-avatar :size="40" :src="user.avatar" class="user-avatar" />
                    <div class="user-info">
                        <div class="user-name">{{ user.name }}</div>
                        <div class="user-description">{{ user.description }}</div>
                    </div>
                </div>
            </div>
        </div>

        <!-- 右侧聊天区域 -->
        <div class="chat-content">
            <!-- 消息列表 -->
            <div ref="messageRef" class="message-list">
                <div v-for="message in messages" :key="message.id" class="message-wrapper"
                    :class="{ 'message-self': message.senderId === currentUserId }">
                    <template v-if="message.senderId !== currentUserId">
                        <el-avatar :size="40" :src="chatInfo.sellerAvatar" class="message-avatar" />
                        <div class="message-content">
                            <div class="message-time">{{ formatTime(message.timestamp) }}</div>
                            <div class="message-bubble message-partner">
                                {{ message.content }}
                            </div>
                            <div v-if="message.isAutoReply" class="auto-reply-hint">
                                {{ seller }} 自动回复
                            </div>
                        </div>
                    </template>

                    <template v-else>
                        <div class="message-content">
                            <div class="message-time">{{ formatTime(message.timestamp) }}</div>
                            <div class="message-bubble message-self">
                                {{ message.content }}
                            </div>
                        </div>
                    </template>
                </div>
            </div>

            <!-- 输入区域 -->
            <div class="input-area">
                <el-input v-model="newMessage" placeholder="输入消息..." :rows="1" @keyup.enter="sendMessage">
                    <template #append>
                        <el-button type="primary" @click="sendMessage">
                            发送
                        </el-button>
                    </template>
                </el-input>
            </div>
        </div>
    </div>
</template>

<style scoped>
.chat-container {
    display: flex;
    height: 100vh;
    background-color: #f5f7fa;
}

.user-list {
    width: 280px;
    background-color: #fff;
    border-right: 1px solid #e6e6e6;
    display: flex;
    flex-direction: column;
}

.user-list-header {
    padding: 16px;
    border-bottom: 1px solid #e6e6e6;
}

.user-list-header h2 {
    margin: 0;
    font-size: 16px;
    color: #333;
}

.user-items {
    flex: 1;
    overflow-y: auto;
}

.user-item {
    display: flex;
    padding: 12px 16px;
    cursor: pointer;
    transition: background-color 0.2s;
    align-items: center;
    border-bottom: 1px solid #f0f0f0;
}

.user-item:hover {
    background-color: #f5f7fa;
}

.user-item-active {
    background-color: #e6f4ff;
}

.user-avatar {
    flex-shrink: 0;
}

.user-info {
    margin-left: 12px;
    overflow: hidden;
}

.user-name {
    font-size: 14px;
    color: #333;
    margin-bottom: 4px;
}

.user-description {
    font-size: 12px;
    color: #999;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.chat-content {
    flex: 1;
    display: flex;
    flex-direction: column;
}

.message-list {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
}

.message-wrapper {
    display: flex;
    margin-bottom: 20px;
    align-items: flex-start;
}

.message-wrapper.message-self {
    flex-direction: row-reverse;
}

.message-avatar {
    margin: 0 12px;
}

.message-content {
    max-width: 70%;
}

.message-time {
    font-size: 12px;
    color: #999;
    margin-bottom: 4px;
    text-align: center;
}

.message-bubble {
    padding: 10px 16px;
    border-radius: 4px;
    word-break: break-all;
}

.message-partner {
    background-color: #fff;
    margin-right: 48px;
}

.message-self {
    background-color: #95ccff;
    color: #fff;
    margin-left: 48px;
}

.auto-reply-hint {
    font-size: 12px;
    color: #999;
    margin-top: 4px;
}

.input-area {
    padding: 16px;
    background-color: #fff;
    border-top: 1px solid #e6e6e6;
}

:deep(.el-input-group__append) {
    padding: 0;
}

:deep(.el-button) {
    border: none;
    margin: 0;
    padding: 0 20px;
    height: 32px;
}
</style>