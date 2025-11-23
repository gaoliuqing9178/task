import { createRouter,createWebHistory } from "vue-router";

import Todos from "../component/Todos.vue"
import login from "../component/login.vue";

const router = createRouter({
    history: createWebHistory(),
    routes:[
        {
            path: `/`,
            name: "login",
            component: login
        },
        {
            path: `/Todos`,
            name: "register",
            component: Todos
        }
    ]
})

export default router