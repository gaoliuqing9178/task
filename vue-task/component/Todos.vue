<script setup>
import { ref } from 'vue'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

const BaseURL = `/api/v1`
const router = useRouter()

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
  } else {
    router.push('/')
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
  //检查输入是否为空
  if (newUserName.value == '' || newUserPassword.value == '') {
    alert('用户名或密码不能为空')
    return
  }
  
  const res = await fetch(req)
  const data = await res.json()
  alert(`${data.message}`)
  isChange.value = !isChange.value
}

//退出登录
function signout() {
  localStorage.removeItem('userName')
  localStorage.removeItem('userToken')
  router.push('/')
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

  //检查输入是否为空
  if (todoTitle.value == '' || todoDescription.value == '' || todoEndTime.value == '') {
    alert(`不可以输入空白内容`)
    return
  }

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
    // 检查ID是否为空
    if (todoID.value == '') {
      alert(`请输入ID`)
      return
    }

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
  //检查输入是否为空
  if (newDescription.value == '' || newEndTime.value == '' || newTitle.value == '') {
    alert(`请输入完整信息`)
    return
  }

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
      <h1>欢迎 {{ userName }}</h1>
      <div class="user">
        <button @click="signout">退出登录</button>
        <button @click="changeWindow">更改用户信息</button>
      </div>
      <!--更改用户信息-->
      <div class="changeData" v-if="isChange">
        <input
          type="text"
          v-model="newUserName"
          placeholder="输入用户名"
        />
        <input type="password" v-model="newUserPassword" placeholder="输入新密码" />
        <p class="tips">更改密码</p>
        <button @click="changeUser">确认</button>
        <button @click="changeWindow">取消</button>
      </div>
    </div>
    <!---创建待办事项-->
    <div class="creatToDo">
      <h3>创建待办事项</h3>
      <input type="text" v-model="todoTitle" placeholder="输入标题" />
      <input type="text" v-model="todoDescription" placeholder="输入描述" />
      <input type="datetime-local" v-model="todoEndTime" />
      <button @click="createToDo">添加待办事项</button>
    </div>
    <!---通过ID查询待办事项-->
    <div class="queryToDo">
      <h3>通过ID查询待办事项</h3>
      <input type="text" v-model="todoID" placeholder="输入ID" />
      <button @click="queryToDo">查询</button>
    </div>
    <div class="Todos" v-if="isQuery">
      <p>标题： {{ queryResult.title }}</p>
      <p>描述： {{ queryResult.description }}</p>
      <p>截止日期: {{ queryResult.end_time }}</p>
    </div>
    <!---通过ID更新指定待办事项-->
    <h3>通过ID更改指定待办事项</h3>
    <div class="updateTodo" v-if="isUpdate">
      <input type="text" placeholder="输入ID" v-model="Id">
      <input type="text" placeholder="输入新的标题" v-model="newTitle" />
      <input type="text" placeholder="输入新的描述" v-model="newDescription" />
      <input type="datetime-local" v-model="newEndTime" />
      <button @click="updateToDoList()">确认</button>
      <button @click="Toupdate">取消</button>
    </div>
    <button class="updateWindow" v-if="!isUpdate" @click="Toupdate">更改</button>
<!------------------------------------------------------------------------------------------------------------------------------------------->
    <hr>
    <!---待办事项列表-->
    <div class="ToDoList">
      <h3>待办事项列表</h3>
      <div v-if="source.length === 0">
        <p>暂无待办事项</p>
      </div>
      <div id="ToDoList" v-for="item in source" :key="item.id" class="Todos">
        <p>标题: {{ item.title }}</p>
        <p>描述: {{ item.description }}</p>
        <time>截止日期: {{ item.end_time }}</time>
        <p>ID: {{ item.id }}</p>
        <button @click="deleteToDoList(item)" :id="item.id">删除</button>
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

.Todos button {
  background-color: brown;
}

@media (max-width: 500px) {
    .creatToDo button, .queryToDo button, .Todos button, .updateTodo button, .updateWindow, .changeData button{
        width: 50%;
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