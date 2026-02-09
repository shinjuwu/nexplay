import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

export class RouletteGameLog {
  constructor(data) {
    // 開獎號碼
    this.openNumber = data.openNumber
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new HundredPlayerLog(d))
  }
}
