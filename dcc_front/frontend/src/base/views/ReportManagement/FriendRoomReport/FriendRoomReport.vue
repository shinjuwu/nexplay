<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textGame') }}</label>
          <FormGameListDropdown
            v-model="formInput.game"
            class="mb-1 w-full md:w-3/12"
            :game-type="constant.GameType.FriendsRoom"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoomId') }}</label>
          <input
            v-model="formInput.roomId"
            type="text"
            class="form-input mb-1 md:w-3/12"
            maxlength="7"
            :placeholder="t('placeHolderTextRoomId')"
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
            @click="searchRecords()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>

      <hr class="my-2" />

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
                  <PageTableSortableTh column="id" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOrderNumber') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textAgentName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textPlayerAccount') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="gameId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roomId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomId') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalGameTax') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="taxpercent" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameTaxpercent') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomCreateTime') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="endTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomEndTime') }}</template>
                  </PageTableSortableTh>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="10">{{ t('textTableEmpty') }}</td>
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
                    <td>{{ record.id }}</td>
                    <td>{{ record.agentName || '-' }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ t(`game__${record.gameId}`) }}</td>
                    <td>{{ record.roomId }}</td>
                    <td>{{ numberToStr(record.tax) }}</td>
                    <td>{{ record.taxpercent * 100 + '%' }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createTime) }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.endTime) }}</td>
                    <td>
                      <ViewButton @click="showFriendRoomDetail(record)" />
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

      <FriendRoomDetail
        :visible="isShowFriendRoomDetail"
        :friend-room-info="friendRoomInfo"
        @close="closeFriendRoomDetail()"
      />
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useFriendRoomReport } from '@/base/composable/reportManagement/friendRoomReport/useFriendRoomReport'
import constant from '@/base/common/constant'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormGameListDropdown from '@/base/components/Form/Dropdown/FormGameListDropdown.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import FriendRoomDetail from '@/base/views/ReportManagement/FriendRoomReport/FriendRoomDetail.vue'

const { t } = useI18n()
const {
  formInput,
  tableInput,
  records,
  searchRecords,
  isShowFriendRoomDetail,
  friendRoomInfo,
  showFriendRoomDetail,
  closeFriendRoomDetail,
} = useFriendRoomReport()
</script>
