<template>
    <div class="product-detail">
        <Navbar />
        <div class="detail-container">
            <div class="product-content">
                <div class="left-section">
                    <el-carousel :interval="4000" type="card" height="400px">
                        <el-carousel-item v-for="image in product.images" :key="image">
                            <img :src="image" class="product-image" />
                        </el-carousel-item>
                    </el-carousel>
                </div>

                <div class="right-section">
                    <div class="product-header">
                        <h1 class="product-title">{{ product.title }}</h1>
                        <div class="product-price">¥{{ product.price }}</div>
                    </div>

                    <div class="seller-info-card">
                        <el-avatar :size="50" :src="product.sellerAvatar" />
                        <div class="seller-details">
                            <div class="seller-name">{{ product.sellerName }}</div>
                            <div class="seller-meta">
                                <span>信用度: {{ product.sellerCredit }}</span>
                                <span>发布时间: {{ product.publishTime }}</span>
                            </div>
                        </div>
                        <el-button type="primary" @click="followSeller">关注</el-button>
                    </div>

                    <div class="product-info">
                        <h3>商品信息</h3>
                        <el-descriptions :column="1" border>
                            <el-descriptions-item label="成色">{{ product.condition }}</el-descriptions-item>
                            <el-descriptions-item label="购买时间">{{ product.purchaseTime }}</el-descriptions-item>
                            <el-descriptions-item label="所在地">{{ product.location }}</el-descriptions-item>
                            <el-descriptions-item label="品牌">{{ product.brand }}</el-descriptions-item>
                        </el-descriptions>
                    </div>

                    <div class="product-description">
                        <h3>详细描述</h3>
                        <p>{{ product.description }}</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="action-bar">
            <div class="action-container">
                <div class="action-left">
                    <el-button type="default" icon="el-icon-star-off">收藏</el-button>
                    <el-button type="default" icon="el-icon-warning-outline">举报</el-button>
                </div>
                <div class="action-right">
                    <el-button type="primary" @click="openChat">联系卖家</el-button>
                    <el-button type="danger" @click="handleBuy">立即购买</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Navbar from '../components/Navbar.vue'

// 模拟商品数据
const product = ref({
    id: 1,
    title: '二手iPhone 12 Pro Max',
    price: 4999,
    images: [
        'path/to/image1.jpg',
        'path/to/image2.jpg',
        'path/to/image3.jpg'
    ],
    sellerName: '用户_54989905',
    sellerAvatar: 'path/to/avatar.jpg',
    sellerCredit: '优秀',
    publishTime: '2024-01-20',
    condition: '95新',
    purchaseTime: '2023-01',
    location: '厦门',
    brand: 'Apple',
    description: '去年1月购买，一直有套保护壳使用，无划痕，电池健康度90%，配件齐全...'
})

const followSeller = () => {
    // 实现关注卖家逻辑
}

const openChat = () => {
    // 实现打开聊天窗口逻辑
}

const handleBuy = () => {
    // 实现购买逻辑
}
</script>

<style scoped>
.product-detail {
    min-height: 100vh;
    background-color: #f5f5f5;
}

.detail-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
}

.product-content {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 40px;
    margin-bottom: 80px;
}

.left-section {
    background: white;
    padding: 20px;
    border-radius: 8px;
}

.product-image {
    width: 100%;
    height: 100%;
    object-fit: contain;
}

.right-section {
    background: white;
    padding: 20px;
    border-radius: 8px;
}

.product-header {
    margin-bottom: 20px;
}

.product-title {
    font-size: 24px;
    color: #333;
    margin-bottom: 10px;
}

.product-price {
    font-size: 32px;
    color: #ff4d4f;
    font-weight: bold;
}

.seller-info-card {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 15px;
    border: 1px solid #eee;
    border-radius: 8px;
    margin: 20px 0;
}

.seller-details {
    flex-grow: 1;
}

.seller-name {
    font-size: 16px;
    font-weight: bold;
    margin-bottom: 4px;
}

.seller-meta {
    font-size: 12px;
    color: #999;
    display: flex;
    gap: 10px;
}

.product-info {
    margin: 20px 0;
}

.product-description {
    margin: 20px 0;
}

.action-bar {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background: white;
    box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
    padding: 10px 0;
}

.action-container {
    max-width: 1200px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
}

.action-left {
    display: flex;
    gap: 10px;
}

.action-right {
    display: flex;
    gap: 10px;
}

@media (max-width: 768px) {
    .product-content {
        grid-template-columns: 1fr;
        gap: 20px;
    }
    
    .action-container {
        padding: 0 10px;
    }
}
</style> 