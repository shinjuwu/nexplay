import { computed, inject, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { menuItemKey, roleItemKey } from '@/base/common/menuConstant'
import { GetBatchUserPlayLogListInput } from '@/base/common/table/getUserPlayLogListInput'
import { useBreadcrumbStore } from '@/base/store/breadcrumbStore'
import { useUserStore } from '@/base/store/userStore'
import { round } from '@/base/utils/math'
import { getMenuItemFromMenu } from '@/base/utils/menu'
import { columnSorting } from '@/base/utils/table'
import time from '@/base/utils/time'
import xlsx from '@/base/utils/xlsx'
import { roomTypeNameIndex } from '@/base/utils/room'

export function useWinLoseReport() {
  const { t } = useI18n()
  const warn = inject('warn')

  const { isAdminUser, isAgentSingleWalletUser, isJackpotOpen } = useUserStore()

  const isAdmin = isAdminUser()
  const isShowKillDiveRelative = isAdmin
  const isShowJackpotRelative = isAdmin || isJackpotOpen()
  const isShowSingleWalletRelative = isAdmin || isAgentSingleWalletUser()

  const calDirections = computed(() => {
    let ret = []

    ret.push(t('textWinLoseReportPlatformCalDirection'), t('textGamerWinLoseCalDirection'))

    if (isShowJackpotRelative) {
      ret.push(t('textJPInjectWaterScoreDirection'), t('textTotalJPInjectWaterScoreDirection'))
    }

    ret.push(t('textValueDisplayDirection'))

    return ret
  })

  const tableColSpan = (() => {
    let span = 23
    if (!isShowKillDiveRelative) {
      span = span - 4
    }
    if (!isShowJackpotRelative) {
      span = span - 2
    }
    if (!isShowSingleWalletRelative) {
      span = span - 1
    }
    return span
  })()

  const betslipStatuses = [
    { name: 'All', value: 0 },
    { name: 'Normal', value: 1 },
    { name: 'Abnormal', value: 2 },
  ]

  const timeRange = time.getCurrentTimeRageByMinutesAndStep(
    time.winloseReportTimeRange <= 30 ? time.winloseReportTimeRange : 30
  )
  const formInput = reactive({
    userName: '',
    logNumber: '',
    agent: {},
    game: {},
    roomType: constant.RoomType.All,
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
    betId: '',
    betslipStatus: betslipStatuses[0],
    singleWalletId: '',
    roomId: '',
  })
  const tableInput = reactive(
    new GetBatchUserPlayLogListInput(
      formInput.agent.id,
      formInput.game.id,
      formInput.roomType,
      formInput.userName,
      formInput.logNumber,
      formInput.betId,
      formInput.betslipStatus.value,
      formInput.startTime,
      formInput.endTime,
      formInput.singleWalletId,
      formInput.roomId,
      constant.TableDefaultLength,
      'bettime',
      constant.TableSortDirection.Desc
    )
  )

  const isRedirectToGameLogParseEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GameLogParseRead)
  })

  const records = reactive({
    items: [],
  })

  const totalValidScore = ref(0)
  const totalWinLoseScore = ref(0)
  const totalKilledRecordCount = ref(0)
  const totalDivedRecordCount = ref(0)
  const totalPlayerWinRecordCount = ref(0)
  const totalJackpotInjectWaterScore = ref(0)

  function validateForm() {
    let errors = []
    if (!formInput.betId && !formInput.logNumber && !formInput.startTime && !formInput.endTime) {
      errors.push(t('textLogNumberOrTimeRequired'))
      return errors
    }

    if (!formInput.logNumber && !formInput.betId) {
      const startTime = formInput.startTime
      const endTime = formInput.endTime
      if (endTime - startTime < 0) {
        errors.push(t('textStartTimeLaterThanEndTime'))
        return errors
      } else if (endTime - startTime > time.winloseReportTimeRange * 60 * 1000) {
        const hour = Math.floor(time.winloseReportTimeRange / 60)
        const minute = time.winloseReportTimeRange % 60

        let timeTextKey = ''
        let timeTextArgs = []
        if (hour > 0 && minute > 0) {
          timeTextKey = 'fmtTextHoursMinutes'
          timeTextArgs.push(hour, minute)
        } else if (hour > 0) {
          timeTextKey = 'fmtTextHours'
          timeTextArgs.push(hour)
        } else {
          timeTextKey = 'fmtTextMinutes'
          timeTextArgs.push(minute)
        }

        errors.push(t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t(timeTextKey, timeTextArgs)]))
        return errors
      }
    }

    return errors
  }

  async function searchRecords(input) {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    if (!input) {
      input = new GetBatchUserPlayLogListInput(
        formInput.agent.id,
        formInput.game.id,
        formInput.roomType,
        formInput.userName,
        formInput.logNumber,
        formInput.betId,
        formInput.betslipStatus.value,
        formInput.startTime,
        formInput.endTime,
        formInput.singleWalletId,
        formInput.roomId,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )
    }

    tableInput.showProcessing = true

    try {
      const resp = await api.getBatchUserPlayLogList(input.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.winloseReportTimeBeforeDays]))
        return
      }

      const pageData = resp.data.data
      if (pageData.draw === 0) {
        input.totalRecords = pageData.recordsTotal

        totalValidScore.value = pageData.total_valid_score
        totalWinLoseScore.value = pageData.total_platform_winlose_score
        totalKilledRecordCount.value = pageData.total_kill_games
        totalDivedRecordCount.value = pageData.total_dive_games
        totalPlayerWinRecordCount.value = pageData.total_player_win_games
        totalJackpotInjectWaterScore.value = pageData.total_jp_inject_water_score
      }

      records.items = pageData.data.map((record) => new WinLoseReport(record))

      Object.assign(tableInput, input)
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

  async function downloadRecords() {
    tableInput.showProcessing = true

    try {
      const resp = await api.getUserPlayLogList(tableInput.parseInputJson())
      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.winloseReportTimeBeforeDays]))
        return
      }

      let data = resp.data.data.map((d) => new WinLoseReport(d))
      // 0:全部 1:正常 2:異常
      if (tableInput.betslipStatus !== betslipStatuses[0].value) {
        data = data.filter(
          (d) =>
            tableInput.betslipStatus === betslipStatuses[1].value
              ? d.agentId !== -1 && d.agentName && d.userName // 正常訂單
              : d.agentId === -1 || !d.agentName || !d.userName // 異常訂單
        )
      }
      const sortColumn = getWinLoseReportSortColumn(tableInput.column)
      data = data.sort((a, b) => columnSorting(tableInput.dir, a[sortColumn], b[sortColumn]))

      /* notice: 下方header跟rows裡面要依照順序所以if的順序兩邊必須一致 */
      const headers = [[]]

      headers[0].push(
        t('textOrderNumber'),
        t('textAgentName'),
        t('textCreateTime'),
        t('textPlayerAccount'),
        t('textGameName'),
        t('textRoomType'),
        t('textBeforeBetScore'),
        t('textBet'),
        t('textValidScore'),
        t('textAfterBetScore')
      )

      if (isShowJackpotRelative) {
        headers[0].push(t('textJPInjectWaterPercent'), t('textJPInjectWaterScore'))
      }

      headers[0].push(
        t('textGamerWinScore'),
        t('textGameTax'),
        t('textBonus'),
        t('textGamerWinLoseScore'),
        t('textSettleScore'),
        t('textRoundId')
      )

      if (isShowKillDiveRelative) {
        headers[0].push(t('textRiskControl'))
        headers[0].push(t('textKillProb'))
        headers[0].push(t('textKillLevel'))
        headers[0].push(t('textRealPlayers'))
      }

      if (isShowSingleWalletRelative) {
        headers[0].push(t('textSingleWalletId'))
      }

      headers[0].push(t('textRoomId'))

      const rows = data.map((d) => {
        let ret = {
          betId: d.betId,
          agentName: d.agentName,
          betTime: time.utcTimeStrToLocalTimeFormat(d.betTime),
          userName: d.userName,
          gameName: t(`game__${d.gameId}`),
          roomTypeName: t(`roomType__${roomTypeNameIndex(d.gameId, d.roomType)}`),
          beforeBetScore: d.startScore,
          yaScore: d.yaScore,
          validYaScore: d.validScore,
          afterBetScore: d.afterBetScore,
        }

        if (isShowJackpotRelative) {
          ret.jpInjectRate = d.jpInjectRate
          ret.jpInjectScore = d.jpInjectScore
        }

        ret.deScore = d.deScore
        ret.tax = d.tax
        ret.bonus = d.bonus
        ret.winLoseAmount = d.winLoseAmount
        ret.endScore = d.endScore
        ret.roundId = d.roundId

        if (isShowKillDiveRelative) {
          ret.killType = t(`killType__${d.killType}`)
          ret.killProb = d.killProb
          ret.killLevel = t(`killLevel__${d.killLevel}`)
          ret.realPlayers = d.realPlayers
        }

        if (isShowSingleWalletRelative) {
          ret.walletLedgerId = d.walletLedgerId
        }

        ret.roomId = d.roomId

        return ret
      })

      let fileName = '[UspReport]'
      if (tableInput.betId) {
        fileName = `${fileName}[${tableInput.betId}]`
      } else if (tableInput.logNumber) {
        fileName = `${fileName}[${tableInput.logNumber}]`
      } else if (tableInput.singleWalletId) {
        fileName = `${fileName}[${tableInput.singleWalletId}]`
      } else {
        if (tableInput.agentId !== constant.Agent.All) {
          fileName = `${fileName}[A-${tableInput.agentId}]`
        }
        if (tableInput.gameId !== constant.Game.All) {
          fileName = `${fileName}[G-${tableInput.gameId}]`
        }
        if (tableInput.roomType !== constant.RoomType.All) {
          fileName = `${fileName}[R-${tableInput.roomType}]`
        }
        if (tableInput.betslipStatus !== betslipStatuses[0].value) {
          const betslipStatus = betslipStatuses.find((bs) => bs.value === tableInput.betslipStatus)
          fileName = `${fileName}[${betslipStatus.name}]`
        }
      }
      fileName = `${fileName}${time.recordTimeFormat(tableInput.startTime)}-${time.recordTimeFormat(
        tableInput.endTime
      )}`

      xlsx.createAndDownloadFile(fileName, headers, rows)
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

  function redirectToGameLogParse(logNumber, userName) {
    let item = getMenuItemFromMenu(menuItemKey.GameLogParse)
    item.props.logNumber = logNumber
    item.props.userName = userName

    const { addBreadcrumbItem } = useBreadcrumbStore()
    addBreadcrumbItem(item)
  }

  return {
    betslipStatuses,
    calDirections,
    formInput,
    isShowJackpotRelative,
    isShowKillDiveRelative,
    isShowSingleWalletRelative,
    isRedirectToGameLogParseEnabled,
    records,
    tableInput,
    totalDivedRecordCount,
    totalKilledRecordCount,
    totalPlayerWinRecordCount,
    totalValidScore,
    totalWinLoseScore,
    totalJackpotInjectWaterScore,
    tableColSpan,
    downloadRecords,
    redirectToGameLogParse,
    searchRecords,
  }
}

