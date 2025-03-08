<template>
  <Navbar />
  <el-card class="profile-container">
    <!-- 个人信息头部 -->
    <div class="profile-header">
      <div class="avatar-section">
        <el-upload
          class="avatar-uploader"
          :show-file-list="false"
          :auto-upload="false"
          :on-change="handleAvatarChange"
          accept="image/*"
        >
          <el-avatar
            :size="100"
            :src="userStore.userInfo.avatar"
            class="hover:opacity-80"
          >
            <el-icon><UserFilled /></el-icon>
          </el-avatar>
          <div class="avatar-hover-text">
            <el-icon><Edit /></el-icon>
            <span>修改头像</span>
          </div>
        </el-upload>
        <h2 class="welcome-text">欢迎回来, {{ userStore.userInfo?.name }}</h2>
      </div>
      
      <el-descriptions :column="2" border>
        <!-- <el-descriptions-item label="用户ID">{{ userStore.userInfo?.id }}</el-descriptions-item> -->
        <el-descriptions-item label="昵称">{{ userStore.userInfo?.name }}</el-descriptions-item>
        <el-descriptions-item label="电子邮件">{{ userStore.userInfo?.email }}</el-descriptions-item>
        <!-- <el-descriptions-item label="角色">
          {{ userStore.userInfo?.role === 'ADMIN' ? '管理员' : '用户' }}
        </el-descriptions-item> -->
        <el-descriptions-item label="手机号码">{{ userStore.userInfo?.phone }}</el-descriptions-item>
        <el-descriptions-item label="地址">{{ userStore.userInfo?.address }}</el-descriptions-item>
      </el-descriptions>
    </div>

    <!-- 功能卡片区域 -->
    <el-row :gutter="20" class="profile-cards">
      <el-col :span="6" v-for="(item, index) in menuItems" :key="index">
        <el-card class="profile-card" shadow="hover" @click="router.push(item.path)">
          <el-icon :size="24" class="card-icon">
            <component :is="item.icon" />
          </el-icon>
          <span class="card-text">{{ item.title }}</span>
        </el-card>
      </el-col>
    </el-row>

    <!-- 操作按钮 -->
    <div class="profile-actions">
      <el-button type="primary" @click="showEditModal = true">修改信息</el-button>
      <el-button type="warning" @click="showPasswordModal = true">修改密码</el-button>
      <el-button type="danger" @click="handleLogout">退出登录</el-button>
    </div>

    <!-- 编辑弹窗 -->
    <el-dialog v-model="showEditModal" title="修改信息" width="500px">
      <el-form :model="formData" label-width="100px">
        <!-- <el-form-item label="电子邮件">
          <el-input v-model="formData.email" />
        </el-form-item> -->
        <el-form-item label="昵称">
          <el-input v-model="formData.name" />
        </el-form-item>
        <el-form-item label="手机号码">
          <el-input v-model="formData.phone" />
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="formData.address" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showEditModal = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">保存</el-button>
      </template>
    </el-dialog>

    <!-- 修改密码弹窗 -->
    <el-dialog v-model="showPasswordModal" title="修改密码" width="500px">
      <el-form :model="passwordForm" label-width="100px">
        <el-form-item label="新密码">
          <el-input v-model="passwordForm.newPassword" type="password" show-password />
        </el-form-item>
        <el-form-item label="确认新密码">
          <el-input v-model="passwordForm.confirmPassword" type="password" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showPasswordModal = false">取消</el-button>
        <el-button type="primary" @click="handlePasswordSubmit">保存</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useUserStore } from '@/store/modules/user'
import { useRouter } from 'vue-router'
import { Edit, UserFilled } from '@element-plus/icons-vue'
import { userApi } from '@/api/user'
import { mediaApi } from '@/api/media'
import { ElMessage } from 'element-plus'
import Navbar from '../components/Navbar.vue'

const userStore = useUserStore()
const router = useRouter()
console.log(userStore.userInfo)

