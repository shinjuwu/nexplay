<template>
  <div class="pyrtreasureslot-MGboard relative">
    <div v-for="(symbols, row) in props.board" :key="`row__${row}`" class="pyrtreasureslotreelsresult">
      <PyrtreausreslotPic
        v-for="(symbol, column) in symbols"
        :key="`row_${row}_column__${column}_symbol__${symbol}`"
        :symbol="symbol"
      />
    </div>
  </div>
</template>

<script setup>
import PyrtreausreslotPic from '@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotPic.vue'
const props = defineProps({
  board: {
    type: Array,
    default: () => [],
  },
})
</script>
