import fc from '@/base/common/featureCodes'
import constant from '@/base/common/constant'
import { roleMenuFolders, roleGroups } from '@/base/common/menuConstant'
import { getRoleMenu, roleMenuToRoles } from '@/base/utils/groupRole'
import time from '@/base/utils/time'
import { roomTypeNameIndex } from '@/base/utils/room'

export function compileActionLog(vueI18nT, featureCode, actionLogJson, userStoreToRefs) {
  const compilers = {
    [fc.AgentSetAgentGameState]: compileAgentSetAgentGameStateAtionLog,
    [fc.AgentSetAgentGameRoomState]: compileAgentSetAgentGameRoomStateAtionLog,
    [fc.AgentSetAgentPermission]: compileAgentSetAgentPermission,
    [fc.GameSetGameState]: compileGameSetGameState,
    [fc.UserUpdateGameUserInfo]: compileUserUpdateGameUserInfo,
    [fc.GameSetGameServerState]: compileGameSetGameServerState,
    [fc.GameNotifyGameServer]: compileGameNotifyGameServer,
    [fc.RiskControlSetGameUsersCustomTag]: compileRiskControlSetGameUsersCustomTag,
    [fc.GameSetGameIconList]: compileGameSetGameIconList,
    [fc.AgentSetAgentIpWhitelist]: compileAgentSetAgentIpWhitelist,
    [fc.RiskControlSetIncomeRatio]: compileRiskControlSetIncomeRatio,
    [fc.AgentSetAgentApiIpWhitelist]: compileAgentSetAgentApiIpWhitelist,
    [fc.SystemSetExchangeDataList]: compileSystemSetExchangeDataList,
    [fc.RiskControlSetAutoRiskControlSetting]: compileRiskControlSetAutoRiskControlSetting,
    [fc.RiskControlSetGameUserRiskControlTag]: compileRiskControlSetGameUserRiskControlTag,
    [fc.UserResetPassword]: compileUserResetPassword,
    [fc.JackpotSetAgentJackpot]: compileJackpotSetting,
    [fc.JackpotCreateJackpotToken]: compileJackpotToken,
    [fc.RiskControlSetGameSetting]: compileRiskControlSetGameSetting,
    [fc.RiskControlSetIncomeRatios]: compileRiskControlSetIncomeRatios,
    [fc.AgentSetAgentCoinSupplyInfo]: compileAgentSetAgentCoinSupplyInfo,
    [fc.GameSetCannedList]: compileGameSetCannedList,
  }
  const compiler = compilers[featureCode] || baseCompileActionLog
  return compiler(vueI18nT, featureCode, actionLogJson, userStoreToRefs)
}

function baseCompileActionLog(vueI18nT, featureCode, actionLogJson) {
  return {
    title: baseCompileActionLogTitle(vueI18nT, featureCode, actionLogJson.title),
    detail: baseCompileActionLogDetail(vueI18nT, actionLogJson),
  }
}

function baseCompileActionLogTitle(vueI18nT, featureCode, title) {
  return vueI18nT(`backendActionLogTitle__${featureCode}`, [title])
}

