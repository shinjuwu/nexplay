import { inject, reactive, ref, computed, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { GetBackendGameUserWalletLedgerListInput } from '@/base/common/table/getBackendGameUserWalletLedgerList'
import { useUserStore } from '@/base/store/userStore'
import { round, roundDown } from '@/base/utils/math'

export function useBackendGameUserWalletLedger() {
  const { t } = useI18n()
  const uStore = useUserStore()
  const { user } = storeToRefs(uStore)

  const warn = inject('warn')

  const isSettingEnabled = computed(() => {
    const { isInRole } = uStore
    return isInRole(roleItemKey.BackendUpdateGameUserWalletUpdate)
  })

  const formInput = reactive({
    username: '',
  })
  const tableInput = reactive(
    new GetBackendGameUserWalletLedgerListInput(formInput.username, 'account', constant.TableDefaultLength)
  )

  const records = reactive({ items: [] })
  const gameUserInfo = reactive({})

  const agentBalance = ref(0)
  const scoreMode = ref('')
  const showEditScoreDialog = ref(false)

  const dialogContent = computed(() => {
    let title, amount, afterAmount, btn

    if (scoreMode.value === 'up') {
      title = t('textBackendGameUserAddScore')
      amount = t('textAddScoreBalance')
      afterAmount = t('textAfterAddScoreBalance')
      btn = t('textAddScore')
    } else {
      title = t('textBackendGameUserDeductScore')
      amount = t('textDeductScoreBalance')
      afterAmount = t('textAfterDeductScoreBalance')
      btn = t('textDeductScore')
    }
    return { title, amount, afterAmount, btn }
  })

  const isAdmin = computed(() => user.value.accountType === constant.AccountType.Admin)

  function editRecordScore(mode, record) {
    scoreMode.value = mode
    Object.assign(gameUserInfo, { ...record, amount: 0, afterAmount: 0, remark: '' })
    showEditScoreDialog.value = true
  }

  function validateForm() {
    const errors = []
    const { amount, afterAmount } = gameUserInfo

    if (isNaN(amount) || amount <= 0) {
      scoreMode.value === 'down'
        ? errors.push(t('textDownScoreNegativeError'))
        : errors.push(t('textUpScoreNegativeError'))
      return errors
    } else if (amount > 999999999999) {
      errors.push(t('textScoreRangeError'))
      return errors
    } else if (amount < 1) {
      errors.push(t('textUpDownScoreThan1'))
      return errors
    }

    if (scoreMode.value === 'down' && afterAmount < 0) {
      errors.push(t('textDownScoreErrorText'))
      return errors
    }

    if (
      !isAdmin.value &&
      user.value.cooperation === constant.AgentCooperation.BuyPoint &&
      scoreMode.value === 'up' &&
      amount > agentBalance.value
    ) {
      errors.push(t(`errorCode__${constant.ErrorCode.ErrorAgentWalletAmountNotEnough}`))
      return errors
    }

    if (!gameUserInfo.remark) {
      errors.push(t('textRemarkErrorText'))
      return errors
    }

    return errors
  }

  async function setGameUserWallet() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const formatAmount = gameUserInfo.amount
      const changeAmount =
        scoreMode.value === 'down' ? roundDown(round(formatAmount, 4)) * -1 : roundDown(round(formatAmount, 4)) * 1

      const resp = await api.setGameUserWallet({
        user_id: gameUserInfo.id,
        change_amount: changeAmount,
        info: gameUserInfo.remark,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        showEditScoreDialog.value = false
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

  async function searchRecords(input) {
    if (!input) {
      input = new GetBackendGameUserWalletLedgerListInput(formInput.username, 'account', tableInput.length)
    }

    tableInput.showProcessing = true
    try {
      const resp = await api.getGameUserWalletList(input.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const pageData = resp.data.data

      if (user.value.accountType !== 1) {
        agentBalance.value = pageData.agent_balance
      }

      if (pageData.draw === 0) {
        input.totalRecords = pageData.recordsTotal
      }

      records.items = pageData.data.map((d) => {
        return {
          id: d.user_id,
          account: d.username,
          agentName: d.agent_name,
          balance: d.gold,
          lockScore: d.lock_gold,
          state: d.is_enabled,
          createTime: d.create_time,
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
  watch(
    () => gameUserInfo.amount,
    () => {
      gameUserInfo.afterAmount =
        scoreMode.value === 'up'
          ? roundDown(round(gameUserInfo.balance + gameUserInfo.amount, 4))
          : roundDown(round(gameUserInfo.balance - gameUserInfo.amount, 4))
    }
  )

  return {
    agentBalance,
    dialogContent,
    formInput,
    gameUserInfo,
    isAdmin,
    isSettingEnabled,
    records,
    showEditScoreDialog,
    tableInput,
    editRecordScore,
    searchRecords,
    setGameUserWallet,
  }
}
