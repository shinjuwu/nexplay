<template>
  <input
    v-model="inputModel"
    class="w-full rounded border border-gray-200 py-2 px-3.5 text-gray-500 outline-none placeholder:text-gray-400 focus-visible:border-gray-300"
  />
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number, Date, Object],
    required: true,
  },
})

const emit = defineEmits(['update:modelValue'])

const inputModel = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  },
})
</script>
