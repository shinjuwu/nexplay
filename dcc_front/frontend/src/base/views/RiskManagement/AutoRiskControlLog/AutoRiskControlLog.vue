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
          <label class="form-label mb-1 w-full self-start pr-2 md:mt-2 md:w-1/12 md:text-right">
            {{ t('textTime') }}
          </label>
          <div class="mb-1 flex w-full flex-col md:w-3/12">
            <div class="flex flex-1">
              <FormDateTimeInput
                v-model="formInput.startTime"
                class="flex-1"
                set-max-date
                :before-days="time.commonReportTimeBeforeDays"
                :minute-increment="time.commonReportTimeMinuteIncrement"
              />
              <div class="border border-gray-200 bg-gray-300 py-2 px-3.5 text-center">
                {{ t('textTimeTo') }}
              </div>
              <FormDateTimeInput
                v-model="formInput.endTime"
                class="flex-1"
                calendar-align="right"
                set-max-date
                :before-days="time.commonReportTimeBeforeDays"
                :minute-increment="time.commonReportTimeMinuteIncrement"
              />
            </div>
          </div>
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="getAutoRiskControlLogList()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageTable
        :records="autoRiskControlLog.items"
        :total-records="autoRiskControlLog.items.length"
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
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textDisposeTime') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textAgentName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textPlayerAccount') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="riskCode" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textAutoRisCondition') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="way" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textDisposeWay') }}</template>
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
                    <td>{{ record.createTime }}</td>
                    <td>{{ record.agentName }}</td>
                    <td>
                      <a
                        class="text-primary cursor-pointer underline"
                        @click="redirectToPlayerAccount(record.agentId, record.userName)"
                      >
                        {{ record.userName }}
                      </a>
                    </td>
                    <td>{{ t(`autoRiskSetting__${record.riskCode}`) }}</td>
                    <td>{{ riskDisposeWay[record.riskCode] }}</td>
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
import { useAutoRiskControlLog } from '@/base/composable/riskManagement/autoRiskControlLog/useAutoRiskControlLog'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'

const { t } = useI18n()
const {
  time,
  formInput,
  autoRiskControlLog,
  tableInput,
  riskDisposeWay,
  getAutoRiskControlLogList,
  redirectToPlayerAccount,
} = useAutoRiskControlLog()
</script>
