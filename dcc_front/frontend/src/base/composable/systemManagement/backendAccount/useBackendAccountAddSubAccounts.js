import { reactive, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import regex from '@/base/common/regex'

export function useBackendAccountAddSubAccounts(emit) {
  const warn = inject('warn')
  const { t } = useI18n()

  const newAccountForm = reactive({
    account: '',
    password: '',
    nickName: '',
    role: '',
    remark: '',
  })

  async function createAccount() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const resp = await api.createAdminUser({
        username: newAccountForm.account,
        password: newAccountForm.password,
        nickname: newAccountForm.nickName,
        role: newAccountForm.role,
        info: newAccountForm.remark,
      })
      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        emit('searchRecords')
        emit('close')

        newAccountForm.account = ''
        newAccountForm.password = ''
        newAccountForm.nickName = ''
        newAccountForm.role = ''
        newAccountForm.remark = ''
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
    emit('close', false)
  }

  function validateForm() {
    const errors = []

    if (!(regex.account.test(newAccountForm.account) && regex.password.test(newAccountForm.password))) {
      errors.push(t('textAccountErrorMessage'))
    }
    if (!newAccountForm.nickName) {
      errors.push(t('textNicknameRequired'))
    }
    if (!newAccountForm.role) {
      errors.push(t('textRoleGroupRequired'))
    }

    return errors
  }

  return {
    t,
    newAccountForm,
    createAccount,
    close,
  }
}
