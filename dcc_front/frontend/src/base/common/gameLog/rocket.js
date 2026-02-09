import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

export class RocketGameLog {
  constructor(data) {
    // 火箭爆炸距離
    this.explodePayout = data.explodePayout
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new RocketGamePlayerLog(d))
  }
}

class RocketGamePlayerLog extends HundredPlayerLog {
  constructor(data) {
    super(data)

    this.fleePayout = data.fleePayout
  }
}
