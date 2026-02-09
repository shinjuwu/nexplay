<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'

import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/account'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { useUserStore } from '@/store/userStore'
import { validateRules, required, englishAndNumber8To16 } from '@/utils/validation'

import BaseFormInput from '@/components/BaseFormInput.vue'

const userStore = useUserStore()
const { user } = storeToRefs(userStore)
const { signOut, updateUserInfo } = userStore

const account = user.value.name
const nickName = ref(user.value.nickName)
const nickNameErrMsg = ref('')

const oldPassword = ref('')
const oldPasswordErrMsg = ref('')
const showOldPassword = ref(false)

const newPassword = ref('')
const newPasswordErrMsg = ref('')
const showNewPassword = ref(false)

const confirmPassword = ref('')
const confirmPasswordErrMsg = ref('')
const showConfirmPassword = ref(false)

const passwordSummaryErrMsg = ref('')

const { warn } = useDialogStore()

function requestModifyinfo() {
  processApiRequest(async () => {
    const axiosResp = await api.modifyinfo({
      nickname: nickName.value,
    })

    if (axiosResp.data.code !== responseCode.Success) {
      warn(parseServerErrorMessage(axiosResp))
      return
    } else {
      warn('昵称修改成功')
    }

    updateUserInfo(nickName.value)
  }, warn)
}

function validateFormPasswordForm() {
  let isValid = true

  let validateResult = validateRules(required(), englishAndNumber8To16())(oldPassword.value, '旧密码')
  isValid = isValid && validateResult.isValid
  oldPasswordErrMsg.value = validateResult.errorMessage

  validateResult = validateRules(required(), englishAndNumber8To16())(newPassword.value, '新密码')
  isValid = isValid && validateResult.isValid
  newPasswordErrMsg.value = validateResult.errorMessage

  validateResult = validateRules(required(), englishAndNumber8To16())(confirmPassword.value, '确认新密码')
  isValid = isValid && validateResult.isValid
  confirmPasswordErrMsg.value = validateResult.errorMessage

  return isValid
}

function specialValidatePasswordForm() {
  if (oldPassword.value === newPassword.value) {
    passwordSummaryErrMsg.value = '新密码不能与旧密码相同'
    return false
  }

  if (newPassword.value !== confirmPassword.value) {
    passwordSummaryErrMsg.value = '新密码与确认新密码不相同'
    return false
  }

  passwordSummaryErrMsg.value = ''

  return true
}

function requestModifyPassword() {
  if (!validateFormPasswordForm() && !specialValidatePasswordForm()) {
    return
  }

  processApiRequest(async () => {
    const axiosResp = await api.modifypassword({
      password: newPassword.value,
    })

    if (axiosResp.data.code !== responseCode.Success) {
      warn(parseServerErrorMessage(axiosResp))
      return
    }

    signOut()
  }, warn)
}
</script>

<template>
  <section class="rounded bg-white p-4">
    <header>
      <h2 class="text-xl font-bold">我的帐户</h2>
    </header>

    <hr class="my-2" />

    <form>
      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">帐号:</label>
        <span class="mb-1 w-full sm:w-10/12 sm:px-4 sm:py-2">{{ account }}</span>
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">昵称:</label>
        <BaseFormInput
          v-model="nickName"
          :maxlength="16"
          class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12"
          :error-message="nickNameErrMsg"
        />
      </div>

      <div class="mt-4">
        <div class="text-right sm:w-10/12 md:w-8/12 lg:w-6/12">
          <button class="btn btn-primary w-full sm:w-2/12" type="button" @click="requestModifyinfo()">提交</button>
        </div>
      </div>
    </form>

    <hr class="my-2" />

    <form>
      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">旧密码:</label>
        <BaseFormInput
          v-model="oldPassword"
          :maxlength="16"
          class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12"
          :error-message="oldPasswordErrMsg"
          :type="showOldPassword ? 'text' : 'password'"
          right-icon
          @click-icon="showOldPassword = !showOldPassword"
        >
          <template #icon>
            <font-awesome-icon v-if="showOldPassword" icon="fa-regular fa-eye-slash" />
            <font-awesome-icon v-else icon="fa-regular fa-eye" />
          </template>
        </BaseFormInput>
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">新密码:</label>
        <BaseFormInput
          v-model="newPassword"
          :maxlength="16"
          class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12"
          :error-message="newPasswordErrMsg"
          :type="showNewPassword ? 'text' : 'password'"
          right-icon
          @click-icon="showNewPassword = !showNewPassword"
        >
          <template #icon>
            <font-awesome-icon v-if="showNewPassword" icon="fa-regular fa-eye-slash" />
            <font-awesome-icon v-else icon="fa-regular fa-eye" />
          </template>
        </BaseFormInput>
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">确认新密码:</label>
        <BaseFormInput
          v-model="confirmPassword"
          :maxlength="16"
          class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12"
          :error-message="confirmPasswordErrMsg"
          :type="showConfirmPassword ? 'text' : 'password'"
          right-icon
          @click-icon="showConfirmPassword = !showConfirmPassword"
        >
          <template #icon>
            <font-awesome-icon v-if="showConfirmPassword" icon="fa-regular fa-eye-slash" />
            <font-awesome-icon v-else icon="fa-regular fa-eye" />
          </template>
        </BaseFormInput>
      </div>

      <div class="text-danger">
        <div>* 密码8~16位数，仅能使用大小写英文数字</div>
        <div>* 成功修改密码后将强制登出，请您用新密码再次登入</div>
      </div>

      <div class="mt-4">
        <div class="text-right sm:w-10/12 md:w-8/12 lg:w-6/12">
          <button class="btn btn-primary w-full sm:w-2/12" type="button" @click="requestModifyPassword()">提交</button>
        </div>
      </div>
    </form>
  </section>
</template>
