export interface CaptchaResponse {
  captchaId: string
  picPath: string
  captchaLength: number
  expiredTime: string
}

export interface UsersLoginRequest {
  username: string
  password: string
  captcha: string
  captchaId: string
}

export interface UserLoginData {
  top_code: string
  username: string
  nickname: string
  user_metadata: string
  create_time: string
  last_login_time: string
  is_admin: boolean
  permissions: string[]
}

export interface UsersLoginResponse {
  UserData: UserLoginData
  token: string
  expiresAt: number
}
