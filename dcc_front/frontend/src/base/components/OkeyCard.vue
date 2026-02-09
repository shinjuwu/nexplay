<template>
  <span class="okeycard-container">
    <span v-if="isSahteOkey && card.color === indicatorColor" class="pb-2 text-xl">&#9733;</span>
    <span
      v-else-if="card.number !== 52"
      class="okeycard"
      :class="{
        'okeycard-blue': card.color === 'blue',
        'okeycard-red': card.color === 'red',
        'okeycard-yellow': card.color === 'yellow',
      }"
    >
      {{ card.number }}
    </span>
    <span v-else-if="card.number === 52 && card.color === 'blank'" class="okeycard"></span>
  </span>
</template>

<script setup>
import { computed } from 'vue'
const props = defineProps({
  index: {
    type: Number,
    default: 0,
    validator: (i) => Number.isInteger(i) && i >= 0 && i <= 52,
  },
  cardColor: {
    type: Array,
    default: () => ['blue', 'red', 'black', 'yellow', 'blank'],
  },
  isJoker: {
    type: Function,
    default: (cardIdx) => cardIdx === 52,
  },
  indicatorIndex: {
    type: Number,
    default: -1,
  },
})

const card = computed(() => {
  return {
    number: props.isJoker(props.index) ? props.index : (props.index % 13) + 1,
    color: props.isJoker(props.index) ? props.cardColor[4] : props.cardColor[Math.floor(props.index / 13)],
  }
})

const isSahteOkey = computed(() => {
  const nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]
  return nums[props.index % 13] === nums[(props.indicatorIndex + 1) % 13]
})

const indicatorColor = computed(() => props.cardColor[Math.floor((props.indicatorIndex + 1) / 13)])
</script>
