<script setup>
import { ref } from 'vue'
const BaseURL = `/api/v1`
const accountName = ref('')
const accountPassword = ref('')
const userName = ref('')
let token = ref(``)
const isRegister = ref(false)
const isSignin = ref(true)
const todoTitle = ref('')
const todoDescription = ref('')
const todoEndTime = ref('')
const source = ref([])
const isUser = ref(false)
let queryResult = ref(``)
let todoID = ref(``)

//注册
async function Register() {
  try {
    const req = new Request(`${BaseURL}/user/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        account: accountName.value,
        name: userName.value,
        password: accountPassword.value,
      }),
    })
    console.log(accountName.value)

    const res = await fetch(req)
    // 如果请求失败，抛出错误
    if (!res.ok) {
      throw new Error(`${res.status} ${res.statusText}`)
    }
    const data = await res.json()
    alert(data.message)
    console.log(res)
    //清空输入
    userName.value = ` `
    accountName.value = ` `
    accountPassword.value = ` `
  } catch (error) {
    console.log(error)
  }
}

//登录
async function Signin() {
  try {
    const req = new Request(`${BaseURL}/user/login`, {
      method: 'POST',
      body: JSON.stringify({
        account: accountName.value,
        password: accountPassword.value,
      }),
    })
    const res = await fetch(req)
    if (!res.ok) {
      throw new Error(`${res.status} ${res.statusText}`)
    }
    const data = await res.json()
    alert(data.message)

    //获取token&userName
    userName.value = data.user.name
    token.value = data.token
    //关闭登陆页面
    isSignin.value = !isSignin.value
    fetchToDoList()
    isUser.value = !isUser.value
  } catch (error) {
    console.log(error)
  }
}
function RegisterWindow() {
  isRegister.value = !isRegister.value
}

//获取待办事项
async function fetchToDoList() {
  const req = new Request(`${BaseURL}/todos/list`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
  })
  const res = await fetch(req)
  const data = await res.json()
  if (data == null || data.length === 0) {
    console.log('No data received')
    return
  }
  // 处理待办事项数据
  source.value = data
}

//创建待办事项
async function createToDo() {
  const req = new Request(`${BaseURL}/todos/add`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
    body: JSON.stringify({
      title: todoTitle.value,
      description: todoDescription.value,
      end_time: `${todoEndTime.value}` + `:00.000Z`,
    }),
  })

  try {
    const res = await fetch(req)
    const data = await res.json()
    console.log(data)
    // 刷新待办事项列表
    fetchToDoList()
  } catch (error) {
    console.log(error)
  }
}
//删除待办事项
async function deleteToDoList(item) {
  const req = new Request(`${BaseURL}/todos/delete/${item.id}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
  })

  try {
    const res = await fetch(req)
    const data = await res.json()
    console.log(data)
    //删除本地立即刷新（没招了）
    source.value = source.value.filter((todo) => todo.id !== item.id)
    // 刷新待办事项列表
    fetchToDoList()
  } catch (error) {
    console.log(error)
  }
}

//通过ID查询待办事项
async function queryToDo() {
  const req = new Request(`${BaseURL}/todos/get/${todoID.value}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
  })

  try {
    const res = await fetch(req)
    const data = await res.json()
    console.log(data)
    // 更新查询结果
    queryResult.value = data
  } catch (error) {
    console.log(error)
  }
}
</script>
<template>
  <h1>Welcome {{ userName }}</h1>
  <div class="Sign&register" v-if="isSignin">
    <!---注册和登录页面-->
    <!---登录-->
    <div class="sign" v-if="!isRegister">
      <p>Sign in</p>
      <input type="text" v-model="accountName" placeholder="Enter account name" />
      <input type="password" v-model="accountPassword" placeholder="Enter password" />
      <button @click="Signin">Signin</button>
      <button @click="RegisterWindow">Register</button>
    </div>
    <!---注册-->
    <div class="register" v-if="isRegister">
      <p>Register</p>
      <input type="text" v-model="accountName" placeholder="Enter account name" />
      <input type="password" v-model="accountPassword" placeholder="Enter password" />
      <input type="text" v-model="userName" placeholder="Enter user name" />
      <button @click="Register">Register</button>
      <button @click="RegisterWindow">Back to signin</button>
    </div>
  </div>
  <!---创建待办事项-->
  <div class="creatToDo" v-if="isUser">
    <h2>Create ToDo</h2>
    <input type="text" v-model="todoTitle" placeholder="Enter title" />
    <input type="text" v-model="todoDescription" placeholder="Enter description" />
    <input type="datetime-local" v-model="todoEndTime" />
    <button @click="createToDo">Create</button>
  </div>
  <!---通过ID查询待办事项-->
  <div class="queryToDo" v-if="isUser">
    <h2>Query ToDo</h2>
    <input type="text" v-model="todoID" placeholder="Enter ID" />
    <button @click="queryToDo">Query</button>
  </div>
  <div class="Todos">
    <p>{{ queryResult.title }}</p>
    <p>{{ queryResult.description }}</p>
    <p>{{ queryResult.end_time }}</p>
    <p>{{ queryResult.id }}</p>
  </div>
  <!---待办事项列表-->
  <div class="ToDoList" v-if="isUser">
    <h2>ToDo List</h2>
    <div v-if="source.length === 0">
      <p>No Todos</p>
    </div>
    <div id="ToDoList" v-for="item in source" :key="item.id" class="Todos">
      <p>{{ item.title }}</p>
      <p>{{ item.description }}</p>
      <p>{{ item.end_time }}</p>
      <p>{{ item.id }}</p>
      <button @click="deleteToDoList(item)" :id="item.id">Delete</button>
      <button @click="updateToDoList" :id="item.id">Update</button>
    </div>
  </div>
</template>

<style scoped></style>
