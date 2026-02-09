import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetAgentIpWhitelistListInput extends BaseTableInput {
  constructor(agentId, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
    }
  }
}
