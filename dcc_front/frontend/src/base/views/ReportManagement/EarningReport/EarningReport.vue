<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown
            v-model="formInput.agent"
            class="mb-1 w-full md:w-3/12"
            :include-self="!isAdminUser"
            :include-grandson="false"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }}</label>
          <div class="mb-1 flex w-full md:w-3/12">
            <FormDateTimeInput
              v-model="formInput.startTime"
              class="flex-1"
              :before-days="time.commonReportTimeBeforeDays"
              set-max-date
              :display-times="false"
            />
            <div class="border border-gray-200 bg-gray-300 py-2 px-3.5 text-center">
              {{ t('textTimeTo') }}
            </div>
            <FormDateTimeInput
              v-model="formInput.endTime"
              class="flex-1"
              calendar-align="right"
              :before-days="time.commonReportTimeBeforeDays"
              set-max-date
              :display-times="false"
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
        <span>{{ t('fmtTextTotalBetInfo', [numberToStr(sumValidBet)]) }}, </span>
        <span
          :class="{
            'text-danger': sumScore > 0,
            'text-teal-500': sumScore < 0,
          }"
          >{{ t('fmtTextTotalWinLoseInfo', [numberToStr(sumScore)]) }}</span
        >,
        <template v-if="isShowJackpotRelative">
          <span>{{ t('fmtTextJPTotalInjectWaterScore', [numberToStr(sumJPInjectWaterScore, 4)]) }},</span>
          <span>{{ t('fmtTextJPTotalWinningScore', [numberToStr(sumJPWinningScore)]) }},</span>
        </template>
        <span>{{ t('fmtTextTotalAgentSettleUp', [numberToStr(sumAgentSettleUp)]) }}</span>
      </div>
      <PageTable :records="sumRecords.items" :total-records="sumRecords.items.length" :table-input="sumAgentTableInput">
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
                    <template #default> {{ t('textAgent') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="betCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBetCounts') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    v-if="isShowJackpotRelative"
                    column="jpBetCount"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textJPBetCount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bet" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBetTotal') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="validBet" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textValidBetTotal') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalGamerWinScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalGameTax') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bonus" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalBonus') }} </template>
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
                  <PageTableSortableTh column="ratio" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRatio') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="agentSettleUp"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textSettleUpScore') }} </template>
                  </PageTableSortableTh>
                  <template v-if="isShowJackpotRelative">
                    <PageTableSortableTh
                      column="totalJPInjectWaterScore"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default> {{ t('textTotalJPInjectWaterScore') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="totalJPlWinningScore"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default> {{ t('textTotalJPWinScore') }} </template>
                    </PageTableSortableTh>
                  </template>
                  <PageTableSortableTh column="currency" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCurrency') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="exchangeRate" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textExchangeRate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="exchangeCurrency"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default>{{ t('textAgentSettleUp') }}</template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="isShowJackpotRelative ? 17 : 14">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="(record, index) in currentRecords" :key="`sumRecord_${index}`">
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.betCount.toLocaleString() }}</td>
                    <td v-if="isShowJackpotRelative">{{ record.jpBetCount.toLocaleString() }}</td>
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
                    <td>{{ numberToStr(record.bonus) }}</td>
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
                    <td>{{ record.ratio + '%' }}</td>
                    <td>
                      <span
                        :class="{
                          'text-danger': record.agentSettleUp > 0,
                          'text-teal-500': record.agentSettleUp < 0,
                        }"
                      >
                        {{
                          record.agentSettleUp > 0
                            ? `+${numberToStr(record.agentSettleUp)}`
                            : numberToStr(record.agentSettleUp)
                        }}
                      </span>
                    </td>
                    <template v-if="isShowJackpotRelative">
                      <td>{{ numberToStr(record.totalJPInjectWaterScore, 4) }}</td>
                      <td>{{ numberToStr(record.totalJPlWinningScore) }}</td>
                    </template>
                    <td>
                      {{ t(`currency__${record.currency}`) }}
                    </td>
                    <td>
                      {{ record.toCoin }}
                    </td>
                    <td>{{ numberToStr(record.exchangeCurrency) }}</td>
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
import { useEarningReport } from '@/base/composable/reportManagement/earningReport/useEarningReport'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'

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
  calDirections,
  formInput,
  isAdminUser,
  isShowJackpotRelative,
  sumAgentTableInput,
  sumValidBet,
  sumRecords,
  sumScore,
  sumAgentSettleUp,
  sumJPInjectWaterScore,
  sumJPWinningScore,
  downloadRecords,
  searchRecords,
} = useEarningReport()
</script>
