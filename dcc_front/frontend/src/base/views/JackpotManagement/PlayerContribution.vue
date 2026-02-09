<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="getJackpotLeaderBoard()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <ToggleHeader :tips-title="t('textPlayerContributionDirections')">
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
      <PageTable :records="rankingList.items" :total-records="rankingList.items.length" :table-input="tableInput">
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
                  <PageTableSortableTh column="ranking" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRanking') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgent') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="playCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalGameSession') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="totalBet" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textValidBetTotal') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="loseCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalLoseSession') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="winCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textTotalWinSession') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="contribute" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textContribution') }} </template>
                  </PageTableSortableTh>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="9">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.ranking }}</td>
                    <td>{{ record.agentName }}</td>
                    <td>{{ record.userName }}</td>
                    <td>{{ record.playCount }}</td>
                    <td>{{ numberToStr(record.totalBet) }}</td>
                    <td>{{ record.loseCount }}</td>
                    <td>{{ record.winCount }}</td>
                    <td>
                      <span class="text-danger">{{ numberToStr(record.contribute) }}</span>
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

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import { usePlayerContribution } from '@/base/composable/jackpotManagement/usePlayerContribution'
import { numberToStr } from '@/base/utils/formatNumber'

const { t } = useI18n()
const { tableInput, formInput, rankingList, pageDirections, getJackpotLeaderBoard } = usePlayerContribution()
</script>
