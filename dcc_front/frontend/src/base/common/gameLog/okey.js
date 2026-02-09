import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

const OkeyAction = {
  Deal: 0,
  PickCard: 1,
  PickCardFromDiscard: 2,
  Discard: 3,
  Win: 4,
}

export class OkeyGameLog {
  constructor(data) {
    // 底注
    this.perPoint = data.ante
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)
    // 指示牌
    this.indicator = data.indicator
    // 玩家資訊
    this.playerLogs = []
    // 遊戲歷程
    this.history = []

    const playerNames = {}
    data.playerlog.forEach((d) => {
      playerNames[d.seatId] = d.username
      this.playerLogs.push(new OkeyPlayerLog(d, data.gameend))
    })

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
              ? OkeyAction.PickCard
              : OkeyAction.PickCardFromDiscard
            : OkeyAction[history.Action],
        seatId: history.SeatId,
        seat: history.SeatId + 1,
        card: history.Card,
        cardSet: generateHandCardSet(cards),
        userName: userName,
      })
    }
  }
}

function generateHandCardSet(cards) {
  const cardSet = []
  if (cards) {
    if (cards.CouldRuns) {
      cardSet.push(...cards.CouldRuns)
    }
    if (cards.Runs) {
      cardSet.push(...cards.Runs)
    }
    if (cards.ImRuns) {
      cardSet.push(...cards.ImRuns)
    }
    if (cards.CouldSets) {
      cardSet.push(...cards.CouldSets)
    }
    if (cards.Sets) {
      cardSet.push(...cards.Sets)
    }
    if (cards.ImSets) {
      cardSet.push(...cards.ImSets)
    }
    if (cards.ImPairs) {
      cardSet.push(...cards.ImPairs)
    }
    if (cards.Pairs) {
      cardSet.push(...cards.Pairs)
    }
    if (cards.Others) {
      cardSet.push(cards.Others)
    }
  }
  return cardSet
}

class OkeyPlayerLog extends BattlePlayerLog {
  constructor(data, isGameEnd) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 結果 (0:獲勝, 1:失敗, 2:棄牌, 3:流局)
    this.result = data.status ? 2 : isGameEnd ? 3 : this.profitScore > 0 ? 0 : 1

    // 手牌
    this.cardSet = generateHandCardSet(data.cards)
  }
}
