import { apiV1Request } from '@/base/utils/request'

const basePath = 'cal'

// 取得指定時間區段代理時間區段的統計資料
export function getPerformanceReport(input) {
  return apiV1Request.post(`/${basePath}/getperformancereport`, input)
}

// 取得指定時間區段代理總和的資料
export function getPerformanceReportList(input) {
  return apiV1Request.post(`/${basePath}/getperformancereportlist`, input)
}
