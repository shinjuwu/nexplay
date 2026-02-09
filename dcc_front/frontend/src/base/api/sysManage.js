import { apiV1Request } from '@/base/utils/request'

const basePath = 'manage'

// 取得跑馬燈列表
export function getMarqueeList() {
  return apiV1Request.post(`/${basePath}/getmarqueelist`)
}

// 取得跑馬燈設定
export function getMarquee(input) {
  return apiV1Request.post(`/${basePath}/getmarquee`, input)
}

// 新增跑馬燈
export function createMarquee(input) {
  return apiV1Request.post(`/${basePath}/createmarquee`, input)
}

// 刪除跑馬燈
export function deleteMarquee(input) {
  return apiV1Request.post(`/${basePath}/deletemarquee`, input)
}

// 更新跑馬燈
export function updateMarquee(input) {
  return apiV1Request.post(`/${basePath}/updatemarquee`, input)
}

// 取得今日排行榜
export function getGameLeaderBoards(input) {
  return apiV1Request.get(
    `/${basePath}/getgameleaderboards?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 取得今日前100名風險玩家清單
export function getRiskUserList(input) {
  return apiV1Request.get(
    `/${basePath}/getriskuserlist?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 取得當天資訊總覽
export function getStatData(input) {
  return apiV1Request.get(
    `/${basePath}/getstatdata?level_code=${input.levelCode}&date_type=${input.timeType}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 最近一段時間(預設30日)的活躍玩家數&日投注人數
export function getIntervalDaysBettorData(input) {
  return apiV1Request.get(
    `/${basePath}/getintervaldaysbettordata?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 今日各時段輸贏(昨日&今日)
export function getIntervalTotalScoreData(input) {
  return apiV1Request.get(
    `/${basePath}/getintervaltotalscoredata?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 今日各時段投注人數
export function getIntervalTotalBettorInfoData(input) {
  return apiV1Request.get(
    `/${basePath}/getintervaltotalbettorinfodata?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

export function getIntervalRealTimeUserData(input) {
  return apiV1Request.get(
    `/${basePath}/getintervalrealtimeuserdata?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 取得後台公告列表
export function getAnnouncementList() {
  return apiV1Request.post(`${basePath}/getannouncementlist`)
}

// 取得某筆後台公告資訊
export function getAnnouncement(input) {
  return apiV1Request.post(`${basePath}/getannouncement`, input)
}

// 新增後台公告
export function createAnnouncement(input) {
  return apiV1Request.post(`${basePath}/createannouncement`, input)
}

// 編輯後台公告
export function updateAnnouncement(input) {
  return apiV1Request.post(`${basePath}/updateannouncement`, input)
}

// 刪除後台公告
export function deleteAnnouncement(input) {
  return apiV1Request.post(`${basePath}/deleteannouncement`, input)
}

// 取得維護頁設定
export function getMaintainPageSetting(input) {
  return apiV1Request.post(`${basePath}/getmaintainpagesetting`, input)
}

// 設定維護頁設定
export function setMaintainPageSetting(input) {
  return apiV1Request.post(`${basePath}/setmaintainpagesetting`, input)
}

// 取得今日用戶登入裝置使用比例
export function getUserLoginDeviceUsageRatioData(input) {
  return apiV1Request.get(
    `/${basePath}/getuserlogindeviceusageratiodata?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}

// 取得用戶登入來源位置排行
export function getUserLoginSourceLocRankingData(input) {
  return apiV1Request.get(
    `/${basePath}/getuserloginsourcelocrankingdata?level_code=${input.levelCode}&time_zone=${input.timeZone}&is_search_all=${input.isSearchAll}`
  )
}
