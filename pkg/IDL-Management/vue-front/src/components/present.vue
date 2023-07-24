<script setup>
import { request } from '../utils/request';

const props = defineProps(['tableData', 'presentState'])

const deleteRow = (index) => {
  request({
    url: '/api/delete?name=' + props.tableData[index].Name,
    method: 'delete',
  }).then((res) => {
    props.tableData.splice(index, 1)
    console.log(res)
    window.alert('删除成功')
  }).catch((err) => {
    console.log(err)
    window.alert(err.message)
  })
}

const downloadRow = (index) => {
  request({
    url: '/api/download?name=' + props.tableData[index].Name,
    method: 'get',
    responseType: 'blob',
  }).then((res)=>{
        console.log('下载的文件',res)
        const link=document.createElement('a');
        try{
	          let blob =  res.data
            console.log(res.data)
	          let _fileName = res.headers['content-disposition'].split(';')[1].split('=')[1]
	          link.style.display='none';
	          // 兼容不同浏览器的URL对象
	          const url = window.URL || window.webkitURL || window.moxURL;
	          link.href=url.createObjectURL(blob);
	          link.download = _fileName;
	          link.click();
	          window.URL.revokeObjectURL(url);
        }catch (e) {
          console.log('下载的文件出错',e)
        }
      }).catch(()=>{
        console.log('下载的文件出错')
      })  
    }
</script>

<template>
    <el-table v-show="presentState"
            :data="tableData"
            style="width: 100%">
            <el-table-column
              prop="Date"
              label="服务上线日期"
              width="180">
            </el-table-column>
            <el-table-column
              prop="Name"
              label="服务名称"
              width="180">
            </el-table-column>
            <el-table-column
              prop="Description"
              label="相关描述">
            </el-table-column>
            <el-table-column fixed="right" label="Operations" width="200">
            <template #default="scope">
              <el-button
                link
                type="primary"
                size="small"
                @click.prevent="deleteRow(scope.$index)"
              >
                Remove
              </el-button>
              <el-button
                link
                type="primary"
                size="small"
                @click.prevent="downloadRow(scope.$index)"
              >
                Download
              </el-button>
            </template>
          </el-table-column>
          </el-table>
</template>