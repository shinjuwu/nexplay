<template>
  <div class="megsharkslot-board megsharkslot-fg-board relative">
    <MegsharkslotPic
      v-for="(symbol, idx) in props.leftTransformBoard"
      :key="`tranform_symbol__${symbol}__${idx}`"
      :class="{
        'transform p1': idx === 0,
        'transform p2': idx === 1,
        'transform p3': idx === 2,
      }"
      :symbol="symbol"
    />

    <div v-for="(symbols, row) in props.board" :key="`row__${row}`" class="flex">
      <MegsharkslotPic
        v-for="(symbol, column) in symbols"
        :key="`row_${row}_column__${column}_symbol__${symbol}`"
        :symbol="symbol"
        :active="props.boardActive ? props.boardActive[row][column] : false"
      />
    </div>
  </div>
</template>

<script setup>
import MegsharkslotPic from '@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotPic.vue'

const props = defineProps({
  leftTransformBoard: {
    type: Array,
    default: () => [],
  },
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
