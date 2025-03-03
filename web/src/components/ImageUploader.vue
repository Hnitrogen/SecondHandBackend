<template>
    <div class="image-uploader">
      <!-- 触发按钮 -->
      <div class="upload-trigger" @mouseenter="showUploadBtn = true" @mouseleave="showUploadBtn = false">
        <slot>
          <!-- 默认显示 -->
          <img :src="modelValue || defaultImage" :alt="alt" class="preview-image">
        </slot>
        
        <!-- 悬浮上传按钮 -->
        <div v-if="showUploadBtn" class="upload-btn" @click="openUploadModal">
          <span>更换图片</span>
        </div>
      </div>
  
      <!-- 上传模态框 -->
      <div v-if="showModal" class="modal-overlay" @click.self="closeModal2">
        <div class="modal-content">
          <div class="modal-header">
            <h3>上传图片</h3>
            <button class="close-btn" @click="closeModal2">×</button>
          </div>
          
          <div class="upload-area">
            <div class="drop-zone" 
                 @drop.prevent="handleDrop"
                 @dragover.prevent="dragover = true"
                 @dragleave.prevent="dragover = false"
                 :class="{ 'dragover': dragover }"
                 @click="triggerFileInput">
              <input type="file" 
                     ref="fileInput" 
                     class="file-input" 
                     accept="image/*"
                     @change="handleFileChange">
              <div v-if="!previewUrl" class="upload-hint">
                <i class="upload-icon">📁</i>
                <p>点击或拖拽图片至此处上传</p>
                <p class="upload-tip">支持 jpg、png、gif 格式，最大 5MB</p>
              </div>
              <img v-else :src="previewUrl" class="preview-img" alt="预览图">
            </div>
          </div>
  
          <div class="modal-footer">
            <button class="cancel-btn" @click="closeModal2">取消</button>
            <button class="confirm-btn" 
                    :disabled="!previewUrl" 
                    @click="handleUpload">确认上传</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps, defineEmits } from 'vue'
  import {mediaApi} from '@/api/media.ts'
  import upload from '@/assets/upload.jpg'
  
  const props = defineProps({
    modelValue: {
      type: String,
      default: ''
    },
    alt: {
      type: String,
      default: '图片'
    },
    defaultImage: {
      type: String,
      default: upload
    }
  })
  
  const emit = defineEmits(['update:modelValue', 'upload-success', 'upload-error', 'confirm'])
  
  const showModal = ref(false)
  const showUploadBtn = ref(false)
  const dragover = ref(false)
  const previewUrl = ref('')
  const fileInput = ref<HTMLInputElement | null>(null)
  const selectedFile = ref<File | null>(null)
  
  // 打开上传模态框
  const openUploadModal = () => {
    showModal.value = true
  }
  
  // 关闭模态框
  const closeModal2 = () => {
    showModal.value = false
    previewUrl.value = ''
    selectedFile.value = null
  }
  
  // 触发文件选择
  const triggerFileInput = () => {
    fileInput.value?.click()
  }
  
  // 处理文件选择
  const handleFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement
    if (target.files && target.files[0]) {
      const file = target.files[0]
      if (validateFile(file)) {
        selectedFile.value = file
        createPreview(file)
      }
    }
  }
  
  // 处理拖拽
  const handleDrop = (event: DragEvent) => {
    dragover.value = false
    const file = event.dataTransfer?.files[0]
    if (file && validateFile(file)) {
      selectedFile.value = file
      createPreview(file)
    }
  }
  
  // 验证文件
  const validateFile = (file: File) => {
    const validTypes = ['image/jpeg', 'image/png', 'image/gif']
    const maxSize = 5 * 1024 * 1024 // 5MB
  
    if (!validTypes.includes(file.type)) {
      alert('请上传 jpg、png 或 gif 格式的图片')
      return false
    }
  
    if (file.size > maxSize) {
      alert('图片大小不能超过 5MB')
      return false
    }
  
    return true
  }
  
  // 创建预览
  const createPreview = (file: File) => {
    const reader = new FileReader()
    reader.onload = (e) => {
      previewUrl.value = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
  
  // 处理上传
  const handleUpload = async () => {
    if (!selectedFile.value) {
      return
    }
    
    try {
      const formData = new FormData()
      formData.append('file', selectedFile.value)
      
      const resp = await mediaApi.uploadImage(formData)
      // 把结果传递回父组件，并添加事件对象
      emit('upload-success', resp, event)
      emit('confirm', event)  // 添加新的事件
      closeModal2()
    } catch (error) {
      console.error('上传错误:', error)
      emit('upload-error', error)
    }
  }
  </script>
  
  <style scoped>
  .image-uploader {
    position: relative;
    display: inline-block;
  }
  
  .upload-trigger {
    position: relative;
    display: inline-block;
    width: 100px;  /* 改小尺寸 */
    height: 100px; /* 改小尺寸 */
    /* border-radius: 50%; */
    overflow: hidden;
  }
  
  .preview-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    /* border-radius: 50%; */
  }
  
  .upload-btn {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, 0.6);
    color: white;
    padding: 4px; /* 减小内边距 */
    font-size: 14px; /* 减小字体大小 */
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
  }
  
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
  }
  
  .modal-content {
    background: white;
    border-radius: 8px;
    padding: 20px;
    width: 90%;
    max-width: 400px; /* 减小弹窗最大宽度 */
  }
  
  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
  
  .close-btn {
    background: none;
    border: none;
    font-size: 24px;
    cursor: pointer;
    color: #666;
  }
  
  .drop-zone {
    border: 2px dashed var(--light-green);
    border-radius: 8px;
    padding: 20px;
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
  }
  
  .drop-zone.dragover {
    border-color: var(--primary-green);
    background: rgba(107, 142, 115, 0.1);
  }
  
  .file-input {
    display: none;
  }
  
  .upload-hint {
    color: #666;
  }
  
  .upload-icon {
    font-size: 40px;
    margin-bottom: 10px;
  }
  
  .upload-tip {
    font-size: 12px;
    color: #999;
    margin-top: 8px;
  }
  
  .preview-img {
    max-width: 100%;
    max-height: 300px;
    object-fit: contain;
  }
  
  .modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
    margin-top: 20px;
  }
  
  .cancel-btn,
  .confirm-btn {
    padding: 8px 16px;
    border-radius: 4px;
    cursor: pointer;
    transition: all 0.3s ease;
  }
  
  .cancel-btn {
    background: #f5f5f5;
    border: 1px solid #ddd;
    color: #666;
  }
  
  .confirm-btn {
    background: var(--primary-green);
    border: none;
    color: white;
  }
  
  .confirm-btn:disabled {
    background: #ccc;
    cursor: not-allowed;
  }
  </style> 