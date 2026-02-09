<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }} </label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textOrderNumber') }} </label>
          <input
            v-model="formInput.id"
            type="text"
            class="form-input mb-1 w-full md:w-3/12"
            :placeholder="t('placeHolderTextReportOrderNumber')"
          />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPlayerAccount') }} </label>
          <input
            v-model="formInput.userName"
            type="text"
            class="form-input mb-1 w-full md:w-3/12"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
          />
        </div>

        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textKind') }} </label>
          <FormDropdown
            v-model="formInput.kind"
            :fmt-item-text="
              (kind) => (kind === constant.WalletLedgerKind.All ? t('textKindAll') : t(`walletLedgerKind__${kind}`))
            "
            class="mb-1 w-full md:w-3/12"
            :fmt-item-key="(kind) => `kind__${kind}`"
            :items="kinds"
            :use-font-awsome="false"
          />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }} </label>
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

          <template v-if="isShowSingleWalletRelative">
            <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textSingleWalletId') }} </label>
            <input
              v-model="formInput.singleWalletId"
              type="text"
              class="form-input mb-1 w-full md:w-3/12"
              :placeholder="t('placeHolderTextSingleWalletId')"
            />
          </template>
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
      <div class="my-2 md:flex md:justify-end">
        <div class="text-danger">
          {{ t('fmtTextWalletLedgerReportTotalInfo', [numberToStr(totalUpScore), numberToStr(totalDownScore)]) }}
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
                  <PageTableSortableTh column="betId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOrderNumber') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="account" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCreateTime') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="kind" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textKind') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="coinAmountSort"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textCoinAmount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="beforeScoreSort"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textBeforeScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="changeScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textChangeScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="afterScoreSort"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textAfterScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="operator" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOperator') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="status" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOrderState') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="errorCode" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textErrorMessage') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="remark" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRemark') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    v-if="isShowSingleWalletRelative"
                    column="singleWalletId"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textSingleWalletId') }} </template>
                  </PageTableSortableTh>
                  <!-- <th>{{ t('textOperate') }}</th> -->
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="isShowSingleWalletRelative ? 14 : 13">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.betId}`">
                    <td>{{ record.betId }}</td>
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.account }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createTime) }}</td>
                    <td>{{ t(`walletLedgerKind__${record.kind}`) }}</td>
                    <td>{{ record.coinAmount }}</td>
                    <td>{{ record.beforeScore }}</td>
                    <td>
                      {{
                        record.changeScore > 0 ? `+${numberToStr(record.changeScore)}` : numberToStr(record.changeScore)
                      }}
                    </td>
                    <td>{{ record.afterScore }}</td>
                    <td>{{ record.operator }}</td>
                    <td>{{ t(`orderType__${record.status}`) }}</td>
                    <td>{{ t(`errorCode__${record.errorCode}`) }}</td>
                    <td>{{ record.remark }}</td>
                    <td v-if="isShowSingleWalletRelative">{{ record.singleWalletId }}</td>
                    <!-- <td>
                      <ButtonTooltips :tips-text="record.statusTips">
                        <template #content>
                          <CheckButton :checked="record.status === 1" @click="confirm(record)" />
                        </template>
                      </ButtonTooltips>
                    </td> -->
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
import { useWalletLedger } from '@/base/composable/reportManagement/walletLedger/useWalletLedger'
import constant from '@/base/common/constant'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
// import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'
// import CheckButton from '@/base/components/Button/CheckButton.vue'

const { t } = useI18n()
const {
  isShowSingleWalletRelative,
  kinds,
  records,
  totalUpScore,
  totalDownScore,
  formInput,
  tableInput,
  searchRecords,
  // confirm,
  downloadRecords,
} = useWalletLedger()
</script>
