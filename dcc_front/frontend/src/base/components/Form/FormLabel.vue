<template>
  <label class="text-gray-500">{{ props.content }}</label>
</template>

<script setup>
const props = defineProps({
  content: {
    type: String,
    default: '',
  },
})
</script>
