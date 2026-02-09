<script setup lang="ts">
import type { GetUserInfoListResponse } from '@/types/types.api-admin'
import type { ReactiveArr } from '@/types/types.common'
import type { TableEmitSearchPayload } from '@/types/types.table'

import { reactive, ref, onBeforeMount, computed } from 'vue'
import { storeToRefs } from 'pinia'

import { processApiRequest, parseServerErrorMessage } from '@/api/base'
import * as api from '@/api/admin'
import { responseCode } from '@/common/constant'
import { useDialogStore } from '@/store/dialogStore'
import { useUserStore } from '@/store/userStore'
import time from '@/utils/time'

import BaseTable from '@/components/BaseTable.vue'

const { user } = storeToRefs(useUserStore())

const filterAccount = ref('')
const showForm = ref(true)
interface Member {
  account: string
  nickName: string
  status: boolean
  createTime: Date
  lastLoginTime: Date
}
const memberArr = reactive<ReactiveArr<Member>>({ items: [] })
const filterMembers = computed(() => memberArr.items.filter((member) => member.account.includes(filterAccount.value)))

const tableFastPagingContainer = ref<HTMLElement | null>(null)
const tableInfoContainer = ref<HTMLElement | null>(null)
const tableLengthMenuContainer = ref<HTMLElement | null>(null)
const tablePaginationContainer = ref<HTMLElement | null>(null)
const tableProcessingContainer = ref<HTMLElement | null>(null)
const tableSearchButtonContainer = ref<HTMLElement | null>(null)

const { warn } = useDialogStore()

function requestGetUsersInfoList(payload?: TableEmitSearchPayload) {
  processApiRequest(
    async () => {
      const axiosResp = await api.getUsersInfoList()

      if (axiosResp.data.code !== responseCode.Success) {
        warn(parseServerErrorMessage(axiosResp))
        return
      }

      const list = JSON.parse(axiosResp.data.data) as GetUserInfoListResponse

      memberArr.items = list.data.map((d) => ({
        account: d.username,
        nickName: d.nickname,
        status: d.is_enabled,
        createTime: new Date(d.create_time),
        lastLoginTime: new Date(d.last_login_time),
      }))
      payload?.resetTable()
    },
    warn,
    payload?.showProcessing,
    payload?.closeProcessing
  )
}

onBeforeMount(() => {
  requestGetUsersInfoList()
})
</script>

<template>
  <section ref="tableProcessingContainer" class="relative rounded bg-white p-4">
    <header class="flex items-center justify-between">
      <h2 class="text-xl font-bold">会员管理</h2>
      <span class="cursor-pointer" @click="showForm = !showForm">
        <font-awesome-icon v-if="showForm" icon="fa-solid fa-circle-minus" />
        <font-awesome-icon v-else icon="fa-solid fa-circle-plus" />
      </span>
    </header>

    <hr class="my-2" />
    <form v-show="showForm">
      <div class="flex flex-wrap items-center justify-between">
        <router-link v-slot="{ navigate }" custom to="/member-management/create-account">
          <button class="btn btn-danger mb-1 w-full md:w-2/12 xl:w-1/12" type="button" @click="navigate">
            <font-awesome-icon icon="fa-solid fa-user-plus" class="mr-0.5" />
            新增帐号
          </button>
        </router-link>
        <div ref="tableSearchButtonContainer" class="mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"></div>
      </div>
    </form>

    <hr v-show="showForm" class="my-2" />

    <BaseTable
      :items="filterMembers"
      :processing-container="tableProcessingContainer"
      :search-button-container="tableSearchButtonContainer"
      :fast-paging-container="tableFastPagingContainer"
      :info-container="tableInfoContainer"
      :length-menu-container="tableLengthMenuContainer"
      :pagination-container="tablePaginationContainer"
      @search="requestGetUsersInfoList"
    >
      <template #default="{ currentItems }">
        <div class="flex flex-wrap items-center justify-end">
          <label
            class="mb-1 flex w-full items-center whitespace-nowrap rounded bg-gray-200 px-3 py-1 text-white sm:w-auto"
          >
            搜索:
            <input
              v-model="filterAccount"
              type="text"
              class="ml-3 box-border w-full min-w-0 border px-2 py-1 text-black"
            />
          </label>
        </div>

        <table class="tbl tbl-rwd tbl-striped">
          <thead class="tbl-thead">
            <tr class="tbl-tr">
              <th class="tbl-th">帐号</th>
              <th class="tbl-th">昵称</th>
              <th class="tbl-th">状态</th>
              <th class="tbl-th">建立时间</th>
              <th class="tbl-th">最后登录时间</th>
              <th class="tbl-th">操作</th>
            </tr>
          </thead>
          <tbody class="tbl-tbody">
            <tr v-if="currentItems.length === 0" class="tbl-tr">
              <td class="block border px-2 py-1 lg:table-cell lg:p-3.5" colspan="6">没有找到匹配的纪录</td>
            </tr>
            <tr v-for="member in currentItems" v-else :key="member.account" class="tbl-tr">
              <td class="tbl-td tbl-td-ellipsis before:content-['帐号']">{{ member.account }}</td>
              <td class="tbl-td tbl-td-ellipsis before:content-['昵称']">{{ member.nickName }}</td>
              <td class="tbl-td before:content-['状态']">{{ member.status ? '开启' : '封停' }}</td>
              <td class="tbl-td before:content-['建立时间']">{{ time.format(member.createTime) }}</td>
              <td class="tbl-td before:content-['最后登录时间']">{{ time.format(member.lastLoginTime) }}</td>
              <td class="tbl-td before:content-['操作'] lg:space-x-1">
                <template v-if="user.name !== member.account">
                  <router-link v-slot="{ navigate }" custom :to="`/member-management/edit-account/${member.account}`">
                    <button class="btn-tool btn-light my-1" type="button" @click="navigate">修改</button>
                  </router-link>
                  <router-link v-slot="{ navigate }" custom :to="`/member-management/edit-password/${member.account}`">
                    <button class="btn-tool btn-light my-1" type="button" @click="navigate">修改密码</button>
                  </router-link>
                </template>
              </td>
            </tr>
          </tbody>
        </table>
      </template>
    </BaseTable>

    <div class="tbl-dom-wrapper">
      <div ref="tableLengthMenuContainer" class="tbl-length-wrapper"></div>
      <div ref="tableInfoContainer" class="tbl-info-wrapper"></div>
      <div ref="tableFastPagingContainer" class="tbl-fast-paginate-wrapper"></div>
      <div ref="tablePaginationContainer" class="tbl-paginate-wrapper"></div>
    </div>
  </section>
</template>
