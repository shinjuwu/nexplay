<template>
  <div class="ml-auto flex items-center justify-center whitespace-nowrap py-2">
    <button class="mr-1" :disabled="totalPages <= 1 || page < 1" @click="page = page - 1">
      <ChevronLeftIcon class="h-5 w-5" :class="{ 'text-slate-400': totalPages <= 1 || page < 1 }" />
    </button>
    <button
      v-for="(paginate, idx) in paginateArray"
      :key="`paginateArray__${idx}__${paginate}`"
      class="mr-1 min-w-[1.5rem] rounded-full px-1"
      :class="page === paginate ? 'bg-indigo-400 text-white' : 'hover:bg-slate-300'"
      :disabled="paginate === '...' || page === paginate"
      @click="page = paginate"
    >
      {{ paginate === '...' ? paginate : paginate + 1 }}
    </button>
    <button class="mr-1" :disabled="totalPages <= 1 || page >= totalPages - 1" @click="page = page + 1">
      <ChevronRightIcon class="h-5 w-5" :class="{ 'text-slate-400': totalPages <= 1 || page >= totalPages - 1 }" />
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ChevronRightIcon, ChevronLeftIcon } from '@heroicons/vue/24/outline'
import * as math from '@/base/utils/math'

const page = computed({
  get() {
    return math.round(props.recordStart / props.pageLength, 0)
  },
  set(value) {
    emit('pageChange', props.pageLength * value)
  },
})
const paginateArray = computed(() => {
  let paginateArray = Array.apply(null, { length: props.totalPages }).map(Number.call, Number)
  if (props.totalPages > 8) {
    const currentPage = page.value
    if (currentPage < 4) {
      paginateArray.splice(5, props.totalPages - 1 - 5, '...')
    } else if (props.totalPages - 1 - currentPage < 4) {
      paginateArray.splice(1, props.totalPages - 5 - 1, '...')
    } else {
      paginateArray.splice(currentPage + 2, props.totalPages - 1 - (currentPage + 2), '...')
      paginateArray.splice(1, currentPage - 2, '...')
    }
  }
  return paginateArray
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
