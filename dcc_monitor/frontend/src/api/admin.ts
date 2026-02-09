import type { Response } from '@/types/types.api-base'
import type {
  UsersRegisterRequest,
  GetUserInfoRequest,
  ModifyUsersInfoRequest,
  ModifyUsersPasswordRequest,
} from '@/types/types.api-admin'
import { axiosInstance } from '@/api/base'

const adminPath = '/admin.api/v1'

export function register(req: UsersRegisterRequest) {
  return axiosInstance.post<Response<string>>(`${adminPath}/register`, req)
}

export function getUsersInfoList() {
  return axiosInstance.post<Response<string>>(`${adminPath}/getusersinfolist`)
}

export function getUsersInfo(req: GetUserInfoRequest) {
  return axiosInstance.post<Response<string>>(`${adminPath}/getusersinfo`, req)
}

export function modifyUsersInfo(req: ModifyUsersInfoRequest) {
  return axiosInstance.post<Response<string>>(`${adminPath}/modifyusersinfo`, req)
}

export function modifyUsersPassword(req: ModifyUsersPasswordRequest) {
  return axiosInstance.post<Response<string>>(`${adminPath}/modifyuserspassword`, req)
}
