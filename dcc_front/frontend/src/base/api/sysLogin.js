import { apiV1Request } from '@/base/utils/request'

const basePath = 'login'

// 取得驗證碼
export function captcha() {
  return apiV1Request.post(`/${basePath}/captcha`)
}

// 用戶登錄
export function login(data) {
  return apiV1Request.post(`/${basePath}/login`, data)
}
