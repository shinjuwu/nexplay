import constant from '@/base/common/constant'
import { round } from '@/base/utils/math'

export function roomTypeNameIndex(gameId, roomType) {
  const gameType = Math.floor(round(gameId / 1000, 4))

  switch (roomType) {
    case constant.RoomType.Newbie:
      return gameType === constant.GameType.Slot || gameType === constant.GameType.FriendsRoom ? '0_1' : '0_0'
    default:
      return roomType
  }
}
