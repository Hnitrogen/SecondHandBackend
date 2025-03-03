import { http } from '@/utils/request'

// 定义接口返回类型
interface TagResult {
    id: string
    name: string
}

export const tagApi = {
    // 获取标签列表
    getTagList() {
        return http.get<TagResult>('/tags/')
    },

    // 添加标签
    addTag(data: { name: string }) {
        return http.post<TagResult>('/tags/', data)
    },

    // 更新标签
    updateTag(data: { id: string; name: string }) {
        return http.put<TagResult>('/tags/' + data.id, { name: data.name, description: "edited" })
    },

    // 删除标签
    deleteTag(data: { id: string }) {
        return http.delete<TagResult>('/tags/' + data.id)
    }
}
