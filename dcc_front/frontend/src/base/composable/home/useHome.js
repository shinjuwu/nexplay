import { onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { dragscroll } from 'vue-dragscroll'
import { storeToRefs } from 'pinia'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'

export function useHome() {
  // create v-dragscroll Directives
  const vDragscroll = dragscroll

  const { t } = useI18n()

  const bStore = useBreadcrumbStore()
  const { breadcrumbs, currentBreadcrumb } = storeToRefs(bStore)
  const { addBreadcrumbItem, removeBreadcrumbItem, clearOtherBreadcrumbItems, clearBreadcrumbItems } = bStore

  onUnmounted(() => {
    clearBreadcrumbItems()
  })

  return {
    vDragscroll,
    t,
    breadcrumbs,
    currentBreadcrumb,
    addBreadcrumbItem,
    removeBreadcrumbItem,
    clearOtherBreadcrumbItems,
  }
}
