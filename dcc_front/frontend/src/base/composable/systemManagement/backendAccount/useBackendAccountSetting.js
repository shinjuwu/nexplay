import { reactive, watch, inject, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'

export function useBackendAccountSetting(props, emit) {
  const { t } = useI18n()

  const warn = inject('warn')
  const confirm = inject('confirm')

  const show = ref(false)

  const account = reactive({
    userName: '',
    state: '',
    role: '',
    roleName: '',
    remark: '',
  })

  const stateOptions = [constant.AccountStatus.Open, constant.AccountStatus.Disable]

  function close() {
    show.value = false
    emit('close', false)
  }

  function submit() {
    confirm(t('textUpdateGroupRoleReminder2')).then(async () => {
      try {
        const resp = await api.updateAdminUserInfo({
          username: account.userName,
          is_enabled: account.state,
          role: account.role,
          info: account.remark,
        })

        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }
          emit('searchRecords')
          emit('close')
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

  async function searchAdminUserInfo() {
    try {
      const resp = await api.getAdminUserInfo({ username: props.selectRecord.userName })
      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          close()
        })
        return
      }

      const data = resp.data.data
      account.userName = props.selectRecord.userName
      account.state = data.is_enabled
      account.remark = data.info
      account.role = data.role
      account.roleName = props.selectRecord.role

      show.value = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }

      emit('close', false)
    }
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        await searchAdminUserInfo()
      } else {
        show.value = false
      }
    }
  )

  return {
    account,
    show,
    stateOptions,
    close,
    submit,
  }
}