const specialActionLogKey = [
  'title',
  'permissions',
  'username',
  'agent_name',
  'game_id',
  'room_type',
  'lobby_switch_info',
]
function baseCompileActionLogDetail(vueI18nT, actionLogJson) {
  let detail = ''

  for (const key in actionLogJson) {
    if (specialActionLogKey.indexOf(key) >= 0) {
      continue
    }

    let detailKey = key
    let before = actionLogJson[key].before
    let after = actionLogJson[key].after

    switch (key) {
      case 'bool_status':
        before = vueI18nT(`status__${before ? 1 : 0}`)
        after = vueI18nT(`status__${after ? 1 : 0}`)
        detailKey = 'status'
        break
      case 'number_status':
        before = vueI18nT(`status__${before}`)
        after = vueI18nT(`status__${after}`)
        detailKey = 'status'
        break
      case 'level':
        before = vueI18nT(`accountType__${before}`)
        after = vueI18nT(`accountType__${after}`)
        break
      case 'start_time':
      case 'end_time':
        before = time.utcTimeStrToLocalTimeFormat(before)
        after = time.utcTimeStrToLocalTimeFormat(after)
        break
      case 'lang':
        before = vueI18nT(`langType__${before}`)
        after = vueI18nT(`langType__${after}`)
        break
      case 'marquee_type':
        before = vueI18nT(`typeCode__${before}`)
        after = vueI18nT(`typeCode__${after}`)
        detailKey = 'type'
        break
      case 'announcement_type':
        before = vueI18nT(`typeCode__${before}`)
        after = vueI18nT(`typeCode__${after}`)
        detailKey = 'type'
        break
      case 'ratio':
      case 'new_ratio':
        before = parseFloat(before * 100).toFixed(1)
        after = parseFloat(after * 100).toFixed(1)
        break
      case 'wallet_conninfo':
        before = covertUrl(before.scheme, before.domain, before.path)
        after = covertUrl(after.scheme, after.domain, after.path)
        break
    }

    if (!before && isNaN(before)) {
      before = vueI18nT(`textEmpty`)
    }
    if (!after && isNaN(after)) {
      after = vueI18nT(`textEmpty`)
    }

    detail += `${vueI18nT(`backendActionLogDetail_${detailKey}`, [before, after])}\n`
  }

  return detail
}

function covertUrl(scheme, domain, path) {
  return `${scheme}://${domain}${path}`
}

/** 遊戲管理-修改代理遊戲 */
function compileAgentSetAgentGameStateAtionLog(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  let detail = ''
  if (actionLogJson.agent_games) {
    for (const agentGame of actionLogJson.agent_games) {
      const agentName = agentGame.agent_name
      const gameName = vueI18nT(`game__${agentGame.game_id}`)
      const beforeState = vueI18nT(`state__${agentGame.state.before}`)
      const afterState = vueI18nT(`state__${agentGame.state.after}`)
      detail += `${vueI18nT('backendActionLogDetail_agentsetagentgame', [
        agentName,
        gameName,
        beforeState,
        afterState,
      ])}\n`
    }
  }

  return { title, detail }
}

/** 遊戲管理-修改代理遊戲房間 */
function compileAgentSetAgentGameRoomStateAtionLog(vueI18nT, featureCode, actionLogJson) {
  const agentName = actionLogJson.agent_name
  const gameName = vueI18nT(`game__${actionLogJson.game_id}`)
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [agentName, gameName])

  let detail = ''
  for (const agentGameRoom of actionLogJson.agent_game_rooms) {
    const roomType = vueI18nT(`roomType__${roomTypeNameIndex(actionLogJson.game_id, agentGameRoom.room_type)}`)
    const beforeState = vueI18nT(`state__${agentGameRoom.state.before}`)
    const afterState = vueI18nT(`state__${agentGameRoom.state.after}`)
    detail += `${vueI18nT('backendActionLogDetail_agentsetagentgameroom', [roomType, beforeState, afterState])}\n`
  }

  return { title, detail }
}

/** 權限群組管理-修改權限 */
function compileAgentSetAgentPermission(vueI18nT, featureCode, actionLogJson, userStoreToRefs) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [actionLogJson.title])

  let detail = baseCompileActionLogDetail(vueI18nT, actionLogJson)
  if (actionLogJson.permissions) {
    detail += `${vueI18nT('textPermission')}:\n`

    const beforePermissions = actionLogJson.permissions.before
    const afterPermissions = actionLogJson.permissions.after
    const totalPermissions = beforePermissions.concat(afterPermissions)

    const beforeRoles = roleMenuToRoles(
      getRoleMenu(roleMenuFolders, roleGroups, totalPermissions, beforePermissions, userStoreToRefs)
    )
    const afterRoles = roleMenuToRoles(
      getRoleMenu(roleMenuFolders, roleGroups, totalPermissions, afterPermissions, userStoreToRefs)
    )

    for (const roleKey in beforeRoles) {
      const beforeChecked = vueI18nT(`checked__${beforeRoles[roleKey].checked ? 1 : 0}`)
      const afterChecked = vueI18nT(`checked__${afterRoles[roleKey].checked ? 1 : 0}`)
      const pageName = vueI18nT(beforeRoles[roleKey].folderNameKey)
      const roleName = vueI18nT(beforeRoles[roleKey].nameKey)
      detail += `${vueI18nT('backendActionLogDetail_agentsetpermission', [
        pageName,
        roleName,
        beforeChecked,
        afterChecked,
      ])}\n`
    }
  }

  return { title, detail }
}

