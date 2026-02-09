<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textMarqueeType') }}</label>
          <FormDropdown
            v-model="formInput.type"
            class="mb-1 md:w-3/12"
            :fmt-item-text="(type) => (type === constant.MarqueeType.All ? t('textTypeAll') : t(`typeCode__${type}`))"
            :fmt-item-key="(type) => `type__${type}`"
            :items="types"
            :use-font-awsome="false"
          />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textLanguage') }}</label>
          <FormDropdown
            v-model="formInput.lang"
            class="mb-1 md:w-3/12"
            :fmt-item-text="(lang) => (lang === constant.LangType.All ? t('textLangAll') : t(`langType__${lang}`))"
            :fmt-item-key="(lang) => `lang__${lang}`"
            :items="langs"
            :use-font-awsome="false"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <div class="flex-1">
            <button
              v-if="isEnabled.MarqueeSettingCreate"
              type="button"
              class="btn btn-danger mb-1 flex w-full items-center md:ml-2 md:w-2/12 xl:w-fit"
              @click="createNewMarquee"
            >
              <SquaresPlusIcon class="mr-2 inline-block h-5 w-5" />
              {{ t('textCreateMarquee') }}
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
                  <PageTableSortableTh column="createDate" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textCreatingDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="type" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textMarqueeType') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="content" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textContentText') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="lang" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textLanguage') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="startDate" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textStartDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="endDate" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textEndDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="order" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPriority') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="freq" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textDisplayFrequency') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="isEnabled" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textIsStart') }} </template>
                  </PageTableSortableTh>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="12">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`marqueeInfo__${record.id}`">
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.createDate) }}</td>
                    <td>{{ t(`typeCode__${record.type}`) }}</td>
                    <td width="25%">
                      {{ ellipsisString(record.content) }}
                    </td>
                    <td>{{ t(`langType__${record.lang}`) }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.startDate) }}</td>
                    <td>{{ time.utcTimeStrToLocalTimeFormat(record.endDate) }}</td>
                    <td>{{ record.order }}</td>
                    <td>{{ record.freq }}</td>
                    <td>
                      {{ record.isEnabled ? t('textOpened') : t('textClose') }}
                    </td>
                    <td>
                      <ViewButton @click="getMarqueeInfo(record.id, 'view')" />
                      <EditButton
                        v-if="isEnabled.MarqueeSettingUpdate"
                        class="mx-2"
                        @click="getMarqueeInfo(record.id, 'update')"
                      />
                      <DeleteButton
                        v-if="isEnabled.MarqueeSettingDelete"
                        @click="getMarqueeInfo(record.id, 'delete')"
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
  <MarqueeSettingComponent
    :visible="showDialog"
    :mode="showMode"
    :marquee-info="marqueeInfo"
    @refresh-marquee-list="searchRecords"
    @close="showDialog = false"
  />
  <PageTableDialog :visible="deleteDialog" @close="deleteDialog = false">
    <template #header>
      {{ t('roleItemMarqueeSettingDelete') }}
    </template>
    <template #default>
      <p class="mb-3">{{ t('fmtTextDeleteMarquee', [ellipsisString(marqueeInfo.content)]) }}</p>
      <ul class="pl-5">
        <li style="list-style: disc" class="text-danger">{{ t('textDeleteMarqueeReminder') }}</li>
      </ul>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="deleteDialog = false">{{ t('textCancel') }}</button>
        <LoadingButton class="btn btn-primary" :button-click="() => deleteMarquee()">
          {{ t('textConfirm') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>
<script setup>
import { useMarqueeSetting } from '@/base/composable/operationManagement/marqueeSetting/useMarqueeSetting'
import time from '@/base/utils/time'
import { SquaresPlusIcon } from '@heroicons/vue/20/solid'
import { ellipsisString } from '@/base/utils/string'
import constant from '@/base/common/constant'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import DeleteButton from '@/base/components/Button/DeleteButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'
import MarqueeSettingComponent from './MarqueeSettingComponent.vue'

const {
  t,
  isEnabled,
  records,
  types,
  langs,
  tableInput,
  formInput,
  marqueeInfo,
  showDialog,
  deleteDialog,
  showMode,
  searchRecords,
  getMarqueeInfo,
  createNewMarquee,
  deleteMarquee,
} = useMarqueeSetting()
</script>
