<template>
  <span class="poker-container">
    <span
      class="poker"
      :class="{
        'poker-club': poker.suitName === 'club',
        'poker-diamond': poker.suitName === 'diamond',
        'poker-heart': poker.suitName === 'heart',
        'poker-spade': poker.suitName === 'spade',
        'poker-joker poker-red-joker': poker.suitName === 'red-joker',
        'poker-joker poker-black-joker': poker.suitName === 'black-joker',
      }"
    >
      <span>{{ poker.number }}</span>
    </span>
  </span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  index: {
    type: Number,
    default: 0,
    validator: (i) => Number.isInteger(i) && i >= 0 && i <= 53,
  },
  cardPointTexts: {
    type: Array,
    default: () => ['A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K', 'Joker'],
  },
  cardSuitNames: {
    type: Array,
    default: () => ['club', 'diamond', 'heart', 'spade', 'black-joker', 'red-joker'],
  },
  isJoker: {
    type: Function,
    default: (cardIdx) => cardIdx === 52 || cardIdx === 53,
  },
  isBlackJoker: {
    type: Function,
    default: (cardIdx) => cardIdx === 52,
  },
  isRedJoker: {
    type: Function,
    default: (cardIdx) => cardIdx === 53,
  },
})

const poker = computed(() => {
  return {
    number: props.isJoker(props.index) ? props.cardPointTexts[13] : props.cardPointTexts[props.index % 13],
    suitName:
      props.cardSuitNames[
        props.isBlackJoker(props.index) ? 4 : props.isRedJoker(props.index) ? 5 : Math.floor(props.index / 13)
      ],
  }
})
</script>
