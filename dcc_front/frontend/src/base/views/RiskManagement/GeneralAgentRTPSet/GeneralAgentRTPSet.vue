<template>
  <div>
    <ToggleHeader>
      <template #default="{ show }">
        <form v-show="show">
          <div class="flex flex-nowrap items-center">
            <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }} </label>
            <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" :include-grandson="false" />
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
        <CheckButton
          v-for="state in accountStates"
          :key="`accountState__${state}`"
          :content="t(`enabledType__${state}`)"
          class="mx-1"
          :checked="formInput.state === state"
          @click="formInput.state = state"
        ></CheckButton>
        <PageTable :records="filterRecords" :total-records="filterRecords.length" :table-input="tableInput">
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
                    <PageTableSortableTh column="ratio" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textRTP') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="updateTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textLastUpdateTime') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="state" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textState') }} </template>
                    </PageTableSortableTh>
                    <th v-if="isSettingEnabled">
                      {{ t('textOperate') }}
                    </th>
                  </tr>
                </thead>
                <tbody class="text-slate-500">
                  <template v-if="currentRecords.length === 0">
                    <tr>
                      <td class="table-td text-center" :colspan="isSettingEnabled ? 7 : 6">
                        {{ t('textTableEmpty') }}
                      </td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="(record, index) in currentRecords" :key="`record_${index}`" class="hover:bg-slate-100">
                      <td>{{ record.agentName }}</td>
                      <td>{{ numberToStr(record.ratio, 1) + '%' }}</td>
                      <td>{{ time.utcTimeStrToLocalTimeFormat(record.updateTime) }}</td>
                      <td>{{ t(`enabledType__${record.state}`) }}</td>
                      <td v-if="isSettingEnabled">
                        <EditButton @click="getAgentRatio(record)" />
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
  </div>
  <PageTableDialog :visible="visible" @close="visible = false">
    <template #header>{{ t('menuItemGeneralAgentRTPSet') }}</template>
    <template #default>
      <div class="grid grid-cols-4 gap-2">
        <div class="col-span-4">
          <label class="form-label">
            {{ t('fmtTextGeneralAgentName', [agentRatioInfo.agentName]) }}
          </label>
        </div>
        <div class="col-span-2 flex items-center">
          <label class="form-label">{{ `${t('textState')}：` }}</label>
          <CheckButton
            v-for="state in checkStates"
            :key="`agentAccountState__${state}`"
            :content="t(`enabledType__${state}`)"
            class="form-label mx-1"
            :checked="agentRatioInfo.state === state"
            @click="agentRatioInfo.state = state"
          ></CheckButton>
        </div>
        <div class="col-span-2 flex items-center">
          <label class="form-label min-w-fit">{{ t('textRTP') }}</label>
          <input v-model="agentRatioInfo.ratio" max="100" min="0" step="0.1" class="form-input mx-2" type="number" />%
        </div>
      </div>
      <div>
        <label class="form-label">{{ t('textRemark') }}</label>
        <textarea v-model="agentRatioInfo.remark" rows="2" class="form-input"></textarea>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="visible = false">
          {{ t('textClose') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="agentRatioInfo"
          :button-click="() => setAgentRatio(agentRatioInfo)"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useGeneralAgentRTPSet } from '@/base/composable/riskManagement/generalAgentRTPSet/useGeneralAgentRTPSet'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const {
  accountStates,
  agentRatioInfo,
  checkStates,
  filterRecords,
  formInput,
  tableInput,
  visible,
  isSettingEnabled,
  getAgentRatio,
  searchRecords,
  setAgentRatio,
} = useGeneralAgentRTPSet()
</script>
