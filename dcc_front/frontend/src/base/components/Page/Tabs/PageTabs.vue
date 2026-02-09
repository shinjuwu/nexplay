<template>
  <div class="bg-white px-2 pt-4">
    <ul class="flex items-center border-b pl-2">
      <li
        v-for="tabItem in props.tabs"
        :key="`tab__${tabItem}`"
        :model-value="props.modelValue"
        class="text-primary mb-[-1px] cursor-pointer rounded-t px-10 py-2"
        :class="{ 'border border-b-white text-black': props.modelValue === tabItem }"
        @click="emit('update:modelValue', tabItem)"
      >
        {{ tabItem }}
      </li>
    </ul>
  </div>
</template>

<script setup>
const props = defineProps({
  modelValue: {
    type: String,
    required: true,
  },
  tabs: {
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits(['update:modelValue'])
</script>
