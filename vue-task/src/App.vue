<script setup>
import { ref } from 'vue'
const BaseURL = `/api/v1`

const accountName = ref('')
const accountPassword = ref('')
const userName = ref('')

const isRegister = ref(false)
const isSignin = ref(true)
const isUser = ref(false)

const todoTitle = ref('')
const todoDescription = ref('')
const todoEndTime = ref('')
const source = ref([])

const newTitle = ref('')
const newDescription = ref('')
const newEndTime = ref('')
const Id = ref(``)

let token = ref(``)
let isChange = ref(false)
let queryResult = ref(``)
let todoID = ref(``)
let isQuery = ref(false)
let isUpdate = ref(false)

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
    alert(`Username already exists`)
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
    console.log(token.value);
    
    //关闭登陆页面
    isSignin.value = !isSignin.value
    fetchToDoList()
    isUser.value = !isUser.value
  } catch (error) {
    alert(`Account or password is wrong`)
    console.log(error)
  }
}
function RegisterWindow() {
  isRegister.value = !isRegister.value
}

//更改用户信息

const newUserName = ref('')
const newUserPassword = ref('')

function changeWindow() {
  isChange.value = !isChange.value
}

async function changeUser() {
  const req = new Request(`${BaseURL}/user/update`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
    body: JSON.stringify({
      name: userName.value,
      password: newUserPassword.value,
    }),
  })
  const res = await fetch(req)
  const data = await res.json()
  alert(`${data.message}`)
  isChange.value = !isChange.value
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
    todoTitle.value = ``
    todoDescription.value = ``
    todoEndTime.value = ``
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
    // 更新查询结果
    queryResult.value = data
    console.log(queryResult.value.ID)

    isQuery.value = !isQuery.value
  } catch (error) {
    console.log(error)
  }
}

//更新待办事项

function Toupdate() {
  isUpdate.value = !isUpdate.value
}

