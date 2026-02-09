<template>
  <div class="flex">
    <input
      :id="props.id"
      v-model="inputModel"
      class="flex-1 rounded-l border border-gray-200 py-2 px-3.5 text-gray-500 outline-none placeholder:text-gray-400 focus-visible:border-gray-300"
      :type="props.type"
      :placeholder="props.placeholder"
      :autocomplete="props.type === 'password'"
    />
    <div
      class="ml-[-1px] flex cursor-pointer items-center rounded-r border border-gray-200 bg-gray-100 py-2 px-3.5"
      @click="emit('clickIcon')"
    >
      <span class="text-gray-500">
        <slot name="icon"></slot>
      </span>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number, Date],
    required: true,
  },
  type: {
    type: String,
    default: 'text',
  },
  placeholder: {
    type: String,
    default: '',
  },
  id: {
    type: String,
    default: '',
  },
})

const emit = defineEmits(['update:modelValue', 'clickIcon'])

const inputModel = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  },
})
</script>
