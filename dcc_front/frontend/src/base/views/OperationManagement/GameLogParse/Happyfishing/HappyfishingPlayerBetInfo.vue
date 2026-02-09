<template>
  <div>
    <div>{{ t('textSeatPlayers') }}:</div>
    <i18n-t keypath="fmtTextBlackjackPlayerLog" tag="p" scope="global">
      <template #seat>{{ props.playerLog.seat }}(seatId:{{ props.playerLog.seatId }})</template>
      <template #userName>
        <span :class="{ 'text-primary': props.targetUserName === playerLog.userName }">
          {{ playerLog.userName }}
        </span>
      </template>
      <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
      <template #betDetail>{{ numberToStr(playerLog.yaScore) }}</template>
    </i18n-t>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'

const props = defineProps({
  playerLog: {
    type: Object,
    default: () => {},
  },
  targetUserName: {
    type: String,
    default: '',
  },
})
const { t } = useI18n()
</script>
