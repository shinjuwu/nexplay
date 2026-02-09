import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

const RummyAction = {
  Deal: 0,
  PickCard: 1,
  PickCardFromDiscard: 2,
  Discard: 3,
  Drop: 4,
}

function generateHandCardSet(cards) {
  const cardSet = []
  if (cards) {
    if (cards.Runs) {
      cardSet.push(...cards.Runs)
    }
    if (cards.ImpureRuns) {
      cardSet.push(...cards.ImpureRuns)
    }
    if (cards.Sets) {
      cardSet.push(...cards.Sets)
    }
    if (cards.ImpureSets) {
      cardSet.push(...cards.ImpureSets)
    }
    if (cards.Others) {
      cardSet.push(cards.Others)
    }
  }
  return cardSet
}

export class RummyGameLog {
  constructor(data) {
    // per point
    this.perPoint = data.ante
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)
    // 特殊百搭排
    this.wildCard = data.wild
    // 玩家資訊
    this.playerLogs = []
    // 遊戲歷程
    this.history = []

    const playerNames = {}
    data.playerlog.forEach((d) => {
      playerNames[d.seatId] = d.username
      this.playerLogs.push(new RummyPlayerLog(d, data.gameend))
    })

    // golang map[順序(數字string)]{}
    for (let i = 0; i < Object.keys(data.history).length; i++) {
      const history = data.history[i]
      const userName = playerNames[history.SeatId]

      const cards = Object.keys(history)
        .filter((key) => Array.isArray(history[key]))
        .reduce((ret, cur) => {
          ret[cur] = history[cur]
          return ret
        }, {})

      this.history.push({
        action:
          history.Action === 'PickCard'
            ? history.Deck === 0
              ? RummyAction.PickCard
              : RummyAction.PickCardFromDiscard
            : RummyAction[history.Action],
        seatId: history.SeatId,
        seat: history.SeatId + 1,
        card: history.Card,
        cardSet: generateHandCardSet(cards),
        userName: userName,
      })
    }
  }
}

class RummyPlayerLog extends BattlePlayerLog {
  constructor(data, isGameEnd) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 結果 (0:獲勝, 1:失敗, 2:棄牌, 3:流局)
    this.result = data.status ? 2 : isGameEnd ? 3 : this.profitScore > 0 ? 0 : 1

    // 手牌
    this.cardSet = generateHandCardSet(data.cards)
    // 積分
    this.cardPoint = data.cardpoint
  }
}
