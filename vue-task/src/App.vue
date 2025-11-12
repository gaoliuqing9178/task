<script setup>
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import { useRouter } from 'vue-router'

const BaseURL = `/api/v1`
const router = useRouter()

const accountName = ref('')
const accountPassword = ref('')
const userName = ref('')

const isRegister = ref(false)
const isSignin = ref(true)
const isUser = ref(false)

let token = ref(``)

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
    } else if (res.status === 500) {
      alert(`Username already exists`)
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
    localStorage.setItem('userToken', data.token)
    localStorage.setItem('userName', data.user.name)

    //关闭登陆页面
    isSignin.value = !isSignin.value
    isUser.value = !isUser.value

    router.push('/Todos')
  } catch (error) {
    alert(`Account or password is wrong`)
    console.log(error)
  }
}
function RegisterWindow() {
  isRegister.value = !isRegister.value
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
      <button class="acbtn" @click="Signin">Sign in</button>
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
  <RouterView></RouterView>
</template>

<style>
html,
body {
  display: flex;
  justify-content: center;
  min-height: 100%;
  width: 100%;
}

.changeData button {
  margin: 10px;
  width: 20%;
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

.sign button,
.register button {
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
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

button {
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  transition: all 0.15s ease-in-out;
}

button:hover {
  transform: scale(1.05);
}

button:active {
  transform: scale(0.95);
}

@media (max-width: 610px) {
  .Sign\&register button {
    width: 50%;
    font-size: 10px;
}
}
</style>
