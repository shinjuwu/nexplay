import { ElectronicPlayerLog } from '@/base/common/gameLog/electronicGame'
import { round } from '@/base/utils/math'

export class RcfishingGameLog {
  constructor(data) {
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new RcfishingPlayerLog(d))
  }
}

export class RcfishingPlayerLog extends ElectronicPlayerLog {
  constructor(data) {
    super(data)

    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId

    const fishStatistic = data.fishstatistic

    // 子彈資訊
    this.bullets = Object.keys(fishStatistic.BulletValue).map((index) => {
      const value = fishStatistic.BulletValue[index]
      const count = fishStatistic.BulletCount[index] || 0
      return {
        value: value,
        count: count,
        totalBulletGold: value * count,
      }
    })

    /*
      golang struct

      type HitInfo struct {
        Dead     int       // 死亡記數
        TotalBet float64   // 總押
        TotalWin float64   // 總得
        Fish     []HitFish // 個別魚的資料
      }

      type HitFish struct {
        BulletGold float64 // 子彈的值
        Wingold    float64 // 贏分
        Reward     string  // 有沒有得到技能(技能id: s1,s2,s3)
        Ftag       int     // 魚種身分(0:普通  1:特殊  2:Boss)
      }
    */

    // 一般魚資訊
    this.normalFishes = []
    // 特殊魚資訊
    this.specialFishes = []
    // boss魚資訊
    this.bossFishes = []

    if (fishStatistic.HitInfo) {
      const tmpFishes = {}

      // map[魚種]map[子彈index]HitInfo
      Object.keys(fishStatistic.HitInfo).forEach((fishType) => {
        Object.keys(fishStatistic.HitInfo[fishType]).forEach((bulletIndex) => {
          const hitInfo = fishStatistic.HitInfo[fishType][bulletIndex]
          for (const fish of hitInfo.Fish) {
            const odds = round(fish.Wingold / fish.BulletGold, 4)

            let fishKey = `${fishType}_${odds}`
            if (fish.Reward) {
              fishKey = `${fishKey}_${fish.Reward}`
            }

            if (!tmpFishes[fishKey]) {
              tmpFishes[fishKey] = {
                name: fishType,
                odds: odds,
                catchInfo: [],
                fTag: fish.Ftag,
              }

              if (fish.Reward) {
                tmpFishes[fishKey].reward = fish.Reward
              }

              this.bullets.forEach(() => {
                tmpFishes[fishKey].catchInfo.push({
                  count: 0,
                  totalGold: 0,
                })
              })
            }

            tmpFishes[fishKey].catchInfo[bulletIndex].count++
            tmpFishes[fishKey].catchInfo[bulletIndex].totalGold = round(
              tmpFishes[fishKey].catchInfo[bulletIndex].totalGold + fish.Wingold,
              4
            )
          }
        })
      })

      Object.values(tmpFishes).forEach((tmpFish) => {
        if (tmpFish.fTag === 0) {
          this.normalFishes.push(tmpFish)
        } else if (tmpFish.fTag === 1) {
          this.specialFishes.push(tmpFish)
        } else if (tmpFish.fTag === 2) {
          this.bossFishes.push(tmpFish)
        }
      })
    }

    // 技能資訊
    this.skills = []

    if (fishStatistic.SkillCastInfo) {
      const tmpSkills = {}
      const tmpFishes = {}

      // map[技能id]map[魚種]map[子彈index]HitInfo
      Object.keys(fishStatistic.SkillCastInfo).forEach((skillId) => {
        Object.keys(fishStatistic.SkillCastInfo[skillId]).forEach((fishType) => {
          Object.keys(fishStatistic.SkillCastInfo[skillId][fishType]).forEach((bulletIndex) => {
            const bulletGold = fishStatistic.BulletValue[bulletIndex]
            const skillKey = `${skillId}_${bulletGold}`

            if (!tmpSkills[skillKey]) {
              tmpSkills[skillKey] = { fishes: [], bulletGold: bulletGold, name: skillId }
            }

            const hitInfo = fishStatistic.SkillCastInfo[skillId][fishType][bulletIndex]

            for (const fish of hitInfo.Fish) {
              const odds = round(fish.Wingold / fish.BulletGold, 4)

              let fishKey = `${skillId}_${fishType}_${odds}`
              if (fish.Reward) {
                fishKey = `${fishKey}_${fish.Reward}`
              }

              if (!tmpFishes[fishKey]) {
                tmpFishes[fishKey] = {
                  name: fishType,
                  odds: odds,
                  count: 0,
                  totalGold: 0,
                }

                if (fish.Reward) {
                  tmpFishes[fishKey].reward = fish.Reward
                }

                tmpSkills[skillKey].fishes.push(tmpFishes[fishKey])
              }

              tmpFishes[fishKey].count++
              tmpFishes[fishKey].totalGold = round(tmpFishes[fishKey].totalGold + fish.Wingold, 4)
            }
          })
        })
      })

      this.skills = this.skills.concat(Object.values(tmpSkills))
    }
  }
}
