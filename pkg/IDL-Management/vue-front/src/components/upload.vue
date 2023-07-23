<template>
    <el-upload v-show="addState"
      ref="upload"
      class="upload-demo"
      :action="action"
      :limit="1"
      :on-exceed="handleExceed"
      :auto-upload="false"
      :on-error="handleError"
    >
      <template #trigger>
        <el-button type="primary">select file</el-button>
      </template>
      <el-button class="ml-3" type="success" @click="submitUpload">
        upload to server
      </el-button>
      <template #tip>
        <div class="el-upload__tip text-red">
          limit 1 file, new file will cover the old file
        </div>
      </template>
    </el-upload>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import { genFileId } from 'element-plus'
  import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
  import {SERVER_ADDR} from '../utils/config'

  defineProps(['addState'])

  const action = SERVER_ADDR + '/api/add'
  
  const upload = ref<UploadInstance>()
  
  const handleExceed: UploadProps['onExceed'] = (files) => {
    upload.value!.clearFiles()
    const file = files[0] as UploadRawFile
    file.uid = genFileId()
    upload.value!.handleStart(file)
  }

  const handleError: UploadProps['onError'] = (err, file) => {
    console.error(err, file)
    window.alert('upload failed')
  }
  
  const submitUpload = () => {
    upload.value!.submit()
  }
  </script>
  