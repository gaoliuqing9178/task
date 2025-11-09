<script setup lang="js">
defineOptions({
  name: 'PersonalInfo',
})
import { reactive, watch } from 'vue'

const person = reactive({
  name: 'zhangsan',
  age: 20,
})

async function getPerson() {
  try {
    const u1 = await fetch(`https://jsonplaceholder.typicode.com/users`)
    if (!u1.ok) {
      throw new Error('Failed to fetch user data')
    }
    const u1d = await u1.json()
    return u1d[0].name
  } catch (error) {
    console.error(error)
  }
}

function changeName() {
  person.name += '~'
}

function changeAge() {
  person.age += 1
}

async function changePerson() {
  Object.assign(person, { name: await getPerson() })
  console.log(person)
}

watch(
  () => {
    return person.name
  },
  (newValue, oldValue) => {
    console.log(`person changed`)
    console.log(newValue, oldValue)
  },
  { deep: true, immediate: true },
)
</script>

<template>
  <p>{{ person.name }}</p>
  <p>{{ person.age }}</p>
  <button @click="changeName">1</button>
  <button @click="changeAge">2</button>
  <button @click="changePerson">3</button>
</template>

<style>
.case2 {
  padding: 5px;
  border-style: solid;
  border-radius: 10px;
}
</style>
