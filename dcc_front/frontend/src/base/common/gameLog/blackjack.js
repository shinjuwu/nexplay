import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

const blackjackCardType = {
  bust: 1,
  none: 2,
  fivedragon: 3,
  '21point': 4,
  blackjack: 5,
}

export class BlackjackGameLog {
  constructor(data) {
    // 莊家牌資訊
    this.bankerPlayCardInfo = new BlackjackPlayCardInfo(data.result, data.cardtype, data.cardpoint)
    // 是否為保險局
    this.isInsurance = data.playerlog.reduce((r, c) => r || c.isInsurance, false)
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new BlackjackPlayerLog(d))
  }
}

class BlackjackPlayCardInfo {
  constructor(cards, cardType, cardPoint, isDouble) {
    // 手牌
    this.cards = cards
    // 牌型
    this.cardType = blackjackCardType[cardType]
    // 點數
    this.cardPoint = cardPoint
    // 是否加倍押注
    this.isDouble = isDouble
  }
}

class BlackjackPlayerLog extends BattlePlayerLog {
  constructor(data) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId
    // 有無買保險
    this.isInsurance = data.isInsurance ? data.isInsurance : false
    // 玩家牌資訊
    this.playCardInfos = data.cards
      .map(
        (cards, index) =>
          new BlackjackPlayCardInfo(cards, data.cardtype[index], data.cardpoint[index], data.isDouble[index])
      )
      .filter((playCardInfo) => playCardInfo.cards.length > 0)
  }
}
