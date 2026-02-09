import { inject, reactive, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysCal'
import constant from '@/base/common/constant'
import { GetPerformanceReportInput } from '@/base/common/table/getPerformanceReportInput'
import { round } from '@/base/utils/math'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'

export function useEarningReportDetail(props, emit) {
  const { t } = useI18n()
  const warn = inject('warn')

  const show = ref(false)

  const formInput = reactive({
    agentId: constant.Agent.All,
    startTime: new Date(),
    endTime: new Date(),
  })

  const tableInput = reactive(
    new GetPerformanceReportInput(
      formInput.agentId,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'date',
      constant.TableSortDirection.Asc
    )
  )

  const records = reactive({ items: [] })

  async function searchRecords() {
    tableInput.showProcessing = true

    const tmpTableInput = new GetPerformanceReportInput(
      formInput.agentId,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getPerformanceReport(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          emit('close')
        })
        return
      }

      const data = resp.data.data
      records.items = data.map((d) => {
        const winScore = round(d.sum_de + d.sum_bonus)
        const gamerWinLoseScore = round(winScore - d.sum_ya - d.sum_tax, 4)
        return {
          date: d.log_time,
          agentId: d.agentId,
          agentName: d.agent_name,
          betCount: d.bet_count,
          bet: d.sum_ya, // 總投注
          validBet: d.sum_valid_ya, // 有效投注
          tax: d.sum_tax, // 抽水
          winScore: winScore, // 贏分
          gamerWinLoseScore: gamerWinLoseScore, // 總玩家輸贏
          agentWinLoseScore: round(gamerWinLoseScore * -1, 4), // 代理總輸贏
        }
      })

      Object.assign(tableInput, tmpTableInput)

      show.value = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }

      emit('close')
    } finally {
      tableInput.showProcessing = false
    }
  }

  function downloadRecords() {
    const headers = [
      [
        t('textDate'),
        t('textAgent'),
        t('textBetCounts'),
        t('textBetTotal'),
        t('textValidBetTotal'),
        t('textTotalGamerWinScore'),
        t('textTotalGameTax'),
        t('textTotalGamerWinLoseScore'),
        t('textTotalAgentWinLoseScore'),
      ],
    ]

    const rows = records.items.map((d) => {
      return {
        date: time.utcTimeStrNoneISO8601ToLocalTimeFormat(d.date),
        agentName: d.agentName,
        betCount: d.betCount.toLocaleString(),
        bet: d.bet,
        validBet: d.validBet,
        winScore: d.winScore,
        tax: d.tax,
        gamerWinLoseScore: d.gamerWinLoseScore,
        agentWinLoseScore: d.agentWinLoseScore,
      }
    })

    xlsx.createAndDownloadFile(
      `${time.localTimeFormat(tableInput.startTime)}-${time.localTimeFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  function close() {
    show.value = false
    emit('close')
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        formInput.agentId = props.agentId
        formInput.startTime = props.startTime
        formInput.endTime = props.endTime

        await searchRecords()
      }
    }
  )

  return {
    records,
    show,
    tableInput,
    close,
    downloadRecords,
  }
}
