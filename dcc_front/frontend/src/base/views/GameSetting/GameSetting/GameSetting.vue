<template>
  <ToggleHeader class="relative mt-4">
    <template #default="{ show }">
      <div v-show="show">
        <div class="my-2 md:flex md:justify-between">
          <div class="text-gray-500">
            <template v-for="state in gameStates" :key="`gameState__${state}`">
              <CheckButton
                class="mr-1"
                size="md"
                :content="state === constant.GameState.All ? t('textStateAll') : t(`state__${state}`)"
                :checked="gameState === state"
                @click="gameState = state"
              />
            </template>
          </div>
          <div class="text-danger">
            {{ t('fmtTextGameSettingInfo', [games.items.length, onlineGameCount, maintainGameCount]) }}
          </div>
        </div>
        <PageTable :records="filterGames" :total-records="filterGames.length" :table-input="gameTaleInput">
          <template #default="{ currentRecords, isSortIconActive, sorting }">
            <div class="tbl-container">
              <table class="tbl tbl-hover">
                <thead>
                  <tr>
                    <PageTableSortableTh column="id" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textGameId') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="code" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textGameCode') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="id" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textGameName') }} </template>
                    </PageTableSortableTh>
                    <PageTableSortableTh column="state" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                      <template #default> {{ t('textGameState') }} </template>
                    </PageTableSortableTh>
                    <th v-if="isSettingEnabled">{{ t('textOperate') }}</th>
                  </tr>
                </thead>
                <tbody>
                  <template v-if="currentRecords.length === 0">
                    <tr class="no-hover">
                      <td :colspan="isSettingEnabled ? 5 : 4">
                        {{ t('textTableEmpty') }}
                      </td>
                    </tr>
                  </template>
                  <template v-else>
                    <tr v-for="game in currentRecords" :key="`game_${game.id}`" class="hover:bg-slate-100">
                      <td>{{ game.id }}</td>
                      <td>{{ game.code }}</td>
                      <td>{{ t(`game__${game.id}`) }}</td>
                      <td :class="{ 'space-x-2': isSettingEnabled }">
                        <template v-if="isSettingEnabled">
                          <button
                            type="button"
                            class="btn"
                            :class="game.state === constant.GameState.Online ? 'btn-success' : 'btn-secondary'"
                            @click="setGameState(game, constant.GameState.Online)"
                          >
                            {{ t(`state__${constant.GameState.Online}`) }}
                          </button>
                          <button
                            type="button"
                            class="btn"
                            :class="game.state === constant.GameState.Maintain ? 'btn-warning' : 'btn-secondary'"
                            @click="setGameState(game, constant.GameState.Maintain)"
                          >
                            {{ t(`state__${constant.GameState.Maintain}`) }}
                          </button>
                          <button
                            type="button"
                            class="btn"
                            :class="game.state === constant.GameState.Offline ? 'btn-danger' : 'btn-secondary'"
                            @click="setGameState(game, constant.GameState.Offline)"
                          >
                            {{ t(`state__${constant.GameState.Offline}`) }}
                          </button>
                        </template>
                        <template v-else>
                          <span
                            :class="{
                              'text-success': game.state === constant.GameState.Online,
                              'text-warning': game.state === constant.GameState.Maintain,
                              'text-danger': game.state === constant.GameState.Offline,
                            }"
                          >
                            {{ t(`state__${game.state}`) }}
                          </span>
                        </template>
                      </td>
                      <td v-if="isSettingEnabled">
                        <ButtonTooltips :tips-text="t('textSyncGameServer')">
                          <template #content>
                            <label
                              class="inline-flex cursor-pointer items-center justify-center"
                              @click="notifyGameToGameServer(game)"
                            >
                              <ArrowUpOnSquareIcon class="inline h-5 w-5 rounded" />
                            </label>
                          </template>
                        </ButtonTooltips>
                      </td>
                    </tr>
                  </template>
                </tbody>
              </table>
            </div>
          </template>
        </PageTable>
      </div>
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { ArrowUpOnSquareIcon } from '@heroicons/vue/24/outline'
import constant from '@/base/common/constant'

import { useGameSetting } from '@/base/composable/gameSetting/gameSetting/useGameSetting'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'

const { t } = useI18n()
const {
  filterGames,
  games,
  gameState,
  gameStates,
  gameTaleInput,
  isSettingEnabled,
  maintainGameCount,
  onlineGameCount,
  notifyGameToGameServer,
  setGameState,
} = useGameSetting()
</script>
