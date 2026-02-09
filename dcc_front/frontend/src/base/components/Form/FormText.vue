<template>
  <span class="min-h-[42px] w-full rounded border border-gray-200 py-2 px-3.5 text-gray-500">
    {{ props.content }}
  </span>
</template>

<script setup>
const props = defineProps({
  content: {
    type: String,
    default: '',
  },
})
</script>
