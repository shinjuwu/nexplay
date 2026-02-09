import { PlayerLog } from '@/base/common/gameLog/common'

export class ElectronicPlayerLog extends PlayerLog {
  constructor(data) {
    super(data)

    // 起始金額
    this.startScore = data.start_score
  }
}
