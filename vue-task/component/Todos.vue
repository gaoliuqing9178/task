<script setup>
defineOptions({
  name: 'TodosView',
})

import { computed, ref } from 'vue'
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
const isListLoading = ref(true)
const isCreating = ref(false)
const isQueryingTodo = ref(false)
const isUpdatingTodo = ref(false)
const deletingId = ref(null)
const keyword = ref('')
const statusFilter = ref('all')
const isUpdatingProfile = ref(false)
const feedback = ref({
  type: '',
  message: '',
})

const newTitle = ref('')
const newDescription = ref('')
const newEndTime = ref('')
const Id = ref(``)
const token = ref(``)
const isChange = ref(false)
const queryResult = ref(``)
const todoID = ref(``)
const isQuery = ref(false)
const isUpdate = ref(false)
let feedbackTimer

const filteredTodos = computed(() => {
  if (!source.value?.length) return []
  const kw = keyword.value.trim().toLowerCase()
  return source.value.filter((item) => {
    const matchKeyword =
      !kw ||
      item.title?.toLowerCase().includes(kw) ||
      item.description?.toLowerCase().includes(kw)
    const matchStatus =
      statusFilter.value === 'all'
        ? true
        : statusFilter.value === 'done'
          ? item.status === 1
          : item.status !== 1
    return matchKeyword && matchStatus
  })
})

function setFeedback(type, message) {
  feedback.value = { type, message }
  clearTimeout(feedbackTimer)
  if (message) {
    feedbackTimer = setTimeout(() => {
      feedback.value = { type: '', message: '' }
    }, 4000)
  }
}

function formatDate(dateString) {
  if (!dateString) return '--'
  try {
    return new Intl.DateTimeFormat('zh-CN', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    }).format(new Date(dateString))
  } catch (error) {
    console.log(error)
    return dateString
  }
}

function getStatusLabel(status) {
  return status === 1 ? '已完成' : '进行中'
}

function getStatusClass(status) {
  return status === 1 ? 'status-badge--done' : 'status-badge--pending'
}

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
  if (isUpdatingProfile.value) return
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
    setFeedback('error', '用户名或密码不能为空')
    return
  }
  try {
    isUpdatingProfile.value = true
    const res = await fetch(req)
    const data = await res.json()
    if (!res.ok) {
      throw new Error(data.message || '更新失败')
    }
    setFeedback('success', data.message || '用户信息已更新')
    isChange.value = !isChange.value
    newUserName.value = ''
    newUserPassword.value = ''
  } catch (error) {
    console.log(error)
    setFeedback('error', error.message || '更新失败，请稍后再试')
  } finally {
    isUpdatingProfile.value = false
  }
}

//退出登录
function signout() {
  localStorage.removeItem('userName')
  localStorage.removeItem('userToken')
  router.push('/')
}

//获取待办事项
async function fetchToDoList() {
  isListLoading.value = true
  const req = new Request(`${BaseURL}/todos/list`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
  })
  try {
    const res = await fetch(req)
    const data = await res.json()
    if (!data || data.length === 0) {
      source.value = []
      setFeedback('info', '当前暂无待办事项')
      return
    }
    source.value = data
  } catch (error) {
    console.log(error)
    setFeedback('error', '获取待办事项失败，请稍后再试')
  } finally {
    isListLoading.value = false
  }
}

//创建待办事项
async function createToDo() {
  if (isCreating.value) return
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
    setFeedback('error', '请输入完整的标题、描述与截止时间')
    return
  }

  try {
    isCreating.value = true
    const res = await fetch(req)
    if (!res.ok) {
      throw new Error('创建失败')
    }
    const data = await res.json()
    console.log(data)
    setFeedback('success', data.message || '已添加新的待办事项')
    // 刷新待办事项列表
    await fetchToDoList()
    todoTitle.value = ``
    todoDescription.value = ``
    todoEndTime.value = ``
  } catch (error) {
    console.log(error)
    setFeedback('error', error.message || '创建失败，请稍后重试')
  } finally {
    isCreating.value = false
  }
}

//删除待办事项

