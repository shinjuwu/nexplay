import { apiV1Request } from '@/base/utils/request'

const basePath = 'user'

// ping
export function ping() {
  return apiV1Request.get(`/${basePath}/ping`)
}

// 依照查詢者角色權限列出遊戲會員帳號清單
export function getGameUsers() {
  return apiV1Request.post(`/${basePath}/getgameusers`)
}

// 指定查詢某遊戲會員帳號信息
export function getGameUsersInfo(input) {
  return apiV1Request.post(`/${basePath}/getgameuserinfo`, input)
}

// 指定修改某遊戲會員帳號信息
export function updateGameUserInfo(data) {
  return apiV1Request.post(`/${basePath}/updategameuserinfo`, data)
}

// 依照查詢者角色權限列出自身權限下的子帳號列表
export function getAdminUsers() {
  return apiV1Request.post(`/${basePath}/getadminusers`)
}

// 指定查詢某後台帳號狀態
export function getAdminUserInfo(username) {
  return apiV1Request.post(`/${basePath}/getadminuserinfo`, username)
}

// 指定設定某後台帳號狀態
export function updateAdminUserInfo(data) {
  return apiV1Request.post(`/${basePath}/updateadminuserinfo`, data)
}

// 創建後台帳號
export function createAdminUser(data) {
  return apiV1Request.post(`/${basePath}/createadminuser`, data)
}

// 取得玩家錢包餘額列表
export function getGameUserWalletList(input) {
  return apiV1Request.post(`/${basePath}/getgameuserwalletlist`, input)
}

// 設置玩家錢包餘額
export function setGameUserWallet(input) {
  return apiV1Request.post(`/${basePath}/setgameuserwallet`, input)
}

// 修改個人資訊
export function setPersonalInfo(input) {
  return apiV1Request.post(`/${basePath}/setpersonalinfo`, input)
}

// 修改個人密碼
export function setPersonalPassword(input) {
  return apiV1Request.post(`/${basePath}/setpersonalpassword`, input)
}

// 重置密碼
export function resetPassword(input) {
  return apiV1Request.post(`/${basePath}/resetpassword`, input)
}

// 取得玩家目前餘額
export function getGameUserBalance(input) {
  return apiV1Request.post(`/${basePath}/getgameuserbalance`, input)
}

// 取得玩家目前餘額
export function getGameUserPlayCountData(input) {
  return apiV1Request.post(`/${basePath}/getgameuserplaycountdata`, input)
}

// 取得玩家登入資訊列表
export function getGameUserLoginData(input) {
  return apiV1Request.post(`${basePath}/getgameuserlogindata`, input)
}
