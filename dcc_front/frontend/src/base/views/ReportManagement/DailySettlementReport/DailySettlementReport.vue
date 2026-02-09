<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textGame') }}</label>
          <FormAllGameListDropdown v-model="formInput.game" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoomType') }}</label>
          <FormRoomTypeListDropdown v-model="formInput.roomType" class="mb-1 w-full md:w-3/12" />
        </div>
        <div class="flex flex-wrap items-center">
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
                :display-times="false"
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
                :display-times="false"
              />
            </div>
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
          <button type="button" class="btn btn-info mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12" @click="downloadRecords()">
            {{ t('textExportExcel') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <ToggleHeader :tips-title="t('textCalDirections')">
        <template #tipsSlot="{ tipsShow }">
          <div
            v-for="(direction, dirIdx) in calDirections"
            v-show="tipsShow"
            :key="`directions__${dirIdx}`"
            class="text-danger"
          >
            {{ direction }}
          </div>
        </template>
      </ToggleHeader>
      <hr class="my-2" />
      <div class="text-danger my-2 text-right">
        <span
          :class="{
            'text-danger': sumScore > 0,
            'text-teal-500': sumScore < 0,
          }"
          >{{ t('fmtTextTotalWinLoseInfo', [numberToStr(sumScore)]) }}</span
        >
      </div>
      <PageServersideTable :records="agentDailyStat.items" :table-input="tableInput" @search="searchRecords">
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
                  <PageTableSortableTh column="agentname" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgent') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="gamename" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textGameName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roomtype" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRoomType') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="playcount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBetCounts') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="yascore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBetTotal') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="descore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalGamerWinScore') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalGameTax') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bonus" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalBonus') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="playerwinlose"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textTotalGamerWinLoseScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="taxedrtp" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textDeductTaxRtp') }} </template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="13">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="(record, index) in currentRecords" :key="`sumRecord_${index}`">
                    <td>{{ record.agentName }}</td>
                    <td>{{ t(`game__${record.gameId}`) }}</td>
                    <td>{{ t(`roomType__${roomTypeNameIndex(record.gameId, record.roomType)}`) }}</td>
                    <td>{{ record.betCounts }}</td>
                    <td>{{ numberToStr(record.bet) }}</td>
                    <td>{{ numberToStr(record.score) }}</td>
                    <td>{{ numberToStr(record.tax) }}</td>
                    <td>{{ numberToStr(record.bonus) }}</td>
                    <td :class="{ 'text-danger': record.winLoseScore > 0, 'text-success': record.winLoseScore < 0 }">
                      {{ numberToStr(record.winLoseScore) }}
                    </td>
                    <td>{{ `${numberToStr(record.deductTaxRtp * 100)}%` }}</td>
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
      </PageServersideTable>
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useDailySettlementReport } from '@/base/composable/reportManagement/dailySettlementReport/useDailySettlementReport'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'
import { roomTypeNameIndex } from '@/base/utils/room'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import FormAllGameListDropdown from '@/base/components/Form/Dropdown/FormAllGameListDropdown.vue'
import FormRoomTypeListDropdown from '@/base/components/Form/Dropdown/FormRoomTypeListDropdown.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageServersideTable from '@/base/components/Page/Table/PageServersideTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'

const { t } = useI18n()
const { formInput, tableInput, agentDailyStat, sumScore, calDirections, downloadRecords, searchRecords } =
  useDailySettlementReport()
</script>
