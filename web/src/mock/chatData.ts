import type { User, Message, ChatRoom } from '../types/chat'

export const mockUsers: User[] = [
    {
        id: '1',
        name: 'BAUIMIN',
        avatar: '/avatars/1.jpg',
        isOnline: true
    },
    {
        id: '2',
        name: 'UnicornPhantom',
        avatar: '/avatars/2.jpg',
        isOnline: true
    },
    {
        id: '3',
        name: '清竹莫叶',
        avatar: '/avatars/3.jpg',
        isOnline: false
    },
    // ... 可以继续添加更多用户
]

export const mockMessages: Message[] = [
    {
        id: '1',
        userId: '1',
        content: '大家好！',
        timestamp: Date.now() - 3600000,
        type: 'text'
    },
    {
        id: '2',
        userId: '2',
        content: '你好啊！',
        timestamp: Date.now() - 3500000,
        type: 'text'
    },
    // ... 可以继续添加更多消息
]

export const mockChatRoom: ChatRoom = {
    id: 'room1',
    users: mockUsers,
    messages: mockMessages
} 