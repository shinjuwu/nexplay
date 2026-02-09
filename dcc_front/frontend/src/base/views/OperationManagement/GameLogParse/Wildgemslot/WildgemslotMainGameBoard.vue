<template>
  <WildgemslotGameBoard :board="props.board" :draw-lines="props.drawLines" />
</template>

<script setup>
import WildgemslotGameBoard from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotGameBoard.vue'

const props = defineProps({
  board: {
    type: Array,
    default: () => [],
  },
  drawLines: {
    type: Array,
    default: () => [],
  },
})
</script>
