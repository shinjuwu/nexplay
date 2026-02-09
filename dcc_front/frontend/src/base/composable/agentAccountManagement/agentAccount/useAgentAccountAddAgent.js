import { reactive, inject, computed, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import * as api from '@/base/api/sysAgent'
import axios from 'axios'
import constant from '@/base/common/constant'
import regex from '@/base/common/regex'
import { useUserStore } from '@/base/store/userStore'

export function useAgentAccountAddAgent(emit) {
  const warn = inject('warn')
  const { t } = useI18n()

  const { user } = storeToRefs(useUserStore())
  const accountType = computed(() => {
    return user.value.accountType + 1
  })
  const header = computed(() =>
    user.value.accountType > 1 ? t('textAgentSettingAddSubAgent') : t('textAgentSettingAddGeneralAgent')
  )
  const isNotAdmin = computed(() => user.value.accountType > 1)

  const newAgentForm = reactive({
    account: '',
    password: '',
    agentName: '',
    ratio: 0,
    allowList: '',
    cooperate: user.value.accountType > 1 ? user.value.cooperation : constant.AgentCooperation.BuyPoint,
    remark: '',
    role: '',
    currency: user.value.accountType > 1 ? user.value.currency : constant.CurrencyType.CNY,
    walletType: user.value.accountType > 1 ? user.value.walletType : constant.AgentWallet.Transfer,
    walletConnInfo: '',
    lobbySwitch: [constant.AgentLobbySwitch.Normal],
    cannedSwitch: false,
  })
  watch(
    () => newAgentForm.walletType,
    (newWalletType) => {
      if (newWalletType !== constant.AgentWallet.Single) {
        newAgentForm.walletConnInfo = ''
      }
    }
  )

  const cooperateOptions = [constant.AgentCooperation.BuyPoint, constant.AgentCooperation.Trust]
  const currencyOptions = Object.values(constant.CurrencyType)
  const agentWalletOptions = Object.values(constant.AgentWallet)
  const agentLobbySwitches = Object.values(constant.AgentLobbySwitch)

  async function createNewAgent() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join('\n'))
      return
    }

    try {
      const resp = await api.createAgent({
        account: newAgentForm.account,
        password: newAgentForm.password,
        commission: newAgentForm.ratio,
        cooperation: newAgentForm.cooperate,
        info: newAgentForm.remark,
        ip_whitelist: newAgentForm.allowList,
        nickname: newAgentForm.agentName,
        role: newAgentForm.role,
        currency: newAgentForm.currency,
        wallet_type: newAgentForm.walletType,
        wallet_url: newAgentForm.walletConnInfo,
        lobby_switch_info: newAgentForm.lobbySwitch.reduce((acc, cur) => acc | cur, 0),
        canned_switch: newAgentForm.cannedSwitch,
      })
      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        emit('searchRecords')
        close()

        newAgentForm.account = ''
        newAgentForm.password = ''
        newAgentForm.agentName = ''
        newAgentForm.cooperate = cooperateOptions[0]
        newAgentForm.ratio = 0
        newAgentForm.remark = ''
        newAgentForm.allowList = ''
        newAgentForm.currency = currencyOptions[0]
        newAgentForm.walletType = user.value.accountType > 1 ? user.value.walletType : constant.AgentWallet.Transfer
        newAgentForm.walletConnInfo = ''
        newAgentForm.lobbySwitch = [constant.AgentLobbySwitch.Normal]
        newAgentForm.cannedSwitch = false
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

  function close() {
    emit('close', false)
  }

  function validateForm() {
    const errors = []

    if (!regex.agentName.test(newAgentForm.agentName)) {
      errors.push(t('textAgentNameErrorMessage'))
    }

    if (!(regex.account.test(newAgentForm.account) && regex.password.test(newAgentForm.password))) {
      errors.push(t('textAccountErrorMessage'))
    }
    if (!newAgentForm.allowList) {
      errors.push(t('textAllowListIPRequired'))
    }

    if (!newAgentForm.role) {
      errors.push(t('textAgentRoleErrorMessage'))
    }

    if (newAgentForm.walletType === constant.AgentWallet.Single) {
      try {
        new URL(newAgentForm.walletConnInfo)
      } catch {
        errors.push(t('textWalletConnInfoErrorMessage'))
      }
    }

    return errors
  }

  return {
    accountType,
    agentWalletOptions,
    agentLobbySwitches,
    header,
    isNotAdmin,
    newAgentForm,
    cooperateOptions,
    currencyOptions,
    createNewAgent,
    close,
  }
}