async function updateToDoList() {
  const req = new Request(`${BaseURL}/todos/update/${Id.value}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
    body: JSON.stringify({
      description: newDescription.value,
      end_time: `${newEndTime.value}` + `:00.000Z`,
      status: 1,
      title: newTitle.value
    }),
  })
  const res = await fetch(req)
  console.log(res);
  const data = await res.json()
  console.log(data);
  isUpdate.value = !isUpdate.value
  newDescription.value = ``
  newEndTime.value = ``
  newTitle.value = ``
  fetchToDoList()
}
</script>
<!------------------------------------------------------------------------------------------------------>
<template>
  <div class="Sign&register" v-if="isSignin">
    <!---注册和登录页面-->
    <!---登录-->
    <div class="sign" v-if="!isRegister">
      <h3>Sign in</h3>
      <input class="ac" type="text" v-model="accountName" placeholder="Enter account name" />
      <input class="ac" type="password" v-model="accountPassword" placeholder="Enter password" />
      <button class="acbtn" @click="Signin">Signin</button>
      <button class="acbtn" @click="RegisterWindow">Register</button>
    </div>
    <!---注册-->
    <div class="register" v-if="isRegister">
      <h3>Register</h3>
      <input class="ac" type="text" v-model="accountName" placeholder="Enter account name" />
      <input class="ac" type="password" v-model="accountPassword" placeholder="Enter password" />
      <input class="ac" type="text" v-model="userName" placeholder="Enter user name" />
      <button class="acbtn" @click="Register">Register</button>
      <button class="acbtn" @click="RegisterWindow">Back to sign in</button>
    </div>
  </div>
  <!--------------------------------------------------------------------------------------------------------------------------------------->
  <div class="LIST" v-if="isUser">
    <!---用户信息-->
    <div class="user">
      <h1>Welcome {{ userName }}</h1>
      <button @click="changeWindow">Change user data</button>
      <!--更改用户信息-->
      <div class="changeData" v-if="isChange">
        <input
          type="text"
          v-model="newUserName"
          placeholder="Enter the username to change pointed account"
        />
        <input type="password" v-model="newUserPassword" placeholder="Enter new password" />
        <p class="tips">you can't change your account name</p>
        <button @click="changeUser">change</button>
        <button @click="changeWindow">cancel</button>
      </div>
    </div>
    <!---创建待办事项-->
    <div class="creatToDo">
      <h3>Create todo</h3>
      <input type="text" v-model="todoTitle" placeholder="Enter title" />
      <input type="text" v-model="todoDescription" placeholder="Enter description" />
      <input type="datetime-local" v-model="todoEndTime" />
      <button @click="createToDo">Add todo</button>
    </div>
    <!---通过ID查询待办事项-->
    <div class="queryToDo">
      <h3>Query todo by id</h3>
      <input type="text" v-model="todoID" placeholder="Enter id" />
      <button @click="queryToDo">Query</button>
    </div>
    <div class="Todos" v-if="isQuery">
      <p>Title: {{ queryResult.title }}</p>
      <p>Description: {{ queryResult.description }}</p>
      <p>Deadline: {{ queryResult.end_time }}</p>
    </div>
    <!---通过ID更新指定待办事项-->
    <h3>Change todo by id</h3>
    <div class="updateTodo" v-if="isUpdate">
      <input type="text" placeholder="Enter id" v-model="Id">
      <input type="text" placeholder="Enter new title" v-model="newTitle" />
      <input type="text" placeholder="Enter new description" v-model="newDescription" />
      <input type="datetime-local" v-model="newEndTime" />
      <button @click="updateToDoList()">Update</button>
      <button @click="Toupdate">Cancel</button>
    </div>
    <button class="updateWindow" v-if="!isUpdate" @click="Toupdate">Update</button>
<!------------------------------------------------------------------------------------------------------------------------------------------->
    <hr>
    <!---待办事项列表-->
    <div class="ToDoList">
      <h3>Todo List</h3>
      <div v-if="source.length === 0">
        <p>No Todos</p>
      </div>
      <div id="ToDoList" v-for="item in source" :key="item.id" class="Todos">
        <p>Title: {{ item.title }}</p>
        <p>Description: {{ item.description }}</p>
        <time>Deadline: {{ item.end_time }}</time>
        <p>ID: {{ item.id }}</p>
        <button @click="deleteToDoList(item)" :id="item.id">Delete</button>
      </div>
    </div>
  </div>

</template>

<style>
html, body {
  display: flex;
  justify-content: center;
  min-height: 100%;
  width: 100%;
}

.changeData, .user {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 60%;
}

.changeData button {
  margin: 10px;
  width: 20%;
}

.user button {
  margin: 10px;
  background-color: rgb(0, 136, 255);
  color: white;
  border: none;
  border-radius: 15px;
  font-size: 15px;
  padding: 15px;
}

.Sign\&register {
  display: flex;
  justify-content: center;
  height: 95%;
  margin: 20px;
}

.sign,
.register {
  width: 60%;
  align-self: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #f9f9f9; /* 添加背景方便观察 */
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  transition: all 0.25s ease-in-out;
}

.sign button, .register button{
  display: block;
  margin: 10px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 10px;
  background-color: rgb(0, 136, 255);
  font-weight: 500;
  width: 20%;
  transition: all 0.25s ease-in-out;
}

input {
  width: 80%;
  padding: 15px;
  margin: 10px 0;
  border: 4px;
  border-radius: 15px;
  border: 1px solid #ccc;
  transition: all 0.25s ease-in-out;
}

.LIST {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.creatToDo button, .queryToDo button, .Todos button, .updateTodo button, .updateWindow{
  margin: 10px;
  width: 40%;
  background-color: rgb(0, 136, 255);
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 12px;
  padding: 8px;
}

.updateWindow {
  width: 20%;
}

.Todos {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  border-style: solid;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 20px;
  margin: 20px;
  padding: 20px;
}

.queryToDo, .updateTodo, .creatToDo {
  width: 60%;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 20px;
}

hr {
  width: 99%;
  margin: 20px;
}

.ToDoList h3 {
  width: 100%;
  text-align: center;
}

.ToDoList {
  width: 99%;
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
}
</style>
