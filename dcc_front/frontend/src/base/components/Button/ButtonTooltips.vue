<template>
  <div class="group relative inline-block cursor-pointer">
    <div class="flex items-center">
      <slot name="content"></slot>
    </div>
    <div class="tooltips text-secondary">{{ props.tipsText }}</div>
  </div>
</template>

<script setup>
const props = defineProps({
  tipsText: {
    type: String,
    default: '',
  },
})
</script>
