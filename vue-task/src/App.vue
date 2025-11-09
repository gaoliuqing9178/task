<script setup>

import sign_register from '../component/sign_register.vue'

import { ref, reactive, watch } from 'vue'
const accountn = ref(``)
const accountp = ref(``)

function closeWindow() {
  const W = document.querySelector('.sign')
  W.style.display = 'none'
}
</script>

<template>
  <div class="sign">
    <input class="username" id="ac" type="text" placeholder="Enter usename" v-model="accountn" />
    <input
      class="password"
      id="ac"
      type="password"
      placeholder="Enter password"
      v-model="accountp"
    />
    <div class="btnarea">
      <button @click="closeWindow">Sign in</button>
      <button @click="closeWindow">Register</button>
    </div>
  </div>
  <div class="signWindow">
  <sign_register/>
  </div>
</template>

<style>
  html, body {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .sign {
  border-style: solid;
  border-radius: 15px;
  padding: 50px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 20px;
}

#ac {
  border-style: solid;
  border-radius: 25px;
  padding: 10px;
  margin: 10px;
}

.btnarea {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}

button {
  background-color: rgb(68, 171, 255);
  border: none;
  border-radius: 15px;
  padding: 10px;
  margin: 10px;
  color: white;
  font-size: 15px;
  font-weight: 500;
  transition: all 0.25s;
}

button:hover {
  background-color: rgb(58, 151, 235);
  transform: scale(1.1);
}

.sign {
  display: none;
}
</style>
