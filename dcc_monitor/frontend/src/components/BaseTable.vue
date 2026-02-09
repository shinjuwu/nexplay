<script setup lang="ts" generic="T">
import type { TableProps } from '@/types/types.table'
import { computed, ref, watch } from 'vue'
import { round } from '@/utils/math'

const props = withDefaults(defineProps<TableProps<T>>(), {
  items: () => [] as T[],
  lengthMenu: () => [10, 25, 50, 100],
  pageLength: 10,
  infoContainer: '',
  fastPagingContainer: '',
  lengthMenuContainer: '',
  paginationContainer: '',
  processingContainer: '',
  searchButtonContainer: '',
})
const emit = defineEmits(['search'])

const currentItems = computed(() => props.items.slice(start.value, start.value + length.value))

const draw = ref(0)
const start = ref(0)
watch(
  () => props.items.length,
  (newValue) => {
    if (newValue <= 0) {
      return
    }

    while (start.value >= newValue) {
      start.value -= length.value
    }
  }
)

const length = ref(props.pageLength)
watch(
  () => props.pageLength,
  (newLength) => {
    length.value = newLength
  }
)
watch(length, (newLength) => {
  start.value = start.value - (start.value % newLength)
  draw.value = draw.value + 1
})

const infoStart = computed(() => (props.items.length > 0 ? start.value + 1 : 0))
const infoEnd = computed(() => (props.items.length > 0 ? start.value + currentItems.value.length : 0))
const infoTotal = computed(() => props.items.length)

const page = computed({
  get() {
    return round(start.value / length.value, 0)
  },
  set(newPage) {
    start.value = newPage * length.value
    draw.value = draw.value + 1
  },
})
const totalPages = computed(() => (props.items.length > 0 ? Math.ceil(round(props.items.length / length.value, 3)) : 0))
const paginateArray = computed(() => {
  const paginateArray: number[] = Array.from({ length: totalPages.value }, (_, i) => i)
  if (totalPages.value > 8) {
    const currentPage = page.value
    if (currentPage < 4) {
      paginateArray.splice(5, totalPages.value - 1 - 5, -1)
    } else if (totalPages.value - 1 - currentPage < 4) {
      paginateArray.splice(1, totalPages.value - 5 - 1, -1)
    } else {
      paginateArray.splice(currentPage + 2, totalPages.value - 1 - (currentPage + 2), -1)
      paginateArray.splice(1, currentPage - 2, -1)
    }
  }
  return paginateArray
})

const skipPageInput = computed({
  get() {
    return round(start.value / length.value, 0) + 1
  },
  set(value) {
    if (!value || value < 1) {
      value = 1
    }
    if (value > totalPages.value) {
      value = totalPages.value
    }
    page.value = value - 1
  },
})

const isProcessing = ref(false)
function showProcessing() {
  isProcessing.value = true
}
function closeProcessing() {
  isProcessing.value = false
}

function resetTable() {
  start.value = 0
  draw.value = 0
}
</script>

<template>
  <slot :current-items="currentItems"></slot>

  <Teleport v-if="props.searchButtonContainer" :to="props.searchButtonContainer">
    <button
      type="button"
      class="btn btn-primary w-full"
      @click="emit('search', { start, length, showProcessing, closeProcessing, resetTable })"
    >
      搜寻
    </button>
  </Teleport>

  <Teleport v-if="props.infoContainer" :to="props.infoContainer">
    <div class="hidden sm:block">显示第 {{ infoStart }} 到第 {{ infoEnd }} 条纪录，总共 {{ infoTotal }} 条纪录</div>
    <div class="sm:hidden">第 {{ infoStart }} ~ {{ infoEnd }} 条纪录，共 {{ infoTotal }} 条纪录</div>
  </Teleport>

  <Teleport v-if="props.lengthMenuContainer" :to="props.lengthMenuContainer">
    显示
    <select v-model="length" class="mx-1 w-full rounded border py-2 pl-1 pr-3.5">
      <option v-for="len in props.lengthMenu" :key="`lengthMenu__${len}`" :value="len">{{ len }}条/页</option>
    </select>
    项结果
  </Teleport>

  <Teleport v-if="props.paginationContainer" :to="props.paginationContainer">
    <button class="mr-1 w-[1.5rem]" :disabled="totalPages <= 1 || page < 1" @click="page = page - 1">
      <font-awesome-icon icon="fa-solid fa-chevron-left" :class="{ 'text-slate-400': totalPages <= 1 || page < 1 }" />
    </button>
    <button
      v-for="(paginate, idx) in paginateArray"
      :key="`paginateArray__${idx}__${paginate}`"
      class="mr-1 min-w-[1.5rem] rounded-full px-1"
      :class="page === paginate ? 'bg-primary text-white' : 'hover:bg-slate-300'"
      :disabled="paginate === -1 || page === paginate"
      @click="page = paginate"
    >
      {{ paginate === -1 ? '...' : paginate + 1 }}
    </button>
    <button class="mr-1 w-[1.5rem]" :disabled="totalPages <= 1 || page >= totalPages - 1" @click="page = page + 1">
      <font-awesome-icon
        icon="fa-solid fa-chevron-right"
        :class="{ 'text-slate-400': totalPages <= 1 || page >= totalPages - 1 }"
      />
    </button>
  </Teleport>

  <Teleport v-if="props.fastPagingContainer && totalPages > 0" :to="props.fastPagingContainer">
    跳转到第
    <span class="relative mx-1 w-full">
      <input
        v-model="skipPageInput"
        class="w-full appearance-none rounded border p-2 pr-4 outline-none"
        type="number"
        min="1"
        :max="totalPages"
      />
      <button
        class="absolute right-1 top-2 border-[5px] border-transparent border-b-slate-500 disabled:border-b-slate-400"
        :disabled="skipPageInput >= totalPages"
        @click="skipPageInput = skipPageInput + 1"
      ></button>
      <button
        class="absolute bottom-2 right-1 border-[5px] border-transparent border-t-slate-500 disabled:border-t-slate-400"
        :disabled="skipPageInput <= 1"
        @click="skipPageInput = skipPageInput - 1"
      ></button>
    </span>
    页
  </Teleport>

  <Teleport v-if="props.processingContainer" :to="props.processingContainer">
    <div v-show="isProcessing" class="absolute inset-0 z-10">
      <div class="absolute inset-0 rounded bg-white opacity-75"></div>
      <div class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
        <div class="h-12 w-12 animate-spin rounded-full border-4 border-black border-b-transparent"></div>
      </div>
    </div>
  </Teleport>
</template>
