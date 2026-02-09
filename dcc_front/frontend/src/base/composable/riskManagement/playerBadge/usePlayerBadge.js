import { reactive, inject, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'
import * as api from '@/base/api/sysRiskControl'
import { defaultBadgeInfo } from '@/base/common/badge'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { GetGameUsersCustomTagInput } from '@/base/common/table/getGameUsersCustomTagInput'
import { useUserStore } from '@/base/store/userStore'
import time from '@/base/utils/time'
import { round } from '@/base/utils/math'

export function usePlayerBadge() {
  const { t } = useI18n()
  const warn = inject('warn')

  const calDirections = [t('textWinScoreDirection'), t('textRTPDirection'), t('textWinRateDirection')]

  const isAdminUser = computed(() => {
    const { isAdminUser } = useUserStore()
    return isAdminUser()
  })
  const myAgentId = computed(() => {
    const { user } = storeToRefs(useUserStore())
    return user.value.agentId
  })

  const timeRange = time.getCurrentTimeRageByMinutesAndStep(60, 60)

  const isEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return {
      PlayerBadgeUpdate: isInRole(roleItemKey.PlayerBadgeUpdate),
      PlayerAccountInfoUpdate: isInRole(roleItemKey.PlayerAccountInfoUpdate),
    }
  })

  const dialog = reactive({
    badge: false,
    setting: false,
  })
  const formInput = reactive({
    agent: { id: constant.Agent.All },
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
    tag: { index: 0, name: '' },
    winScore: null,
    rtp: null,
    winRate: null,
  })

  const tableInput = reactive(
    new GetGameUsersCustomTagInput(
      formInput.agent.id,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'tagList',
      constant.TableSortDirection.Desc
    )
  )

  const agentsCustomTagInfo = reactive({})

  const records = reactive({ items: [] })

  const playerInfo = reactive({ id: 0, userName: '' })
  function selectPlayerInfo(info, targetDialog) {
    playerInfo.id = info.id
    playerInfo.userName = info.userName
    dialog[targetDialog] = true
  }

  function validateForm() {
    let errors = []

    const startTime = formInput.startTime
    const endTime = formInput.endTime
    if (endTime - startTime < 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
    } else if (endTime - startTime > time.commonReportTimeRange * 24 * 60 * 60 * 1000) {
      errors.push(
        t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t('fmtTextDays', [time.commonReportTimeRange])])
      )
    }
    if (formInput.winScore && (formInput.winScore < 0 || formInput.winScore > 999999999)) {
      errors.push(t('textWinScoreRequired'))
    }
    if (formInput.rtp && (formInput.rtp < 0 || formInput.rtp > 100)) {
      errors.push(t('textRTPRequired'))
    }
    if (formInput.winRate && (formInput.winRate < 0 || formInput.winRate > 100)) {
      errors.push(t('textWinRateRequired'))
    }

    return errors
  }

  async function getGameUserTagList() {
    const errors = validateForm()

    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    tableInput.showProcessing = true
    const tmpTableInput = new GetGameUsersCustomTagInput(
      formInput.agent.id,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getGameUsersCustomTagList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const tmpRecords = []
      const tagPosition = formInput.tag.index < 0 ? formInput.tag.index + 3 : formInput.tag.index
      Object.entries(resp.data.data).forEach(([agentId, detail]) => {
        const customBadges = [
          new Badges('', '', '', 0),
          new Badges('', '', '', 1),
          new Badges('', '', '', 2),
          new Badges('', '', '', 3),
          new Badges('', '', '', 4),
          new Badges('', '', '', 5),
          new Badges('', '', '', 6),
          new Badges('', '', '', 7),
        ]

        for (const data of detail.data_list) {
          const winScore = round(data.de + data.bonus, 4)
          tmpRecords.push({
            agentId: data.agent_id,
            id: data.game_users_id,
            agentName: data.agent_name,
            userName: data.game_users_name,
            state: data.is_enabled,
            winScore: data.de,
            bonus: data.bonus,
            rtp: data.ya > 0 ? (winScore / data.ya) * 100 : 0,
            winRate: (data.win_count / data.play_count) * 100,
            highRisk: data.high_risk,
            killDiveState: data.kill_dive_state,
            killDiveValue: data.kill_dive_value,
            tagList: `${data.kill_dive_state === 2 ? '1' : '0'}${data.high_risk ? '1' : '0'}${
              data.kill_dive_state === 1 ? '1' : '0'
            }${data.tag_list}`,
          })
        }

        Object.values(detail.custom_tag_info).forEach((tagInfo) => {
          const customBadge = customBadges[tagInfo.idx]
          customBadge.name = tagInfo.name
          customBadge.info = tagInfo.info

          if (tagInfo.name !== '') {
            const colorObj = JSON.parse(tagInfo.color)
            customBadge.bgColor = colorObj.bg_color
            customBadge.txtColor = colorObj.txt_color
          }
        })
        agentsCustomTagInfo[agentId] = customBadges
      })

      records.items = tmpRecords.filter((r) => {
        const tagListArr = [...r.tagList]

        if (
          (formInput.winScore && r.winScore < formInput.winScore) ||
          (formInput.rtp && r.rtp < formInput.rtp) ||
          (formInput.winRate && r.winRate < formInput.winRate) ||
          (formInput.tag.index !== 101 && tagListArr[tagPosition] === '0')
        ) {
          return false
        }

        return true
      })

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

  return {
    agentsCustomTagInfo,
    calDirections,
    defaultBadgeInfo,
    dialog,
    formInput,
    isAdminUser,
    isEnabled,
    myAgentId,
    playerInfo,
    records,
    tableInput,
    getGameUserTagList,
    selectPlayerInfo,
  }
}

class Badges {
  constructor(name, bgColor, txtColor, index) {
    this.name = name
    this.bgColor = bgColor
    this.txtColor = txtColor
    this.index = index
  }
}
