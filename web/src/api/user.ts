import { http } from '@/utils/request'

// 定义接口返回类型
interface LoginResult {
    token: string
    userInfo: {
        id: number
        username: string
        email: string
    }
}

interface RegisterResult {
    code: number
    reason: string
    message: string
    metadata: {}
}

// 用户相关接口
export const userApi = {
    // 登录
    login(data: { email: string; password: string }) {
        return http.post<LoginResult>('/user/login', data)
    },

    // 注册
    register(data: { name: string; password: string; email: string }) {
        return http.post<RegisterResult>('/user/register', data)
    },

    // 用户信息
    updateUserInfo(userId: number, data: { username: string; email: string; avatar: string }) {
        return http.put(`/users/${userId}`, data)
    }
}

