import { BaseServersideTableInput } from '@/base/common/table/tableInput'

export class GetBackendGameUserWalletLedgerListInput extends BaseServersideTableInput {
  constructor(username, column, length) {
    super(length)

    this.username = username
    this.column = column
  }

  parseInputJson() {
    return {
      username: this.username,
      length: this.length,
      sort_column: this.column,
      sort_direction: this.dir,
      start: this.start,
    }
  }
}
