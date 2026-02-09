<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textJPTokenId') }}</label>
          <input
            v-model="formInput.tokenId"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextJackpotTokenId')"
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
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPlayerAccount') }}</label>
          <input
            v-model="formInput.userName"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
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
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="getJackpotPrizeRecords()"
          >
            {{ t('textSearch') }}
          </button>
          <button type="button" class="btn btn-info mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12" @click="downloadRecords()">
            {{ t('textExportExcel') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <div class="text-danger text-right">
        <span>{{ t('fmtTextTotalPrizeScore', [numberToStr(totalWinningScore)]) }}</span>
      </div>
      <PageTable
        :records="jackpotPrizeList.items"
        :total-records="jackpotPrizeList.items.length"
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
                  <PageTableSortableTh column="orderId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOrderNumber') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roundId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRoundId') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winningTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textWinningTime') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgent') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tokenId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJPTokenId') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tokenGetTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJPTokenGetTime') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winningItem" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textWinningItem') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winningScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJPWinningScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    v-if="isAdmin"
                    column="isBOT"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> BOT </template>
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
                    <td>{{ record.orderId }}</td>
                    <td>{{ record.roundId }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.winningTime) }}</td>
                    <td>{{ record.isBOT === 1 ? '-' : record.agentName }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ record.tokenId }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.tokenGetTime) }}</td>
                    <td>{{ t(`jackpotWinningItem__${record.winningItem}`) }}</td>
                    <td>{{ numberToStr(record.winningScore) }}</td>
                    <td v-if="isAdmin">{{ record.isBOT === 1 ? t('textYes') : t('textNo') }}</td>
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

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import { useJackpotPrizeRecord } from '@/base/composable/jackpotManagement/useJackpotPrizeRecord'
import time from '@/base/utils/time'
import { numberToStr } from '@/base/utils/formatNumber'

const { t } = useI18n()
const { tableInput, formInput, jackpotPrizeList, isAdmin, totalWinningScore, getJackpotPrizeRecords, downloadRecords } =
  useJackpotPrizeRecord()
</script>
