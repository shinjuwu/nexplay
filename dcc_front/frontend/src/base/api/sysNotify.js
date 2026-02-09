import { apiV1Request } from '@/base/utils/request'

const basePath = 'notify'

export function getChatServiceConnInfo() {
  return apiV1Request.get(`/${basePath}/getchatserviceconnInfo`)
}

export function notifyBroadcastMessage(input) {
  return apiV1Request.post(`/${basePath}/notifybroadcastmessage`, input)
}
