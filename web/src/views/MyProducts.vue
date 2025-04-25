<template>
    <div class="app">
        <Navbar />
        <div class="main-content" style="margin-top: 80px;">
            <h2>我的商品</h2>
            <!-- Loading state -->
            <div v-if="loading" class="loading-state">
                <el-skeleton :rows="3" animated />
            </div>

            <!-- Empty state -->
            <div v-else-if="products.length === 0" class="empty-state">
                <el-empty description="您还没有发布任何商品">
                    <el-button type="primary" @click="router.push('/sell-stuff')">发布商品</el-button>
                </el-empty>
            </div>

            <!-- Products list -->
            <div v-else class="products-list">
                <el-card v-for="product in products" :key="product.id" class="product-card"
                    @click="navigateToDetail(product.id)">
                    <div class="product-content">
                        <img :src="product.photos" class="product-image" />
                        <div class="product-info">
                            <div class="product-title">{{ product.name }}</div>
                            <div class="product-price">
                                ¥{{ product.price }}
                                <span class="condition">{{ product.condition }}</span>
                            </div>
                            <div class="product-meta">
                                <!-- <span class="upload-date">上传时间：{{ product.create_time }}</span> -->
                            </div>
                        </div>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/modules/user'
import { stuffApi } from '@/api/stuff'
import Navbar from '../components/NavBar.vue'

const router = useRouter()
const userStore = useUserStore()
const products = ref([])
const loading = ref(true)

const navigateToDetail = (productId: number) => {
    router.push(`/product-detail/${productId}`)
}

// const userId = userStore.userInfo?.id

const fetchUserProducts = async () => {
    try {
        loading.value = true
        const userId = userStore.userInfo?.id

        const response = await stuffApi.getUserStuffs(userId)
        products.value = response.stuffs || []
    } catch (error) {
        console.error('Failed to fetch user products:', error)
        products.value = []
    } finally {
        loading.value = false
    }
}

onMounted(() => {
    fetchUserProducts()
})
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

h2 {
    color: var(--text-color);
    margin-bottom: 24px;
}

.products-list {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.product-card {
    width: 100%;
    cursor: pointer;
    transition: all 0.3s ease;
}

.product-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.product-content {
    display: flex;
    gap: 24px;
    align-items: center;
    padding: 16px;
}

.product-image {
    width: 180px;
    height: 180px;
    object-fit: cover;
    border-radius: 8px;
}

.product-info {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.product-title {
    font-size: 22px;
    color: #333;
    font-weight: 600;
    line-height: 1.4;
}

.product-price {
    color: #ff4d4f;
    font-size: 28px;
    font-weight: bold;
    display: flex;
    align-items: center;
    gap: 16px;
}

.product-price .condition {
    color: var(--primary-color);
    font-size: 16px;
    font-weight: 500;
    background: rgba(24, 144, 255, 0.1);
    padding: 4px 12px;
    border-radius: 4px;
}

.product-meta {
    margin-top: 8px;
}

.upload-date {
    color: #999;
    font-size: 14px;
}

.loading-state {
    padding: 20px;
}

.empty-state {
    padding: 40px;
    text-align: center;
}

@media (max-width: 768px) {
    .main-content {
        padding: 10px;
    }

    .product-content {
        flex-direction: column;
        gap: 12px;
    }

    .product-image {
        width: 100%;
        height: 200px;
    }

    .product-info {
        width: 100%;
        text-align: center;
    }

    .product-meta {
        flex-direction: column;
        gap: 8px;
    }
}
</style>
