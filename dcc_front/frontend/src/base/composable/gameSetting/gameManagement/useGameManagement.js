import { computed, inject, reactive, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { GetAgentGameListInput } from '@/base/common/table/getAgentGameListInput'
import { useUserStore } from '@/base/store/userStore'

export function useGameManagement() {
  const { t } = useI18n()

  const confirm = inject('confirm')
  const warn = inject('warn')

  const isGameStateUpdateEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GameManagementGameStateUpdate)
  })
  const isGameRoomStateUpdateEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GameManagementGameRoomStateUpdate)
  })

  const formAgentGameInput = reactive({
    agent: {},
    game: {},
    state: constant.GameState.All,
  })
  const tableAgentGameInput = reactive(
    new GetAgentGameListInput(
      formAgentGameInput.agent.id,
      formAgentGameInput.game.id,
      formAgentGameInput.state,
      constant.TableDefaultLength
    )
  )

  const agentGames = reactive({ items: [] })
  const selectedAgentGames = computed(() =>
    agentGames.items.filter((agentGame) => agentGame.selected && !agentGame.disabled)
  )

  const isAgentGameAllSelected = ref(false)
  watch(isAgentGameAllSelected, (newValue) => {
    const cache = []

    for (const agentGame of agentGames.items) {
      const cacheValue = `${agentGame.gameId}_${agentGame.agentLevelCode}`

      let findTop = false
      for (const data of cache) {
        if (cacheValue.startsWith(data)) {
          findTop = true
          break
        }
      }

      if (!findTop) {
        cache.push(cacheValue)
      }

      agentGame.selected = newValue && !findTop
      agentGame.disabled = newValue && findTop
    }
  })

  function toggleAgentGameCheckbox(targetAgentGame) {
    if (!targetAgentGame || targetAgentGame.disabled) {
      return
    }

    targetAgentGame.selected = !targetAgentGame.selected
    // 如果勾選起來將下級的勾選取消減少送的資料
    for (const agentGame of agentGames.items) {
      if (
        agentGame === targetAgentGame ||
        agentGame.gameId !== targetAgentGame.gameId ||
        !agentGame.agentLevelCode.startsWith(targetAgentGame.agentLevelCode)
      ) {
        continue
      }

      agentGame.selected = false
      agentGame.disabled = targetAgentGame.selected
    }
  }

  async function searchAgentGames(input) {
    if (!input) {
      input = new GetAgentGameListInput(
        formAgentGameInput.agent.id,
        formAgentGameInput.game.id,
        formAgentGameInput.state,
        tableAgentGameInput.length
      )
    }

    try {
      tableAgentGameInput.showProcessing = true

      const resp = await api.getAgentGameList(input.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const pageData = resp.data.data
      if (pageData.draw === 0) {
        input.totalRecords = pageData.recordsTotal
      }

      agentGames.items = pageData.data.map((r) => {
        return {
          agentId: r.agent_id,
          agentName: r.agent_name,
          agentLevelCode: r.agent_level_code,
          gameId: r.game_id,
          gameCode: r.game_code,
          state: r.state,
          selected: false,
          disabled: false,
        }
      })

      Object.assign(tableAgentGameInput, input)

      isAgentGameAllSelected.value = false
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableAgentGameInput.showProcessing = false
    }
  }

  function setAgentGameState(state) {
    if (selectedAgentGames.value.length === 0) {
      return
    }

    const info = {
      0: [t('textUpdateGameStateOffline'), t('textUpdateGameStateReturnLobby')],
      1: [t('textUpdateGameStateOnline')],
      2: [t('textUpdateGameStateMaintain'), t('textUpdateGameStateReturnLobby')],
    }

    const msg = info[state].join('\n')
    const title = t('fmtTextUpdateAgentGameStateTitle', [t(`state__${state}`)])

    confirm(msg, { title }).then(async () => {
      try {
        tableAgentGameInput.showProcessing = true

        const resp = await api.setAgentGameState({
          list: selectedAgentGames.value.map((i) => {
            return {
              agent_id: i.agentId,
              game_id: i.gameId,
            }
          }),
          state: state,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(async () => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          await searchAgentGames(tableAgentGameInput)
          isAgentGameAllSelected.value = false
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
        tableAgentGameInput.showProcessing = false
      }
    })
  }

  const dialogAgentGameRoomInput = reactive({
    agentId: constant.Agent.All,
    agentName: '',
    gameId: constant.Game.All,
    show: false,
    showProcessing: false,
    mode: 'view',
  })

  const agentGameRooms = reactive({ items: [] })

  async function searchAgentGameRooms(agentId, gameId, mode) {
    if (!agentId || agentId <= constant.Agent.All || !gameId || gameId <= constant.Game.All) {
      return
    }

    if (!mode || (mode !== 'view' && mode !== 'edit')) {
      mode = 'view'
    }

    try {
      tableAgentGameInput.showProcessing = true

      const resp = await api.getAgentGameRoomList({
        agent_id: agentId,
        game_id: gameId,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      agentGameRooms.items = resp.data.data
        .map((r) => {
          return {
            agentId: r.agent_id,
            agentName: r.agent_name,
            gameId: r.game_id,
            gameRoomId: r.game_room_id,
            roomType: r.room_type,
            state: r.state,
          }
        })
        .sort((a, b) => a.gameRoomId - b.gameRoomId)

      dialogAgentGameRoomInput.agentId = agentId
      dialogAgentGameRoomInput.gameId = gameId
      dialogAgentGameRoomInput.agentName = agentGameRooms.items.length > 0 ? agentGameRooms.items[0].agentName : ''
      dialogAgentGameRoomInput.mode = mode
      dialogAgentGameRoomInput.show = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableAgentGameInput.showProcessing = false
    }
  }

  async function setAgentGameRoomState() {
    try {
      dialogAgentGameRoomInput.showProcessing = true

      const resp = await api.setAgentGameRoomState({
        agent_id: dialogAgentGameRoomInput.agentId,
        game_id: dialogAgentGameRoomInput.gameId,
        list: agentGameRooms.items.map((i) => {
          return {
            agent_id: i.agentId,
            game_room_id: i.gameRoomId,
            state: i.state,
          }
        }),
      })

      warn(t(`errorCode__${resp.data.code}`)).then(async () => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        await searchAgentGames(tableAgentGameInput)
        dialogAgentGameRoomInput.show = false
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
      dialogAgentGameRoomInput.showProcessing = false
    }
  }

  return {
    agentGameRooms,
    agentGames,
    dialogAgentGameRoomInput,
    formAgentGameInput,
    isAgentGameAllSelected,
    isGameRoomStateUpdateEnabled,
    isGameStateUpdateEnabled,
    selectedAgentGames,
    tableAgentGameInput,
    searchAgentGames,
    searchAgentGameRooms,
    setAgentGameState,
    setAgentGameRoomState,
    toggleAgentGameCheckbox,
  }
}
