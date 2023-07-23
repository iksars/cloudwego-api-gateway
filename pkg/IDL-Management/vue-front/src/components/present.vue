<script setup>
import { request } from '../utils/request';

const props = defineProps(['tableData', 'presentState'])

const deleteRow = (index) => {
  request({
    url: '/api/delete?name=' + props.tableData[index].name,
    method: 'delete',
  }).then((res) => {
    props.tableData.splice(index, 1)
    console.log(res)
  }).catch((err) => {
    console.log(err)
    window.alert('删除失败')
  })
}

const downloadRow = (index) => {
  request({
    url: '/api/download?name=' + props.tableData[index].name,
    method: 'get',
    responseType: 'blob',
  }).then((res)=>{
        console.log('下载的文件',res)
        const link=document.createElement('a');
        try{
	          let blob =  res.data
	          let _fileName = res.headers['content-disposition'].split(';')[1].split('=')[1];//文件名，中文无法解析的时候会显示 _(下划线),生产环境获取不到
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
              prop="date"
              label="服务上线日期"
              width="180">
            </el-table-column>
            <el-table-column
              prop="name"
              label="服务名称"
              width="180">
            </el-table-column>
            <el-table-column
              prop="IDLfile"
              label="API">
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