<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/admin'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { useUserStore } from '@/store/userStore'
import { validateRules, required, lowercaseEnglishAndNumber4To16, englishAndNumber8To16 } from '@/utils/validation'

import BaseFormInput from '@/components/BaseFormInput.vue'

const { user } = storeToRefs(useUserStore())

const account = ref('')
const accountErrMsg = ref('')
const password = ref('')
const passwordErrMsg = ref('')
const nickName = ref('')
const nickNameErrMsg = ref('')
const permissions = reactive({ items: [] as string[] })
const note = ref('')

function togglePermission(permission: string) {
  const index = permissions.items.indexOf(permission)
  if (index !== -1) {
    permissions.items.splice(index, 1)
  } else {
    permissions.items.push(permission)
  }
}

function checkPermission(permission: string) {
  return permissions.items.indexOf(permission) >= 0
}

function validateForm() {
  let isValid = true

  let validateResult = validateRules(required(), lowercaseEnglishAndNumber4To16())(account.value, '帐号')
  isValid = isValid && validateResult.isValid
  accountErrMsg.value = validateResult.errorMessage

  validateResult = validateRules(required(), englishAndNumber8To16())(password.value, '密码')
  isValid = isValid && validateResult.isValid
  passwordErrMsg.value = validateResult.errorMessage

  return isValid
}

const router = useRouter()
const { warn } = useDialogStore()

function requestRegister() {
  if (!validateForm()) {
    return
  }

  processApiRequest(async () => {
    const axiosResp = await api.register({
      username: account.value,
      password: password.value,
      nickname: nickName.value,
      permissions: user.value.permissions.filter((permission) => permissions.items.indexOf(permission) >= 0),
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
      <h2 class="text-xl font-bold">新增会员帐号</h2>
    </header>

    <hr class="my-2" />

    <form>
      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 before:text-red-500 before:content-['*'] sm:w-2/12">帐号:</label>
        <BaseFormInput
          v-model="account"
          :maxlength="16"
          class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12"
          :error-message="accountErrMsg"
        />
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 before:text-red-500 before:content-['*'] sm:w-2/12">密码:</label>
        <BaseFormInput
          v-model="password"
          type="password"
          :maxlength="16"
          class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12"
          :error-message="passwordErrMsg"
        />
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

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">权限:</label>
        <div class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12">
          <label
            v-for="permission in user.permissions"
            :key="`permission_${permission}`"
            class="mr-1 inline-flex cursor-pointer items-center"
            @click="togglePermission(permission)"
          >
            <font-awesome-icon
              v-show="checkPermission(permission)"
              icon="fa-solid fa-square-check"
              class="text-primary mr-0.5"
            />
            <font-awesome-icon v-show="!checkPermission(permission)" icon="fa-regular fa-square" class="mr-0.5" />
            <span>{{ permission.toUpperCase() }}站监控</span>
          </label>
        </div>
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">备注:</label>
        <textarea
          v-model="note"
          class="form-input focus:border-primary mb-1 rounded border sm:w-8/12 md:w-6/12 lg:w-4/12"
          maxlength="50"
          rows="4"
        ></textarea>
      </div>

      <div class="text-danger">
        <div>* 帐号4~16位数，仅能使用小写英文数字</div>
        <div>* 密码8~16位数，仅能使用大小写英文数字</div>
        <div>* 昵称最大支援16位数</div>
        <div>* 备注最大支援50位数</div>
      </div>

      <div class="mt-4">
        <div class="text-right sm:w-10/12 md:w-8/12 lg:w-6/12">
          <router-link v-slot="{ navigate }" custom to="/member-management">
            <button class="btn btn-light mb-1 mr-1 w-full sm:w-4/12" type="button" @click="navigate">
              返回会员管理列表
            </button>
          </router-link>
          <button class="btn btn-primary mb-1 w-full sm:w-2/12" type="button" @click="requestRegister()">新增</button>
        </div>
      </div>
    </form>
  </section>
</template>