async function deleteToDoList(item) {
  if (deletingId.value) return
  const confirmed = confirm(`确认删除待办事项「${item.title}」吗？`)
  if (!confirmed) return
  deletingId.value = item.id
  const req = new Request(`${BaseURL}/todos/delete/${item.id}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token.value}`,
    },
  })

  try {
    const res = await fetch(req)
    if (!res.ok) {
      throw new Error('删除失败')
    }
    const data = await res.json()
    console.log(data)
    //删除本地立即刷新（没招了）
    source.value = source.value.filter((todo) => todo.id !== item.id)
    // 刷新待办事项列表
    fetchToDoList()
    setFeedback('success', data.message || '已删除该待办事项')
  } catch (error) {
    console.log(error)
    setFeedback('error', error.message || '删除失败，请稍后再试')
  } finally {
    deletingId.value = null
  }
}

//通过ID查询待办事项

async function queryToDo() {
  if (isQueryingTodo.value) return
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
      setFeedback('error', '请输入待办事项ID')
      return
    }

    isQueryingTodo.value = true
    const res = await fetch(req)
    if (!res.ok) {
      throw new Error('查询失败')
    }
    const data = await res.json()
    // 更新查询结果
    queryResult.value = data
    console.log(queryResult.value.ID)

    isQuery.value = !isQuery.value
    setFeedback('info', '已加载该待办事项信息')
  } catch (error) {
    console.log(error)
    setFeedback('error', error.message || '查询失败，请稍后重试')
  } finally {
    isQueryingTodo.value = false
  }
}

//更新待办事项

function Toupdate() {
  isUpdate.value = !isUpdate.value
}

async function updateToDoList() {
  if (isUpdatingTodo.value) return
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
    setFeedback('error', '请填写完整的更新信息')
    return
  }

  try {
    isUpdatingTodo.value = true
    const res = await fetch(req)
    if (!res.ok) {
      throw new Error('更新失败')
    }
    console.log(res)
    const data = await res.json()
    console.log(data)
    isUpdate.value = !isUpdate.value
    newDescription.value = ``
    newEndTime.value = ``
    newTitle.value = ``
    await fetchToDoList()
    setFeedback('success', data.message || '待办事项已更新')
  } catch (error) {
    console.log(error)
    setFeedback('error', error.message || '更新失败，请稍后再试')
  } finally {
    isUpdatingTodo.value = false
  }
}

</script>

