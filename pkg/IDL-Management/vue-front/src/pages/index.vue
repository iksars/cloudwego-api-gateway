<script setup>
import { ref, onMounted } from 'vue'
import Search from '../components/search.vue'
import Upload from '../components/upload.vue'
import Present from '../components/present.vue'
import { request } from '../utils/request.js'

const tableData = ref([
  {
    date: '2016-05-02',
    name: '王小虎',
    IDLfile: '上海市普陀区金沙江路 1518 弄'
  },
  {
    date: '2016-05-04',
    name: '王小虎',
    IDLfile: '上海市普陀区金沙江路 1517 弄'
  },
  {
    date: '2016-05-01',
    name: '王小虎',
    IDLfile: '上海市普陀区金沙江路 1519 弄'
  },
  {
    date: '2016-05-03',
    name: '王小虎',
    IDLfile: '上海市普陀区金沙江路 1516 弄'
  }
])

const present = ref(true)

const search = ref(false)

const add = ref(false)


onMounted(() => {
  request({
    url: '/api/getAll',
    method: 'get',
  }).then(res => {
    console.log(res)
    tableData.value = res.data.data
  }).catch(err => {
    console.log(err)
    window.alert('查询失败')
  })
})



</script>


<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">IDL管理平台</el-header>
      <el-container>
        <el-aside width="200px" class="aside">
          <el-button type="primary" class="asideButton" plain @Click="present=true;search=false;add=false">主页</el-button>
          <br>
          <el-button type="primary" class="asideButton" plain @Click="search=true;add=false;present=false">查询</el-button>
          <br>
          <el-button type="primary" class="asideButton" plain @Click="add=true;search=false;present=false">增加</el-button>
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

