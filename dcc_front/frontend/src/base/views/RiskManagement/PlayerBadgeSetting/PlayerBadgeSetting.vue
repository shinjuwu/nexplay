<template>
  <div class="rounded bg-white p-4">
    <PageTable :records="badges" :total-records="badges.length" :table-input="tableInput">
      <template #default="{ currentRecords }">
        <div class="tbl-container">
          <table class="tbl">
            <thead>
              <tr>
                <th>{{ t('textBadgeType') }}</th>
                <th>{{ t('textBadgeName') }}</th>
                <th>{{ t('textBackgroundColor') }}</th>
                <th>{{ t('textTextColor') }}</th>
                <th>{{ t('textPreview') }}</th>
                <th>{{ t('textRemark') }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="badge in currentRecords" :key="`badge_${badge.index}`">
                <td>{{ badge.isEditable ? t('textCustomize') : t('textDefault') }}</td>
                <td>
                  <input
                    v-if="badge.isEditable && isSettingEnabled"
                    v-model="badge.name"
                    class="form-input"
                    type="text"
                    maxlength="4"
                    :placeholder="t('placeHolderTextPlayerBadgeSettingBadgeName')"
                  />
                  <span v-else>{{ badge.name }}</span>
                </td>
                <td>
                  <span
                    class="inline-flex h-5 w-5 rounded border border-black"
                    :style="{ backgroundColor: badge.bgColor }"
                  >
                    <input
                      v-if="badge.isEditable && isSettingEnabled"
                      v-model="badge.bgColor"
                      class="h-5 w-5 cursor-pointer rounded border-none opacity-0"
                      type="color"
                    />
                  </span>
                </td>
                <td>
                  <span
                    class="inline-flex h-5 w-5 rounded border border-black"
                    :style="{ backgroundColor: badge.txtColor }"
                  >
                    <input
                      v-if="badge.isEditable && isSettingEnabled"
                      v-model="badge.txtColor"
                      class="h-5 w-5 cursor-pointer rounded border-none opacity-0"
                      type="color"
                    />
                  </span>
                </td>
                <td>
                  <span
                    class="rounded-full px-1 text-xs"
                    :style="{ backgroundColor: badge.bgColor, color: badge.txtColor }"
                  >
                    {{ badge.name }}
                  </span>
                </td>
                <td>
                  <input
                    v-if="badge.isEditable && isSettingEnabled"
                    v-model="badge.info"
                    class="form-input"
                    type="text"
                    maxlength="30"
                    :placeholder="t('placeHolderTextPlayerBadgeSettingBadgeInfo')"
                  />
                  <template v-else>{{ badge.info }}</template>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div v-if="isSettingEnabled" class="mt-4 flex justify-end space-x-2">
          <button type="button" class="btn btn-light" @click="resetAllCustomBadges()">
            {{ t('textResetDefaultValue') }}
          </button>
          <LoadingButton
            class="btn btn-primary"
            :is-get-data="!tableInput.showProcessing"
            :parent-data="badges"
            :button-click="() => setAgentCustomTagSettingList()"
          >
            {{ t('textOnSave') }}
          </LoadingButton>
        </div>
      </template>
    </PageTable>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { usePlayerBadgeSetting } from '@/base/composable/riskManagement/playerBadgeSetting/usePlayerBadgeSetting'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const { badges, isSettingEnabled, tableInput, resetAllCustomBadges, setAgentCustomTagSettingList } =
  usePlayerBadgeSetting()
</script>
