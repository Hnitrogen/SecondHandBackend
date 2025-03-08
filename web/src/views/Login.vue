<template>
    <div class="login-container">
        <el-card class="login-card">
            <template #header>
                <div class="card-header">
                    <img src="../assets/logo.png" alt="Logo" class="logo">
                    <h2>欢迎回来</h2>
                    <p>登录您的二手商城账号</p>
                </div>
            </template>
            
            <el-form 
                ref="loginFormRef"
                :model="loginForm"
                :rules="rules"
                label-position="top"
            >
                <el-form-item prop="email" label="邮箱">
                    <el-input 
                        v-model="loginForm.email"
                        prefix-icon="el-icon-message"
                        placeholder="请输入邮箱"
                    />
                </el-form-item>

                <el-form-item prop="password" label="密码">
                    <el-input 
                        v-model="loginForm.password"
                        type="password"
                        prefix-icon="el-icon-lock"
                        placeholder="请输入密码"
                        show-password
                    />
                </el-form-item>

                <div class="remember-forgot">
                    <el-link type="primary" :underline="false">忘记密码？</el-link>
                </div>

                <el-button type="primary" class="login-button" @click="handleLogin" :loading="loading">
                    登录
                </el-button>

                <div class="register-link">
                    还没有账号？ <router-link to="/register">立即注册</router-link>
                </div>
            </el-form>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import type { FormInstance, FormRules } from 'element-plus'
import { userApi } from '@/api/user'
import { useUserStore } from '@/store/modules/user'
const loading = ref(false)
const loginForm = reactive({
    email: '',
    password: '',
})

const userStore = useUserStore()
const router = useRouter()
const rules = reactive<FormRules>({
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { 
            pattern: /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/, 
            message: '请输入正确的邮箱格式', 
            trigger: 'blur' 
        }
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 3, message: '密码长度不能小于3位', trigger: 'blur' }
    ]
})

const loginFormRef = ref<FormInstance>()

const handleLogin = async () => {
    if (!loginFormRef.value) return
    
    try {
        await loginFormRef.value.validate((valid) => {
            if (valid) {
                loading.value = true
                userApi.login(loginForm).then((res: any) => {  
                    ElMessage.success('登录成功')
                   
                    localStorage.setItem('token', res.token)
                    // userStore
                    userStore.setLoginState(res)
                    console.log(res)
                    router.push('/') // 登录成功后跳转到首页
                })
            }
        })
    } catch (error) {
        ElMessage.error('请检查输入信息是否正确')
    } finally {
        loading.value = false
    }
}
</script>

<style scoped>
.login-container {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, #e3f2fd 0%, #bbdefb 100%);
    padding: 20px;
}

.login-card {
    width: 100%;
    max-width: 400px;
}

.card-header {
    text-align: center;
    margin-bottom: 20px;
}

.logo {
    width: 80px;
    height: 80px;
    margin-bottom: 16px;
    border-radius: 50%;
}

h2 {
    margin: 0;
    color: var(--text-color);
    font-size: 24px;
}

.card-header p {
    margin: 8px 0 0;
    color: #666;
}

.remember-forgot {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 16px 0;
}

.login-button {
    width: 100%;
    margin-top: 16px;
    padding: 12px 0;
}

.register-link {
    text-align: center;
    margin-top: 16px;
}

:deep(.el-input__wrapper) {
    box-shadow: 0 0 0 1px #dcdfe6 inset;
}

:deep(.el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px var(--primary-color) inset;
}
</style> 