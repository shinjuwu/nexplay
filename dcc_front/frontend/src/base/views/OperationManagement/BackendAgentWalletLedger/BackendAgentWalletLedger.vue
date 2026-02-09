<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown
            v-model="formInput.agent"
            class="mb-1 w-full md:w-3/12"
            :include-self="false"
            :include-grandson="false"
            :agent-cooperation="constant.AgentCooperation.BuyPoint"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
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
      <div class="text-danger text-right">
        <span v-if="!isAdmin">
          {{ t('fmtTextRemainingScore', [numberToStr(agentBalance)]) }}
        </span>
        <span v-else>{{ t('fmtTextRemainingScore', ['-']) }}</span>
      </div>
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
                  <PageTableSortableTh column="adminUser" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('menuItemAgentAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="name" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="ratio" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRatio') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="balance" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBalance') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="state" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textState') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCreateDate') }} </template>
                  </PageTableSortableTh>
                  <th v-if="isSettingEnabled">
                    {{ t('textOperate') }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="isSettingEnabled ? 7 : 6">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.adminUser }}</td>
                    <td>{{ record.name }}</td>
                    <td>{{ record.ratio + '%' }}</td>
                    <td>{{ numberToStr(record.balance) }}</td>
                    <td>{{ record.state === 1 ? t('textOpened') : t('textDisabled') }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createTime) }}</td>
                    <td v-if="isSettingEnabled">
                      <ArrowUpCircleIcon
                        class="text-success mr-2 inline h-6 w-6 cursor-pointer rounded-full"
                        @click="editRecordScore('up', record)"
                      />
                      <ArrowDownCircleIcon
                        class="text-danger inline h-6 w-6 cursor-pointer rounded-full"
                        @click="editRecordScore('down', record)"
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
  <TableDialog :visible="showEditScoreDialog" @close="showEditScoreDialog = false">
    <template #header> {{ dialogContent.title }} </template>
    <template #default>
      <div class="flex flex-wrap justify-between">
        <label class="form-label my-2 w-full md:w-1/2">{{ t('fmtTextAgentSettingAgentId', [agentInfo.id]) }}</label>
        <label class="form-label my-2 w-full md:w-1/2">
          {{ t('fmtTextCurrentBalance', [numberToStr(agentInfo.balance)]) }}
        </label>
        <label class="form-label my-2 w-full">{{ t('fmtTextAgentName', [agentInfo.name]) }}</label>
        <div class="flex w-full flex-col md:w-1/2">
          <label class="form-label my-2 w-full text-red-500 after:ml-0.5 after:content-['*'] md:w-1/2">
            {{ dialogContent.amount }}
          </label>
          <DecimalNumberInput
            class="w-full md:w-11/12"
            :model-value="agentInfo.amount"
            :align="'right'"
            @update:model-value="(newValue) => (agentInfo.amount = newValue)"
          />
        </div>
        <div class="flex w-full flex-col md:w-1/2">
          <label class="form-label my-2 w-full md:w-1/2">{{ dialogContent.afterAmount }}</label>
          <DecimalNumberInput
            disabled
            :align="'right'"
            :model-value="agentInfo.afterAmount"
            class="w-full md:w-11/12"
          />
        </div>
        <div class="flex w-full flex-col">
          <label class="form-label my-2 w-full text-red-500 after:content-['*'] md:w-1/2">{{ t('textRemark') }}</label>
          <textarea
            v-model="agentInfo.remark"
            rows="5"
            maxlength="100"
            :placeholder="t('placeHolderTextAreaLimitNoUndo')"
            class="resize-none rounded border border-gray-400 p-2 outline-none placeholder:text-gray-400 focus-visible:border-gray-300"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="showEditScoreDialog = false">
          {{ t('textClose') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="showEditScoreDialog"
          :parent-data="agentInfo"
          :button-click="() => setAgentWallet()"
        >
          {{ dialogContent.btn }}
        </LoadingButton>
      </div>
    </template>
  </TableDialog>
</template>

<script setup>
import { useBackendAgentWalletLedger } from '@/base/composable/operationManagement/backendAgentWalletLedger/useBackendAgentWalletLedger'
import { useI18n } from 'vue-i18n'
import { ArrowUpCircleIcon, ArrowDownCircleIcon } from '@heroicons/vue/24/outline'
import constant from '@/base/common/constant'
import time from '@/base/utils/time'
import { numberToStr } from '@/base/utils/formatNumber'

import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import DecimalNumberInput from '@/base/components/NumberInput/DecimalNumberInput.vue'
import TableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const {
  agentBalance,
  agentInfo,
  dialogContent,
  formInput,
  isAdmin,
  isSettingEnabled,
  records,
  showEditScoreDialog,
  tableInput,
  editRecordScore,
  searchRecords,
  setAgentWallet,
} = useBackendAgentWalletLedger()
</script>
