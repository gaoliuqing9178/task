<script setup>
import { ref, watch } from 'vue'
import personal from '../component/personal.vue'

const sum = ref(0)

const stop = watch(sum, (newValue, oldValue) => {
  console.log(`sum was changed`)
  console.log(newValue, oldValue)
  if (newValue > 10) {
    stop()
  }
})
</script>

<template>
  <div class="case1">
    <p>{{ sum }}</p>
    <button @click="sum++">+</button>
    <button @click="sum--">-</button>
  </div>
  <div class="case2">
    <personal />
  </div>
</template>

<style>
html body {
  display: flex;
  justify-content: center;
  background-image: url(https://free.picui.cn/free/2025/11/06/690c8ffcadcec.jpg);
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
  margin: 40px;
}

.case1 {
  padding: 5px;
  border-style: solid;
  border-radius: 10px;
}
</style>
