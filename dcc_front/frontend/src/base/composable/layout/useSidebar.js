import { reactive } from 'vue'
import { storeToRefs } from 'pinia'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'
import { useUserStore } from '@/base/store/userStore'
import { getMenu } from '@/base/utils/menu'

export function useSidebar(menuFolders, menuItems, menuItemMap) {
  const { addBreadcrumbItem } = useBreadcrumbStore()
  const { user } = storeToRefs(useUserStore())

  const menu = reactive({
    items: getMenu(menuFolders, menuItems, user),
  })

  function refreshOrShowBreadcrumbItem(item) {
    const menuItem = menuItemMap[item.key]
    const defaultProps = menuItem.createProps()

    Object.keys(defaultProps).forEach((key) => {
      item.props[key] = defaultProps[key]
    })

    addBreadcrumbItem(item)
  }

  function toggleMenu(menuIdx) {
    menu.items.forEach((item, idx) => {
      item.show = idx === menuIdx ? !item.show : false
    })
  }

  return {
    menu,
    refreshOrShowBreadcrumbItem,
    toggleMenu,
  }
}
