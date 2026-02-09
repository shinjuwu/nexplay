<template>
  <div class="megsharkslot-board megsharkslot-mg-board">
    <div v-for="(symbols, row) in props.board" :key="`row__${row}`" class="flex">
      <MegsharkslotPic
        v-for="(symbol, column) in symbols"
        :key="`row_${row}_column__${column}_symbol__${symbol}`"
        :symbol="symbol"
        :active="props.boardActive[row][column]"
      />
    </div>
  </div>
</template>

<script setup>
import MegsharkslotPic from '@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotPic.vue'

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
