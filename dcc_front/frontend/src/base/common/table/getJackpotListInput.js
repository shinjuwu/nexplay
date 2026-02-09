import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetJackpotListInput extends BaseTableInput {
  constructor(agentId, startTime, endTime, logNumber, tokenId, userName, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.startTime = startTime
    this.endTime = endTime
    this.logNumber = logNumber
    this.tokenId = tokenId
    this.userName = userName
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      lognumber: this.logNumber,
      token_id: this.tokenId,
      username: this.userName,
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
