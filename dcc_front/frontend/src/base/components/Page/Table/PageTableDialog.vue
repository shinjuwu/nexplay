<template>
  <div v-show="props.visible" class="fixed inset-x-0 top-0 z-10">
    <div
      class="relative z-10 m-2 flex flex-col bg-white md:my-7 md:mx-auto"
      :class="{
        'md:max-w-[600px]': props.size === 'md',
        'md:max-w-[600px] lg:max-w-[800px]': props.size === 'lg',
        'md:max-w-[600px] lg:max-w-[800px] xl:max-w-[1200px]': props.size === 'xl',
      }"
    >
      <div class="bg-indigo-400 p-4 text-xl font-bold text-white">
        <slot name="header"></slot>
      </div>
      <div class="p-3">
        <slot></slot>
      </div>
      <hr />
      <div v-if="props.hasFooter" class="flex items-center justify-between p-3">
        <slot name="footer"></slot>
      </div>
    </div>
    <div class="fixed top-0 h-screen w-full bg-black/50" @click="emit('close')"></div>
  </div>
</template>

<script setup>
const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  size: {
    type: String,
    default: 'lg',
    validator: (value) => ['sm', 'md', 'lg', 'xl', '2xl'].indexOf(value) >= 0,
  },
  hasFooter: {
    type: Boolean,
    default: true,
  },
})

const emit = defineEmits(['close'])
</script>
