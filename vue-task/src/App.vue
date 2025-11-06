<script setup>
import { ref } from 'vue'

// 给每个 todo 对象一个唯一的 id
let id = 0

const newTodo = ref('')
const todos = ref([
  { id: id++, text: 'Learn HTML' },
  { id: id++, text: 'Learn JavaScript' },
  { id: id++, text: 'Learn Vue' },
])

function addTodo() {
  todos.value.push({ id: id++, text: newTodo.value })
  newTodo.value = ''
}

function removeTodo(tod) {
  todos.value = todos.value.filter((t) => t !== tod)
}
</script>

<template>
  <form @submit.prevent="addTodo">
    <input v-model="newTodo" required placeholder="new todo" />
    <button>Add Todo</button>
  </form>
  <ul>
    <li v-for="tod in todos" :key="tod.id">
      {{ tod.text }}
      <button @click="removeTodo(todo)">X</button>
    </li>
  </ul>
</template>

<style>
html body {
  display: flex;
  justify-content: center;
  background-image: url(https://free.picui.cn/free/2025/11/06/690c8ffcadcec.jpg);
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
}

.cv-done {
  width: 700px;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.655);
  transition: all 0.3s ease-in-out;
}

.cv-done:hover {
  transform: scale(1.1);
  transition: all 0.3s ease-in-out;
}
</style>
