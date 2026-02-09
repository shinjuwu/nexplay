<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show" @submit.prevent>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAccount') }}</label>
          <input
            v-model="formInput.account"
            class="form-input mb-1 md:w-3/12"
            type="text"
            :placeholder="t('placeHolderTextAccount')"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <div class="flex flex-1">
            <button
              v-if="isEnabled.BackendAccountCreate"
              type="button"
              class="btn btn-danger mb-1 flex w-full items-center justify-center md:ml-2 md:w-fit"
              @click="dialog.addSubAccounts = true"
            >
              <UserPlusIcon class="mr-2 inline-block h-5 w-5" />
              {{ t('textBackendAccountAddBackendAccount') }}
            </button>
          </div>
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchRecords()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageTable :records="records.items" :total-records="records.items.length" :table-input="tableInput">
        <template
          #default="{
            currentRecords,
            pageLength,
            recordStart,
            totalPages,
            totalRecords,
            isSortIconActive,
            lengthChange,
            pageChange,
            sorting,
          }"
        >
          <div class="tbl-container">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="nickName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textNickName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="parentAccount"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textParentAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="role" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textGroupRole') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="state" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textState') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCreateDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="lastLoginTime"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textLastLoginDate') }} </template>
                  </PageTableSortableTh>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr>
                    <td colspan="8">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="(record, index) in currentRecords" :key="`record_${index}`">
                    <td>{{ record.userName }}</td>
                    <td>{{ record.nickName }}</td>
                    <td>{{ record.parentAccount }}</td>
                    <td>{{ record.role }}</td>
                    <td>
                      <span :class="{ 'text-danger': !record.state }">
                        {{ record.state ? t('textOpened') : t('textDisabled') }}
                      </span>
                    </td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createTime) }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.lastLoginTime) }}</td>
                    <td class="space-x-2">
                      <ViewButton @click="accountSetting(record, false)" />
                      <EditButton
                        v-if="isEnabled.BackendAccountUpdate && record.parentAccount !== '-'"
                        @click="accountSetting(record, true)"
                      />
                      <ResetPasswordButton
                        v-if="isEnabled.ResetPassword && record.parentAccount !== '-'"
                        @click="resetPassword(record)"
                      />
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
          <div class="mt-6 text-center text-slate-500 md:flex md:flex-wrap md:items-center">
            <PageTableMenuLength class="mb-1.5 md:mr-6" :display-length="pageLength" @length-change="lengthChange" />
            <PageTableInfo
              class="mb-1.5 md:mr-6"
              :display-records="currentRecords.length"
              :record-start="recordStart"
              :total-records="totalRecords"
            />
            <PageTableQuickPage
              class="mb-1.5"
              :page-length="pageLength"
              :record-start="recordStart"
              :total-pages="totalPages"
              @page-change="pageChange"
            />
            <PageTablePagination
              class="mb-1.5"
              :page-length="pageLength"
              :record-start="recordStart"
              :total-pages="totalPages"
              @page-change="pageChange"
            />
          </div>
        </template>
      </PageTable>
    </template>
  </ToggleHeader>
  <AddSubAccounts
    :visible="dialog.addSubAccounts"
    @close="(newValue) => (dialog.addSubAccounts = newValue)"
    @search-records="searchRecords"
  />
  <Setting
    :visible="dialog.setting"
    :select-record="selectRecord"
    :is-edit-enabled="isSettingEditEnabled"
    @close="(newValue) => (dialog.setting = newValue)"
    @search-records="searchRecords"
  />
  <ResetPassword
    :visible="dialog.resetPassword"
    :user-name="selectRecord.userName"
    :type="'BackendAccount'"
    @close="(newValue) => (dialog.resetPassword = newValue)"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { UserPlusIcon } from '@heroicons/vue/20/solid'
import { useBackendAccount } from '@/base/composable/systemManagement/backendAccount/useBackendAccount'
import time from '@/base/utils/time'

import ViewButton from '@/base/components/Button/ViewButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import ResetPasswordButton from '@/base/components/Button/ResetPasswordButton.vue'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import AddSubAccounts from '@/base/views/SystemManagement/BackendAccount/BackendAccountAddSubAccounts.vue'
import Setting from '@/base/views/SystemManagement/BackendAccount/BackendAccountSetting.vue'
import ResetPassword from '@/base/views/Common/Dialog/ResetPasswordDialog.vue'

const { t } = useI18n()
const {
  dialog,
  formInput,
  isEnabled,
  isSettingEditEnabled,
  records,
  selectRecord,
  tableInput,
  searchRecords,
  accountSetting,
  resetPassword,
} = useBackendAccount()
</script>
