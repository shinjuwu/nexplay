<template>
  <i18n-t :keypath="props.keypath" tag="div" scope="global">
    <template #userName>
      <span :class="{ 'text-primary': props.userName === props.playerLog.userName }">
        {{ props.playerLog.userName }}
      </span>
    </template>
    <template #profitScore>
      <span :class="{ 'text-danger': playerLog.profitScore > 0, 'text-teal-500': props.playerLog.profitScore < 0 }">
        {{ numberToStr(props.playerLog.profitScore) }}
      </span>
    </template>
    <template #validScore>{{ numberToStr(props.playerLog.validScore) }}</template>
    <template #balance>{{ numberToStr(props.playerLog.endScore) }}</template>
    <template #tax>{{ numberToStr(props.playerLog.tax) }}</template>
    <template #bonus>
      <span :class="{ 'text-danger': playerLog.bonus > 0, 'text-teal-500': props.playerLog.bonus < 0 }">
        {{ numberToStr(props.playerLog.bonus) }}
      </span>
    </template>
    <template #jackpot>
      <span v-if="playerLog.jackpotBet > 0">
        {{ t('fmtTextPlayerResultGetJackpotBet', [numberToStr(playerLog.jackpotBet)]) }}
      </span>
    </template>
  </i18n-t>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { PlayerLog } from '@/base/common/gameLog/common'
import { numberToStr } from '@/base/utils/formatNumber'

const props = defineProps({
  playerLog: {
    type: PlayerLog,
    default: null,
  },
  userName: {
    type: String,
    default: '',
  },
  keypath: {
    type: String,
    default: 'fmtTextPlayerResult',
  },
})
const { t } = useI18n()
</script>
