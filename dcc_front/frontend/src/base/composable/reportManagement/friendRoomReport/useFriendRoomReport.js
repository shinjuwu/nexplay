import { inject, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { GetFriendRoomReportListInput } from '@/base/common/table/getFriendRoomReportList'
import time from '@/base/utils/time'

export function useFriendRoomReport() {
  const { t } = useI18n()
  const warn = inject('warn')

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()
  const formInput = reactive({
    agent: {},
    game: {},
    userName: '',
    roomId: '',
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })
  const tableInput = reactive(
    new GetFriendRoomReportListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.game.id,
      formInput.roomId,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'createTime',
      constant.TableSortDirection.Desc
    )
  )

  const records = reactive({
    items: [],
  })

  function validateForm() {
    let errors = []

    const startTime = formInput.startTime
    const endTime = formInput.endTime
    if (endTime - startTime < 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
      return errors
    } else if (endTime - startTime > time.commonReportTimeRange * 24 * 60 * 60 * 1000) {
      errors.push(
        t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t('fmtTextDays', [time.commonReportTimeRange])])
      )
      return errors
    }

    return errors
  }

  async function searchRecords() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    tableInput.showProcessing = true

    const tmpTableInput = new GetFriendRoomReportListInput(
      formInput.agent.id,
      formInput.userName,
      formInput.game.id,
      formInput.roomId,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await api.getFriendRoomLogList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.winloseReportTimeBeforeDays]))
        return
      }

      records.items = resp.data.data.map((record) => ({
        id: record.id,
        agentId: record.agent_id,
        agentName: record.agent_name,
        userName: record.username,
        gameId: record.game_id,
        roomId: record.room_id,
        tax: record.tax,
        taxpercent: record.taxpercent,
        createTime: record.create_time,
        endTime: record.end_time,
        detail: record.detail,
      }))

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

  const isShowFriendRoomDetail = ref(false)
  const friendRoomInfo = reactive({
    id: '',
    agentId: 0,
    agentName: '',
    userName: '',
    gameId: 0,
    roomId: '',
    tax: 0,
    taxpercent: 0,
    createTime: '',
    endTime: '',
    detail: '',
  })

  function showFriendRoomDetail(record) {
    friendRoomInfo.id = record.id
    friendRoomInfo.agentId = record.agentId
    friendRoomInfo.agentName = record.agentName
    friendRoomInfo.userName = record.userName
    friendRoomInfo.gameId = record.gameId
    friendRoomInfo.roomId = record.roomId
    friendRoomInfo.tax = record.tax
    friendRoomInfo.taxpercent = record.taxpercent
    friendRoomInfo.createTime = record.createTime
    friendRoomInfo.endTime = record.endTime
    friendRoomInfo.detail = record.detail

    isShowFriendRoomDetail.value = true
  }

  function closeFriendRoomDetail() {
    isShowFriendRoomDetail.value = false
  }

  return {
    formInput,
    tableInput,
    records,
    searchRecords,
    isShowFriendRoomDetail,
    friendRoomInfo,
    showFriendRoomDetail,
    closeFriendRoomDetail,
  }
}
