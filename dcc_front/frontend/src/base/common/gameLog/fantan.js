import { HundredPlayerLog, getWinAreas } from '@/base/common/gameLog/hundredGame'

export class FantanGameLog {
  constructor(data) {
    // 押注區開獎結果
    this.winAreas = getWinAreas(data.winarea)
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new HundredPlayerLog(d))
  }
}
