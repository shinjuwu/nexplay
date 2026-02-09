import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

const result = {
  REDWIN: 0, // 紅方勝
  BLUEWIN: 1, // 藍方勝
  TIE: 2, // 和局
  BIGTIE: 3, // 大和局
}

export class CockfightGameLog {
  constructor(data) {
    // 戰鬥結果(對應betArea可以直接使用)
    this.result = result[data.result]
    // 戰鬥組別[紅方id,藍方id]
    this.fightGroup = data.fightgroup
    // 賠率
    this.odds = data.rate
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new HundredPlayerLog(d))
  }
}
