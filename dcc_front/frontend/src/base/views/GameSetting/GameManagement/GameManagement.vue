<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <FormLabel :content="t('textAgent')" class="mb-1 w-full pr-2 md:w-1/12 md:text-right" />
          <FormAgentListDropdown v-model="formAgentGameInput.agent" class="mb-1 w-full md:w-3/12" />
          <FormLabel :content="t('textGame')" class="mb-1 w-full pr-2 md:w-1/12 md:text-right" />
          <FormGameListDropdown v-model="formAgentGameInput.game" class="mb-1 w-full md:w-3/12" />
          <FormLabel :content="t('textGameState')" class="mb-1 w-full pr-2 md:w-1/12 md:text-right" />
          <FormStateListDropdown v-model="formAgentGameInput.state" class="mb-1 w-full md:w-3/12" />
        </div>
        <div class="mt-4 flex flex-wrap">
          <div class="flex flex-1">
            <div v-show="selectedAgentGames.length > 0" class="flex flex-1 flex-col md:flex-row">
              <button
                type="button"
                class="btn btn-danger mb-1 flex w-full items-center justify-center md:w-2/12 xl:w-fit"
                @click="setAgentGameState(constant.GameState.Online)"
              >
                <CheckCircleIcon class="mr-2 inline-block h-5 w-5" />
                {{ t(`state__${constant.GameState.Online}`) }}
              </button>
              <button
                type="button"
                class="btn btn-danger mb-1 flex w-full items-center justify-center md:ml-2 md:w-2/12 xl:w-fit"
                @click="setAgentGameState(constant.GameState.Maintain)"
              >
                <ExclamationCircleIcon class="mr-2 inline-block h-5 w-5" />
                {{ t(`state__${constant.GameState.Maintain}`) }}
              </button>
              <button
                type="button"
                class="btn btn-danger mb-1 flex w-full items-center justify-center md:ml-2 md:w-2/12 xl:w-fit"
                @click="setAgentGameState(constant.GameState.Offline)"
              >
                <XCircleIcon class="mr-2 inline-block h-5 w-5" />
                {{ t(`state__${constant.GameState.Offline}`) }}
              </button>
            </div>
          </div>
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchAgentGames()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageServersideTable :records="agentGames.items" :table-input="tableAgentGameInput" @search="searchAgentGames">
        <template
          #default="{ currentRecords, pageLength, recordStart, totalPages, totalRecords, lengthChange, pageChange }"
        >
          <div class="tbl-container">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <th v-if="isGameStateUpdateEnabled">
                    <CheckButton
                      :checked="isAgentGameAllSelected"
                      @click="isAgentGameAllSelected = !isAgentGameAllSelected"
                    />
                  </th>
                  <th>{{ t('textAgentName') }}</th>
                  <th>{{ t('textGameCode') }}</th>
                  <th>{{ t('textGameName') }}</th>
                  <th>{{ t('textGameState') }}</th>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="isGameStateUpdateEnabled ? 6 : 5">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="agentGame in currentRecords" :key="`record_${agentGame.agentId}_${agentGame.gameId}`">
                    <td v-if="isGameStateUpdateEnabled">
                      <CheckButton
                        :checked="agentGame.selected"
                        :disabled="agentGame.disabled"
                        @click="toggleAgentGameCheckbox(agentGame)"
                      />
                    </td>
                    <td>{{ agentGame.agentName }}</td>
                    <td>{{ agentGame.gameCode }}</td>
                    <td>{{ t(`game__${agentGame.gameId}`) }}</td>
                    <td>
                      <span
                        :class="{
                          'text-success': agentGame.state === constant.GameState.Online,
                          'text-warning': agentGame.state === constant.GameState.Maintain,
                          'text-danger': agentGame.state === constant.GameState.Offline,
                        }"
                      >
                        {{ t(`state__${agentGame.state}`) }}
                      </span>
                    </td>
                    <td class="space-x-2">
                      <ViewButton
                        :tips-text="t('textViewRoomState')"
                        @click="searchAgentGameRooms(agentGame.agentId, agentGame.gameId, 'view')"
                      />
                      <EditButton
                        v-if="isGameRoomStateUpdateEnabled"
                        :tips-text="t('textSettingRoomState')"
                        @click="searchAgentGameRooms(agentGame.agentId, agentGame.gameId, 'edit')"
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
      </PageServersideTable>
    </template>
  </ToggleHeader>
  <PageTableDialog :visible="dialogAgentGameRoomInput.show" @close="dialogAgentGameRoomInput.show = false">
    <template #header>
      <span>{{ dialogAgentGameRoomInput.agentName }}</span>
      &nbsp;
      <template v-if="dialogAgentGameRoomInput.gameId > constant.Game.All">
        <span>{{ t(`game__${dialogAgentGameRoomInput.gameId}`) }}</span>
        &nbsp;
      </template>
      <span>{{ t('textRoomSetting') }}</span>
    </template>
    <template #default>
      <div class="tbl-container">
        <table class="tbl">
          <thead>
            <tr>
              <th>{{ t('textRoomCode') }}</th>
              <th>{{ t('textRoomType') }}</th>
              <th>{{ t('textRoomState') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="agentGameRoom in agentGameRooms.items"
              :key="`agentGameRoom__${agentGameRoom.agentId}__${agentGameRoom.gameRoomId}`"
            >
              <td>{{ agentGameRoom.gameRoomId }}</td>
              <td>{{ t(`roomType__${roomTypeNameIndex(agentGameRoom.gameId, agentGameRoom.roomType)}`) }}</td>
              <td>
                <template v-if="isGameRoomStateUpdateEnabled && dialogAgentGameRoomInput.mode === 'edit'">
                  <button
                    type="button"
                    class="btn"
                    :class="agentGameRoom.state === constant.GameState.Online ? 'btn-success' : 'btn-secondary'"
                    @click="agentGameRoom.state = constant.GameState.Online"
                  >
                    {{ t(`state__${constant.GameState.Online}`) }}
                  </button>
                  <button
                    type="button"
                    class="btn"
                    :class="agentGameRoom.state === constant.GameState.Maintain ? 'btn-warning' : 'btn-secondary'"
                    @click="agentGameRoom.state = constant.GameState.Maintain"
                  >
                    {{ t(`state__${constant.GameState.Maintain}`) }}
                  </button>
                  <button
                    type="button"
                    class="btn"
                    :class="agentGameRoom.state === constant.GameState.Offline ? 'btn-danger' : 'btn-secondary'"
                    @click="agentGameRoom.state = constant.GameState.Offline"
                  >
                    {{ t(`state__${constant.GameState.Offline}`) }}
                  </button>
                </template>
                <template v-else>
                  <span
                    :class="{
                      'text-success': agentGameRoom.state === constant.GameState.Online,
                      'text-warning': agentGameRoom.state === constant.GameState.Maintain,
                      'text-danger': agentGameRoom.state === constant.GameState.Offline,
                    }"
                  >
                    {{ t(`state__${agentGameRoom.state}`) }}
                  </span>
                </template>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light" @click="dialogAgentGameRoomInput.show = false">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          v-show="isGameRoomStateUpdateEnabled && dialogAgentGameRoomInput.mode === 'edit'"
          class="btn btn-primary ml-2"
          :is-get-data="dialogAgentGameRoomInput.show"
          :parent-data="agentGameRooms.items"
          :button-click="() => setAgentGameRoomState()"
        >
          {{ t('textSend') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { CheckCircleIcon, ExclamationCircleIcon, XCircleIcon } from '@heroicons/vue/20/solid'
import { useGameManagement } from '@/base/composable/gameSetting/gameManagement/useGameManagement'
import constant from '@/base/common/constant'
import { roomTypeNameIndex } from '@/base/utils/room'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormLabel from '@/base/components/Form/FormLabel.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import FormGameListDropdown from '@/base/components/Form/Dropdown/FormGameListDropdown.vue'
import FormStateListDropdown from '@/base/components/Form/Dropdown/FormStateListDropdown.vue'
import PageServersideTable from '@/base/components/Page/Table/PageServersideTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const {
  agentGameRooms,
  agentGames,
  dialogAgentGameRoomInput,
  formAgentGameInput,
  isAgentGameAllSelected,
  isGameRoomStateUpdateEnabled,
  isGameStateUpdateEnabled,
  selectedAgentGames,
  tableAgentGameInput,
  searchAgentGames,
  searchAgentGameRooms,
  setAgentGameState,
  setAgentGameRoomState,
  toggleAgentGameCheckbox,
} = useGameManagement()
</script>
