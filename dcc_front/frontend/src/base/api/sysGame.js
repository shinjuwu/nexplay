import { apiV1Request } from '@/base/utils/request'

const basePath = 'game'

// 取得遊戲列表
export function getGameList() {
  return apiV1Request.get(`/${basePath}/getgamelist`)
}

// 修改遊戲狀態
export function setGameState(input) {
  return apiV1Request.post(`/${basePath}/setgamestate`, input)
}

// 取得遊戲server狀態
export function getGameServerState() {
  return apiV1Request.get(`/${basePath}/getgameserverstate`)
}

// 設置遊戲server狀態
export function setGameServerState(input) {
  return apiV1Request.post(`/${basePath}/setgameserverstate`, input)
}

// 創建更新遊戲相關設定(遊戲server維護中才可以使用)
export function notifyGameServer(input) {
  return apiV1Request.post(`/${basePath}/notifygameserver`, input)
}

// 取得遊戲icon list
export function getGameIconList() {
  return apiV1Request.get(`/${basePath}/getgameiconlist`)
}

// 取得遊戲icon list
export function setGameIconList(input) {
  return apiV1Request.post(`/${basePath}/setgameiconlist`, input)
}

// 取得遊戲icon list
export function getCannedList(input) {
  return apiV1Request.post(`/${basePath}/getcannedlist`, input)
}

// 取得遊戲icon list
export function setCannedList(input) {
  return apiV1Request.post(`/${basePath}/setcannedlist`, input)
}
