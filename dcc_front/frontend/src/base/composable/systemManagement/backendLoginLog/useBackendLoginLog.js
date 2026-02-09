import { computed, inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { GetBackendLoginLogListInput } from '@/base/common/table/getBackendLoginLogListInput'
import { useUserStore } from '@/base/store/userStore'
import time from '@/base/utils/time'

export function useBackendLoginLog() {
  const warn = inject('warn')
  const { t } = useI18n()
  const { user } = storeToRefs(useUserStore())

  const timeRange = time.getCurrentTimeRageByMinutesAndStep(30, time.earningReportTimeMinuteIncrement)
  const formInput = reactive({
    agent: { id: user.value.agentId },
    userName: '',
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })
  const tableInput = reactive(
    new GetBackendLoginLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'loginTime',
      constant.TableSortDirection.Asc
    )
  )

  const records = reactive({ items: [] })
  const filterRecords = computed(() => {
    let finalRecords = records.items
    if (tableInput.ip !== '') {
      finalRecords = finalRecords.filter((r) => r.ip.includes(tableInput.ip))
    }
    tableInput.filterAdjust(finalRecords.length)
    return finalRecords
  })

  async function searchRecords() {
    tableInput.showProcessing = true

    const tmpTableInput = new GetBackendLoginLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getBackendLoginLogList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      records.items = data.map((d) => {
        return {
          agentId: d.agent_id,
          agentName: d.agent_name,
          userName: d.username,
          ip: d.ip,
          loginTime: d.login_time,
          status: d.error_code === constant.ErrorCode.Success,
          errorCode: d.error_code,
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

  return {
    filterRecords,
    formInput,
    tableInput,
    searchRecords,
  }
}
