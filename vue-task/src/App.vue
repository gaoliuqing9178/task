<script setup>
import { ref } from 'vue'
const name = ref('555664')
const unumber = ref(777899)
let data = ref(null)
function addunumber() {
  unumber.value++
}
async function getData() {
  try {
    const res = await fetch('https://jsonplaceholder.typicode.com/todos/1')
    if (!res.ok) {
      throw new Error(`请求失败: ${res.status}`)
    }
    const udata = await res.json()
    data.value = udata.title
    console.log(data.value)
  } catch (error) {
    console.error(error)
  }
}
</script>

<template>
  <input type="text" v-model="name" />
  <h1>hello world</h1>
  <h1>{{ name }}</h1>
  <h1>{{ unumber }}</h1>
  <button @click="addunumber">+1</button>
  <button @click="getData">getData</button>
  <h1>{{ data }}</h1>
</template>

<style>
h1 {
  color: #42b983;
  box-shadow: 0 0 10px #42b983;
}
</style>
