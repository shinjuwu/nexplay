<template>
  <div class="midasslot-board relative">
    <div v-for="(symbols, row) in props.board" :key="`row__${row}`" class="flex">
      <MidasslotPic
        v-for="(symbol, column) in symbols"
        :key="`row_${row}_column__${column}_symbol__${symbol}`"
        :class="`column-${column}`"
        :symbol="symbol"
      />
    </div>

    <MidasslotWildMultiPic
      v-for="(symbol, posIdx) in props.multiBoard"
      :key="`multiSymbol__${symbol}__${posIdx}`"
      :class="[`pos-${posIdx}`, props.multiBoard.length > 1 ? 'double' : 'single']"
      :symbol="symbol"
    />
    <MidasslotDrawLine
      v-for="drawLine in drawLines"
      :key="`drawLine_${drawLine.key}`"
      :stroke-color="drawLine.color"
      :path="drawLine.path"
    />
  </div>
</template>

<script setup>
import MidasslotPic from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotPic.vue'
import MidasslotWildMultiPic from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotWildMultiPic.vue'
import MidasslotDrawLine from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotDrawLine.vue'
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
