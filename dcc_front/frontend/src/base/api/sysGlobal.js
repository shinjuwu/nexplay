import { apiV1Request } from '@/base/utils/request'

const basePath = 'global'

export function getAgentList() {
  return apiV1Request.get(`/${basePath}/getagentlist`)
}

export function getAllGameList() {
  return apiV1Request.get(`/${basePath}/getallgamelist`)
}

export function getGameList() {
  return apiV1Request.get(`/${basePath}/getgamelist`)
}

export function getRoomTypeList() {
  return apiV1Request.get(`/${basePath}/getroomtypelist`)
}

export function getAgentPermissionList(accountType) {
  return apiV1Request.get(`/${basePath}/getagentpermissionlist?account_type=${accountType}`)
}

export function getAgentCustomTagSettingList() {
  return apiV1Request.get(`/${basePath}/getagentcustomtagsettinglist`)
}

export function getAgentAdminuserPermissionList() {
  return apiV1Request.get(`/${basePath}/getagentadminuserpermissionlist`)
}

export function getUserLoginLogCountryShortList() {
  return apiV1Request.get(`/${basePath}/getuserloginlogcountryshortlist`)
}

export function getUserLoginLogBroswerList() {
  return apiV1Request.get(`/${basePath}/getuserloginlogbroswerlist`)
}
