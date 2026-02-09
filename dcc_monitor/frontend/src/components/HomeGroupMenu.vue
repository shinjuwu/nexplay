<script setup lang="ts">
const props = defineProps({
  groupName: {
    type: String,
    default: '',
  },
})
</script>

<template>
  <section class="rounded bg-white p-4">
    <header class="text-center">
      <h2 class="text-2xl font-bold">{{ props.groupName }}</h2>
    </header>

    <hr class="my-2" />

    <nav
      class="grid grid-cols-2 gap-2 sm:grid-cols-3 sm:gap-4 md:grid-cols-4 md:gap-6 lg:grid-cols-5 lg:gap-8 xl:grid-cols-6 xl:gap-10"
    >
      <slot></slot>
    </nav>
  </section>
</template>
