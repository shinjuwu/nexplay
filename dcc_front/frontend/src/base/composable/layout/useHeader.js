import { computed, nextTick, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { menuItemKey, roleItemKey } from '@/base/common/menuConstant'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'
import { useChatServiceStore, chatServerMessageSubject } from '@/base/store/chatServiceStore'
import { useUserStore } from '@/base/store/userStore'
import { getMenuItemFromMenu } from '@/base/utils/menu'

export function useHeader() {
  const router = useRouter()

  const userName = computed(() => {
    const { user } = storeToRefs(useUserStore())
    return user.value.nickName || user.value.name
  })

  const isShowUserInfoList = ref(false)
  const userInfoToggleEl = ref()
  const userInfoListEl = ref()

  function onBlurUserInfoList(e) {
    if (
      e.relatedTarget !== null &&
      (e.relatedTarget === userInfoToggleEl.value || userInfoListEl.value.contains(e.relatedTarget))
    ) {
      return
    }
    isShowUserInfoList.value = false
  }

  /** 使用者登出 */
  function userSignOut() {
    const { signOut } = useUserStore()
    signOut(router)
  }

  watch(isShowUserInfoList, (newValue) => {
    if (newValue) {
      nextTick(() => {
        userInfoListEl.value.focus()
      })
    }
  })

  const hasBackendAnnoucementRole = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.BackendAnnouncementRead)
  })
  const hasNewAnnoucements = computed(() => {
    const { unreadMessages } = storeToRefs(useChatServiceStore())
    return unreadMessages.value[chatServerMessageSubject.announcement].length > 0
  })

  function redirectToBackendAnnoucement() {
    let item = getMenuItemFromMenu(menuItemKey.BackendAnnouncement)

    const { addBreadcrumbItem } = useBreadcrumbStore()
    addBreadcrumbItem(item, hasNewAnnoucements.value)
  }

  return {
    hasBackendAnnoucementRole,
    hasNewAnnoucements,
    isShowUserInfoList,
    userInfoListEl,
    userInfoToggleEl,
    userName,
    onBlurUserInfoList,
    redirectToBackendAnnoucement,
    userSignOut,
  }
}
