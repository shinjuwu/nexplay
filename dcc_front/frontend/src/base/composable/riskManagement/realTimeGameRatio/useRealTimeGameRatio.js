import { inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'

import * as api from '@/base/api/sysRiskControl'
import constant from '@/base/common/constant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import * as math from '@/base/utils/math'

export function useRealTimeGameRatio() {
  const warn = inject('warn')
  const { t } = useI18n()

  const formInput = reactive({
    game: null,
  })
  const tableInput = reactive(new BaseTableInput(constant.TableDefaultLength, 'rtp', constant.TableSortDirection.Desc))

  const realTimeGameRatios = reactive({ items: [] })

  async function searchRecords() {
    tableInput.showProcessing = true

    try {
      const resp = await api.getRealTimeGameRatio({
        game_id: formInput.game.id,
      })
      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const dataObj = JSON.parse(resp.data.data)
      if (!dataObj) {
        warn(t(`errorCode__${constant.ErrorCode.ErrorNotification}`))
        return
      }

      realTimeGameRatios.items = Object.keys(dataObj).map((roomIdStr) => {
        const data = dataObj[roomIdStr]

        const roomId = parseInt(roomIdStr)
        const gameId = Math.floor(roomId / 10)
        const roomType = roomId % 10

        // 得分有扣除抽水，所以要加回來
        const de = math.round(data.de + data.tax, 4)

        return {
          roomId: roomId,
          gameId: gameId,
          roomType: roomType,
          playerCount: data.player_count,
          playCount: data.play_count,
          yaScore: data.ya,
          deScore: de,
          tax: data.tax,
          bonus: data.bonus,
          rtp: data.ya > 0 ? math.round(de / data.ya, 6) : 0,
          deductTaxRtp: data.ya > 0 ? math.round((de - data.tax + data.bonus) / data.ya, 6) : 0,
        }
      })

      tableInput.length = realTimeGameRatios.items.length
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

  return {
    formInput,
    realTimeGameRatios,
    tableInput,
    searchRecords,
  }
}
