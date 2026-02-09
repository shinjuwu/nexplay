import {
  SlotPlayerLog,
  SLOT_TYPE,
  SLOT_FEATURE_GAME_SYMBOLS,
  generateLineBoard,
  generateWayGameLineInfo,
  isWinFreeGame,
  generateLineBoardActive,
} from '@/base/common/gameLog/slotGame'
import * as array from '@/base/utils/array'

const reelStartIndex = 1
const reelEndIndex = 3

export class JumphighslotGameLog {
  constructor(data) {
    // main game 內容
    this.mainGameInfo = generateMGInfo(data.result)
    // free game 內容
    this.freeGameInfo = generateFGInfo(data.result, data.mg_gid)
    // 本上內容
    this.slotType = data.result.slotType

    const finalBoard = this.mainGameInfo ? this.mainGameInfo.board : this.freeGameInfo.originalBoard
    const finalReels = array.transpose2D(finalBoard)
    const fgMulti = this.mainGameInfo ? 1 : data.result.fg_multi

    const { lines, lineCombinationSymbolPositions } = generateWayGameLineInfo(
      finalReels,
      data.result.windetail,
      data.result.bet,
      fgMulti
    )

    if (isWinFreeGame(data.result.case)) {
      generateFGAwardLine(lines)
    }

    // 得獎連線資訊
    this.lines = lines
    this.wildMulti = data.result.fg_multi
    this.boardActive = generateLineBoardActive(finalBoard, lineCombinationSymbolPositions)
    // wild倍率
    this.multiBoard = generateMultiBoard(data.result.fg_multi)

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

  return {
    board: board,
  }
}

function generateFGInfo(result, maingGameLognumber) {
  if (result.slotType !== SLOT_TYPE.FG) {
    return null
  }

  const originalBoard = generateLineBoard(result.PayReels, reelStartIndex, reelEndIndex)

  return {
    maingGameLognumber: maingGameLognumber,
    stage: result.stages,
    originalBoard: originalBoard,
  }
}

function generateFGAwardLine(lines) {
  // 遊戲出現五個free game圖案才會觸發免費遊戲12次
  lines.push({
    key: `${SLOT_FEATURE_GAME_SYMBOLS.SF},${SLOT_FEATURE_GAME_SYMBOLS.SF},${SLOT_FEATURE_GAME_SYMBOLS.SF},${SLOT_FEATURE_GAME_SYMBOLS.SF},${SLOT_FEATURE_GAME_SYMBOLS.SF}`,
    winWay: '-',
    multi: '-',
    symbols: [
      SLOT_FEATURE_GAME_SYMBOLS.SF,
      SLOT_FEATURE_GAME_SYMBOLS.SF,
      SLOT_FEATURE_GAME_SYMBOLS.SF,
      SLOT_FEATURE_GAME_SYMBOLS.SF,
      SLOT_FEATURE_GAME_SYMBOLS.SF,
    ],
    wins: '-',
    freeGameCount: 12,
    bonusGameCount: '-',
  })
}

function generateMultiBoard(wildMulti) {
  const symbols = []

  if (wildMulti >= 10) {
    symbols.push(Math.floor(wildMulti / 10), wildMulti % 10)
  } else {
    symbols.push(wildMulti)
  }

  return symbols
}

class JumphighslotPlayerLog extends SlotPlayerLog {
  constructor(data) {
    super(data)
  }
}
