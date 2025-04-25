import { http } from '@/utils/request'


export const stuffApi = {
    // 发布闲置商品

    publishStuff(data: {
        name: string; price: number; category: string;
        description: string; photos: string, publisher: string; condition: string;
    }) {
        // Example:
        // {
        //     "name": "安提戈鲁思家族的笔记",
        //     "category": "神奇物品",
        //     "price": 3.22,
        //     "photos": "xx.png",
        //     "publisher": "1791422575@qq.com"
        // }
        return http.post('/stuff/create', data)
    },
    getStuffCategory() {
        return http.get('/stuff/category/list')
    },
    getStuffDetail(id: string) {
        return http.get(`/stuff/${id}`)
    },
    getStuffList(data: {
        category: string;
        page: number;
        pageSize: number;
    }) {
        return http.post('/stuff/category', data)
    },
    getStuffListAll(data: {
        category: string;
        page: number;
        pageSize: number;
    }) {
        return http.post('/stuff/all', data)
    },
    getUserStuffs(userId: number) {
        return http.get(`/stuff/user/${userId}?page=1&page_size=1000`)
    }
}
