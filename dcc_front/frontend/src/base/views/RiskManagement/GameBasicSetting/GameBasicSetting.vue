<template>
  <ToggleHeader>
    <template #default="{ show }">
      <div v-show="show">
        <div v-for="(direction, dirIdx) in pageDirections.items" :key="`directions__${dirIdx}`" class="text-danger">
          {{ direction }}
        </div>
        <div v-if="isEditEnabled" class="mt-1 flex flex-1 justify-end">
          <button
            class="btn btn-primary flex w-full items-center justify-around md:w-auto"
            type="button"
            @click="setGameSettings()"
          >
            {{ t('textOnSave') }}
          </button>
        </div>
      </div>

      <hr class="my-2" />

      <div class="tbl-container">
        <table class="tbl tbl-hover">
          <thead>
            <tr>
              <th>{{ t('textGameId') }}</th>
              <th>{{ t('textGameName') }}</th>
              <th>{{ t('textMatchGameRTP') }}</th>
              <th>{{ t('textMatchGameKillRate') }}</th>
              <th>{{ t('textMatchGames') }}</th>
              <th>{{ t('textNormalMatchGameRTP') }}</th>
              <th>{{ t('textNormalMatchGameKillRate') }}</th>
              <th>{{ t('textLowBoundRTP') }}</th>
              <th>{{ t('textLimitOdds') }}</th>
              <th>{{ t('textLimitMoney') }}</th>
            </tr>
          </thead>
          <tbody>
            <template v-if="gameSettings.items.length === 0">
              <tr class="no-hover">
                <td colspan="10">{{ t('textTableEmpty') }}</td>
              </tr>
            </template>
            <template v-else>
              <tr v-for="gameSetting in gameSettings.items" :key="`gameSetting_${gameSetting.gameId}`">
                <td>{{ gameSetting.gameId }}</td>
                <td>{{ t(`game__${gameSetting.gameId}`) }}</td>
                <td>
                  <template v-if="isMatchGameRTPEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.matchGameRTP"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="100"
                      step="0.1"
                    />
                    <span class="mr-2">%</span>
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isMatchGameKillRateEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.matchGameKillRate"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="100"
                      step="0.1"
                    />
                    <span class="mr-2">%</span>
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isMatchGamesEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.matchGames"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="9999"
                      step="1"
                    />
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isNormalMatchGameRTPEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.normalMatchGameRTP"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="100"
                      step="0.1"
                    />
                    <span class="mr-2">%</span>
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isNormalMatchGameKillRateEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.normalMatchGameKillRate"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="100"
                      step="0.1"
                    />
                    <span class="mr-2">%</span>
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isLowBoundRTPEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.lowBoundRTP"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="100"
                      step="0.1"
                    />
                    <span class="mr-2">%</span>
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isLimitOddsEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.limitOdds"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="9999"
                      step="0.01"
                    />
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
                <td>
                  <template v-if="isLimitMoneyEnabled(gameSetting.gameId)">
                    <input
                      v-model="gameSetting.limitMoney"
                      class="form-input"
                      type="number"
                      :disabled="!isEditEnabled"
                      min="0"
                      max="999999"
                      step="0.01"
                    />
                  </template>
                  <template v-else>
                    <span>-</span>
                  </template>
                </td>
              </tr>
            </template>
          </tbody>
        </table>

        <PageTableLoading v-show="showProcessing" />
      </div>
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useGameBasicSetting } from '@/base/composable/riskManagement/gameBasicSetting/useGameBasicSetting'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'

const { t } = useI18n()
const {
  gameSettings,
  isEditEnabled,
  isMatchGameRTPEnabled,
  isMatchGameKillRateEnabled,
  isMatchGamesEnabled,
  isNormalMatchGameRTPEnabled,
  isNormalMatchGameKillRateEnabled,
  isLowBoundRTPEnabled,
  isLimitOddsEnabled,
  isLimitMoneyEnabled,
  pageDirections,
  showProcessing,
  setGameSettings,
} = useGameBasicSetting()
</script>
