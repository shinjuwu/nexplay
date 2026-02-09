<template>
  <PageTableDialog :visible="visible" @close="close()">
    <template #header>{{ t('textMemberAccountDetail') }}</template>
    <template #default>
      <div class="grid gap-2 md:grid-cols-2 md:gap-4">
        <div>
          <label class="form-label">{{ t('fmtTextUserName', [playerInfo.userName]) }}</label>
        </div>
        <div>
          <label class="form-label">{{ t('fmtTextState', [t(`status__${playerInfo.state}`)]) }}</label>
        </div>
        <div>
          <label class="form-label">
            {{ t('textValidBetTotal') }}
            <span class="text-danger">{{ t('textMemberAccountIncludeTodayData') }}</span>
          </label>
          <input v-model="playerInfo.sumValidBet" class="form-input" type="text" disabled />
        </div>
        <div>
          <label class="form-label">
            {{ t('textTotalWinLoseScore') }}
            <span class="text-danger">{{ t('textMemberAccountIncludeTodayData') }}</span>
          </label>
          <input v-model="playerInfo.sumProfit" class="form-input" type="text" disabled />
        </div>
        <div>
          <label class="form-label">
            {{ t('textMemberAccountMonthValidBet') }}
            <span class="text-danger">{{ t('textMemberAccountIncludeTodayData') }}</span>
          </label>
          <input v-model="playerInfo.monthValidBet" class="form-input" type="text" disabled />
        </div>
        <div>
          <label class="form-label">
            {{ t('textMonthWinLoseScore') }}
            <span class="text-danger">{{ t('textMemberAccountIncludeTodayData') }}</span>
          </label>
          <input v-model="playerInfo.monthProfit" class="form-input" type="text" disabled />
        </div>
        <div class="md:col-span-2">
          <label class="form-label">
            {{ t('textRemark') }}
          </label>
          <textarea v-model="playerInfo.remark" class="form-input" type="text" maxlength="100" disabled />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto">
        <button type="button" class="btn btn-light" @click="close()">
          {{ t('textClose') }}
        </button>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { usePlayerAccountDetail } from '@/base/composable/operationManagement/playerAccount/usePlayerAccountDetail'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  playerInfo: {
    type: Object,
    default: () => {},
  },
})
const emit = defineEmits(['close'])

const { t } = useI18n()
const { playerInfo, visible, close } = usePlayerAccountDetail(props, emit)
</script>
