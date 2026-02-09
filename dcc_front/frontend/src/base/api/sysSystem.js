import { apiV1Request } from '@/base/utils/request'

const basePath = 'system'

// 此接口用來取得server相關設定的參數
export function getServerSetting() {
  return apiV1Request.post(`/${basePath}/getserversetting`)
}

// 取得匯率設定
export function getExchangeDataSetting() {
  return apiV1Request.post(`/${basePath}/getexchangedatalist`)
}

// 設定匯率設定
export function setExchangeDataSetting(input) {
  return apiV1Request.post(`/${basePath}/setexchangedatalist`, input)
}
