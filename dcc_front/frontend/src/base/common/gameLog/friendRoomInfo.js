export function createBaseFriendRoomInfo(friendRoomInfo) {
  return friendRoomInfo
    ? {
        roomId: friendRoomInfo.room_id,
        hostUserName: friendRoomInfo.username,
        taxpercent: friendRoomInfo.taxpercent,
      }
    : null
}
