<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" class="mb-1 w-full md:w-3/12" :include-all="false" />

          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">
            {{ t('textAgentSettingSupplyModeCycle') }}
          </label>
          <FormDropdown
            v-model="formInput.timeType"
            class="mb-1 w-full md:w-3/12"
            :fmt-item-text="(type) => (type === timeTypes[0] ? t('inputTimeType__day') : t(`inputTimeType__${type}`))"
            :fmt-item-key="(type) => `inputTimeType__${type}`"
            :items="timeTypes"
            :use-font-awsome="false"
          />

          <div
            v-if="isShowIsSearchAll"
            class="flex w-full cursor-pointer items-center md:w-3/12"
            @click="formInput.isSearchAll = !formInput.isSearchAll"
          >
            <Checkbox class="mr-1 mb-1 md:ml-3" :checked="formInput.isSearchAll" />
            <label class="form-label mb-1 w-full cursor-pointer">{{ t('textIncludeChildAgents') }}</label>
          </div>
        </div>

        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchDataOverview()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>

      <hr class="my-6" />

      <div>
        <keep-alive>
          <div id="RealTimeUserCountChart" class="h-[500px]"></div>
        </keep-alive>
      </div>

      <hr class="my-6" />

      <div class="grid grid-cols-3 gap-4">
        <div class="col-span-full grid grid-cols-2 gap-4 md:col-span-2">
          <div class="col-span-full md:col-span-1">
            <DataCardComponent
              :title="t('textActivePlayer')"
              :data="[stats.thisTime.activePlayer.toLocaleString(), stats.lastTime.activePlayer.toLocaleString()]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <UserGroupIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>

          <div class="col-span-full md:col-span-1">
            <DataCardComponent
              :title="t('textBetPeopleNumbers')"
              :data="[stats.thisTime.bettorsNum.toLocaleString(), stats.lastTime.bettorsNum.toLocaleString()]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <UserIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>

          <div class="col-span-full md:col-span-1">
            <DataCardComponent
              :title="t('textRegisterNumbers')"
              :data="[stats.thisTime.registerNum.toLocaleString(), stats.lastTime.registerNum.toLocaleString()]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <UserIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>

          <div class="col-span-full md:col-span-1">
            <DataCardComponent
              :title="t('textBetCounts')"
              :data="[stats.thisTime.oddNum.toLocaleString(), stats.lastTime.oddNum.toLocaleString()]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <ArrowDownCircleIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>

          <div class="col-span-full md:col-span-1">
            <DataCardComponent
              :title="t('textBetTotal')"
              :data="[numberToStr(stats.thisTime.totalBet), numberToStr(stats.lastTime.totalBet)]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <BanknotesIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>

          <div class="col-span-full md:col-span-1">
            <DataCardComponent
              :title="t('textGameTax')"
              :data="[numberToStr(stats.thisTime.gameTax), numberToStr(stats.lastTime.gameTax)]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <CurrencyDollarIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>

          <div class="col-span-full md:col-span-2">
            <DataCardComponent
              :title="t('textPlatformWinLoseScore')"
              :data="[numberToStr(stats.thisTime.totalScore), numberToStr(stats.lastTime.totalScore)]"
              :compare-text="timesText.than"
            >
              <template #icon>
                <CurrencyDollarIcon class="h-6 w-6" />
              </template>
            </DataCardComponent>
          </div>
        </div>

        <div class="col-span-full md:col-span-1">
          <div>
            <div class="mb-1">{{ t('fmtTextPeriodRiskPlayers') }}</div>
            <div class="tbl-container max-h-[250px] overflow-y-scroll">
              <table class="tbl tbl-hover">
                <thead class="sticky top-0">
                  <tr>
                    <th>{{ t('textPlayerAccount') }}</th>
                    <th>{{ t('textBetTotal') }}</th>
                    <th>{{ t('textTotalGamerWinScore') }}</th>
                  </tr>
                </thead>
                <tbody>
                  <template v-if="tableData.riskList.length === 0">
                    <tr class="no-hover">
                      <td colspan="7">{{ t('textTableEmpty') }}</td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="record in tableData.riskList" :key="`record_${record.id}`">
                      <td>{{ record.userName }}</td>
                      <td>{{ numberToStr(record.totalYa) }}</td>
                      <td>
                        <span class="text-danger">+{{ numberToStr(record.totalDe) }}</span>
                      </td>
                    </tr>
                  </template>
                </tbody>
              </table>
            </div>
          </div>
          <div class="mt-4">
            <div class="mb-1">{{ t('fmtTextPeriodGamesLeaderBoard') }}</div>
            <div class="tbl-container max-h-[250px] overflow-y-scroll">
              <table class="tbl tbl-hover">
                <thead class="sticky top-0">
                  <tr>
                    <th>{{ t('textGame') }}</th>
                    <th>{{ t('textBetTotal') }}</th>
                    <th>{{ t('textPlatformWinLoseScore') }}</th>
                  </tr>
                </thead>
                <tbody>
                  <template v-if="tableData.leaderBoards.length === 0">
                    <tr class="no-hover">
                      <td colspan="7">{{ t('textTableEmpty') }}</td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="record in tableData.leaderBoards" :key="`record_${record.id}`">
                      <td>{{ t(`game__${record.gameID}`) }}</td>
                      <td>{{ numberToStr(record.totalYa) }}</td>
                      <td>
                        <div :class="{ 'text-danger': record.totalBet > 0, 'text-success': record.totalBet < 0 }">
                          <span v-if="record.totalBet > 0">+</span>
                          <span>{{ numberToStr(record.totalBet) }}</span>
                        </div>
                      </td>
                    </tr>
                  </template>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <hr class="my-6" />

      <div>
        <keep-alive>
          <div id="thirtyDaysChart" class="h-[500px]"></div>
        </keep-alive>
      </div>

      <hr class="my-2" />

      <div>
        <keep-alive>
          <div id="AllDaysWinLoseChart" class="h-[500px]"></div>
        </keep-alive>

        <table class="mb-10 block w-full overflow-x-auto">
          <tr>
            <th></th>
            <th v-for="hour in hours" :key="`hour_${hour} `" class="border text-center">
              {{ `${hour}:00` }}
            </th>
          </tr>
          <tr v-for="series in chartTable.AllDaysWinLoseChart" :key="`winlose_series_${series.name} `">
            <td class="border text-center md:min-w-[80px]">
              {{ series.name }}
            </td>
            <td
              v-for="(item, winLoseItemKey) in series.data.slice(24, 48)"
              :key="`winlose_series_item_${winLoseItemKey} `"
              class="border text-center md:w-[70px]"
            >
              <span v-if="item.isTime" :class="[parseInt(item.value) > 0 ? 'text-danger' : 'text-success']">
                {{ item.value !== 0 ? numberToStr(item.value) : 0 }}
              </span>
              <span v-else>-</span>
            </td>
          </tr>
        </table>
      </div>

      <hr class="my-5" />

      <div>
        <keep-alive>
          <div id="AllDaysPlayersChart" class="h-[500px]"></div>
        </keep-alive>

        <table class="block w-full overflow-x-auto" :class="{ 'mb-10': isDeviceLocationInfoEnabled }">
          <tr>
            <th></th>
            <th v-for="hour in hours" :key="`hour_${hour} `" class="border text-center">
              {{ `${hour}:00` }}
            </th>
          </tr>
          <tr v-for="series in chartTable.AllDaysPlayersChart" :key="`players_series_${series.name} `">
            <td class="border text-center md:min-w-[80px]">
              {{ series.name }}
            </td>
            <td
              v-for="(item, playerSeriesItemKey) in series.data.slice(24, 48)"
              :key="`players_series_item_${playerSeriesItemKey} `"
              class="border text-center md:w-[70px]"
            >
              <span v-if="item.isTime" :class="[parseInt(item.value) > 0 ? 'text-danger' : 'text-success']">
                {{ item.value !== 0 ? item.value.toLocaleString() : 0 }}</span
              >
              <span v-else>-</span>
            </td>
          </tr>
        </table>
      </div>

      <template v-if="isDeviceLocationInfoEnabled">
        <hr class="my-5" />

        <div class="grid gap-4 md:grid-cols-3">
          <div class="col-span-full md:col-span-2">
            <keep-alive>
              <div id="AllDaysDeviceChart" class="h-[500px]"></div>
            </keep-alive>
          </div>

          <div class="col-span-full md:col-span-1">
            <div class="mb-1">{{ t('textTodaySourceLocationRanking') }}</div>
            <div class="tbl-container max-h-[450px] overflow-y-scroll">
              <table class="tbl tbl-hover">
                <thead class="sticky top-0">
                  <tr>
                    <th>{{ t('textCountry') }}</th>
                    <th>{{ t('textRegion') }}</th>
                    <th>{{ t('textCount') }}</th>
                  </tr>
                </thead>
                <tbody>
                  <template v-if="tableData.locationList.length === 0">
                    <tr class="no-hover">
                      <td colspan="3">{{ t('textTableEmpty') }}</td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="record in tableData.locationList" :key="`record_${record.country}_${record.region}`">
                      <td>{{ record.country }}</td>
                      <td>{{ record.region }}</td>
                      <td>{{ numberToStr(record.count) }}</td>
                    </tr>
                  </template>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </template>
    </template>
  </ToggleHeader>

  <PageTableDialog :visible="showChartDialog" size="xl" @close="showChartDialog = false">
    <template #header>
      <div class="text-center text-2xl">{{ dialogTitle }}</div>
    </template>
    <template #default>
      <div id="dialogChart" class="h-[500px]"></div>
      <table class="-mt-5 block w-full overflow-x-auto">
        <tr>
          <th></th>
          <th v-for="hour in hours" :key="`hour_${hour} `" class="border text-center">
            {{ `${hour}:00` }}
          </th>
        </tr>
        <tr v-for="series in chartTable.dialogChart" :key="`dialog_series_${series.name} `">
          <td class="sticky left-0 z-[1] min-w-[60px] border bg-white text-center">
            {{ series.name }}
          </td>
          <td
            v-for="(item, idx) in series.data"
            :key="`dialog_series_item_${item.value}_${idx}`"
            class="min-w-[50px] border px-2 text-center"
          >
            <span v-if="item.isTime" :class="[parseInt(item.value) > 0 ? 'text-danger' : 'text-success']">
              {{ item.value.toLocaleString() }}</span
            >
            <span v-else>-</span>
          </td>
        </tr>
      </table>
    </template>
  </PageTableDialog>
</template>
<script setup>
import { useI18n } from 'vue-i18n'
import {
  UserGroupIcon,
  UserIcon,
  ArrowDownCircleIcon,
  BanknotesIcon,
  CurrencyDollarIcon,
} from '@heroicons/vue/20/solid'
import { useDataOverview } from '@/base/composable/operationManagement/dataOverview/useDataOverview'
import { numberToStr } from '@/base/utils/formatNumber'

import Checkbox from '@/base/components/Button/CheckButton.vue'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import DataCardComponent from '@/base/views/OperationManagement/DataOverview/DataCard.vue'

const { t } = useI18n()
const {
  chartTable,
  dialogTitle,
  formInput,
  hours,
  isShowIsSearchAll,
  showChartDialog,
  stats,
  tableData,
  timesText,
  timeTypes,
  searchDataOverview,
  isDeviceLocationInfoEnabled,
} = useDataOverview()
</script>
