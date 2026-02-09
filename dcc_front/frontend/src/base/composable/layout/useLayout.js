import { computed, inject, onBeforeMount, onMounted, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useChatServiceStore } from '@/base/store/chatServiceStore'
import { useUserStore } from '@/base/store/userStore'
import * as utilRequest from '@/base/utils/request'

export function useLayout() {
  const warn = inject('warn')

  const { t } = useI18n()
  const router = useRouter()
  const chatServiceStore = useChatServiceStore()

  const isSidebarOpen = ref(window.innerWidth >= 768)
  const isShowPersonalInfo = ref(false)

  const headerEl = ref()
  const mainEl = ref()
  const footerEl = ref()

  // Notice: 目前只有公告會用到，如果以後有開放聊天需要調整判斷
  const hasChatService = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.BackendAnnouncementRead)
  })

  /** 重新計算content高度 */
  function resizeMainContenHeight() {
    const header = headerEl.value.$el
    const main = mainEl.value
    const footer = footerEl.value.$el

    const windowHeight = window.innerHeight
    const headerHeight = parseFloat(getComputedStyle(header).height)
    const footerHeight = parseFloat(getComputedStyle(footer).height)

    main.style.minHeight = `${windowHeight - headerHeight - footerHeight}px`
  }

  onBeforeMount(async () => {
    utilRequest.setAxiosInstanceConfig(router)

    try {
      await api.ping()

      if (hasChatService.value) {
        const { init: chatServiceInit, start: chatServiceStart } = chatServiceStore
        await chatServiceInit()
        await chatServiceStart()
      }
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    }
  })

  onMounted(() => {
    resizeMainContenHeight()
    window.addEventListener('resize', resizeMainContenHeight)
  })

  onUnmounted(() => {
    if (hasChatService.value) {
      const { stop: chatServiceStop } = chatServiceStore
      chatServiceStop()
    }

    window.removeEventListener('resize', resizeMainContenHeight)
  })

  return {
    isSidebarOpen,
    isShowPersonalInfo,
    headerEl,
    mainEl,
    footerEl,
  }
}
