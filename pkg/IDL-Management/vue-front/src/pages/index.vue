<script setup>
import { ref, onMounted } from 'vue'
import Search from '../components/search.vue'
import Upload from '../components/upload.vue'
import Present from '../components/present.vue'
import { request } from '../utils/request.js'

const tableData = ref([
  
])

const present = ref(true)

const search = ref(false)

const add = ref(false)


onMounted(() => {
  pullAllData()
}
)

const pullAllData = () => {
  request({
    url: '/api/getAll',
    method: 'get',
  }).then(res => {
    console.log(res.data)
    tableData.value = res.data.Ls
  }).catch(err => {
    console.log(err)
    window.alert('获取数据失败')
  })
}

const jumpToSearch = () => {
  present.value = false
  search.value = true
  add.value = false
}

const jumpToAdd = () => {
  present.value = false
  search.value = false
  add.value = true
}

const jumpToPresent = () => {
  present.value = true
  search.value = false
  add.value = false
  pullAllData()
}



</script>


<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">IDL管理平台</el-header>
      <el-container>
        <el-aside width="200px" class="aside">
          <el-button type="primary" class="asideButton" plain @Click="jumpToPresent">主页</el-button>
          <br>
          <el-button type="primary" class="asideButton" plain @Click="jumpToSearch">查询</el-button>
          <br>
          <el-button type="primary" class="asideButton" plain @Click="jumpToAdd">增加</el-button>
          </el-aside>
        <el-main class="main">
          <Present :tableData="tableData" :presentState="present"></present>
          <Search :searchState="search"></search>
          <Upload :addState="add"></Upload>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>


<style scoped>
.common-layout {
  height: 100%;
}

.el-container {
  height: 100%;
}


.aside {
  text-align: center;
  background-color: #d3dce6;
}

.header {
  background-color:  #a0cfff;
  text-align: center;
  line-height: 60px;
}

.main {
  background-color: #e9eef3;
  color: #333;
  text-align: center;
  line-height: 120px;
}

.asideButton {
  width: 100%;
  height: 50px;
}

</style>
```

