<template>
  <ToggleHeader>
    <template #default="{ show }">
      <div v-show="show">
        <div class="text-danger">{{ t('textExchangeRateDirections') }}</div>
        <div
          v-for="(direction, dirIdx) in exchangeDataDirections.items"
          :key="`directions__${dirIdx}`"
          class="text-danger"
        >
          {{ direction }}
        </div>
      </div>
      <hr class="my-2" />
      <PageTable :records="exchangeData.items" :total-records="exchangeData.items.length" :table-input="tableInput">
        <template #default="{ currentRecords }">
          <div class="tbl-container">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <th>{{ t('textCurrency') }}</th>
                  <th>{{ t('textExchangeRate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="2">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.currency}`">
                    <td>{{ `${t(`currency__${record.currency}`)}(${record.currency})` }}</td>
                    <td>
                      <template v-if="isSettingEnabled">
                        <input v-model.number="record.toCoin" step="0.0001" class="form-input" type="number" />
                      </template>
                      <template v-else>{{ record.toCoin }}</template>
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
          <div v-if="isSettingEnabled" class="mt-4 flex justify-end space-x-2">
            <LoadingButton
              class="btn btn-primary"
              :is-get-data="!tableInput.showProcessing"
              :parent-data="exchangeData"
              :button-click="() => setExchangeDataList()"
            >
              {{ t('textOnSave') }}
            </LoadingButton>
          </div>
        </template>
      </PageTable>
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useExchangeRateSetting } from '@/base/composable/systemManagement/exchangeRateSetting/useExchangeRateSetting'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const { exchangeData, exchangeDataDirections, isSettingEnabled, tableInput, setExchangeDataList } =
  useExchangeRateSetting()
</script>
