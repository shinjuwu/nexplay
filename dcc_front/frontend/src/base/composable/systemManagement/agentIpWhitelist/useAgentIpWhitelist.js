import { computed, inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'
import { GetAgentIpWhitelistListInput } from '@/base/common/table/getAgentIpWhitelistListInput'

export function useAgentIpWhitelist() {
  const { t } = useI18n()
  const warn = inject('warn')

  const isEditEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.AgentIpWhitelistUpdate)
  })

  const formAgentIpWhitelistListInput = reactive({
    agent: { id: constant.Agent.All },
  })
  const tableAgentIpWhitelistListInput = reactive(
    new GetAgentIpWhitelistListInput(
      formAgentIpWhitelistListInput.agent.id,
      constant.TableDefaultLength,
      'id',
      constant.TableSortDirection.Asc
    )
  )

  const agentIpWhitelistList = reactive({ items: [] })

  async function searchAgentIpWhitelistInfos() {
    try {
      tableAgentIpWhitelistListInput.showProcessing = true

      const tmpTableAgentIpWhitelistListInput = new GetAgentIpWhitelistListInput(
        formAgentIpWhitelistListInput.agent.id,
        tableAgentIpWhitelistListInput.length,
        tableAgentIpWhitelistListInput.column,
        tableAgentIpWhitelistListInput.dir
      )

      const resp = await api.getAgentIpWhitelistList(tmpTableAgentIpWhitelistListInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      agentIpWhitelistList.items = resp.data.data.map((d) => {
        return {
          id: d.id,
          name: d.name,
          levelCode: d.level_code,
          ipAddressCount: d.ip_address_count,
          apiIpAddressCount: d.api_ip_address_count,
          adminUsername: d.admin_username,
        }
      })

      Object.assign(tableAgentIpWhitelistListInput, tmpTableAgentIpWhitelistListInput)
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableAgentIpWhitelistListInput.showProcessing = false
    }
  }

  const backendIpInfo = reactive({ agentId: -1, agentName: '', mode: '', visible: false })
  const apiIpInfo = reactive({ id: -1, name: '', mode: '' })

  function showDialog(reactiveObj, agentId, agentName, mode) {
    reactiveObj.agentId = agentId
    reactiveObj.agentName = agentName
    reactiveObj.mode = mode
    reactiveObj.visible = true
  }

  function showBackendIpDialog(agentId, agentName, mode) {
    showDialog(backendIpInfo, agentId, agentName, mode)
  }

  function showApiIpDialog(agentId, agentName, mode) {
    showDialog(apiIpInfo, agentId, agentName, mode)
  }

  return {
    agentIpWhitelistList,
    apiIpInfo,
    backendIpInfo,
    formAgentIpWhitelistListInput,
    isEditEnabled,
    tableAgentIpWhitelistListInput,
    searchAgentIpWhitelistInfos,
    showBackendIpDialog,
    showApiIpDialog,
  }
}