/** 遊戲設置-修改遊戲 */
function compileGameSetGameState(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  const gameName = vueI18nT(`game__${actionLogJson.game_id}`)
  const beforeState = vueI18nT(`state__${actionLogJson.state.before}`)
  const afterState = vueI18nT(`state__${actionLogJson.state.after}`)
  const detail = `${vueI18nT(`backendActionLogDetail_gamesetgamestate`, [gameName, beforeState, afterState])}\n`

  return { title, detail }
}

/** 玩家帳號-修改 */
function compileUserUpdateGameUserInfo(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [actionLogJson.agent_name, actionLogJson.username])
  const detail = baseCompileActionLogDetail(vueI18nT, actionLogJson)
  return { title, detail }
}

/** 遊戲設置-修改全局狀態 */
function compileGameSetGameServerState(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  const gameServerText = vueI18nT('textGameServer')
  const globalText = vueI18nT('textGlobal')
  const beforeState = vueI18nT(`state__${actionLogJson.state.before}`)
  const afterState = vueI18nT(`state__${actionLogJson.state.after}`)
  const body = `${vueI18nT(`backendActionLogDetail_gamesetgameserverstate`, [
    gameServerText,
    globalText,
    beforeState,
    afterState,
  ])}\n`

  return { title, body }
}

/** 遊戲設置-創建更新 */
function compileGameNotifyGameServer(vueI18nT, featureCode, actionLogJson) {
  const gameName = vueI18nT(`game__${actionLogJson.game_id}`)
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [gameName])
  return { title, detail: null }
}

/** 風控功能-玩家標示 */
function compileRiskControlSetGameUsersCustomTag(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  const agentNameRow = vueI18nT('fmtTextAgentName', [actionLogJson.agent_name])
  const userNameRow = vueI18nT('fmtTextPlayerAccount', [actionLogJson.username])

  let detail = `${agentNameRow}\n${userNameRow}\n`

  if (actionLogJson.is_risk) {
    const beforeChecked = vueI18nT(`checked__${actionLogJson.is_risk.before ? 1 : 0}`)
    const afterChecked = vueI18nT(`checked__${actionLogJson.is_risk.after ? 1 : 0}`)
    detail += `${vueI18nT('backendActionLogDetail_riskcontrolsetgameuserscustomtag_is_risk', [
      beforeChecked,
      afterChecked,
    ])}\n`
  }

  if (actionLogJson.kill_dive_state) {
    // before表示從有到無所以是勾選 => 未勾選，after則相反
    if (actionLogJson.kill_dive_state.before > 0) {
      const beforeChecked = vueI18nT(`checked__1`)
      const afterChecked = vueI18nT(`checked__0`)
      detail += `${vueI18nT(
        `backendActionLogDetail_riskcontrolsetgameuserscustomtag_kill_dive_state__${actionLogJson.kill_dive_state.before}`,
        [beforeChecked, afterChecked]
      )}\n`
    }
    if (actionLogJson.kill_dive_state.after > 0) {
      const beforeChecked = vueI18nT(`checked__0`)
      const afterChecked = vueI18nT(`checked__1`)
      detail += `${vueI18nT(
        `backendActionLogDetail_riskcontrolsetgameuserscustomtag_kill_dive_state__${actionLogJson.kill_dive_state.after}`,
        [beforeChecked, afterChecked]
      )}\n`
    }
  }

  if (actionLogJson.kill_dive_value) {
    detail += `${vueI18nT('backendActionLogDetail_riskcontrolsetgameuserscustomtag_kill_dive_value', [
      actionLogJson.kill_dive_value.before,
      actionLogJson.kill_dive_value.after,
    ])}\n`
  }

  if (actionLogJson.custom_status) {
    for (let i = 0; i < actionLogJson.custom_status.before.length; i++) {
      const before = actionLogJson.custom_status.before[i]
      const after = actionLogJson.custom_status.after[i]

      if (before === after) {
        continue
      }

      const customTagName = actionLogJson.custom_tag_info[i].name
      const beforeChecked = vueI18nT(`checked__${before}`)
      const afterChecked = vueI18nT(`checked__${after}`)
      detail += `${vueI18nT('backendActionLogDetail_riskcontrolsetgameuserscustomtag_custom_status', [
        customTagName,
        beforeChecked,
        afterChecked,
      ])}\n`
    }
  }

  return { title, detail }
}

