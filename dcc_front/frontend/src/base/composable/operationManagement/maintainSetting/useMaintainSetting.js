import { computed, inject, reactive, ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { format } from 'date-fns'
import * as gameApi from '@/base/api/sysGame'
import * as manageApi from '@/base/api/sysManage'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'

export function useMaintainSetting() {
  const { t } = useI18n()

  const confirm = inject('confirm')
  const warn = inject('warn')

  const isSettingEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.MaintainSettingUpdate)
  })

  const gameServerState = ref(-1)
  const showGameServerProcessing = ref(false)

  async function searchGameServerState() {
    try {
      showGameServerProcessing.value = true

      const resp = await gameApi.getGameServerState()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      gameServerState.value = resp.data.data.state
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showGameServerProcessing.value = false
    }
  }

  function setGameServerState(state) {
    if (gameServerState.value === state) {
      return
    }

    const msg = t('fmtTextUpdateStateTitle', [t('textGlobal'), t(`state__${state}`)])
    confirm(msg).then(async () => {
      try {
        showGameServerProcessing.value = true

        const resp = await gameApi.setGameServerState({
          state: state,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          gameServerState.value = state
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
        showGameServerProcessing.value = false
      }
    })
  }

  const maintainPageSetting = reactive({
    startTime: '',
    endTime: '',
    timezone: '',
  })
  const maintainPageStartTime = computed({
    get() {
      if (!maintainPageSetting.startTime) {
        return null
      }
      return new Date(maintainPageSetting.startTime)
    },
    set(value) {
      maintainPageSetting.startTime = format(value, 'yyyy-MM-dd HH:mm')
    },
  })
  const maintainPageEndTime = computed({
    get() {
      if (!maintainPageSetting.endTime) {
        return null
      }
      return new Date(maintainPageSetting.endTime)
    },
    set(value) {
      maintainPageSetting.endTime = format(value, 'yyyy-MM-dd HH:mm')
    },
  })
  const showMaintainPageProcessing = ref(false)

  async function searchMaintainPageSetting() {
    try {
      showMaintainPageProcessing.value = true

      const resp = await manageApi.getMaintainPageSetting()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      maintainPageSetting.startTime = data.start_time
      maintainPageSetting.endTime = data.end_time
      maintainPageSetting.timezone = data.timezone
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showMaintainPageProcessing.value = false
    }
  }

  function validateMaintainPageSetting() {
    const errors = []

    if (!maintainPageSetting.timezone) {
      errors.push(t('textTimezoneRequired'))
    }

    if (!maintainPageSetting.startTime) {
      errors.push(t('textStartTime'))
    }
    if (!maintainPageSetting.endTime) {
      errors.push(t('textEndTime'))
    }

    return errors
  }

  async function setMaintainPageSetting() {
    showMaintainPageProcessing.value = true

    const errors = validateMaintainPageSetting()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      showMaintainPageProcessing.value = false
      return
    }

    try {
      const resp = await manageApi.setMaintainPageSetting({
        start_time: maintainPageSetting.startTime,
        end_time: maintainPageSetting.endTime,
        timezone: maintainPageSetting.timezone,
      })

      warn(t(`errorCode__${resp.data.code}`))
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showMaintainPageProcessing.value = false
    }
  }

  onMounted(async () => {
    await searchGameServerState()
    await searchMaintainPageSetting()
  })

  return {
    gameServerState,
    isSettingEnabled,
    maintainPageEndTime,
    maintainPageSetting,
    maintainPageStartTime,
    showGameServerProcessing,
    showMaintainPageProcessing,
    setGameServerState,
    setMaintainPageSetting,
  }
}
