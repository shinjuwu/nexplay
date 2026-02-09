import { HundredPlayerLog } from '@/base/common/gameLog/hundredGame'

const result = {
  Winner: 3, // 0:和 1:莊贏 2:閒贏
  BankerPair: 4, // 莊對
  PlayerPair: 8, // 閒對
  WinSmallThanOther: 16, // 上莊贏
  WinBiggerThanOther: 32, // 上莊輸
  Big: 64, // 大
  Small: 128, // 小
}

export class BaccaratGameLog {
  constructor(data) {
    // 莊家手牌資訊
    this.banker = new BaccaratCards(data.banker)
    // 閒家手牌資訊
    this.player = new BaccaratCards(data.player)
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new HundredPlayerLog(d))
    // 莊閒和資訊
    this.winner = data.result & result.Winner
    // 是否開莊對
    this.bankerPair = data.result & result.BankerPair
    // 是否開閒對
    this.playerPair = data.result & result.PlayerPair
    // 是否開上莊贏
    this.winSmallThanOther = data.result & result.WinSmallThanOther
    // 是否開上莊輸
    this.winBiggerThanOther = data.result & result.WinBiggerThanOther
    // 是否開大
    this.big = data.result & result.Big
    // 是否開小
    this.small = data.result & result.Small
    // 目前第幾局
    this.round = data.round
  }
}

class BaccaratCards {
  constructor(cards) {
    this.cards = cards
    this.point = cards.reduce((point, cardIdx) => {
      let card = (cardIdx + 1) % 13
      return card < 10 ? (point + card) % 10 : point
    }, 0)
  }
}