/** 遊戲設置-修改遊戲排序 */
function compileGameSetGameIconList(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)
  let detail = ''

  if (actionLogJson.is_default.before !== actionLogJson.is_default.after) {
    const beforeSetting = actionLogJson.is_default.before
      ? vueI18nT('textUseSuperiorSorting')
      : vueI18nT('textUseCustomSorting')
    const afterSetting = actionLogJson.is_default.after
      ? vueI18nT('textUseSuperiorSorting')
      : vueI18nT('textUseCustomSorting')

    detail += `${vueI18nT('backendActionLogDetail_gamesetgameiconlist_is_default', [beforeSetting, afterSetting])}\n`
  }

  for (let i = 0; i < actionLogJson.icon_list.before.length; i++) {
    const before = actionLogJson.icon_list.before[i]
    const after = actionLogJson.icon_list.after.find((icon) => icon.gameId === before.gameId)

    Object.keys(before).forEach((key) => {
      if (before[key] !== after[key]) {
        switch (key) {
          case 'hot':
            detail += `${vueI18nT('backendActionLogDetail_gamesetgameiconlist_hot', [
              vueI18nT(`game__${before.gameId}`),
              vueI18nT(`enabledType__${before[key]}`),
              vueI18nT(`enabledType__${after[key]}`),
            ])}\n`
            break
          case 'newest':
            detail += `${vueI18nT('backendActionLogDetail_gamesetgameiconlist_newest', [
              vueI18nT(`game__${before.gameId}`),
              vueI18nT(`enabledType__${before[key]}`),
              vueI18nT(`enabledType__${after[key]}`),
            ])}\n`
            break
          case 'rank':
            detail += `${vueI18nT('backendActionLogDetail_gamesetgameiconlist_rank', [
              vueI18nT(`game__${before.gameId}`),
              before[key],
              after[key],
            ])}\n`
            break
          case 'push':
            detail += `${vueI18nT('backendActionLogDetail_gamesetgameiconlist_push', [
              vueI18nT(`game__${before.gameId}`),
              vueI18nT(`gameIconPush__${before[key]}`),
              vueI18nT(`gameIconPush__${after[key]}`),
            ])}\n`
            break
        }
      }
    })
  }

  return { title, detail }
}

function compileAgentSetBaseIpWhitelist(vueI18nT, featureCode, actionLogJson, ipKey) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [actionLogJson['title']])

  let detail = ''

  const logIp = (ipArr) => {
    for (const ipObj of ipArr) {
      detail += `${vueI18nT('backendActionLogDetail_ip', [
        time.localTimeFormat(ipObj.create_time * 1000),
        ipObj.ip_address,
        ipObj.info || vueI18nT('textEmpty'),
        ipObj.creator,
      ])}\n`
    }
  }

  logIp(actionLogJson[ipKey].before)
  detail += `${vueI18nT('textModifyTo')}\n`
  logIp(actionLogJson[ipKey].after)

  return { title, detail }
}
/** 系統管理-後台IP白名單 */
function compileAgentSetAgentIpWhitelist(vueI18nT, featureCode, actionLogJson) {
  return compileAgentSetBaseIpWhitelist(vueI18nT, featureCode, actionLogJson, 'ip')
}

