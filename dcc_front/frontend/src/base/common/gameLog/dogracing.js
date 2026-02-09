import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

export class DogracingGameLog {
  constructor(data) {
    const racingDogs = data.result.slice().map((i) => i + 1)

    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new DogracingPlayerLog(d))
    // 前兩名
    this.topTwoDogs = racingDogs.slice(0, 2)

    racingDogs.sort((a, b) => a - b)
    // betArea組合
    this.betAreaDogs = []
    for (let i = 0; i < racingDogs.length; i++) {
      this.betAreaDogs.push([racingDogs[i]])
      for (let j = i + 1; j < racingDogs.length; j++) {
        this.betAreaDogs.push([racingDogs[i], racingDogs[j]])
      }
    }
  }
}

class DogracingPlayerLog extends HundredPlayerLog {
  constructor(data) {
    super(data)

    this.odds = data.odds
  }
}
