<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textGame') }}</label>
          <FormGameListDropdown v-model="formInput.game" :include-all="false" class="mb-1 w-full md:w-3/12" />
        </div>

        <div class="flex flex-wrap items-center justify-end">
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

      <PageTable
        :records="realTimeGameRatios.items"
        :total-records="realTimeGameRatios.items.length"
        :table-input="tableInput"
      >
        <template #default="{ currentRecords, isSortIconActive, sorting }">
          <div class="tbl-container">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <PageTableSortableTh column="gameId" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textGameName') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="roomType" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textRoomType') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="playerCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textOnlinePlayerCount') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="playCount" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textPlayCount') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="yaScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textBetTotal') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="deScore" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalGamerWinScore') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="tax" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalGameTax') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="bonus" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textTotalBonus') }}</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="rtp" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>RTP</template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="deductTaxRtp" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default>{{ t('textDeductTaxRtp') }}</template>
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
                  <tr v-for="record in currentRecords" :key="`record_${record.roomId}`">
                    <td>{{ t(`game__${record.gameId}`) }}</td>
                    <td>{{ t(`roomType__${roomTypeNameIndex(record.gameId, record.roomType)}`) }}</td>
                    <td>{{ record.playerCount.toLocaleString() }}</td>
                    <td>{{ record.playCount.toLocaleString() }}</td>
                    <td>{{ numberToStr(record.yaScore) }}</td>
                    <td>{{ numberToStr(record.deScore) }}</td>
                    <td>{{ numberToStr(record.tax) }}</td>
                    <td>{{ numberToStr(record.bonus) }}</td>
                    <td :class="record.rtp >= 1 ? 'text-danger' : 'text-success'">
                      {{ `${numberToStr(record.rtp * 100)}%` }}
                    </td>
                    <td :class="record.deductTaxRtp >= 1 ? 'text-danger' : 'text-success'">
                      {{ `${numberToStr(record.deductTaxRtp * 100)}%` }}
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
        </template>
      </PageTable>
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useRealTimeGameRatio } from '@/base/composable/riskManagement/realTimeGameRatio/useRealTimeGameRatio'
import { numberToStr } from '@/base/utils/formatNumber'
import { roomTypeNameIndex } from '@/base/utils/room'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormGameListDropdown from '@/base/components/Form/Dropdown/FormGameListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'

const { t } = useI18n()
const { formInput, realTimeGameRatios, tableInput, searchRecords } = useRealTimeGameRatio()
</script>
