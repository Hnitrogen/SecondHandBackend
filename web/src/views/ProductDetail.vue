<template>
    <div class="product-detail">
        <Navbar />
        <div class="detail-container" style="margin-top: 100px;">
            <!-- 卖家信息区域 -->
            <div class="seller-section">
                <div class="seller-info">
                    <div class="seller-left">
                        <el-avatar :size="50" :src="productDetail.publisher.avatar" />
                        <div class="seller-details">
                            <div class="seller-name">{{ productDetail.publisher.name }}</div>
                            <div class="seller-location">{{ productDetail.publisher.address }}</div>
                        </div>
                    </div>
                    <el-button type="primary" plain @click="followSeller">+ 关注</el-button>
                </div>
            </div>


            <!-- 商品信息区域 -->
            <div class="product-content">
                <div class="left-section">
                    <el-carousel trigger="click" height="400px">
                        <img :src="productDetail.photos" class="carousel-image" />
                        <!-- <el-carousel-item v-for="(image, index) in productDetail.images" :key="index">
                            <img :src="image" class="carousel-image" />
                        </el-carousel-item> -->
                    </el-carousel>
                    
                    <div class="product-description">
                        <h3>商品详情</h3>
                        <div class="description-text">{{ productDetail.description }}</div>
                        <div class="image-list">
                            <img v-for="(image, index) in productDetail.detailImages" 
                                :key="index" 
                                :src="image" 
                                class="detail-image"
                            />
                        </div>
                    </div>
                </div>

                <div class="right-section">
                    <div class="price-section">
                        <div class="stuff-name">{{ productDetail.name }}</div>
                        <div class="current-price" style="margin-top: 10px;">¥{{ productDetail.price }}</div>
                    </div>

                    <!-- <div class="product-meta">
                        <div class="views">
                            <span>{{ productDetail.views }}人浏览</span>
                            <span>{{ productDetail.likes }}人想要</span>
                        </div>
                    </div>

                    <div class="product-info">
                        <el-descriptions :column="1" border>
                            <el-descriptions-item label="商品状态">{{ productDetail.condition }}</el-descriptions-item>
                            <el-descriptions-item label="购买渠道">{{ productDetail.purchaseChannel }}</el-descriptions-item>
                            <el-descriptions-item label="购买时间">{{ productDetail.purchaseTime }}</el-descriptions-item>
                            <el-descriptions-item label="品牌型号">{{ productDetail.brand }}</el-descriptions-item>
                            <el-descriptions-item label="发布时间">{{ productDetail.publishTime }}</el-descriptions-item>
                        </el-descriptions>
                    </div> -->
                </div>
            </div>
        </div>

        <!-- 底部操作栏 -->
        <div class="action-bar">
            <div class="action-container">
                <div class="action-left">
                    <el-button :icon="Star" @click="handleCollect">收藏</el-button>
                    <el-button :icon="Warning" @click="handleReport">举报</el-button>
                </div>
                <div class="action-right">
                    <el-button type="primary" @click="handleChat">联系卖家</el-button>
                    <el-button type="danger" @click="handleBuy">立即购买</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Star, Warning } from '@element-plus/icons-vue'
import Navbar from '../components/Navbar.vue'
import { stuffApi } from '@/api/stuff'
const route = useRoute()
const productId = route.params.id

const productDetail = ref({})
stuffApi.getStuffDetail(productId).then((res: any) => {
    productDetail.value = res
})

// Mock数据
// const productDetail = ref({
//     id: productId,
//     sellerName: '用户_54989905',
//     sellerAvatar: 'https://placeholder.com/50',
//     location: '厦门',
//     price: 11.99,
//     originalPrice: 99,
//     discount: 6.0,
//     views: 1700,
//     likes: 27,
//     images: [
//         'https://placeholder.com/800x400',
//         'https://placeholder.com/800x400',
//         'https://placeholder.com/800x400'
//     ],
//     detailImages: [
//         'https://placeholder.com/800x600',
//         'https://placeholder.com/800x600'
//     ],
//     description: '原价99元巨资购买，但被远程锁定。给钱就出，现在被锁无法正常充电，会的人可以买去破解',
//     condition: '95新',
//     purchaseChannel: '正品商城',
//     purchaseTime: '2023年12月',
//     brand: 'other/其他',
//     publishTime: '9小时前',
//     tags: ['小巧充电宝', '街电充电宝']
// })

// 处理关注卖家
const followSeller = () => {
    ElMessage.success('关注成功')
}

// 处理收藏
const handleCollect = () => {
    ElMessage.success('收藏成功')
}

// 处理举报
const handleReport = () => {
    ElMessage.info('举报已提交')
}

// 处理聊天
const handleChat = () => {
    ElMessage.info('正在打开聊天窗口')
}

// 处理购买
const handleBuy = () => {
    ElMessage.success('正在跳转到支付页面')
}

onMounted(() => {
    console.log('商品ID:', productId)
    // 这里可以调用API获取商品详情
})
</script>

<style scoped>
.product-detail {
    min-height: 100vh;
    background-color: #f5f7fa;
    padding-bottom: 60px;
}

.detail-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
}

.seller-section {
    background: white;
    padding: 20px;
    border-radius: 8px;
    margin-bottom: 20px;
}

.seller-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.seller-left {
    display: flex;
    align-items: center;
    gap: 15px;
}

.seller-details {
    display: flex;
    flex-direction: column;
    gap: 4px;
}

.seller-name {
    font-size: 16px;
    font-weight: bold;
}

.seller-location {
    font-size: 14px;
    color: #666;
}

.product-content {
    display: grid;
    grid-template-columns: 2fr 1fr;
    gap: 20px;
}

.left-section {
    background: white;
    border-radius: 8px;
    overflow: hidden;
}

.carousel-image {
    width: 100%;
    height: 100%;
    object-fit: contain;
}

.product-description {
    padding: 20px;
}

.description-text {
    margin: 15px 0;
    line-height: 1.6;
    color: #333;
}

.image-list {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.detail-image {
    width: 100%;
    border-radius: 4px;
}

.right-section {
    background: white;
    padding: 20px;
    border-radius: 8px;
    height: fit-content;
}

.price-section {
    margin-bottom: 20px;
}

.current-price {
    font-size: 28px;
    color: #ff4d4f;
    font-weight: bold;
}

.original-price {
    color: #999;
    text-decoration: line-through;
    font-size: 14px;
    margin: 5px 0;
}

.discount {
    display: inline-block;
    padding: 2px 8px;
    background-color: #fff1f0;
    color: #ff4d4f;
    border-radius: 4px;
    font-size: 12px;
}

.product-meta {
    margin: 15px 0;
    padding: 15px 0;
    border-top: 1px solid #eee;
    border-bottom: 1px solid #eee;
}

.views {
    display: flex;
    justify-content: space-between;
    color: #666;
    font-size: 14px;
}

.action-bar {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background: white;
    box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
    padding: 10px 0;
    z-index: 100;
}

.action-container {
    max-width: 1200px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
}

.action-left, .action-right {
    display: flex;
    gap: 10px;
}

@media (max-width: 768px) {
    .product-content {
        grid-template-columns: 1fr;
    }
    
    .detail-container {
        padding: 10px;
    }
    
    .action-container {
        padding: 0 10px;
    }
}
</style> 