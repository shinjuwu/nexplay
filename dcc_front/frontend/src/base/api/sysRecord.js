import { apiV1Request } from '@/base/utils/request'

const basePath = 'record'

// 取得個人遊戲紀錄列表
export function getUserPlayLogList(input) {
  return apiV1Request.post(`/${basePath}/getuserplayloglist`, input)
}

// 分批取得個人遊戲紀錄列表
export function getBatchUserPlayLogList(input) {
  return apiV1Request.post(`/${basePath}/getbatchuserplayloglist`, input)
}

// 取得遊戲局記錄
export function getPlayLogCommon(input) {
  return apiV1Request.post(`/${basePath}/getplaylogcommon`, input)
}

// 取得玩家分數紀錄列表
export function getWalletLedgerList(input) {
  return apiV1Request.post(`/${basePath}/getwalletledgerlist`, input)
}

// 更新玩家分數紀錄狀態
export function confirmWalletLedger(input) {
  return apiV1Request.post(`/${basePath}/confirmwalletledger`, input)
}

// 取得代理分數紀錄列表
export function getAgentWalletLedgerList(input) {
  return apiV1Request.post(`/${basePath}/getagentwalletledgerlist`, input)
}

// 取得後台操作紀錄列表
export function getBackendActionLogList(input) {
  return apiV1Request.post(`/${basePath}/getbackendactionloglist`, input)
}

// 取得後台登入紀錄列表
export function getBackendLoginLogList(input) {
  return apiV1Request.post(`/${basePath}/getbackendloginloglist`, input)
}

// 取得自動風控紀錄
export function getAutoRiskControlLogList(input) {
  return apiV1Request.post(`/${basePath}/getautoriskcontrolloglist`, input)
}

// 取得日結算報表列表
export function getAgentGameRatioStatList(input) {
  return apiV1Request.post(`/${basePath}/getagentgameratiostatlist`, input)
}

// 取得玩家遊玩紀錄列表
export function getGameUsersStatHourList(input) {
  return apiV1Request.post(`/${basePath}/getgameusersstathourlist`, input)
}

// 取得好友房建房紀錄列表
export function getFriendRoomLogList(input) {
  return apiV1Request.post(`/${basePath}/getfriendroomloglist`, input)
}

// 取得玩家帳變紀錄列表
export function getUserCreditLogList(input) {
  return apiV1Request.post(`/${basePath}/getusercreditloglist`, input)
}
