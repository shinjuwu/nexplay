div
<template>
  <PageTableDialog :visible="show" class="modal-lg" @close="close">
    <template #header>
      {{ t('textBackendAccountBackendAccountSetting') }}
    </template>
    <template #default>
      <div class="grid grid-cols-4 gap-4">
        <div class="col-span-4">{{ t('fmtTextUserName', [account.userName]) }}</div>
        <div class="col-span-1">
          <label class="form-label">{{ t('textState') }}</label>
          <select v-model="account.state" class="form-input" :disabled="!props.isEditEnabled">
            <option
              v-for="option in stateOptions"
              :key="`status__${option}`"
              :value="option"
              :label="t(`status__${option}`)"
            ></option>
          </select>
        </div>
        <div class="col-span-2">
          <label class="form-label">{{ t('textGroupRole') }}</label>
          <FormAgentPermissionListDropdown v-if="props.isEditEnabled" v-model="account.role" class="form-input" />
          <input v-else class="form-input" type="text" disabled :value="account.roleName" />
        </div>
        <div class="col-span-4">
          <label class="form-label">{{ t('textRemark') }}</label>
          <input
            v-model="account.remark"
            class="form-input"
            type="text"
            maxlength="100"
            :disabled="!props.isEditEnabled"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close">{{ t('textCancel') }}</button>
        <LoadingButton
          v-show="props.isEditEnabled"
          class="btn btn-primary"
          :is-get-data="show"
          :parent-data="account"
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
import { useBackendAccountSetting } from '@/base/composable/systemManagement/backendAccount/useBackendAccountSetting'
import FormAgentPermissionListDropdown from '@/base/components/Form/Dropdown/FormAgentPermissionListDropdown.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
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
const emit = defineEmits(['close', 'searchRecords'])

const { t } = useI18n()
const { account, show, stateOptions, close, submit } = useBackendAccountSetting(props, emit)
</script>
