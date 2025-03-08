<template>
    <div class="app">
        <Navbar />
        <div class="main-content" style="margin-top: 80px;">
            <div class="category-tabs">
                <el-tabs v-for="category in categoryList" v-model="activeCategory" class="demo-tabs" @click="getProductsJudgement">
                    <el-tab-pane :label="category.name" :name="category.id" />
                </el-tabs>
            </div>

            <!-- Item组件 -->
            <div class="products-grid">
                <el-card 
                    v-for="product in products" 
                    :key="product.id" 
                    class="product-card"
                    @click="navigateToDetail(product.id)"
                >
                    <img :src="product.photos" class="product-image" />
                    <div class="product-info">
                        <div class="product-price">¥{{ product.price }}</div>
                        <div class="product-title">{{ product.name }}</div>
                        <div class="product-meta">
                            <span class="seller-info">
                                <el-avatar :size="20" :src="product.publisher.avatar" />
                                {{ product.publisher.name }}
                            </span>
                            <span style="color: darkblue; font-size: 12px;">{{ product.condition }}</span>
                        </div>
                    </div>
                </el-card>
            </div>
            
            <!-- 添加分页组件 -->
            <div class="pagination-container">
                <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :page-sizes="[10]"
                    :total="total"
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    layout="total, sizes, prev, pager, next"
                />
            </div>
        </div>
    </div>


</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from '../components/NavBar.vue'
import { stuffApi } from '@/api/stuff'
const router = useRouter()
const activeCategory = ref('all')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const categoryList = ref([
    {name:'全部',id:'all'}
])

// taglist
stuffApi.getStuffCategory().then((res:any) =>{
    categoryList.value = res.category
    categoryList.value.unshift({name:'全部',id:'all'})
})  


const products = ref([])

const navigateToDetail = (productId: number) => {
    router.push(`/product-detail/${productId}`)
}

// 处理每页显示数量变化
const handleSizeChange = (val: number) => {
    pageSize.value = val
    currentPage.value = 1 // 重置到第一页
    getProductsJudgement()
}

// 处理页码变化
const handleCurrentChange = (val: number) => {
    currentPage.value = val
    getProductsJudgement()
}

const getProductsJudgement = () => {
    if(activeCategory.value === 'all'){
        getProductsAll()
    }else{
        getProducts()
    }
}

// 修改获取商品列表函数
const getProducts = () => {
    stuffApi.getStuffList({
        category: activeCategory.value,
        page: currentPage.value,
        pageSize: pageSize.value
    }).then((res:any) =>{
        products.value = res.stuffs
        total.value = parseInt(res.total) // 添加总数
    })
}

const getProductsAll = () => {
    stuffApi.getStuffListAll({
        category: activeCategory.value,
        page: currentPage.value,
        pageSize: pageSize.value
    }).then((res:any) =>{
        products.value = res.stuffs
        total.value = parseInt(res.total) // 添加总数
    })
}

onMounted(() => {
    // getProducts()
    getProductsAll()
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

.category-tabs {
    margin: 1rem 0 2rem;
    padding: 0.5rem 0;
    display: flex;
    gap: 1.5rem;
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
    .pagination-container {
        margin-top: 20px;
    }
}

/* 添加新的tab样式 */
:deep(.el-tabs__nav-wrap::after) {
    display: none;  /* 移除 Element Plus 默认的底部线条 */
}

:deep(.el-tabs__item) {
    color: #909399;
    font-size: 1rem;
    padding: 0.5rem 1.5rem;
    transition: all 0.3s ease;
}

:deep(.el-tabs__item:hover) {
    color: #409EFF;
}

:deep(.el-tabs__item.is-active) {
    color: #409EFF;
    font-weight: 500;
}

/* 为激活的tab添加背景效果 */
:deep(.el-tabs__item.is-active) {
    background-color: rgba(64, 158, 255, 0.1);
    border-radius: 4px;
}

.pagination-container {
    margin-top: 30px;
    display: flex;
    justify-content: center;
}
</style>