/** 系統管理-API IP白名單 */
function compileAgentSetAgentApiIpWhitelist(vueI18nT, featureCode, actionLogJson) {
  return compileAgentSetBaseIpWhitelist(vueI18nT, featureCode, actionLogJson, 'api_ip')
}

/** 風控功能-遊戲風控設定 */
function compileRiskControlSetIncomeRatio(vueI18nT, featureCode, actionLogJson) {
  const agentName = actionLogJson['agent_name'] || vueI18nT('textDefault')
  const gameName = vueI18nT(`game__${actionLogJson['game_id']}`)
  const roomTypeName = vueI18nT(`roomType__${roomTypeNameIndex(actionLogJson['game_id'], actionLogJson['room_type'])}`)

  return {
    title: vueI18nT(`backendActionLogTitle__${featureCode}`, [agentName, gameName, roomTypeName]),
    detail: baseCompileActionLogDetail(vueI18nT, actionLogJson),
  }
}

/** 系統管理-匯率設定 */
function compileSystemSetExchangeDataList(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  let detail = ''
  for (const exchangeData of actionLogJson['exchange_data']) {
    const currencyName = vueI18nT(`currency__${exchangeData.currency}`)
    const toCoinBefore = exchangeData.to_coin.before
    const toCoinAfter = exchangeData.to_coin.after
    detail += `${vueI18nT('backendActionLogDetail_systemsetexchangedatalist_to_coin', [
      currencyName,
      toCoinBefore,
      toCoinAfter,
    ])}\n`
  }

  return { title, detail }
}

/** 風控管理-自動風控設定 */
function compileRiskControlSetAutoRiskControlSetting(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  let detail = ''
  for (const [key, value] of Object.entries(actionLogJson)) {
    let before, after

    if (key === 'game_user_win_rate_limit') {
      before = `${value.before * 100}%`
      after = `${value.after * 100}%`
    } else if (key === 'is_enabled') {
      before = vueI18nT(`enabledType__${actionLogJson.is_enabled.before ? 1 : 0}`)
      after = vueI18nT(`enabledType__${actionLogJson.is_enabled.after ? 1 : 0}`)
    } else {
      before = value.before
      after = value.after
    }

    detail += `${vueI18nT(`backendActionLogDetail_${key}`, [before, after])}\n`
  }

  return {
    title,
    detail,
  }
}

/** 玩家帳號-處置玩家設定 */
function compileRiskControlSetGameUserRiskControlTag(vueI18nT, featureCode, actionLogJson) {
  const agentName = actionLogJson.agent_name
  const userName = actionLogJson.username
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [agentName, userName])

  const before = [...actionLogJson.risk_control_status.before]
  const after = [...actionLogJson.risk_control_status.after]
  const beforeTags = []
  const afterTags = []

  const riskTags = ['1000', '0100', '0010', '0001']
  for (let i = 0; i < before.length; i++) {
    if (before[i] === '1') {
      beforeTags.push(vueI18nT(`riskControlTag__${riskTags[i]}`))
    }
    if (after[i] === '1') {
      afterTags.push(vueI18nT(`riskControlTag__${riskTags[i]}`))
    }
  }

  if (!beforeTags.length) {
    beforeTags.push(vueI18nT('riskControlTag__0000'))
  }
  if (!afterTags.length) {
    afterTags.push(vueI18nT('riskControlTag__0000'))
  }

  let detail = ''
  detail += `${vueI18nT(`backendActionLogDetail_risk_control_status`, [
    `${beforeTags.join(`${vueI18nT('symbolComma')}`)}`,
    `${afterTags.join(`${vueI18nT('symbolComma')}`)}`,
  ])}`

  return {
    title,
    detail,
  }
}

/** 重置密碼 */
function compileUserResetPassword(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [actionLogJson['username']])

  return { title, detail: null }
}

