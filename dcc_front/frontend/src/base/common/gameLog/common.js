import { round } from '@/base/utils/math'

export function getTotalTax(playerlog) {
  return playerlog.reduce((s, c) => round(s + c.tax, 4), 0)
}

export class PlayerLog {
  constructor(data) {
    // 玩家名稱
    this.userName = data.username
    // 結算分數
    this.endScore = data.end_score
    // 有效投注
    this.validScore = data.valid_score
    // 壓分
    this.yaScore = data.ya_score
    // 紅利(喜錢)
    this.bonus = data.bonus || 0
    // 得分
    this.deScore = round(data.de_score + this.bonus, 4)
    // 輸贏金額
    this.profitScore = round(this.deScore - data.ya_score - data.tax, 4)
    // 抽水
    this.tax = data.tax
    // JP代幣
    this.jackpotBet = data?.jp_token?.bet || 0
  }
}
