<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textDeviceType') }}</label>
          <FormDropdown
            v-model="formInput.device"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(item) => t(`device_${item}`)"
            :fmt-item-key="(item) => `device_${item}`"
            :items="deviceList"
          />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textBrowserType') }}</label>
          <FormDropdown
            v-model="formInput.browser"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(item) => (item === constant.Browser.All ? t(`browser_${item}`) : item)"
            :fmt-item-key="(item) => `browser_${item}`"
            :items="browserList.items"
          />
        </div>

        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textSourceLocation') }}</label>
          <FormDropdown
            v-model="formInput.country"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(item) => (item === constant.Country.All ? t(`country_${item}`) : item)"
            :fmt-item-key="(item) => `country_${item}`"
            :items="countryList.items"
          />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPlayerAccount') }}</label>
          <input
            v-model="formInput.userName"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
          />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textIpAddress') }}</label>
          <input
            v-model="formInput.ip"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextIpAddress')"
          />
        </div>

        <div class="flex flex-wrap items-center">
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
        </div>

        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchPlayerLoginInfo()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>

      <hr class="my-2" />

      <PageTable
        :records="playerLoginInfo.items"
        :total-records="playerLoginInfo.items.length"
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
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="deviceName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textDeviceType') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="osName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textSystemVersion') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="browser" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBrowserType') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="browserVer" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBrowserVersion') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="resolutionName"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textScreenResolution') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="loginTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textLoginDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="ip" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textIpAddress') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="location" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textSourceLocation') }} </template>
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
                  <tr v-for="record in currentRecords" :key="`record_${record.token}_${record.loginTime}`">
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ record.deviceName }}</td>
                    <td>{{ record.osName }}</td>
                    <td>{{ record.browser }}</td>
                    <td>{{ record.browserVer }}</td>
                    <td>{{ record.resolutionName }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.loginTime) }}</td>
                    <td>{{ record.ip }}</td>
                    <td>{{ record.location }}</td>
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
import { usePlayerLoginInfo } from '@/base/composable/operationManagement/playerLoginInfo/usePlayerLoginInfo'
import constant from '@/base/common/constant'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'

const { t } = useI18n()
const { deviceList, browserList, countryList, formInput, tableInput, playerLoginInfo, searchPlayerLoginInfo } =
  usePlayerLoginInfo()
</script>
