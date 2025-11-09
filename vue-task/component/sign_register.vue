<script setup lang="js">
defineOptions({
  name: 'sign_register',
})
import { ref } from 'vue'
const BaseURL = `http://127.0.0.1:4523:8080`
const List = document.querySelector('.todoList')
const description = ref('')
const end_time = ref('')
const title = ref('')
console.log(List);

//POST ToDo
async function addTodo() {
  const req = new Request(BaseURL +'/todos', {
    method: 'POST',
    body: JSON.stringify({
      description: description.value,
      end_time: end_time.value,
      title: title.value,
    }),
  })
  const res = await fetch(req)
  const data = await res.json()
  console.log(data)
}

//GET ToDo
async function getTodo() {
  const data = await fetch(BaseURL + '/todos/list')
  if (!data.ok) {
    throw new Error(`HTTP error! status: ${data.status}`)
  }
  const res = await data.json()
  res.forEach((item) => {
    const todos = document.createElement('div')
    todos.innerHTML = `${item.title} - ${item.description}`
    List.appendChild(todos)
  })
}
getTodo()
</script>

<template>
  <div class="addTodo">
    <textarea v-model="description" placeholder="Enter description"></textarea>
    <input type="text" placeholder="Enter end_time" v-model="end_time" />
    <input type="text" placeholder="Enter title" v-model="title" />
    <button @click="addTodo">Add</button>
  </div>
  <div class="todoList">

  </div>
</template>

<style scoped></style>
