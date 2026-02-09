import { inject, ref, reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import { getAgentCoinsSupplyInfo, setAgentCoinSupplyInfo } from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import { useUserStore } from '@/base/store/userStore'
import regex from '@/base/common/regex'

export function useAgentAccountSetting(props, emit) {
  const { t } = useI18n()

  const warn = inject('warn')
  const confirm = inject('confirm')

  const { user } = storeToRefs(useUserStore())
  const show = ref(false)

  const agent = reactive({
    id: 0,
    agentId: '',
    topAgentId: '',
    balance: '',
    name: '',
    ratio: 0,
    cooperate: 1,
    remark: '',
    role: '',
    roleName: '',
    walletType: 0,
    walletConnInfo: '',
    lobbySwitch: [],
    cannedSwitch: false,
  })
  const agentLobbySwitches = Object.values(constant.AgentLobbySwitch)

  function validateForm() {
    const errors = []

    if (!regex.agentName.test(agent.name)) {
      errors.push(t('textAgentNameErrorMessage'))
    }

    if (agent.walletType === constant.AgentWallet.Single) {
      try {
        new URL(agent.walletConnInfo)
      } catch {
        errors.push(t('textWalletConnInfoErrorMessage'))
      }
    }

    return errors
  }

  function submit() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    confirm(t('textUpdateGroupRoleReminder2')).then(async () => {
      try {
        const resp = await setAgentCoinSupplyInfo({
          id: agent.id,
          commission: agent.ratio,
          info: agent.remark,
          name: agent.name,
          role: agent.role,
          wallet_conninfo: agent.walletConnInfo,
          lobby_switch_info: agent.lobbySwitch.reduce((acc, cur) => acc | cur, 0),
          canned_switch: agent.cannedSwitch,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }
          emit('searchRecords')
          close()
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
    })
  }

  function close() {
    show.value = false
    emit('close', false)
  }

  async function searchAgentCoinsSupplyInfo() {
    try {
      const resp = await getAgentCoinsSupplyInfo({ id: props.selectRecord.id })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          close()
        })
        return
      }

      const data = resp.data.data
      agent.id = props.selectRecord.id
      agent.topAgentId = data.top_agent_id
      agent.balance = (data.coin_limit - data.coin_use).toFixed(2)
      agent.name = data.name
      agent.ratio = data.commission
      agent.cooperate = data.cooperation
      agent.remark = data.info
      agent.role = data.role
      agent.roleName = data.role_name
      agent.walletType = data.wallet_type
      agent.walletConnInfo = data.wallet_conninfo
      agent.lobbySwitch = agentLobbySwitches.filter((v) => (v & data.lobby_switch_info) === v)
      agent.cannedSwitch = data.canned_switch

      show.value = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }

      emit('close', false)
    }
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        await searchAgentCoinsSupplyInfo()
      }
    }
  )

  return {
    user,
    show,
    agent,
    agentLobbySwitches,
    submit,
    close,
  }
}
