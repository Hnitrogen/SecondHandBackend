<template>
  <el-card class="sell-container">
    <template #header>
      <div class="card-header">
        <h2>发布闲置商品</h2>
        <p class="subtitle">填写商品信息，让它找到新主人吧</p>
      </div>
    </template>

    <el-form ref="formRef" :model="itemForm" :rules="rules" label-position="top">
      <el-row :gutter="20">
        <!-- 左侧表单 -->
        <el-col :span="14">
          <el-form-item prop="name" label="商品名称">
            <el-input v-model="itemForm.name" placeholder="请输入商品名称" />
          </el-form-item>

          <el-form-item prop="category" label="商品类别">
            <el-select v-model="itemForm.category" placeholder="请选择商品类别" class="w-full">
              <el-option v-for="category in categoryList" :key="category.id" :label="category.name" :value="category.id" />
            </el-select>
          </el-form-item>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item prop="price" label="价格">
                <el-input-number v-model="itemForm.price" :min="0" :precision="2" class="w-full" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item prop="condition" label="商品成色">
                <el-select v-model="itemForm.condition" placeholder="选择成色" class="w-full">
                  <el-option v-for="condition in conditionList" :key="condition.id" :label="condition.name"
                    :value="condition.name" />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item prop="description" label="商品描述">
            <el-input type="textarea" v-model="itemForm.description" rows="4" placeholder="描述商品的具体情况，例如：使用时长、磨损程度、附带配件等" />
          </el-form-item>
        </el-col>

        <!-- 右侧图片上传 -->
        <el-col :span="10">
          <el-form-item prop="imageUrl" label="商品图片">
            <div class="upload-container">
              <el-upload class="image-uploader" :show-file-list="false" :before-upload="beforeUpload"
                :http-request="customUpload">
                <div v-if="itemForm.imageShow" class="image-preview">
                  <img :src="itemForm.imageShow" class="preview-img" />
                  <div class="image-actions">
                    <el-button type="primary" size="small" @click.stop="handleReupload">重新上传</el-button>
                  </div>
                </div>
                <div v-else class="upload-placeholder">
                  <el-icon class="upload-icon"><Plus /></el-icon>
                  <div class="upload-text">点击上传商品图片</div>
                  <div class="upload-tip">建议尺寸 800x800px，最大 5MB</div>
                </div>
              </el-upload>
            </div>
          </el-form-item>
        </el-col>
      </el-row>

      <div class="form-actions">
        <el-button @click="router.back()">取消</el-button>
        <el-button type="primary" @click="handleSubmit">发布商品</el-button>
      </div>
    </el-form>
  </el-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import upload from '@/assets/upload.jpg'
// import { bookApi } from '@/api/book'
import { ElMessage } from 'element-plus'
import { mediaApi } from '@/api/media'
import { useUserStore } from '@/store/modules/user'
import { useRouter } from 'vue-router'
import { stuffApi } from '@/api/stuff'
const router = useRouter()
const userStore = useUserStore()
console.log(userStore.userInfo)
const itemForm = ref({
    name: '',
    category: '',
    price: 0,
    condition: '',
    description: '',
    photos: '',   // filename
    imageShow: '',
    publisher: userStore.userInfo?.id || '',
})

const formRef = ref()

const rules = {
    name: [{ required: true, message: '请输入商品名称', trigger: 'blur' }],
    photos: [{ required: true, message: '请上传商品图片', trigger: 'change' }],
    category: [{ required: true, message: '请选择商品类别', trigger: 'change' }],
    condition: [{ required: true, message: '请选择商品成色', trigger: 'change' }],
    price: [{ required: true, message: '请输入价格', trigger: 'change' }],
    description: [{ required: true, message: '请输入商品描述', trigger: 'blur' }]
}

const handleSubmit = () => {
    if (!formRef.value) return

    formRef.value.validate(async (valid: boolean) => {
        if (valid) {
            try {
                stuffApi.publishStuff(itemForm.value).then((res: any) => {
                    console.log(res)
                    ElMessage.success('发布成功')
                    router.push('/personal')
                })
            } catch (error) {
                ElMessage.error('发布失败')
            }
        } else {
            ElMessage.error('请填写完整所有必填项')
            return false
        }
    })
}

const handleReupload = () => {
    itemForm.value.photos = ''
    itemForm.value.imageShow = ''
}

const categoryList = ref([])
stuffApi.getStuffCategory().then((res: any) => {
    categoryList.value = res.category
    console.log(categoryList.value)
})

const conditionList = ref([
    { id: 1, name: '全新' },
    { id: 2, name: '九成新' },
    { id: 3, name: '八成新' },
    { id: 4, name: '七成新' },
    { id: 5, name: '六成新' },
    { id: 6, name: '五成新及以下' }
])

const beforeUpload = (file: File) => {
    const isImage = /^image\//.test(file.type)
    const isLt5M = file.size / 1024 / 1024 < 5

    if (!isImage) {
        ElMessage.error('只能上传图片文件！')
        return false
    }
    if (!isLt5M) {
        ElMessage.error('图片大小不能超过 5MB！')
        return false
    }
    return true
}

const customUpload = async (options: any) => {
    try {
        const formData = new FormData()
        formData.append('file', options.file)
        formData.append('type',"stuff")

        const response = await mediaApi.uploadImage(formData)
        itemForm.value.photos = response.filename
        itemForm.value.imageShow = response.preview
        ElMessage.success('图片上传成功')
    } catch (error) {
        ElMessage.error('图片上传失败')
    }
}

</script>

<style scoped>
.sell-container {
    max-width: 1200px;
    margin: 2rem auto;
}

.card-header {
    text-align: center;
    padding: 1rem 0;
}

.card-header h2 {
    color: var(--el-text-color-primary);
    margin: 0;
    font-size: 1.8rem;
}

.subtitle {
    color: var(--el-text-color-secondary);
    margin-top: 0.5rem;
    font-size: 0.9rem;
}

.upload-container {
    border: 2px dashed var(--el-border-color);
    border-radius: 8px;
    overflow: hidden;
}

.image-uploader {
    width: 100%;
}

.upload-placeholder {
    padding: 2rem;
    text-align: center;
    color: var(--el-text-color-secondary);
}

.upload-icon {
    font-size: 2rem;
    margin-bottom: 1rem;
}

.upload-text {
    font-size: 1rem;
    margin-bottom: 0.5rem;
}

.upload-tip {
    font-size: 0.8rem;
}

.image-preview {
    position: relative;
    width: 100%;
    aspect-ratio: 1;
}

.preview-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.image-actions {
    position: absolute;
    bottom: 1rem;
    left: 50%;
    transform: translateX(-50%);
    background: rgba(0, 0, 0, 0.5);
    padding: 0.5rem;
    border-radius: 4px;
}

.form-actions {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
    margin-top: 2rem;
    padding-top: 1rem;
    border-top: 1px solid var(--el-border-color-lighter);
}

:deep(.el-form-item__label) {
    font-weight: 500;
}

:deep(.el-input-number) {
    width: 100%;
}

.w-full {
    width: 100%;
}
</style>