<template>
  <div class="mb-4 border-y-2 py-2 text-center text-xl font-bold">
    <div>{{ t('fmtTextSlotFreeGameStage', [props.stage]) }}</div>
    <div>{{ t('fmtTextSlotFreeGameMainGameLognumber', [props.mainGameLognumber]) }}</div>
  </div>

  <WildgemslotGameBoard :board="props.board" :draw-lines="props.drawLines" />
</template>

<script setup>
import { useI18n } from 'vue-i18n'

import WildgemslotGameBoard from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotGameBoard.vue'

const props = defineProps({
  stage: {
    type: Number,
    default: 0,
  },
  mainGameLognumber: {
    type: String,
    default: '',
  },
  board: {
    type: Array,
    default: () => [],
  },
  drawLines: {
    type: Array,
    default: () => [],
  },
})

const { t } = useI18n()
</script>
