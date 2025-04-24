import { http } from '@/utils/request'

export const chatApi = {
    sendMessage(data: { sender_id: number, receiver_id: number, item_id: number, content: string }) {
        return http.post<any>('/chat/message', data)
    },
    getMessageList(data: { userId: number, targetId: number, itemId: number }) {
        return http.get<any>('/chat/messages', { params: data })
    },
    getConversationList(data: { userId: number }) {
        return http.get<any>('/chat/conversations', { params: data })
    }
}