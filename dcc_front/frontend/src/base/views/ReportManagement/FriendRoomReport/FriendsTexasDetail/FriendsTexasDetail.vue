<template>
  <div class="grid max-h-[70vh] grid-cols-2 gap-1 overflow-y-auto">
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextOrderNumber', [friendRoomInfo.id]) }}</div>
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextAgentName', [friendRoomInfo.agentName]) }}</div>
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextFriendRoomInfoHostUserName', [friendRoomInfo.userName]) }}</div>
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextGameName', [t(`game__${friendRoomInfo.gameId}`)]) }}</div>
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextFriendRoomInfoRoomId', [friendRoomInfo.roomId]) }}</div>
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextTotalGameTax', [numberToStr(friendRoomInfo.tax)]) }}</div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomInfoTaxpercent', [friendRoomInfo.taxpercent * 100]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomCreateTime', [time.utcTimeStrToLocalTimeFormat(friendRoomInfo.createTime)]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomEndTime', [time.utcTimeStrToLocalTimeFormat(friendRoomInfo.endTime)]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomInfoCompleteGame', [friendRoomInfo.detail.completeGame]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomInfoTotalGame', [friendRoomInfo.detail.totalGame]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomInfoAnte', [numberToStr(friendRoomInfo.detail.ante)]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomInfoEnterLimit', [numberToStr(friendRoomInfo.detail.enterLimit)]) }}
    </div>
    <div class="col-span-2 md:col-span-1">
      {{ t('fmtTextFriendRoomInfoBetLimit', [numberToStr(friendRoomInfo.detail.betLimit)]) }}
    </div>
    <div class="col-span-2 md:col-span-1">{{ t('fmtTextFriendRoomInfoBetSec', [friendRoomInfo.detail.betSec]) }}</div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'

const { t } = useI18n()

const props = defineProps({
  friendRoomInfo: {
    type: Object,
    default: () => {},
  },
})

const friendRoomInfo = computed(() => {
  const detail = JSON.parse(props.friendRoomInfo.detail)
  return Object.assign({}, props.friendRoomInfo, {
    detail: {
      completeGame: detail.current_game,
      totalGame: detail.total_game,
      ante: detail.ante,
      enterLimit: detail.enter_limit,
      betLimit: detail.bet_limit,
      betSec: detail.bet_sec,
    },
  })
})
</script>
