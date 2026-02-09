<template>
  <div class="mb-4 border-y-2 py-2 text-center text-xl font-bold">
    <div>{{ t('fmtTextSlotFreeGameStage', [props.stage]) }}</div>
    <div>{{ t('fmtTextSlotFreeGameMainGameLognumber', [props.mainGameLognumber]) }}</div>
  </div>
  <JumphighslotGameBoard
    class="fg"
    :multi-display="true"
    :board="props.originalBoard"
    :board-active="props.boardActive"
    :multi-board="props.multiBoard"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import JumphighslotGameBoard from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotGameBoard.vue'

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
  originalBoard: {
    type: Array,
    default: () => [],
  },
  boardActive: {
    type: Array,
    default: () => [],
  },
  multiBoard: {
    type: Array,
    default: () => [],
  },
})
</script>
