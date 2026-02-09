import { BaseTableInput } from '@/base/common/table/tableInput'

export class GetWalletLedgerListInput extends BaseTableInput {
  constructor(id, agentId, userName, kind, startTime, endTime, singleWalletId, length, column, dir) {
    super(length, column, dir)

    this.id = id
    this.agentId = agentId
    this.userName = userName
    this.kind = kind
    this.startTime = startTime
    this.endTime = endTime
    this.singleWalletId = singleWalletId
  }

  parseInputJson() {
    return {
      id: this.id,
      agent_id: this.agentId,
      username: this.userName,
      kind: this.kind,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
      single_wallet_id: this.singleWalletId,
    }
  }
}
