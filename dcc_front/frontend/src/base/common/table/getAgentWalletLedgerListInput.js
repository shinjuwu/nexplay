import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetAgentWalletLedgerListInput extends BaseTableInput {
  constructor(agentId, id, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.id = id
    this.agentId = agentId
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      id: this.id,
      agent_id: this.agentId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
