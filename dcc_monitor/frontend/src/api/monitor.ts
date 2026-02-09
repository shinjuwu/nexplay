import type { Response } from '@/types/types.api-base'
import type {
  CoinInOutStatusRequest,
  AbnormalWinAndLoseStatusRequest,
  PlatformRTPStatusRequest,
  ServiceStatusRequest,
} from '@/types/types.api-monitor'
import { axiosInstance } from '@/api/base'

const monitorPath = '/monitor.api/v1'

export function coinInOutStatus(req: CoinInOutStatusRequest) {
  return axiosInstance.post<Response<string>>(`${monitorPath}/coininoutstatus`, req)
}

export function abnormalWinAndLoseStatus(req: AbnormalWinAndLoseStatusRequest) {
  return axiosInstance.post<Response<string>>(`${monitorPath}/abnormalwinandlosestatus`, req)
}

export function platformRtpStatus(req: PlatformRTPStatusRequest) {
  return axiosInstance.post<Response<string>>(`${monitorPath}/platformrtpstatus`, req)
}

export function serviceStatus(req: ServiceStatusRequest) {
  return axiosInstance.post<Response<string>>(`${monitorPath}/servicestatus`, req)
}
