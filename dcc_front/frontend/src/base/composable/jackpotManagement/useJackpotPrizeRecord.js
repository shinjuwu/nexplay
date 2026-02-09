import { reactive, computed, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import { GetJackpotListInput } from '@/base/common/table/getJackpotListInput'
import constant from '@/base/common/constant'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'
import { useUserStore } from '@/base/store/userStore'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import * as api from '@/base/api/sysJackpot'
import { round } from '@/base/utils/math'
import { numberToStr } from '@/base/utils/formatNumber'

export function useJackpotPrizeRecord() {
  const { t } = useI18n()
  const warn = inject('warn')

  const { user } = storeToRefs(useUserStore())
  const isAdmin = computed(() => user.value.accountType === constant.AccountType.Admin)

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()
  const formInput = reactive({
    agent: { id: constant.Agent.All },
    tokenId: '',
    roundId: '',
    userName: '',
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })
  const tableInput = reactive(
    new GetJackpotListInput(
      formInput.agent.id,
      formInput.startTime,
      formInput.endTime,
      formInput.roundId,
      formInput.tokenId,
      formInput.userName,
      constant.TableDefaultLength,
      'winningTime',
      constant.TableSortDirection.Desc
    )
  )

  const jackpotPrizeList = reactive({
    items: [],
  })
  const totalWinningScore = computed(() =>
    jackpotPrizeList.items.reduce((sum, record) => round(sum + record.winningScore, 4), 0)
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

  async function getJackpotPrizeRecords() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      tableInput.showProcessing = true

      const tmpTableInput = new GetJackpotListInput(
        formInput.agent.id,
        formInput.startTime,
        formInput.endTime,
        formInput.roundId,
        formInput.tokenId,
        formInput.userName,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )

      const resp = await api.getJackpotList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      jackpotPrizeList.items = data.map((d) => {
        return {
          orderId: d.bet_id,
          roundId: d.lognumber,
          winningTime: d.winning_time,
          agentName: d.agent_name,
          userName: d.username,
          tokenId: d.token_id,
          tokenGetTime: d.token_create_time,
          winningItem: d.prize_item,
          winningScore: d.prize_score,
          isBOT: d.is_robot,
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
        t('textRoundId'),
        t('textWinningTime'),
        t('textAgent'),
        t('textPlayerAccount'),
        t('textJPTokenId'),
        t('textJPTokenGetTime'),
        t('textWinningItem'),
        t('textWinningScore'),
        'BOT',
      ],
    ]
    const rows = jackpotPrizeList.items.map((d) => {
      return {
        orderId: d.orderId,
        roundId: d.roundId,
        winningTime: time.utcTimeStrToLocalTimeFormat(d.winningTime),
        agentName: d.isBOT ? '-' : d.agentName,
        userName: d.userName,
        tokenId: d.tokenId,
        tokenGetTime: time.utcTimeStrToLocalTimeFormat(d.tokenGetTime),
        winningItem: d.prizeItem,
        winningScore: numberToStr(d.winningScore),
        isBOT: d.isBOT === 1 ? t('textYes') : t('textNo'),
      }
    })

    xlsx.createAndDownloadFile(
      `[JackpotPrizeReport]${time.recordTimeFormat(tableInput.startTime)}-${time.recordTimeFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  return {
    tableInput,
    formInput,
    jackpotPrizeList,
    isAdmin,
    totalWinningScore,
    getJackpotPrizeRecords,
    downloadRecords,
  }
}
