<template>
  <div v-show="props.visible" class="z-[2] w-[240px] bg-white">
    <div ref="scrollCalendar" class="h-[200px] overflow-y-scroll text-center">
      <button
        v-for="year in yearRange"
        :key="`yearKey__${year}`"
        type="button"
        :class="{
          'bg-indigo-400 text-white hover:bg-indigo-500 current-year': isSameYear(year, props.modelValue),
        }"
        :disabled="!isSameYear(year, minDay) && props.setMaxDate"
        class="h-10 w-12 rounded hover:bg-indigo-400 hover:text-white disabled:text-gray-400 disabled:hover:bg-gray-300"
        @click="selectedYear = year"
      >
        {{ year.getFullYear() }}
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
import { eachYearOfInterval, getYear, isSameYear } from 'date-fns'
import { ref, computed, onUpdated } from 'vue'
import { useI18n } from 'vue-i18n'
const emit = defineEmits(['update:modelValue', 'close', 'getMaxAndMinDate'])
const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  minDay: {
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

const currentYear = ref(new Date())
const scrollCalendar = ref(null)
const yearRange = computed(() => {
  return eachYearOfInterval({
    start: new Date(2000, 0, 1),
    end: new Date(2099, 0, 1),
  })
})

const selectedYear = computed({
  get() {
    return currentYear
  },
  set(value) {
    currentYear.value = value
    emit('update:modelValue', value)
  },
})

onUpdated(() => {
  if (props.visible) {
    const yearsPerRow = 4
    const rowsHeight = 40
    const yearDifference = getYear(props.modelValue) - getYear(yearRange.value[0])
    const currentScroll = Math.floor(yearDifference / yearsPerRow)

    scrollCalendar.value.scrollTop = rowsHeight * currentScroll
  }
})
</script>
