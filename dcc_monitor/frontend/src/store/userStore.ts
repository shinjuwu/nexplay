import type { UserLoginData } from '@/types/types.api-login'

import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { defineStore } from 'pinia'
import storage from '@/utils/storage'
import * as token from '@/utils/token'

interface UserData {
  topCode: string
  name: string
  nickName: string
  permissions: string[]
  isAdmin: boolean
}

export const useUserStore = defineStore('user', () => {
  const userFromStorage = storage.local.get(storage.keys.INFO) as UserData

  const user = reactive<UserData>({
    topCode: userFromStorage ? userFromStorage.topCode : '',
    name: userFromStorage ? userFromStorage.name : '',
    nickName: userFromStorage ? userFromStorage.nickName : '',
    permissions: userFromStorage ? userFromStorage.permissions : [],
    isAdmin: userFromStorage ? userFromStorage.isAdmin : false,
  })

  function setUser(userData: UserLoginData) {
    user.topCode = userData.top_code
    user.name = userData.username
    user.nickName = userData.nickname
    user.permissions = userData.permissions
    user.isAdmin = userData.is_admin

    saveUserToLocalStorage()
  }

  function updateUserInfo(newNickname: string) {
    user.nickName = newNickname

    saveUserToLocalStorage()
  }

  function saveUserToLocalStorage() {
    storage.local.set(storage.keys.INFO, {
      topCode: user.topCode,
      name: user.name,
      nickName: user.nickName,
      permissions: user.permissions,
      isAdmin: user.isAdmin,
    })
  }

  function hasPermission(permission: string) {
    return user.permissions.indexOf(permission) >= 0
  }

  function isAdminUser() {
    return user.isAdmin
  }

  const router = useRouter()
  function signOut() {
    token.clear()
    router.push('/login')
  }

  return {
    user,
    hasPermission,
    isAdminUser,
    setUser,
    signOut,
    updateUserInfo,
  }
})
