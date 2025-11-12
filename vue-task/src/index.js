import { createRouter,createWebHistory } from "vue-router";

import Todos from "../component/Todos.vue"
import App from "./App.vue";

const router = createRouter({
    history: createWebHistory(),
    routes:[
        {
            path: `/`,
            name: "Home",
            component: App
        },
        {
            path: `/Todos`,
            name: "register",
            component: Todos
        }
    ]
})

export default router