<template>
  <div
    class="megsharkslot"
    :class="{
      'megsharkslot-1': props.symbol === 1,
      'megsharkslot-2': props.symbol === 2,
      'megsharkslot-3': props.symbol === 3,
      'megsharkslot-4': props.symbol === 4,
      'megsharkslot-11': props.symbol === 11,
      'megsharkslot-12': props.symbol === 12,
      'megsharkslot-13': props.symbol === 13,
      'megsharkslot-14': props.symbol === 14,
      'megsharkslot-15': props.symbol === 15,
      'megsharkslot-16': props.symbol === 16,
      'megsharkslot-21': props.symbol === 21,
      'megsharkslot-31': props.symbol === 31,
      active: props.active,
    }"
  ></div>
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
