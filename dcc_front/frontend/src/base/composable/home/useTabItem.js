import { computed, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'

export function useTabItem(itemKey) {
  const { currentBreadcrumb } = storeToRefs(useBreadcrumbStore())

  const isInitialized = ref(false)
  const isShow = computed(() => {
    return currentBreadcrumb.value.item !== null ? currentBreadcrumb.value.item.key === itemKey : false
  })

  watch(
    isShow,
    () => {
      if (isShow.value && !isInitialized.value) {
        isInitialized.value = true
      }
    },
    { immediate: true }
  )

  return {
    isInitialized,
    isShow,
  }
}
