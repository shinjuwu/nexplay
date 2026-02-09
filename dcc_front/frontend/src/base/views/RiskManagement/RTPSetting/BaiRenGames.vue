<template>
  <div>
    <ToggleHeader>
      <template #default="{ show }">
        <form v-show="show">
          <div class="flex flex-wrap items-center">
            <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textGame') }}</label>
            <FormGameListDropdown
              v-model="formInput.gameName"
              class="mb-1 w-full md:w-3/12"
              :game-type="props.gameType"
            />
            <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoomType') }}</label>
            <FormBaiRenRoomTypeListDropdown v-model="formInput.roomType" class="mb-1 w-full md:w-3/12" />
          </div>
          <div class="mt-4 flex flex-wrap items-center" :class="[isEditEnabled ? 'justify-between' : 'justify-end']">
            <button
              v-if="isEditEnabled"
              type="button"
              class="btn btn-danger mb-1 flex w-full items-center md:ml-2 md:w-2/12 xl:w-fit"
              @click="showBatchSettingDialog = true"
            >
              {{ t('roleItemRTPBatchSettingUpdate') }}
            </button>
            <button
              type="button"
              class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
              @click="searchRecords(props.gameType)"
            >
              {{ t('textSearch') }}
            </button>
          </div>
        </form>
        <hr class="my-2" />
        <ToggleHeader :tips-title="t('textSpecialDirections')">
          <template #tipsSlot="{ tipsShow }">
            <div v-show="tipsShow" class="text-danger">
              <div v-for="(direction, index) in BaiRenDirections" :key="`BaiRenDirection__${index}`">
                {{ direction }}
              </div>
            </div>
            <hr class="my-2" />
          </template>
        </ToggleHeader>
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
                    <PageTableSortableTh column="gameCode" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textGameCode') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="gameId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textGameName') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="roomType" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textRoomType') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="killRatioSort"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default>{{ t('textRTP') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="newKillRatio"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default>{{ t('textNewKillRatio') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="activeNum" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default>{{ t('textActiveNum') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh
                      column="lastUpdateTime"
                      :is-sort-icon-active="isSortIconActive"
                      @sorting="sorting"
                    >
                      <template #default> {{ t('textLastUpdateTime') }} </template>
                    </PageTableSortableTh>
                    <th>
                      {{ t('textOperate') }}
                    </th>
                  </tr>
                </thead>
                <tbody class="text-slate-500">
                  <template v-if="currentRecords.length === 0">
                    <tr>
                      <td class="table-td text-center" colspan="8">{{ t('textTableEmpty') }}</td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="(record, index) in currentRecords" :key="`bairenRecord_${index}`">
                      <td>{{ record.gameCode }}</td>
                      <td>{{ t(`game__${record.gameId}`) }}</td>
                      <td>{{ t(`roomType__${roomTypeNameIndex(record.gameId, record.roomType)}`) }}</td>
                      <td>
                        <span>{{ numberToStr(record.killRatio * 100, 1) + '%' }}</span>
                      </td>
                      <td>
                        <span>{{ numberToStr(record.newKillRatio * 100, 1) + '%' }}</span>
                      </td>
                      <td>
                        <span>{{ record.activeNum }}</span>
                      </td>
                      <td>{{ time.utcTimeStrToLocalTimeFormat(record.lastUpdateTime) }}</td>
                      <td>
                        <EditButton @click="getRatioSet(record)" />
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
    <GamesRTPSettingDialog
      :select-record="selectRecord"
      :visible="showSettingDialog"
      :game-type="props.gameType"
      @close="showSettingDialog = false"
      @set-game-ratio="(newValue) => setGameRatio(newValue)"
    />
    <GameRatioSettingDialog
      :visible="showBatchSettingDialog"
      :game-type="props.gameType"
      @close="showBatchSettingDialog = false"
    />
  </div>
</template>

<script setup>
import { useRTPGameSetting } from '@/base/composable/riskManagement/RTPSetting/useRTPGameSetting'
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'
import time from '@/base/utils/time'
import { roomTypeNameIndex } from '@/base/utils/room'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormGameListDropdown from '@/base/components/Form/Dropdown/FormGameListDropdown.vue'
import FormBaiRenRoomTypeListDropdown from '@/base/components/Form/Dropdown/FormBaiRenRoomTypeListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import GamesRTPSettingDialog from './GamesRTPSettingDialog.vue'
import GameRatioSettingDialog from './GameRatioSettingDialog.vue'

const props = defineProps({
  gameType: {
    type: Number,
    default: 0,
  },
})

const { t } = useI18n()
const {
  searchRecords,
  getRatioSet,
  setGameRatio,
  formInput,
  tableInput,
  showSettingDialog,
  showBatchSettingDialog,
  records,
  selectRecord,
  BaiRenDirections,
  isEditEnabled,
} = useRTPGameSetting(props)
</script>
