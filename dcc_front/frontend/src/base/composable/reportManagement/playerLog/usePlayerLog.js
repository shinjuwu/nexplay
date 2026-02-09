import { watch, inject, reactive, ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { round } from '@/base/utils/math'
import { columnSorting } from '@/base/utils/table'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'
import { GetPlayerLogListInput } from '@/base/common/table/getPlayerLogListInput'

export function usePlayerLog(props) {
  const { t } = useI18n()
  const warn = inject('warn')

  const calDirections = computed(() => {
    return [
      t('textTotalWinLoseScoreDirection'),
      t('textPlayerWinLoseScoreDirection'),
      t('textDailySettlementReportDeductTaxRtpDirection'),
      t('textValueDisplayDirection'),
    ]
  })
  const timeRange = time.getCurrentTimeRageByMinutesAndStep(60, 60)

  const records = reactive({ items: [] })

  const formInput = reactive({
    agent: { id: constant.Agent.All },
    userName: '',
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })

  const tableInput = reactive(
    new GetPlayerLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'deductTaxRtp',
      constant.TableSortDirection.Desc
    )
  )

  const sumWinLoseScore = ref(0)
  const avgDeductTaxRtp = ref(0)

  function validateForm() {
    let errors = []

    const userName = formInput.userName
    if (!userName) {
      errors.push(t('textUserNameEmptyError'))
      return errors
    }

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

  async function searchRecords() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }
    tableInput.showProcessing = true
    let totalGamerBet = 0
    let totalTax = 0
    let totalGamerWinScore = 0
    let totalWinLoseScore = 0

    const tmpTableInput = new GetPlayerLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getGameUsersStatHourList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      records.items = resp.data.data.map((d) => {
        totalGamerBet += round(d.ya_score, 4)
        totalTax += round(d.tax, 4)
        totalGamerWinScore += round(d.de_score + d.bonus, 4)
        totalWinLoseScore += round(d.de_score - d.ya_score - d.tax + d.bonus, 4)
        return {
          agentId: d.agent_id,
          userId: d.userId,
          agentName: d.agent_name,
          userName: d.username,
          gameId: d.game_id,
          betCounts: d.play_count,
          bet: d.ya_score,
          score: d.de_score,
          tax: d.tax,
          bonus: d.bonus,
          winLoseScore: round(d.de_score - d.ya_score - d.tax + d.bonus, 4),
          deductTaxRtp: round((d.de_score - d.tax + d.bonus) / d.ya_score, 6),
        }
      })
      sumWinLoseScore.value = totalWinLoseScore
      avgDeductTaxRtp.value = round(((totalGamerWinScore - totalTax) / totalGamerBet) * 100, 4) || 0

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
    if (!records.items.length) {
      warn(t('textNoRecordForExport'))
      return
    }

    const headers = [
      [
        t('textAgentName'),
        t('textPlayerAccount'),
        t('textGameName'),
        t('textBetCounts'),
        t('textBetTotal'),
        t('textTotalGamerWinScore'),
        t('textTotalGameTax'),
        t('textTotalBonus'),
        t('textGamerWinLoseScore'),
        t('textDeductTaxRtp'),
      ],
    ]

    const dir = tableInput.dir
    const column = tableInput.column
    const rows = records.items
      .sort((a, b) => columnSorting(dir, a[column], b[column]))
      .map((d) => {
        return {
          agentName: d.agentName,
          userName: d.userName,
          gameId: t(`game__${d.gameId}`),
          betCounts: d.betCounts,
          bet: d.bet,
          score: d.score,
          tax: d.tax,
          bonus: d.bonus,
          winLoseScore: d.winLoseScore,
          deductTaxRtp: d.deductTaxRtp,
        }
      })

    xlsx.createAndDownloadFile(
      `[PlReport]${time.localTimeFormat(tableInput.startTime)}-${time.localTimeFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  watch(
    props,
    async () => {
      if (props.userName && props.agentId !== -1) {
        formInput.agent.id = props.agentId
        formInput.userName = props.userName
      }
    },
    { immediate: true }
  )

  return {
    avgDeductTaxRtp,
    calDirections,
    formInput,
    records,
    sumWinLoseScore,
    tableInput,
    downloadRecords,
    searchRecords,
  }
}
