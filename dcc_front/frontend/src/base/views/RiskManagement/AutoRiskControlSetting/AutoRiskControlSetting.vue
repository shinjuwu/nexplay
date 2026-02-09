<template>
  <div class="rounded bg-white p-4">
    <CheckButton
      class="mb-4"
      :checked="isEnabledAutoRiskControl"
      :content="t('textEnabledAutoRiskControl')"
      @click="isEnabledAutoRiskControl = !isEnabledAutoRiskControl"
    />
    <ToggleHeader :tips-title="t('textAutoRiskControlSettingDirectionTitle')">
      <template #tipsSlot="{ tipsShow }">
        <div
          v-for="(direction, dirIdx) in directions"
          v-show="tipsShow"
          :key="`directions__${dirIdx}`"
          class="text-danger"
        >
          {{ direction }}
        </div>
        <hr class="my-2" />
      </template>
    </ToggleHeader>
    <PageTable :records="autoSetting.items" :total-records="autoSetting.items.length" :table-input="tableInput">
      <template #default="{ currentRecords }">
        <div class="tbl-container">
          <table class="tbl tbl-hover">
            <thead>
              <tr>
                <th>{{ t('textAutoRisCondition') }}</th>
                <th>{{ t('textConditionParams') }}</th>
                <th>{{ t('textDisposeWay') }}</th>
                <th>{{ t('textDisposeTime') }}</th>
                <th>{{ t('textPreviewDirection') }}</th>
              </tr>
            </thead>
            <tbody>
              <template v-if="currentRecords.length === 0">
                <tr class="no-hover">
                  <td colspan="10">{{ t('textTableEmpty') }}</td>
                </tr>
              </template>
              <template v-else>
                <tr v-for="(setting, index) in currentRecords" :key="`setting__${index}`">
                  <td>{{ setting.name }}</td>
                  <td>
                    <input
                      v-if="isSettingEnabled.AutoRiskControlSettingUpdate"
                      v-model="setting.limit"
                      :disabled="!isEnabledAutoRiskControl"
                      type="number"
                      step="1"
                      class="text-danger w-24 rounded-md border-2 border-gray-200 text-center"
                      @change="setting.limit = Math.floor(setting.limit)"
                    />
                    <span v-else>{{ setting.limit }}</span>
                    <span class="ml-1">{{ setting.unit }}</span>
                  </td>
                  <td>{{ setting.dispose }}</td>
                  <td>{{ setting.time }}</td>
                  <i18n-t :keypath="`fmtAutoRiskSettingPreviewDirection__${index}`" tag="td" scope="global">
                    <template #default>
                      <span class="text-danger">
                        {{ setting.limit }}
                      </span>
                    </template>
                  </i18n-t>
                </tr>
              </template>
            </tbody>
          </table>
        </div>
        <div v-if="isSettingEnabled.AutoRiskControlSettingUpdate" class="mt-4 flex justify-end space-x-2">
          <button type="button" class="btn btn-light" @click="resetDefaultAll()">
            {{ t('textResetDefaultValue') }}
          </button>
          <LoadingButton
            class="btn btn-primary"
            :is-get-data="!tableInput.showProcessing"
            :parent-data="autoRiskSettingDetail"
            :button-click="() => setAutoRiskControlSetting()"
          >
            {{ t('textOnSave') }}
          </LoadingButton>
        </div>
      </template>
    </PageTable>
  </div>
</template>

<script setup>
import { useAutoRiskControlSetting } from '@/base/composable/riskManagement/autoRiskControlSetting/useAutoRiskControlSetting.js'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'

const {
  t,
  autoSetting,
  directions,
  tableInput,
  isSettingEnabled,
  isEnabledAutoRiskControl,
  autoRiskSettingDetail,
  setAutoRiskControlSetting,
  resetDefaultAll,
} = useAutoRiskControlSetting()
</script>
