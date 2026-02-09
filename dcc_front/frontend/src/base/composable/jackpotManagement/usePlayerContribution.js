import { reactive, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import { BaseTableInput } from '@/base/common/table/tableInput'
import constant from '@/base/common/constant'
import axios from 'axios'
import * as api from '@/base/api/sysJackpot'

export function usePlayerContribution() {
  const { t } = useI18n()
  const warn = inject('warn')
  const pageDirections = [
    t('textPlayerContributionDirections__1'),
    t('textPlayerContributionDirections__2'),
    t('textPlayerContributionDirections__3'),
    t('textPlayerContributionDirections__4'),
    t('textPlayerContributionDirections__5'),
  ]
  const tableInput = reactive(
    new BaseTableInput(constant.TableDefaultLength, 'ranking', constant.TableSortDirection.Asc)
  )
  const formInput = reactive({
    agent: {},
  })

  const rankingList = reactive({
    items: [],
  })

  async function getJackpotLeaderBoard() {
    try {
      tableInput.showProcessing = true

      const resp = await api.getJackpotLeaderBoard({
        agent_id: formInput.agent.id,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const sortRankingData = resp.data.data.sort((a, b) => b.score - a.score)
      rankingList.items = sortRankingData.map((player, idx) => {
        return {
          ranking: idx + 1,
          agentName: player.agent_name,
          userName: player.username,
          playCount: player.play_num,
          totalBet: player.total_bet,
          loseCount: player.lose,
          winCount: player.win,
          contribute: player.score,
        }
      })

      tableInput.start = 0
      tableInput.draw = 0
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
    tableInput,
    formInput,
    rankingList,
    pageDirections,
    getJackpotLeaderBoard,
  }
}
