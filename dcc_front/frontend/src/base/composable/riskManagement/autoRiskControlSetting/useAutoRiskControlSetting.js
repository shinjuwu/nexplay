import { reactive, ref, inject, onBeforeMount, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRiskControl'
import constant from '@/base/common/constant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useUserStore } from '@/base/store/userStore'
import { roleItemKey } from '@/base/common/menuConstant'

export function useAutoRiskControlSetting() {
  const { t } = useI18n()
  const warn = inject('warn')

  const isEnabledAutoRiskControl = ref(false)
  const tableInput = ref(new BaseTableInput(constant.TableDefaultLength, '', constant.TableSortDirection.Asc))
  const directions = [
    t('textAutoRiskControlSettingDirection1'),
    t('textAutoRiskControlSettingDirection2'),
    t('textAutoRiskControlSettingDirection3'),
    t('textAutoRiskControlSettingDirection4'),
    t('textAutoRiskControlSettingDirection5'),
  ]

  const uStore = useUserStore()
  const isSettingEnabled = computed(() => {
    const { isInRole } = uStore
    return {
      AutoRiskControlSettingUpdate: isInRole(roleItemKey.AutoRiskControlSettingUpdate),
    }
  })

  const autoRiskSettingDetail = computed(() => {
    return {
      autoSetting: autoSetting.items,
      isEnabled: isEnabledAutoRiskControl.value,
    }
  })

  const autoSetting = reactive({
    items: [
      new AutoSetting(
        '',
        t('autoRiskSetting__0'),
        0,
        0,
        t('autoRiskSettingDisposeUpScoreAndDownScore'),
        t('autoRiskSettingDeposeTime__1'),
        t('textFreq')
      ),
      new AutoSetting(
        '',
        t('autoRiskSetting__1'),
        0,
        0,
        t('textDisabled'),
        t('autoRiskSettingDeposeTime__2'),
        t('textFreq')
      ),
      new AutoSetting(
        '',
        t('autoRiskSetting__2'),
        0,
        0,
        t('riskControlTag__0001'),
        t('autoRiskSettingDeposeTime__2'),
        t('textPoint')
      ),
      new AutoSetting(
        '',
        t('autoRiskSetting__3'),
        0,
        0,
        t('riskControlTag__0100'),
        t('autoRiskSettingDeposeTime__2'),
        '%'
      ),
    ],
  })

  async function getAutoRiskControlSetting() {
    tableInput.value.showProcessing = true
    try {
      const resp = await api.getAutoRiskControlSetting()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      Object.entries(resp.data.data.current).forEach(([key, limit], idx) => {
        if (autoSetting.items[idx]) {
          autoSetting.items[idx].key = key
          autoSetting.items[idx].limit = key === 'game_user_win_rate_limit' ? limit * 100 : limit
        }
      })
      isEnabledAutoRiskControl.value = resp.data.data.current.is_enabled

      Object.entries(resp.data.data.default).forEach(([key, limit], idx) => {
        if (autoSetting.items[idx]) {
          autoSetting.items[idx].defaultLimit = key === 'game_user_win_rate_limit' ? limit * 100 : limit
        }
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
      tableInput.value.showProcessing = false
    }
  }

  async function setAutoRiskControlSetting() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      tableInput.value.showProcessing = true

      const params = {
        is_enabled: isEnabledAutoRiskControl.value,
      }
      autoSetting.items.forEach((item) => {
        params[item.key] = item.key === 'game_user_win_rate_limit' ? item.limit / 100 : item.limit
      })

      const resp = await api.setAutoRiskControlSetting(params)

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
      tableInput.value.showProcessing = false
    }
  }

  function resetDefaultAll() {
    autoSetting.items.forEach((item) => item.resetDefaultAll())
  }

  function validateForm() {
    const errors = []

    const result = autoSetting.items.every((item) => item.limit >= 0 && item.limit <= 999999999)
    if (!result) {
      errors.push(t('textAutoRiskSettingValueError'))
    }

    return errors
  }

  onBeforeMount(async () => {
    await getAutoRiskControlSetting()
  })

  return {
    t,
    autoSetting,
    directions,
    tableInput,
    isSettingEnabled,
    isEnabledAutoRiskControl,
    autoRiskSettingDetail,
    setAutoRiskControlSetting,
    resetDefaultAll,
  }
}

class AutoSetting {
  constructor(key, name, limit, defaultLimit, dispose, time, unit) {
    this.key = key
    this.name = name
    this.limit = limit
    this.defaultLimit = defaultLimit
    this.dispose = dispose
    this.time = time
    this.unit = unit
  }

  resetDefaultAll() {
    this.limit = this.defaultLimit
  }
}
