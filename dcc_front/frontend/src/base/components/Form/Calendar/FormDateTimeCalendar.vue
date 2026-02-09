<template>
  <div
    v-show="props.visible"
    ref="scopeElRef"
    class="absolute top-[50px] z-[1] inline-block max-h-[320px] w-[250px] rounded border border-gray-200 bg-white"
    :class="props.calendarAlign === 'left' ? 'left-0' : 'right-0'"
    tabindex="-1"
    @blur="(e) => onBlurCalendar(e)"
  >
    <div
      class="absolute -top-4 inline-block border-8 border-gray-200 border-x-transparent border-t-transparent"
      :class="props.calendarAlign === 'left' ? 'left-4' : 'right-4'"
    ></div>
    <div class="flex items-center px-4 pt-2.5">
      <button
        type="button"
        class="border-8 border-transparent border-r-gray-400 disabled:border-r-gray-300"
        :disabled="
          (props.beforeDays > 0 && isSameMonth(firstDayCurrentMonth, minDay)) ||
          isBefore(firstDayCurrentMonth, new Date(2000, 0, 1))
        "
        @click="firstDayCurrentMonth = add(firstDayCurrentMonth, { months: -1 })"
        @blur="(e) => onBlurCalendar(e)"
      ></button>
      <div class="flex flex-1 justify-center">
        <span class="cursor-pointer p-1 text-lg hover:rounded-md hover:bg-gray-200" @click="changeCalendar('month')">
          {{ t(`calendarMonth__${firstDayCurrentMonth.getMonth()}`) }}
        </span>
        &nbsp;
        <span class="cursor-pointer p-1 text-lg hover:rounded-md hover:bg-gray-200" @click="changeCalendar('year')">
          {{ firstDayCurrentMonth.getFullYear() }}
        </span>
      </div>
      <button
        type="button"
        class="border-8 border-transparent border-l-gray-400 disabled:border-l-gray-300"
        :disabled="
          (props.setMaxDate && isSameMonth(firstDayCurrentMonth, today)) ||
          isBefore(new Date(2099, 11, 31), endOfMonth(firstDayCurrentMonth))
        "
        @click="firstDayCurrentMonth = add(firstDayCurrentMonth, { months: 1 })"
        @blur="(e) => onBlurCalendar(e)"
      ></button>
    </div>
    <FormYearCalendar
      :visible="showYearCalender"
      :before-days="beforeDays"
      :set-max-date="props.setMaxDate"
      :model-value="props.modelValue"
      @update:model-value="(newValue) => (selectedYear = newValue)"
      @close="showYearCalender = false"
    />
    <FormMonthCalendar
      :visible="showMonthCalendar"
      :selected-day="selectedDay"
      :set-max-date="props.setMaxDate"
      :model-value="props.modelValue"
      :min-day="minDay"
      :before-days="beforeDays"
      @update:model-value="(newValue) => (selectedMonth = newValue)"
      @close="showMonthCalendar = false"
    />
    <div :class="{ hidden: showYearCalender || showMonthCalendar }" class="grid grid-cols-7 p-1.5 text-center">
      <div v-for="i in 7" :key="`calendarWeekOfDay1__${i}`" class="text-gray-500">
        {{ t(`calendarWeekOfDay1__${i - 1}`) }}
      </div>
      <div v-for="day in currentDays" :key="day">
        <button
          type="button"
          class="h-7 w-7 rounded hover:bg-indigo-400 hover:text-white disabled:text-gray-400 disabled:hover:bg-gray-300"
          :class="{
            'bg-indigo-400 text-white hover:bg-indigo-500': isSameDay(day, selectedDay),
            'text-gray-400 hover:bg-gray-300': !isSameMonth(day, firstDayCurrentMonth) && !isSameDay(day, selectedDay),
          }"
          :disabled="
            (props.setMaxDate && isAfter(day, today)) ||
            (props.beforeDays > 0 && isBefore(day, minDay)) ||
            isBefore(day, new Date(2000, 0, 1)) ||
            isBefore(new Date(2099, 11, 31), day)
          "
          @click="selectedDay = day"
          @blur="(e) => onBlurCalendar(e)"
        >
          {{ format(day, 'd') }}
        </button>
      </div>
    </div>
    <div
      :class="{ hidden: showYearCalender || showMonthCalendar }"
      class="flex items-center border-t-[1px] border-gray-200 p-2"
    >
      <template v-if="props.displayTimes">
        <ClockIcon class="h-4 w-4 text-gray-500" />
        <div class="ml-2 pl-1 pb-1 text-sm">
          <div class="pl-1 pb-1 leading-5">
            <span>{{ t(`calendarWeekOfDay2__${getDay(selectedDay)}`) }}</span>
            <span class="text-primary ml-2">
              {{ t('calendarDate', [getMonth(selectedDay) + 1, getDate(selectedDay)]) }}
            </span>
            <span class="ml-2 text-gray-400">{{ getYear(selectedDay) }}</span>
          </div>
          <div class="flex items-center">
            <div class="relative inline-block">
              <select
                v-model="selectedHour"
                class="h-[30px] w-[52px] cursor-pointer appearance-none rounded bg-gray-200 pl-2 pr-5 text-center text-sm outline-none"
                @blur="(e) => onBlurCalendar(e)"
              >
                <option v-for="hour in hours" :key="`hour__${hour}`" :value="hour">{{ `0${hour}`.slice(-2) }}</option>
              </select>
              <div class="pointer-events-none absolute top-[12px] right-[7px]">
                <div class="border-[6px] border-transparent border-t-gray-400 disabled:border-r-gray-300"></div>
              </div>
            </div>
            <span class="mx-1">:</span>
            <div class="relative mr-1 inline-block">
              <select
                v-model="selectedMinute"
                class="h-[30px] w-[52px] cursor-pointer appearance-none rounded bg-gray-200 pl-2 pr-5 text-center text-sm outline-none"
                @blur="(e) => onBlurCalendar(e)"
              >
                <option v-for="minute in minutes" :key="`minute__${minute}`" :value="minute">
                  {{ `0${minute}`.slice(-2) }}
                </option>
              </select>
              <div class="pointer-events-none absolute top-[12px] right-[7px]">
                <div class="border-[6px] border-transparent border-t-gray-400 disabled:border-r-gray-300"></div>
              </div>
            </div>
          </div>
        </div>
      </template>
      <div class="self-end pb-1" :class="{ 'ml-auto': !props.displayTimes }">
        <button
          type="button"
          class="btn btn-primary mr-1 ml-auto px-2 py-1 text-sm"
          @click="emit('closeCalendar')"
          @blur="(e) => onBlurCalendar(e)"
        >
          {{ t('textClose') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  add,
  set,
  eachDayOfInterval,
  endOfMonth,
  endOfWeek,
  format,
  getDate,
  getDay,
  getHours,
  getMinutes,
  getMonth,
  getYear,
  isAfter,
  isBefore,
  isSameDay,
  isSameMonth,
  startOfDay,
  startOfMonth,
  startOfToday,
  startOfWeek,
} from 'date-fns'
import { ClockIcon } from '@heroicons/vue/24/outline'
import FormYearCalendar from './FormYearCalendar.vue'
import FormMonthCalendar from './FormMonthCalendar.vue'

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
  visible: {
    type: Boolean,
    default: false,
  },
  displayTimes: {
    type: Boolean,
    default: true,
  },
})
const emit = defineEmits(['update:modelValue', 'closeCalendar'])

