<template>
  <div class="h-full rounded border bg-gray-100 p-3 drop-shadow-md">
    <div class="grid h-full grid-flow-col grid-cols-10">
      <div class="col-span-3 flex">
        <slot name="icon"> </slot>
        <span class="ml-2">{{ props.title }}</span>
      </div>
      <div class="col-span-4 flex items-center justify-center">
        <span class="text-xl font-medium">{{ props.data[0] }}</span>
      </div>
      <div class="text-danger col-span-3 break-words text-right">
        <div>{{ compareText }}</div>
        <div class="text-lg">{{ props.data[1] }}</div>
      </div>
    </div>
  </div>
</template>
<script setup>
const props = defineProps({
  icon: {
    type: String,
    default: '',
  },
  title: {
    type: String,
    default: '',
  },
  data: {
    type: Array,
    default: () => [0, 0],
  },
  compareText: {
    type: String,
    default: '',
  },
})
</script>
