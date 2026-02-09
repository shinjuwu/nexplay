import { reactive, ref, inject, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import constant from '@/base/common/constant'
import { GetRTPSettingListInput } from '@/base/common/table/getRTPSettingListInput'
import * as api from '@/base/api/sysRiskControl'
import { round } from '@/base/utils/math'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'

export function useRTPGameSetting(props) {
  const { t } = useI18n()
  const warn = inject('warn')
  const confirm = inject('confirm')
  const uStore = useUserStore()

  const isEditEnabled = computed(() => {
    const { isInRole } = uStore
    return isInRole(roleItemKey.RTPBatchSettingUpdate)
  })

  const BaiRenDirections = [t('textBairenGamesDirection1'), t('textBairenGamesDirection2')]

  const showSettingDialog = ref(false)
  const showBatchSettingDialog = ref(false)

  const formInput = reactive({
    agent: { id: constant.Agent.All },
    gameName: { id: 0 },
    roomType: constant.RoomType.All,
  })

  const tableInput = reactive(
    new GetRTPSettingListInput(
      formInput.agent.id,
      formInput.gameName.id,
      props.gameType,
      formInput.roomType,
      constant.TableDefaultLength,
      'gameCode',
      constant.TableSortDirection.Asc
    )
  )

  const records = reactive({ items: [] })

  const selectRecord = reactive({
    agentName: '',
    gameName: '',
    roomType: '',
    killRatio: 0,
    newKillRatio: 0,
    activeNum: 0,
    remark: '',
  })

  async function searchRecords(gameType) {
    tableInput.showProcessing = true

    try {
      const tmpTableInput = new GetRTPSettingListInput(
        formInput.agent.id,
        formInput.gameName.id,
        gameType,
        formInput.roomType,
        constant.TableDefaultLength,
        'gameCode',
        constant.TableSortDirection.Asc
      )

      const resp = await api.getIncomeRatioList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      records.items = data.map((d) => {
        return {
          id: d.id,
          agentId: d.agent_id,
          agentName: d.agent_name,
          gameType: d.game_type,
          gameCode: d.game_code,
          gameId: d.game_id,
          roomType: d.room_type,
          killRatio: round(d.kill_ratio, 4),
          killRatioSort: d.is_parent ? -1 : d.kill_ratio,
          newKillRatio: round(d.new_kill_ratio, 4),
          activeNum: d.active_num,
          lastUpdateTime: d.last_update_time,
          isParent: d.is_parent,
        }
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

  async function getRatioSet(record) {
    try {
      const resp = await api.getIncomeRatio({
        agent_id: record.agentId,
        game_id: record.gameId,
        game_type: record.gameType,
        room_type: record.roomType,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      selectRecord.id = record.id
      selectRecord.agentName = record.agentName
      selectRecord.isParent = record.isParent
      selectRecord.agentId = data.agent_id
      selectRecord.gameId = data.game_id
      selectRecord.gameType = data.game_type
      selectRecord.roomType = data.room_type
      selectRecord.killRatio = round(data.kill_ratio, 4)
      selectRecord.newKillRatio = round(data.new_kill_ratio, 4)
      selectRecord.activeNum = data.active_num
      selectRecord.remark = data.info

      showSettingDialog.value = true
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

  function validateForm(payload) {
    const errors = []

    if (!Number.isInteger(payload.killRatio) || payload.killRatio < 0 || payload.killRatio > 100) {
      errors.push(t('textRTPRangeError'))
    }

    if (!Number.isInteger(payload.newKillRatio) || payload.newKillRatio < 0 || payload.newKillRatio > 100) {
      errors.push(t('textNewKillRatioRangeError'))
    }

    if (!Number.isInteger(payload.activeNum) || payload.activeNum < 0 || payload.activeNum > 999) {
      errors.push(t('textActiveNumRangeError'))
    }

    return errors
  }

  async function setGameRatio(payload) {
    const errors = validateForm(payload)
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    const msg = t('textRTPSettingReminder')
    const title =
      payload.gameType === constant.GameType.BaiRen
        ? t('fmtTextUpdateBaiRenRTPSetting', [t(`game__${payload.gameId}`), t(`roomType__${payload.roomType}`)])
        : t('fmtTextUpdateChiPaiRTPSetting', [
            payload.agentName,
            t(`game__${payload.gameId}`),
            t(`roomType__${payload.roomType}`),
          ])

    confirm(msg, { title }).then(async () => {
      try {
        const resp = await api.setIncomeRatio({
          id: payload.id,
          kill_ratio: round(payload.killRatio / 100, 4),
          new_kill_ratio: round(payload.newKillRatio / 100, 4),
          active_num: payload.activeNum,
          info: payload.remark,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          searchRecords(payload.gameType)
          showSettingDialog.value = false
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
    })
  }

  return {
    formInput,
    tableInput,
    records,
    showSettingDialog,
    showBatchSettingDialog,
    selectRecord,
    BaiRenDirections,
    isEditEnabled,
    searchRecords,
    getRatioSet,
    setGameRatio,
  }
}
