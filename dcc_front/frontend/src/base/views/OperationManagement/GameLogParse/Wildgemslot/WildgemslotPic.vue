<template>
  <div
    class="wildgemslot"
    :class="{
      'wildgemslot-1': props.symbol === 1,
      'wildgemslot-2': props.symbol === 2,
      'wildgemslot-3': props.symbol === 3,
      'wildgemslot-4': props.symbol === 4,
      'wildgemslot-5': props.symbol === 5,
      'wildgemslot-11': props.symbol === 11,
      'wildgemslot-12': props.symbol === 12,
      'wildgemslot-13': props.symbol === 13,
      'wildgemslot-14': props.symbol === 14,
      'wildgemslot-15': props.symbol === 15,
      'wildgemslot-21': props.symbol === 21,
      'wildgemslot-31': props.symbol === 31,
      'wildgemslot-32 wildgemslot-wild': props.symbol === 32,
      'wildgemslot-33 wildgemslot-wild': props.symbol === 33,
      'wildgemslot-34 wildgemslot-wild': props.symbol === 34,
      'wildgemslot-35 wildgemslot-wild': props.symbol === 35,
      'wildgemslot-36 wildgemslot-wild': props.symbol === 36,
      'wildgemslot-37 wildgemslot-wild': props.symbol === 37,
      'wildgemslot-38 wildgemslot-wild': props.symbol === 38,
      'wildgemslot-39 wildgemslot-wild': props.symbol === 39,
      'wildgemslot-40 wildgemslot-wild': props.symbol === 40,
      'wildgemslot-41 wildgemslot-wild': props.symbol === 41,
      'wildgemslot-42 wildgemslot-wild': props.symbol === 42,
      'wildgemslot-43 wildgemslot-wild': props.symbol === 43,
      'wildgemslot-44 wildgemslot-wild': props.symbol === 44,
      'wildgemslot-45 wildgemslot-wild': props.symbol === 45,
      'wildgemslot-46 wildgemslot-wild': props.symbol === 46,
      'wildgemslot-47 wildgemslot-wild': props.symbol === 47,
      'wildgemslot-48 wildgemslot-wild': props.symbol === 48,
      'wildgemslot-49 wildgemslot-wild': props.symbol === 49,
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
