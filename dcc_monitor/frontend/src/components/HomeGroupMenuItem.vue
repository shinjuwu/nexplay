<script setup lang="ts">
const props = defineProps({
  name: {
    type: String,
    default: '',
  },
  linkTo: {
    type: String,
    default: '/',
  },
  iconContainerClass: {
    type: String,
    default: '',
  },
})
</script>

<template>
  <router-link v-slot="{ navigate }" custom :to="props.linkTo">
    <div class="cursor-pointer" @click="navigate">
      <div class="relative rounded border-4">
        <div class="mt-[100%]"></div>
        <div class="absolute inset-0 flex items-center justify-center" :class="iconContainerClass">
          <slot icon-size-class="h-3/5"></slot>
        </div>
      </div>
      <div class="text-center text-xl">{{ props.name }}</div>
    </div>
  </router-link>
</template>
