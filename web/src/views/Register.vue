<template>
    <div class="register-container">
        <el-card class="register-card">
            <template #header>
                <div class="card-header">
                    <img src="../assets/logo.png" alt="Logo" class="logo">
                    <h2>创建账号</h2>
                    <p>加入二手商城，开启您的购物之旅</p>
                </div>
            </template>
            
            <el-form 
                ref="registerFormRef"
                :model="registerForm"
                :rules="rules"
                label-position="top"
            >
                <el-form-item prop="name" label="用户名">
                    <el-input 
                        v-model="registerForm.name"
                        prefix-icon="el-icon-user"
                        placeholder="请输入用户名"
                    />
                </el-form-item>

                <el-form-item prop="email" label="邮箱">
                    <el-input 
                        v-model="registerForm.email"
                        prefix-icon="el-icon-message"
                        placeholder="请输入邮箱"
                    />
                </el-form-item>

                <el-form-item prop="password" label="密码">
                    <el-input 
                        v-model="registerForm.password"
                        type="password"
                        prefix-icon="el-icon-lock"
                        placeholder="请输入密码"
                        show-password
                    />
                </el-form-item>

                <el-form-item prop="confirmPassword" label="确认密码">
                    <el-input 
                        v-model="registerForm.confirmPassword"
                        type="password"
                        prefix-icon="el-icon-lock"
                        placeholder="请再次输入密码"
                        show-password
                    />
                </el-form-item>

                <el-button type="primary" class="register-button" @click="handleRegister" :loading="loading">
                    注册
                </el-button>

                <div class="login-link">
                    已有账号？ <router-link to="/login">立即登录</router-link>
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
const router = useRouter()
const loading = ref(false)
const registerFormRef = ref<FormInstance>()
const registerForm = reactive({
    name: '',
    email: '',
    password: '',
    confirmPassword: ''
})

const validatePass2 = (rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('请再次输入密码'))
    } else if (value !== registerForm.password) {
        callback(new Error('两次输入密码不一致!'))
    } else {
        callback()
    }
}

const rules = reactive<FormRules>({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 2, message: '用户名长度不能小于2位', trigger: 'blur' }
    ],
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 3, message: '密码长度不能小于3位', trigger: 'blur' }
    ],
    confirmPassword: [
        { required: true, message: '请再次输入密码', trigger: 'blur' },
        { validator: validatePass2, trigger: 'blur' }
    ]
})

const handleRegister = async () => {
    if (!registerFormRef.value) return
    
    await registerFormRef.value.validate((valid) => {
        if (valid) {
            userApi.register(registerForm).then((res: any) => {
                ElMessage.success('注册成功')
                router.push('/login')
          
            })
        }
    })
}
</script>

<style scoped>
.register-container {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, #e3f2fd 0%, #bbdefb 100%);
    padding: 20px;
}

.register-card {
    width: 100%;
    max-width: 400px;
}

.card-header {
    text-align: center;
    margin-bottom: 20px;
}

.logo {
    width: 80px;
    margin-bottom: 16px;
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

.register-button {
    width: 100%;
    margin-top: 16px;
    padding: 12px 0;
}

.login-link {
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