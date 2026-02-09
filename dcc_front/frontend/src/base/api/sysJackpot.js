import { apiV1Request } from '@/base/utils/request'

const basePath = 'jackpot'

// 建立JACKPOT代幣
export function createJackpotToken(input) {
  return apiV1Request.post(`/${basePath}/createjackpottoken`, input)
}

// 取得總代理JACKPOT設定列表
export function getAgentJackpotList(input) {
  return apiV1Request.post(`/${basePath}/getagentjackpotlist`, input)
}

// 取得指定總代理JACKPOT設定
export function getAgentJackpot(input) {
  return apiV1Request.post(`/${basePath}/getagentjackpot`, input)
}

// 取得JACKPOT玩家貢獻度
export function getJackpotLeaderBoard(input) {
  return apiV1Request.post(`/${basePath}/getjackpotleaderboard`, input)
}

// 取得JACKPOT紀錄列表
export function getJackpotList(input) {
  return apiV1Request.post(`/${basePath}/getjackpotlist`, input)
}

// 取得JACKPOT獎池資訊
export function getJackpotPoolData() {
  return apiV1Request.get(`/${basePath}/getjackpotpooldata`)
}

// 取得平台JACKPOT設定
export function getJackpotSetting() {
  return apiV1Request.get(`/${basePath}/getjackpotsetting`)
}

// 取得JACKPOT代幣紀錄列表
export function getJackpotTokenList(input) {
  return apiV1Request.post(`/${basePath}/getjackpottokenlist`, input)
}

// SERVER同步JACKPOT資訊
export function notifyGameServerAgentJackpotInfo(input) {
  return apiV1Request.post(`/${basePath}/notifygameserveragentjackpotinfo`, input)
}

// 設定總代理JACKPOT
export function setAgentJackpot(input) {
  return apiV1Request.post(`/${basePath}/setagentjackpot`, input)
}

// 設定平台JACKPOT設定
export function setJackpotSetting(input) {
  return apiV1Request.post(`/${basePath}/setjackpotsetting`, input)
}
