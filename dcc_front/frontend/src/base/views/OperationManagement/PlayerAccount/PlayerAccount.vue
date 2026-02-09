<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formInput.agent" :agent-id="props.agentId" class="mb-1 w-full md:w-3/12" />
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textPlayerAccount') }}</label>
          <input
            v-model="formInput.userName"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportPlayerAccount')"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchGameUsers()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>

      <hr class="my-2" />

      <ToggleHeader :tips-title="t('textSpecialDirections')">
        <template #tipsSlot="{ tipsShow }">
          <div v-show="tipsShow" class="text-danger">
            <div>
              {{ t('textPlayerAccountIsNewbieTotalDirection') }}
            </div>
            <div>
              {{ t('textPlayerAccountIsNewbieSingleGameDirection') }}
            </div>
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
                  <PageTableSortableTh column="userName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textPlayerAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="badgeInfo" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
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
                  <PageTableSortableTh
                    column="riskControlTagList"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textDispose') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="agentName" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="coinIn" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textMemberAccountCoinIn') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="coinOut" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textMemberAccountCoinOut') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="createTime" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textCreateDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="lastLoginTime"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textLastLoginDate') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="isOnline" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textIsOnline') }} </template>
                  </PageTableSortableTh>
                  <th v-if="showRealTimeBalanceColumn">{{ t('textRealTimeBalance') }}</th>
                  <th>{{ t('textIsNewbie') }}</th>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td :colspan="tableColumns">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="gameUser in currentRecords" :key="`gameUser_${gameUser.id}`">
                    <td>{{ gameUser.userName }}</td>
                    <td>
                      <template v-if="gameUser.killDiveState === 2">
                        <span
                          class="rounded-full px-1 text-xs"
                          :style="{
                            backgroundColor: defaultBadgeInfo.Blacklist.BackgroundColor,
                            color: defaultBadgeInfo.Blacklist.TextColor,
                          }"
                          >{{ t(`riskType__${defaultBadgeInfo.Blacklist.Index + 4}`) }}
                        </span>
                        <br />
                      </template>
                      <template v-if="gameUser.isHighRisk">
                        <span
                          class="rounded-full px-1 text-xs"
                          :style="{
                            backgroundColor: defaultBadgeInfo.HighRisk.BackgroundColor,
                            color: defaultBadgeInfo.HighRisk.TextColor,
                          }"
                          >{{ t(`riskType__${defaultBadgeInfo.HighRisk.Index + 4}`) }}
                        </span>
                        <br />
                      </template>
                      <template v-if="gameUser.killDiveState === 1">
                        <span
                          class="rounded-full px-1 text-xs"
                          :style="{
                            backgroundColor: defaultBadgeInfo.TargetKill.BackgroundColor,
                            color: defaultBadgeInfo.TargetKill.TextColor,
                          }"
                          >{{ t(`riskType__${defaultBadgeInfo.TargetKill.Index + 4}`) }}
                        </span>
                        <br />
                      </template>
                      <template v-for="customBadgeIdx in gameUser.badges" :key="`customBadge__${customBadgeIdx}`">
                        <template v-if="myCustomBadges.items[customBadgeIdx].name !== ''">
                          <span
                            class="rounded-full px-1 text-xs"
                            :style="{
                              backgroundColor: myCustomBadges.items[customBadgeIdx].backgroundColor,
                              color: myCustomBadges.items[customBadgeIdx].textColor,
                            }"
                          >
                            {{ myCustomBadges.items[customBadgeIdx].name }}
                          </span>
                          <br />
                        </template>
                      </template>
                    </td>
                    <td :class="{ 'text-danger': !gameUser.state }">
                      {{
                        t(`status__${gameUser.state ? constant.AccountStatus.Open : constant.AccountStatus.Disable}`)
                      }}
                    </td>
                    <td>
                      {{
                        gameUser.killDiveState === constant.KillDive.ConfigKill && gameUser.killDiveValue > 0
                          ? numberToStr(gameUser.killDiveValue)
                          : '-'
                      }}
                    </td>
                    <td>
                      <template v-if="gameUser.riskControlTag.length">
                        <div v-for="(tag, index) in gameUser.riskControlTag" :key="`gameUserRiskTags__${index}`">
                          {{ tag }}
                        </div>
                      </template>
                      <template v-else>-</template>
                    </td>
                    <td>{{ gameUser.agentName }}</td>
                    <td>{{ numberToStr(gameUser.coinIn) }}</td>
                    <td>
                      <span :class="{ 'text-danger': gameUser.coinOut > 0 }">{{ numberToStr(gameUser.coinOut) }}</span>
                    </td>
                    <td>{{ gameUser.createTimeStr }}</td>
                    <td>{{ gameUser.lastLoginTimeStr }}</td>
                    <td>
                      <span :class="{ 'text-danger': gameUser.isOnline }">
                        {{ gameUser.isOnline ? t('textYes') : t('textNo') }}
                      </span>
                    </td>
                    <td v-if="showRealTimeBalanceColumn">
                      <button
                        v-if="gameUser.walletType === constant.AgentWallet.Single"
                        class="btn btn-secondary cursor-default"
                      >
                        {{ t(`agentWalletOption__${constant.AgentWallet.Single}`) }}
                      </button>
                      <button
                        v-else
                        class="btn btn-success mx-auto flex items-center"
                        @click="getPlayerWalletBalance(gameUser)"
                      >
                        <span class="mr-2">
                          {{
                            gameUser.isSearchWalletBalance ? numberToStr(gameUser.walletBalance) : t('textClickUpdate')
                          }}
                        </span>
                        <ArrowPathIcon class="h-4 w-4" />
                      </button>
                    </td>
                    <td>
                      <button
                        class="btn btn-success mx-auto flex items-center"
                        @click="getPlayerPlayCountData(gameUser)"
                      >
                        <span class="mr-2">
                          {{
                            gameUser.isSearchGameUserPlayCount
                              ? gameUser.isNewbie
                                ? t('textYes')
                                : t('textNo')
                              : t('textClickUpdate')
                          }}
                        </span>
                        <ArrowPathIcon class="h-4 w-4" />
                      </button>
                    </td>
                    <td class="space-x-2">
                      <ViewButton @click="selectPlayerInfo(gameUser, 'detail')" />
                      <EditButton
                        v-if="isSettingEnabled.PlayerAccountInfoUpdate"
                        @click="selectPlayerInfo(gameUser, 'setting')"
                      />
                      <BadgeButton
                        v-if="isSettingEnabled.PlayerBadgeUpdate && myAgentId == gameUser.agentId"
                        @click="selectPlayerInfo(gameUser, 'badge')"
                      />
                      <LockButton
                        v-if="isSettingEnabled.PlayerDisposeSettingUpdate"
                        @click="selectPlayerInfo(gameUser, 'dispose')"
                      >
                      </LockButton>
                      <ButtonTooltips
                        v-if="isSettingEnabled.PlayerLogRead"
                        :tips-text="t('menuItemPlayerLog')"
                        @click="redirectToPlayerLog(gameUser.agentId, gameUser.userName)"
                      >
                        <template #content>
                          <CurrencyDollarIcon class="inline-block h-5 w-5" />
                        </template>
                      </ButtonTooltips>
                      <ButtonTooltips :tips-text="t('textPlayInfo')" @click="selectPlayerInfo(gameUser, 'playInfo')">
                        <template #content>
                          <ClipboardDocumentListIcon class="inline-block h-5 w-5" />
                        </template>
                      </ButtonTooltips>
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
  <PlayerAccountDetail
    :visible="dialog.detail"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.detail = newValue)"
  />
  <PlayerAccountSetting
    v-if="isSettingEnabled"
    :visible="dialog.setting"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.setting = newValue)"
    @search-game-users="searchGameUsers()"
  />
  <PlayerBadgeDialog
    :visible="dialog.badge"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.badge = newValue)"
    @search-game-users="searchGameUsers()"
  />
  <PlayerDisposeSettingDialog
    :visible="dialog.dispose"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.dispose = newValue)"
    @search-game-users="searchGameUsers()"
  />
  <PlayerAccountPlayInfo
    :visible="dialog.playInfo"
    :player-info="playerInfo"
    @close="(newValue) => (dialog.playInfo = newValue)"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { usePlayerAccount } from '@/base/composable/operationManagement/playerAccount/usePlayerAccount'
