import { PlayerLog } from '@/base/common/gameLog/common'

import { round } from '@/base/utils/math'

export class HundredPlayerLog extends PlayerLog {
  constructor(data) {
    super(data)

    // 起始金額
    this.startScore = round(data.start_score + data.ya_score, 4)

    // 押注區資訊
    this.betAreas = []
    if (data.bet_area) {
      for (let i = 0; i < data.bet_area.length; i++) {
        if (data.bet_area[i] > 0) {
          this.betAreas.push({ betScore: data.bet_area[i], areaId: i })
        }
      }
    }
  }
}

export function getWinAreas(winArea) {
  let winAreas = []
  for (let i = 0; i < winArea.length; i++) {
    if (winArea[i] === 1) {
      winAreas.push(i)
    }
  }
  return winAreas
}
