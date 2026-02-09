import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

const betArea = {
  Odd: 0, // 單
  Even: 1, // 雙
  FourRed: 2, // 4紅
  ThreeRed: 3, // 3洪
  ThreeWhite: 4, // 3白
  FourWhite: 5, // 4白
}

export class ColordiscGameLog {
  constructor(data) {
    // 骰子結果
    this.result = data.result
    // 白色骰子數量
    this.whiteCount = data.result.reduce((count, i) => count + i, 0)
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new HundredPlayerLog(d))
  }

  getBetAreaRedWhiteResult(betAreaId) {
    switch (betAreaId) {
      case betArea.FourRed:
        return [0, 0, 0, 0]
      case betArea.ThreeRed:
        return [0, 0, 0, 1]
      case betArea.ThreeWhite:
        return [1, 1, 1, 0]
      case betArea.FourWhite:
        return [1, 1, 1, 1]
      default:
        return []
    }
  }
}
