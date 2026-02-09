<template>
  <aside
    class="fixed inset-y-0 z-[9] w-[250px] pt-14 md:ml-0"
    :class="{ 'md:w-[4.6rem] ml-[-250px]': !props.isSidebarOpen }"
  >
    <div
      class="fixed top-0 hidden h-[56px] bg-gray-700 px-2 py-[0.8125rem] md:flex md:w-[4.6rem] md:items-center md:justify-center"
      :class="{ 'md:w-[250px]': props.isSidebarOpen }"
    >
      <!-- logo removed -->
    </div>
    <nav class="h-full bg-gray-700">
      <ul class="text-[18px] text-gray-400">
        <li v-for="(folder, index) in menu.items" :key="`menuFolder${folder.key}`" class="group relative">
          <template v-if="props.isSidebarOpen">
            <a
              class="flex h-[43px] cursor-pointer items-center px-4 py-2"
              :class="{
                'text-gray-300': folder.show,
              }"
              @click="toggleMenu(index)"
            >
              <component :is="folder.icon" class="mx-2 w-[1.6rem]" />
              <span>{{ t(`menuFolder${folder.key}`) }}</span>
              <ChevronRightIcon
                class="absolute right-2 w-5"
                :class="{ 'md:hidden': !props.isSidebarOpen, 'rotate-90': folder.show }"
              />
            </a>
            <SlideTransition>
              <ul v-show="folder.show" class="bg-gray-800">
                <li v-for="item in folder.items" :key="`menuItem${item.key}`">
                  <a
                    class="flex cursor-pointer items-center py-2 px-4 hover:text-gray-300"
                    @click="refreshOrShowBreadcrumbItem(item)"
                  >
                    <StopIcon class="mx-2 w-[1.6rem] rotate-45" />
                    <span>{{ t(`menuItem${item.key}`) }}</span>
                  </a>
                </li>
              </ul>
            </SlideTransition>
          </template>
          <template v-else>
            <a
              class="flex h-[43px] cursor-pointer items-center px-4 py-2 md:w-[4.6rem] md:group-hover:w-[250px] md:group-hover:bg-indigo-400 md:group-hover:text-white"
            >
              <component :is="folder.icon" class="mx-2 w-[1.6rem] md:group-hover:mr-6" />
              <span :class="{ 'md:hidden md:group-hover:inline-block md:group-hover:ml-2': !props.isSidebarOpen }">{{
                t(`menuFolder${folder.key}`)
              }}</span>
            </a>
            <ul
              class="hidden bg-gray-800 md:group-hover:absolute md:group-hover:left-[4.6rem] md:group-hover:top-[43px] md:group-hover:block md:group-hover:w-[calc(250px-4.6rem)]"
            >
              <li v-for="item in folder.items" :key="`menuItem${item.key}`">
                <a
                  class="ml-2 flex cursor-pointer items-center py-2 hover:text-gray-300"
                  @click="refreshOrShowBreadcrumbItem(item)"
                >
                  <span>{{ t(`menuItem${item.key}`) }}</span>
                </a>
              </li>
            </ul>
          </template>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useSidebar } from '@/base/composable/layout/useSidebar'
import { ChevronRightIcon, StopIcon } from '@heroicons/vue/20/solid'
import SlideTransition from '@/base/components/SlideTransition.vue'

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
  isSidebarOpen: {
    type: Boolean,
    default: true,
  },
})

const { t } = useI18n()
const { menu, refreshOrShowBreadcrumbItem, toggleMenu } = useSidebar(
  props.menuFolders,
  props.menuItems,
  props.menuItemMap
)
</script>
