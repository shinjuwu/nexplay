import { getTotalTax } from '@/base/common/gameLog/common'
import { BattlePlayerLog } from '@/base/common/gameLog/battleGame'
import { createBaseFriendRoomInfo } from '@/base/common/gameLog/friendRoomInfo'

const foldState = {
  Deal: 1,
  PreFlop: 1,
  FlopRound: 2,
  TurnRound: 3,
  RiverRound: 4,
}

export class TexasGameLog {
  constructor(data) {
    // 好友房資訊
    this.friendRoomInfo = data.friend_room_info ? new TexasFriendRoomInfo(data.friend_room_info) : null
    // 小盲
    this.smallBlind = this.friendRoomInfo ? this.friendRoomInfo.detail.ante : data.ante
    // 大盲
    this.bigBlind = this.smallBlind * 2
    // 公牌
    this.pulicCards = data.board
    // 總抽水
    this.totalTax = getTotalTax(data.playerlog)
    // 發牌位
    this.dealerSeat = data.dealer + 1
    this.dealerSeatId = data.dealer
    // 玩家資訊
    this.playerLogs = []
    // 遊戲歷程(翻牌前、翻牌圈、轉牌圈、河牌圈)
    this.history = [[], [], [], []]

    const playerNames = {}
    data.playerlog.forEach((d) => {
      playerNames[d.seatId] = d.username
      this.playerLogs.push(new TexasPlayerLog(d))
    })

    // golang map[順序(數字string)]{}
    for (let i = 1; i <= Object.keys(data.history).length; i++) {
      const history = data.history[i]
      const index = foldState[history.State] - 1
      const userName = playerNames[history.SeatId]

      this.history[index].push({
        action: history.Action,
        bet: history.Bet,
        seatId: history.SeatId,
        seat: history.SeatId + 1,
        userName: userName,
        index: i,
      })
    }
  }
}

class TexasPlayerLog extends BattlePlayerLog {
  constructor(data) {
    super(data)
    // 座位
    this.seat = data.seatId + 1
    this.seatId = data.seatId
    // 下注金額
    this.bet = data.bet
    // 手牌
    this.handCards = data.cards
    // 最好的手牌組合
    this.bestCards = data.bestcard
    // 最好的手牌組合牌型
    this.bestCardType = data.cardtype
    // 玩家狀態 (1:棄牌,5:ALL In,9:比牌到最後,10:比牌前贏)
    this.status = data.status
    // 棄牌狀態
    this.foldState = foldState[data.fold_state] || foldState.RiverRound
  }
}

class TexasFriendRoomInfo {
  constructor(friendRoomInfo) {
    const baseFriendRoomInfo = createBaseFriendRoomInfo(friendRoomInfo)
    Object.entries(baseFriendRoomInfo).forEach(([key, value]) => {
      this[key] = value
    })

    const detail = JSON.parse(friendRoomInfo.detail)
    this.detail = {
      currentGame: detail.current_game,
      totalGame: detail.total_game,
      ante: detail.ante,
      enterLimit: detail.enter_limit,
      betLimit: detail.bet_limit,
      betSec: detail.bet_sec,
    }
  }
}
