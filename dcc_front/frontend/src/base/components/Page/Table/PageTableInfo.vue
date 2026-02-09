<template>
  <i18n-t keypath="fmtTextTableInfo" tag="div" scope="global" class="whitespace-nowrap py-2">
    <template #infoStart>
      <strong class="px-1">{{ infoStart }}</strong>
    </template>
    <template #infoEnd>
      <strong class="px-1">{{ infoEnd }}</strong>
    </template>
    <template #infoTotal>
      <strong class="px-1">{{ infoTotal }}</strong>
    </template>
  </i18n-t>
</template>

<script setup>
import { computed } from 'vue'

const infoStart = computed(() => {
  return props.totalRecords > 0 ? props.recordStart + 1 : 0
})
const infoEnd = computed(() => {
  return props.totalRecords > 0 ? props.recordStart + props.displayRecords : 0
})
const infoTotal = computed(() => {
  return props.totalRecords
})

const props = defineProps({
  recordStart: {
    type: Number,
    default: 0,
  },
  displayRecords: {
    type: Number,
    default: 0,
  },
  totalRecords: {
    type: Number,
    default: 0,
  },
})
</script>
