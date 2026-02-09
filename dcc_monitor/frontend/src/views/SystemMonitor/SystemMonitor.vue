<script setup lang="ts">
import type { ReactiveArr } from '@/types/types.common'
import type { NotificationTransactionRecord, NotificationGameRecord, MonitorRtpItem } from '@/types/types.monitorSystem'
import type {
  CoinInOutStatusResponse,
  CoinInOutStatusChangeSet,
  AbnormalWinAndLoseStatusResponse,
  PlatformRTPStatusResponse,
} from '@/types/types.api-monitor'

import { computed, onBeforeMount, onUnmounted, reactive, ref } from 'vue'
import { useRoute } from 'vue-router'
import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/monitor'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { round } from '@/utils/math'
import timer from '@/utils/timer'

import SystemMonitorTransactionRecord from '@/components/SystemMonitorTransactionRecord.vue'
import SystemMonitorGameRecord from '@/components/SystemMonitorGameRecord.vue'
import SystemMonitorRtp from '@/components/SystemMonitorRtp.vue'

const route = useRoute()
const platform = computed(() => route.params.platform as string)

const transactionRecords = reactive<ReactiveArr<NotificationTransactionRecord>>({ items: [] })
const gameRecords = reactive<ReactiveArr<NotificationGameRecord>>({ items: [] })

const { warn } = useDialogStore()

async function requestCoinInOutStatus() {
  const axiosResp = await api.coinInOutStatus({
    filter: platform.value,
  })

  if (axiosResp.data.code !== responseCode.Success) {
    warn(parseServerErrorMessage(axiosResp))
    return
  }

  const data = JSON.parse(axiosResp.data.data) as CoinInOutStatusResponse
  if (data.coin_inout_status_list === null) {
    transactionRecords.items = []
    return
  }

  transactionRecords.items = data.coin_inout_status_list.map((d) => {
    const changeSet = JSON.parse(d.changeset) as CoinInOutStatusChangeSet
    const coin = changeSet.add_coin > 0 ? changeSet.add_coin : -changeSet.add_coin
    return {
      id: d.id,
      agentName: d.agent_name,
      playerName: d.username,
      info: `${changeSet.add_coin > 0 ? '上分' : '下分'} ${coin.toLocaleString()}`,
      occurredTime: new Date(d.create_time),
    }
  })
}

async function requestAbnormalWinAndLoseStatus() {
  const axiosResp = await api.abnormalWinAndLoseStatus({
    filter: platform.value,
  })

  if (axiosResp.data.code !== responseCode.Success) {
    warn(parseServerErrorMessage(axiosResp))
    return
  }

  const data = JSON.parse(axiosResp.data.data) as AbnormalWinAndLoseStatusResponse
  if (data.abnormal_win_and_lose_status_list === null) {
    gameRecords.items = []
    return
  }

  gameRecords.items = data.abnormal_win_and_lose_status_list.map((d) => {
    return {
      id: d.bet_id,
      agentName: d.agent_name,
      playerName: d.username,
      gameName: d.game_name,
      roomName: d.room_name,
      info: `赢 ${d.de_score.toLocaleString()}`,
      occurredTime: new Date(d.create_time),
    }
  })
}

const gameTypes = reactive({
  items: [
    {
      text: '全部游戏',
      value: 0,
    },
    {
      text: '百人游戏',
      value: 1,
    },
    {
      text: '棋牌游戏',
      value: 2,
    },
    {
      text: '电子游戏',
      value: 3,
    },
    {
      text: '老虎机',
      value: 4,
    },
  ],
})
const gameType = ref(gameTypes.items[0].value)

