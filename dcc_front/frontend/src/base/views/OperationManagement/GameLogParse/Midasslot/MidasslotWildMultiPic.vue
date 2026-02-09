<template>
  <div
    class="midasslotmulti absolute"
    :class="{
      'midasslotmulti-0': props.symbol === 0,
      'midasslotmulti-1': props.symbol === 1,
      'midasslotmulti-2': props.symbol === 2,
      'midasslotmulti-3': props.symbol === 3,
      'midasslotmulti-4': props.symbol === 4,
      'midasslotmulti-5': props.symbol === 5,
      'midasslotmulti-6': props.symbol === 6,
      'midasslotmulti-7': props.symbol === 7,
      'midasslotmulti-8': props.symbol === 8,
      'midasslotmulti-9': props.symbol === 9,
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
