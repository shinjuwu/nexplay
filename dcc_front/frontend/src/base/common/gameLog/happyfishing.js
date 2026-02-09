import { RcfishingPlayerLog } from '@/base/common/gameLog/rcfishing'

export class HappyfishingGameLog {
  constructor(data) {
    // 玩家資訊
    this.playerLog = data.playerlog.map((d) => new RcfishingPlayerLog(d))[0]
  }
}
