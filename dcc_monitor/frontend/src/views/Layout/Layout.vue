<script setup lang="ts">
import type { ComponentPublicInstance } from 'vue'
import type { ServiceStatusResponse } from '@/types/types.api-monitor'

import { useRoute, useRouter } from 'vue-router'
import { computed, onBeforeMount, onMounted, onUnmounted, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'

import {
  setAxiosInstanceConfig,
  clearAxiosInstanceConfig,
  processApiRequest,
  parseServerErrorMessage,
} from '@/api/base'
import { verifyToken } from '@/api/account'
import { serviceStatus } from '@/api/monitor'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { useServiceStatusStore } from '@/store/serviceStatusStore'
import timer from '@/utils/timer'

import LayoutTopBar from '@/components/LayoutTopBar.vue'
import LayoutRightAside from '@/components/LayoutRightAside.vue'
import { useUserStore } from '@/store/userStore'

const isShowRightAside = ref(false)

function toggleRightAside() {
  isShowRightAside.value = !isShowRightAside.value
}

watch(isShowRightAside, (newValue) => {
  if (newValue) {
    document.body.classList.add('overflow-hidden')
  } else {
    document.body.classList.remove('overflow-hidden')
  }
})

const layoutTopBarRef = ref<ComponentPublicInstance | null>(null)
const mainRef = ref<HTMLElement | null>(null)

/** 重新計算content高度 */
function resizeMainContenHeight() {
  if (layoutTopBarRef.value == null || mainRef.value == null) {
    return
  }

  const layoutTopBar = layoutTopBarRef.value.$el as HTMLElement
  const main = mainRef.value

  const windowHeight = window.innerHeight
  const headerHeight = parseFloat(getComputedStyle(layoutTopBar).height)

  main.style.minHeight = `${windowHeight - headerHeight}px`
}

onMounted(() => {
  resizeMainContenHeight()
  window.addEventListener('resize', resizeMainContenHeight)
})

onUnmounted(() => {
  window.removeEventListener('resize', resizeMainContenHeight)
})

const route = useRoute()
const routePath = computed(() => route.path)

watch(routePath, () => {
  isShowRightAside.value = false
})

const { updateServiceStatus, clearServiceStatus } = useServiceStatusStore()
const { warn } = useDialogStore()
const userStore = useUserStore()
const { user } = storeToRefs(userStore)
const { hasPermission } = userStore

async function executeServiceStatus() {
  const filter = user.value.permissions.length > 1 ? 'all' : user.value.permissions[0]
  const axiosResp = await serviceStatus({ filter: filter })

  if (axiosResp.data.code !== responseCode.Success) {
    warn(parseServerErrorMessage(axiosResp))
    return
  }

  const data = JSON.parse(axiosResp.data.data) as ServiceStatusResponse
  if (data.status_list === null || data.status_list.length === 0) {
    return
  }

  updateServiceStatus(data.status_list)

  for (let i = 0; i < data.status_list.length; i++) {
    const serviceStatus = data.status_list[i]

    if (!hasPermission(serviceStatus.name)) {
      continue
    }

    if (serviceStatus.status === 0) {
      continue
    }

    warn(`${serviceStatus.info}无法连线，请立即通知工作人员`)
  }
}

function requestServiceStatus() {
  processApiRequest(executeServiceStatus, warn)
}

const router = useRouter()

onBeforeMount(() => {
  setAxiosInstanceConfig(router)

  processApiRequest(async () => {
    await verifyToken()

    const permissions = user.value.permissions.length
    if (permissions > 0) {
      await executeServiceStatus()
      timer.register(60 * 1000, requestServiceStatus)
    }
  }, warn)
})

onUnmounted(() => {
  clearAxiosInstanceConfig()
  clearServiceStatus(user.value.permissions)

  timer.removeAll()
})
</script>

<template>
  <LayoutTopBar
    ref="layoutTopBarRef"
    :is-show-right-aside="isShowRightAside"
    @toggle-right-aside="toggleRightAside()"
  />
  <LayoutRightAside :is-show="isShowRightAside" @toggle-is-show="toggleRightAside()" />

  <main
    ref="mainRef"
    class="mt-10 space-y-4 bg-gray-300 p-4 transition-transform"
    :class="{ 'translate-x-[-250px]': isShowRightAside }"
  >
    <router-view />
  </main>

  <div
    v-show="isShowRightAside"
    class="fixed inset-0 z-30 bg-black/50 transition-transform"
    :class="{ 'translate-x-[-250px]': isShowRightAside }"
    @click="isShowRightAside = false"
  ></div>
</template>
