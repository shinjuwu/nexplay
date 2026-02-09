import { BaseServersideTableInput } from '@/base/common/table/tableInput'
import { endOfDay, startOfDay } from 'date-fns'

export class GetDailySettlementReportListInput extends BaseServersideTableInput {
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
      start_time: startTime.toISOString(),
      end_time: endTime.toISOString(),
      start: this.start,
      length: this.length,
      draw: this.draw,
      sort_column: this.column,
      sort_direction: this.dir,
    }
  }
}
