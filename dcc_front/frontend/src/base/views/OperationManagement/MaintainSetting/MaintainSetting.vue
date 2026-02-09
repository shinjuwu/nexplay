<template>
  <ToggleHeader class="relative">
    <template #default="{ show }">
      <div v-show="show" class="tbl-container">
        <table class="tbl tbl-hover">
          <thead>
            <tr>
              <th>{{ t('textGameServer') }}</th>
              <th>{{ t('textGameState') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{{ t('textGameAll') }}</td>
              <td class="space-x-2">
                <template v-if="isSettingEnabled">
                  <button
                    type="button"
                    class="btn"
                    :class="gameServerState === constant.GameState.Online ? 'btn-success' : 'btn-secondary'"
                    @click="setGameServerState(constant.GameState.Online)"
                  >
                    {{ t(`state__${constant.GameState.Online}`) }}
                  </button>
                  <button
                    type="button"
                    class="btn"
                    :class="gameServerState === constant.GameState.Maintain ? 'btn-warning' : 'btn-secondary'"
                    @click="setGameServerState(constant.GameState.Maintain)"
                  >
                    {{ t(`state__${constant.GameState.Maintain}`) }}
                  </button>
                </template>
                <template v-else>
                  <span
                    :class="{
                      'text-success': gameServerState === constant.GameState.Online,
                      'text-warning': gameServerState === constant.GameState.Maintain,
                    }"
                  >
                    {{ t(`state__${gameServerState}`) }}
                  </span>
                </template>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <PageTableLoading v-show="showGameServerProcessing" />
    </template>
  </ToggleHeader>
  <ToggleHeader class="relative mt-4">
    <template #default="{ show }">
      <table v-show="show" class="tbl tbl-hover">
        <thead>
          <tr>
            <th>{{ t('textTimezone') }}</th>
            <th>{{ t('textStartTime') }}</th>
            <th>{{ t('textEndTime') }}</th>
            <th>{{ t('textPreviewDirection') }}</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>
              <FormInput v-model="maintainPageSetting.timezone" maxlength="20" :disabled="!isSettingEnabled" />
            </td>
            <td>
              <FormDateTimeInput v-model="maintainPageStartTime" :disabled="!isSettingEnabled" />
            </td>
            <td>
              <FormDateTimeInput v-model="maintainPageEndTime" :disabled="!isSettingEnabled" />
            </td>
            <td>
              {{ maintainPageSetting.timezone }} {{ maintainPageSetting.startTime }} ~ {{ maintainPageSetting.endTime }}
            </td>
          </tr>
        </tbody>
      </table>
      <div class="mt-4 flex justify-end">
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="!showMaintainPageProcessing"
          :parent-data="maintainPageSetting"
          :button-click="() => setMaintainPageSetting()"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
      <PageTableLoading v-show="showMaintainPageProcessing" />
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'

import { useMaintainSetting } from '@/base/composable/operationManagement/maintainSetting/useMaintainSetting'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'
import FormInput from '@/base/components/Form/FormInput.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const {
  gameServerState,
  isSettingEnabled,
  maintainPageEndTime,
  maintainPageSetting,
  maintainPageStartTime,
  showGameServerProcessing,
  showMaintainPageProcessing,
  setGameServerState,
  setMaintainPageSetting,
} = useMaintainSetting()
</script>
