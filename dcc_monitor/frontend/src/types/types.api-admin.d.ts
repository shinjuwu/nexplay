export interface UsersRegisterRequest {
  username: string
  password: string
  nickname: string
  permissions: string[]
}

export interface UsersRegisterResponse {
  username: string
  nickname: string
  user_metadata: string
}

export interface GetUserInfoListResponse {
  data: GetUserInfoResponse[]
}

export interface GetUserInfoResponse {
  username: string
  nickname: string
  user_metadata: string
  is_enabled: boolean
  last_login_time: string
  info: string
  create_time: string
  permissions: string[]
}

export interface GetUserInfoRequest {
  username: string
}

export interface ModifyUsersInfoRequest {
  username: string
  nickname: string
  is_enabled: boolean
  permissions: string[]
  info: string
}

export interface ModifyUsersPasswordRequest {
  username: string
  password: string
}
