import { inject, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { GetAgentWalletLedgerListInput } from '@/base/common/table/getAgentWalletLedgerListInput'
import { round } from '@/base/utils/math'
import { columnSorting } from '@/base/utils/table'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'

export function useAgentWalletLedger() {
  const { t } = useI18n()

  const warn = inject('warn')

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()

  const totalUpScore = ref(0)
  const totalDownScore = ref(0)
  const records = reactive({ items: [] })

  const formInput = reactive({
    id: '',
    agent: { id: constant.Agent.All },
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })
  const tableInput = reactive(
    new GetAgentWalletLedgerListInput(
      formInput.agent.id,
      formInput.id,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'updateTime',
      constant.TableSortDirection.Desc
    )
  )

  function validateForm() {
    let errors = []

    if (!formInput.id) {
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
    }

    return errors
  }

  async function searchRecords() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }
    tableInput.showProcessing = true

    totalUpScore.value = 0
    totalDownScore.value = 0

    const tmpTableInput = new GetAgentWalletLedgerListInput(
      formInput.agent.id,
      formInput.id,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getAgentWalletLedgerList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      const data = resp.data.data
      records.items = data.map((d) => {
        const changeSet = JSON.parse(d.changeset)
        if (changeSet.add_coin > 0) {
          totalUpScore.value += round(changeSet.add_coin, 4)
        } else {
          totalDownScore.value += round(changeSet.add_coin, 4)
        }
        return {
          id: d.id,
          agentName: d.agent_name,
          updateTime: d.update_time,
          kind: d.kind,
          beforeScore: changeSet.before_coin,
          changeScore: changeSet.add_coin,
          afterScore: changeSet.after_coin,
          operator: d.creator,
          remark: d.info,
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

  function downloadRecords() {
    const headers = [
      [
        t('textOrderNumber'),
        t('textAgentName'),
        t('textCoinChangeTime'),
        t('textKind'),
        t('textBeforeScore'),
        t('textChangeScore'),
        t('textAfterScore'),
        t('textOperator'),
        t('textRemark'),
      ],
    ]

    const dir = tableInput.dir
    const column = tableInput.column
    const rows = records.items
      .sort((a, b) => columnSorting(dir, a[column], b[column]))
      .map((d) => {
        return {
          id: d.id,
          agentName: d.agentName,
          updateTime: time.utcTimeStrToLocalTimeFormat(d.updateTime),
          kind: t(`walletLedgerKind__${d.kind}`),
          beforeScore: d.beforeScore,
          changeScore: d.changeScore,
          afterScore: d.afterScore,
          operator: d.operator,
          remark: d.remark,
        }
      })

    xlsx.createAndDownloadFile(
      `[AwlReport]${time.localTimeFormat(tableInput.startTime)}-${time.localTimeFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  return {
    records,
    totalUpScore,
    totalDownScore,
    formInput,
    tableInput,
    searchRecords,
    downloadRecords,
  }
}