/** JACKPOT功能管理-參加設定 */
function compileJackpotSetting(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [actionLogJson['title']])

  let detail = ''
  let firstStatus
  const defaultTime = '1970-01-01T00:00:00Z'
  const isNeverJoin = (status, time) => status === 0 && time === defaultTime

  for (const [key, status] of Object.entries(actionLogJson)) {
    if (key === 'jackpot_status') {
      firstStatus = status.before
      detail += `${vueI18nT(`backendActionLogDetail_${key}`, [
        `${vueI18nT(`jackpotStatus__${status.before}`)}`,
        `${vueI18nT(`jackpotStatus__${status.after}`)}`,
      ])}\n`
    } else if (key === 'jackpot_start_time' || key === 'jackpot_end_time') {
      const beforeTime = isNeverJoin(firstStatus, status.before) ? '-' : time.utcTimeStrToLocalTimeFormat(status.before)
      const afterTime = isNeverJoin(firstStatus, status.after) ? '-' : time.utcTimeStrToLocalTimeFormat(status.after)
      detail += `${vueI18nT(`backendActionLogDetail_${key}`, [beforeTime, afterTime])}\n`
    }
  }

  return { title, detail }
}

/** JACKPOT功能管理-添加代幣 */
function compileJackpotToken(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [
    actionLogJson['agent_name'],
    actionLogJson['username'],
  ])

  let detail = `${vueI18nT(`backendActionLogDetail_jackpot_token_id`, [actionLogJson['token_id']])}`
  return { title, detail }
}

function compileRiskControlSetGameSetting(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)

  let detail = ''

  const items = actionLogJson.game_settings
  for (let i = 0; i < items.length; i++) {
    const item = items[i]

    detail += `【${vueI18nT(`game__${item.game_id}`)}】\n`

    const before = JSON.parse(item.before)
    const after = JSON.parse(item.after)

    for (const [key, value] of Object.entries(before)) {
      let beforeValue = value
      let afterValue = after[key]

      if (beforeValue === afterValue) {
        continue
      }

      switch (key) {
        case 'MatchGameRTP':
        case 'MatchGameKillRate':
        case 'NormalMatchGameRTP':
        case 'NormalMatchGameKillRate':
        case 'LowBoundRTP':
          beforeValue = beforeValue * 100
          afterValue = afterValue * 100
          break
      }

      detail += `${vueI18nT(`backendActionLogDetail_game_setting_${key}`, [beforeValue, afterValue])}\n`
    }

    if (items.length > 1 && i + 1 !== items.length) {
      detail += `------------------------------\n`
    }
  }

  return { title, detail }
}

/** 遊戲風控設定-批次設定 */
function compileRiskControlSetIncomeRatios(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)
  let detail = ''

  const logs = actionLogJson.agent_game_ratios
  for (let i = 0; i < logs.length; i++) {
    const log = logs[i]

    for (const [key, value] of Object.entries(log)) {
      let before = log[key].before
      let after = log[key].after

      if (!before && isNaN(before)) {
        before = vueI18nT(`textEmpty`)
      }
      if (!after && isNaN(after)) {
        after = vueI18nT(`textEmpty`)
      }

      let gameId
      switch (key) {
        case 'agent_name':
          detail += `${vueI18nT(`backendActionLogDetail_${key}`, [value])}\n`
          break
        case 'info':
        case 'active_num':
          detail += `${vueI18nT(`backendActionLogDetail_${key}`, [before, after])}\n`
          break
        case 'game_id':
          detail += `${vueI18nT(`backendActionLogDetail_${key}`, [vueI18nT(`game__${value}`)])}\n`
          gameId = value
          break
        case 'room_type':
          detail += `${vueI18nT(`backendActionLogDetail_${key}`, [
            vueI18nT(`roomType__${roomTypeNameIndex(gameId, value)}`),
          ])}\n`
          break
        default:
          detail += `${vueI18nT(`backendActionLogDetail_${key}`, [before * 100, after * 100])}\n`
          break
      }
    }
    detail += `------------------------------\n`
  }
  return { title, detail }
}

