import { computed, inject, reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import { defaultBadgeInfo } from '@/base/common/badge'
import constant from '@/base/common/constant'
import { menuItemKey, roleItemKey } from '@/base/common/menuConstant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useUserStore } from '@/base/store/userStore'
import { getMenuItemFromMenu } from '@/base/utils/menu'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'
import time from '@/base/utils/time'

export function usePlayerAccount(props) {
  const { t } = useI18n()
  const warn = inject('warn')

  const uStore = useUserStore()
  const { user } = storeToRefs(uStore)

  const myAgentId = user.value.agentId
  const myCustomBadges = reactive({ items: [] })

  const isSettingEnabled = computed(() => {
    const { isInRole } = uStore
    return {
      PlayerAccountInfoUpdate: isInRole(roleItemKey.PlayerAccountInfoUpdate),
      PlayerBadgeUpdate: isInRole(roleItemKey.PlayerBadgeUpdate),
      PlayerDisposeSettingUpdate: isInRole(roleItemKey.PlayerDisposeSettingUpdate),
      PlayerLogRead: isInRole(roleItemKey.PlayerLogRead),
    }
  })

  const formInput = reactive({
    agent: {},
    userName: '',
  })
  const tableInput = reactive(
    new BaseTableInput(constant.TableDefaultLength, 'userName', constant.TableSortDirection.Asc)
  )

  const records = reactive({ items: [] })

  async function searchGameUsers() {
    try {
      tableInput.showProcessing = true

      const resp = await api.getGameUsers()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      records.items = resp.data.data
        .filter((d) => {
          if (formInput.agent.id !== constant.Agent.All && formInput.agent.id !== d.agent_id) {
            return false
          }
          if (formInput.userName !== '' && !d.username.includes(formInput.userName)) {
            return false
          }
          return true
        })
        .map((d) => {
          if (d.agent_id == user.value.agentId && myCustomBadges.items.length === 0) {
            myCustomBadges.items = Object.values(d.custom_tag_info)
              .map((c) => {
                const colorInfo = c.color !== '' ? JSON.parse(c.color) : { txt_color: '', bg_color: '' }
                return {
                  idx: c.idx,
                  name: c.name,
                  backgroundColor: colorInfo.bg_color,
                  textColor: colorInfo.txt_color,
                }
              })
              .sort((a, b) => a.idx - b.idx)
          }

          const createTime = new Date(d.create_time)
          const lastLoginTime = new Date(d.last_login_time)

          return {
            id: d.id,
            userName: d.username,
            agentName: d.agent_name,
            state: d.id_enabled,
            agentId: d.agent_id,
            coinIn: d.coin_in,
            coinOut: d.coin_out,
            createTime: createTime,
            createTimeStr: time.localTimeFormat(createTime),
            lastLoginTime: d.last_login_time,
            lastLoginTimeStr: lastLoginTime - createTime >= 0 ? time.localTimeFormat(lastLoginTime) : '-',
            isOnline: d.is_online,
            badges: myCustomBadges.items.map((i) => i.idx).filter((idx) => d.tag_list[idx] !== '0'),
            isHighRisk: d.high_risk,
            killDiveState: d.kill_dive_state,
            killDiveValue: d.kill_dive_state === constant.KillDive.ConfigKill ? d.kill_dive_value : 0,
            badgeInfo: `${d.kill_dive_state === 2 ? 1 : 0}${d.high_risk ? 1 : 0}${d.kill_dive_state === 1 ? 1 : 0}${
              myCustomBadges.items.length > 0 ? d.tag_list : d.tag_list.replace('1', '0')
            }`,
            riskControlTag: setRiskTag(d.risk_control_tag_list),
            riskControlTagList: d.risk_control_tag_list,
            walletType: d.wallet_type,
            walletBalance: 0,
            isSearchWalletBalance: false,
            isNewbie: false,
            isSearchGameUserPlayCount: false,
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

  function setRiskTag(playerTags) {
    let tags = []
    const riskTags = ['1000', '0100', '0010', '0001']
    for (let i = 0; i < playerTags.length; i++) {
      if (playerTags[i] === '1') {
        tags.push(t(`riskControlTag__${riskTags[i]}`))
      }
    }
    return tags
  }

  function redirectToPlayerLog(agentId, userName) {
    let item = getMenuItemFromMenu(menuItemKey.PlayerLog)
    item.props.agentId = agentId
    item.props.userName = userName

    const { addBreadcrumbItem } = useBreadcrumbStore()
    addBreadcrumbItem(item)
  }

  const dialog = reactive({
    detail: false, // 信息
    setting: false, // 設定
    badge: false, // 標示
    dispose: false, // 處置
    playInfo: false, // 遊戲資訊
  })

  const playerInfo = reactive({ id: 0, userName: '', agentName: '' })
  function selectPlayerInfo(targetPlayerInfo, targetDialog) {
    playerInfo.id = targetPlayerInfo.id
    playerInfo.userName = targetPlayerInfo.userName
    playerInfo.agentName = targetPlayerInfo.agentName
    dialog[targetDialog] = true
  }

  watch(
    props,
    async () => {
      if (props.userName !== '' && props.agentId !== -1) {
        formInput.userName = props.userName
        await searchGameUsers()
      }
    },
    { immediate: true }
  )

  const showRealTimeBalanceColumn =
    user.value.accountType === constant.AccountType.Admin || user.value.walletType === constant.AgentWallet.Transfer
  const tableColumns = (() => {
    // 基礎有13個欄位，其餘有多的再自行增加
    let tableColumns = 13

    if (showRealTimeBalanceColumn) {
      tableColumns++
    }

    return tableColumns
  })()

  async function getPlayerWalletBalance(gameUser) {
    try {
      tableInput.showProcessing = true

      const resp = await api.getGameUserBalance({ user_id: gameUser.id })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      gameUser.walletBalance = resp.data.data.wallet_balance
      gameUser.isSearchWalletBalance = true
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

  async function getPlayerPlayCountData(gameUser) {
    try {
      tableInput.showProcessing = true

      const resp = await api.getGameUserPlayCountData({ user_id: gameUser.id })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      gameUser.isNewbie = data.data
        ? data.data.reduce((sum, cur) => sum + cur.play_count, 0) <= data.total_newbie_limit
        : true
      gameUser.isSearchGameUserPlayCount = true
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
    defaultBadgeInfo,
    dialog,
    formInput,
    isSettingEnabled,
    myAgentId,
    myCustomBadges,
    playerInfo,
    records,
    showRealTimeBalanceColumn,
    tableInput,
    tableColumns,
    getPlayerWalletBalance,
    getPlayerPlayCountData,
    searchGameUsers,
    selectPlayerInfo,
    redirectToPlayerLog,
  }
}
