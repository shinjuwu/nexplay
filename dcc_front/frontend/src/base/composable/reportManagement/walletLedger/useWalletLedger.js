import { inject, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { GetWalletLedgerListInput } from '@/base/common/table/getWalletLedgerListInput'
import { useUserStore } from '@/base/store/userStore'
import { numberToStr } from '@/base/utils/formatNumber'
import { round } from '@/base/utils/math'
import { columnSorting } from '@/base/utils/table'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'

export function useWalletLedger() {
  const { t } = useI18n()
  const warn = inject('warn')

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()

  const uStore = useUserStore()
  const { isAdminUser, isAgentSingleWalletUser, isAgentTransferUser } = uStore

  const isAdmin = isAdminUser()
  const kinds = (() => {
    const k = []

    k.push(constant.WalletLedgerKind.All)

    if (isAdmin || isAgentTransferUser()) {
      k.push(
        constant.WalletLedgerKind.ApiUp,
        constant.WalletLedgerKind.ApiDown,
        constant.WalletLedgerKind.BackendUp,
        constant.WalletLedgerKind.BackendDown
      )
    }

    if (isAdmin || isAgentSingleWalletUser()) {
      k.push(constant.WalletLedgerKind.SingleWalletUp, constant.WalletLedgerKind.SingleWalletDown)
    }

    return k
  })()
  const isShowSingleWalletRelative = isAdmin || isAgentSingleWalletUser()

  const records = reactive({ items: [] })
  const totalUpScore = ref(0)
  const totalDownScore = ref(0)

  const formInput = reactive({
    id: '',
    agent: { id: constant.Agent.All },
    userName: '',
    kind: kinds[0],
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
    singleWalletId: '',
  })
  const tableInput = reactive(
    new GetWalletLedgerListInput(
      formInput.id,
      formInput.agent.id,
      formInput.userName,
      formInput.kind,
      formInput.startTime,
      formInput.endTime,
      formInput.singleWalletId,
      constant.TableDefaultLength,
      'createTime',
      constant.TableSortDirection.Desc
    )
  )

  function validateForm() {
    let errors = []

    const startTime = formInput.startTime
    const endTime = formInput.endTime

    if (endTime - startTime < 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
      return errors
    }

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

  async function confirm(record) {
    if (record.status === 1) return

    try {
      const resp = await api.confirmWalletLedger({
        id: record.betId,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        record.checked = true
        searchRecords()
      })
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    }
  }

  async function searchRecords() {
    totalUpScore.value = 0
    totalDownScore.value = 0

    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    tableInput.showProcessing = true

    try {
      const tmpTableInput = new GetWalletLedgerListInput(
        formInput.id,
        formInput.agent.id,
        formInput.userName,
        formInput.kind,
        formInput.startTime,
        formInput.endTime,
        formInput.singleWalletId,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )

      const resp = await api.getWalletLedgerList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      const pageData = resp.data.data

      records.items = pageData.map((r) => {
        const isSingleWalletKind =
          r.kind === constant.WalletLedgerKind.SingleWalletUp || r.kind === constant.WalletLedgerKind.SingleWalletDown

        const changeset =
          r.changeset !== null && r.changeset !== '' && r.changeset !== '{}'
            ? JSON.parse(r.changeset)
            : { before_coin: 0, add_coin: 0, after_coin: 0 }

        if (r.error_code === constant.ErrorCode.Success) {
          if (changeset.add_coin > 0) {
            totalUpScore.value += round(changeset.add_coin, 4)
          } else {
            totalDownScore.value += round(changeset.add_coin, 4)
          }
        }

        return {
          betId: r.id,
          agentName: r.agent_name,
          account: r.username,
          createTime: r.create_time,
          kind: r.kind,
          coinAmount: isSingleWalletKind ? '-' : numberToStr(r.coin_amount),
          coinAmountSrc: r.coin_amount,
          coinAmountSort: isSingleWalletKind ? Number.MIN_SAFE_INTEGER : r.coin_amount,
          beforeScore: isSingleWalletKind ? '-' : numberToStr(changeset.before_coin),
          beforeScoreSrc: changeset.before_coin,
          beforeScoreSort: isSingleWalletKind ? Number.MIN_SAFE_INTEGER : changeset.before_coin,
          changeScore: changeset.add_coin,
          afterScore: isSingleWalletKind ? '-' : numberToStr(changeset.after_coin),
          afterScoreSrc: changeset.after_coin,
          afterScoreSort: isSingleWalletKind ? Number.MIN_SAFE_INTEGER : changeset.after_coin,
          operator: r.creator,
          status: r.status,
          errorCode: r.error_code,
          remark: r.info,
          statusTips: r.status === 1 ? t('textOrderSuccess') : t('textOrderUnsuccess'),
          singleWalletId: r.single_wallet_id,
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
        t('textPlayerAccount'),
        t('textOrderTime'),
        t('textKind'),
        t('textCoinAmount'),
        t('textBeforeScore'),
        t('textChangeScore'),
        t('textAfterScore'),
        t('textOperator'),
        t('textOrderState'),
        t('textErrorMessage'),
        t('textRemark'),
      ],
    ]

    if (isShowSingleWalletRelative) {
      headers[0].push(t('textSingleWalletId'))
    }

    const dir = tableInput.dir
    const column = tableInput.column
    const rows = records.items
      .sort((a, b) => columnSorting(dir, a[column], b[column]))
      .map((d) => {
        let ret = {
          betId: d.betId,
          agentName: d.agentName,
          account: d.account,
          createTime: time.utcTimeStrToLocalTimeFormat(d.createTime),
          kind: t(`walletLedgerKind__${d.kind}`),
          coinAmount: d.coinAmount,
          beforeScore: d.beforeScore,
          changeScore: d.changeScore,
          afterScore: d.afterScore,
          operator: d.operator,
          status: t(`orderType__${d.status}`),
          errorCode: t(`errorCode__${d.errorCode}`),
          remark: d.remark,
        }

        if (isShowSingleWalletRelative) {
          ret.singleWalletId = d.singleWalletId
        }

        return ret
      })

    xlsx.createAndDownloadFile(
      `[WlReport]${time.localTimeFormat(tableInput.startTime)}-${time.localTimeFormat(tableInput.endTime)}`,
      headers,
      rows
    )
  }

  return {
    isShowSingleWalletRelative,
    kinds,
    records,
    totalUpScore,
    totalDownScore,
    formInput,
    tableInput,
    searchRecords,
    confirm,
    downloadRecords,
  }
}
