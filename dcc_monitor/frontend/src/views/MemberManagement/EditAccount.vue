<script setup lang="ts">
import type { GetUserInfoResponse } from '@/types/types.api-admin'

import { onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/admin'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { useUserStore } from '@/store/userStore'
import BaseFormInput from '@/components/BaseFormInput.vue'

const route = useRoute()
const { user } = storeToRefs(useUserStore())

const account = ref(route.params.account as string)
const status = ref(false)
const nickName = ref('')
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

const { warn } = useDialogStore()

function requestGetUsersInfo() {
  processApiRequest(async () => {
    const axiosResp = await api.getUsersInfo({
      username: account.value,
    })

    if (axiosResp.data.code !== responseCode.Success) {
      warn(parseServerErrorMessage(axiosResp))
      return
    }

    const data = JSON.parse(axiosResp.data.data) as GetUserInfoResponse
    nickName.value = data.nickname
    status.value = data.is_enabled
    permissions.items = data.permissions
    note.value = data.info
  }, warn)
}

onMounted(() => {
  requestGetUsersInfo()
})

const router = useRouter()

function requestModifyUsersInfo() {
  processApiRequest(async () => {
    const axiosResp = await api.modifyUsersInfo({
      username: account.value,
      nickname: nickName.value,
      is_enabled: status.value,
      permissions: user.value.permissions.filter((permission) => permissions.items.indexOf(permission) >= 0),
      info: note.value,
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
      <h2 class="text-xl font-bold">修改会员帐号</h2>
    </header>

    <hr class="my-2" />

    <form>
      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">帐号:</label>
        <span class="mb-1 w-full sm:w-10/12 sm:px-4 sm:py-2">{{ account }}</span>
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">昵称:</label>
        <BaseFormInput v-model="nickName" :maxlength="16" class="mb-1 w-full sm:w-8/12 md:w-6/12 lg:w-4/12" />
      </div>

      <div class="flex flex-wrap items-center">
        <label class="mb-1 w-full pr-2 sm:w-2/12">状态:</label>
        <select
          v-model="status"
          class="form-input focus:border-primary mb-1 rounded border sm:w-8/12 md:w-6/12 lg:w-4/12"
        >
          <option :value="true">开启</option>
          <option :value="false">封停</option>
        </select>
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
          <button class="btn btn-primary mb-1 w-full sm:w-2/12" type="button" @click="requestModifyUsersInfo()">
            修改
          </button>
        </div>
      </div>
    </form>
  </section>
</template>
