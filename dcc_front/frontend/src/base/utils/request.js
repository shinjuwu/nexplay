import axios from 'axios'
import { get, clear } from '@/base/utils/token'
import constant from '@/base/common/constant'

export const apiV1Request = axios.create({
  baseURL: '/api/v1',
})

let customRespInterceptor
export function setAxiosInstanceConfig(router) {
  // when error response data code is ErrorCode.ErrorJWT
  // token is invalid
  // user need re-login
  // redirect to login page
  customRespInterceptor = apiV1Request.interceptors.response.use(
    (response) => {
      return response
    },
    (error) => {
      if (error.response.data.code === constant.ErrorCode.ErrorJWT) {
        // clear token
        clear()
        router.push('/Login')
        return Promise.reject(new Error('token expried '))
      } else {
        return Promise.reject(error)
      }
    }
  )

  // set token
  apiV1Request.defaults.headers.common['Dcc-Token'] = get()
}

export function clearAxiosInstanceConfig() {
  // clear customRespInterceptor
  apiV1Request.interceptors.response.eject(customRespInterceptor)

  // remove token
  delete apiV1Request.defaults.headers.common['Dcc-Token']
}
