import { PlayerLog } from '@/base/common/gameLog/common'
import * as array from '@/base/utils/array'
import * as math from '@/base/utils/math'

export const SLOT_HIGH_SYMBOLS = {
  H1: 1,
  H2: 2,
  H3: 3,
  H4: 4,
  H5: 5,
  H6: 6,
  H7: 7,
  H8: 8,
  H9: 9,
  H10: 10,
}

export const SLOT_LOW_SYMBOLS = {
  L1: 11,
  L2: 12,
  L3: 13,
  L4: 14,
  L5: 15,
  L6: 16,
  L7: 17,
  L8: 18,
  L9: 19,
  L10: 20,
}

export const SLOT_FEATURE_GAME_SYMBOLS = {
  SF: 21,
  SB: 22,
}

export const SLOT_WILD_SYMBOLS = {
  WW: 31,
  WW_2: 32,
  WW_3: 33,
}

export const SLOT_TYPE = {
  NG: 0,
  FG: 1,
}

export const SLOT_CASE = {
  Lose: 0x0000,
  Win: 0x0001,
  FreeGame: 0x0010,
  Bonus: 0x0020,
}

export const SLOT_LINE_COLORS = [
  '#ff0000', // red
  '#ffff00', // yellow
  '#00ff00', // lime
  '#00ffff', // aqua (cyan)
  '#0000ff', // blue
  '#ff00ff', // magenta (fuchsia)
  '#4682b4', // steelblue
  '#bdb76b', // darkkhaki
  '#b0e0e6', // powderblue
  '#b22222', // firebrick
  '#cd5c5c', // indianred
  '#7fffd4', // aquamarine
  '#db7093', // palevioletred
  '#c71585', // mediumvioletred
  '#66cdaa', // mediumaquamarine
  '#f08080', // lightcoral
  '#b0e0e6', // powderblue
  '#d2691e', // chocolate
  '#00bfff', // deepskyblue
  '#8b0000', // darkred
]

export class SlotPlayerLog extends PlayerLog {
  constructor(data) {
    super(data)

    // 起始金額
    this.startScore = data.start_score
  }
}

export function isWinFreeGame(slotCase) {
  return (slotCase & SLOT_CASE.FreeGame) === SLOT_CASE.FreeGame
}

export function isWinBonusGame(slotCase) {
  return (slotCase & SLOT_CASE.Bonus) === SLOT_CASE.Bonus
}

export function generateLineBoard(reels, startIndex, endIndex) {
  return array.transpose2D(reels).filter((_, i) => i >= startIndex && i <= endIndex)
}

export function generateDrawLines(anchor2DArr, lineColumns, lineDiffs) {
  const drawLines = []
  for (let i = 0; i < lineColumns.length; i++) {
    const diff = lineDiffs[i]
    const path = []
    for (let j = 0; j < lineColumns[i].length - 1; j++) {
      const curAnchor = anchor2DArr[lineColumns[i][j]][j]
      const nextAnchor = anchor2DArr[lineColumns[i][j + 1]][j + 1]

      path.push({
        x1: curAnchor.x + diff.x,
        y1: curAnchor.y + diff.y,
        x2: nextAnchor.x + diff.x,
        y2: nextAnchor.y + diff.y,
      })
    }

    drawLines.push({ key: i, color: SLOT_LINE_COLORS[i], path: path })
  }
  return drawLines
}

function generateWayGameLineCombinations(
  symbol,
  match,
  wwSymbolIdxArr,
  reels,
  curReelIdx,
  currentLineSymbols,
  currentLineSymbolPositions,
  resultLineCombinations,
  resultLineCombinationSymbolPositions
) {
  if (curReelIdx === match) {
    resultLineCombinations.push(currentLineSymbols.slice())
    resultLineCombinationSymbolPositions.push(currentLineSymbolPositions.slice())
    return
  }

  const subSymbols = reels[curReelIdx]
  for (let i = 0; i < subSymbols.length; i++) {
    const subSymbol = subSymbols[i]
    if (subSymbol !== symbol && wwSymbolIdxArr.indexOf(subSymbol) < 0) {
      continue
    }

    currentLineSymbols.push(subSymbol)
    currentLineSymbolPositions.push(i)
    generateWayGameLineCombinations(
      symbol,
      match,
      wwSymbolIdxArr,
      reels,
      curReelIdx + 1,
      currentLineSymbols,
      currentLineSymbolPositions,
      resultLineCombinations,
      resultLineCombinationSymbolPositions
    )
    currentLineSymbols.pop()
    currentLineSymbolPositions.pop()
  }
}

export function generateWayGameLineInfo(reels, winDetail, bet, wildMulti) {
  if (!winDetail) {
    return { lines: [], lineCombinationSymbolPositions: [] }
  }

  const cacheSymbolLineCombinations = {}
  const lineCombinationSymbolPositions = []
  for (let wIdx = 0; wIdx < winDetail.length; wIdx++) {
    const match = winDetail[wIdx].Match
    const winWay = winDetail[wIdx].WinWay
    const multi = math.round(parseFloat(winDetail[wIdx].Multi) / winWay / bet, 4)
    const symbol = winDetail[wIdx].Symbol
    const lineCombinations = []
    generateWayGameLineCombinations(
      symbol,
      match,
      Object.values(SLOT_WILD_SYMBOLS),
      reels,
      0,
      [],
      [],
      lineCombinations,
      lineCombinationSymbolPositions
    )
    if (lineCombinations.length !== winWay) {
      console.error(
        `Line win detail error, detail:${JSON.stringify(winDetail[wIdx])}, lines:${JSON.stringify(
          cacheSymbolLineCombinations[symbol]
        )}`
      )
    }

    cacheSymbolLineCombinations[symbol] = {}
    for (let lIdx = 0; lIdx < lineCombinations.length; lIdx++) {
      const combinationKey = lineCombinations[lIdx].join(',')
      if (!cacheSymbolLineCombinations[symbol][combinationKey]) {
        cacheSymbolLineCombinations[symbol][combinationKey] = {
          key: combinationKey,
          winWay: 0,
          multi: multi,
          symbols: lineCombinations[lIdx],
          wins: 0,
          freeGameCount: '-',
          bonusGameCount: '-',
          wildMulti: '-',
        }
      }
      cacheSymbolLineCombinations[symbol][combinationKey].wins = math.round(
        cacheSymbolLineCombinations[symbol][combinationKey].wins + math.round(multi * bet * wildMulti, 4),
        4
      )
      cacheSymbolLineCombinations[symbol][combinationKey].winWay++
    }
  }

  return {
    lines: Object.values(cacheSymbolLineCombinations).flatMap((obj) => Object.values(obj)),
    lineCombinationSymbolPositions: lineCombinationSymbolPositions,
  }
}

export function generateLineBoardActive(board, lineCombinationSymbolPositions) {
  const boardActive = board.map((v) => v.map(() => false))
  for (let i = 0; i < lineCombinationSymbolPositions.length; i++) {
    const lineSymbolPositions = lineCombinationSymbolPositions[i]
    for (let j = 0; j < lineSymbolPositions.length; j++) {
      boardActive[lineSymbolPositions[j]][j] = true
    }
  }
  return boardActive
}
