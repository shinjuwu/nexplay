<template>
  <div class="rounded bg-white" :class="{ 'p-4': !hasTips }">
    <div :class="{ 'flex justify-between': hasTips }">
      <div v-if="hasTips" class="text-danger">{{ props.tipsTitle }}</div>

      <div class="text-right">
        <span class="inline-block cursor-pointer" @click="showSlot = !showSlot">
          <MinusCircleIcon v-if="showSlot" class="h-5 w-5" />
          <PlusCircleIcon v-else class="h-5 w-5" />
        </span>
      </div>
    </div>
    <slot :show="showSlot"> </slot>
    <slot name="tipsSlot" :tips-show="showSlot"></slot>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { MinusCircleIcon, PlusCircleIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  tipsTitle: {
    type: String,
    default: '',
  },
})

const hasTips = computed(() => props.tipsTitle.length)
const showSlot = ref(true)
</script>