const showEditModal = ref(false)
const showPasswordModal = ref(false)

const formData = reactive({
  id: userStore.userInfo?.id || 0,
  email: userStore.userInfo?.email || '',
  name: userStore.userInfo?.name || '',
  phone: userStore.userInfo?.phone || '',
  address: userStore.userInfo?.address || ''
})

const passwordForm = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

// 添加菜单项配置
const menuItems = [
  { title: '我要卖闲置', path: '/sell-stuff', icon: 'Sell' },
  { title: '我的闲置', path: '/MyBooks', icon: 'Reading' },
  { title: '我的收藏', path: '/MyFavourite', icon: 'Star' },
  { title: '我的评价', path: '/MyRating', icon: 'ChatSquare' }
]

const handleLogout = () => {
  // 处理退出登录
  console.log("退出...")
  userStore.clearLoginState()
  router.push('/')
}

const handleSubmit = async () => {
  try {
    userApi.updateUserInfo(formData).then(() => {
      ElMessage.success("修改成功")
      showEditModal.value = false

      // migrate to userStore
      userStore.userInfo.phone = formData.phone
      userStore.userInfo.address = formData.address
      userStore.userInfo.name = formData.name
    })
  } catch (error) {
    ElMessage.error("修改失败")
  }
}

const handlePasswordSubmit = async () => {
  try {
    if (passwordForm.newPassword !== passwordForm.confirmPassword) {
      ElMessage.error("两次输入的新密码不一致")
      return
    }

    await userApi.updatePassword({
      password: passwordForm.newPassword
    })

    ElMessage.success("密码修改成功，请重新登录")
    showPasswordModal.value = false

    // 清除登录状态并跳转到登录页
    setTimeout(() => {
      handleLogout()
    }, 1500)
  } catch (error) {
    ElMessage.error("密码修改失败")
  }
}

const handleAvatarChange = async (file: any) => {
  try {
    if (!file.raw) return
    
    const formData = new FormData()
    formData.append('file', file.raw)
    formData.append('type',"avatar")
    
    // 上传头像
    const resp = await mediaApi.uploadImage(formData)
    
    userStore.userInfo.avatar = resp.preview  
    // 更新头像数据
    const userId:number = userStore.userInfo.id
    await userApi.updateUserAvatar({
      id: userId,
      avatar: "avatar/" + resp.filename
    })
    
    ElMessage.success('头像更新成功')
  } catch (error) {
    console.error('头像更新失败', error)
    ElMessage.error('头像更新失败')
  }
}

</script>

<style scoped>
.profile-container {
  padding-top: 100px;
  max-width: 1200px;
  margin: 2rem auto;
  background: var(--white);
}

.profile-header {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  margin-bottom: 2rem;
}

.avatar-section {
  display: flex;
  align-items: center;
  gap: 2rem;
  margin-bottom: 1rem;
}

.welcome-text {
  color: var(--text-green);
  margin: 0;
}

.profile-cards {
  margin: 2rem 0;
}

.profile-card {
  height: 150px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s;
}

.profile-card:hover {
  transform: translateY(-5px);
}

.card-icon {
  color: var(--primary-green);
  margin-bottom: 1rem;
}

.card-text {
  color: var(--text-green);
  font-size: 1.1rem;
}

.profile-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 2rem;
}

:deep(.el-descriptions) {
  padding: 1rem;
}

:deep(.el-descriptions__label) {
  color: var(--text-green);
}

.avatar-uploader {
  position: relative;
  display: inline-block;
  cursor: pointer;
}

.avatar-hover-text {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: rgba(0, 0, 0, 0.5);
  color: white;
  opacity: 0;
  transition: opacity 0.3s;
  border-radius: 50%;
}

.avatar-uploader:hover .avatar-hover-text {
  opacity: 1;
}

.avatar-hover-text .el-icon {
  font-size: 20px;
  margin-bottom: 4px;
}
</style>