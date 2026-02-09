<template>
  <div class="mb-4 border-y-2 py-2 text-center text-xl font-bold">
    <div>{{ t('fmtTextSlotFreeGameStage', [props.stage]) }}</div>
    <div>{{ t('fmtTextSlotFreeGameMainGameLognumber', [props.mainGameLognumber]) }}</div>
  </div>

  <MegsharkslotFreeGameBoard
    :left-transform-board="props.leftTransformBoard"
    :board="props.originalBoard"
    :board-active="props.sameBoard ? props.boardActive : null"
  />

  <template v-if="!props.sameBoard">
    <div class="text-center text-6xl">&#8681;</div>

    <MegsharkslotFreeGameBoard
      :left-transform-board="props.leftTransformBoard"
      :board="props.wwBoard"
      :board-active="props.boardActive"
    />
  </template>
</template>

<script setup>
import { useI18n } from 'vue-i18n'

import MegsharkslotFreeGameBoard from '@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotFreeGameBoard.vue'

const { t } = useI18n()
const props = defineProps({
  stage: {
    type: Number,
    default: 0,
  },
  mainGameLognumber: {
    type: String,
    default: '',
  },
  leftTransformBoard: {
    type: Array,
    default: () => [],
  },
  originalBoard: {
    type: Array,
    default: () => [],
  },
  wwBoard: {
    type: Array,
    default: () => [],
  },
  boardActive: {
    type: Array,
    default: () => [],
  },
  sameBoard: {
    type: Boolean,
    default: false,
  },
})
</script>
