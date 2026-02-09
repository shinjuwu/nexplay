export interface NotificationTableProps {
  items: NotificationItem[]
  tableName: string
}

export interface NotificationTransactionRecord {
  id: string
  agentName: string
  playerName: string
  info: string
  occurredTime: Date
}

export interface NotificationGameRecord {
  id: string
  agentName: string
  playerName: string
  gameName: string
  roomName: string
  info: string
  occurredTime: Date
}

export interface MonitorRtpItem {
  rtpType: string
  name: string
  gameType: number
  deScore: number
  yaScore: number
  tax: number
  bonus: number
  rtp: number
  playCount: number
}
