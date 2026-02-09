<template>
  <svg height="338" width="661" class="absolute top-[15px] left-[22px]">
    <line
      v-for="(line, lineIdx) in path"
      :key="`line__${lineIdx}`"
      :x1="line.x1"
      :y1="line.y1"
      :x2="line.x2"
      :y2="line.y2"
      style="stroke-width: 3"
      :style="`stroke: ${props.strokeColor}`"
    />
  </svg>
</template>

<script setup>
const props = defineProps({
  strokeColor: {
    type: String,
    default: '',
  },
  path: {
    type: Array,
    default: () => [],
  },
})
</script>
