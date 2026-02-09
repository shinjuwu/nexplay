import { computed, inject, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { GetJackpotListInput } from '@/base/common/table/getJackpotListInput'
import { useUserStore } from '@/base/store/userStore'
import { storeToRefs } from 'pinia'
import constant from '@/base/common/constant'
import time from '@/base/utils/time'
import axios from 'axios'
import * as api from '@/base/api/sysJackpot'
import { roleItemKey } from '@/base/common/menuConstant'

export function useJackpotTokenRecord() {
  const { t } = useI18n()
  const warn = inject('warn')
  const confirm = inject('confirm')

  const { user } = storeToRefs(useUserStore())
  const isSettingEditEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.JackpotTokenUpdate)
  })

  const showTokenDialog = ref(false)

  const timeRange = time.getCurrentTimeRageByMinutesAndStep()
  const formInput = reactive({
    agent: { id: constant.Agent.All },
    tokenId: '',
    roundId: '',
    userName: '',
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
  })
  const tableInput = reactive(
    new GetJackpotListInput(
      formInput.agent.id,
      formInput.startTime,
      formInput.endTime,
      formInput.roundId,
      formInput.tokenId,
      formInput.userName,
      constant.TableDefaultLength,
      'tokenGetTime',
      constant.TableSortDirection.Desc
    )
  )
  const tokenForm = reactive({
    agent: { id: user.value.agentId },
    userName: '',
    quota: 1,
    remark: '',
  })

  const jackpotTokenList = reactive({
    items: [],
  })

  function validateInputForm() {
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

  async function getJackpotTokenList() {
    const errors = validateInputForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      tableInput.showProcessing = true

      const tmpTableInput = new GetJackpotListInput(
        formInput.agent.id,
        formInput.startTime,
        formInput.endTime,
        formInput.roundId,
        formInput.tokenId,
        formInput.userName,
        tableInput.length,
        tableInput.column,
        tableInput.dir
      )

      const resp = await api.getJackpotTokenList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      jackpotTokenList.items = data.map((d) => {
        return {
          agentId: d.agent_id,
          orderId: d.id,
          tokenId: d.token_id,
          roundId: d.source_lognumber,
          tokenGetTime: d.token_create_time,
          jackpotBet: d.jp_bet,
          agentName: d.agent_name,
          userName: d.username,
          operator: d.creator,
          orderState: d.status,
          errorCode: d.error_code,
          remark: d.info,
          isUsed: d.usage_count,
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

  function validateForm() {
    const errors = []

    if (!tokenForm.userName) {
      errors.push(t('textUserNameEmptyError'))
    }

    if (tokenForm.quota <= 0) {
      errors.push(t('textCoinQuotaError'))
    }

    if (!tokenForm.remark) {
      errors.push(t('textRemarkErrorText'))
    }

    return errors
  }

  function addToToken() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    const title = t('textAddJPToken')
    const info = [
      t('fmtTextAddToCoin', [tokenForm.quota]),
      t('fmtTextAgentName', [tokenForm.agent.name]),
      t('fmtTextPlayerAccount', [tokenForm.userName]),
      t('textUpdateCoinQuotaReminder'),
    ]

    const message = info.join('\n')
    confirm(message, { title }).then(async () => {
      try {
        const resp = await api.createJackpotToken({
          agent_id: tokenForm.agent.id,
          username: tokenForm.userName,
          jp_bet: tokenForm.quota,
          info: tokenForm.remark,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(async () => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }
          await getJackpotTokenList()
          closeDialog()
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
        showTokenDialog.value = false
      }
    })
  }

  function closeDialog() {
    showTokenDialog.value = false
    Object.assign(tokenForm, {
      agent: { id: user.value.id },
      userName: '',
      quota: 1,
      remark: '',
    })
  }

  return {
    tableInput,
    formInput,
    jackpotTokenList,
    showTokenDialog,
    tokenForm,
    isSettingEditEnabled,
    getJackpotTokenList,
    addToToken,
    closeDialog,
  }
}
