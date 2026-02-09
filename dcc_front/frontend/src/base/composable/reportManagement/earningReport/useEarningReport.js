import { computed, inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysCal'
import constant from '@/base/common/constant'
import { GetPerformanceReportListInput } from '@/base/common/table/getPerformanceReportListInput'
import { useUserStore } from '@/base/store/userStore'
import { round } from '@/base/utils/math'
import { columnSorting } from '@/base/utils/table'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'

export function useEarningReport() {
  const { t } = useI18n()
  const warn = inject('warn')

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()

  const { isAdminUser, isJackpotOpen } = useUserStore()

  const isAdmin = isAdminUser()
  const isShowJackpotRelative = isAdmin || isJackpotOpen()

  const calDirections = computed(() => {
    let ret = []

    ret.push(
      t('textPlatformCalDirection'),
      t('textTotalAgentSettleUpDirection'),
      t('textTotalGamerWinScoreDirection'),
      t('textTotalGamerWinLoseScoreDirection'),
      t('textTotalAgentWinLoseDirection')
    )

    if (isShowJackpotRelative) {
      ret.push(t('textSettleUpIncludeJPDirection'))
    } else {
      ret.push(t('textSettleUpDirection'))
    }

    ret.push(t('textAgentSettleUpDirection'), t('textEarningReportRefreshDirection'), t('textValueDisplayDirection'))

    return ret
  })

  const formInput = reactive({
    agent: { id: constant.Agent.All },
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })

  const sumAgentTableInput = reactive(
    new GetPerformanceReportListInput(
      formInput.agent.id,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'exchangeCurrency',
      constant.TableSortDirection.Desc
    )
  )
  const singleAgentTableInput = reactive({
    agentId: formInput.agent.id,
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
    show: false,
  })

  const sumRecords = reactive({ items: [] })
  const sumValidBet = computed(() => sumRecords.items.reduce((sum, record) => round(sum + record.validBet, 4), 0))
  const sumScore = computed(() => sumRecords.items.reduce((sum, record) => round(sum - record.gamerWinLoseScore, 4), 0))
  const sumAgentSettleUp = computed(() =>
    sumRecords.items.reduce((sum, record) => round(sum + record.agentSettleUp, 4), 0)
  )
  const sumJPInjectWaterScore = computed(() =>
    sumRecords.items.reduce((sum, record) => round(sum + record.totalJPInjectWaterScore, 4), 0)
  )
  const sumJPWinningScore = computed(() =>
    sumRecords.items.reduce((sum, record) => round(sum + record.totalJPlWinningScore, 4), 0)
  )

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

  async function searchRecords() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    sumAgentTableInput.showProcessing = true

    const tmpSumAgentTableInput = new GetPerformanceReportListInput(
      formInput.agent.id,
      formInput.startTime,
      formInput.endTime,
      sumAgentTableInput.length,
      sumAgentTableInput.column,
      sumAgentTableInput.dir
    )

    try {
      const resp = await api.getPerformanceReportList(tmpSumAgentTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      const data = resp.data.data
      sumRecords.items = data.map((d) => {
        const gamerWinLoseScore = round(d.sum_de - d.sum_ya - d.sum_tax + d.sum_bonus, 4)
        const agentSettleUp = round(gamerWinLoseScore * -1 * (d.agent_commission / 100), 4)
        return {
          agentId: d.agentId,
          agentName: d.agent_name,
          betCount: d.bet_count,
          bet: d.sum_ya,
          validBet: d.sum_valid_ya,
          tax: d.sum_tax,
          winScore: d.sum_de,
          bonus: d.sum_bonus,
          gamerWinLoseScore: gamerWinLoseScore,
          agentWinLoseScore: round(gamerWinLoseScore * -1, 4),
          ratio: d.agent_commission,
          agentSettleUp: agentSettleUp < 0 ? 0 : agentSettleUp,
          currency: d.currency,
          exchangeCurrency: agentSettleUp < 0 ? 0 : round(agentSettleUp, 2) * d.to_coin,
          toCoin: d.to_coin,
          jpBetCount: d.jackpot_count,
          totalJPlWinningScore: d.sum_jp_prize_score,
          totalJPInjectWaterScore: d.sum_jp_inject_water_score,
        }
      })

      Object.assign(sumAgentTableInput, tmpSumAgentTableInput)
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      sumAgentTableInput.showProcessing = false
    }
  }

  function downloadRecords() {
    /* notice: 下方header跟rows裡面要依照順序所以if的順序兩邊必須一致 */
    const headers = [[]]

    headers[0].push(t('textAgent'), t('textBetCounts'))

    if (isShowJackpotRelative) {
      headers[0].push(t('textJPBetCount'))
    }

    headers[0].push(
      t('textBetTotal'),
      t('textValidBetTotal'),
      t('textTotalGamerWinScore'),
      t('textTotalGameTax'),
      t('textTotalBonus'),
      t('textTotalGamerWinLoseScore'),
      t('textTotalAgentWinLoseScore'),
      t('textRatio'),
      t('textSettleUpScore')
    )

    if (isShowJackpotRelative) {
      headers[0].push(t('textTotalJPInjectWaterScore'), t('textTotalJPWinScore'))
    }

    headers[0].push(t('textCurrency'), t('textExchangeRate'), t('textAgentSettleUp'))

    const dir = sumAgentTableInput.dir
    const column = sumAgentTableInput.column
    const rows = sumRecords.items
      .sort((a, b) => columnSorting(dir, a[column], b[column]))
      .map((d) => {
        let ret = {}

        ret.agentName = d.agentName
        ret.betCount = d.betCount

        if (isShowJackpotRelative) {
          ret.jpBetCount = d.jpBetCount
        }

        ret.sumBet = d.bet
        ret.validBet = d.validBet
        ret.winScore = d.winScore
        ret.tax = d.tax
        ret.bonus = d.bonus
        ret.gamerWinLoseScore = d.gamerWinLoseScore
        ret.agentWinLoseScore = d.agentWinLoseScore
        ret.ratio = d.ratio
        ret.agentSettleUp = d.agentSettleUp

        if (isShowJackpotRelative) {
          ret.totalJPlWinningScore = d.totalJPlWinningScore
          ret.totalJPInjectWaterScore = d.totalJPInjectWaterScore
        }

        ret.currency = t(`currency__${d.currency}`)
        ret.toCoin = d.toCoin
        ret.exchangeCurrency = d.exchangeCurrency

        return ret
      })

    xlsx.createAndDownloadFile(
      `[PerfReport]${time.recordTimeFormat(sumAgentTableInput.startTime)}-${time.recordTimeFormat(
        sumAgentTableInput.endTime
      )}`,
      headers,
      rows
    )
  }

  function showDetail(agentId) {
    singleAgentTableInput.agentId = agentId
    singleAgentTableInput.startTime = sumAgentTableInput.startTime
    singleAgentTableInput.endTime = sumAgentTableInput.endTime
    singleAgentTableInput.show = true
  }

  return {
    calDirections,
    formInput,
    isAdminUser: isAdmin,
    isShowJackpotRelative,
    singleAgentTableInput,
    sumAgentTableInput,
    sumValidBet,
    sumRecords,
    sumScore,
    sumAgentSettleUp,
    sumJPInjectWaterScore,
    sumJPWinningScore,
    downloadRecords,
    searchRecords,
    showDetail,
  }
}
