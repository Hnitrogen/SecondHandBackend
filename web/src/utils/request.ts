import axios from 'axios'
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { useUserStore } from '@/store/modules/user'
import { ElMessage } from 'element-plus'

// 创建 axios 实例
const service: AxiosInstance = axios.create({
    baseURL: 'http://localhost/',
    timeout: 15000, // 请求超时时间
    headers: {
        'Content-Type': 'application/json;charset=utf-8',
    }
})

// 请求拦截器
service.interceptors.request.use(
    (config) => {
        const userStore = useUserStore()
        if (userStore.token) {
            config.headers.Authorization = `${userStore.token}`
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

// 响应拦截器
service.interceptors.response.use(
    (response: AxiosResponse) => {
        const res = response.data

        // 这里可以根据后端的响应结构进行调整
        if (response.status >= 200 && response.status < 300) {
            return res
        } else {
            // 处理其他状态码
            console.log(res)
            handleErrorResponse(res)
            return Promise.reject(new Error(res.message || '未知错误'))
        }
    },
    (error) => {
        console.log('响应错误：', error)
        const { response } = error
        if (response) {
            handleErrorResponse(response.data)
        }
        return Promise.reject(error)
    }
)

// 错误处理函数
function handleErrorResponse(res: any) {
    // 处理特定错误码
    switch (res.code) {
        case 401:
            // token 过期或未登录
            localStorage.removeItem('token')
            window.location.href = '/login'
            break
        case 403:
            // 无权限
            ElMessage.error('无权限访问')
            break
        default:
            ElMessage.error(res.message || '未知错误')
            console.log(res.message || '未知错误')
    }
}

// 封装请求方法
export const http = {
    get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return service.get(url, config)
    },

    post<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
        return service.post(url, data, config)
    },

    put<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
        return service.put(url, data, config)
    },

    delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return service.delete(url, config)
    }
}

export default service
