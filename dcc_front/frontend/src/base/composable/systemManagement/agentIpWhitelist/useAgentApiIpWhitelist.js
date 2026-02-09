import { inject, reactive, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import time from '@/base/utils/time'

export function useAgentApiIpWhitelist(props, emit) {
  const warn = inject('warn')
  const { t } = useI18n()

  const visible = ref(false)

  const agentIpWhitelist = reactive({ items: [] })

  async function searchAgentApiIpWhitelist() {
    try {
      const resp = await api.getAgentApiIpWhitelist({ agent_id: props.agentId })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      agentIpWhitelist.items = resp.data.data
        ? resp.data.data.map((d) => {
            return {
              createTime: d.create_time,
              createTimeStr: time.localTimeFormat(d.create_time * 1000),
              ipAddress: d.ip_address,
              info: d.info,
              creator: d.creator,
            }
          })
        : []

      visible.value = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }

      emit('close')
    }
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        await searchAgentApiIpWhitelist()
      } else {
        visible.value = false
      }
    }
  )

  async function updateAgentApiIpWhitelist(agentIpWhitelistItems) {
    try {
      const resp = await api.setAgentApiIpWhitelist({
        agent_id: props.agentId,
        ip_whitelist: agentIpWhitelistItems.map((d) => {
          return {
            create_time: d.createTime,
            ip_address: d.ipAddress,
            info: d.info,
            creator: d.creator,
          }
        }),
      })

      warn(t(`errorCode__${resp.data.code}`)).then(async () => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        emit('refreshIpList')
        emit('close')
      })
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    }
  }

  return {
    agentIpWhitelist,
    visible,
    updateAgentApiIpWhitelist,
  }
}
