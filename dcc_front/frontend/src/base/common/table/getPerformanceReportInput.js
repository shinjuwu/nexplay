import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetPerformanceReportInput extends BaseTableInput {
  constructor(agentId, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
