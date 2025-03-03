<template>
    <el-dropdown trigger="click" @command="handleCommand">
        <div class="user-info">
            <el-avatar :size="32" :src="userStore.userInfo?.avatar || defaultAvatar" />
            <span class="username">{{ userStore.userInfo?.username }}</span>
        </div>

        <template #dropdown>
            <el-dropdown-menu class="custom-dropdown">
                <el-dropdown-item command="settings">
                    <el-icon>
                        <Setting />
                    </el-icon>
                    个人中心
                </el-dropdown-item>
                <el-dropdown-item command="logout" @click="handleLogout">
                    <el-icon>
                        <SwitchButton />
                    </el-icon>
                    退出登录
                </el-dropdown-item>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>

<script setup lang="ts">
import { User, Setting, SwitchButton } from '@element-plus/icons-vue'
import { useUserStore } from '@/store/modules/user'
import { useRouter } from 'vue-router'
import defaultAvatar from '@/assets/default-avatar.png'

const userStore = useUserStore()
const router = useRouter()

console.log(userStore.userInfo)

const handleCommand = async (command: string) => {
    switch (command) {
        case 'profile':
            router.push('/profile')
            break
        case 'settings':
            router.push('/personal')
            break
        case 'logout':
            try {
                userStore.clearLoginState()
                router.push('/')
            } catch (error) {
                console.error('退出登录失败：', error)
            }
            break
    }
}

const handleLogout = () => {
    userStore.clearLoginState()
    router.push('/login')
}
</script>

<style scoped lang="scss">
.user-info {
    display: flex;
    align-items: center;
    cursor: pointer;

    .username {
        margin-left: 8px;
        font-size: 14px;
    }
}
</style>

<style lang="scss">
.custom-dropdown {
    background-color: #213547 !important;
    border: none !important;

    .el-dropdown-menu__item {
        color: white !important;

        .el-icon {
            margin-right: 8px;
            color: white !important;
        }

        &:not(.el-dropdown-menu__item--divided):hover {
            background-color: #1976d2 !important;
        }

        &.el-dropdown-menu__item--divided {
            border-top-color: rgba(255, 255, 255, 0.1) !important;
        }
    }
}

.el-popper.is-light {
    border: none !important;

    .el-popper__arrow::before {
        background-color: #1867c0 !important;
        border: none !important;
    }
}
</style>