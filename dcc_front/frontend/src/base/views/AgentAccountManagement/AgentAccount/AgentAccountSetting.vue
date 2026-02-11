<template>
  <PageTableDialog :visible="show" @close="close">
    <template #header> {{ t('textAgentSettingAgentSet') }} </template>
    <template #default>
      <div class="grid grid-cols-2 gap-4">
        <label class="form-label col-span-2 md:col-span-1">{{ t('fmtTextAgentSettingAgentId', [agent.id]) }}</label>
        <label class="form-label col-span-2 md:col-span-1">
          {{ t('fmtTextCurrentBalance') }}<span class="text-danger">{{ agent.balance }}</span>
        </label>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textAgentName') }}</label>
          <input
            v-model="agent.name"
            class="form-input"
            type="text"
            :placeholder="t('placeHolderTextFormAgentName')"
            :disabled="!props.isEditEnabled"
            maxlength="16"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label">{{ t('textRatio') }}</label>
          <PercentageNumberInput
            v-model="agent.ratio"
            class="form-input"
            :is-disabled="!props.isEditEnabled"
            @update:model-value="(newValue) => (agent.ratio = newValue)"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">
            {{ t('textCooperateMethod') }}
          </label>
          <input class="form-input" disabled :value="t(`agentCooperation__${agent.cooperate}`)" />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textGroupRole') }}</label>
          <FormAgentPermissionListDropdown
            v-if="props.isEditEnabled"
            v-model="agent.role"
            class="form-input"
            :account-type="user.accountType + 1"
          />
          <input v-else class="form-input" type="text" disabled :value="agent.roleName" />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">
            {{ t('textWalletType') }}
          </label>
          <input class="form-input" disabled :value="t(`agentWalletOption__${agent.walletType}`)" />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label" :class="{ 'required-mark': agent.walletType === constant.AgentWallet.Single }">
            {{ t('textWalletApiPath') }}
          </label>
          <input
            v-model="agent.walletConnInfo"
            :disabled="!props.isEditEnabled || agent.walletType !== constant.AgentWallet.Single"
            class="form-input"
            type="text"
          />
        </div>
        <!-- <div class="col-span-2 md:col-span-1">
          <div v-if="props.isEditEnabled && user.accountType === constant.AccountType.Admin" class="text-danger">
            {{ t('textLobbySwitchDirection') }}
          </div>
          <div class="space-x-3">
            <label
              v-for="agentLobbySwitch in agentLobbySwitches"
              :key="`agentLobbySwitch__${agentLobbySwitch}`"
              class="form-label space-x-1"
            >
              <input
                v-model="agent.lobbySwitch"
                type="checkbox"
                :value="agentLobbySwitch"
                :disabled="!props.isEditEnabled || user.accountType > constant.AccountType.Admin"
              />
              <span>{{ t(`lobbySwitch__${agentLobbySwitch}`) }}</span>
            </label>
          </div>
        </div>
        <div class="col-span-2 md:col-span-1">
          <div v-if="props.isEditEnabled && user.accountType === constant.AccountType.Admin" class="text-danger">
            {{ t('textEmojiCannedDirection') }}
          </div>
          <div>
            <label class="form-label space-x-1">
              <input
                v-model="agent.cannedSwitch"
                type="checkbox"
                :disabled="!props.isEditEnabled || user.accountType > constant.AccountType.Admin"
              />
              <span>{{ t('textEmojiCannedLanguage') }}</span>
            </label>
          </div>
        </div> -->
        <div class="col-span-2">
          <label class="form-label">{{ t('textRemark') }}</label>
          <textarea
            v-model="agent.remark"
            class="form-input"
            type="text"
            :disabled="!props.isEditEnabled"
            maxlength="100"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          v-show="props.isEditEnabled"
          class="btn btn-primary"
          :is-get-data="show"
          :parent-data="agent"
          :button-click="() => submit()"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useAgentAccountSetting } from '@/base/composable/agentAccountManagement/agentAccount/useAgentAccountSetting'
import constant from '@/base/common/constant'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import FormAgentPermissionListDropdown from '@/base/components/Form/Dropdown/FormAgentPermissionListDropdown.vue'
import PercentageNumberInput from '@/base/components/NumberInput/PercentageNumberInput.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  selectRecord: {
    type: Object,
    default: () => {},
  },
  isEditEnabled: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['searchRecords', 'close'])

const { t } = useI18n()
const { user, show, agent, submit, close } = useAgentAccountSetting(props, emit)
</script>
