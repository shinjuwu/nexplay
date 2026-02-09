<template>
  <PageTable
    :table-input="props.tableInput"
    :total-records="props.tableInput.totalRecords"
    @length-change="(tableInput) => emit('search', tableInput)"
    @page-change="(tableInput) => emit('search', tableInput)"
    @sorting="(tableInput) => emit('search', tableInput)"
  >
    <template
      #default="{
        pageLength,
        recordStart,
        totalPages,
        totalRecords,
        isSortIconActive,
        lengthChange,
        pageChange,
        sorting,
      }"
    >
      <slot
        :current-records="props.records"
        :page-length="pageLength"
        :record-start="recordStart"
        :total-pages="totalPages"
        :total-records="totalRecords"
        :is-sort-icon-active="isSortIconActive"
        :length-change="lengthChange"
        :page-change="pageChange"
        :sorting="sorting"
      ></slot>
    </template>
  </PageTable>
</template>

<script setup>
import { BaseServersideTableInput } from '@/base/common/table/tableInput'

import PageTable from '@/base/components/Page/Table/PageTable.vue'

const props = defineProps({
  tableInput: {
    type: BaseServersideTableInput,
    default: () => new BaseServersideTableInput(),
  },
  records: {
    type: Array,
    default: () => [],
  },
})
const emit = defineEmits(['search'])
</script>
