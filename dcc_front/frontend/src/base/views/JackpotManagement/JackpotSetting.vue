<template>
  <ToggleHeader>
    <template #default="{ show }">
      <div v-show="show" class="tbl-container">
        <table class="tbl tbl-hover">
          <thead>
            <tr>
              <th>{{ t('textGameName') }}</th>
              <th>{{ t('textGameState') }}</th>
              <th v-if="isSettingEditEnabled">{{ t('textOperate') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>JACKPOT</td>
              <td class="space-x-2">
                <template v-if="isSettingEditEnabled">
                  <button
                    :class="[jackpotGameState ? 'btn-success' : 'btn-secondary']"
                    class="btn"
                    @click="setJackpotGameState(constant.GameState.Online)"
                  >
                    {{ t(`state__${constant.GameState.Online}`) }}
                  </button>
                  <button
                    :class="[!jackpotGameState ? 'btn-warning' : 'btn-secondary']"
                    class="btn"
                    @click="setJackpotGameState(constant.GameState.Offline)"
                  >
                    {{ t(`state__${constant.GameState.Offline}`) }}
                  </button>
                </template>
                <template v-else>
                  <span :class="[jackpotGameState ? 'text-success' : 'text-warning']">
                    {{ t(`state__${Number(jackpotGameState)}`) }}
                  </span>
                </template>
              </td>
              <td v-if="isSettingEditEnabled">
                <ButtonTooltips :tips-text="t('textSyncAgentData')">
                  <template #content>
                    <label
                      class="inline-flex cursor-pointer items-center justify-center"
                      @click="notifyGameToGameServer()"
                    >
                      <ArrowUpOnSquareIcon class="inline h-5 w-5 rounded" />
                    </label>
                  </template>
                </ButtonTooltips>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <PageTableLoading v-show="showJackpotServerProcessing" />
    </template>
  </ToggleHeader>
  <ToggleHeader class="mt-4">
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" :include-grandson="false" class="mb-1 w-full md:w-3/12" />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="getAgentJackpotList()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <ToggleHeader :tips-title="t('textJackpotSettingDirections')">
        <template #tipsSlot="{ tipsShow }">
          <div
            v-for="(direction, dirIdx) in pageDirections"
            v-show="tipsShow"
            :key="`directions__${dirIdx}`"
            class="text-danger"
          >
            {{ direction }}
          </div>
        </template>
      </ToggleHeader>
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
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textGeneralAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="subAgentCount"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textSubAgents') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="status" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textJackpotStatus') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="startTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textStartDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="endTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textEndDate') }} </template>
                  </PageTableSortableTh>
                  <th v-if="isSettingEditEnabled">{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="6">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.subAgentCount }}</td>
                    <td>
                      <span :class="{ 'text-danger': record.status === 1 }">
                        {{ t(`jackpotStatus__${record.status}`) }}
                      </span>
                    </td>
                    <td>
                      {{
                        record.status === 0 && record.startTime === '1970-01-01T00:00:00Z'
                          ? '-'
                          : time.utcTimeStrToLocalTimeFormat(record.startTime)
                      }}
                    </td>
                    <td>
                      {{
                        record.status === 0 && record.endTime === '1970-01-01T00:00:00Z'
                          ? '-'
                          : time.utcTimeStrToLocalTimeFormat(record.endTime)
                      }}
                    </td>
                    <td v-if="isSettingEditEnabled">
                      <EditButton @click="getAgentJackpotSetting(record)"></EditButton>
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
  <PageTableDialog :visible="visible" @close="visible = false">
    <template #header>
      <span class="px-4">
        {{ selectedAgent.agentName }}
      </span>
      {{ t('menuItemJackpotSetting') }}
    </template>
    <template #default>
      <div class="flex items-center justify-around bg-slate-100 p-3.5 font-bold">
        <div>{{ t('textJackpotStatus') }}</div>
        <CheckButton
          :content="t(`jackpotStatus__0`)"
          :checked="selectedAgent.newStatus === 0"
          @click="selectedAgent.newStatus = 0"
        />
        <CheckButton
          :content="t(`jackpotStatus__1`)"
          :checked="selectedAgent.newStatus === 1"
          @click="selectedAgent.newStatus = 1"
        />
        <CheckButton
          :content="t(`jackpotStatus__2`)"
          :checked="selectedAgent.newStatus === 2"
          @click="selectedAgent.newStatus = 2"
        />
      </div>
      <div class="mt-4 flex justify-between space-x-4">
        <div class="w-full">
          <label class="form-label">{{ t('textStartDate') }}</label>
          <FormDateTimeInput
            v-if="isShowDateTimeInput"
            :key="selectedAgent.newStatus"
            v-model="selectedAgent.startTime"
            :disabled="isTimeEditable"
            class="flex-1"
            calendar-align="left"
          />
          <input
            v-else
            disabled
            class="block w-full border border-gray-200 py-2 px-3.5 text-gray-500 outline-none"
            value="-"
          />
        </div>
        <div class="w-full">
          <label class="form-label">{{ t('textEndDate') }}</label>
          <FormDateTimeInput
            v-if="isShowDateTimeInput"
            :key="selectedAgent.newStatus"
            v-model="selectedAgent.endTime"
            :disabled="isTimeEditable"
            class="flex-1"
            calendar-align="right"
          />
          <input
            v-else
            disabled
            class="block w-full border border-gray-200 py-2 px-3.5 text-gray-500 outline-none"
            value="-"
          />
        </div>
      </div>
      <div class="text-danger mt-2">
        <p>{{ t('textJPSettingDirection__1') }}</p>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="visible = false">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="selectedAgent"
          :button-click="
            () => {
              updateAgentJackpotSetting()
            }
          "
        >
          {{ t('textSend') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { ArrowUpOnSquareIcon } from '@heroicons/vue/24/outline'
import { useJackpotSetting } from '@/base/composable/jackpotManagement/useJackpotSetting'
import time from '@/base/utils/time'
import constant from '@/base/common/constant'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'
import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'

const { t } = useI18n()

const {
  records,
  formInput,
  tableInput,
  visible,
  selectedAgent,
  pageDirections,
  isShowDateTimeInput,
  isTimeEditable,
  isSettingEditEnabled,
  jackpotGameState,
  showJackpotServerProcessing,
  getAgentJackpotList,
  getAgentJackpotSetting,
  updateAgentJackpotSetting,
  setJackpotGameState,
  notifyGameToGameServer,
} = useJackpotSetting()
</script>