const sorts = reactive({
  items: [
    {
      text: 'RTP数值递减排序',
      value: {
        sort: 'desc',
        column: 'rtp',
      },
    },
    {
      text: 'RTP数值递增排序',
      value: {
        sort: 'asc',
        column: 'rtp',
      },
    },
    {
      text: '场次数值递减排序',
      value: {
        sort: 'desc',
        column: 'playCount',
      },
    },
    {
      text: '场次数值递增排序',
      value: {
        sort: 'asc',
        column: 'playCount',
      },
    },
    {
      text: '得分数值递减排序',
      value: {
        sort: 'desc',
        column: 'deScore',
      },
    },
    {
      text: '得分数值递增排序',
      value: {
        sort: 'asc',
        column: 'deScore',
      },
    },
    {
      text: '押分数值递减排序',
      value: {
        sort: 'desc',
        column: 'yaScore',
      },
    },
    {
      text: '押分数值递增排序',
      value: {
        sort: 'asc',
        column: 'yaScore',
      },
    },
  ],
})
const sort = ref(sorts.items[2].value)

const monitorDay = reactive<MonitorRtpItem>({
  rtpType: 'day',
  name: '',
  gameType: -1,
  deScore: 0,
  yaScore: 0,
  tax: 0,
  bonus: 0,
  rtp: 0,
  playCount: 0,
})
const monitorWeek = reactive<MonitorRtpItem>({
  rtpType: 'week',
  name: '',
  gameType: -1,
  deScore: 0,
  yaScore: 0,
  tax: 0,
  bonus: 0,
  rtp: 0,
  playCount: 0,
})
const monitorMonth = reactive<MonitorRtpItem>({
  rtpType: 'month',
  name: '',
  gameType: -1,
  deScore: 0,
  yaScore: 0,
  tax: 0,
  bonus: 0,
  rtp: 0,
  playCount: 0,
})
const monitorGames = reactive<ReactiveArr<MonitorRtpItem>>({
  items: [],
})
const viewMonitorGames = computed(() => {
  let result = monitorGames.items.slice()
  if (gameType.value !== 0) {
    result = result.filter((g) => g.gameType === gameType.value)
  }
  result.sort((a, b) => {
    switch (sort.value.column) {
      case 'playCount':
        return sort.value.sort === 'asc' ? a.playCount - b.playCount : b.playCount - a.playCount
      case 'deScore':
        return sort.value.sort === 'asc' ? a.deScore - b.deScore : b.deScore - a.deScore
      case 'yaScore':
        return sort.value.sort === 'asc' ? a.yaScore - b.yaScore : b.yaScore - a.yaScore
      case 'rtp':
        return sort.value.sort === 'asc' ? a.rtp - b.rtp : b.rtp - a.rtp
      default:
        return 0
    }
  })
  return result
})

async function requestPlatformRtpStatus() {
  const axiosResp = await api.platformRtpStatus({
    filter: platform.value,
    time_zone: new Date().getTimezoneOffset(),
  })

  if (axiosResp.data.code !== responseCode.Success) {
    warn(parseServerErrorMessage(axiosResp))
    return
  }

  const data = JSON.parse(axiosResp.data.data) as PlatformRTPStatusResponse
  if (data.rtp_status_list === null) {
    monitorDay.rtp = 0
    monitorWeek.rtp = 0
    monitorMonth.rtp = 0
    monitorGames.items = []
    return
  }

  const games = [] as MonitorRtpItem[]
  for (let i = 0; i < data.rtp_status_list.length; i++) {
    const d = data.rtp_status_list[i]
    const rtp = d.ya > 0 ? round(d.de / d.ya, 4) : 0

    switch (d.rtp_type) {
      case monitorDay.rtpType:
        monitorDay.rtp = rtp
        monitorDay.name = d.title
        monitorDay.deScore = d.de
        monitorDay.yaScore = d.ya
        monitorDay.tax = d.tax
        monitorDay.bonus = d.bonus
        monitorDay.playCount = d.play_count
        break
      case monitorWeek.rtpType:
        monitorWeek.rtp = rtp
        monitorWeek.name = d.title
        monitorWeek.deScore = d.de
        monitorWeek.yaScore = d.ya
        monitorWeek.tax = d.tax
        monitorWeek.bonus = d.bonus
        monitorWeek.playCount = d.play_count
        break
      case monitorMonth.rtpType:
        monitorMonth.rtp = rtp
        monitorMonth.name = d.title
        monitorMonth.deScore = d.de
        monitorMonth.yaScore = d.ya
        monitorMonth.tax = d.tax
        monitorMonth.bonus = d.bonus
        monitorMonth.playCount = d.play_count
        break
      default:
        games.push({
          rtpType: d.rtp_type,
          name: d.title,
          gameType: d.game_type,
          deScore: d.de,
          yaScore: d.ya,
          tax: d.tax,
          bonus: d.bonus,
          playCount: d.play_count,
          rtp,
        })
        break
    }
  }
  monitorGames.items = games
}

