import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetUserCreditLogListInput extends BaseTableInput {
  constructor(agentId, userName, gameId, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.userName = userName
    this.gameId = gameId
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      username: this.userName,
      game_id: this.gameId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
