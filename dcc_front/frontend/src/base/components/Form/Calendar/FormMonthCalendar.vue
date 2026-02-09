<template>
  <div v-show="props.visible" class="z-[3] w-[240px] bg-white">
    <div class="flex flex-wrap justify-around px-2">
      <button
        v-for="month in monthRange"
        :key="month"
        type="button"
        :class="{
          'bg-indigo-400 text-white hover:bg-indigo-500': isSameMonth(month, modelValue),
        }"
        :disabled="
          (props.setMaxDate && isAfter(month, currentDays)) ||
          (props.beforeDays > 0 && isBefore(month, startOfMonth(props.minDay)))
        "
        class="h-10 w-12 rounded hover:bg-indigo-400 hover:text-white disabled:text-gray-400 disabled:hover:bg-gray-300"
        @click="selectedMonth = month"
      >
        {{ t(`calendarMonth__${month.getMonth()}`) }}
      </button>
    </div>
    <hr class="my-1" />
    <div class="pb-1 text-right">
      <button type="button" class="btn btn-primary px-2 py-1 text-sm" @click="emit('close')">
        {{ t('textClose') }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { startOfMonth, eachMonthOfInterval, startOfYear, endOfYear, isBefore, isAfter, isSameMonth } from 'date-fns'
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
const emit = defineEmits(['update:modelValue', 'close'])
const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  minDay: {
    type: Date,
    default: new Date(),
  },
  beforeDays: {
    type: Number,
    default: 0,
    validator: (value) => value >= 0,
  },
  selectedDay: {
    type: Date,
    default: new Date(),
  },
  setMaxDate: {
    type: Boolean,
    default: false,
  },
  modelValue: {
    type: Date,
    default: new Date(),
  },
})
const { t } = useI18n()
const currentDays = new Date()
const currentMonth = ref(startOfMonth(currentDays))
const monthRange = computed(() => {
  return eachMonthOfInterval({
    start: startOfYear(props.selectedDay),
    end: endOfYear(props.selectedDay),
  })
})

const selectedMonth = computed({
  get() {
    return currentMonth
  },
  set(value) {
    currentMonth.value = value
    emit('update:modelValue', value)
  },
})
</script>
