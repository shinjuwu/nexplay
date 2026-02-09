import { inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'

import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { menuItemKey, roleItemKey } from '@/base/common/menuConstant'
import { GetUserCreditLogListInput } from '@/base/common/table/getUserCreditLogListInput'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'
import { useUserStore } from '@/base/store/userStore'
import { numberToStr } from '@/base/utils/formatNumber'
import { round } from '@/base/utils/math'
import { getMenuItemFromMenu } from '@/base/utils/menu'
import { columnSorting } from '@/base/utils/table'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'

export function useUserCreditReport() {
  const warn = inject('warn')
  const { t } = useI18n()

  const isAdminUser = (() => {
    const { isAdminUser } = useUserStore()
    return isAdminUser()
  })()

  const isInRoleGameLogParse = (() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GameLogParseRead)
  })()

  const timeRange = time.getCurrentTimeRageByMinutesAndStep(
    time.winloseReportTimeRange <= 30 ? time.winloseReportTimeRange : 30
  )
  const formInput = reactive({
    userName: '',
    agent: {},
    game: {},
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })
  const tableInput = reactive(
    new GetUserCreditLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.game.id,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'creditTime',
      constant.TableSortDirection.Desc
    )
  )

  const records = reactive({
    items: [],
  })

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
    } else if (endTime - startTime > time.winloseReportTimeRange * 60 * 1000) {
      const hour = Math.floor(time.winloseReportTimeRange / 60)
      const minute = time.winloseReportTimeRange % 60

      let timeTextKey = ''
      let timeTextArgs = []
      if (hour > 0 && minute > 0) {
        timeTextKey = 'fmtTextHoursMinutes'
        timeTextArgs.push(hour, minute)
      } else if (hour > 0) {
        timeTextKey = 'fmtTextHours'
        timeTextArgs.push(hour)
      } else {
        timeTextKey = 'fmtTextMinutes'
        timeTextArgs.push(minute)
      }

      errors.push(t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t(timeTextKey, timeTextArgs)]))
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

    const tmpTableInput = new GetUserCreditLogListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.game.id,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getUserCreditLogList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.winloseReportTimeBeforeDays]))
        return
      }

      records.items = resp.data.data.map((d) => {
        let isCreditIdSupportShowSource = false
        let showCreditIdSourceFunc = null

        let reasonI18nKey = ''
        const resultI18nKeys = []

        let beforeScore = '-'
        let changeScore = '-'
        let afterScore = '-'

        let beforeScoreSrc = 0
        let changeScoreSrc = 0
        let afterScoreSrc = 0

        let beforeScoreSort = Number.MIN_SAFE_INTEGER
        let changeScoreSort = Number.MIN_SAFE_INTEGER
        let afterScoreSort = Number.MIN_SAFE_INTEGER

        // JP紀錄
        if (d.jp_item !== -1) {
          changeScoreSrc = d.jp_score
          changeScore = numberToStr(changeScoreSrc)
          changeScoreSort = changeScoreSrc

          reasonI18nKey = `jackpotWinningItemWithJPPrefix__${d.jp_item}`
        }
        // 遊戲紀錄
        else if (d.game_id !== -1) {
          beforeScoreSrc = d.start_score

          const gameId = d.game_id
          const gameType = Math.floor(gameId / 1000)
          if (gameType === constant.GameType.BaiRen) {
            beforeScoreSrc = round(beforeScoreSrc + d.ya_score, 4)
          }

          afterScoreSrc = d.end_score
          changeScoreSrc = round(afterScoreSrc - beforeScoreSrc, 4)

          beforeScore = numberToStr(beforeScoreSrc)
          changeScore = numberToStr(changeScoreSrc)
          afterScore = numberToStr(afterScoreSrc)

          beforeScoreSort = beforeScoreSrc
          changeScoreSort = changeScoreSrc
          afterScoreSort = afterScoreSrc

          reasonI18nKey = `game__${gameId}`

          if (isInRoleGameLogParse) {
            isCreditIdSupportShowSource = true
            showCreditIdSourceFunc = showGameLogParse
          }
        }
        // 上下分紀錄
        else if (d.wallet_kind !== -1) {
          const isSingleWalletKind =
            d.wallet_kind === constant.WalletLedgerKind.SingleWalletUp ||
            d.wallet_kind === constant.WalletLedgerKind.SingleWalletDown
          const changeset =
            d.wallet_changeset !== null && d.wallet_changeset !== '' && d.wallet_changeset !== '{}'
              ? JSON.parse(d.wallet_changeset)
              : { before_coin: 0, add_coin: 0, after_coin: 0 }
          beforeScoreSrc = changeset.before_coin
          changeScoreSrc = changeset.add_coin
          afterScoreSrc = changeset.after_coin

          changeScore = numberToStr(changeScoreSrc)
          changeScoreSort = changeScoreSrc

          if (!isSingleWalletKind) {
            beforeScore = numberToStr(beforeScoreSrc)
            afterScore = numberToStr(afterScoreSrc)

            beforeScoreSort = beforeScoreSrc
            afterScoreSort = afterScoreSrc
          }

          reasonI18nKey = `walletLedgerKind__${d.wallet_kind}`
          resultI18nKeys.push(`orderType__${d.wallet_status}`, `errorCode__${d.wallet_error_code}`)
        }

        return {
          agentId: d.agent_id,
          agentName: d.agent_name,
          userName: d.username,
          creditTime: d.credit_time,
          beforeScore: beforeScore,
          beforeScoreSrc: beforeScoreSrc,
          beforeScoreSort: beforeScoreSort,
          changeScore: changeScore,
          changeScoreSrc: changeScoreSrc,
          changeScoreSort: changeScoreSort,
          afterScore: afterScore,
          afterScoreSrc: afterScoreSrc,
          afterScoreSort: afterScoreSort,
          reasonI18nKey: reasonI18nKey,
          creditId: d.credit_id,
          resultI18nKeys: resultI18nKeys,
          creditCode: d.credit_code,
          isCreditIdSupportShowSource: isCreditIdSupportShowSource,
          showCreditIdSourceFunc: showCreditIdSourceFunc,
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

  async function downloadRecords() {
    const headers = [
      [
        t('textAgentName'),
        t('textPlayerAccount'),
        t('textCreateTime'),
        t('textBeforeScore'),
        t('textChangeScore'),
        t('textAfterScore'),
        t('textScoreChangeReason'),
        `${t('textRoundId')}/${t('textBetId')}`,
        t('textOperateResult'),
        `${t('textSingleWalletId')}/${t('textJPTokenId')}`,
      ],
    ]

    const dir = tableInput.dir
    const column = tableInput.column
    const rows = records.items
      .sort((a, b) => columnSorting(dir, a[column], b[column]))
      .map((d) => ({
        agentName: d.agentName,
        userName: d.userName,
        creditTime: time.utcTimeStrToLocalTimeFormat(d.creditTime),
        beforeScore: d.beforeScore,
        changeScore: d.changeScore,
        afterScore: d.afterScore,
        reason: t(d.reasonI18nKey),
        creditId: d.creditId,
        result: d.resultI18nKeys.length > 0 ? d.resultI18nKeys.map((k) => t(k)).join('/') : '-',
        creditCode: d.creditCode,
      }))

    xlsx.createAndDownloadFile(
      `[UcReport]${time.localTimeFormat(tableInput.startTime)}-${time.localTimeFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  function showGameLogParse(record) {
    let item = getMenuItemFromMenu(menuItemKey.GameLogParse)
    item.props.logNumber = record.creditId
    item.props.userName = record.userName

    const { addBreadcrumbItem } = useBreadcrumbStore()
    addBreadcrumbItem(item)
  }

  return {
    isAdminUser,
    formInput,
    tableInput,
    records,
    searchRecords,
    downloadRecords,
  }
}
