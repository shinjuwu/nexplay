import {
  SlotPlayerLog,
  SLOT_TYPE,
  SLOT_WILD_SYMBOLS,
  SLOT_FEATURE_GAME_SYMBOLS,
  generateDrawLines,
  generateLineBoard,
  isWinFreeGame,
} from '@/base/common/gameLog/slotGame'
import * as array from '@/base/utils/array'
import * as math from '@/base/utils/math'

const reelStartIndex = 1
const reelEndIndex = 3

const paylineAnchors = [
  [
    { x: 62.87, y: 56 },
    { x: 197.62, y: 56 },
    { x: 333.37, y: 56 },
    { x: 465.12, y: 56 },
    { x: 600.87, y: 56 },
  ],
  [
    { x: 62.87, y: 168 },
    { x: 197.62, y: 168 },
    { x: 333.37, y: 168 },
    { x: 465.12, y: 168 },
    { x: 600.87, y: 168 },
  ],
  [
    { x: 62.87, y: 280 },
    { x: 197.62, y: 280 },
    { x: 333.37, y: 280 },
    { x: 465.12, y: 280 },
    { x: 600.87, y: 280 },
  ],
]
const paylineLineColumns = [
  [1, 1, 1, 1, 1],
  [0, 0, 0, 0, 0],
  [2, 2, 2, 2, 2],
  [0, 1, 2, 1, 0],
  [2, 1, 0, 1, 2],
  [1, 0, 0, 0, 1],
  [1, 2, 2, 2, 1],
  [0, 0, 1, 2, 2],
  [2, 2, 1, 0, 0],
  [1, 2, 1, 0, 1],
  [1, 0, 1, 2, 1],
  [0, 1, 1, 1, 0],
  [2, 1, 1, 1, 2],
  [0, 1, 0, 1, 0],
  [2, 1, 2, 1, 2],
  [1, 1, 0, 1, 1],
  [1, 1, 2, 1, 1],
  [0, 0, 2, 0, 0],
  [2, 2, 0, 2, 2],
  [0, 2, 2, 2, 0],
]
const paylineLineDiffs = [
  { x: 0, y: 0 },
  { x: 0, y: 0 },
  { x: 0, y: 0 },
  { x: 0, y: 2 },
  { x: 0, y: 4 },
  { x: 0, y: 6 },
  { x: 0, y: 8 },
  { x: 0, y: 10 },
  { x: 0, y: 12 },
  { x: 0, y: 14 },
  { x: 0, y: 16 },
  { x: 0, y: 18 },
  { x: 0, y: -16 },
  { x: 0, y: -14 },
  { x: 0, y: -12 },
  { x: 0, y: -10 },
  { x: 0, y: -8 },
  { x: 0, y: -6 },
  { x: 0, y: -4 },
  { x: 0, y: -2 },
]
const paylines = generateDrawLines(paylineAnchors, paylineLineColumns, paylineLineDiffs)

export class MidasslotGameLog {
  constructor(data) {
    // main game 內容
    this.mainGameInfo = generateMGInfo(data.result)
    // free game 內容
    this.freeGameInfo = generateFGInfo(data.result, data.mg_gid)
    // 本上內容
    this.slotType = data.result.slotType

    const finalBoard = this.mainGameInfo ? this.mainGameInfo.board : this.freeGameInfo.board
    const finalReels = array.transpose2D(finalBoard)

    const { lines } = generateLineGameLineInfo(
      finalReels,
      data.result.windetail,
      paylineLineColumns,
      data.result.bet,
      data.result.wwMulti
    )

    // 畫線資訊
    this.drawLines = Object.values(lines).map((line) => paylines[line.id - 1])

    if (isWinFreeGame(data.result.case)) {
      generateFGAwardLine(lines, data.result.scatterInfo)
    }
    // 線獎資訊
    this.lines = lines

    // wild倍率
    this.multiBoard = generateMultiBoard(data.result.wwMulti)

    // 總押分 (free game的總押分會是其觸發main game的總押分)
    this.bet = data.result.bet

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new MidasslotPlayerLog(d))
  }
}

function generateLineGameLineInfo(reels, winDetail, paylines, bet, wildMulti) {
  if (!winDetail) {
    return { lines: [] }
  }

  const lines = []
  for (let wIdx = 0; wIdx < winDetail.length; wIdx++) {
    const match = winDetail[wIdx].Match
    const multi = math.round(winDetail[wIdx].Multi / bet, 4)
    const winLineIdx = winDetail[wIdx].WinLine

    const symbols = paylines[winLineIdx].slice(0, match).map((posIdx, reelIdx) => reels[reelIdx][posIdx])
    lines.push({
      key: `${winLineIdx}|${symbols.join(',')}`,
      id: winLineIdx + 1,
      symbols: symbols,
      symbolMulti: multi,
      wildMulti: wildMulti,
      wins: math.round(multi * bet * wildMulti, 4),
      freeGameCount: '-',
    })
  }

  return {
    lines: lines,
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

  const reels = result.PayReels
  if (result.WildReel) {
    result.WildReel.After.forEach((wildSymbols, reelIdx) => {
      wildSymbols.forEach((wildSymbol, posIdx) => {
        switch (wildSymbol) {
          case 1:
            reels[reelIdx][posIdx] = SLOT_WILD_SYMBOLS.WW
            return
          case 2:
            reels[reelIdx][posIdx] = SLOT_WILD_SYMBOLS.WW_2
            return
          case 3:
            reels[reelIdx][posIdx] = SLOT_WILD_SYMBOLS.WW_3
            return
        }
      })
    })
  }

  const board = generateLineBoard(reels, reelStartIndex, reelEndIndex)

  return {
    maingGameLognumber: maingGameLognumber,
    stage: result.stages,
    board: board,
  }
}

function generateFGAwardLine(lines, scatterInfo) {
  const symbols = Array.from({ length: scatterInfo.ScatterCount }, () => SLOT_FEATURE_GAME_SYMBOLS.SF)
  lines.push({
    key: `-|${symbols.join(',')}`,
    id: '-',
    symbols: symbols,
    symbolMulti: '-',
    wildMulti: '-',
    wins: '-',
    freeGameCount: scatterInfo.ScSmallReel.reduce((sumCount, count) => sumCount + count, 0),
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

class MidasslotPlayerLog extends SlotPlayerLog {
  constructor(data) {
    super(data)
  }
}
