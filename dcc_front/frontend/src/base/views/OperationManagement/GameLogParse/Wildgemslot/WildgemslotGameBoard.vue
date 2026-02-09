<template>
  <div class="wildgemslot-board relative">
    <template v-for="(symbols, row) in props.board" :key="`row__${row}`">
      <WildgemslotPic
        v-for="(symbol, column) in symbols"
        :key="`row_${row}_column__${column}_symbol__${symbol}`"
        :symbol="symbol"
      />
    </template>

    <WildgemslotDrawLine
      v-for="drawLine in drawLines"
      :key="`drawLine_${drawLine.key}`"
      :stroke-color="drawLine.color"
      :path="drawLine.path"
    />
  </div>
</template>

<script setup>
import WildgemslotPic from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotPic.vue'
import WildgemslotDrawLine from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotDrawLine.vue'

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
