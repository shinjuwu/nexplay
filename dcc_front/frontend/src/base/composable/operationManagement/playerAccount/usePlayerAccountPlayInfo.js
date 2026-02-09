import { watch, ref, inject, reactive, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'

export function usePlayerAccountPlayInfo(props, emit) {
  const { t } = useI18n()
  const warn = inject('warn')

  const visibleGameDialog = ref(false)

  function closeGameDialog() {
    visibleGameDialog.value = false
    emit('close', false)
  }

  const visibleGameRoomDialog = ref(false)

  function closeGameRoomDialog() {
    visibleGameRoomDialog.value = false
  }

  const playerInfo = reactive({
    id: 0,
    userName: '',
    totalPlays: 0,
    gamePlayInfo: {},
  })
  const gamePlayInfo = computed(() => {
    return Object.values(playerInfo.gamePlayInfo).sort((a, b) => a.gameId - b.gameId)
  })
  const gamePlayDetail = reactive({
    items: [],
  })

  async function searchGameUsersPlayInfo() {
    try {
      const resp = await api.getGameUserPlayCountData({
        user_id: props.playerInfo.id,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`)).then(() => emit('close', false))
        return
      }

      const data = resp.data.data
      playerInfo.id = props.playerInfo.id
      playerInfo.userName = props.playerInfo.userName

      let tmpTotalPlays = 0
      const tmpGamePlayInfo = {}
      if (data.data) {
        for (const playInfo of data.data) {
          const gameId = Math.floor(playInfo.room_id / 10)
          const roomType = playInfo.room_id % 10

          if (!tmpGamePlayInfo[gameId]) {
            tmpGamePlayInfo[gameId] = {
              gameId: gameId,
              gameCode: playInfo.game_code,
              newbieLimit: playInfo.newbie_limit,
              totalCount: 0,
              details: [],
            }
          }

          tmpGamePlayInfo[gameId].totalCount += playInfo.play_count
          tmpGamePlayInfo[gameId].details.push({
            roomId: playInfo.room_id,
            roomType: roomType,
            gameId: gameId,
            playCount: playInfo.play_count,
          })
          tmpTotalPlays += playInfo.play_count
        }
      }
      playerInfo.totalPlays = tmpTotalPlays
      playerInfo.gamePlayInfo = tmpGamePlayInfo

      visibleGameDialog.value = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`)).then(() => emit('close', false))
      }
    }
  }

  function showGameUserGamePlayDetail(gameId) {
    gamePlayDetail.items = playerInfo.gamePlayInfo[gameId].details.sort((a, b) => a.roomId - b.roomId)
    visibleGameRoomDialog.value = true
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        await searchGameUsersPlayInfo()
      } else {
        visibleGameDialog.value = false
      }
    }
  )

  return {
    gamePlayInfo,
    gamePlayDetail,
    playerInfo,
    visibleGameDialog,
    visibleGameRoomDialog,
    closeGameDialog,
    closeGameRoomDialog,
    showGameUserGamePlayDetail,
  }
}
