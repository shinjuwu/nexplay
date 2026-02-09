import { inject, reactive, ref, watch, computed, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useUserStore } from '@/base/store/userStore'
import { roleItemKey } from '@/base/common/menuConstant'
import { add, endOfMonth, endOfYear, startOfMonth, isSameDay, isBefore, isAfter } from 'date-fns'
import axios from 'axios'
import * as api from '@/base/api/sysJackpot'

export function useJackpotSetting() {
  const { t } = useI18n()

  const pageDirections = [
    t('textJackpotSettingDirections__1'),
    t('textJackpotSettingDirections__2'),
    t('textJackpotSettingDirections__3'),
    t('textJackpotSettingDirections__4'),
    t('textJackpotSettingDirections__5'),
  ]
  const confirm = inject('confirm')
  const warn = inject('warn')

  const isSettingEditEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.JackpotSettingUpdate)
  })
  const isShowDateTimeInput = computed(() => {
    const defaultTime = new Date('1970-01-01T00:00:00Z')
    return !isSameDay(selectedAgent.startTime, defaultTime) && !isSameDay(selectedAgent.endTime, defaultTime)
  })

  const isTimeEditable = computed(() => selectedAgent.newStatus === 0 || selectedAgent.newStatus === 1)

  const showJackpotServerProcessing = ref(false)
  const jackpotGameState = ref(null)
  const visible = ref(false)

  const tableInput = reactive(
    new BaseTableInput(constant.TableDefaultLength, 'startDate', constant.TableSortDirection.Asc)
  )
  const formInput = reactive({
    agent: { id: constant.Agent.All },
  })
  const records = reactive({
    items: [],
  })

  const selectedAgent = reactive({
    id: -1,
    agentName: '',
    status: 0,
    newStatus: -1,
    startTime: new Date(),
    endTime: new Date(),
    tmpStartTime: new Date(),
    tmpEndTime: new Date(),
  })

  function updateAgentJackpotSetting() {
    const msg = t('fmtTextUpdateJackpotSetting', [
      selectedAgent.agentName,
      t(`jackpotStatus__${selectedAgent.status}`),
      t(`jackpotStatus__${selectedAgent.newStatus}`),
    ])

    confirm(msg).then(async () => {
      try {
        const resp = await api.setAgentJackpot({
          agent_id: selectedAgent.id,
          jackpot_status: selectedAgent.newStatus,
          jackpot_start_time: selectedAgent.startTime.toISOString(),
          jackpot_end_time: selectedAgent.endTime.toISOString(),
        })

        warn(t(`errorCode__${resp.data.code}`)).then(async () => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          await getAgentJackpotList()
        })
      } catch (err) {
        console.error(err)

        if (axios.isAxiosError(err)) {
          const errorCode = err.response.data
            ? err.response.data.code || constant.ErrorCode.ErrorLocal
            : constant.ErrorCode.ErrorLocal
          warn(t(`errorCode__${errorCode}`))
        }
      } finally {
        visible.value = false
      }
    })
  }

  function setJackpotGameState(state) {
    const isEnabled = Boolean(state)
    if (jackpotGameState.value === isEnabled) {
      return
    }

    const info = {
      0: [t('textUpdateGameStateOffline'), t('textUpdateGameStateReturnLobby')],
      1: [t('textUpdateGameStateOnline')],
    }
    const message = info[state].join('\n')
    const title = t('fmtTextUpdateStateTitle', ['JACKPOT', t(`state__${state}`)])

    confirm(message, { title }).then(async () => {
      try {
        showJackpotServerProcessing.value = true
        const resp = await api.setJackpotSetting({
          jackpot_switch: isEnabled,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          jackpotGameState.value = isEnabled
        })
      } catch (err) {
        console.error(err)

        if (axios.isAxiosError(err)) {
          const errorCode = err.response.data
            ? err.response.data.code || constant.ErrorCode.ErrorLocal
            : constant.ErrorCode.ErrorLocal
          warn(t(`errorCode__${errorCode}`))
        }
      } finally {
        showJackpotServerProcessing.value = false
      }
    })
  }

  async function getAgentJackpotList() {
    try {
      tableInput.showProcessing = true

      const resp = await api.getAgentJackpotList({
        agent_id: formInput.agent.id,
      })

      const data = resp.data.data
      records.items = data.map((d) => {
        return {
          id: d.id,
          agentName: d.name,
          subAgentCount: d.child_agent_count,
          status: d.jackpot_status,
          startTime: d.jackpot_start_time,
          endTime: d.jackpot_end_time,
        }
      })
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

  async function notifyGameToGameServer() {
    try {
      tableInput.showProcessing = true

      const resp = await api.notifyGameServerAgentJackpotInfo()

      warn(t(`errorCode__${resp.data.code}`))
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

  async function getJackpotGameState() {
    try {
      showJackpotServerProcessing.value = true
      const resp = await api.getJackpotSetting()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      jackpotGameState.value = resp.data.data.jackpot_switch
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showJackpotServerProcessing.value = false
    }
  }

  async function getAgentJackpotSetting(agent) {
    try {
      const resp = await api.getAgentJackpot({
        agent_id: agent.id,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      selectedAgent.id = data.id
      selectedAgent.agentName = data.name
      selectedAgent.startTime = new Date(data.jackpot_start_time)
      selectedAgent.endTime = new Date(data.jackpot_end_time)
      selectedAgent.status = data.jackpot_status
      selectedAgent.newStatus = data.jackpot_status
      selectedAgent.tmpStartTime = new Date(data.jackpot_start_time)
      selectedAgent.tmpEndTime = new Date(data.jackpot_end_time)

      visible.value = true
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

  watch(
    () => selectedAgent.newStatus,
    () => {
      // 還原
      if (selectedAgent.status === selectedAgent.newStatus) {
        selectedAgent.startTime = selectedAgent.tmpStartTime
        selectedAgent.endTime = selectedAgent.tmpEndTime
        return
      }

      const now = new Date()
      const defaultTime = new Date('1970-01-01T00:00:00.000Z')
      const nextMonthFirstDate = add(startOfMonth(now), { months: 1 })
      const maxTime = endOfYear(new Date(2099, 11, 31))
      const isJackpotActive = !isBefore(now, selectedAgent.startTime) && !isAfter(now, selectedAgent.endTime)

      if (selectedAgent.newStatus === 1) {
        selectedAgent.startTime = isJackpotActive ? selectedAgent.tmpStartTime : nextMonthFirstDate
        selectedAgent.endTime = maxTime
      } else if (selectedAgent.newStatus === 2) {
        selectedAgent.startTime = isJackpotActive ? selectedAgent.tmpStartTime : nextMonthFirstDate
        selectedAgent.endTime = isJackpotActive ? selectedAgent.tmpEndTime : maxTime
      } else if (selectedAgent.newStatus === 0) {
        selectedAgent.startTime = isJackpotActive ? selectedAgent.tmpStartTime : defaultTime
        selectedAgent.endTime = isJackpotActive ? endOfMonth(selectedAgent.startTime) : defaultTime
      }
    }
  )

  onBeforeMount(async () => {
    await getJackpotGameState()
  })

  return {
    tableInput,
    formInput,
    records,
    visible,
    selectedAgent,
    pageDirections,
    jackpotGameState,
    isShowDateTimeInput,
    isTimeEditable,
    isSettingEditEnabled,
    showJackpotServerProcessing,
    getAgentJackpotList,
    getAgentJackpotSetting,
    updateAgentJackpotSetting,
    setJackpotGameState,
    notifyGameToGameServer,
  }
}
