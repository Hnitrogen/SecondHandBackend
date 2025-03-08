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
    updateUserInfo(data: { id: number; name: string; address: string; phone: string }) {
        return http.put(`/user/${data.id}`, data)
    },

    // 更新头像
    updateUserAvatar(data: { id: number; avatar: string }) {
        return http.put(`/user/${data.id}/avatar`, data)
    }
}

