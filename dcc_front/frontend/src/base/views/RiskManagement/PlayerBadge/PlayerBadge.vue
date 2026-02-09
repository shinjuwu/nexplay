<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 lg:w-1/12 lg:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full lg:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 lg:w-1/12 lg:text-right">{{ t('textRiskType') }}</label>
          <FormAgentTagsDropdown v-model="formInput.tag" class="mb-1 w-full lg:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 lg:w-1/12 lg:text-right">{{ t('textTime') }}</label>
          <FormDateTimeInput
            v-model="formInput.startTime"
            class="flex-1"
            :minute-increment="60"
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
            :minute-increment="60"
            :before-days="time.commonReportTimeBeforeDays"
            set-max-date
          />
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 lg:w-1/12 lg:text-right">{{ t('textWinScoreMoreThan') }}</label>
          <input
            v-model="formInput.winScore"
            type="number"
            class="form-input w-full lg:w-3/12"
            min="0"
            max="999999999"
            step="1"
            :placeholder="t('placeHolderTextValue')"
          />
          <label class="form-label mb-1 w-full pr-2 lg:w-1/12 lg:text-right">{{ t('textRTPMoreThan') }}</label>
          <input
            v-model="formInput.rtp"
            type="number"
            min="0"
            max="100"
            step="0.1"
            :placeholder="t('placeHolderTextInputPercentage')"
            class="form-input w-full lg:w-3/12"
          />
          <label class="form-label mb-1 w-full pr-2 lg:w-1/12 lg:text-right">{{ t('textWinRateMoreThan') }}</label>
          <input
            v-model="formInput.winRate"
            type="number"
            min="0"
            max="100"
            step="0.1"
            :placeholder="t('placeHolderTextInputPercentage')"
            class="form-input w-full lg:w-3/12"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full lg:ml-2 lg:w-2/12 xl:w-1/12"
            @click="getGameUserTagList"
          >
            {{ t('textSearch') }}
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
      <!-- <TogglePageTips :title="t('textCalDirections')">
        <template #default="{ tipsShow }">
          <div v-show="tipsShow" class="mt-2">
            <div v-for="(direction, dirIdx) in calDirections" :key="`directions__${dirIdx}`" class="text-danger">
              {{ direction }}
            </div>
          </div>
        </template>
      </TogglePageTips> -->
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
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tagList" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textMark') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="state" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textState') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="killDiveValue"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textKillDiveValueBalance') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textWinScore') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bonus" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textBonus') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="rtp" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> RTP </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winRate" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textWinRate') }} </template>
                  </PageTableSortableTh>
                  <th>
                    {{ t('textOperate') }}
                  </th>
                </tr>
              </thead>
              <tbody class="text-slate-500">
                <template v-if="currentRecords.length === 0">
                  <tr>
                    <td class="text-center" colspan="11">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="(record, index) in currentRecords" :key="`record_${index}`" class="hover:bg-slate-100">
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.userName }}</td>
                    <td class="space-x-2">
                      <template v-if="record.killDiveState === 2">
                        <span
                          class="rounded-full px-1 text-xs"
                          :style="{
                            backgroundColor: defaultBadgeInfo.Blacklist.BackgroundColor,
                            color: defaultBadgeInfo.Blacklist.TextColor,
                          }"
                          >{{ t(`riskType__${defaultBadgeInfo.Blacklist.Index + 4}`) }}
                        </span>
                      </template>
                      <template v-if="record.highRisk">
                        <span
                          class="rounded-full px-1 text-xs"
                          :style="{
                            backgroundColor: defaultBadgeInfo.HighRisk.BackgroundColor,
                            color: defaultBadgeInfo.HighRisk.TextColor,
                          }"
                          >{{ t(`riskType__${defaultBadgeInfo.HighRisk.Index + 4}`) }}
                        </span>
                      </template>
                      <template v-if="record.killDiveState === 1">
                        <span
                          class="rounded-full px-1 text-xs"
                          :style="{
                            backgroundColor: defaultBadgeInfo.TargetKill.BackgroundColor,
                            color: defaultBadgeInfo.TargetKill.TextColor,
                          }"
                          >{{ t(`riskType__${defaultBadgeInfo.TargetKill.Index + 4}`) }}
                        </span>
                      </template>
                      <template v-if="!isAdminUser">
                        <template
                          v-for="customBadge in agentsCustomTagInfo[record.agentId]"
                          :key="`customBadge__${customBadge.index}`"
                        >
                          <span
                            v-if="record.tagList[customBadge.index + 3] !== '0'"
                            :style="{ backgroundColor: customBadge.bgColor, color: customBadge.txtColor }"
                            class="rounded-full px-1 text-xs"
                          >
                            {{ customBadge.name }}
                          </span>
                        </template>
                      </template>
                    </td>
                    <td>
                      <span :class="{ 'text-danger': !record.state }">
                        {{ record.state ? t('textOpened') : t('textDisabled') }}
                      </span>
                    </td>
                    <td>{{ record.killDiveValue > 0 ? numberToStr(record.killDiveValue) : '-' }}</td>
                    <td>{{ numberToStr(record.winScore) }}</td>
                    <td>{{ numberToStr(record.bonus) }}</td>
                    <td>{{ numberToStr(record.rtp, 1) + '%' }}</td>
                    <td>{{ numberToStr(record.winRate, 1) + '%' }}</td>
                    <td class="space-x-2">
                      <BadgeButton
                        v-if="isEnabled.PlayerBadgeUpdate && myAgentId === record.agentId"
                        @click="selectPlayerInfo(record, 'badge')"
                      />
                      <EditButton
                        v-if="isEnabled.PlayerAccountInfoUpdate"
                        @click="selectPlayerInfo(record, 'setting')"
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
  <PlayerBadgeDialog
    :visible="dialog.badge"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.badge = newValue)"
    @search-game-users="getGameUserTagList"
  />
  <PlayerAccountSetting
    :visible="dialog.setting"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.setting = newValue)"
    @search-game-users="getGameUserTagList"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { usePlayerBadge } from '@/base/composable/riskManagement/playerBadge/usePlayerBadge'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormAgentTagsDropdown from '@/base/components/Form/Dropdown/FormAgentTagsDropdown.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import BadgeButton from '@/base/components/Button/BadgeButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import PlayerAccountSetting from '@/base/views/Common/Dialog/PlayerAccountSettingDialog.vue'
import PlayerBadgeDialog from '@/base/views/Common/Dialog/PlayerBadgeDialog.vue'

const { t } = useI18n()
const {
  agentsCustomTagInfo,
  calDirections,
  defaultBadgeInfo,
  dialog,
  formInput,
  isAdminUser,
  isEnabled,
  myAgentId,
  playerInfo,
  records,
  tableInput,
  getGameUserTagList,
  selectPlayerInfo,
} = usePlayerBadge()
</script>
