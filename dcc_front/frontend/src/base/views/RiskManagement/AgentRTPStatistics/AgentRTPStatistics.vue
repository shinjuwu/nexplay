<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textGame') }}</label>
          <FormAllGameListDropdown v-model="formInput.game" class="mb-1 w-full md:w-3/12" :agent-rtp="true" />
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
                :minute-increment="time.earningReportTimeMinuteIncrement"
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
                :minute-increment="time.earningReportTimeMinuteIncrement"
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
        </div>
      </form>
      <hr class="my-2" />
      <ToggleHeader :tips-title="t('textCalDirections')">
        <template #tipsSlot="{ tipsShow }">
          <div
            v-for="(direction, dirIdx) in calDirections.items"
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
        {{
          t('fmtTextAgentRTPStatisticsInfo', [
            numberToStr(recordsInfo.totalGamerBet),
            numberToStr(recordsInfo.totalGamerWinScore),
            numberToStr(recordsInfo.totalTax),
            numberToStr(recordsInfo.totalBonus),
            numberToStr(recordsInfo.avgRtp),
            numberToStr(recordsInfo.avgDeductTaxRtp),
          ])
        }}
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
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textAgentName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="gameType" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameCategory') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="gameId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roomType" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomType') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="lastRTPSetting"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default>{{ t('textRTP') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="betCounts" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalBetCounts') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="yaScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalPlayerBet') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="deScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalPlayerWinLose') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalTax') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bonus" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textSumTotalBonus') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="rtp" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>RTP</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="deductTaxRtp" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textDeductTaxRtp') }}</template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="12">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.agentName }}</td>
                    <td>{{ t(`gameType__${record.gameType}`) }}</td>
                    <td>{{ t(`game__${record.gameId}`) }}</td>
                    <td>{{ t(`roomType__${roomTypeNameIndex(record.gameId, record.roomType)}`) }}</td>
                    <td>{{ numberToStr(record.lastRTPSetting, 1) + '%' }}</td>
                    <td>{{ record.betCounts.toLocaleString() }}</td>
                    <td>{{ numberToStr(record.yaScore) }}</td>
                    <td>{{ numberToStr(record.deScore) }}</td>
                    <td>{{ numberToStr(record.tax) }}</td>
                    <td>{{ numberToStr(record.bonus) }}</td>
                    <td>{{ `${numberToStr(record.rtp)}%` }}</td>
                    <td>{{ `${numberToStr(record.deductTaxRtp)}%` }}</td>
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
import { useAgentRTPStatistics } from '@/base/composable/riskManagement/agentRTPStatistics/useAgentRTPStatistics'
import { numberToStr } from '@/base/utils/formatNumber'
import { roomTypeNameIndex } from '@/base/utils/room'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormAllGameListDropdown from '@/base/components/Form/Dropdown/FormAllGameListDropdown.vue'
import FormRoomTypeListDropdown from '@/base/components/Form/Dropdown/FormRoomTypeListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'

const { t } = useI18n()
const { calDirections, formInput, records, recordsInfo, tableInput, searchRecords } = useAgentRTPStatistics()
</script>
