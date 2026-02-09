import { inject, reactive, ref, computed, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { GetBackendAgentWalletLedger } from '@/base/common/table/getBackendAgentWalletLedgerList'
import { useUserStore } from '@/base/store/userStore'
import { round, roundDown } from '@/base/utils/math'

export function useBackendAgentWalletLedger() {
  const { t } = useI18n()
  const uStore = useUserStore()
  const { user } = storeToRefs(uStore)

  const warn = inject('warn')

  const isSettingEnabled = computed(() => {
    const { isInRole } = uStore
    return isInRole(roleItemKey.BackendUpdateAgentWalletUpdate)
  })

  const formInput = reactive({
    agent: { id: constant.Agent.All },
  })
  const tableInput = reactive(
    new GetBackendAgentWalletLedger(
      formInput.agent.id,
      constant.TableDefaultLength,
      'name',
      constant.TableSortDirection.Asc
    )
  )

  const records = reactive({ items: [] })
  const agentInfo = reactive({})

  const agentBalance = ref(0)
  const scoreMode = ref('')
  const showEditScoreDialog = ref(false)
  const errorTexts = reactive({
    amount: '',
    remark: '',
  })

  const dialogContent = computed(() => {
    let title, amount, afterAmount, btn

    if (scoreMode.value === 'up') {
      title = t('textBackendAgentAddScore')
      amount = t('textAddScoreBalance')
      afterAmount = t('textAfterAddScoreBalance')
      btn = t('textAddScore')
    } else {
      title = t('textBackendAgentDeductScore')
      amount = t('textDeductScoreBalance')
      afterAmount = t('textAfterDeductScoreBalance')
      btn = t('textDeductScore')
    }
    return { title, amount, afterAmount, btn }
  })

  const isAdmin = computed(() => user.value.accountType === constant.AccountType.Admin)

  function editRecordScore(mode, record) {
    scoreMode.value = mode
    Object.assign(agentInfo, { ...record, amount: 0, afterAmount: 0, remark: '' })
    showEditScoreDialog.value = true
  }

  function validateForm() {
    const errors = []
    const { amount, afterAmount } = agentInfo

    if (isNaN(amount) || amount <= 0) {
      errors.push(scoreMode.value === 'down' ? t('textDownScoreNegativeError') : t('textUpScoreNegativeError'))
      return errors
    } else if (amount > 999999999999) {
      errors.push(t('textScoreRangeError'))
      return errors
    }

    if (scoreMode.value === 'down' && afterAmount < 0) {
      errors.push(t('textDownScoreErrorText'))
      return errors
    }

    if (!isAdmin.value && scoreMode.value === 'up' && amount > agentBalance.value) {
      errors.push(t('textUpScoreErrorText'))
      return errors
    }

    if (!agentInfo.remark) {
      errors.push(t('textRemarkErrorText'))
      return errors
    }

    return errors
  }

  async function setAgentWallet() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const formatAmount = agentInfo.amount
      const changeAmount =
        scoreMode.value === 'down' ? roundDown(round(formatAmount, 4)) * -1 : roundDown(round(formatAmount, 4)) * 1

      const resp = await api.setAgentWalletList({
        agent_id: agentInfo.id,
        change_amount: changeAmount,
        info: agentInfo.remark,
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

  async function searchRecords() {
    tableInput.showProcessing = true
    try {
      const tmpTableInput = new GetBackendAgentWalletLedger(
        formInput.agent.id,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )

      const resp = await api.getAgentWalletList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      records.items = data
        .map((d) => {
          if (d.id === user.value.agentId && d.level_code.length > 4) agentBalance.value = d.balance
          return {
            id: d.id,
            adminUser: d.admin_user_username,
            name: d.name,
            ratio: d.commission,
            balance: d.balance,
            state: d.is_enabled,
            createTime: d.create_time,
          }
        })
        .filter((r) => r.id !== user.value.agentId)

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

  watch(showEditScoreDialog, () => {
    if (!showEditScoreDialog.value) {
      errorTexts.amount = ''
      errorTexts.remark = ''
    }
  })

  watch(
    () => agentInfo.amount,
    () => {
      agentInfo.afterAmount =
        scoreMode.value === 'up'
          ? roundDown(round(agentInfo.balance + agentInfo.amount, 4))
          : roundDown(round(agentInfo.balance - agentInfo.amount, 4))
    }
  )

  return {
    agentBalance,
    agentInfo,
    dialogContent,
    formInput,
    isAdmin,
    isSettingEnabled,
    records,
    showEditScoreDialog,
    tableInput,
    editRecordScore,
    searchRecords,
    setAgentWallet,
  }
}
