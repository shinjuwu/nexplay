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
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textBetId') }}</label>
          <input
            v-model="formInput.betId"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportBetId')"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoundId') }}</label>
          <input
            v-model="formInput.logNumber"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportRoundId')"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPlayerAccount') }}</label>
          <input
            v-model="formInput.userName"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
          />
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textBetIssue') }}</label>
          <FormDropdown
            v-model="formInput.betslipStatus"
            class="mb-1 w-full md:w-3/12"
            :items="betslipStatuses"
            :fmt-item-key="(betslipStatus) => `betslipStatuses__${betslipStatus.value}`"
            :fmt-item-text="(betslipStatus) => t(`text${betslipStatus.name}Betslip`)"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }}</label>
          <div class="mb-1 flex w-full md:w-3/12">
            <FormDateTimeInput
              v-model="formInput.startTime"
              class="flex-1"
              :minute-increment="time.commonReportTimeMinuteIncrement"
              :before-days="time.winloseReportTimeBeforeDays"
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
              :before-days="time.winloseReportTimeBeforeDays"
              set-max-date
            />
          </div>
          <template v-if="isShowSingleWalletRelative">
            <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textSingleWalletId') }}</label>
            <input
              v-model="formInput.singleWalletId"
              type="text"
              class="form-input mb-1 md:w-3/12"
              :placeholder="t('placeHolderTextSingleWalletId')"
            />
          </template>
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoomId') }}</label>
          <input
            v-model="formInput.roomId"
            type="text"
            class="form-input mb-1 md:w-3/12"
            maxlength="7"
            :placeholder="t('placeHolderTextRoomId')"
          />
        </div>
        <div class="flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchRecords()"
          >
            {{ t('textSearch') }}
          </button>
          <button
            type="button"
            class="btn btn-info mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            :disabled="tableInput.agentId === undefined || tableInput.gameId === undefined"
            @click="downloadRecords()"
          >
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
      <div class="text-righ my-2 md:justify-between lg:flex">
        <div class="text-gray-500"></div>
        <div class="text-danger flex flex-wrap">
          <div v-if="isShowKillDiveRelative" class="mr-4">
            {{
              t('fmtTextWinLoseReportKillDiveInfo', [
                totalKilledRecordCount,
                totalDivedRecordCount,
                totalPlayerWinRecordCount,
              ])
            }}
          </div>
          <div>
            <span>{{ t('fmtTextTotalBetInfo', [numberToStr(totalValidScore)]) }}</span
            >,
            <span v-if="isShowJackpotRelative"
              >{{ t('fmtTextTotalJackpotInjectWaterScore', [numberToStr(totalJackpotInjectWaterScore, 4)]) }},</span
            >
            <span
              :class="{
                'text-danger': totalWinLoseScore > 0,
                'text-teal-500': totalWinLoseScore < 0,
              }"
            >
              {{ t('fmtTextTotalWinLoseInfo', [numberToStr(totalWinLoseScore)]) }}
            </span>
          </div>
        </div>
      </div>
      <PageServersideTable :records="records.items" :table-input="tableInput" @search="searchRecords">
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
                  <PageTableSortableTh column="betid" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBetId') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentname" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textAgentName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bettime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textCreateTime') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="username" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textPlayerAccount') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="gameid" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roomtype" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomType') }}</template>
                  </PageTableSortableTh>
                  <th>{{ t('textBeforeBetScore') }}</th>
                  <PageTableSortableTh column="yascore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textBet') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="validscore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textValidScore') }}</template>
                  </PageTableSortableTh>
                  <template v-if="isShowJackpotRelative">
                    <PageTableSortableTh
                      column="jpinjectwaterrate"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default>{{ t('textJPInjectWaterPercent') }}</template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="jpinjectwaterscore"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default>{{ t('textJPInjectWaterScore') }}</template>
                    </PageTableSortableTh>
                  </template>
                  <PageTableSortableTh column="descore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGamerWinScore') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameTax') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bonus" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textBonus') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winlosescore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGamerWinLoseScore') }}</template>
                  </PageTableSortableTh>
                  <th>{{ t('textSettleScore') }}</th>
                  <PageTableSortableTh column="roundid" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoundId') }}</template>
                  </PageTableSortableTh>
                  <template v-if="isShowKillDiveRelative">
                    <PageTableSortableTh column="killtype" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default>{{ t('textRiskControl') }}</template>
                    </PageTableSortableTh>
                    <th>{{ t('textKillProb') }}</th>
                    <th>{{ t('textKillLevel') }}</th>
                    <th>{{ t('textRealPlayers') }}</th>
                  </template>
                  <PageTableSortableTh
                    v-if="isShowSingleWalletRelative"
                    column="wallet_ledger_id"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default>{{ t('textSingleWalletId') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roomid" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomId') }}</template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="tableColSpan">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr
                    v-for="record in currentRecords"
                    :key="`record_${record.betId}`"
                    :class="{
                      'bg-red-400/20': record.agentId === -1 || !record.agentName || !record.userName,
                    }"
                  >
                    <td>{{ record.betId }}</td>
                    <td>{{ record.agentName || '-' }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.betTime) }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ t(`game__${record.gameId}`) }}</td>
                    <td>{{ t(`roomType__${roomTypeNameIndex(record.gameId, record.roomType)}`) }}</td>
                    <td>{{ numberToStr(record.startScore) }}</td>
                    <td>{{ numberToStr(record.yaScore) }}</td>
                    <td>{{ numberToStr(record.validScore) }}</td>
                    <template v-if="isShowJackpotRelative">
                      <td>{{ record.jpInjectRate * 100 + '%' }}</td>
                      <td>{{ numberToStr(record.jpInjectScore, 4) }}</td>
                    </template>
                    <td>{{ numberToStr(record.deScore) }}</td>
                    <td>{{ numberToStr(record.tax) }}</td>
                    <td>{{ numberToStr(record.bonus) }}</td>
                    <td>
                      <span
                        :class="{
                          'text-danger': record.winLoseAmount > 0,
                          'text-teal-500': record.winLoseAmount < 0,
                        }"
                      >
                        {{
                          record.winLoseAmount > 0
                            ? `+${numberToStr(record.winLoseAmount)}`
                            : numberToStr(record.winLoseAmount)
                        }}
                      </span>
                    </td>
                    <td>{{ numberToStr(record.endScore) }}</td>
                    <td>
                      <button
                        :class="{
                          'text-cyan-500 underline underline-offset-1': isRedirectToGameLogParseEnabled,
                        }"
                        :disabled="!isRedirectToGameLogParseEnabled"
                        @click="redirectToGameLogParse(record.roundId, record.userName)"
                      >
                        {{ record.roundId }}
                      </button>
                    </td>
                    <template v-if="isShowKillDiveRelative">
                      <td>
                        <span
                          :class="{
                            'text-danger': record.killType !== 0 && record.killType !== 2,
                            'text-teal-500': record.killType === 2,
                          }"
                        >
                          {{ t(`killType__${record.killType}`) }}
                        </span>
                      </td>
                      <td>{{ record.killProb }}</td>
                      <td>{{ t(`killLevel__${record.killLevel}`) }}</td>
                      <td>{{ record.realPlayers }}</td>
                    </template>
                    <td v-if="isShowSingleWalletRelative">{{ record.walletLedgerId }}</td>
                    <td>{{ record.roomId }}</td>
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
import { useWinLoseReport } from '@/base/composable/reportManagement/winLoseReport/useWinLoseReport'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'
import { roomTypeNameIndex } from '@/base/utils/room'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormAllGameListDropdown from '@/base/components/Form/Dropdown/FormAllGameListDropdown.vue'
import FormRoomTypeListDropdown from '@/base/components/Form/Dropdown/FormRoomTypeListDropdown.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageServersideTable from '@/base/components/Page/Table/PageServersideTable.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'

const { t } = useI18n()
const {
  betslipStatuses,
  calDirections,
  formInput,
  isShowJackpotRelative,
  isShowKillDiveRelative,
  isShowSingleWalletRelative,
  isRedirectToGameLogParseEnabled,
  records,
  tableInput,
  totalDivedRecordCount,
  totalKilledRecordCount,
  totalPlayerWinRecordCount,
  totalValidScore,
  totalWinLoseScore,
  totalJackpotInjectWaterScore,
  tableColSpan,
  downloadRecords,
  redirectToGameLogParse,
  searchRecords,
} = useWinLoseReport()
</script>
