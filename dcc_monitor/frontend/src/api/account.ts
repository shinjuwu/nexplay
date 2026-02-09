import type { Response } from '@/types/types.api-base'
import type { ModifyInfoRequest, ModifyPasswordRequest } from '@/types/types.api-account'
import { axiosInstance } from '@/api/base'

const accountPath = '/account.api/v1'

export function modifyinfo(req: ModifyInfoRequest) {
  return axiosInstance.post<Response<string>>(`${accountPath}/modifyinfo`, req)
}

export function modifypassword(req: ModifyPasswordRequest) {
  return axiosInstance.post<Response<string>>(`${accountPath}/modifypassword`, req)
}

export function verifyToken() {
  return axiosInstance.post<Response<string>>(`${accountPath}/verifytoken`)
}
