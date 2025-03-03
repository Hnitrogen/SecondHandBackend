import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'



const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home
        },
        {
            path: '/login',
            name: 'Login',
            component: () => import('../views/Login.vue')
        },
        {
            path: '/register',
            name: 'Register',
            component: () => import('../views/Register.vue')
        },
        {
            path: '/personal',
            name: 'Personal',
            component: () => import('../views/Personal.vue')
        },
        {
            path: '/sell-stuff',
            name: 'sellStuff',
            component: () => import('../components/SellStuffForm.vue')
        }
    ]
})

export default router 