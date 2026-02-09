import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'

export class ChinesepokerGameLog {
  constructor(data) {
    // 底注
    this.baseBet = data.basebet
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => {
      const shoot = data.playerrecord.shoot[d.seatId]
      const getShoot = data.playerrecord.getshoot[d.seatId]
      const playerRecord = data.playerrecord.playerRecord.find((r) => r.seatId === d.seatId)
      return new ChinesepokerPlayerLog(d, shoot, getShoot, playerRecord)
    })
  }
}

class ChinesepokerPlayerLog extends BattlePlayerLog {
  constructor(data, shoot, getShoot, playerRecord) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    // 打槍
    this.shoot = shoot || 0
    // 被打槍
    this.getShoot = getShoot || 0

    // 整副手牌牌型
    this.handCardType = playerRecord.playInfo.handCardType
    // 是否為特殊牌型
    this.isSpecialCardType = this.handCardType > 0

    // [頭墩,中墩,尾墩]擺牌
    this.cardsArray = playerRecord.playInfo.cardsArray
    // [頭墩,中墩,尾墩]牌型
    this.arrayType = playerRecord.playInfo.arrayType

    // 總計水數
    this.point = playerRecord.point
    // 結果
    this.resultScore = playerRecord.resultScore
  }
}
