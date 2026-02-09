import type { ServiceStatus } from '@/types/types.api-monitor'

import { reactive } from 'vue'
import { defineStore } from 'pinia'

export const useServiceStatusStore = defineStore('serviceStatus', () => {
  const serviceStatus = reactive<Record<string, boolean | undefined>>({})

  function updateServiceStatus(statusItems: ServiceStatus[]) {
    for (const statusItem of statusItems) {
      if (serviceStatus[statusItem.name] === undefined) {
        serviceStatus[statusItem.name] = true
      }
      serviceStatus[statusItem.name] = serviceStatus[statusItem.name] && statusItem.status === 0
    }
  }

  function clearServiceStatus(platforms: string[]) {
    for (const platform of platforms) {
      if (serviceStatus[platform] === undefined) {
        continue
      }
      serviceStatus[platform] = undefined
    }
  }

  return {
    serviceStatus,
    updateServiceStatus,
    clearServiceStatus,
  }
})
