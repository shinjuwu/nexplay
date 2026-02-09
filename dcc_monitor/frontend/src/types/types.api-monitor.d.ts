export interface CoinInOutStatusRequest {
  filter: string
}

export interface CoinInOutStatus {
  id: string
  agent_id: number
  agent_name: string
  user_id: number
  username: string
  create_time: string
  kind: number
  changeset: string
  status: number
  info: string
}

export interface CoinInOutStatusResponse {
  coin_inout_status_list: CoinInOutStatus[]
}

export interface CoinInOutStatusChangeSet {
  add_coin: number
  after_coin: number
  before_coin: number
}

export interface AbnormalWinAndLoseStatusRequest {
  filter: string
}

export interface AbnormalWinAndLoseStatus {
  lognumber: string
  agent_id: number
  agent_name: string
  user_id: number
  username: string
  game_id: number
  game_name: string
  room_type: number
  room_name: string
  de_score: number
  bonus: number
  bet_id: string
  bet_time: string
  create_time: string
}

export interface AbnormalWinAndLoseStatusResponse {
  abnormal_win_and_lose_status_list: AbnormalWinAndLoseStatus[]
}

export interface PlatformRTPStatusRequest {
  filter: string
  time_zone: number
}

export interface RTPStatus {
  rtp_type: string
  title: string
  game_type: number
  de: number
  ya: number
  tax: number
  bonus: number
  play_count: number
}

export interface PlatformRTPStatusResponse {
  rtp_status_list: RTPStatus[]
}

export interface ServiceStatus {
  name: string
  info: string
  sub_name: string
  status: number
}

export interface ServiceStatusRequest {
  filter: string
}

export interface ServiceStatusResponse {
  status_list: ServiceStatus[]
}
