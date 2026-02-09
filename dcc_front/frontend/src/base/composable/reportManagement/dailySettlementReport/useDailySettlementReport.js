import { ref, inject, reactive, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { GetDailySettlementReportListInput } from '@/base/common/table/getDailySettlementReportListInput'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'
import { roomTypeNameIndex } from '@/base/utils/room'

export function useDailySettlementReport() {
  const { t } = useI18n()
  const warn = inject('warn')

  const calDirections = computed(() => {
    return [
      t('textPlatformCalDirection'),
      t('textTotalGamerWinLoseScoreDirection'),
      t('textDailySettlementReportDeductTaxRtpDirection'),
      t('textValueDisplayDirection'),
    ]
  })

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()
  const formInput = reactive({
    agent: { id: constant.Agent.All },
    game: { id: 0 },
    roomType: '',
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })

  const tableInput = reactive(
    new GetDailySettlementReportListInput(
      formInput.agent.id,
      formInput.game.id,
      formInput.roomType,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'taxedrtp',
      constant.TableSortDirection.Desc
    )
  )

  const agentDailyStat = reactive({ items: [] })

  const sumScore = ref(0)

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

  async function searchRecords(input) {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      tableInput.showProcessing = true

      if (!input) {
        input = new GetDailySettlementReportListInput(
          formInput.agent.id,
          formInput.game.id,
          formInput.roomType,
          formInput.startTime,
          formInput.endTime,
          tableInput.length,
          tableInput.column,
          tableInput.dir
        )
      }

      const resp = await api.getAgentGameRatioStatList(input.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      const pageData = resp.data.data
      if (pageData.draw === 0) {
        sumScore.value = pageData.total_platform_winlose_score
        input.totalRecords = pageData.recordsTotal
      }

      agentDailyStat.items = pageData.data.map((d) => {
        return {
          agentId: d.agent_id,
          agentName: d.agent_name,
          roomType: d.room_type,
          gameId: d.game_id,
          betCounts: d.play_count,
          bet: d.ya_score,
          score: d.de_score,
          tax: d.tax,
          bonus: d.bonus,
          winLoseScore: d.player_winlose,
          deductTaxRtp: d.taxed_rtp,
        }
      })

      Object.assign(tableInput, input)
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
    if (!agentDailyStat.items.length) {
      warn(t('textNoRecordForExport'))
      return
    }

    const headers = [
      [
        t('textAgent'),
        t('textGameName'),
        t('textRoomType'),
        t('textBetCounts'),
        t('textBetTotal'),
        t('textTotalGamerWinScore'),
        t('textTotalGameTax'),
        t('textTotalBonus'),
        t('textTotalGamerWinLoseScore'),
        t('textDeductTaxRtp'),
      ],
    ]

    const rows = agentDailyStat.items.map((d) => {
      return {
        agentName: d.agentName,
        roomType: t(`roomType__${roomTypeNameIndex(d.gameId, d.roomType)}`),
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
      `[DsReport]${time.localDateFormat(tableInput.startTime)}-${time.localDateFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  return {
    formInput,
    tableInput,
    sumScore,
    calDirections,
    agentDailyStat,
    downloadRecords,
    searchRecords,
  }
}
