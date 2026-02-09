import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetGeneralAgentRTPSetListInput extends BaseTableInput {
  constructor(agentId, state, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.state = state
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      state_type: this.state,
    }
  }
}