function requestAll() {
  processApiRequest(async () => {
    await requestCoinInOutStatus()
    await requestAbnormalWinAndLoseStatus()
    await requestPlatformRtpStatus()
  }, warn)
}

let timerId: number

onBeforeMount(() => {
  processApiRequest(async () => {
    await requestCoinInOutStatus()
    await requestAbnormalWinAndLoseStatus()
    await requestPlatformRtpStatus()
    timerId = timer.register(60 * 1000, requestAll)
  }, warn)
})

onUnmounted(() => {
  timer.unregister(timerId)
})
</script>

<template>
  <section class="relative rounded bg-white p-4">
    <div class="grid gap-2 lg:grid-cols-2">
      <SystemMonitorTransactionRecord :items="transactionRecords.items" />
      <SystemMonitorGameRecord :items="gameRecords.items" />
    </div>

    <hr class="my-4 border-8" />

    <div
      class="grid grid-cols-2 gap-2 sm:grid-cols-3 sm:gap-4 md:grid-cols-4 md:gap-6 lg:grid-cols-5 lg:gap-8 xl:grid-cols-6 xl:gap-10"
    >
      <SystemMonitorRtp
        :name="monitorDay.name"
        :de-score="monitorDay.deScore"
        :ya-score="monitorDay.yaScore"
        :tax="monitorDay.tax"
        :bonus="monitorDay.bonus"
        :rtp="monitorDay.rtp"
        :play-count="monitorDay.playCount"
      />
      <SystemMonitorRtp
        :name="monitorWeek.name"
        :de-score="monitorWeek.deScore"
        :ya-score="monitorWeek.yaScore"
        :tax="monitorWeek.tax"
        :bonus="monitorWeek.bonus"
        :rtp="monitorWeek.rtp"
        :play-count="monitorWeek.playCount"
      />
      <SystemMonitorRtp
        :name="monitorMonth.name"
        :de-score="monitorMonth.deScore"
        :ya-score="monitorMonth.yaScore"
        :tax="monitorMonth.tax"
        :bonus="monitorMonth.bonus"
        :rtp="monitorMonth.rtp"
        :play-count="monitorMonth.playCount"
      />

      <div class="text-danger font-bold sm:col-span-3 md:col-span-1 lg:col-span-2 xl:col-span-3">
        <div>RTP监控:</div>
        <div>※ 各资料每 1 分钟自动刷新</div>
      </div>
    </div>

    <hr class="my-4 border-8" />

    <div class="grid sm:grid-cols-2 sm:gap-2 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6">
      <select
        v-model="gameType"
        class="mb-2 w-full rounded border px-4 py-2 md:col-start-2 lg:col-start-4 xl:col-start-5"
      >
        <option v-for="gt in gameTypes.items" :key="`gameType__${gt.value}`" :value="gt.value">{{ gt.text }}</option>
      </select>

      <select v-model="sort" class="mb-2 w-full rounded border px-4 py-2">
        <option v-for="s in sorts.items" :key="`sort__${s.value.column}__${s.value.sort}`" :value="s.value">{{ s.text }}</option>
      </select>
    </div>
    <div
      class="grid grid-cols-2 gap-2 sm:grid-cols-3 sm:gap-4 md:grid-cols-4 md:gap-6 lg:grid-cols-5 lg:gap-8 xl:grid-cols-6 xl:gap-10"
    >
      <SystemMonitorRtp
        v-for="game in viewMonitorGames"
        :key="`monitorGame__${game.rtpType}`"
        :name="game.name"
        :de-score="game.deScore"
        :ya-score="game.yaScore"
        :tax="game.tax"
        :bonus="game.bonus"
        :rtp="game.rtp"
        :play-count="game.playCount"
      />
    </div>
  </section>
</template>