function compileAgentSetAgentCoinSupplyInfo(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`, [actionLogJson['title']])
  let detail = baseCompileActionLogDetail(vueI18nT, actionLogJson)

  const lobbySwitchInfo = actionLogJson.lobby_switch_info
  if (lobbySwitchInfo) {
    Object.values(constant.AgentLobbySwitch).forEach((lobbySwitch) => {
      const beforeChecked = (lobbySwitchInfo.before & lobbySwitch) === lobbySwitch
      const afterChecked = (lobbySwitchInfo.after & lobbySwitch) === lobbySwitch
      if (beforeChecked === afterChecked) {
        return
      }

      detail += `${vueI18nT('backendActionLogDetail_lobby_switch_info', [
        vueI18nT(`lobbySwitch__${lobbySwitch}`),
        vueI18nT(`checked__${beforeChecked ? 1 : 0}`),
        vueI18nT(`checked__${afterChecked ? 1 : 0}`),
      ])}\n`
    })
  }

  return { title, detail }
}

/** 遊戲設置-遊戲罐頭語設定 */
function compileGameSetCannedList(vueI18nT, featureCode, actionLogJson) {
  const title = vueI18nT(`backendActionLogTitle__${featureCode}`)
  let detail = ''

  const agentName = actionLogJson.agent_name
  const contentList = actionLogJson.content_list
  const beforeContentList = contentList.before
  const afterContentList = contentList.after

  for (let i = 0; i < afterContentList.length; i++) {
    const after = afterContentList[i]
    const before = beforeContentList.find((b) => b.serial === after.serial)

    const isDefault = after.canned_type === constant.CannedLanguageType.Default
    const isCreated = !before

    const serial = after.serial
    const cannedType = vueI18nT(`cannedLanguageType__${after.canned_type}`)

    // 標頭
    detail += isDefault
      ? `【${cannedType}】【${vueI18nT('textSlotLineId')}-${serial}】\n`
      : `【${agentName}】【${cannedType}】【${vueI18nT('textSlotLineId')}-${serial}】\n`

    // 表情類型
    detail += isCreated
      ? `${vueI18nT('textEmojiType')}: ${vueI18nT('actionType__1')} ${vueI18nT(
          `cannedEmojiType__${after.emotion_type}`
        )}\n`
      : before.emotion_type !== after.emotion_type
      ? `${vueI18nT('textEmojiType')}: ${vueI18nT(`cannedEmojiType__${before.emotion_type}`)} ${vueI18nT(
          'textModifyTo'
        )} ${vueI18nT(`cannedEmojiType__${after.emotion_type}`)}\n`
      : ''

    // 狀態
    detail += isCreated
      ? `${vueI18nT('textState')}: ${vueI18nT('actionType__1')} ${vueI18nT(`cannedStatus__${after.status}`)}\n`
      : before.status !== after.status
      ? `${vueI18nT('textState')}: ${vueI18nT(`cannedStatus__${before.status}`)} ${vueI18nT('textModifyTo')} ${vueI18nT(
          `cannedStatus__${after.status}`
        )}\n`
      : ''

    // 語言內容
    for (const afterContentItem of after.content_list) {
      const beforeContentItem = before?.content_list.find((b) => b.lang === afterContentItem.lang)
      const isCreatedContent = isCreated || !beforeContentItem

      const afterContentItemContent = afterContentItem.content ? afterContentItem.content : vueI18nT('textEmpty')
      const beforeContentItemContent = beforeContentItem?.content ? beforeContentItem.content : vueI18nT('textEmpty')

      detail += isCreatedContent
        ? `${vueI18nT(`langType__${afterContentItem.lang}`)}: ${vueI18nT('actionType__1')} ${afterContentItemContent}\n`
        : beforeContentItem.content !== afterContentItem.content
        ? `${vueI18nT(`langType__${afterContentItem.lang}`)}: ${beforeContentItemContent} ${vueI18nT(
            'textModifyTo'
          )} ${afterContentItemContent}\n`
        : ''
    }

    if (afterContentList.length > 1 && i + 1 !== afterContentList.length) {
      detail += `------------------------------\n`
    }
  }

  return { title, detail }
}