function getWinLoseReportSortColumn(sortColumn) {
  switch (sortColumn) {
    case 'betid':
      return 'betId'
    case 'agentname':
      return 'agentId'
    case 'username':
      return 'userName'
    case 'gameid':
      return 'gameId'
    case 'roomtype':
      return 'roomType'
    case 'yascore':
      return 'yaScore'
    case 'validscore':
      return 'validScore'
    case 'jpinjectwaterrate':
      return 'jpInjectRate'
    case 'jpinjectwaterscore':
      return 'jpInjectScore'
    case 'descore':
      return 'deScore'
    case 'winlosescore':
      return 'winLoseAmount'
    case 'roundid':
      return 'roundId'
    case 'killtype':
      return 'killType'
    case 'tax':
    case 'bonus':
      return sortColumn
    default:
      return 'betTime'
  }
}

class WinLoseReport {
  constructor(data) {
    this.agentId = data.agent_id
    this.agentName = data.agent_name
    this.betTime = data.bet_time
    this.userName = data.username
    this.gameId = data.game_id
    this.roomType = data.room_type
    this.deskId = data.desk_id
    this.seatId = data.seat_id
    this.startScore = this.calStartScore(data.game_id, data.start_score, data.ya_score)
    this.yaScore = data.ya_score
    this.validScore = data.valid_score
    this.afterBetScore = round(this.startScore - data.ya_score, 4)
    this.jpInjectRate = data.jp_inject_water_rate
    this.jpInjectScore = data.jp_inject_water_score
    this.deScore = data.de_score
    this.tax = data.tax
    this.bonus = data.bonus
    this.winLoseAmount = round(data.de_score - data.ya_score - data.tax + data.bonus, 4)
    this.endScore = data.end_score
    this.roundId = data.lognumber
    this.betId = data.bet_id
    this.killType = data.kill_type
    this.killProb = data.kill_prob
    this.killLevel = data.kill_level
    this.realPlayers = data.real_players
    this.walletLedgerId = data.wallet_ledger_id
    this.roomId = data.room_id
  }

  calStartScore(gameId, startScore, yaScore) {
    const gameType = Math.floor(gameId / 1000)
    if (gameType === constant.GameType.BaiRen) {
      return round(startScore + yaScore, 4)
    } else {
      return startScore
    }
  }
}
