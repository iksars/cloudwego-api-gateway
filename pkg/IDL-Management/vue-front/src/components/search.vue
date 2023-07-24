<script setup>
import { ref,reactive } from 'vue'
import {request} from '../utils/request.js'
import Present from './present.vue'

defineProps(['searchState'])

const inputServiceName = ref('')

const result = ref([])

const present = ref(false)

const query = () => {
  console.log(inputServiceName.value)
    request({
        url: '/api/search?name='+inputServiceName.value,
        method: 'get',
    }).then(res => {
        console.log(res)
        result.value = res.data.Ls
        present.value = true
    }).catch(err => {
        console.log(err)
        window.alert('查询失败')
    })
}
</script>



<template>
<div v-show="searchState">
    <el-input v-model="inputServiceName" placeholder="Please input" maxlength="20"/>
    <el-button type="primary" plain @click="query">查询</el-button>
    <Present :table-data="result" :present-state="present" />
  </div>
</template>
