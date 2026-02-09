<template>
  <BaseHeader
    ref="headerEl"
    :is-sidebar-open="isSidebarOpen"
    @toggle-sidebar="isSidebarOpen = !isSidebarOpen"
    @show-personal-info="isShowPersonalInfo = true"
  />
  <BaseSidebar
    :menu-folders="props.menuFolders"
    :menu-items="props.menuItems"
    :menu-item-map="props.menuItemMap"
    :is-sidebar-open="isSidebarOpen"
  />
  <main ref="mainEl" class="mt-14 bg-gray-200 pb-2 md:ml-[4.6rem]" :class="{ 'md:ml-[250px]': isSidebarOpen }">
    <slot> </slot>
  </main>
  <BaseFooter ref="footerEl" :is-sidebar-open="isSidebarOpen" />
  <PersonalInfo :visible="isShowPersonalInfo" @close="isShowPersonalInfo = false" />
  <div
    class="fixed inset-x-0 top-14 bottom-0 z-[8] bg-black/10 md:hidden"
    :class="{ hidden: !isSidebarOpen }"
    @click="isSidebarOpen = false"
  ></div>
</template>

<script setup>
import { useLayout } from '@/base/composable/layout/useLayout'
import BaseHeader from '@/base/views/Layout/BaseHeader.vue'
import BaseSidebar from '@/base/views/Layout/BaseSidebar.vue'
import BaseFooter from '@/base/views/Layout/BaseFooter.vue'
import PersonalInfo from '@/base/views/Layout/PersonalInfo.vue'

const { isSidebarOpen, isShowPersonalInfo, headerEl, mainEl, footerEl } = useLayout()

const props = defineProps({
  menuFolders: {
    type: Array,
    default: () => [],
  },
  menuItems: {
    type: Array,
    default: () => [],
  },
  menuItemMap: {
    type: Object,
    default: () => {},
  },
})
</script>
