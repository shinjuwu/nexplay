<template>
  <div class="jumphighslotmulti">
    <div
      class="jumphighslotmultinum"
      :class="{
        'jumphighslotmulti-0': props.symbol === 0,
        'jumphighslotmulti-1': props.symbol === 1,
        'jumphighslotmulti-2': props.symbol === 2,
        'jumphighslotmulti-3': props.symbol === 3,
        'jumphighslotmulti-4': props.symbol === 4,
        'jumphighslotmulti-5': props.symbol === 5,
        'jumphighslotmulti-6': props.symbol === 6,
        'jumphighslotmulti-7': props.symbol === 7,
        'jumphighslotmulti-8': props.symbol === 8,
        'jumphighslotmulti-9': props.symbol === 9,
      }"
    ></div>
  </div>
</template>

<script setup>
const props = defineProps({
  symbol: {
    type: Number,
    default: 0,
  },
})
</script>
