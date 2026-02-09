<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" :include-all="false" />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textCategory') }}</label>
          <FormDropdown
            v-model="formInput.category"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(item) => t(item.nameKey)"
            :fmt-item-key="(item) => item.key"
            :items="dropdownCategory.items"
          />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textOperate') }}</label>
          <FormDropdown
            v-model="formInput.actionType"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(actionType) => t(`actionType__${actionType}`)"
            :fmt-item-key="(actionType) => `actionType__${actionType}`"
            :items="dropdownActionTypes"
          />
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textTime') }}</label>
          <div class="mb-1 flex w-full md:w-3/12">
            <FormDateTimeInput
              v-model="formInput.startTime"
              class="flex-1"
              :before-days="time.commonReportTimeBeforeDays"
              :minute-increment="time.earningReportTimeMinuteIncrement"
              set-max-date
            />
            <div class="border border-gray-200 bg-gray-300 py-2 px-3.5 text-center">
              {{ t('textTimeTo') }}
            </div>
            <FormDateTimeInput
              v-model="formInput.endTime"
              class="flex-1"
              calendar-align="right"
              :before-days="time.commonReportTimeBeforeDays"
              :minute-increment="time.earningReportTimeMinuteIncrement"
              set-max-date
            />
          </div>
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchBackendActionLogList()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageTable :records="records.items" :table-input="tableInput" :total-records="records.items.length">
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
                    <template #default> {{ t('textOperatingDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('menuItemBackendAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="actionType" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOperate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="log" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textOperatingLog') }} </template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="4">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.createTime }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ t(`actionType__${record.actionType}`) }}</td>
                    <td class="whitespace-pre text-left">
                      <div>{{ record.log.title }}</div>
                      <template v-if="record.log.detail">
                        <ButtonTooltips
                          :tips-text="record.showLogDetail ? t('textHideExpandedContent') : t('textShowHiddenContent')"
                          @click="record.showLogDetail = !record.showLogDetail"
                        >
                          <template #content>
                            <div class="py-1">
                              <span class="flex h-3 w-6 justify-center rounded-full bg-gray-200">
                                <img class="-my-1" src="@/base/assets/images/pic_3dot.png" />
                              </span>
                            </div>
                          </template>
                        </ButtonTooltips>
                        <div v-show="record.showLogDetail">{{ record.log.detail }}</div>
                      </template>
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
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useBackendActionLog } from '@/base/composable/operationManagement/backendActionLog/useBackendActionLog'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'

const { t } = useI18n()
const { dropdownActionTypes, dropdownCategory, formInput, tableInput, records, searchBackendActionLogList } =
  useBackendActionLog()
</script>
