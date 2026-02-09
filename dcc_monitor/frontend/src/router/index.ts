import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'

import { useUserStore } from '@/store/userStore'
import * as token from '@/utils/token'

const Login = () => import('@/views/Login/Login.vue')
const Layout = () => import('@/views/Layout/Layout.vue')
const Home = () => import('@/views/Home/Home.vue')
const PersonalInfo = () => import('@/views/PersonalInfo/PersonalInfo.vue')
const MemberManagement = () => import('@/views/MemberManagement/MemberManagement.vue')
const MemberManagementCreateAccount = () => import('@/views/MemberManagement/CreateAccount.vue')
const MemberManagementEditAccount = () => import('@/views/MemberManagement/EditAccount.vue')
const MemberManagementEditPassword = () => import('@/views/MemberManagement/EditPassword.vue')
const SystemMonitor = () => import('@/views/SystemMonitor/SystemMonitor.vue')
const Error = () => import('@/views/Error/Error.vue')

const routes: Array<RouteRecordRaw> = [
  {
    name: 'login',
    path: '/login',
    component: Login,
  },
  {
    name: 'layout',
    path: '/',
    component: Layout,
    children: [
      {
        name: 'home',
        path: '/',
        component: Home,
      },
      {
        name: 'personal-info',
        path: '/personal-info',
        component: PersonalInfo,
      },
      {
        name: 'member-management',
        path: '/member-management',
        component: MemberManagement,
        meta: {
          requiresAdmin: true,
        },
      },
      {
        name: 'member-management-create-account',
        path: '/member-management/create-account',
        component: MemberManagementCreateAccount,
        meta: {
          requiresAdmin: true,
        },
      },
      {
        name: 'member-management-edit-account',
        path: '/member-management/edit-account/:account',
        component: MemberManagementEditAccount,
        meta: {
          requiresAdmin: true,
        },
      },
      {
        name: 'member-management-edit-password',
        path: '/member-management/edit-password/:account',
        component: MemberManagementEditPassword,
        meta: {
          requiresAdmin: true,
        },
      },
      {
        name: 'system-monitor',
        path: '/system-monitor/:platform',
        meta: {
          requiresPermission: true,
        },
        component: SystemMonitor,
      },
    ],
  },
  {
    name: 'error',
    path: '/error',
    component: Error,
  },
  {
    name: 'other',
    path: '/:catchAll(.*)',
    redirect: { path: '/error' },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

const whitelist = ['/login']

router.beforeEach((to, from, next) => {
  const path = to.path.toLowerCase()
  const hasToken = token.get() !== null

  // 沒有token
  if (!hasToken) {
    if (whitelist.indexOf(path) !== -1) {
      // 不需要登入的頁面
      next()
    } else {
      // 需要登入的頁面導到登入頁
      next(`/login?redirect=${path}`)
    }

    token.clear()
    return
  }

  // 登入頁面有token則導到首頁
  if (path === '/login') {
    router.push('/')
  } else {
    const { isAdminUser, hasPermission } = useUserStore()

    if (
      (to.meta.requiresAdmin && !isAdminUser()) ||
      (to.meta.requiresPermission && !hasPermission(to.params.platform as string))
    ) {
      router.push('/error')
    } else {
      next()
    }
  }
})

export default router
