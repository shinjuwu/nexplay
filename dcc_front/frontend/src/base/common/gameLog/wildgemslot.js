import {
  SlotPlayerLog,
  SLOT_TYPE,
  SLOT_FEATURE_GAME_SYMBOLS,
  generateDrawLines,
  generateLineBoard,
  isWinFreeGame,
} from '@/base/common/gameLog/slotGame'
import * as array from '@/base/utils/array'
import * as math from '@/base/utils/math'

const reelStartIndex = 1
const reelEndIndex = 3

const freeGameTimes = {
  3: 7,
  4: 10,
  5: 15,
}
const wildSymbolByMulti = {
  0.3: 32,
  0.5: 33,
  0.8: 34,
  1: 35,
  1.5: 36,
  2: 37,
  3: 38,
  4: 39,
  5: 40,
  6: 41,
  7: 42,
  8: 43,
  9: 44,
  10: 45,
  20: 46,
  30: 47,
  50: 48,
  100: 49,
}

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

export class WildgemslotGameLog {
  constructor(data) {
    // main game 內容
    this.mainGameInfo = generateMGInfo(data.result)
    // free game 內容
    this.freeGameInfo = generateFGInfo(data.result, data.mg_gid)

    const finalBoard = this.mainGameInfo ? this.mainGameInfo.board : this.freeGameInfo.board
    const finalReels = array.transpose2D(finalBoard)

    const { lines } = generateLineGameLineInfo(finalReels, data.result.windetail, paylineLineColumns, data.result.bet)

    // 畫線資訊
    this.drawLines = lines.map((line) => paylines[line.id - 1])

    if (Array.isArray(data.result.wildprize) && data.result.wildprize.length > 0) {
      generateWildAwardLine(lines, data.result.wildprize, data.result.bet)
    }

    if (isWinFreeGame(data.result.case)) {
      const fgSymbolCount = finalReels.reduce((count, rowSymbols) => {
        return count + rowSymbols.filter((symbol) => symbol === SLOT_FEATURE_GAME_SYMBOLS.SF).length
      }, 0)
      generateFGAwardLine(lines, fgSymbolCount)
    }

    // 線獎資訊
    this.lines = lines

    // 總押分 (free game的總押分會是其觸發main game的總押分)
    this.bet = data.result.bet

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new WildgemslotPlayerLog(d))
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

  const board = generateLineBoard(result.PayReels, reelStartIndex, reelEndIndex)

  return {
    maingGameLognumber: maingGameLognumber,
    stage: result.stages,
    board: board,
  }
}

function generateLineGameLineInfo(reels, winDetail, paylines, bet) {
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
      wildMulti: '-',
      wins: math.round(multi * bet, 4),
      freeGameCount: '-',
    })
  }

  return {
    lines: lines,
  }
}

function generateWildAwardLine(lines, wildprize, bet) {
  const wildMulti = wildprize.reduce((s, c) => s + c, 0)
  const wins = math.round(bet * wildMulti, 4)

  const symbols = wildprize.map((multi) => wildSymbolByMulti[multi])
  lines.push({
    key: `-|${symbols.join(',')}`,
    id: '-',
    symbols: symbols,
    symbolMulti: '-',
    wildMulti: wildMulti,
    wins: wins,
    freeGameCount: '-',
  })
}

function generateFGAwardLine(lines, fgSymbolCount) {
  const symbols = Array.from({ length: fgSymbolCount }, () => SLOT_FEATURE_GAME_SYMBOLS.SF)
  lines.push({
    key: `-|${symbols.join(',')}`,
    id: '-',
    symbols: symbols,
    symbolMulti: '-',
    wildMulti: '-',
    wins: '-',
    freeGameCount: freeGameTimes[fgSymbolCount],
  })
}

class WildgemslotPlayerLog extends SlotPlayerLog {
  constructor(data) {
    super(data)
  }
}
