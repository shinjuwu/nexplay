<script setup lang="ts">
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/admin'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { validateRules, required, englishAndNumber8To16 } from '@/utils/validation'

import BaseFormInput from '@/components/BaseFormInput.vue'

const route = useRoute()

const account = ref(route.params.account as string)

const newPassword = ref('')
const newPasswordErrMsg = ref('')
const showNewPassword = ref(false)

const confirmPassword = ref('')
const confirmPasswordErrMsg = ref('')
const showConfirmPassword = ref(false)

const passwordSummaryErrMsg = ref('')

function validateFormPasswordForm() {
  let isValid = true

  let validateResult = validateRules(required(), englishAndNumber8To16())(newPassword.value, '新密码')
  isValid = isValid && validateResult.isValid
  newPasswordErrMsg.value = validateResult.errorMessage

  validateResult = validateRules(required(), englishAndNumber8To16())(confirmPassword.value, '确认新密码')
  isValid = isValid && validateResult.isValid
  confirmPasswordErrMsg.value = validateResult.errorMessage

  return isValid
}

function specialValidatePasswordForm() {
  if (newPassword.value !== confirmPassword.value) {
    passwordSummaryErrMsg.value = '新密码与确认新密码不相同'
    return false
  }

  passwordSummaryErrMsg.value = ''

  return true
}

const { warn } = useDialogStore()
const router = useRouter()

function requestModifyUsersPassword() {
  if (!validateFormPasswordForm() && !specialValidatePasswordForm()) {
    return
  }

  processApiRequest(async () => {
    const axiosResp = await api.modifyUsersPassword({
      username: account.value,
      password: newPassword.value,
    })

    if (axiosResp.data.code !== responseCode.Success) {
      warn(parseServerErrorMessage(axiosResp))
      return
    }

    router.push('/member-management')
  }, warn)
}
</script>

<template>
  <section class="rounded bg-white p-4">
    <header>
      <h2 class="text-xl font-bold">修改会员密码</h2>
    </header>

    <hr class="my-2" />

    <form>
      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">帐号:</label>
        <span class="mb-1 w-full sm:w-10/12 sm:px-4 sm:py-2">{{ account }}</span>
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
        <div>* 成功修改密码后将强制该会员登出</div>
      </div>

      <div class="mt-4">
        <div class="text-right sm:w-10/12 md:w-8/12 lg:w-6/12">
          <router-link v-slot="{ navigate }" custom to="/member-management">
            <button class="btn btn-light mb-1 mr-1 w-full sm:w-4/12" type="button" @click="navigate">
              返回会员管理列表
            </button>
          </router-link>
          <button class="btn btn-primary mb-1 w-full sm:w-2/12" type="button" @click="requestModifyUsersPassword()">
            修改
          </button>
        </div>
      </div>
    </form>
  </section>
</template>
