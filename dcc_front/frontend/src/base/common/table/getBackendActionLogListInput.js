import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetBackendActionLogListInput extends BaseTableInput {
  constructor(agentId, actionType, featureCodes, startTime, endTime, length, column, dir) {
    super(length, column, dir)

    this.agentId = agentId
    this.actionType = actionType
    this.featureCodes = featureCodes
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      action_type: this.actionType,
      feature_codes: this.featureCodes,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
    }
  }
}
