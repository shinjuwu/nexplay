import { inject, reactive, ref, watch, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysManage'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useChatServiceStore, chatServerMessageSubject } from '@/base/store/chatServiceStore'
import { useUserStore } from '@/base/store/userStore'
import time from '@/base/utils/time'

export function useBackendAnnouncement() {
  const { t } = useI18n()
  const warn = inject('warn')

  const isEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return {
      BackendAnnouncementCreate: isInRole(roleItemKey.BackendAnnouncementCreate),
      BackendAnnouncementUpdate: isInRole(roleItemKey.BackendAnnouncementUpdate),
      BackendAnnouncementDelete: isInRole(roleItemKey.BackendAnnouncementDelete),
    }
  })

  const showDialog = ref(false)
  const deleteDialog = ref(false)
  const showMode = ref('')

  const types = Object.values(constant.AnnounceType)
  const formInput = reactive({
    type: types[0],
  })
  const tableInput = reactive(
    new BaseTableInput(constant.TableDefaultLength, 'updateTime', constant.TableSortDirection.Desc)
  )

  const announcementInfo = reactive({
    id: '',
    createTime: '',
    updateTime: '',
    type: 1,
    subject: '',
    content: '',
  })

  const records = reactive({ items: [] })

  function createNewAnnounce() {
    showMode.value = 'create'
    showDialog.value = true
  }

  async function searchRecords() {
    tableInput.showProcessing = true

    try {
      const resp = await api.getAnnouncementList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data

      let ret = data.map((d) => {
        return {
          id: d.id,
          createTime: d.create_time,
          updateTime: d.update_time,
          type: d.type,
          subject: d.subject,
        }
      })

      if (formInput.type !== constant.AnnounceType.All) {
        ret = ret.filter((r) => r.type === formInput.type)
      }

      records.items = ret

      const { sendReadAllSubjectMessages } = useChatServiceStore()
      sendReadAllSubjectMessages(chatServerMessageSubject.announcement)

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

  async function getAnnouncementInfo(announceId, mode) {
    try {
      const resp = await api.getAnnouncement({ id: announceId })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data

      announcementInfo.id = data.id
      announcementInfo.createTime = time.utcTimeStrToLocalTimeFormat(data.create_time)
      announcementInfo.updateTime = time.utcTimeStrToLocalTimeFormat(data.update_time)
      announcementInfo.type = data.type
      announcementInfo.subject = data.subject
      announcementInfo.content = data.content

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

  async function deleteAnnouncement() {
    try {
      const resp = await api.deleteAnnouncement({ id: announcementInfo.id })

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

  watch(
    () => showMode.value,
    () => {
      if (showMode.value === 'create') {
        Object.assign(announcementInfo, {
          id: '',
          createTime: '',
          updateTime: '',
          type: 1,
          subject: '',
          content: '',
        })
      }
    }
  )

  onMounted(() => {
    searchRecords()
  })

  return {
    formInput,
    tableInput,
    records,
    types,
    showDialog,
    showMode,
    deleteDialog,
    announcementInfo,
    isEnabled,
    getAnnouncementInfo,
    createNewAnnounce,
    deleteAnnouncement,
    searchRecords,
  }
}
