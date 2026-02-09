import { inject, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import regex from '@/base/common/regex'
import { useUserStore } from '@/base/store/userStore'

export function usePersonalInfo() {
  const warn = inject('warn')

  const router = useRouter()

  const { t } = useI18n()

  const uStore = useUserStore()
  const { user } = storeToRefs(uStore)

  const myName = user.value.name
  const myNickname = ref(user.value.nickName)

  const oldPassword = ref('')
  const newPassword = ref('')
  const confirmPassword = ref('')

  const showOldPassword = ref(false)
  const showNewPassword = ref(false)
  const showConfirmPassword = ref(false)

  async function setPersonalInfo() {
    if (myNickname.value === '') {
      warn(t('textNicknameRequired'))
      return
    }

    try {
      const resp = await api.setPersonalInfo({
        nickname: myNickname.value,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        const { updateUserInfo } = uStore
        updateUserInfo(myNickname.value)
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

  function validPassword() {
    let isValid = false
    let errorCode = constant.ErrorCode.Success

    if (
      !regex.password.test(oldPassword.value) ||
      !regex.password.test(newPassword.value) ||
      !regex.password.test(confirmPassword.value)
    ) {
      errorCode = constant.ErrorCode.ErrorPasswordFormat
      return { isValid, errorCode }
    }

    if (oldPassword.value === newPassword.value) {
      errorCode = constant.ErrorCode.ErrorPasswordSame
      return { isValid, errorCode }
    }

    if (newPassword.value !== confirmPassword.value) {
      errorCode = constant.ErrorCode.ErrorConfirmPassword
      return { isValid, errorCode }
    }

    isValid = true

    return { isValid, errorCode }
  }

  async function setPersonalPassword() {
    const { isValid, errorCode } = validPassword()
    if (!isValid) {
      warn(t(`errorCode__${errorCode}`))
      return
    }

    try {
      const resp = await api.setPersonalPassword({
        old_password: oldPassword.value,
        new_password: newPassword.value,
        confirm_password: confirmPassword.value,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        // 密碼變更成功後要登出
        const { signOut } = useUserStore()
        signOut(router)
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
    confirmPassword,
    myName,
    myNickname,
    newPassword,
    oldPassword,
    showConfirmPassword,
    showNewPassword,
    showOldPassword,
    setPersonalInfo,
    setPersonalPassword,
  }
}
