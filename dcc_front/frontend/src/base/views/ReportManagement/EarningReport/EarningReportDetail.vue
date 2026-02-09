<template>
  <PageTableDialog :visible="show" size="xl" @close="close()">
    <template #header>{{ t('textDetailInfo') }}</template>
    <template #default>
      <div class="max-h-[500px] overflow-y-auto lg:max-h-[680px]">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }}</label>
          <div class="mb-1 flex w-full md:w-3/12">
            <div class="form-input whitespace-nowrap rounded-r-none">
              {{ format(props.startTime, 'yyyy-MM-dd HH:mm') }}
            </div>
            <div class="border-gray-200 bg-gray-300 py-2 px-3.5 text-center">
              {{ t('textTimeTo') }}
            </div>
            <div class="form-input whitespace-nowrap rounded-l-none">
              {{ format(props.endTime, 'yyyy-MM-dd HH:mm') }}
            </div>
          </div>
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
                    <PageTableSortableTh column="date" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textDate') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textAgent') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="betCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textBetCounts') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="bet" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textBetTotal') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="winScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textValidBetTotal') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="winScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textTotalGamerWinScore') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textTotalGameTax') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="gamerWinLoseScore"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default> {{ t('textTotalGamerWinLoseScore') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="agentWinLoseScore"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default> {{ t('textTotalAgentWinLoseScore') }} </template>
                    </PageTableSortableTh>
                  </tr>
                </thead>
                <tbody>
                  <template v-if="currentRecords.length === 0">
                    <tr>
                      <td colspan="8">{{ t('textTableEmpty') }}</td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                      <td>{{ time.utcTimeStrNoneISO8601ToLocalTimeFormat(record.date) }}</td>
                      <td>{{ record.agentName }}</td>
                      <td>{{ record.betCount.toLocaleString() }}</td>
                      <td>{{ numberToStr(record.bet) }}</td>
                      <td>{{ numberToStr(record.validBet) }}</td>
                      <td>
                        <span
                          :class="{
                            'text-danger': record.winScore > 0,
                            'text-teal-500': record.winScore < 0,
                          }"
                        >
                          {{ record.winScore > 0 ? `+${numberToStr(record.winScore)}` : numberToStr(record.winScore) }}
                        </span>
                      </td>
                      <td>{{ numberToStr(record.tax) }}</td>
                      <td>
                        <span
                          :class="{
                            'text-danger': record.gamerWinLoseScore > 0,
                            'text-teal-500': record.gamerWinLoseScore < 0,
                          }"
                        >
                          {{
                            record.gamerWinLoseScore > 0
                              ? `+${numberToStr(record.gamerWinLoseScore)}`
                              : numberToStr(record.gamerWinLoseScore)
                          }}
                        </span>
                      </td>
                      <td>
                        <span
                          :class="{
                            'text-danger': record.agentWinLoseScore > 0,
                            'text-teal-500': record.agentWinLoseScore < 0,
                          }"
                        >
                          {{
                            record.agentWinLoseScore > 0
                              ? `+${numberToStr(record.agentWinLoseScore)}`
                              : numberToStr(record.agentWinLoseScore)
                          }}
                        </span>
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
      </div>
    </template>
    <template #footer>
      <div class="ml-auto space-x-2">
        <button type="button" class="btn btn-light" @click="close()">
          {{ t('textClose') }}
        </button>
        <button type="button" class="btn btn-info" @click="downloadRecords()">
          {{ t('textExportExcel') }}
        </button>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { format } from 'date-fns'
import { useEarningReportDetail } from '@/base/composable/reportManagement/earningReport/useEarningReportDetail'
import { numberToStr } from '@/base/utils/formatNumber'
import { useI18n } from 'vue-i18n'
import time from '@/base/utils/time'

import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  agentId: {
    type: Number,
    default: 0,
  },
  startTime: {
    type: Date,
    default: () => new Date(),
  },
  endTime: {
    type: Date,
    default: () => new Date(),
  },
})
const emit = defineEmits(['close'])

const { t } = useI18n()
const { records, show, tableInput, close, downloadRecords } = useEarningReportDetail(props, emit)
</script>
