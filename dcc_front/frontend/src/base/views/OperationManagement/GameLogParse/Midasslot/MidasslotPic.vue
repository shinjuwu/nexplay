<template>
  <div
    class="midasslot"
    :class="{
      'midasslot-1': props.symbol === 1,
      'midasslot-2': props.symbol === 2,
      'midasslot-3': props.symbol === 3,
      'midasslot-4': props.symbol === 4,
      'midasslot-5': props.symbol === 5,
      'midasslot-11': props.symbol === 11,
      'midasslot-12': props.symbol === 12,
      'midasslot-13': props.symbol === 13,
      'midasslot-14': props.symbol === 14,
      'midasslot-15': props.symbol === 15,
      'midasslot-21': props.symbol === 21,
      'midasslot-31': props.symbol === 31,
      'midasslot-32': props.symbol === 32,
      'midasslot-33': props.symbol === 33,
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
