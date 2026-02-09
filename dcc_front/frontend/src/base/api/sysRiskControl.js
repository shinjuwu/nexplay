import { apiV1Request } from '@/base/utils/request'

const basePath = 'riskcontrol'

// 取得平台機率設定列表
export function getIncomeRatioList(input) {
  return apiV1Request.post(`/${basePath}/getincomeratiolist`, input)
}

// 取得指定ID平台機率設定資料
export function getIncomeRatio(input) {
  return apiV1Request.post(`/${basePath}/getincomeratio`, input)
}

// 設定指定ID平台機率設定
export function setIncomeRatio(input) {
  return apiV1Request.post(`/${basePath}/setincomeratio`, input)
}

// 取得代理設定機率&遊戲輸贏結果
export function getAgentIncomeRatioAndGameData(input) {
  return apiV1Request.post(`/${basePath}/getagentincomeratioandgamedata`, input)
}

// 取得玩家設定機率&遊戲輸贏結果
export function getPlayerIncomeRatioAndGameData(input) {
  return apiV1Request.post(`/${basePath}/getplayerincomeratioandgamedata`, input)
}

// 取得當前總代理風控設定資料列表
export function getAgentIncomeRatioList(input) {
  return apiV1Request.post(`/${basePath}/getagentincomeratiolist`, input)
}

// 取得指定id總代理風控設定資料
export function getAgentIncomeRatio(input) {
  return apiV1Request.post(`/${basePath}/getagentincomeratio`, input)
}

// 設定指定id總代理風控設定資料
export function setAgentIncomeRatio(input) {
  return apiV1Request.post(`/${basePath}/setagentincomeratio`, input)
}

// 取得指定id代理玩家客製化標籤資料
export function getAgentCustomTagSettingList() {
  return apiV1Request.post(`/${basePath}/getagentcustomtagsettinglist`)
}

// 設定指定id代理玩家客製化標籤資料
export function setAgentCustomTagSettingList(input) {
  return apiV1Request.post(`/${basePath}/setagentcustomtagsettinglist`, input)
}

// 取得玩家標示資料列表
export function getGameUsersCustomTagList(input) {
  return apiV1Request.post(`/${basePath}/getgameuserscustomtaglist`, input)
}

// 取得標示玩家資料
export function getGameUsersCustomTag(input) {
  return apiV1Request.post(`/${basePath}/getgameuserscustomtag`, input)
}

// 設定標示玩家資料
export function setGameUsersCustomTag(input) {
  return apiV1Request.post(`/${basePath}/setgameuserscustomtag`, input)
}

// 取得自動風控設定
export function getAutoRiskControlSetting() {
  return apiV1Request.post(`/${basePath}/getautoriskcontrolsetting`)
}

// 設定自動風控設定
export function setAutoRiskControlSetting(input) {
  return apiV1Request.post(`/${basePath}/setautoriskcontrolsetting`, input)
}

// 取得指定玩家處置設定
export function getGameUserRiskControlTag(input) {
  return apiV1Request.post(`/${basePath}/getgameuserriskcontroltag`, input)
}

// 設定指定玩家處置設定
export function setGameUserRiskControlTag(input) {
  return apiV1Request.post(`/${basePath}/setgameuserriskcontroltag`, input)
}

// 取得遊戲基礎設定
export function getGameSetting() {
  return apiV1Request.post(`/${basePath}/getgamesetting`)
}

// 設定遊戲基礎設定
export function setGameSetting(input) {
  return apiV1Request.post(`/${basePath}/setgamesetting`, input)
}

// 取得遊戲即時資訊
export function getRealTimeGameRatio(input) {
  return apiV1Request.post(`/${basePath}/getrealtimegameratio`, input)
}

// 批次設定殺數設定資料
export function setIncomeRatios(input) {
  return apiV1Request.post(`/${basePath}/setincomeratios`, input)
}
