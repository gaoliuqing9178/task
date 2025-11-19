<script setup>
import { reactive, ref } from 'vue'
import { RouterView } from 'vue-router'
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'

const BaseURL = `/api/v1`
const router = useRouter()

const accountName = ref('')
const accountPassword = ref('')
const userName = ref('')

const isRegister = ref(false)
const isSignin = ref(true)
const isUser = ref(false)
const rememberAccount = ref(false)
const authMessage = ref({
  type: '',
  text: '',
})
const authLoading = reactive({
  signin: false,
  register: false,
})
let authTimer

function setAuthMessage(type, text) {
  authMessage.value = { type, text }
  clearTimeout(authTimer)
  if (text) {
    authTimer = setTimeout(() => {
      authMessage.value = { type: '', text: '' }
    }, 4000)
  }
}

//注册
async function Register() {
  if (authLoading.register) return
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

    //检查输入是否为空
    if (accountName.value === '' || accountPassword.value === '' || userName.value === '') {
      setAuthMessage('error', '请输入完整的账户、密码与用户名')
      return
    }

    authLoading.register = true
    const res = await fetch(req)
    // 如果请求失败，抛出错误
    if (res.status === 500) {
      setAuthMessage('error', '用户已存在，请尝试直接登录')
    }

    const data = await res.json()
    console.log(res)
    if (!res.ok) {
      setAuthMessage('error', data.message || '注册失败，请稍后再试')
      return
    }
    setAuthMessage('success', data.message || '注册成功，快去登录吧')
    //清空输入
    userName.value = ``
    accountName.value = ``
    accountPassword.value = ``
    if (!isSignin.value) {
      RegisterWindow()
    }
  } catch (error) {
    console.log(error)
    setAuthMessage('error', '注册失败，请检查网络后重试')
  } finally {
    authLoading.register = false
  }
}

//登录
async function Signin() {
  if (authLoading.signin) return
  try {
    const req = new Request(`${BaseURL}/user/login`, {
      method: 'POST',
      body: JSON.stringify({
        account: accountName.value,
        password: accountPassword.value,
      }),
    })

    //检查输入是否为空
    if (accountName.value === '' || accountPassword.value === '') {
      setAuthMessage('error', '请输入账户与密码')
      return
    }

    authLoading.signin = true
    const res = await fetch(req)
    if (res.status === 401) {
      setAuthMessage('error', '用户名或密码错误')
    }

    const data = await res.json()
    if (!res.ok) {
      setAuthMessage('error', data.message || '登录失败，请稍后再试')
      return
    }
    setAuthMessage('success', '欢迎回来，已成功登录')

    //获取token&userName
    userName.value = data.user.name
    localStorage.setItem('userToken', data.token)
    localStorage.setItem('userName', data.user.name)
    if (rememberAccount.value) {
      localStorage.setItem('savedAccount', accountName.value)
    } else {
      localStorage.removeItem('savedAccount')
    }

    //关闭登陆页面
    isSignin.value = !isSignin.value
    isUser.value = !isUser.value

    router.push('/Todos')
  } catch (error) {
    console.log(error)
    setAuthMessage('error', '登录失败，请检查网络后重试')
  } finally {
    authLoading.signin = false
  }
}

function RegisterWindow() {
  isRegister.value = !isRegister.value
  isSignin.value = !isSignin.value
}

//若本地存在token则自动登录
onMounted(() => {
  const savedToken = localStorage.getItem('userToken')
  const savedAccount = localStorage.getItem('savedAccount')
  if (savedAccount) {
    accountName.value = savedAccount
    rememberAccount.value = true
  }
  if (savedToken) {
    isUser.value = !isUser.value
    router.push('/Todos')
  }
})

</script>
<!------------------------------------------------------------------------------------------------------>
<template>
  <div class="app-shell">
    <div class="Sign&amp;register" v-if="!isUser">
      <section class="hero-panel glass-card">
        <p class="eyebrow">Smart Todo</p>
        <h1>让工作与生活保持从容</h1>
        <p class="hero-text">使用轻盈的任务看板，跨设备同步、随时掌控进度。</p>
        <ul class="hero-highlights">
          <li>✓ 实时同步</li>
          <li>✓ 灵活提醒</li>
          <li>✓ 安全登录</li>
        </ul>
      </section>

      <div class="auth-card glass-card">
        <transition name="fade">
          <p
            v-if="authMessage.text"
            class="form-feedback"
            :class="`form-feedback--${authMessage.type}`"
            aria-live="assertive"
          >
            {{ authMessage.text }}
          </p>
        </transition>
        <!---登录-->
        <div class="sign" v-if="isSignin">
          <div class="card-head">
            <h3>立即登录</h3>
            <p>使用账户信息进入个人待办中心。</p>
          </div>
          <input class="ac" type="text" v-model="accountName" placeholder="输入账户名称" />
          <input class="ac" type="password" v-model="accountPassword" placeholder="输入密码" />
          <label class="remember-row">
            <input type="checkbox" v-model="rememberAccount" />
            <span>记住账户</span>
          </label>
          <button class="acbtn primary" :disabled="authLoading.signin" @click="Signin">
            {{ authLoading.signin ? '登录中...' : '登录' }}
          </button>
          <button class="acbtn ghost" :disabled="authLoading.signin" @click="RegisterWindow">
            注册新账号
          </button>
        </div>
        <!---注册-->
        <div class="register" v-if="isRegister">
          <div class="card-head">
            <h3>创建账号</h3>
            <p>填写简单信息，即可拥有专属待办空间。</p>
          </div>
          <input class="ac" type="text" v-model="accountName" placeholder="输入账户名称" />
          <input class="ac" type="password" v-model="accountPassword" placeholder="输入密码" />
          <input class="ac" type="text" v-model="userName" placeholder="输入用户名称" />
          <button class="acbtn primary" :disabled="authLoading.register" @click="Register">
            {{ authLoading.register ? '注册中...' : '注册' }}
          </button>
          <button class="acbtn ghost" :disabled="authLoading.register" @click="RegisterWindow">
            返回登陆
          </button>
        </div>
      </div>
    </div>
    <RouterView></RouterView>
  </div>
