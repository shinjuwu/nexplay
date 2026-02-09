import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

export class PokdengGameLog {
  constructor(data) {
    // 底注
    this.baseBet = data.basebet
    // 莊家資訊
    this.bankerLog = {
      cards: data.bankercards, // 手牌
      cardType: data.bankertype, // 牌型1
      status: data.bankertype2, // 牌型1
      odds: data.bankerodds, // 賠率
    }
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new PokdengPlayerLog(d))
  }
}

class PokdengPlayerLog extends BattlePlayerLog {
  constructor(data) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 手牌
    this.cards = data.cards
    // 牌型1
    this.cardType = data.cardtype
    // 牌型2
    this.status = data.status
    // 賠率
    this.odds = data.odds
    // 原始投注
    this.bet = data.bet
  }
}
