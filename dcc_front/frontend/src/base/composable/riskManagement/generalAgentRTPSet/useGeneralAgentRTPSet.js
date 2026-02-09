import { reactive, ref, inject, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRiskControl'
import constant from '@/base/common/constant'
import { useUserStore } from '@/base/store/userStore'
import { roleItemKey } from '@/base/common/menuConstant'
import { GetGeneralAgentRTPSetListInput } from '@/base/common/table/getGeneralAgentRTPSetListInput'
import { round } from '@/base/utils/math'

export function useGeneralAgentRTPSet() {
  const { t } = useI18n()
  const warn = inject('warn')

  const visible = ref(false)

  const accountStates = {
    All: -1,
    Disabled: 0,
    Enabled: 1,
  }
  const checkStates = computed(() => Object.values(accountStates).filter((state) => state !== -1))

  const isSettingEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GeneralAgentRTPSetUpdate)
  })

  const formInput = reactive({
    agent: { id: constant.Agent.All },
    state: accountStates.All,
  })

  const tableInput = reactive(
    new GetGeneralAgentRTPSetListInput(
      formInput.agent.id,
      formInput.state,
      constant.TableDefaultLength,
      'createDate',
      constant.TableSortDirection.Asc
    )
  )

  const records = reactive({ items: [] })
  const filterRecords = computed(() => {
    if (formInput.state === accountStates.All) {
      return records.items
    } else {
      return records.items.filter((r) => r.state === formInput.state)
    }
  })

  const agentRatioInfo = reactive({
    agentId: '',
    agentName: '',
    state: '',
    ratio: '',
    remark: '',
  })

  function validateForm(payload) {
    const errors = []

    if (payload.ratio < 0 || payload.ratio > 100) {
      errors.push(t('textRTPRangeError'))
    }

    return errors
  }

  async function searchRecords() {
    tableInput.showProcessing = true
    try {
      const tmpTableInput = new GetGeneralAgentRTPSetListInput(
        formInput.agent.id,
        formInput.state,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )

      const resp = await api.getAgentIncomeRatioList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      records.items = data.map((d) => {
        return {
          agentId: d.agent_id,
          agentName: d.agent_name,
          ratio: round(d.ratio * 100, 4),
          updateTime: d.update_time,
          state: d.state ? 1 : 0,
        }
      })

      Object.assign(tableInput, tmpTableInput)
    } catch (err) {
      console.error(err)
      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableInput.showProcessing = false
    }
  }

  async function getAgentRatio(payload) {
    try {
      const resp = await api.getAgentIncomeRatio({ agent_id: payload.agentId })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      agentRatioInfo.agentId = data.agent_id
      agentRatioInfo.agentName = data.agent_name
      agentRatioInfo.ratio = round(data.ratio * 100, 4)
      agentRatioInfo.state = data.state ? 1 : 0
      agentRatioInfo.remark = data.info

      visible.value = true
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

  async function setAgentRatio(payload) {
    const errors = validateForm(payload)
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const formatVal = round(agentRatioInfo.ratio / 100, 3)

      const resp = await api.setAgentIncomeRatio({
        agent_id: payload.agentId,
        ratio: formatVal,
        info: payload.remark,
        state: payload.state === 1,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        visible.value = false
        searchRecords()
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
    accountStates,
    agentRatioInfo,
    checkStates,
    filterRecords,
    formInput,
    tableInput,
    visible,
    isSettingEnabled,
    getAgentRatio,
    searchRecords,
    setAgentRatio,
  }
}
