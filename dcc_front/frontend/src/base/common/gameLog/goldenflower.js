import { getTotalTax } from '@/base/common/gameLog/common'
import { PokerCardType } from '@/base/common/gameLog/poker'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

const GoldenflowerStatus = {
  Winner: 0,
  LookAndCampareLose: 1,
  CompareLose: 2,
  LookAndFold: 3,
  Fold: 4,
}

export class GoldenflowerGameLog {
  constructor(data) {
    // 底注
    this.baseBet = data.basebet
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)
    // 遊戲歷程
    this.history = []
    // 玩家資訊
    this.playerLogs = []

    const playerNames = {}
    const playerLastOpRound = {}

    const playerLogs = []
    data.playerlog.forEach((d) => {
      playerNames[d.seatId] = d.username
      playerLogs.push(new GoldenflowerPlayerLog(d, data.winner))
    })

    // golang map[順序(數字string)]{}
    let index = 0
    for (let i = 0; i < Object.keys(data.ophistory).length; i++) {
      const roundHistories = []
      for (let j = 0; j < data.ophistory[i].length; j++) {
        const roundHistory = data.ophistory[i][j]
        const userName = playerNames[roundHistory.SeatId]

        roundHistories.push({
          index: index++,
          op: roundHistory.Op,
          bet: roundHistory.BetGold,
          seatId: roundHistory.SeatId,
          seat: roundHistory.SeatId + 1,
          userName: userName,
        })

        playerLastOpRound[roundHistory.SeatId] = i > 0 ? i : 1
      }

      // 第0跟第1需當作同一輪
      if (i === 1) {
        this.history[0] = this.history[0].concat(roundHistories)
        continue
      }

      this.history.push(roundHistories)
    }

    playerLogs.forEach((cur) => {
      cur.lastOpRound = playerLastOpRound[cur.seatId]
      this.playerLogs.push(cur)
    })
  }
}

class GoldenflowerPlayerLog extends BattlePlayerLog {
  constructor(data, winner) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 手牌
    this.cards = data.cards
    // 牌型
    this.cardType = PokerCardType[data.cardtype]

    // 狀態
    if (winner === data.seatId) {
      this.status = GoldenflowerStatus.Winner
    } else if (data.isLookAt) {
      this.status = data.isFold ? GoldenflowerStatus.LookAndFold : GoldenflowerStatus.LookAndCampareLose
    } else {
      this.status = data.isFold ? GoldenflowerStatus.Fold : GoldenflowerStatus.CompareLose
    }
  }
}
