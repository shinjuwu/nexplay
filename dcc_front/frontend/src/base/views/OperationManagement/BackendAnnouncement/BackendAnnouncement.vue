<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPostType') }}</label>
          <FormDropdown
            v-model="formInput.type"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(type) => (type === constant.AnnounceType.All ? t('textTypeAll') : t(`typeCode__${type}`))"
            :fmt-item-key="(type) => `type__${type}`"
            :items="types"
            :use-font-awsome="false"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-between">
          <div class="flex-1">
            <button
              v-if="isEnabled.BackendAnnouncementCreate"
              type="button"
              class="btn btn-danger flex"
              @click="createNewAnnounce"
            >
              <CalendarIcon class="mr-2 h-6 w-6" />
              {{ t('roleItemBackendAnnouncementCreate') }}
            </button>
          </div>
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
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCreatingDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="updateTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textUpdateDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="type" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPostType') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="subject" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textSubject') }} </template>
                  </PageTableSortableTh>
                  <th>
                    {{ t('textOperate') }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="7">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createTime) }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.updateTime) }}</td>
                    <td>{{ t(`typeCode__${record.type}`) }}</td>
                    <td>{{ record.subject }}</td>
                    <td class="space-x-2">
                      <ViewButton @click="getAnnouncementInfo(record.id, 'view')" />
                      <EditButton
                        v-if="isEnabled.BackendAnnouncementUpdate"
                        @click="getAnnouncementInfo(record.id, 'update')"
                      />
                      <DeleteButton
                        v-if="isEnabled.BackendAnnouncementDelete"
                        @click="getAnnouncementInfo(record.id, 'delete')"
                      />
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
  <BackendAnnouncementComponent
    :announcement-info="announcementInfo"
    :visible="showDialog"
    :mode="showMode"
    @refresh-announcement-list="searchRecords"
    @close="showDialog = false"
  />
  <PageTableDialog :visible="deleteDialog" @close="deleteDialog = false">
    <template #header>{{ t('roleItemBackendAnnouncementDelete') }} </template>
    <template #default>
      <div>{{ t('fmtTextDeleteMarquee', [announcementInfo.subject]) }}</div>
      <ul class="ml-6 mt-6 list-disc">
        <li class="text-danger">{{ t('textDeleteAnnouncementReminder') }}</li>
      </ul>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="deleteDialog = false">{{ t('textCancel') }}</button>
        <LoadingButton class="btn btn-primary" :button-click="() => deleteAnnouncement()">
          {{ t('textDelete') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useBackendAnnouncement } from '@/base/composable/operationManagement/backendAnnouncement/useBackendAnnouncement'
import { useI18n } from 'vue-i18n'
import { CalendarIcon } from '@heroicons/vue/20/solid'
import constant from '@/base/common/constant.js'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import BackendAnnouncementComponent from './BackendAnnouncementComponent.vue'

import ViewButton from '@/base/components/Button/ViewButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import DeleteButton from '@/base/components/Button/DeleteButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const {
  formInput,
  tableInput,
  records,
  types,
  showDialog,
  showMode,
  deleteDialog,
  announcementInfo,
  isEnabled,
  searchRecords,
  getAnnouncementInfo,
  createNewAnnounce,
  deleteAnnouncement,
} = useBackendAnnouncement()
</script>
