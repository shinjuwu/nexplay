import { reactive, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import constant from '@/base/common/constant'
import { GetAutoRiskControlLogListInput } from '@/base/common/table/getAutoRiskControlLogListInput.js'
import * as api from '@/base/api/sysRecord'
import time from '@/base/utils/time'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'
import { getMenuItemFromMenu } from '@/base/utils/menu'
import { menuItemKey } from '@/base/common/menuConstant'

export function useAutoRiskControlLog() {
  const { t } = useI18n()
  const warn = inject('warn')

  const timeRange = time.getCurrentTimeRageByMinutesAndStep(60, 60)

  const formInput = reactive({
    agent: { id: constant.Agent.All },
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
    userName: '',
  })

  const tableInput = reactive(
    new GetAutoRiskControlLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'updateTime',
      constant.TableSortDirection.Asc
    )
  )

  const autoRiskControlLog = reactive({ items: [] })

  const riskDisposeWay = {
    1: t('textDisabled'),
    2: t('riskControlTag__0001'),
    3: t('riskControlTag__0100'),
  }

  async function getAutoRiskControlLogList() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      tableInput.showProcessing = true

      const tmpTableInput = new GetAutoRiskControlLogListInput(
        formInput.agent.id,
        formInput.userName,
        formInput.startTime,
        formInput.endTime,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )

      const resp = await api.getAutoRiskControlLogList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      autoRiskControlLog.items = resp.data.data.map((log) => {
        return {
          createTime: time.utcTimeStrToLocalTimeFormat(log.create_time),
          agentId: log.agent_id,
          agentName: log.agent_name,
          userName: log.username,
          riskCode: log.risk_code,
        }
      })
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

  function redirectToPlayerAccount(agentId, userName) {
    let item = getMenuItemFromMenu(menuItemKey.PlayerAccount)
    item.props.agentId = agentId
    item.props.userName = userName

    const { addBreadcrumbItem } = useBreadcrumbStore()
    addBreadcrumbItem(item)
  }

  function validateForm() {
    let errors = []

    const startTime = formInput.startTime
    const endTime = formInput.endTime
    if (endTime - startTime < 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
      return errors
    } else if (endTime - startTime > time.commonReportTimeRange * 24 * 60 * 60 * 1000) {
      errors.push(
        t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t('fmtTextDays', [time.commonReportTimeRange])])
      )
      return errors
    }

    return errors
  }

  return {
    time,
    formInput,
    autoRiskControlLog,
    tableInput,
    riskDisposeWay,
    getAutoRiskControlLogList,
    redirectToPlayerAccount,
  }
}
