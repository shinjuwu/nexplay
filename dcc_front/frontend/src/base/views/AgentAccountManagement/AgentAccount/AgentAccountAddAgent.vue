<template>
  <PageTableDialog :visible="props.visible" @close="close">
    <template #header>
      {{ header }}
    </template>
    <template #default>
      <div class="grid max-h-[70vh] grid-cols-2 gap-6 overflow-y-auto">
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textAccount') }}</label>
          <input
            v-model="newAgentForm.account"
            class="form-input"
            type="text"
            :placeholder="t('placeHolderTextFormAccount')"
            maxlength="16"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textPassword') }}</label>
          <input
            v-model="newAgentForm.password"
            autocomplete="new-password"
            class="form-input"
            type="password"
            :placeholder="t('placeHolderTextFormPassword')"
            maxlength="16"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textAgentName') }}</label>
          <input
            v-model="newAgentForm.agentName"
            class="form-input"
            type="text"
            :placeholder="t('placeHolderTextFormAgentName')"
            maxlength="16"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label">{{ t('textRatio') }}</label>
          <PercentageNumberInput
            v-model="newAgentForm.ratio"
            class="form-input"
            @update:model-value="(newValue) => (newAgentForm.ratio = newValue)"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">
            {{ t('textBackStageAllowListIP') }}
          </label>
          <input
            v-model="newAgentForm.allowList"
            class="form-input"
            type="text"
            :placeholder="t('placeHolderTextFormBackStageAllowList')"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">
            {{ t('textCooperateMethod') }}
          </label>
          <select v-model="newAgentForm.cooperate" class="form-input" :disabled="isNotAdmin">
            <option
              v-for="option in cooperateOptions"
              :key="option"
              :value="option"
              :label="t(`agentCooperation__${option}`)"
            ></option>
          </select>
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">
            {{ t('textCurrency') }}
          </label>
          <select v-model="newAgentForm.currency" class="form-input" :disabled="isNotAdmin">
            <option v-for="cur in currencyOptions" :key="cur" :value="cur" :label="t(`currency__${cur}`)"></option>
          </select>
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textGroupRole') }}</label>
          <FormAgentPermissionListDropdown v-model="newAgentForm.role" class="form-input" :account-type="accountType" />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">
            {{ t('textWalletType') }}
          </label>
          <select v-model="newAgentForm.walletType" class="form-input" :disabled="isNotAdmin">
            <option
              v-for="walletOption in agentWalletOptions"
              :key="`agentWalletOption__${walletOption}`"
              :value="walletOption"
              :label="t(`agentWalletOption__${walletOption}`)"
            ></option>
          </select>
        </div>
        <div class="col-span-2 md:col-span-1">
          <label
            class="form-label"
            :class="{ 'required-mark': newAgentForm.walletType === constant.AgentWallet.Single }"
          >
            {{ t('textWalletApiPath') }}
          </label>
          <input
            v-model="newAgentForm.walletConnInfo"
            class="form-input"
            type="text"
            :disabled="newAgentForm.walletType !== constant.AgentWallet.Single"
          />
        </div>
        <div v-if="!isNotAdmin" class="col-span-2 space-x-3 md:col-span-1">
          <label
            v-for="agentLobbySwitch in agentLobbySwitches"
            :key="`agentLobbySwitch__${agentLobbySwitch}`"
            class="form-label space-x-1"
          >
            <input v-model="newAgentForm.lobbySwitch" type="checkbox" :value="agentLobbySwitch" />
            <span>{{ t(`lobbySwitch__${agentLobbySwitch}`) }}</span>
          </label>
        </div>
        <div v-if="!isNotAdmin" class="col-span-2 md:col-span-1">
          <label class="form-label space-x-1">
            <input v-model="newAgentForm.cannedSwitch" type="checkbox" />
            <span>{{ t('textEmojiCannedLanguage') }}</span>
          </label>
        </div>
        <div class="col-span-2">
          <label class="form-label">{{ t('textRemark') }}</label>
          <textarea v-model="newAgentForm.remark" class="form-input" type="text" maxlength="100" rows="4" />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="props.visible"
          :parent-data="newAgentForm"
          :button-click="() => createNewAgent()"
        >
          {{ t('textIncrease') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useAgentAccountAddAgent } from '@/base/composable/agentAccountManagement/agentAccount/useAgentAccountAddAgent'
import constant from '@/base/common/constant'

import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import PercentageNumberInput from '@/base/components/NumberInput/PercentageNumberInput.vue'
import FormAgentPermissionListDropdown from '@/base/components/Form/Dropdown/FormAgentPermissionListDropdown.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['close', 'searchRecords'])

const { t } = useI18n()
const {
  accountType,
  agentWalletOptions,
  agentLobbySwitches,
  header,
  isNotAdmin,
  newAgentForm,
  cooperateOptions,
  currencyOptions,
  createNewAgent,
  close,
} = useAgentAccountAddAgent(emit)
</script>
