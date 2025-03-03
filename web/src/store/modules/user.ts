import { defineStore } from 'pinia'

interface UserState {
    token: string
    userInfo: {
        id: number
        username: string
        email: string
        avatar?: string
        role: 'admin' | 'author' | 'reader'
    } | null
}

export const useUserStore = defineStore('user', {
    state: (): UserState => ({
        token: '',
        userInfo: null
    }),

    actions: {
        setLoginState(data: { token: string; username: string; email: string; avatar: string; role: string }) {
            this.token = data.token
            this.userInfo = {
                id: 1,
                username: data.username,
                email: data.email,
                avatar: data.avatar,
                role: data.role as 'admin' | 'author' | 'reader'
            }
        },
        clearLoginState() {
            this.token = ''
            this.userInfo = null
            localStorage.removeItem('token')
        }
    },

    persist: {
        key: 'user-store',
        storage: localStorage,
        paths: ['token', 'userInfo'] // 指定要持久化的字段
    }
}) 