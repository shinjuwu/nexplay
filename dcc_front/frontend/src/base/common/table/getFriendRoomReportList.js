import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetFriendRoomReportListInput extends BaseTableInput {
  constructor(agentId, userName, gameId, roomId, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.userName = userName
    this.gameId = gameId
    this.roomId = roomId
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      username: this.userName,
      game_id: this.gameId,
      room_Id: this.roomId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
    }
  }
}
