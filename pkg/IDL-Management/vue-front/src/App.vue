<script setup>
import { onMounted } from 'vue';
import { RouterView } from 'vue-router';

onMounted(() => {
  console.log('app mounted');
});

</script>

<template>
  <RouterView />
</template>


<style>
html,body,#app{
  padding: 0px;
  margin: 0px;
  height: 100%;
}
</style>
