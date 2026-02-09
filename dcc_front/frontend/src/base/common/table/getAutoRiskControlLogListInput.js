import { BaseTableInput } from '@/base/common/table/tableInput'
import { startOfDay } from 'date-fns'

export class GetAutoRiskControlLogListInput extends BaseTableInput {
  constructor(agentId, userName, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.userName = userName
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    const startTime = startOfDay(this.startTime)
    startTime.setMinutes(-startTime.getTimezoneOffset())

    const endTime = startOfDay(this.endTime)
    endTime.setMinutes(-endTime.getTimezoneOffset())

    return {
      agent_id: this.agentId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
      username: this.userName,
    }
  }
}
