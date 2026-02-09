<template>
  <PageTableDialog :visible="props.visible" class="modal-lg" @close="close">
    <template #header>
      {{ t('textBackendAccountAddBackendAccount') }}
    </template>
    <template #default>
      <div class="grid grid-cols-2 gap-4">
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textAccount') }}</label>
          <input
            v-model="newAccountForm.account"
            class="form-input"
            type="text"
            :placeholder="t('placeHolderTextFormAccount')"
            maxlength="16"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textPassword') }}</label>
          <input
            v-model="newAccountForm.password"
            class="form-input"
            autocomplete="new-password"
            type="password"
            :placeholder="t('placeHolderTextFormPassword')"
            maxlength="16"
          />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textNickName') }}</label>
          <input v-model="newAccountForm.nickName" class="form-input" type="text" maxlength="16" />
        </div>
        <div class="col-span-2 md:col-span-1">
          <label class="form-label required-mark">{{ t('textGroupRole') }}</label>
          <FormAgentPermissionListDropdown v-model="newAccountForm.role" class="form-input" />
        </div>
        <div class="col-span-2">
          <label class="form-label">{{ t('textRemark') }}</label>
          <textarea v-model="newAccountForm.remark" class="form-input" type="text" maxlength="100" />
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
          :parent-data="newAccountForm"
          :button-click="() => createAccount()"
        >
          {{ t('textIncrease') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useBackendAccountAddSubAccounts } from '@/base/composable/systemManagement/backendAccount/useBackendAccountAddSubAccounts'
import FormAgentPermissionListDropdown from '@/base/components/Form/Dropdown/FormAgentPermissionListDropdown.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'
import { useI18n } from 'vue-i18n'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['close', 'searchRecords'])

const { t } = useI18n()
const { newAccountForm, createAccount, close } = useBackendAccountAddSubAccounts(emit)
</script>
