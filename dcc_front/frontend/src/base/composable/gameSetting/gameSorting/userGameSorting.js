import { computed, inject, reactive, ref, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysGame'
import constant from '@/base/common/constant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'

export function useGameSorting() {
  const { t } = useI18n()
  const warn = inject('warn')
  const uStore = useUserStore()

  const reminders = [
    t('textGameSortingReminder1'),
    t('textGameSortingReminder2'),
    t('textGameSortingReminder3'),
    t('textGameSortingReminder4'),
    t('textGameSortingReminder5'),
  ]

  const isEditEnabled = computed(() => {
    const { isInRole } = uStore
    return isInRole(roleItemKey.GameSortingUpdate)
  })

  const isAdmin = ref(null)

  const tableInput = reactive(new BaseTableInput(constant.TableDefaultLength, 'id', constant.TableSortDirection.Asc))

  const gameIconList = reactive({ isDefault: null, items: [] })
  const normalGameIconList = computed(() =>
    gameIconList.items.filter(
      (gameIcon) => gameIcon.type !== constant.GameType.Lobby && gameIcon.type !== constant.GameType.FriendsRoom
    )
  )
  const friendRoomGameIconList = computed(() =>
    gameIconList.items.filter((gameIcon) => gameIcon.type === constant.GameType.FriendsRoom)
  )
  let defaultGameIconList = []

  function changeIconLabel(game, label) {
    gameIconList.isDefault = false
    game[label] = game[label] === 0 ? 1 : 0

    if (label === 'isHot' && game[label] === 1) {
      game.isNewest = 0
      return
    }

    if (label === 'isNewest' && game[label] === 1) {
      game.isHot = 0
      return
    }
  }

  async function searchGameIconList() {
    tableInput.showProcessing = true
    try {
      const resp = await api.getGameIconList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      isAdmin.value = resp.data.data.is_admin
      gameIconList.isDefault = resp.data.data.is_default

      const transGameIconList = (gameIconList) =>
        gameIconList.map((gameIcon) => {
          return {
            gameId: gameIcon.id,
            name: gameIcon.name,
            gameCode: gameIcon.code,
            type: gameIcon.type,
            isHot: gameIcon.hot,
            isNewest: gameIcon.newest,
            rank: gameIcon.rank,
            push: gameIcon.push,
          }
        })

      defaultGameIconList = transGameIconList(resp.data.data.default_gameicon_list_data)

      const gameIconListData = gameIconList.isDefault
        ? defaultGameIconList
        : transGameIconList(resp.data.data.gameicon_list_data)
      gameIconList.items = gameIconListData.sort((a, b) => {
        if ((a.push > 0 && b.push > 0) || (a.push === 0 && b.push === 0)) {
          const diff = a.rank - b.rank
          return diff !== 0 ? diff : a.gameId - b.gameId
        } else {
          return a.push > 0 ? -1 : 1
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

  function validateForm() {
    const errors = []

    for (const gameIcon of gameIconList.items) {
      if (gameIcon.rank < 0 || gameIcon.rank > 9999) {
        errors.push(t('textGameIconRankErrorMessage'))
      }
    }

    return errors
  }

  async function setGameIconList() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const sourceGameIconList = gameIconList.isDefault ? defaultGameIconList : gameIconList.items
      const newGameIconList = sourceGameIconList.map((gameIcon) => {
        return {
          gameId: gameIcon.gameId,
          hot: gameIcon.isHot,
          newest: gameIcon.isNewest,
          rank: gameIcon.rank,
          push: gameIcon.push,
        }
      })

      const resp = await api.setGameIconList({
        game_icon_list: newGameIconList,
        is_default: isAdmin.value ? false : gameIconList.isDefault,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        searchGameIconList()
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

  onBeforeMount(async () => {
    await searchGameIconList()
  })

  function getImageUrl(gameId, push) {
    const padPush = `${push}`.padStart(2, '0')
    const url = new URL(`../../../assets/images/promote/icon_${gameId}_${padPush}.png`, import.meta.url)
    return url.href
  }

  return {
    reminders,
    friendRoomGameIconList,
    gameIconList,
    normalGameIconList,
    tableInput,
    isAdmin,
    isEditEnabled,
    changeIconLabel,
    getImageUrl,
    setGameIconList,
  }
}
