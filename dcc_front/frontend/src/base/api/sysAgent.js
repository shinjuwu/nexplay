import { apiV1Request } from '@/base/utils/request'

const basePath = 'agent'

// 取得代理底下所有代理資料
export function getAgentList() {
  return apiV1Request.post(`/${basePath}/getagentlist`)
}

// 創建代理帳號
export function createAgent(data) {
  return apiV1Request.post(`/${basePath}/createagent`, data)
}

// 取得指定代理補分相關資料設定
export function getAgentCoinsSupplyInfo(id) {
  return apiV1Request.post(`/${basePath}/getagentcoinsupplyinfo`, id)
}

// 秘鑰資訊顯示
export function getAgentSecretKey(id) {
  return apiV1Request.post(`/${basePath}/getagentsecretkey`, id)
}

// 修改指定代理補分相關設定
export function setAgentCoinSupplyInfo(data) {
  return apiV1Request.post(`/${basePath}/setagentcoinsupplyinfo`, data)
}

// 取得代理遊戲列表
export function getAgentGameList(input) {
  return apiV1Request.post(`/${basePath}/getagentgamelist`, input)
}

// 設置代理遊戲狀態
export function setAgentGameState(input) {
  return apiV1Request.post(`/${basePath}/setagentgamestate`, input)
}

// 取得代理遊戲房間列表
export function getAgentGameRoomList(input) {
  return apiV1Request.post(`/${basePath}/getagentgameroomlist`, input)
}

// 設置代理遊戲房間狀態
export function setAgentGameRoomState(input) {
  return apiV1Request.post(`/${basePath}/setagentgameroomstate`, input)
}

// 取得代理權限群組權限樣板
export function getAgentPermissionTemplateInfo() {
  return apiV1Request.get(`/${basePath}/getagentpermissiontemplateinfo`)
}

// 取得代理權限群組列表
export function getAgentPermissionList(input) {
  return apiV1Request.post(`/${basePath}/getagentpermissionlist`, input)
}

// 取得代理權限群組列表
export function getAgentPermission(input) {
  return apiV1Request.post(`/${basePath}/getagentpermission`, input)
}

// 創建代理權限群組
export function createAgentPermission(input) {
  return apiV1Request.post(`/${basePath}/createagentpermission`, input)
}

// 修改代理權限群組
export function setAgentPermission(input) {
  return apiV1Request.post(`/${basePath}/setagentpermission`, input)
}

// 刪除代理權限群組
export function deleteAgentPermission(input) {
  return apiV1Request.post(`/${basePath}/deleteagentpermission`, input)
}

// 取得代理後台IP資訊列表
export function getAgentIpWhitelistList(input) {
  return apiV1Request.post(`/${basePath}/getagentipwhitelistlist`, input)
}

// 取得代理後台IP資訊
export function getAgentIpWhitelist(input) {
  return apiV1Request.post(`/${basePath}/getagentipwhitelist`, input)
}

// 設置代理後台IP資訊
export function setAgentIpWhitelist(input) {
  return apiV1Request.post(`/${basePath}/setagentipwhitelist`, input)
}

// 取得代理錢包餘額列表
export function getAgentWalletList(input) {
  return apiV1Request.post(`/${basePath}/getagentwalletlist`, input)
}

// 設置代理錢包餘額
export function setAgentWalletList(input) {
  return apiV1Request.post(`/${basePath}/setagentwallet`, input)
}

// 取得代理API IP資訊
export function getAgentApiIpWhitelist(input) {
  return apiV1Request.post(`/${basePath}/getagentapiipwhitelist`, input)
}

// 設置代理API IP資訊
export function setAgentApiIpWhitelist(input) {
  return apiV1Request.post(`/${basePath}/setagentapiipwhitelist`, input)
}
