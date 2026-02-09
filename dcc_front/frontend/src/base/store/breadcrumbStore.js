import { inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import { defineStore } from 'pinia'

export const useBreadcrumbStore = defineStore('breadcrumb', () => {
  const warn = inject('warn')
  const { t } = useI18n()

  const breadcrumbs = reactive({
    items: [],
  })
  const currentBreadcrumb = reactive({
    item: null,
  })

  function addBreadcrumbItem(item, forceUpdate) {
    const storeItem = breadcrumbs.items.find((i) => i.key === item.key)
    if (storeItem === undefined && breadcrumbs.items.length === 15) {
      warn(t('textTooMuchTab'))
      return
    }

    if (storeItem === undefined) {
      breadcrumbs.items.push(item)
    }
    currentBreadcrumb.item = item
    if (forceUpdate) {
      currentBreadcrumb.item.forceUpdateCount++
    }
  }

  function removeBreadcrumbItem(item) {
    breadcrumbs.items = breadcrumbs.items.filter((i) => i.key !== item.key)

    if (currentBreadcrumb.item.key === item.key) {
      const items = breadcrumbs.items
      currentBreadcrumb.item = items.length > 0 ? items[items.length - 1] : null
    }
  }

  function clearOtherBreadcrumbItems() {
    if (currentBreadcrumb.item === null) {
      return
    }

    breadcrumbs.items = [currentBreadcrumb.item]
  }

  function clearBreadcrumbItems() {
    breadcrumbs.items = []
    currentBreadcrumb.item = null
  }

  return {
    breadcrumbs,
    currentBreadcrumb,
    addBreadcrumbItem,
    removeBreadcrumbItem,
    clearOtherBreadcrumbItems,
    clearBreadcrumbItems,
  }
})
