import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetRTPSettingListInput extends BaseTableInput {
  constructor(agentId, gameId, gameType, roomType, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.gameType = gameType
    this.gameId = gameId
    this.roomType = roomType
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      game_type: this.gameType,
      game_id: this.gameId,
      room_type: this.roomType,
    }
  }
}
