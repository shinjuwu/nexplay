<template>
  <footer
    class="flex border-t border-gray-400/20 bg-gray-100 py-2 pr-4 text-gray-400 md:ml-[4.6rem]"
    :class="{ 'md:ml-[250px]': props.isSidebarOpen }"
  >
    <ul class="flex">
      <li>
        <a class="block px-4 py-2"></a>
      </li>
    </ul>
    <ul class="ml-auto flex">
      <li>
        <a class="block px-4 py-2">{{ appVersion }}</a>
      </li>
    </ul>
  </footer>
</template>

<script setup>
import { useFooter } from '@/base/composable/layout/useFooter'

const props = defineProps({
  isSidebarOpen: {
    type: Boolean,
    default: true,
  },
})

const { appVersion, year } = useFooter()
</script>
