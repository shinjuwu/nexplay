import { reactive } from 'vue'
import { defineStore } from 'pinia'
import constant from '@/base/common/constant'
import { roleItems } from '@/base/common/menuConstant'
import * as request from '@/base/utils/request'
import storage from '@/base/utils/storage'
import * as token from '@/base/utils/token'
import { isAfter, isBefore } from 'date-fns'

export const useUserStore = defineStore('user', () => {
  const userFromStorage = storage.local.get(storage.keys.INFO)

  const user = reactive({
    agentId: userFromStorage ? userFromStorage.agentId : 0,
    name: userFromStorage ? userFromStorage.name : '',
    nickName: userFromStorage ? userFromStorage.nickName : '',
    accountType: userFromStorage ? userFromStorage.accountType : 0,
    permissions: userFromStorage ? userFromStorage.permissions : [],
    roles: userFromStorage ? setRoles(userFromStorage.permissions) : setRoles([]),
    levelCode: userFromStorage ? userFromStorage.levelCode : '',
    cooperation: userFromStorage ? userFromStorage.cooperation : 0,
    currency: userFromStorage ? userFromStorage.currency : '',
    jackpotStartTime: userFromStorage
      ? new Date(userFromStorage.jackpotStartTime)
      : new Date('1970-01-01T00:00:00.000Z'),
    jackpotEndTime: userFromStorage ? new Date(userFromStorage.jackpotEndTime) : new Date('1970-01-01T00:00:00.000Z'),
    walletType: userFromStorage ? userFromStorage.walletType : 0,
  })

  function setUser(input) {
    user.agentId = input.agent_id
    user.name = input.username
    user.nickName = input.nickname
    user.accountType = input.account_type
    user.permissions = input.permission
    user.cooperation = input.cooperation
    user.roles = setRoles(input.permission)
    user.levelCode = input.level_code
    user.currency = input.currency
    user.jackpotStartTime = new Date(input.jackpot_start_time)
    user.jackpotEndTime = new Date(input.jackpot_end_time)
    user.walletType = input.wallet_type

    saveUserToLocalStorage()
  }

  function updateUserInfo(newNickname) {
    user.nickName = newNickname

    saveUserToLocalStorage()
  }

  function saveUserToLocalStorage() {
    storage.local.set(storage.keys.INFO, {
      agentId: user.agentId,
      name: user.name,
      nickName: user.nickName,
      accountType: user.accountType,
      permissions: user.permissions,
      levelCode: user.levelCode,
      cooperation: user.cooperation,
      currency: user.currency,
      jackpotStartTime: user.jackpotStartTime,
      jackpotEndTime: user.jackpotEndTime,
      walletType: user.walletType,
    })
  }

  function isInRole(roleKey) {
    return user.roles[roleKey]
  }

  function setRoles(permissions) {
    return roleItems.reduce((roles, roleItem) => {
      roles[roleItem.key] = roleItem.permissions.every((r) => permissions.find((fc) => fc === r))
      return roles
    }, {})
  }

  function isAdminUser() {
    return user.accountType === constant.AccountType.Admin
  }

  function isAgentTransferUser() {
    return user.walletType === constant.AgentWallet.Transfer
  }

  function isAgentSingleWalletUser() {
    return user.walletType === constant.AgentWallet.Single
  }

  function signOut(router) {
    token.clear()
    request.clearAxiosInstanceConfig()
    router.push('/Login')
  }

  function isJackpotOpen() {
    const today = new Date()
    return !isBefore(today, user.jackpotStartTime) && !isAfter(today, user.jackpotEndTime)
  }

  return {
    user,
    isInRole,
    isAdminUser,
    isAgentSingleWalletUser,
    isAgentTransferUser,
    isJackpotOpen,
    setUser,
    signOut,
    updateUserInfo,
  }
})