const { t } = useI18n()

const today = ref(startOfToday())

const showYearCalender = ref(false)
const showMonthCalendar = ref(false)

watch(
  () => props.visible,
  (newValue) => {
    if (newValue) {
      today.value = startOfToday()
    }
  }
)
const minDay = computed(() => add(today.value, { days: -props.beforeDays }))
const hours = [...Array(24).keys()]
const minutes = [...Array(60).keys()].filter((i) => i % props.minuteIncrement === 0)

const selectedYear = computed({
  get() {
    return getYear(today.value)
  },
  set(value) {
    const dateValue = set(selectedDay.value, {
      year: value.getFullYear(),
      hours: selectedHour.value,
      minutes: selectedMinute.value,
    })
    if (value !== getYear(selectedDay.value)) {
      firstDayCurrentMonth.value = startOfMonth(dateValue)
    }

    showYearCalender.value = false
    showMonthCalendar.value = true
    emit('update:modelValue', dateValue)
  },
})

const selectedMonth = computed({
  get() {
    return getMonth(today.value)
  },
  set(value) {
    const dateValue = set(selectedDay.value, {
      month: getMonth(value),
      hours: selectedHour.value,
      minutes: selectedMinute.value,
    })
    if (value !== getMonth(selectedDay.value)) {
      firstDayCurrentMonth.value = startOfMonth(dateValue)
    }

    showMonthCalendar.value = false
    emit('update:modelValue', dateValue)
  },
})

const selectedDay = computed({
  get() {
    return startOfDay(props.modelValue)
  },
  set(value) {
    if (!isSameMonth(value, firstDayCurrentMonth.value)) {
      firstDayCurrentMonth.value = startOfMonth(value)
    }

    const dateValue = props.displayTimes
      ? add(value, { hours: selectedHour.value, minutes: selectedMinute.value })
      : value

    emit('update:modelValue', dateValue)
  },
})
const selectedHour = computed({
  get() {
    return getHours(props.modelValue)
  },
  set(value) {
    emit('update:modelValue', add(selectedDay.value, { hours: value, minutes: selectedMinute.value }))
  },
})
const selectedMinute = computed({
  get() {
    return getMinutes(props.modelValue)
  },
  set(value) {
    emit('update:modelValue', add(selectedDay.value, { hours: selectedHour.value, minutes: value }))
  },
})

const firstDayCurrentMonth = ref(startOfMonth(props.modelValue))
const currentDays = computed(() => {
  return eachDayOfInterval({
    start: startOfWeek(firstDayCurrentMonth.value),
    end: endOfWeek(endOfMonth(firstDayCurrentMonth.value)),
  })
})

const scopeElRef = ref()
function onBlurCalendar(e) {
  if (e.relatedTarget !== null && scopeElRef.value.contains(e.relatedTarget)) {
    return
  }
  if (e.target.disabled) {
    scopeElRef.value.focus()
    return
  }
  showMonthCalendar.value = false
  showYearCalender.value = false
  emit('closeCalendar')
}

function changeCalendar(type) {
  if (type === 'month') {
    showMonthCalendar.value = !showMonthCalendar.value
    showYearCalender.value = false
    return
  }

  if (type === 'year') {
    showYearCalender.value = !showYearCalender.value
    showMonthCalendar.value = false
    return
  }
}
</script>
