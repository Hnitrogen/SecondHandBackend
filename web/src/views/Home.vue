<template>
    <div class="app">
        <Navbar />
        <div class="main-content">
            <div class="category-tabs">
                <el-tabs v-model="activeCategory" class="demo-tabs">
                    <el-tab-pane label="数码产品" name="digital" />
                    <el-tab-pane label="书籍" name="books" />
                    <el-tab-pane label="服装" name="clothing" />
                    <el-tab-pane label="其他" name="others" />
                </el-tabs>
            </div>

            <div class="products-grid">
                <el-card 
                    v-for="product in products" 
                    :key="product.id" 
                    class="product-card"
                    @click="navigateToDetail(product.id)"
                >
                    <img :src="product.image" class="product-image" />
                    <div class="product-info">
                        <div class="product-price">¥{{ product.price }}</div>
                        <div class="product-title">{{ product.title }}</div>
                        <div class="product-meta">
                            <span class="seller-info">
                                <el-avatar :size="20" :src="product.sellerAvatar" />
                                {{ product.sellerName }}
                            </span>
                            <span class="location">{{ product.location }}</span>
                        </div>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from '../components/Navbar.vue'

const router = useRouter()
const activeCategory = ref('all')

// 模拟数据
const products = ref([
    {
        id: 1,
        title: '二手iPhone 12 Pro Max',
        price: 4999,
        image: 'path/to/image.jpg',
        sellerName: '用户_54989905',
        sellerAvatar: 'path/to/avatar.jpg',
        location: '厦门',
    },
    {
        id: 2,
        title: '二手iPhone 12 Pro Max',
        price: 4999,
        image: 'path/to/image.jpg',
        sellerName: '用户_54989905',
        sellerAvatar: 'path/to/avatar.jpg',
        location: '厦门',
    },

])

const navigateToDetail = (productId: number) => {
    router.push(`/product/${productId}`)
}
</script>

<style scoped>
.app {
    min-height: 100vh;
    padding-top: 2px;
}

.main-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
}

.category-tabs {
    margin: 2rem 0;
    padding: 1rem 0;
    display: flex;
    gap: 2rem;
    border-bottom: 1px solid #2d2d2d;
}

.tab {
    color: #a8a8a8;
    font-size: 1.1rem;
    cursor: pointer;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    transition: all 0.3s ease;
}

.tab:hover {
    color: #FF7F50;
}

.tab.active {
    color: #FF7F50;
    background-color: rgba(255, 127, 80, 0.1);
}

.products-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
    gap: 20px;
    margin-top: 20px;
}

.product-card {
    cursor: pointer;
    transition: all 0.3s ease;
}

.product-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.product-image {
    width: 100%;
    height: 200px;
    object-fit: cover;
    border-radius: 4px;
}

.product-info {
    padding: 12px 0;
}

.product-price {
    color: #ff4d4f;
    font-size: 20px;
    font-weight: bold;
}

.product-title {
    margin: 8px 0;
    font-size: 14px;
    color: #333;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
}

.product-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 12px;
    color: #999;
}

.seller-info {
    display: flex;
    align-items: center;
    gap: 4px;
}

@media (max-width: 768px) {
    .main-content {
        padding: 10px;
    }
    
    .products-grid {
        grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
        gap: 10px;
    }
}
</style>