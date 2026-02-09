import { BaseTableInput } from '@/base/common/table/tableInput'
import { startOfDay, addDays } from 'date-fns'

export class GetPerformanceReportListInput extends BaseTableInput {
  constructor(agentId, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    const startTime = startOfDay(this.startTime)
    const endTime = addDays(startOfDay(this.endTime), 1)
    return {
      agent_id: this.agentId,
      start_time: startTime.toISOString(),
      end_time: endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