<template>
  <div class="LIST page-shell">
    <transition name="fade">
      <div v-if="feedback.message" class="toast" :class="`toast--${feedback.type}`">
        {{ feedback.message }}
      </div>
    </transition>
    <!---用户信息-->
    <div class="user card card--highlight">
      <h1>欢迎 {{ userName }}</h1>
      <p class="subtitle">今天也要好好安排待办事项喔 ✨</p>
      <div class="action-group">
        <button class="ghost-button" @click="signout">退出登录</button>
        <button class="primary-button" @click="changeWindow">更改用户信息</button>
      </div>
      <!--更改用户信息-->
      <div class="changeData" v-if="isChange">
        <input type="text" v-model="newUserName" placeholder="输入用户名" />
        <input type="password" v-model="newUserPassword" placeholder="输入新密码" />
        <p class="tips">更改密码</p>
        <div class="action-group">
          <button class="primary-button" :disabled="isUpdatingProfile" @click="changeUser">
            {{ isUpdatingProfile ? '提交中...' : '确认' }}
          </button>
          <button class="ghost-button" @click="changeWindow">取消</button>
        </div>
      </div>
    </div>

    <div class="panel-grid">
      <!---创建待办事项-->
      <div class="creatToDo card card--form">
        <h3>创建待办事项</h3>
        <p class="section-desc">填写标题、描述与截止时间，快速记录灵感。</p>
        <input type="text" v-model="todoTitle" placeholder="输入标题" />
        <input type="text" v-model="todoDescription" placeholder="输入描述" />
        <input type="datetime-local" v-model="todoEndTime" />
        <button class="primary-button" :disabled="isCreating" @click="createToDo">
          {{ isCreating ? '创建中...' : '添加待办事项' }}
        </button>
      </div>

      <!---通过ID查询待办事项-->
      <div class="queryToDo card card--form">
        <h3>通过ID查询待办事项</h3>
        <p class="section-desc">输入待办事项ID即可快速定位详情。</p>
        <input type="text" v-model="todoID" placeholder="输入ID" />
        <button class="primary-button" :disabled="isQueryingTodo" @click="queryToDo">
          {{ isQueryingTodo ? '查询中...' : '查询' }}
        </button>
      </div>
    </div>

    <div class="Todos card card--soft" v-if="isQuery">
      <div class="result-head">
        <strong>查询结果</strong>
        <span class="status-badge" :class="getStatusClass(queryResult.status)">
          {{ getStatusLabel(queryResult.status) }}
        </span>
      </div>
      <p>标题： {{ queryResult.title }}</p>
      <p>描述： {{ queryResult.description }}</p>
      <p>截止日期: {{ formatDate(queryResult.end_time) }}</p>
    </div>

    <!---通过ID更新指定待办事项-->
    <section class="update-section card card--form">
      <h3>通过ID更改指定待办事项</h3>
      <p class="section-desc">可随时调整标题、描述与截止时间。</p>
      <div class="updateTodo" v-if="isUpdate">
        <input type="text" placeholder="输入ID" v-model="Id" />
        <input type="text" placeholder="输入新的标题" v-model="newTitle" />
        <input type="text" placeholder="输入新的描述" v-model="newDescription" />
        <input type="datetime-local" v-model="newEndTime" />
        <div class="action-group">
          <button class="primary-button" :disabled="isUpdatingTodo" @click="updateToDoList()">
            {{ isUpdatingTodo ? '更新中...' : '确认' }}
          </button>
          <button class="ghost-button" @click="Toupdate">取消</button>
        </div>
      </div>
      <button class="updateWindow ghost-button" v-if="!isUpdate" @click="Toupdate">打开编辑窗口</button>
    </section>

    <hr />

    <!---待办事项列表-->
    <div class="ToDoList card card--list">
      <h3>待办事项列表</h3>
      <div class="list-toolbar">
        <input
          class="search-input"
          type="text"
          v-model.trim="keyword"
          placeholder="搜索标题或描述"
        />
        <div class="filter-group">
          <button
            class="tag-button"
            :class="{ 'tag-button--active': statusFilter === 'all' }"
            @click="statusFilter = 'all'"
          >
            全部
          </button>
          <button
            class="tag-button"
            :class="{ 'tag-button--active': statusFilter === 'pending' }"
            @click="statusFilter = 'pending'"
          >
            进行中
          </button>
          <button
            class="tag-button"
            :class="{ 'tag-button--active': statusFilter === 'done' }"
            @click="statusFilter = 'done'"
          >
            已完成
          </button>
        </div>
      </div>
      <div v-if="isListLoading" class="empty-state">
        正在加载数据，请稍候...
      </div>
      <div v-if="source.length === 0">
        <p>暂无待办事项</p>
      </div>
      <div v-else-if="filteredTodos.length === 0" class="empty-state">
        未找到匹配的待办事项
      </div>
      <div
        id="ToDoList"
        v-for="item in filteredTodos"
        :key="item.id"
        class="Todos card card--shadow"
      >
        <div class="todo-header">
          <p class="todo-title">{{ item.title }}</p>
          <span class="status-badge" :class="getStatusClass(item.status)">
            {{ getStatusLabel(item.status) }}
          </span>
        </div>
        <p class="todo-desc">{{ item.description }}</p>
        <time>截止日期: {{ formatDate(item.end_time) }}</time>
        <p>ID: {{ item.id }}</p>
        <button
          class="danger-button"
          :disabled="deletingId === item.id"
          @click="deleteToDoList(item)"
          :id="item.id"
        >
          {{ deletingId === item.id ? '删除中...' : '删除' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page-shell {
  max-width: 1100px;
  margin: 0 auto;
  padding: 32px 16px 64px;
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 24px;
  border: 1px solid rgba(15, 23, 42, 0.08);
  box-shadow: 0 15px 35px rgba(15, 23, 42, 0.08);
  padding: 24px;
  backdrop-filter: blur(8px);
}

.card--highlight {
  background: linear-gradient(135deg, #2563eb, #7c3aed);
  color: #fff;
  text-align: center;
  gap: 16px;
}

.toast {
  position: sticky;
  top: 16px;
  align-self: center;
  padding: 12px 24px;
  border-radius: 999px;
  font-weight: 600;
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.4);
  box-shadow: 0 10px 25px rgba(15, 23, 42, 0.2);
  z-index: 2;
}

.toast--success {
  background: rgba(34, 197, 94, 0.15);
  color: #14532d;
}

.toast--error {
  background: rgba(248, 113, 113, 0.15);
  color: #7f1d1d;
}

.toast--info {
  background: rgba(59, 130, 246, 0.15);
  color: #1e3a8a;
}

.card--form {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.card--soft {
  border-style: dashed;
  border-color: rgba(37, 99, 235, 0.2);
}

.card--list {
  width: 100%;
}

.card--shadow {
  border: 1px solid rgba(226, 232, 240, 0.8);
}

.LIST {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  width: 100%;
}

.user {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.subtitle {
  font-size: 16px;
  opacity: 0.9;
}

.panel-grid {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
}

.changeData {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.action-group {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  justify-content: center;
}

.section-desc {
  color: #64748b;
  font-size: 14px;
  margin-bottom: 4px;
}

input,
textarea {
  width: 100%;
  padding: 12px 16px;
  border-radius: 14px;
  border: 1px solid rgba(15, 23, 42, 0.15);
  background: rgba(248, 250, 252, 0.9);
  transition: border-color 0.2s, box-shadow 0.2s;
}

input:focus,
textarea:focus {
  border-color: #2563eb;
  outline: none;
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15);
}

.primary-button,
.ghost-button,
.danger-button,
.updateWindow {
  border: none;
  border-radius: 999px;
  padding: 12px 24px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  width: auto;
  min-width: 160px;
}

.primary-button {
  background: linear-gradient(135deg, #2563eb, #9333ea);
  color: #fff;
  box-shadow: 0 10px 20px rgba(79, 70, 229, 0.35);
}

.ghost-button {
  background: rgba(255, 255, 255, 0.15);
  color: inherit;
  border: 1px solid rgba(255, 255, 255, 0.5);
}

.danger-button {
  background: linear-gradient(135deg, #f97316, #ef4444);
  color: #fff;
}

.primary-button:hover,
.ghost-button:hover,
.danger-button:hover,
.updateWindow:hover {
  transform: translateY(-2px);
  box-shadow: 0 12px 25px rgba(15, 23, 42, 0.15);
}

.primary-button:disabled,
.ghost-button:disabled,
.danger-button:disabled,
.updateWindow:disabled {
  opacity: 0.6;
  pointer-events: none;
  box-shadow: none;
}

.tag-button {
  border: 1px solid rgba(99, 102, 241, 0.4);
  background: rgba(99, 102, 241, 0.08);
  color: #312e81;
  border-radius: 999px;
  padding: 8px 16px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tag-button--active {
  background: linear-gradient(135deg, #2563eb, #9333ea);
  color: #fff;
  border-color: transparent;
  box-shadow: 0 8px 18px rgba(79, 70, 229, 0.3);
}

.Todos {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.todo-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
}

.todo-title {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
}

.todo-desc {
  color: #475569;
}

.update-section {
  width: 100%;
}

hr {
  width: 100%;
  border: none;
  border-top: 1px dashed rgba(148, 163, 184, 0.6);
}

.ToDoList {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.ToDoList h3 {
  text-align: center;
  margin-bottom: 4px;
}

.list-toolbar {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.search-input {
  flex: 1 1 220px;
  border-radius: 999px;
}

.filter-group {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: flex-end;
}

.empty-state {
  text-align: center;
  padding: 32px;
  color: #64748b;
}

.status-badge {
  padding: 4px 12px;
  border-radius: 999px;
  font-size: 12px;
  font-weight: 600;
  line-height: 1.4;
}

.status-badge--done {
  background: rgba(34, 197, 94, 0.18);
  color: #166534;
}

.status-badge--pending {
  background: rgba(251, 191, 36, 0.18);
  color: #92400e;
}

.result-head {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-bottom: 8px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

input {
  width: 95%;
  margin: 5px auto;
}

@media (max-width: 640px) {
  .primary-button,
  .ghost-button,
  .danger-button,
  .updateWindow {
    width: 100%;
  }

  .card {
    padding: 20px;
  }
}
</style>