</template>

<style>
:root {
  font-family: 'Inter', 'Segoe UI', 'PingFang SC', sans-serif;
  color: #0f172a;
  background-color: #030712;
}

body {
  margin: 0;
  min-height: 100vh;
  background: radial-gradient(circle at 20% 20%, rgba(59, 130, 246, 0.25), transparent 45%),
    radial-gradient(circle at 80% 0%, rgba(236, 72, 153, 0.2), transparent 40%),
    #030712;
  color: #0f172a;
}

.app-shell {
  min-height: 100vh;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 48px 16px 80px;
  gap: 32px;
}

.Sign\&register {
  width: min(1100px, 100%);
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
  justify-content: center;
  align-items: stretch;
}

.glass-card {
  border-radius: 28px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  box-shadow: 0 25px 50px rgba(15, 23, 42, 0.45);
  backdrop-filter: blur(18px);
}

.hero-panel {
  flex: 1 1 320px;
  padding: 32px;
  background: rgba(15, 23, 42, 0.75);
  color: #e2e8f0;
  display: flex;
  flex-direction: column;
  gap: 12px;
  justify-content: center;
}

.eyebrow {
  text-transform: uppercase;
  letter-spacing: 0.35em;
  font-size: 12px;
  color: #93c5fd;
}

.hero-panel h1 {
  font-size: clamp(28px, 4vw, 40px);
  margin: 0;
}

.hero-text {
  margin: 0;
  color: rgba(226, 232, 240, 0.85);
  line-height: 1.6;
}

.hero-highlights {
  list-style: none;
  padding: 0;
  margin: 12px 0 0;
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  color: #bae6fd;
  font-weight: 600;
}

.hero-highlights li {
  background: rgba(15, 23, 42, 0.55);
  border-radius: 999px;
  padding: 8px 16px;
}

.auth-card {
  flex: 1 1 360px;
  padding: 32px;
  background: rgba(255, 255, 255, 0.98);
  display: flex;
  border-radius: 24px;
  flex-direction: column;
  gap: 16px;
}

.form-feedback {
  border-radius: 14px;
  padding: 12px 16px;
  font-weight: 600;
  text-align: center;
}

.form-feedback--success {
  background: rgba(34, 197, 94, 0.15);
  color: #166534;
}

.form-feedback--error {
  background: rgba(248, 113, 113, 0.15);
  color: #7f1d1d;
}

.form-feedback--info {
  background: rgba(59, 130, 246, 0.12);
  color: #1d4ed8;
}

.sign,
.register {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 14px;
  animation: fadeSlide 0.35s ease;
}

.remember-row {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #475569;
}

.remember-row input {
  width: auto;
  accent-color: #6366f1;
}

.card-head h3 {
  margin: 0;
  font-size: 28px;
}

.card-head p {
  margin: 4px 0 12px;
  color: #475569;
}

.ac {
  width: 90%;
  padding: 14px 16px;
  border-radius: 16px;
  border: 1px solid rgba(148, 163, 184, 0.6);
  background: #f8fafc;
  font-size: 15px;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.ac:focus {
  outline: none;
  border-color: #6366f1;
  box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.15);
}

.acbtn {
  border: none;
  border-radius: 16px;
  padding: 14px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  width: 100%;
}

.acbtn.primary {
  background: linear-gradient(135deg, #2563eb, #a855f7);
  color: #fff;
  box-shadow: 0 10px 25px rgba(79, 70, 229, 0.35);
}

.acbtn.ghost {
  background: rgba(99, 102, 241, 0.08);
  color: #1e1b4b;
  border: 1px solid rgba(99, 102, 241, 0.3);
}

.acbtn:hover {
  transform: translateY(-2px);
}

.acbtn:active {
  transform: scale(0.98);
}

.acbtn:disabled {
  opacity: 0.6;
  pointer-events: none;
  transform: none;
  box-shadow: none;
}

@keyframes fadeSlide {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@media (max-width: 768px) {
  .auth-card,
  .hero-panel {
    padding: 24px;
  }

  .Sign\&register {
    gap: 16px;
  }
}
</style>
