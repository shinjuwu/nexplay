<template>
  <div class="bg-gray-300">
    <div class="flex min-h-[43px] justify-between">
      <ul v-dragscroll.x class="flex overflow-hidden">
        <li v-for="item in breadcrumbs.items" :key="`tab_${item.key}`" class="inline-block">
          <a
            class="inline-block cursor-pointer whitespace-nowrap p-4"
            :class="
              currentBreadcrumb.item.key === item.key
                ? 'bg-gray-200 text-gray-500'
                : 'text-gray-400 hover:bg-gray-400 hover:text-white'
            "
          >
            <p class="text-nowrap flex items-center justify-around">
              <span @click="addBreadcrumbItem(item)">{{ t(`menuItem${item.key}`) }}</span>
              <span class="ml-1" @click="removeBreadcrumbItem(item)">
                <XCircleIcon class="h-6 w-6" />
              </span>
            </p>
          </a>
        </li>
      </ul>
      <div class="flex py-1 px-4">
        <button
          v-show="breadcrumbs.items.length > 0"
          class="btn btn-secondary"
          :disabled="currentBreadcrumb.item === null"
          @click="clearOtherBreadcrumbItems()"
        >
          {{ t('textCloseOtherTabs') }}
        </button>
      </div>
    </div>
  </div>
  <div class="p-4">
    <img v-show="breadcrumbs.items.length === 0" class="w-full" src="@/base/assets/images/pic_welcome.jpg" />
    <BaseTabItem
      v-for="item in breadcrumbs.items"
      :key="`content_${item.key}`"
      :item="item"
      :menu-item-map="props.menuItemMap"
    />
  </div>
</template>

<script setup>
import { useHome } from '@/base/composable/home/useHome'
import { XCircleIcon } from '@heroicons/vue/24/outline'
import BaseTabItem from '@/base/views/Home/BaseTabItem.vue'

const {
  vDragscroll,
  t,
  breadcrumbs,
  currentBreadcrumb,
  addBreadcrumbItem,
  removeBreadcrumbItem,
  clearOtherBreadcrumbItems,
} = useHome()

const props = defineProps({
  menuItemMap: {
    type: Object,
    default: () => {},
  },
})
</script>
