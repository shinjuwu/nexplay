import { BaseTableInput, BaseServersideTableInput } from '@/base/common/table/tableInput'

export class GetUserPlayLogListInput extends BaseTableInput {
  constructor(
    agentId,
    gameId,
    roomType,
    userName,
    logNumber,
    betId,
    singleWalletId,
    startTime,
    endTime,
    length,
    column,
    dir
  ) {
    super(length, column, dir)

    this.agentId = agentId
    this.gameId = gameId
    this.roomType = roomType
    this.userName = userName
    this.logNumber = logNumber
    this.betId = betId
    this.singleWalletId = singleWalletId
    this.startTime = startTime
    this.endTime = endTime
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      game_id: this.gameId,
      room_type: this.roomType,
      username: this.userName,
      lognumber: this.logNumber,
      bet_id: this.betId,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
      single_wallet_id: this.singleWalletId,
    }
  }
}

export class GetBatchUserPlayLogListInput extends BaseServersideTableInput {
  constructor(
    agentId,
    gameId,
    roomType,
    userName,
    logNumber,
    betId,
    betslipStatus,
    startTime,
    endTime,
    singleWalletId,
    roomId,
    length,
    column,
    dir
  ) {
    super(length, column, dir)

    this.agentId = agentId
    this.gameId = gameId
    this.roomType = roomType
    this.userName = userName
    this.logNumber = logNumber
    this.betId = betId
    this.betslipStatus = betslipStatus
    this.startTime = startTime
    this.endTime = endTime
    this.singleWalletId = singleWalletId
    this.roomId = roomId
  }

  parseInputJson() {
    return {
      agent_id: this.agentId,
      game_id: this.gameId,
      room_type: this.roomType,
      username: this.userName,
      lognumber: this.logNumber,
      bet_id: this.betId,
      betslip_status: this.betslipStatus,
      start_time: this.startTime.toISOString(),
      end_time: this.endTime.toISOString(),
      timezone_offset: this.startTime.getTimezoneOffset(),
      start: this.start,
      length: this.length,
      draw: this.draw,
      sort_column: this.column,
      sort_direction: this.dir,
      single_wallet_id: this.singleWalletId,
      room_id: this.roomId,
    }
  }
}
