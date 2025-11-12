<script setup>
import { ref } from 'vue'
import { onMounted } from 'vue'

const BaseURL = `/api/v1`

const userName = ref('')
const newUserName = ref('')
const newUserPassword = ref('')
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

// 从本地存储读取token
onMounted(() => {
  const savedToken = localStorage.getItem('userToken')
  const savedUserName = localStorage.getItem('userName')
  if (savedToken) {
    token.value = savedToken
    userName.value = savedUserName
  }
  //获取待办事项列表
  fetchToDoList()
})
//更改用户信息

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
      name: newUserName.value,
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
    if (res.status == 400) {
      alert(`Please enter compelete information`)
    }
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

<template>
    <div class="LIST">
    <!---用户信息-->
    <div class="user">
      <h1>Welcome {{ userName }}</h1>
      <button @click="changeWindow">Change user data</button>
      <!--更改用户信息-->
      <div class="changeData" v-if="isChange">
        <input
          type="text"
          v-model="newUserName"
          placeholder="Enter the username"
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

<style scoped>

.changeData, .user {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 60%;
}

.changeData button {
  width: 40%;
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

@media (max-width: 500px) {
    .creatToDo button, .queryToDo button, .Todos button, .updateTodo button, .updateWindow, .changeData button{
        width: 70%;
    }
    
    .Todos {
        width: 90%;
    }
    
    .queryToDo, .updateTodo, .creatToDo {
        width: 80%;
    }

    p, button, input, textarea, time {
      font-size: 10px;
    }
}
</style>