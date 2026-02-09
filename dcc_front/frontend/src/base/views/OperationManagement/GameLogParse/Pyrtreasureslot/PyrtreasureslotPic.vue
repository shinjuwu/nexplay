<template>
  <div
    class="pyrtreasureslot"
    :class="{
      'pyrtreasureslot-1': props.symbol === 1,
      'pyrtreasureslot-2': props.symbol === 2,
      'pyrtreasureslot-3': props.symbol === 3,
      'pyrtreasureslot-4': props.symbol === 4,
      'pyrtreasureslot-5': props.symbol === 5,
      'pyrtreasureslot-11': props.symbol === 11,
      'pyrtreasureslot-12': props.symbol === 12,
      'pyrtreasureslot-13': props.symbol === 13,
      'pyrtreasureslot-14': props.symbol === 14,
      'pyrtreasureslot-15': props.symbol === 15,
      'pyrtreasureslot-22': props.symbol === 22,
      'pyrtreasureslot-31': props.symbol === 31,
    }"
  ></div>
</template>

<script setup>
const props = defineProps({
  symbol: {
    type: Number,
    default: 0,
  },
})
</script>
