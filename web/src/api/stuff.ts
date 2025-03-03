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
    }
}

