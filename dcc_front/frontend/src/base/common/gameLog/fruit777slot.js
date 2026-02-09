import { SlotPlayerLog } from '@/base/common/gameLog/slotGame'
import { round } from '@/base/utils/math'

const symbol = {
  blue7: 0, // 藍7
  red7: 1, // 紅7
  bar3: 2, // 3Bar
  bar2: 3, // 2Bar
  bar: 4, // 1Bar
  bell: 5, // 鈴鐺
  mangosteen: 6, // 山竹
  durian: 7, // 榴槤
  mango: 8, // 芒果
  rambutan: 9, // 紅毛丹
}
const special7 = {
  blue: 0,
  red: 1,
  mix: 2,
}

export class Fruit777slotGameLog {
  constructor(data) {
    // 滾輪結果
    this.reels = data.reels

    // 壓分
    this.baseBet = data.basebet

    let totalBet = round(data.basebet * data.line, 4)

    // 一般獎
    if (data.lines) {
      this.lines = []

      for (const key in data.lines) {
        const lineIdx = parseInt(key, 10)
        const lineResult = data.lines[lineIdx]

        this.lines.push({
          index: lineIdx,
          symbol: data.symbol,
          symbols: data.paylines[lineIdx].map((reelIdx, arrIdx) =>
            arrIdx < lineResult.count ? data.reels[reelIdx] : 99
          ),
          odds: lineResult.odds,
          wins: round(data.basebet * lineResult.odds, 4),
        })
      }

      this.lines.sort((a, b) => a.index - b.index)
    }

    // 特殊7獎
    if (!isNaN(data.count7)) {
      this.special7 = {
        index: data.special7,
        symbols: getSpecialSymbols(data.reels, data.special7),
        odds: data.special7odds,
        wins: round(totalBet * data.special7odds, 4),
      }
    }

    // 全盤
    if (!isNaN(data.jp)) {
      this.jp = {
        symbol: data.jp,
        odds: data.jpodds,
        wins: round(totalBet * data.jpodds, 4),
      }
    }

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new FruitslotPlayerLog(d))
  }
}

class FruitslotPlayerLog extends SlotPlayerLog {
  constructor(data) {
    super(data)
  }
}

function getSpecialSymbols(reels, specialIdx) {
  const symbols = []

  if (specialIdx === special7.blue || specialIdx === special7.mix) {
    symbols.push({
      symbol: symbol.blue7,
      count: reels.filter((s) => s === symbol.blue7).length,
    })
  }

  if (specialIdx === special7.red || specialIdx === special7.mix) {
    symbols.push({
      symbol: symbol.red7,
      count: reels.filter((s) => s === symbol.red7).length,
    })
  }

  return symbols
}
