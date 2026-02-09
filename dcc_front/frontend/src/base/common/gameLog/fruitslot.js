import { ElectronicPlayerLog } from '@/base/common/gameLog/electronicGame'

export class FruitslotGameLog {
  constructor(data) {
    // 畫面棋盤
    this.boards = fruitslotBoards
    // 下注按鈕
    this.betButtons = betButtons
    // 底注
    this.baseBet = data.basebet
    // 圖標倍率
    this.rates = data.rate
    // 玩家資訊
    this.playerLogs = data.playerlog.map((d) => new FruitslotPlayerLog(d, this.baseBet))
  }
}

class FruitslotPlayerLog extends ElectronicPlayerLog {
  constructor(data, baseBet) {
    super(data)

    // 下注資訊
    this.betDetails = data.bet_area.map((bet) => {
      return {
        bet: bet,
        betCount: parseInt(bet / baseBet),
      }
    })

    // 水果機歷程
    this.detail = new FruitslotPlayerLogDetail(data.fruit)
  }
}

class FruitslotPlayerLogDetail {
  constructor(data) {
    // 一般獎項
    this.prizeType = data.priceType

    // notice: spPriceType 0 為 0燈獎次處刻意不顯示0燈獎訊息
    if (data.spPriceType) {
      this.spPrizeType = data.spPriceType
      this.spPrizes = getSpPrizes(data.spPriceType, data.spPrice)
    }

    // 猜大小歷程
    this.guessDetails = data.guess ? data.guess.map((detail) => new FruitslotPlayerLogGuessDetail(detail)) : null
  }
}

const spPrizeType = {
  RandomeLamp1: 1, // 隨機1燈
  RandomeLamp2: 2, // 隨機2燈
  RandomeLamp3: 3, // 隨機3燈
  RandomeLamp4: 4, // 隨機4燈
  BellGrapeLemon: 5, // 小三元
  SevenDiamondMangosteen: 6, // 大三元
  FourApples: 7, // 大四喜
  AllTop: 8, // 縱橫四海-上
  AllBottom: 9, // 縱橫四海-下
  AllLeft: 10, // 縱橫四海-左
  AllRight: 11, // 縱橫四海-右
}

function getSpPrizes(inputSpPrizeType, inputSpPrizes) {
  // 請對應borads
  switch (inputSpPrizeType) {
    case spPrizeType.SevenDiamondMangosteen:
      return [4, 12, 16]
    case spPrizeType.FourApples:
      return [1, 7, 13, 19]
    case spPrizeType.AllTop:
      return [0, 1, 2, 3, 4, 20, 21, 22, 23]
    case spPrizeType.AllBottom:
      return [8, 9, 10, 11, 12, 13, 14, 15, 16]
    case spPrizeType.AllLeft:
      return [16, 17, 18, 19, 20]
    case spPrizeType.AllRight:
      return [4, 5, 6, 7, 8]
    default:
      return inputSpPrizes || []
  }
}

class FruitslotPlayerLogGuessDetail {
  constructor(data) {
    // 下注金額
    this.bet = data.Bet
    // 下注的類型(1:大,2:小)
    this.betType = data.BetBig ? 1 : 2
    // 猜大小的數字
    this.result = data.Num
    // 猜大小的類型(1:大,2:小)
    this.resultType = data.IsBig ? 1 : 2
  }
}

class FruitslotSymbol {
  constructor(name, prizeType) {
    this.name = name
    this.prizeType = prizeType
  }
}

class FruitslotBetButton {
  constructor(name, betIndex, rateIndex) {
    this.name = name
    this.betIndex = betIndex
    this.rateIndex = rateIndex
  }
}

// 棋盤空符號
const emptSymbol = new FruitslotSymbol('', -1)
// 棋盤 9 * 5 部分是挖空
const fruitslotBoards = [
  [
    new FruitslotSymbol('3xbell', 20),
    new FruitslotSymbol('lemon', 21),
    new FruitslotSymbol('bell', 22),
    new FruitslotSymbol('50xcrown', 23),
    new FruitslotSymbol('crown', 0),
    new FruitslotSymbol('apple', 1),
    new FruitslotSymbol('3xapple', 2),
    new FruitslotSymbol('grape', 3),
    new FruitslotSymbol('mangosteen', 4),
  ],
  [
    new FruitslotSymbol('apple', 19),
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    new FruitslotSymbol('3xmangosteen', 5),
  ],
  [
    new FruitslotSymbol('clover', 18),
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    new FruitslotSymbol('clover', 6),
  ],
  [
    new FruitslotSymbol('3xdiamond', 17),
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    emptSymbol,
    new FruitslotSymbol('apple', 7),
  ],
  [
    new FruitslotSymbol('diamond', 16),
    new FruitslotSymbol('grape', 15),
    new FruitslotSymbol('3xgrape', 14),
    new FruitslotSymbol('apple', 13),
    new FruitslotSymbol('seven', 12),
    new FruitslotSymbol('3xseven', 11),
    new FruitslotSymbol('bell', 10),
    new FruitslotSymbol('lemon', 9),
    new FruitslotSymbol('3xlemon', 8),
  ],
]

const betButtons = [
  new FruitslotBetButton('crown', 7, 15),
  new FruitslotBetButton('seven', 6, 13),
  new FruitslotBetButton('diamond', 5, 11),
  new FruitslotBetButton('mangosteen', 4, 9),
  new FruitslotBetButton('bell', 3, 7),
  new FruitslotBetButton('grape', 2, 5),
  new FruitslotBetButton('lemon', 1, 3),
  new FruitslotBetButton('apple', 0, 1),
]
