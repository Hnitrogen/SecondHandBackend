<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ElInput, ElButton, ElAvatar, ElMessage } from 'element-plus'
import { format } from 'date-fns'
import { el, zhCN } from 'date-fns/locale'
import { useUserStore } from '@/store/modules/user'
import { chatApi } from '../api/chat'

import { useRoute } from 'vue-router'

const route = useRoute()
const userStore = useUserStore()
const chatInfo = ref({})
const currentMsgId = ref(0)     // 消息Id

const ws = ref<WebSocket | null>(null)

const itemId = ref<number>(0)
const userId = ref<number>(0) 
const sellerId = ref<number>(0)
const users = ref<User[]>([])
const ChatUserInfo = ref({})

const getMessages = (userId: number, sellerId: number, itemId: number) => {
    chatApi.getMessageList({
        userId: userId,
        targetId: sellerId,
        itemId: itemId
    }).then((res: any) => {
        console.log(res)
        messages.value = res.messages
        productInfo.value = res.productInfo
        ChatUserInfo.value = {
            "receiverInfo": res.receiverInfo,
            "senderInfo": res.senderInfo
        }

        // await 滚动到最新消息
        setTimeout(() => {
            if (messageRef.value) {
                messageRef.value.scrollTop = messageRef.value.scrollHeight
            }
        }, 0)
    })
}


onMounted(() => {
    if (route.query.chatInfo) {
        chatInfo.value = JSON.parse(decodeURIComponent(route.query.chatInfo as string))
        ElMessage.success("商品ID： " + chatInfo.value.productId)
        ElMessage.success("用户ID： " + parseInt(userStore.userInfo.id))
        ElMessage.success("卖家ID： " + chatInfo.value.sellerId)
        itemId.value = chatInfo.value.productId
        userId.value = userStore.userInfo.id
        sellerId.value = chatInfo.value.sellerId
        // 建立WebSocket连接
        connectWebSocket(userId,itemId)

        // 获取消息
        getMessages(userId.value, sellerId.value, itemId.value)
    }


    chatApi.getConversationList({
        userId: parseInt(userId.value || '1')
    }).then((res: any) => {
        console.log(res)
        users.value = res.conversations
    })
})


onUnmounted(() => {
    // 组件卸载时关闭WebSocket连接
    if (ws.value) {
        ws.value.close()
    }
})

interface User {
    id: number
    user_id: number
    target_user_id: number
    item_id: number
    last_message_id: number
    unread_count: number
    updated_at: string
    created_at: string
    user_info: {
        id: number
        username: string
        avatar: string
    }
    last_message: string
}

interface ChatMessage {
    id: number
    sender_id: number
    receiver_id: number
    content: string
    item_id: number
    status: number
    created_at: string
    updated_at: string
}

interface ProductInfo {
    name: string
    price: number
    photos: string
}

const productInfo = ref<ProductInfo>({
    name: '',
    price: 0,
    photos: ''
})

// 当前选中的聊天对象
const selectedUser = ref<User>(users.value[0])

const messages = ref<ChatMessage[]>([])

const newMessage = ref('')
const messageRef = ref<HTMLDivElement>()

function formatTime(timestamp: string): string {
    try {   
        return format(new Date(timestamp), 'yyyy年MM月dd日 HH:mm', { locale: zhCN })
    } catch (error) {
        // 1741788826
        console.log(timestamp)
        console.log(new Date(timestamp * 1000))
        return '未知时间'
    }
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
        console.log(message)
        ElMessage.success("收到消息： " + message.data)

        // 判断是否是当前选中的聊天对象
        if(message.item_id === itemId.value) {
            // 将收到的消息添加到消息列表
            messages.value.push({
                id: parseInt(Date.now().toString()),
                senderId: message.from.toString(),
                content: message.data,
                created_at: message.time * 1000
            })
            // 滚动到最新消息
            setTimeout(() => {
                if (messageRef.value) {
                    messageRef.value.scrollTop = messageRef.value.scrollHeight
                }
            }, 0)
        }else if(message.data === '连接成功') {
            return
        }
        else {
            ElMessage.warning("收到消息，但不是当前选中的聊天对象")

            // for users update unread_count
            users.value.forEach((user) => {
                if(user.item_id === message.item_id) {
                    user.unread_count++
                }
            })
        }
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
        receiver_id: parseInt(sellerId.value),
        item_id: parseInt(itemId.value || '1'),
        content: newMessage.value
    }).then((res: any) => {
        newMessage.value = ''
        ElMessage.success(res.message)
        // 更新消息列表
        getMessages(userId.value, sellerId.value, itemId.value)
    })
}

