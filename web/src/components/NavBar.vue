<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import UserDropdown from './UserDropdown.vue'
import { useUserStore } from '@/store/modules/user'

const router = useRouter()
const categories = []
const searchQuery = ref('')
const userStore = useUserStore()

const goToLogin = () => {
    router.push('/login')
}


const handleCategoryClick = (category: string) => {
    if (category === '我的收藏') {
        router.push('/favorites')
    } else if (category === '分类') {
        router.push('/category')
    }
}

const handleSearch = () => {
    // 这里处理搜索逻辑
    console.log('搜索:', searchQuery.value)
    router.push(`/search?q=${searchQuery.value}`)
}
</script>

<template>
    <div class="nav-wrapper">
        <nav class="navbar">
            <div class="nav-left">
                <router-link to="/" class="logo">校园二手商城</router-link>
                <div class="nav-links">
                    <router-link to="/" class="nav-item">首页</router-link>
                    <a href="#" class="nav-item" v-for="category in categories" :key="category"
                        @click="handleCategoryClick(category)">
                        {{ category }}
                    </a>
                </div>
            </div>
            <div class="nav-right">
                <div class="search-box">
                    <input type="text" v-model="searchQuery" placeholder="搜索商品" @keyup.enter="handleSearch">
                    <button class="search-btn" @click="handleSearch">
                        <el-icon><Search /></el-icon>
                    </button>
                </div>
                <a @click="goToLogin" class="login-btn" v-if="!userStore.token">登录</a>
                <UserDropdown v-else />
            </div>
        </nav>
    </div>
</template>

<style scoped>
.nav-wrapper {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    background-color: var(--el-bg-color);
    display: flex;
    justify-content: center;
    box-shadow: var(--el-box-shadow-light);
    z-index: 1000;
    margin-bottom: 200px;
}

.navbar {
    max-width: var(--max-width);
    width: 100%;
    padding: 0.8rem 2rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.nav-left {
    display: flex;
    align-items: center;
    gap: 3rem;
}

.logo {
    color: var(--el-color-primary);
    font-size: 1.8rem;
    text-decoration: none;
    font-weight: bold;
    letter-spacing: 1px;
}

.nav-links {
    display: flex;
    gap: 2rem;
}

.nav-item {
    color: var(--el-text-color-regular);
    text-decoration: none;
    font-size: 1.1rem;
    transition: all 0.3s ease;
    padding: 0.5rem 0;
    position: relative;
}

.nav-item:hover {
    color: var(--el-color-primary);
}

.nav-item::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    width: 0;
    height: 2px;
    background-color: #3b82f6;
    transition: width 0.3s ease;
}

.nav-item:hover::after {
    width: 100%;
}

.nav-right {
    display: flex;
    align-items: center;
    gap: 1.5rem;
}

/* 搜索框样式 */
.search-box {
    display: flex;
    align-items: center;
    background-color: var(--el-bg-color-page);
    border-radius: var(--el-border-radius-base);
    padding: 0.3rem;
    border: 1px solid var(--el-border-color);
    transition: all 0.3s ease;
}

.search-box:focus-within {
    border-color: var(--el-color-primary);
    box-shadow: 0 0 0 2px var(--el-color-primary-light-9);
}

.search-box input {
    background: none;
    border: none;
    outline: none;
    color: #fff;
    padding: 0.5rem 1rem;
    width: 200px;
    font-size: 0.95rem;
}

.search-box input::placeholder {
    color: #64748b;
}

.search-btn {
    background: none;
    border: none;
    color: #64748b;
    padding: 0.5rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: color 0.3s ease;
}

.search-btn:hover {
    color: #3b82f6;
}

.login-btn {
    background-color: var(--el-color-primary);
    color: white;
    padding: 0.6rem 1.8rem;
    border-radius: var(--el-border-radius-base);
    text-decoration: none;
    font-size: 1rem;
    transition: all 0.3s ease;
    cursor: pointer;
}

.login-btn:hover {
    background-color: var(--el-color-primary-dark-2);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px var(--el-color-primary-light-5);
}

/* 响应式调整 */
@media (max-width: 1024px) {
    .navbar {
        max-width: 960px;
    }

    .nav-left {
        gap: 2rem;
    }

    .nav-links {
        gap: 1.5rem;
    }

    .search-box input {
        width: 150px;
    }
}

@media (max-width: 768px) {
    .navbar {
        max-width: 100%;
        padding: 0.8rem 1rem;
    }

    .nav-links {
        display: none;
        /* 在移动端可以改为下拉菜单 */
    }

    .search-box input {
        width: 120px;
    }
}
</style>