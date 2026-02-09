<template>
  <th class="sorting" @click="emit('sorting', props.column)">
    <slot></slot>
    <div class="asc" :class="{ active: isSortIconActive(props.column, constant.TableSortDirection.Asc) }"></div>
    <div class="desc" :class="{ active: isSortIconActive(props.column, constant.TableSortDirection.Desc) }"></div>
  </th>
</template>

<script setup>
import constant from '@/base/common/constant'

const props = defineProps({
  column: {
    type: String,
    default: '',
  },
  isSortIconActive: {
    type: Function,
    default: () => false,
  },
})
const emit = defineEmits(['sorting'])
</script>
