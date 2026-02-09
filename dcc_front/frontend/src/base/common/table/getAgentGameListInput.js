import { BaseServersideTableInput } from '@/base/common/table/tableInput'

export class GetAgentGameListInput extends BaseServersideTableInput {
  constructor(agentId, gameId, state, length) {
    super(length)

    this.agentId = agentId
    this.gameId = gameId
    this.state = state
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      game_id: this.gameId,
      state: this.state,
      start: this.start,
      length: this.length,
      draw: this.draw,
    }
  }
}
