<script setup>
import { ref } from 'vue'
const BaseURL = `http://localhost:8848/api/v1`
const accountName = ref("")
const accountPassword = ref("")
const userName = ref("")
let token = ref(``)

//注册
async function Register() {
  const req = new Request(`${BaseURL}/user/register`, {
    method: 'POST',
    body: JSON.stringify({
      "name": userName.value,
      "account": accountName.value,
      "password": accountPassword.value,
    }),
  })
  const res = await fetch(req)
  const data = await res.json()
  alert(data.message)
  userName.value = ` `
  accountName.value = ` `
  accountPassword.value = ` `
}

//登录
async function Signin() {
  const req = new Request(`${BaseURL}/user/signin`, {
    method: 'POST',
    body: JSON.stringify({
      "account": accountName.value,
      "password": accountPassword.value,
    }),
  })
  const res = await fetch(req)
  const data = await res.json()
  token.value = data.token
  console.log(token.value)
}
</script>

<template>
  <p>Register</p>
  <input type="text" v-model="accountName" placeholder="Enter name" />
  <input type="password" v-model="accountPassword" placeholder="Enter password"/>
  <input type="text" v-model="userName" placeholder="Enter user name"/>
  <button @click="Register">Register</button>
  <button @click="Signin">Signin</button>
</template>

<style scoped></style>
