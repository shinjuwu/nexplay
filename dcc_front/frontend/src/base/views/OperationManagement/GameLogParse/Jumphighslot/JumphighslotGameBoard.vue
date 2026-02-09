<template>
  <div class="jumphighslot-board relative">
    <div v-if="props.multiDisplay" class="multiBoard">
      <JumphighslotWildMultiPic
        v-for="(symbol, posIdx) in props.multiBoard"
        :key="`multiSymbol__${symbol}__${posIdx}`"
        :class="[`pos-${posIdx}`, props.multiBoard.length > 1 ? 'double' : 'single']"
        :symbol="symbol"
      />
      <div class="jumphighslotmulti-x"></div>
    </div>

    <div v-for="(symbols, row) in props.board" :key="`row__${row}`" class="jumphighslotreelsresult">
      <JumphighslotPic
        v-for="(symbol, column) in symbols"
        :key="`row_${row}_column__${column}_symbol__${symbol}`"
        :symbol="symbol"
        :class="{ blueframe: props.boardActive[row][column] }"
      />
    </div>
  </div>
</template>

<script setup>
import JumphighslotPic from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotPic.vue'
import JumphighslotWildMultiPic from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotWildMultiPic.vue'
const props = defineProps({
  board: {
    type: Array,
    default: () => [],
  },
  boardActive: {
    type: Array,
    default: () => [],
  },
  multiDisplay: {
    type: Boolean,
    default: false,
  },
  multiBoard: {
    type: Array,
    default: () => [],
  },
})
</script>
