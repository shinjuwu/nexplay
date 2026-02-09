<template>
  <div ref="scopeElRef" class="relative">
    <input
      :class="props.calendarAlign === 'left' ? 'rounded-l' : 'rounded-r'"
      type="text"
      readonly
      :value="props.modelValue ? format(props.modelValue, 'yyyy-MM') : ''"
      class="w-full cursor-pointer border border-gray-200 py-2 px-3.5 text-gray-500 outline-none placeholder:text-gray-400 focus-visible:border-gray-300"
      @click="showCalendar = !showCalendar"
      @blur="(e) => onBlurCalendar(e)"
    />
    <FormMonthCalendar
      v-show="showCalendar"
      :calendar-align="props.calendarAlign"
      :before-days="props.beforeDays"
      :set-max-date="props.setMaxDate"
      @close-calendar="showCalendar = false"
      @update:model-value="(newValue) => emit('update:modelValue', newValue)"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { format } from 'date-fns'
import FormMonthCalendar from '@/base/components/Form/Calendar/FormMonthCalendar.vue'

const props = defineProps({
  modelValue: {
    type: Date,
    default: () => new Date(),
  },
  calendarAlign: {
    type: String,
    default: 'left',
    validator: (value) => value === 'left' || value === 'right',
  },
})

const emit = defineEmits(['update:modelValue'])

const showCalendar = ref(false)

const scopeElRef = ref()
function onBlurCalendar(e) {
  if (e.relatedTarget !== null && scopeElRef.value.contains(e.relatedTarget)) {
    return
  }
  showCalendar.value = false
}
</script>
