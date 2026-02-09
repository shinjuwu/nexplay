<template>
  <div ref="scopeElRef" class="relative">
    <input
      :class="[props.calendarAlign === 'left' ? 'rounded-l' : 'rounded-r', { 'cursor-pointer': !props.disabled }]"
      type="text"
      readonly
      :disabled="props.disabled"
      :value="formatDate"
      class="w-full border border-gray-200 py-2 px-3.5 text-gray-500 outline-none placeholder:text-gray-400 focus-visible:border-gray-300"
      @click="showCalendar = !showCalendar"
      @blur="(e) => onBlurCalendar(e)"
    />
    <FormDateTimeCalendar
      :visible="showCalendar"
      :model-value="props.modelValue || new Date()"
      :calendar-align="props.calendarAlign"
      :before-days="props.beforeDays"
      :set-max-date="props.setMaxDate"
      :minute-increment="props.minuteIncrement"
      :display-times="props.displayTimes"
      @close-calendar="showCalendar = false"
      @update:model-value="(newValue) => emit('update:modelValue', newValue)"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { format } from 'date-fns'
import FormDateTimeCalendar from '@/base/components/Form/Calendar/FormDateTimeCalendar.vue'

const props = defineProps({
  modelValue: {
    type: Date,
    default: new Date(),
  },
  calendarAlign: {
    type: String,
    default: 'left',
    validator: (value) => value === 'left' || value === 'right',
  },
  beforeDays: {
    type: Number,
    default: 0,
    validator: (value) => value >= 0,
  },
  setMaxDate: {
    type: Boolean,
    default: false,
  },
  minuteIncrement: {
    type: Number,
    default: 1,
    validator: (value) => value > 0,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  displayTimes: {
    type: Boolean,
    default: true,
  },
})

const emit = defineEmits(['update:modelValue'])

const showCalendar = ref(false)
const formatDate = computed(() => {
  if (!props.modelValue) {
    return ''
  }

  return props.displayTimes ? format(props.modelValue, 'yyyy-MM-dd HH:mm') : format(props.modelValue, 'yyyy-MM-dd')
})
const scopeElRef = ref()
function onBlurCalendar(e) {
  if (e.relatedTarget !== null && scopeElRef.value.contains(e.relatedTarget)) {
    return
  }
  showCalendar.value = false
}
</script>
