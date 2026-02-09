import {
  SlotPlayerLog,
  SLOT_HIGH_SYMBOLS,
  SLOT_LOW_SYMBOLS,
  SLOT_FEATURE_GAME_SYMBOLS,
  SLOT_TYPE,
  generateLineBoard,
  generateLineBoardActive,
  generateWayGameLineInfo,
  isWinFreeGame,
} from '@/base/common/gameLog/slotGame'
import * as array from '@/base/utils/array'

const reelStartIndex = 1
const reelEndIndex = 3

const transformSymbols = [
  SLOT_HIGH_SYMBOLS.H1,
  SLOT_HIGH_SYMBOLS.H2,
  SLOT_HIGH_SYMBOLS.H3,
  SLOT_HIGH_SYMBOLS.H4,
  SLOT_LOW_SYMBOLS.L1,
  SLOT_LOW_SYMBOLS.L2,
  SLOT_LOW_SYMBOLS.L3,
  SLOT_LOW_SYMBOLS.L4,
  SLOT_LOW_SYMBOLS.L5,
  SLOT_LOW_SYMBOLS.L6,
]

export class MegsharkslotGameLog {
  constructor(data) {
    // main game 內容
    this.mainGameInfo = generateMGInfo(data.result)
    // free game 內容
    this.freeGameInfo = generateFGInfo(data.result, data.mg_gid)

    const finalBoard = this.mainGameInfo
      ? this.mainGameInfo.board
      : this.freeGameInfo.wwBoard
      ? this.freeGameInfo.wwBoard
      : this.freeGameInfo.originalBoard
    const finalReels = array.transpose2D(finalBoard)

    const { lines, lineCombinationSymbolPositions } = generateWayGameLineInfo(
      finalReels,
      data.result.windetail,
      data.result.bet
    )
    if (isWinFreeGame(data.result.case)) {
      generateFGAwardLine(lines)
    }

    // 得獎連線資訊
    this.lines = lines
    this.boardActive = generateLineBoardActive(finalBoard, lineCombinationSymbolPositions)

    // 總押分 (free game的總押分會是其觸發main game的總押分)
    this.bet = data.result.bet

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new MegsharkslotPlayerLog(d))
  }
}

function generateMGInfo(result) {
  if (result.slotType !== SLOT_TYPE.NG) {
    return null
  }

  const board = generateLineBoard(result.originalReels, reelStartIndex, reelEndIndex)

  return {
    board: board,
  }
}

function generateFGLeftTransformBoard(midSymbol) {
  const midIdx = transformSymbols.findIndex((symbol) => symbol === midSymbol)
  const topIdx = midIdx > 0 ? midIdx - 1 : transformSymbols.length - 1
  const bottomIdx = midIdx < transformSymbols.length - 1 ? midIdx + 1 : 0
  return [transformSymbols[topIdx], transformSymbols[midIdx], transformSymbols[bottomIdx]]
}

function generateFGInfo(result, maingGameLognumber) {
  if (result.slotType !== SLOT_TYPE.FG) {
    return null
  }

  const originalBoard = generateLineBoard(result.originalReels, reelStartIndex, reelEndIndex)
  const wwBoard = result.wwReels ? generateLineBoard(result.wwReels, reelStartIndex, reelEndIndex) : null
  const sameBoard = array.equal2D(originalBoard, wwBoard)

  return {
    maingGameLognumber: maingGameLognumber,
    stage: result.stages,
    originalBoard: originalBoard,
    wwBoard: wwBoard,
    leftTransformBoard: generateFGLeftTransformBoard(result.wwSym),
    sameBoard: sameBoard,
  }
}

function generateFGAwardLine(lines) {
  // 遊戲出現三個free game圖案才會觸發免費遊戲10次
  lines.push({
    key: `${SLOT_FEATURE_GAME_SYMBOLS.SF},${SLOT_FEATURE_GAME_SYMBOLS.SF},${SLOT_FEATURE_GAME_SYMBOLS.SF}`,
    winWay: '-',
    multi: '-',
    symbols: [SLOT_FEATURE_GAME_SYMBOLS.SF, SLOT_FEATURE_GAME_SYMBOLS.SF, SLOT_FEATURE_GAME_SYMBOLS.SF],
    wins: '-',
    freeGameCount: 10,
    bonusGameCount: '-',
  })
}

class MegsharkslotPlayerLog extends SlotPlayerLog {
  constructor(data) {
    super(data)
  }
}
