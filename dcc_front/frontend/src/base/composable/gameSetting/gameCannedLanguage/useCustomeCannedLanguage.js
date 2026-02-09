import { inject, computed, ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import axios from 'axios'
import { cloneDeep } from 'lodash'

import { useUserStore } from '@/base/store/userStore'
import { useBaseCannedLanguage } from '@/base/composable/gameSetting/gameCannedLanguage/useBaseCannedLanguage'
import constant from '@/base/common/constant'

export function useCustomeCannedLanguage() {
  const uStore = useUserStore()
  const { user } = storeToRefs(uStore)
  const isAdminUser = uStore.isAdminUser()

  const formInput = reactive({ agent: {} })
  const cannedListAgentId = ref(-1)
  const isEditable = computed(() => user.value.agentId === cannedListAgentId.value)

  const warn = inject('warn')
  const { t } = useI18n()
  const showTableProcessing = ref(false)
  const baseFunc = useBaseCannedLanguage()
  const cannedList = reactive({ items: [] })
  let backupCannedList = []

  async function getCannedList() {
    showTableProcessing.value = true
    try {
      const resp = await baseFunc.getCannedList(constant.CannedLanguageType.Custome, formInput.agent.id)

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      cannedListAgentId.value = formInput.agent.id
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

  async function setCannedList() {
    showTableProcessing.value = true

    const updateList = baseFunc.getUpdateList(backupCannedList, cannedList.items)
    if (updateList.length === 0) {
      showTableProcessing.value = false
      warn(t(`errorCode__${constant.ErrorCode.ErrorNoChangeInData}`))
      return
    }

    try {
      const resp = await baseFunc.setCannedList(constant.CannedLanguageType.Custome, formInput.agent.id, updateList)
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
    isAdminUser,
    formInput,
    cannedListAgentId,
    showTableProcessing,
    isEditable,
    cannedList,
    getCannedList,
    setCannedList,
  }
}
