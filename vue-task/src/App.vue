<template>
  <audio v-if="auto" autoplay="autoplay" src="../music/DJ阿轩 - 半点心.mp3" ></audio>
  <input type="text" placeholder="Where is this?"></input>
<button @click="play">我确定</button>
  <h1>{{ uinput }}</h1>
</template>

<script setup>
import { ref } from 'vue'
import { reactive } from 'vue'
let auto = ref(false)

function play() {
  const uinput = document.querySelector('input').value
  if (uinput === `canton` || uinput === `Canton` || uinput === `CANTON`) {
  auto.value = !auto.value
  }
}
</script>

<style>
html body {
  display: flex;
  justify-content: center;
  background-image: url(https://free.picui.cn/free/2025/11/06/690c8ffcadcec.jpg);
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
}
</style>
