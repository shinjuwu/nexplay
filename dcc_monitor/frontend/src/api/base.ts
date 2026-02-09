import type { Router } from 'vue-router'
import type { AxiosError, AxiosResponse } from 'axios'
import type { Response } from '@/types/types.api-base'
import type { DialogShowFunc } from '@/types/types.dialog'

import axios from 'axios'
import { responseCode } from '@/common/constant'
import * as token from '@/utils/token'

export const axiosInstance = axios.create({
  baseURL: '/monitor',
})

let customRespInterceptor: number
export function setAxiosInstanceConfig(router: Router) {
  // when error response data code is ErrorCode.ErrorJWT
  // token is invalid
  // user need re-login
  // redirect to login page
  customRespInterceptor = axiosInstance.interceptors.response.use(
    (response) => {
      return response
    },
    (error: AxiosError<Response<string>>) => {
      if (error.response?.data.code === responseCode.Unauth) {
        // clear token
        token.clear()
        router.push('/login')
      } else {
        return Promise.reject(error)
      }
    }
  )

  // set token
  axiosInstance.defaults.headers.common['Token'] = token.get()
}

export function clearAxiosInstanceConfig() {
  // clear customRespInterceptor
  axiosInstance.interceptors.response.eject(customRespInterceptor)

  // remove token
  delete axiosInstance.defaults.headers.common['Token']
}

export async function processApiRequest(
  apiHandler: () => void,
  warn: DialogShowFunc,
  startFunc?: () => void,
  endFunc?: () => void
) {
  startFunc?.()

  try {
    await apiHandler()
  } catch (err) {
    console.error(err)

    if (axios.isAxiosError<Response<string>>(err)) {
      warn(parseServerErrorMessage(err.response))
    }
  }

  endFunc?.()
}

export function parseServerErrorMessage(response?: AxiosResponse<Response<string>>) {
  const errorCode: number | string = (response?.data.code as number) || 'network error'
  const errorMsg = (response?.data.data as string) || '伺服器发生错误，请重新尝试或联系客服'
  return `错误: ${errorMsg} (${errorCode}) `
}
