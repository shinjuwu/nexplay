<template>
  <MidasslotGameBoard :board="props.board" :multi-board="props.multiBoard" :draw-lines="props.drawLines" class="mg" />
</template>

<script setup>
import MidasslotGameBoard from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotGameBoard.vue'

const props = defineProps({
  board: {
    type: Array,
    default: () => [],
  },
  multiBoard: {
    type: Array,
    default: () => [],
  },
  drawLines: {
    type: Array,
    default: () => [],
  },
})
</script>
