import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

export class AndarbaharGameLog {
  constructor(data) {
    // joker牌
    this.joker = data.gameRecord.joker
    // 安達區牌
    this.andarAreaCards = data.gameRecord.andarAreaCards
    // 巴哈區牌
    this.baharAreaCards = data.gameRecord.baharAreaCards
    // 開牌張數
    this.openCardCount = data.gameRecord.openCardCount
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new HundredPlayerLog(d))
  }
}
