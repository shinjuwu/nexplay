<template>
  <label
    class="inline-flex items-center"
    :class="{ 'cursor-pointer': !props.disabled, 'pointer-events-none': props.disabled }"
  >
    <CheckIcon
      class="inline rounded border border-slate-500"
      :class="[
        {
          'bg-light': props.disabled,
          'mr-1': props.content.length > 0,
          'h-4 w-4': props.size === 'md',
          'h-5 w-5': props.size === 'lg',
        },
        props.checked ? 'bg-primary text-white' : 'text-transparent',
      ]"
    />
    <span v-if="props.content">{{ props.content }}</span>
  </label>
</template>

<script setup>
import { CheckIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  checked: {
    type: Boolean,
    default: false,
  },
  content: {
    type: String,
    default: '',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  size: {
    type: String,
    default: 'lg',
    validator: (value) => ['sm', 'md', 'lg', 'xl', '2xl'].indexOf(value) >= 0,
  },
})
</script>
