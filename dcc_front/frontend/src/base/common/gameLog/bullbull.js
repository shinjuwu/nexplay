import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

export class BullbullGameLog {
  constructor(data) {
    // 底注
    this.baseBet = data.basebet
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new BullbullPlayerLog(d, data.bankerid, data.bankerBet))
  }
}

class BullbullPlayerLog extends BattlePlayerLog {
  constructor(data, bankerId, bankerBet) {
    super(data)

    // 下注倍数
    this.bet = data.seatId === bankerId ? bankerBet : data.bet
    // 手牌
    this.cards = data.cards
    // 牌型
    this.cardType = data.cardtype
    // 是否為莊家
    this.isBanker = data.seatId === bankerId
  }
}