import constant from '@/base/common/constant'
import { numberToStr } from '@/base/utils/formatNumber'
import { CurrencyDollarIcon } from '@heroicons/vue/24/outline'
import { ArrowPathIcon } from '@heroicons/vue/24/solid'
import { ClipboardDocumentListIcon } from '@heroicons/vue/24/outline'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import BadgeButton from '@/base/components/Button/BadgeButton.vue'
import LockButton from '@/base/components/Button/LockButton.vue'
import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'
import PlayerAccountDetail from '@/base/views/OperationManagement/PlayerAccount/PlayerAccountDetail.vue'
import PlayerAccountPlayInfo from '@/base/views/OperationManagement/PlayerAccount/PlayerAccountPlayInfo.vue'
import PlayerAccountSetting from '@/base/views/Common/Dialog/PlayerAccountSettingDialog.vue'
import PlayerBadgeDialog from '@/base/views/Common/Dialog/PlayerBadgeDialog.vue'
import PlayerDisposeSettingDialog from '@/base/views/Common/Dialog/PlayerDisposeSettingDialog.vue'

const props = defineProps({
  userName: {
    type: String,
    default: '',
  },
  agentId: {
    type: Number,
    default: -1,
  },
})

const { t } = useI18n()
const {
  defaultBadgeInfo,
  dialog,
  formInput,
  isSettingEnabled,
  myAgentId,
  myCustomBadges,
  playerInfo,
  records,
  showRealTimeBalanceColumn,
  tableInput,
  tableColumns,
  getPlayerWalletBalance,
  getPlayerPlayCountData,
  searchGameUsers,
  selectPlayerInfo,
  redirectToPlayerLog,
} = usePlayerAccount(props)
</script>
