import {
  SlotPlayerLog,
  SLOT_TYPE,
  SLOT_FEATURE_GAME_SYMBOLS,
  generateLineBoard,
  generateWayGameLineInfo,
  isWinBonusGame,
} from '@/base/common/gameLog/slotGame'
import * as array from '@/base/utils/array'

const reelStartIndex = 1
const reelEndIndex = 3

export class PyrtreasureslotGameLog {
  constructor(data) {
    // main game 內容
    this.mainGameInfo = generateMGInfo(data.result)
    // free game 內容
    this.freeGameInfo = generateFGInfo(data.result, data.mg_gid)
    // 本上內容
    this.slotType = data.result.slotType

    const finalBoard = this.mainGameInfo ? this.mainGameInfo.board : this.freeGameInfo.board
    const finalReels = this.mainGameInfo ? array.transpose2D(finalBoard) : []

    const { lines } = generateWayGameLineInfo(finalReels, data.result.windetail, data.result.bet, 1)
    console.log(lines)

    if (isWinBonusGame(data.result.case)) {
      generateBGAwardLine(lines)
    }

    // 得獎連線資訊
    this.lines = lines
    this.wildMulti = data.result.fg_multi
    this.bonusWins = data.result.case
    // 總押分 (free game的總押分會是其觸發main game的總押分)
    this.bet = data.result.bet

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new JumphighslotPlayerLog(d))
  }
}

function generateMGInfo(result) {
  if (result.slotType !== SLOT_TYPE.NG) {
    return null
  }

  const board = generateLineBoard(result.PayReels, reelStartIndex, reelEndIndex)

  function checkWild(originBoard) {
    for (let i = 0; i < originBoard.length; i++) {
      originBoard[i].forEach((item, index) => {
        if (item === 31) {
          extendWild(originBoard, index)
          return
        }
      })
    }
  }

  function extendWild(originBoard, i) {
    for (let j = 0; j < originBoard.length; j++) {
      originBoard[j][i] = 31
    }
  }

  if (result.windetail) {
    checkWild(board)
  }

  return {
    board: board,
  }
}

function generateFGInfo(result, maingGameLognumber) {
  if (result.slotType !== SLOT_TYPE.FG) {
    return null
  }

  const multi = result.BonusHistory[5] ? result.BonusHistory[5][0] : -2
  console.log(multi)

  const originBoardLine1 = result.BonusHistory[1]
    ? generateBoardLine1(result.BonusHistory[1], 7)
    : [-2, -2, -2, -2, -2, -2, -2]
  const originBoardLine2 = result.BonusHistory[2]
    ? generateBoardLine1(result.BonusHistory[2], 6)
    : [-2, -2, -2, -2, -2, -2]
  const originBoardLine3 = result.BonusHistory[3] ? generateBoardLine1(result.BonusHistory[3], 5) : [-2, -2, -2, -2, -2]
  const originBoardLine4 = result.BonusHistory[4] ? generateBoardLine1(result.BonusHistory[4], 4) : [-2, -2, -2, -2]
  const bonusBoard = [originBoardLine4, originBoardLine3, originBoardLine2, originBoardLine1]
  const awardLine1 = originBoardLine1.filter((item) => item > 0)
  const awardLine2 = originBoardLine2.filter((item) => item > 0)
  const awardLine3 = originBoardLine3.filter((item) => item > 0)
  const awardLine4 = originBoardLine4.filter((item) => item > 0)
  const bonusLineInfo = [
    {
      level: 1,
      lines: awardLine1,
      bonus: awardLine1.reduce((accumulator, currentValue) => accumulator + currentValue, 0),
    },
    {
      level: 2,
      lines: awardLine2,
      bonus: awardLine2.reduce((accumulator, currentValue) => accumulator + currentValue, 0),
    },
    {
      level: 3,
      lines: awardLine3,
      bonus: awardLine3.reduce((accumulator, currentValue) => accumulator + currentValue, 0),
    },
    {
      level: 4,
      lines: awardLine4,
      bonus: awardLine4.reduce((accumulator, currentValue) => accumulator + currentValue, 0),
    },
  ]

  console.log(bonusBoard)

  return {
    maingGameLognumber: maingGameLognumber,
    bonusBoard: bonusBoard,
    bonusLineInfo: bonusLineInfo,
    multi: multi,
  }
}

function generateBoardLine1(bonusHistory, box) {
  for (let i = 0; i < box; i++) {
    if (!(i in bonusHistory)) {
      bonusHistory[i] = -2
    }
  }

  return (bonusHistory = Object.keys(bonusHistory).map((i) => bonusHistory[i]))
}

function generateBGAwardLine(lines) {
  // 遊戲出現三個bonus game圖案才會觸發紅利遊戲1次
  lines.push({
    key: `${SLOT_FEATURE_GAME_SYMBOLS.SB},${SLOT_FEATURE_GAME_SYMBOLS.SB},${SLOT_FEATURE_GAME_SYMBOLS.SB}`,
    winWay: '-',
    multi: '-',
    symbols: [SLOT_FEATURE_GAME_SYMBOLS.SB, SLOT_FEATURE_GAME_SYMBOLS.SB, SLOT_FEATURE_GAME_SYMBOLS.SB],
    wins: '-',
    freeGameCount: '-',
    bonusGameCount: 1,
  })
}

class JumphighslotPlayerLog extends SlotPlayerLog {
  constructor(data) {
    super(data)
  }
}
