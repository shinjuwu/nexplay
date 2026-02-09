<template>
  <JumphighslotGameBoard class="mg" :multi-display="false" :board="props.board" :board-active="props.boardActive" />
</template>

<script setup>
import JumphighslotGameBoard from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotGameBoard.vue'

const props = defineProps({
  board: {
    type: Array,
    default: () => [],
  },
  boardActive: {
    type: Array,
    default: () => [],
  },
})
</script>
