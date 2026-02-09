<template>
  <div
    class="jumphighslot"
    :class="{
      'jumphighslot-1': props.symbol === 1,
      'jumphighslot-2': props.symbol === 2,
      'jumphighslot-3': props.symbol === 3,
      'jumphighslot-4': props.symbol === 4,
      'jumphighslot-11': props.symbol === 11,
      'jumphighslot-12': props.symbol === 12,
      'jumphighslot-13': props.symbol === 13,
      'jumphighslot-14': props.symbol === 14,
      'jumphighslot-21': props.symbol === 21,
      'jumphighslot-31': props.symbol === 31,
    }"
  >
    <div class="blueframe_jumphighslot"></div>
  </div>
</template>

<script setup>
const props = defineProps({
  symbol: {
    type: Number,
    default: 0,
  },
  active: {
    type: Boolean,
    default: false,
  },
})
</script>
