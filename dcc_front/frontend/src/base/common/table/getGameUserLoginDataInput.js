import { BaseTableInput } from '@/base/common/table/tableInput'

export class getGameUserLoginDataInput extends BaseTableInput {
  constructor(agentId, device, browser, country, userName, ip, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.device = device
    this.browser = browser
    this.country = country
    this.userName = userName
    this.ip = ip
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      device: this.device,
      browser: this.browser,
      country_s: this.country,
      username: this.userName,
      ip: this.ip,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
