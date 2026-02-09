<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show" @submit.prevent>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" :include-all="false" />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textBackendAccount') }}</label>
          <input v-model="formInput.userName" type="text" class="form-input mb-1 w-full md:w-3/12" />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }}</label>
          <div class="mb-1 flex w-full md:w-3/12">
            <FormDateTimeInput
              v-model="formInput.startTime"
              class="flex-1"
              :minute-increment="time.earningReportTimeMinuteIncrement"
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
              :minute-increment="time.earningReportTimeMinuteIncrement"
              :before-days="time.commonReportTimeBeforeDays"
              set-max-date
            />
          </div>
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

      <PageTable :records="filterRecords" :total-records="filterRecords.length" :table-input="tableInput">
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
          <div class="flex items-center justify-end">
            <input
              v-model="tableInput.ip"
              type="text"
              class="form-input mb-1 md:w-4/12"
              :placeholder="t('placeHolderTextIpAddress')"
            />
          </div>
          <div class="tbl-container">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <PageTableSortableTh column="loginTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textLoginDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBackendAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="ip" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textIpAddress') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="status" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textLoginStatus') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="errorCode" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRemark') }} </template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr>
                    <td colspan="5">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="(record, index) in currentRecords" :key="`record_${index}`">
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.loginTime) }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ record.ip }}</td>
                    <td>{{ record.status ? t('textSuccess') : t('textFailure') }}</td>
                    <td>
                      {{ record.errorCode === constant.ErrorCode.Success ? '-' : t(`errorCode__${record.errorCode}`) }}
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
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useBackendLoginLog } from '@/base/composable/systemManagement/backendLoginLog/useBackendLoginLog'
import constant from '@/base/common/constant'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'

const { t } = useI18n()
const { filterRecords, formInput, tableInput, searchRecords } = useBackendLoginLog()
</script>
