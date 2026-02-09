<script setup lang="ts">
import type { CaptchaResponse, UsersLoginResponse } from '@/types/types.api-login'

import { onBeforeMount, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/login'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { useUserStore } from '@/store/userStore'
import {
  validateRules,
  required,
  stringLength,
  lowercaseEnglishAndNumber4To16,
  englishAndNumber8To16,
} from '@/utils/validation'
import * as token from '@/utils/token'

import BaseFormInput from '@/components/BaseFormInput.vue'

const account = ref('')
const accountErrMsg = ref('')
const password = ref('')
const passwordErrMsg = ref('')
const showPassword = ref(false)
const captcha = ref('')
const captchaErrMsg = ref('')

interface GraphicsCaptcha {
  captchaId: string
  imageBase64: string
  captchaLength: number
  expiredTime: Date | null
}
const graphicsCaptcha: GraphicsCaptcha = reactive({
  captchaId: '',
  imageBase64: '',
  captchaLength: 0,
  expiredTime: null,
})

const { warn } = useDialogStore()

function requestCaptcha() {
  processApiRequest(async () => {
    const axiosResp = await api.captcha()

    if (axiosResp.data.code !== responseCode.Success) {
      warn(parseServerErrorMessage(axiosResp))
      return
    }

    const data = JSON.parse(axiosResp.data.data) as CaptchaResponse

    graphicsCaptcha.captchaId = data.captchaId
    graphicsCaptcha.imageBase64 = data.picPath
    graphicsCaptcha.captchaLength = data.captchaLength
    graphicsCaptcha.expiredTime = new Date(data.expiredTime)
  }, warn)
}

function clearAndUpdateCaptcha() {
  captcha.value = ''
  requestCaptcha()
}

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

function validateForm() {
  let isValid = true

  let validateResult = validateRules(required(), lowercaseEnglishAndNumber4To16())(account.value, '帐号')
  isValid = isValid && validateResult.isValid
  accountErrMsg.value = validateResult.errorMessage

  validateResult = validateRules(required(), englishAndNumber8To16())(password.value, '密码')
  isValid = isValid && validateResult.isValid
  passwordErrMsg.value = validateResult.errorMessage

  validateResult = validateRules(required(), stringLength(graphicsCaptcha.captchaLength))(captcha.value, '验证码')
  isValid = isValid && validateResult.isValid
  captchaErrMsg.value = validateResult.errorMessage

  return isValid
}

function requestLogin() {
  if (!validateForm()) {
    return
  }

  processApiRequest(async () => {
    const axiosResp = await api.login({
      username: account.value,
      password: password.value,
      captcha: captcha.value,
      captchaId: graphicsCaptcha.captchaId,
    })

    if (axiosResp.data.code !== responseCode.Success) {
      warn(parseServerErrorMessage(axiosResp)).then(() => {
        clearAndUpdateCaptcha()
      })
      return
    }

    const data = JSON.parse(axiosResp.data.data) as UsersLoginResponse

    token.set(data.token, data.expiresAt)
    userStore.setUser(data.UserData)

    let path = route.query.redirect as string | undefined
    if (path === undefined) {
      path = '/'
    }
    router.push(path)
  }, warn)
}

onBeforeMount(() => {
  requestCaptcha()
})
</script>

<template>
  <section class="flex min-h-screen flex-col justify-center bg-green-100">
    <div class="flex min-h-screen flex-col sm:mx-auto sm:block sm:min-h-0 sm:w-[450px]">
      <div class="flex-1 bg-white px-10 pb-9 pt-12 sm:rounded-lg sm:border sm:border-gray-200">
        <div>
          <img class="mx-auto" src="@/assets/images/pic_logo_rtp.png" />
        </div>
        <div class="text-center">
          <div class="pt-4 text-2xl">登录</div>
          <div class="pt-[7px]">使用您的监控平台帐户</div>
        </div>
        <div class="pt-6">
          <form @keyup.enter="requestLogin()">
            <BaseFormInput
              v-model="account"
              :error-message="accountErrMsg"
              class="mt-3"
              placeholder="请输入您的帐号"
              :maxlength="16"
            />
            <BaseFormInput
              v-model="password"
              class="mt-3"
              placeholder="请输入您的密码"
              :type="showPassword ? 'text' : 'password'"
              :maxlength="16"
              right-icon
              :error-message="passwordErrMsg"
              @click-icon="showPassword = !showPassword"
            >
              <template #icon>
                <font-awesome-icon v-if="showPassword" icon="fa-regular fa-eye-slash" />
                <font-awesome-icon v-else icon="fa-regular fa-eye" />
              </template>
            </BaseFormInput>
            <BaseFormInput
              v-model="captcha"
              class="mt-3"
              placeholder="请输入您的验证码"
              right-icon
              :maxlength="graphicsCaptcha.captchaLength"
              :error-message="captchaErrMsg"
              @click-icon="requestCaptcha()"
            >
              <template #icon>
                <font-awesome-icon icon="fa-solid fa-rotate" />
              </template>
            </BaseFormInput>
            <img
              class="mt-1 w-full cursor-pointer rounded border border-gray-200"
              :src="graphicsCaptcha.imageBase64"
              @click="requestCaptcha()"
            />
          </form>
          <div class="mt-3 flex justify-end pb-5">
            <button class="btn btn-primary my-1.5 h-9 px-6 py-0 text-sm" @click="requestLogin()">登录</button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
