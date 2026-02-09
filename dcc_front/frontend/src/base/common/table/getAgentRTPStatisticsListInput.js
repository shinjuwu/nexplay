import { BaseTableInput } from '@/base/common/table/tableInput'
import { endOfDay, startOfDay } from 'date-fns'

export class GetAgentRTPStatisticsListInput extends BaseTableInput {
  constructor(agentId, gameId, roomType, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.gameId = gameId
    this.roomType = roomType
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    const startTime = startOfDay(this.startTime)
    const endTime = endOfDay(this.endTime)

    return {
      agent_id: this.agentId,
      game_id: this.gameId,
      room_type: this.roomType,
      start_time: startTime,
      end_time: endTime,
    }
  }
}
