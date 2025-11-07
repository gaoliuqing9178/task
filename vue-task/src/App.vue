<script setup>
import { ref } from 'vue'
import { reactive } from 'vue'
import personal from '../component/personal.vue'
components: {
  personal
}

const a = reactive({sadas: 4654})

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
  todos.value.splice(todos.value.indexOf(tod), 1)
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
  <personal />
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
</style>
