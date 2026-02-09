import { computed, inject, reactive, ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysGame'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useUserStore } from '@/base/store/userStore'
import { useDropdownListStore } from '@/base/store/dropdownStore'

export function useGameSetting() {
  const { t } = useI18n()

  const confirm = inject('confirm')
  const warn = inject('warn')

  const isSettingEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GameSettingGameStateUpdate)
  })

  const gameState = ref(constant.GameState.All)
  const gameStates = [
    constant.GameState.All,
    constant.GameState.Online,
    constant.GameState.Maintain,
    constant.GameState.Offline,
  ]

  const games = reactive({ items: [] })
  const gameTaleInput = reactive(new BaseTableInput(0, 'id', constant.TableSortDirection.Asc))
  const filterGames = computed(() => {
    if (gameState.value === constant.GameState.All) {
      return games.items
    } else {
      return games.items.filter((g) => g.state === gameState.value)
    }
  })
  const onlineGameCount = computed(() => {
    return games.items.filter((g) => g.state === constant.GameState.Online).length
  })
  const maintainGameCount = computed(() => {
    return games.items.filter((g) => g.state === constant.GameState.Maintain).length
  })

  async function searchGames() {
    try {
      gameTaleInput.showProcessing = true

      const resp = await api.getGameList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      games.items = resp.data.data

      gameTaleInput.length = games.items.length
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      gameTaleInput.showProcessing = false
    }
  }

  function setGameState(game, state) {
    if (game.state === state) {
      return
    }

    const info = {
      0: [t('textUpdateGameStateOffline'), t('textUpdateGameStateReturnLobby')],
      1: [t('textUpdateGameStateOnline')],
      2: [t('textUpdateGameStateMaintain'), t('textUpdateGameStateReturnLobby')],
    }

    const msg = info[state].join('\n')
    const title = t('fmtTextUpdateStateTitle', [t(`game__${game.id}`), t(`state__${state}`)])

    confirm(msg, { title }).then(async () => {
      try {
        gameTaleInput.showProcessing = true

        const resp = await api.setGameState({
          game_id: game.id,
          state: state,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(async () => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          await searchGames()

          const { getGameList } = useDropdownListStore()
          await getGameList()
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
        gameTaleInput.showProcessing = false
      }
    })
  }

  async function notifyGameToGameServer(game) {
    try {
      gameTaleInput.showProcessing = true

      const resp = await api.notifyGameServer({
        game_id: game.id,
      })

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
      gameTaleInput.showProcessing = false
    }
  }

  onMounted(async () => {
    await searchGames()
  })

  return {
    filterGames,
    games,
    gameState,
    gameStates,
    gameTaleInput,
    isSettingEnabled,
    maintainGameCount,
    onlineGameCount,
    notifyGameToGameServer,
    setGameState,
  }
}
