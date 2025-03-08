import { defineStore } from 'pinia'

interface UserState {
    token: string
    userInfo: {
        id: number
        name: string
        email: string
        avatar?: string
        address?: string
        phone?: string
    } | null
}

export const useUserStore = defineStore('user', {
    state: (): UserState => ({
        token: '',
        userInfo: null
    }),

    actions: {
        setLoginState(data: { id: number; token: string; username: string; email: string; avatar: string; address: string; phone: string; role: string }) {
            this.token = data.token
            this.userInfo = {
                id: data.id,
                name: data.username,
                email: data.email,
                avatar: data.avatar,
                address: data.address,
                phone: data.phone,
            }
        },
        clearLoginState() {
            this.token = ''
            this.userInfo = null
            localStorage.removeItem('token')
        },
        updateAvatar(avatar: string) {
            if (this.userInfo) {
                this.userInfo.avatar = avatar
            }
        }
    },

    persist: {
        key: 'user-store',
        storage: localStorage,
        paths: ['token', 'userInfo'] // 指定要持久化的字段
    }
}) 