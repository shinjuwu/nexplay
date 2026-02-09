import { inject, reactive, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'
import * as apiR from '@/base/api/sysRecord'
import * as apiG from '@/base/api/sysGlobal'
import constant from '@/base/common/constant'
import { compileActionLog } from '@/base/common/backendActionLog/compiler'
import { roleGroups, menuItemKey } from '@/base/common/menuConstant'
import { GetBackendActionLogListInput } from '@/base/common/table/getBackendActionLogListInput'
import { useUserStore } from '@/base/store/userStore'
import time from '@/base/utils/time'

export function useBackendActionLog() {
  const { t } = useI18n()
  const warn = inject('warn')

  const { user } = storeToRefs(useUserStore())

  const dropdownActionTypes = [
    constant.ActionLogType.All,
    constant.ActionLogType.Create,
    constant.ActionLogType.Update,
    constant.ActionLogType.Delete,
  ]

  const dropdownCategory = reactive({ items: [{ key: 'All', nameKey: 'textAllCategory', featureCodes: [] }] })
  onBeforeMount(async () => {
    try {
      const resp = await apiG.getAgentAdminuserPermissionList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      const allowFeatureCodes = resp.data.data.reduce((res, fcCode) => {
        res[fcCode] = true
        return res
      }, {})

      for (const roleGroup of roleGroups) {
        // 後台代理上下分不需要顯示
        if (roleGroup.folderKey === menuItemKey.BackendUpdateAgentWallet) {
          continue
        }

        const faltFeatureCodes = roleGroup.items
          .filter((item) => !item.key.endsWith('Read') && item.permissions.every((fcCode) => allowFeatureCodes[fcCode]))
          .flatMap((item) => item.permissions)
        if (faltFeatureCodes.length === 0) {
          continue
        }

        dropdownCategory.items.push({
          key: roleGroup.folderKey,
          nameKey: `menuItem${roleGroup.folderKey}`,
          featureCodes: faltFeatureCodes,
        })
      }
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

  const timeRange = time.getCurrentTimeRageByMinutesAndStep(30, time.earningReportTimeMinuteIncrement)
  const formInput = reactive({
    agent: {},
    actionType: dropdownActionTypes[0],
    category: dropdownCategory.items[0],
    startTime: timeRange.startTime,
    endTime: timeRange.endTime,
    betId: '',
  })
  const tableInput = reactive(
    new GetBackendActionLogListInput(
      formInput.agent.id,
      formInput.actionType,
      formInput.category.featureCodes,
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
    const errors = []
    if (formInput.endTime - formInput.startTime < 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
      return errors
    } else if (formInput.endTime - formInput.startTime > time.commonReportTimeRange * 24 * 60 * 60 * 1000) {
      errors.push(
        t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t('fmtTextDays', [time.commonReportTimeRange])])
      )
      return errors
    }

    return errors
  }

  async function searchBackendActionLogList() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }
    tableInput.showProcessing = true

    const tmpTableInput = new GetBackendActionLogListInput(
      formInput.agent.id,
      formInput.actionType,
      formInput.category.featureCodes,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await apiR.getBackendActionLogList(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`, [time.commonReportTimeBeforeDays]))
        return
      }

      records.items = resp.data.data.map((d) => {
        return {
          id: d.id,
          userName: d.username,
          featureCode: d.feature_code,
          actionType: d.action_type,
          actionLog: d.action_log,
          log: compileActionLog(t, d.feature_code, JSON.parse(d.action_log), user),
          showLogDetail: false,
          createTime: time.utcTimeStrToLocalTimeFormat(d.create_time),
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

  return {
    dropdownActionTypes,
    dropdownCategory,
    formInput,
    tableInput,
    records,
    searchBackendActionLogList,
  }
}
