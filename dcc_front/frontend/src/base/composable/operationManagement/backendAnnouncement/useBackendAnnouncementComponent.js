import { useI18n } from 'vue-i18n'
import { ref, computed, reactive, inject, watch } from 'vue'
import * as api from '@/base/api/sysManage'
import constant from '@/base/common/constant'
import axios from 'axios'

export function useBackendAnnouncementComponent(props, emit) {
  const { t } = useI18n()
  const warn = inject('warn')

  const typeOptions = [
    {
      value: 1,
      label: t('typeCode__1'),
    },
    {
      value: 2,
      label: t('typeCode__2'),
    },
  ]

  const visible = ref(false)
  const content = computed(() => {
    let title = ''
    let submit = ''
    switch (props.mode) {
      case 'create':
        title = t('roleItemBackendAnnouncementCreate')
        submit = t('textIncrease')
        break
      case 'update':
        title = t('roleItemBackendAnnouncementUpdate')
        submit = t('textOnSave')
        break
      case 'view':
        title = t('menuItemBackendAnnouncement')
    }
    return { title, submit }
  })

  const disabled = computed(() => props.mode === 'view')

  const announcementInfo = reactive({
    id: props.announcementInfo.id,
    type: props.announcementInfo.type,
    subject: props.announcementInfo.subject,
    content: props.announcementInfo.content,
  })

  function validateForm() {
    const errors = []

    if (!announcementInfo.subject) {
      errors.push(t('fmtTextInputErrorMessage', [t('textSubject')]))
    }

    if (!announcementInfo.content) {
      errors.push(t('fmtTextInputErrorMessage', [t('textContentText')]))
    }
    return errors
  }

  async function createAnnouncement() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const resp = await api.createAnnouncement({
        type: announcementInfo.type,
        subject: announcementInfo.subject,
        content: announcementInfo.content,
      })
      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        close()
        emit('refreshAnnouncementList')
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

  async function updateAnnouncement() {
    try {
      const resp = await api.updateAnnouncement({ ...announcementInfo })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        close()
        emit('refreshAnnouncementList')
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

  function close() {
    visible.value = false
    emit('close')
  }

  watch(props, async (newValue) => {
    if (newValue) {
      announcementInfo.id = props.announcementInfo.id
      announcementInfo.type = props.announcementInfo.type
      announcementInfo.subject = props.announcementInfo.subject
      announcementInfo.content = props.announcementInfo.content
      announcementInfo.createTime = props.announcementInfo.createTime
    }
    visible.value = props.visible
  })

  return {
    t,
    visible,
    content,
    typeOptions,
    announcementInfo,
    disabled,
    createAnnouncement,
    updateAnnouncement,
    close,
  }
}
