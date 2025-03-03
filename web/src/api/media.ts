import { http } from '@/utils/request'

// 定义接口返回类型
interface MediaResult {
    filename: string
    url: string
}

export const mediaApi = {
    // 上传图片
    uploadImage(formData: FormData) {
        return http.post<MediaResult>('/media/upload', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
    },
}
