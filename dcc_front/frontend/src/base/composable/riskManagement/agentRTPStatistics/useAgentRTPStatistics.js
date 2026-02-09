import { reactive, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import constant from '@/base/common/constant'
import { GetAgentRTPStatisticsListInput } from '@/base/common/table/getAgentRTPStatisticsListInput'
import * as api from '@/base/api/sysRiskControl'
import { isSameMonth } from 'date-fns'
import { round } from '@/base/utils/math'
import time from '@/base/utils/time'

export function useAgentRTPStatistics() {
  const { t } = useI18n()
  const warn = inject('warn')

  const calDirections = reactive({
    items: [t('textAgentGameRatioRtpDirection'), t('textAgentGameRatioDeductTaxRtpDirection')],
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
    new GetAgentRTPStatisticsListInput(
      formInput.agent.id,
      formInput.game.id,
      formInput.roomType,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'agentName',
      constant.TableSortDirection.Asc
    )
  )

  const records = reactive({
    items: [],
  })

  const recordsInfo = reactive({
    totalGamerBet: 0,
    totalGamerWinScore: 0,
    totalTax: 0,
    totalBonus: 0,
    avgRtp: 0,
    avgDeductTaxRtp: 0,
  })

  function validate() {
    let errors = []
    if (formInput.startTime - formInput.endTime > 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
      return errors
    }
    if (!isSameMonth(formInput.startTime, formInput.endTime)) {
      errors.push(t('textAcrossMonthErrorMessage'))
      return errors
    }

    return errors
  }

  async function searchRecords() {
    const errors = validate()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      tableInput.showProcessing = true

      const tmpTableInput = new GetAgentRTPStatisticsListInput(
        formInput.agent.id,
        formInput.game.id,
        formInput.roomType,
        formInput.startTime,
        formInput.endTime,
        constant.TableDefaultLength,
        'agentName',
        constant.TableSortDirection.Asc
      )

      const resp = await api.getAgentIncomeRatioAndGameData(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      recordsInfo.totalGamerBet = 0
      recordsInfo.totalGamerWinScore = 0
      recordsInfo.totalTax = 0
      recordsInfo.totalBonus = 0
      recordsInfo.avgRtp = 0
      recordsInfo.avgDeductTaxRtp = 0

      const data = resp.data.data
      records.items = data
        ? Object.values(data).map((d) => {
            const deScore = round(d.de_score + d.bonus, 4)
            recordsInfo.totalGamerBet += round(d.ya_score, 4)
            recordsInfo.totalGamerWinScore += round(d.de_score, 4)
            recordsInfo.totalTax += round(d.tax, 4)
            recordsInfo.totalBonus += round(d.bonus, 4)
            return {
              agentId: d.agent_id,
              agentName: d.agent_name,
              gameId: d.game_id,
              gameType: d.game_type,
              gameCode: d.game_code,
              roomType: d.room_type,
              validYaScore: d.vaild_ya_score,
              yaScore: d.ya_score,
              deScore: d.de_score,
              tax: d.tax,
              bonus: d.bonus,
              rtp: d.ya_score > 0 ? round((deScore / d.ya_score) * 100, 4) : 0,
              deductTaxRtp: d.ya_score > 0 ? round(((deScore - d.tax) / d.ya_score) * 100, 4) : 0,
              lastRTPSetting: d.kill_ratio * 100,
              betCounts: d.bet_count,
            }
          })
        : []

      let totalScore = round(recordsInfo.totalGamerWinScore + recordsInfo.totalBonus, 4)
      recordsInfo.avgRtp = round((totalScore / recordsInfo.totalGamerBet) * 100, 4) || 0
      recordsInfo.avgDeductTaxRtp =
        round(((totalScore - recordsInfo.totalTax) / recordsInfo.totalGamerBet) * 100, 4) || 0

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
    calDirections,
    formInput,
    records,
    recordsInfo,
    tableInput,
    searchRecords,
  }
}
