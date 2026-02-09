import { inject, ref, reactive, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { cloneDeep } from 'lodash'

import { useUserStore } from '@/base/store/userStore'
import { useBaseCannedLanguage } from '@/base/composable/gameSetting/gameCannedLanguage/useBaseCannedLanguage'
import constant from '@/base/common/constant'

export function useDefaultCannedLanguage() {
  const { isAdminUser } = useUserStore()

  const isEditable = isAdminUser()

  const warn = inject('warn')
  const { t } = useI18n()
  const showTableProcessing = ref(false)
  const baseFunc = useBaseCannedLanguage()
  const cannedList = reactive({ items: [] })
  let backupCannedList = []

  async function getCannedList() {
    showTableProcessing.value = true
    try {
      const resp = await baseFunc.getCannedList(constant.CannedLanguageType.Default, -1)

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      backupCannedList = resp.data.data
        .map((d) => ({
          serial: d.serial,
          cannedType: d.canned_type,
          emotionType: d.emotion_type,
          langContentMap: d.content_list.reduce((map, data) => {
            map.set(data.lang, data.content)
            return map
          }, new Map()),
          status: d.status,
        }))
        .sort((a, b) => a.serial - b.serial)
      cannedList.items = cloneDeep(backupCannedList)
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showTableProcessing.value = false
    }
  }

  onBeforeMount(async () => {
    await getCannedList()
  })

  async function setCannedList() {
    showTableProcessing.value = true

    const updateList = baseFunc.getUpdateList(backupCannedList, cannedList.items)
    if (updateList.length === 0) {
      showTableProcessing.value = false
      warn(t(`errorCode__${constant.ErrorCode.ErrorNoChangeInData}`))
      return
    }

    try {
      const resp = await baseFunc.setCannedList(constant.CannedLanguageType.Default, -1, updateList)
      warn(t(`errorCode__${resp.data.code}`)).then(async () => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        await getCannedList()
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
      showTableProcessing.value = false
    }
  }

  return {
    showTableProcessing,
    isEditable,
    cannedList,
    setCannedList,
  }
}
