import { onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysLogin'
import constant from '@/base/common/constant'
import { useUserStore } from '@/base/store/userStore'
import * as token from '@/base/utils/token'

export function useLogin() {
  const { t } = useI18n()
  const route = useRoute()
  const router = useRouter()
  const uStore = useUserStore()

  const userName = ref('')
  const password = ref('')
  const graphicsCaptcha = reactive({ captchaId: '', captcha: '', imageBase64: '' })
  const errorMessage = ref('')
  const showPassword = ref(false)

  /** 取得server captcha資訊 */
  async function updateGraphicsCaptcha() {
    try {
      const resp = await api.captcha()

      if (resp.data.code !== constant.ErrorCode.Success) {
        errorMessage.value = t(`errorCode__${resp.data.code}`)
        return
      }

      const data = resp.data.data
      graphicsCaptcha.captchaId = data.captchaId
      graphicsCaptcha.imageBase64 = data.picPath
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        errorMessage.value = t(`errorCode__${errorCode}`)
      }
    }
  }

  async function clearAndUpdateCaptcha() {
    graphicsCaptcha.captcha = ''
    await updateGraphicsCaptcha()
  }

  /** 送出server登入請求 */
  async function submit() {
    if (!userName.value || !password.value || !graphicsCaptcha.captcha) {
      errorMessage.value = t('textContainEmptyData')
      return
    }

    try {
      const resp = await api.login({
        username: userName.value,
        password: password.value,
        captcha: graphicsCaptcha.captcha,
        captchaId: graphicsCaptcha.captchaId,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        await clearAndUpdateCaptcha()
        errorMessage.value = t(`errorCode__${resp.data.code}`)
        return
      }

      const data = resp.data.data

      token.set({
        access_token: data.token,
        expires_in: data.expiresAt,
      })

      const { setUser } = uStore
      setUser(data.userData)

      let path = route.query.redirect
      if (path === undefined) {
        path = '/'
      }
      router.push(path)
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        errorMessage.value = t(`errorCode__${errorCode}`)
      }
    }
  }

  onMounted(async () => {
    await updateGraphicsCaptcha()
  })

  return {
    userName,
    password,
    graphicsCaptcha,
    errorMessage,
    showPassword,
    updateGraphicsCaptcha,
    submit,
  }
}
