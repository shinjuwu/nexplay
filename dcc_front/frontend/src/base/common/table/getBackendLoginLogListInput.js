import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetBackendLoginLogListInput extends BaseTableInput {
  constructor(agentId, userName, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.userName = userName
    this.startTime = startTime
    this.endTime = endTime
    this.ip = ''
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      username: this.userName,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
