import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

export class TeenpattiGameLog {
  constructor(data) {
    // 底注
    this.baseBet = data.basebet
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)

    const { playerNames, playerLogs } = transformPlayerLogs(data)
    const { playerLastOpRoundInfo, history } = transformHistory(playerNames, data)

    // 玩家資訊
    this.playerLogs = playerLogs.map((p) => {
      p.status = playerLastOpRoundInfo[p.seatId].status
      p.lastOpRound = playerLastOpRoundInfo[p.seatId].round
      return p
    })
    // 遊戲歷程
    this.history = history
  }
}

function transformPlayerLogs(data) {
  const playerNames = {}
  const playerLogs = []

  data.playerlog.forEach((d) => {
    playerNames[d.seatId] = d.username
    playerLogs.push(new TeenpattiPlayerLog(d))
  })

  return { playerNames, playerLogs }
}

function transformHistory(playerNames, data) {
  const playerLastOpRoundInfo = {}

  const roundIdStrs = Object.keys(data.playerrecord)
  const history = Array.from({ length: roundIdStrs.length }, () => [])

  for (const roundIdStr of roundIdStrs) {
    const roundId = parseInt(roundIdStr, 10)
    const roundRecords = data.playerrecord[roundIdStr]

    for (let i = 0; i < roundRecords.length; i++) {
      const playerRecord = roundRecords[i]
      const playerName = playerNames[playerRecord.seatId]

      history[roundId - 1].push({
        round: roundId,
        roundPlay: i,
        seatId: playerRecord.seatId,
        seat: playerRecord.seatId + 1,
        userName: playerName,
        status: playerRecord.status,
        bet: playerRecord.bet,
      })

      if (!playerLastOpRoundInfo[playerRecord.seatId]) {
        playerLastOpRoundInfo[playerRecord.seatId] = {
          round: roundId,
          roundPlay: i,
          status: playerRecord.status,
        }
      }

      if (
        playerLastOpRoundInfo[playerRecord.seatId].round < roundId ||
        (playerLastOpRoundInfo[playerRecord.seatId].round === roundId &&
          playerLastOpRoundInfo[playerRecord.seatId].roundPlay < i)
      ) {
        playerLastOpRoundInfo[playerRecord.seatId].round = roundId
        playerLastOpRoundInfo[playerRecord.seatId].roundPlay = i
        playerLastOpRoundInfo[playerRecord.seatId].status = playerRecord.status
      }
    }
  }

  return { playerLastOpRoundInfo, history }
}

class TeenpattiPlayerLog extends BattlePlayerLog {
  constructor(data) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 手牌
    this.cards = data.cards
    // 牌型
    this.cardType = data.cardtype

    // 最後動作回合
    this.lastOpRound = 0
  }
}