function selectMsg(user: User) {
    selectedUser.value = user
    // 这里可以加载与该用户的聊天记录
    console.log(user)   

    getMessages(user.user_id, user.target_user_id, user.item_id)
    // 更新选中消息Id
    currentMsgId.value = user.id
    // 更新选中用户
    sellerId.value = user.target_user_id
    itemId.value = user.item_id

    // ElMessage.success("切换消息： " + user.user_info.username + " 商品ID： " + itemId.value)

    if(user.unread_count > 0) {
        user.unread_count = 0
    }
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
                    @click="selectMsg(user)" :class="{ 'user-item-active': currentMsgId === user.id }">    
                    <!-- user.id 是消息id -->

                    <el-avatar :size="40" :src="user.user_info.avatar" class="user-avatar" />
                    <div class="user-info">
                        <div class="user-name">{{ user.user_info.username }} </div>
                        <div class="stuff-info" style="font-size: smaller; color: blue; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;">{{ user.stuff_info.name }}</div>
                        <div class="user-description">{{ user.last_message }}</div>
                    </div>
                    <div style="margin-left: auto;" v-if="user.unread_count > 0">
                        <el-button type="danger" style="text-align: right;"  round>{{ user.unread_count }}</el-button>
                    </div>
                </div>
            </div>
        </div>

        <!-- 右侧聊天区域 -->
        <div class="chat-content">
            <!-- 商品信息栏 -->
            <div class="product-info-bar" v-if="chatInfo.productId || productInfo.name"> 
                <div class="product-info">
                   
                    <img :src="productInfo.photos" class="product-image" alt="商品图片">
                    <div class="product-details">
                        <div class="product-name">{{ productInfo.name }}</div>
                        <div class="product-price">¥{{ productInfo.price }}</div>
                    </div>
                </div>
            </div>

            <!-- 消息列表 -->
            <div ref="messageRef" class="message-list">
                <div v-for="message in messages" :key="message.id" 
                    :class="[
                        'message-wrapper',
                        message.sender_id === parseInt(userId) ? 'message-wrapper-self' : 'message-wrapper-partner'
                    ]">
                    <!-- 对方的消息 -->
                    <template v-if="message.sender_id !== parseInt(userId)">
                        <el-avatar :size="40" :src="ChatUserInfo.senderInfo.avatar" class="message-avatar" />
                        <div class="message-content message-content-partner">
                            <div class="message-time">{{ formatTime(message.created_at) }}</div>
                            <div class="message-bubble message-bubble-partner">
                                {{ message.content }}
                            </div>
                        </div>
                    </template>

                    <!-- 自己的消息 -->
                    <template v-else>
                        <div class="message-content message-content-self">
                            <div class="message-time">{{ formatTime(message.created_at) }}</div>
                            <div class="message-bubble message-bubble-self">
                                {{ message.content }}
                            </div>
                        </div>
                        <el-avatar :size="40" :src="ChatUserInfo.receiverInfo.avatar" class="message-avatar" />
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

.product-info-bar {
    padding: 12px 16px;
    background-color: #fff;
    border-bottom: 1px solid #e6e6e6;
}

.product-info {
    display: flex;
    align-items: center;
    gap: 12px;
}

.product-image {
    width: 48px;
    height: 48px;
    object-fit: cover;
    border-radius: 4px;
}

.product-details {
    flex: 1;
}

.product-name {
    font-size: 14px;
    color: #333;
    margin-bottom: 4px;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
}

.product-price {
    font-size: 16px;
    color: #ff4d4f;
    font-weight: 500;
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

.message-wrapper-partner {
    justify-content: flex-start;
}

.message-wrapper-self {
    justify-content: flex-end;
}

.message-avatar {
    margin: 0 12px;
}

.message-content {
    max-width: 70%;
    display: flex;
    flex-direction: column;
}

.message-content-partner {
    margin-right: 48px;
}

.message-content-self {
    margin-left: 48px;
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
    display: inline-block;
    max-width: 100%;
}

.message-bubble-partner {
    background-color: #fff;
    align-self: flex-start;
}

.message-bubble-self {
    background-color: #95ccff;
    color: #fff;
    align-self: flex-end;
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