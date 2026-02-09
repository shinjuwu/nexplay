<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
        </div>
        <div class="mt-4 flex flex-wrap items-center">
          <div class="flex flex-1 flex-col md:flex-row">
            <button
              v-if="isEnabled.AgentAccountCreate"
              type="button"
              class="btn btn-danger mb-1 flex items-center justify-center md:ml-2 md:w-2/12 xl:w-fit"
              @click="showDialog('addAgent')"
            >
              <UserPlusIcon class="mr-2 inline-block h-5 w-5" />
              {{ t('textAgentSettingAddAgent') }}
            </button>
            <button
              v-if="isEnabled.AgentAccountSecretKeyRead"
              type="button"
              class="btn btn-danger mb-1 flex items-center justify-center disabled:bg-gray-500 md:ml-2 md:w-2/12 xl:w-fit"
              :disabled="!selectRecord.id || !selectRecord.checked"
              @click="dialog.checkKeys = true"
            >
              <KeyIcon class="mr-2 inline-block h-5 w-5" />
              {{ t('textAgentSettingCheckKeys') }}
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
                  <th></th>
                  <PageTableSortableTh column="owner" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOpener') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('menuItemAgentAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="topAgentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textSuperiorName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="memberCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textMemberQuantity') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="ratio" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRatio') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    v-if="isAmountColumnEnabled"
                    column="amountStr"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textBalance') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="currency" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCurrency') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="type" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textState') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="role" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textGroupRole') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCreateDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="updateTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textLastUpdateTime') }} </template>
                  </PageTableSortableTh>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="isAmountColumnEnabled ? 14 : 13">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.code}`">
                    <td>
                      <CheckButton :checked="record.checked" @click="changeSelect(record)" />
                    </td>
                    <td>{{ record.owner }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.topAgentName }}</td>
                    <td>{{ record.memberCount }}</td>
                    <td>{{ `${record.ratio}%` }}</td>
                    <td v-if="isAmountColumnEnabled" :class="{ 'text-danger': record.amount > 0 }">
                      {{ record.amountStr }}
                    </td>
                    <td>{{ t(`currency__${record.currency}`) }}</td>
                    <td>{{ record.isEnabled === 1 ? t('textNormal') : t('textClose') }}</td>
                    <td>{{ record.role }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createTime) }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.updateTime) }}</td>
                    <td class="space-x-2">
                      <ViewButton @click="recordSetting(record, false)" />
                      <EditButton
                        v-if="isEnabled.AgentAccountUpdate && isAgentEditEnabled(record)"
                        @click="recordSetting(record, true)"
                      />
                      <ResetPasswordButton
                        v-if="isEnabled.ResetPassword && isAgentEditEnabled(record)"
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
  <AddAgent
    :visible="dialog.addAgent"
    @close="(newValue) => (dialog.addAgent = newValue)"
    @search-records="searchRecords()"
  />
  <CheckKeys
    :visible="dialog.checkKeys"
    :select-record="selectRecord"
    @close="(newValue) => (dialog.checkKeys = newValue)"
  />
  <Setting
    :visible="dialog.agentSet"
    :select-record="selectRecord"
    :is-edit-enabled="isSettingEditEnabled"
    @close="(newValue) => (dialog.agentSet = newValue)"
    @search-records="searchRecords()"
  />
  <ResetPassword
    :visible="dialog.resetPassword"
    :user-name="selectRecord.userName"
    :type="'AgentAccount'"
    @close="(newValue) => (dialog.resetPassword = newValue)"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { UserPlusIcon, KeyIcon } from '@heroicons/vue/20/solid'
import { useAgentAccount } from '@/base/composable/agentAccountManagement/agentAccount/useAgentAccount'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import ResetPasswordButton from '@/base/components/Button/ResetPasswordButton.vue'
import AddAgent from '@/base/views/AgentAccountManagement/AgentAccount/AgentAccountAddAgent.vue'
import CheckKeys from '@/base/views/AgentAccountManagement/AgentAccount/AgentAccountCheckKeys.vue'
import Setting from '@/base/views/AgentAccountManagement/AgentAccount/AgentAccountSetting.vue'
import ResetPassword from '@/base/views/Common/Dialog/ResetPasswordDialog.vue'

const { t } = useI18n()
const {
  dialog,
  formInput,
  isAmountColumnEnabled,
  isEnabled,
  isSettingEditEnabled,
  records,
  selectRecord,
  tableInput,
  changeSelect,
  isAgentEditEnabled,
  showDialog,
  searchRecords,
  recordSetting,
  resetPassword,
} = useAgentAccount()
</script>
