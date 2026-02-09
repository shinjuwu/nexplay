<template>
  <slot
    :current-records="currentRecords"
    :page-length="props.tableInput.length"
    :record-start="props.tableInput.start"
    :total-pages="totalPages"
    :total-records="props.totalRecords"
    :is-sort-icon-active="isSortIconActive"
    :length-change="lengthChange"
    :page-change="pageChange"
    :sorting="sorting"
  >
  </slot>
  <PageTableLoading v-show="props.tableInput.showProcessing" />
</template>

<script setup>
import { computed } from 'vue'
import { BaseTableInput } from '@/base/common/table/tableInput'
import * as math from '@/base/utils/math'
import * as table from '@/base/utils/table'

import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'

const currentRecords = computed(() => {
  const column = props.tableInput.column
  const dir = props.tableInput.dir
  const start = props.tableInput.start
  const end = start + props.tableInput.length

  return props.records
    .slice()
    .sort((a, b) => table.columnSorting(dir, a[column], b[column]))
    .slice(start, end)
})

const totalPages = computed(() => {
  return props.totalRecords > 0 ? Math.ceil(math.round(props.totalRecords / props.tableInput.length, 3)) : 0
})

function pageChange(start) {
  props.tableInput.pageChange(start)
  emit('pageChange', props.tableInput)
}

function lengthChange(length) {
  props.tableInput.lengthChange(length)
  emit('lengthChange', props.tableInput)
}

function sorting(column) {
  props.tableInput.sorting(column)
  emit('sorting', props.tableInput)
}

function isSortIconActive(column, dir) {
  return column === props.tableInput.column && dir === props.tableInput.dir
}

const props = defineProps({
  tableInput: {
    type: BaseTableInput,
    default: () => new BaseTableInput(),
  },
  records: {
    type: Array,
    default: () => [],
  },
  totalRecords: {
    type: Number,
    default: 0,
  },
})
const emit = defineEmits(['lengthChange', 'pageChange', 'sorting'])
</script>
