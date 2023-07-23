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
        url: '/api/search?'+inputServiceName.value,
        method: 'get',
    }).then(res => {
        console.log(res)
        result = res.data.data
    }).catch(err => {
        console.log(err)
        window.alert('查询失败')
    })
    result.value = [
        {
            date: '2021-10-01',
            name: 'service1',
            IDLfile: 'file1',
        },
        {
            date: '2021-10-02',
            name: 'service2',
            IDLfile: 'file2',
        },
        {
            date: '2021-10-03',
            name: 'service3',
            IDLfile: 'file3',
        },
    ]
    present.value = true
}
</script>



<template>
<div v-show="searchState">
    <el-input v-model="inputServiceName" placeholder="Please input" maxlength="20"/>
    <el-button type="primary" plain @click="query">查询</el-button>
    <Present :table-data="result" :present-state="present" />
  </div>
</template>
