import type { Response } from '@/types/types.api-base'
import type { UsersLoginRequest } from '@/types/types.api-login'
import { axiosInstance } from '@/api/base'

const loginPath = '/login.api/v1'

export function captcha() {
  return axiosInstance.post<Response<string>>(`${loginPath}/captcha`)
}

export function login(req: UsersLoginRequest) {
  return axiosInstance.post<Response<string>>(`${loginPath}/login`, req)
}
