import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

export class CatteGameLog {
  constructor(data) {
    // 底注
    this.baseBet = data.basebet
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)

    const { hasHistory, playerNames, playerLogs } = transformPlayerLogs(data)
    const { history } = transformHistory(hasHistory, playerNames, data)

    // 特殊牌局
    this.isSpecialGame = !hasHistory

    // 玩家資訊
    this.playerLogs = playerLogs
    // 遊戲歷程
    this.history = history
  }
}

function transformPlayerLogs(data) {
  let hasHistory = true
  let playerNames = {}
  let playerLogs = []

  data.playerlog.forEach((d) => {
    hasHistory = hasHistory && d.cardtype === 0
    playerNames[d.seatId] = d.username
    playerLogs.push(new CattePlayerLog(d, data.playerrecord.isWinIn4Round[d.seatId]))
  })

  return { hasHistory, playerNames, playerLogs }
}

function transformHistory(hasHistory, playerNames, data) {
  let history = []

  if (!hasHistory) {
    return { history }
  }

  for (const seatIdStr in data.playerrecord.playRecord) {
    const seatId = parseInt(seatIdStr, 10)
    const playerName = playerNames[seatId]
    const playerActionInfo = data.playerrecord.playRecord[seatId]

    for (let i = 0; i < playerActionInfo.length; i++) {
      if (history.length <= i) {
        history.push([])
      }

      history[i].push({
        index: 0,
        seatId: seatId,
        seat: seatId + 1,
        userName: playerName,
        card: playerActionInfo[i].card,
        op: playerActionInfo[i].isPlay ? 1 : 0,
        roundPlay: playerActionInfo[i].roundPlay,
      })
    }
  }

  let index = 0
  for (let i = 0; i < history.length; i++) {
    // 排序完再確認index
    history[i].sort((a, b) => a.roundPlay - b.roundPlay)

    for (let j = 0; j < history[i].length; j++) {
      history[i][j].index = index++
    }
  }

  return { history }
}

class CattePlayerLog extends BattlePlayerLog {
  constructor(data, isWinFirstFourRound) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 手牌
    this.cards = data.cards
    // 牌型
    this.cardType = data.cardtype

    // 前四回合是否有贏
    this.isWinFirstFourRound = isWinFirstFourRound
  }
}
