import { reactive, ref, inject, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'
import * as api from '@/base/api/sysManage'
import axios from 'axios'
import { BaseTableInput } from '@/base/common/table/tableInput'
import time from '@/base/utils/time'

export function useMarqueeSetting() {
  const { t } = useI18n()
  const warn = inject('warn')
  const timeRange = time.getCurrentTimeRageByMinutesAndStep()

  const isEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return {
      MarqueeSettingCreate: isInRole(roleItemKey.MarqueeSettingCreate),
      MarqueeSettingUpdate: isInRole(roleItemKey.MarqueeSettingUpdate),
      MarqueeSettingDelete: isInRole(roleItemKey.MarqueeSettingDelete),
    }
  })

  const showDialog = ref(false)
  const deleteDialog = ref(false)
  const showMode = ref('')

  const records = reactive({ items: [] })
  const marqueeInfo = reactive({
    id: '',
    type: 1,
    lang: constant.LangType.CHS,
    order: '',
    freq: '',
    startDate: timeRange.startTime,
    endDate: timeRange.endTime,
    content: '',
    isEnabled: '',
  })

  const types = Object.values(constant.MarqueeType)
  const langs = Object.values(constant.LangType)
  const tableInput = reactive(
    new BaseTableInput(constant.TableDefaultLength, 'createDate', constant.TableSortDirection.Asc)
  )

  const formInput = reactive({
    type: types[0],
    lang: langs[0],
  })

  function createNewMarquee() {
    showMode.value = 'create'

    delete marqueeInfo.createDate
    Object.assign(marqueeInfo, {
      id: '',
      type: 1,
      lang: constant.LangType.CHS,
      order: '',
      freq: '',
      startDate: timeRange.startTime,
      endDate: timeRange.endTime,
      content: '',
      isEnabled: '',
    })

    showDialog.value = true
  }

  async function getMarqueeInfo(marqueeId, mode) {
    try {
      const resp = await api.getMarquee({
        id: marqueeId,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      marqueeInfo.id = data.id
      marqueeInfo.type = data.type
      marqueeInfo.lang = data.lang
      marqueeInfo.order = data.order
      marqueeInfo.freq = data.freq
      marqueeInfo.content = data.content
      marqueeInfo.isEnabled = data.is_enabled
      marqueeInfo.createDate = time.utcTimeStrNoneISO8601ToLocalTimeFormat(data.create_time)
      marqueeInfo.startDate = time.utcTimeStrToLocalTimeFormat(data.start_time)
      marqueeInfo.endDate = time.utcTimeStrToLocalTimeFormat(data.end_time)

      showMode.value = mode
      if (mode !== 'delete') {
        showDialog.value = true
      } else {
        deleteDialog.value = true
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
  }

  async function searchRecords() {
    tableInput.showProcessing = true

    try {
      const resp = await api.getMarqueeList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }
      const data = resp.data.data
      records.items = data
        .map((d) => {
          return {
            id: d.id,
            createDate: d.create_time,
            type: d.type,
            content: d.content,
            lang: d.lang,
            startDate: d.start_time,
            endDate: d.end_time,
            order: d.order,
            freq: d.freq,
            isEnabled: d.is_enabled,
          }
        })
        .filter((r) => {
          if (formInput.type === 0 && formInput.lang === 'All') return r
          else if (formInput.type !== 0 && formInput.lang === 'All') return formInput.type === r.type
          else if (formInput.lang !== 'All' && formInput.type === 0) return formInput.lang === r.lang
          else return formInput.type === r.type && formInput.lang === r.lang
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

  async function deleteMarquee() {
    try {
      const resp = await api.deleteMarquee({ id: marqueeInfo.id })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        deleteDialog.value = false
        searchRecords()
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
  }

  return {
    t,
    isEnabled,
    records,
    marqueeInfo,
    types,
    langs,
    formInput,
    tableInput,
    showDialog,
    deleteDialog,
    showMode,
    createNewMarquee,
    searchRecords,
    deleteMarquee,
    getMarqueeInfo,
  }
}
