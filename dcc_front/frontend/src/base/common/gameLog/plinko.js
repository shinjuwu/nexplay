import { ElectronicPlayerLog } from '@/base/common/gameLog/electronicGame'
import { round } from '@/base/utils/math'

export class PlinkoGameLog {
  constructor(data) {
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new PlinkoPlayerLog(d))
  }
}

class PlinkoPlayerLog extends ElectronicPlayerLog {
  constructor(data) {
    super(data)

    const { ballOddsHoleOddsBetCount, ballOddsBetWins } = transformFishstatisticRecord(data.fishstatistic.Record)

    // 球的倍率
    this.ballInfo = data.fishstatistic.BallInfo

    // 不同倍率球的各洞口倍率下注數量
    this.ballOddsHoleOddsBetCount = ballOddsHoleOddsBetCount

    // 可下注金額
    this.betInfo = data.fishstatistic.BetInfo

    // 洞口倍率
    this.holeInfo = data.fishstatistic.HoleInfo.filter(
      (item, index) => data.fishstatistic.HoleInfo.indexOf(item) === index
    ).sort((a, b) => a - b)

    // 不同倍率球的下注金額贏總和
    this.ballOddsBetWins = ballOddsBetWins
  }
}

function transformFishstatisticRecord(record) {
  const ballOddsHoleOddsBetCount = {}
  const ballOddsBetWins = {}

  if (!record) {
    return { ballOddsHoleOddsBetCount, ballOddsBetWins }
  }

  for (const betCollection of Object.values(record)) {
    for (const holeCollection of Object.values(betCollection)) {
      for (const betRecord of Object.values(holeCollection)) {
        const bet = round(betRecord.Win / betRecord.HoleOdds / betRecord.BallOdds / betRecord.Count, 4)
        if (!ballOddsHoleOddsBetCount[betRecord.BallOdds]) {
          ballOddsHoleOddsBetCount[betRecord.BallOdds] = {}
        }
        if (!ballOddsHoleOddsBetCount[betRecord.BallOdds][betRecord.HoleOdds]) {
          ballOddsHoleOddsBetCount[betRecord.BallOdds][betRecord.HoleOdds] = {}
        }
        if (!ballOddsHoleOddsBetCount[betRecord.BallOdds][betRecord.HoleOdds][bet]) {
          ballOddsHoleOddsBetCount[betRecord.BallOdds][betRecord.HoleOdds][bet] = 0
        }
        ballOddsHoleOddsBetCount[betRecord.BallOdds][betRecord.HoleOdds][bet] += round(betRecord.Count, 0)

        if (!ballOddsBetWins[betRecord.BallOdds]) {
          ballOddsBetWins[betRecord.BallOdds] = {}
        }
        if (!ballOddsBetWins[betRecord.BallOdds][bet]) {
          ballOddsBetWins[betRecord.BallOdds][bet] = 0
        }
        ballOddsBetWins[betRecord.BallOdds][bet] += round(betRecord.Win, 4)
      }
    }
  }

  return { ballOddsHoleOddsBetCount, ballOddsBetWins }
}
