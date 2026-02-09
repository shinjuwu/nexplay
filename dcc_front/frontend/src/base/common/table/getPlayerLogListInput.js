import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetPlayerLogListInput extends BaseTableInput {
  constructor(agentId, userName, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.userName = userName
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      username: this.userName,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
    }
  }
}
