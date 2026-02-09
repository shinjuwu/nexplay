import { createRouter, createWebHistory } from 'vue-router'

const Login = () => import('@/base/views/BaseLogin.vue')
const Home = () => import('@/manager/views/Home/ManagerHome.vue')
const routes = [
  {
    name: 'Login',
    path: '/Login',
    component: Login,
  },
  {
    name: 'Home',
    path: '/',
    component: Home,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

export default router
