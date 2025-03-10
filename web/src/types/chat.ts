export interface User {
    id: string
    name: string
    avatar: string
    isOnline?: boolean
}

export interface Message {
    id: string
    userId: string
    content: string
    timestamp: number
    type: 'text' | 'image'
}

export interface ChatRoom {
    id: string
    users: User[]
    messages: Message[]
} 