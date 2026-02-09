<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPlayerAccount') }}</label>
          <input
            v-model="formInput.userName"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoundId') }}</label>
          <input
            v-model="formInput.roundId"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportRoundId')"
          />
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textJPTokenId') }}</label>
          <input
            v-model="formInput.tokenId"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextJackpotTokenId')"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }}</label>
          <div class="mb-1 flex w-full md:w-3/12">
            <FormDateTimeInput
              v-model="formInput.startTime"
              class="flex-1"
              :minute-increment="time.commonReportTimeMinuteIncrement"
              :before-days="time.commonReportTimeBeforeDays"
              set-max-date
            />
            <div class="border border-gray-200 bg-gray-300 py-2 px-3.5 text-center">
              {{ t('textTimeTo') }}
            </div>
            <FormDateTimeInput
              v-model="formInput.endTime"
              class="flex-1"
              calendar-align="right"
              :minute-increment="time.commonReportTimeMinuteIncrement"
              :before-days="time.commonReportTimeBeforeDays"
              set-max-date
            />
          </div>
        </div>
        <div
          :class="[isSettingEditEnabled ? 'justify-between' : 'justify-end']"
          class="mt-4 flex flex-wrap items-center"
        >
          <button
            v-if="isSettingEditEnabled"
            type="button"
            class="btn btn-danger flex items-center"
            @click="showTokenDialog = true"
          >
            <PlusCircleIcon class="mr-2 inline-block h-6 w-6" />
            {{ t('textAddToken') }}
          </button>
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="getJackpotTokenList()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageTable
        :records="jackpotTokenList.items"
        :total-records="jackpotTokenList.items.length"
        :table-input="tableInput"
      >
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
                  <PageTableSortableTh column="roundId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRoundId') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tokenId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJPTokenId') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tokenGetTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJPTokenGetTime') }} </template>
                  </PageTableSortableTh>
                  <!-- <PageTableSortableTh column="jackpotBet" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJPTokenBet') }} </template>
                  </PageTableSortableTh> -->
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="operator" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOperator') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="orderState" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOrderState') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="errorMsg" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textErrorMessage') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="remark" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRemark') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="isUsed" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textIsUsed') }} </template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="10">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.roundId || '-' }}</td>
                    <td>{{ record.tokenId }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.tokenGetTime) }}</td>
                    <!-- <td>{{ numberToStr(record.jackpotBet) }}</td> -->
                    <td>{{ record.agentName || '-' }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ record.operator || '-' }}</td>
                    <td>{{ t(`orderType__${record.orderState}`) }}</td>
                    <td>{{ t(`errorCode__${record.errorCode}`) }}</td>
                    <td>{{ record.remark || '-' }}</td>
                    <td>{{ record.isUsed > 0 ? t('textYes') : t('textNo') }}</td>
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
  <PageTableDialog v-if="showTokenDialog" :visible="showTokenDialog" @close="closeDialog()">
    <template #header>
      {{ t('textAddToken') }}
    </template>
    <template #default>
      <div class="grid grid-cols-12 gap-2">
        <div class="col-span-6 flex items-center">
          <label class="text-danger required-mark mr-2 w-fit">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="tokenForm.agent" :include-all="false" class="mb-1 w-full md:w-3/4" />
        </div>
        <div class="col-span-6">
          <label class="text-danger required-mark mr-2 w-fit">{{ t('textPlayerAccount') }}</label>
          <input
            v-model="tokenForm.userName"
            type="text"
            class="form-input mb-1 w-full md:w-3/4"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
          />
        </div>
        <!-- <div class="col-span-6">
          <label class="text-danger required-mark mr-2 w-fit">{{ t('textCoinQuota') }}</label>
          <DecimalNumberInput v-model:model-value="tokenForm.quota" :align="'right'" class="w-full md:w-11/12" />
        </div> -->
        <div class="col-span-12">
          <label class="text-danger required-mark">{{ t('textRemark') }}</label>
          <textarea
            v-model="tokenForm.remark"
            rows="5"
            maxlength="100"
            :placeholder="t('placeHolderTextAreaLimitNoUndo')"
            class="block w-full resize-none rounded border border-gray-400 p-2 outline-none placeholder:text-gray-400 focus-visible:border-gray-300"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="closeDialog()">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="showTokenDialog"
          :parent-data="tokenForm"
          :button-click="
            () => {
              addToToken()
            }
          "
        >
          {{ t('textSend') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>
<script setup>
import { useI18n } from 'vue-i18n'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
// import DecimalNumberInput from '@/base/components/NumberInput/DecimalNumberInput.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

import time from '@/base/utils/time'
// import { numberToStr } from '@/base/utils/formatNumber'
import { PlusCircleIcon } from '@heroicons/vue/24/solid'
import { useJackpotTokenRecord } from '@/base/composable/jackpotManagement/useJackpotTokenRecord'

const { t } = useI18n()
const {
  tableInput,
  formInput,
  jackpotTokenList,
  showTokenDialog,
  tokenForm,
  isSettingEditEnabled,
  getJackpotTokenList,
  addToToken,
  closeDialog,
} = useJackpotTokenRecord()
</script>
