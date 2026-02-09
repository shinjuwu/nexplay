<template>
  <i18n-t
    v-show="props.totalPages > 0"
    keypath="fmtTextTableSkipPage"
    tag="div"
    scope="global"
    class="flex items-center whitespace-nowrap"
  >
    <template #input>
      <span class="relative mx-1 w-full">
        <input
          v-model="skipPageInput"
          class="w-full appearance-none rounded border p-2 pr-4 outline-none"
          type="number"
          min="1"
          :max="props.totalPages"
        />
        <button
          class="absolute right-1 top-2 border-[5px] border-transparent border-b-slate-500 disabled:border-b-slate-400"
          :disabled="skipPageInput >= props.totalPages"
          @click="skipPageInput = skipPageInput + 1"
        ></button>
        <button
          class="absolute right-1 bottom-2 border-[5px] border-transparent border-t-slate-500 disabled:border-t-slate-400"
          :disabled="skipPageInput <= 1"
          @click="skipPageInput = skipPageInput - 1"
        ></button>
      </span>
    </template>
  </i18n-t>
</template>

<script setup>
import { computed } from 'vue'
import * as math from '@/base/utils/math'

const skipPageInput = computed({
  get() {
    return math.round(props.recordStart / props.pageLength, 0) + 1
  },
  set(value) {
    if (!value || value < 1) {
      value = 1
    }
    if (value > props.totalPages) {
      value = props.totalPages
    }
    emit('pageChange', props.pageLength * (value - 1))
  },
})

const props = defineProps({
  recordStart: {
    type: Number,
    default: 0,
  },
  pageLength: {
    type: Number,
    default: 0,
  },
  totalPages: {
    type: Number,
    default: 0,
  },
})
const emit = defineEmits(['pageChange'])
</script>
