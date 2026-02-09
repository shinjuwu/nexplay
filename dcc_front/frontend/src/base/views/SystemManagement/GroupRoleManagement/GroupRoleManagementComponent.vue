<template>
  <PageTableDialog :visible="visible" size="xl" @close="close()">
    <template #header>{{ contentText.title }}</template>
    <template #default>
      <div class="grid h-[50vh] grid-cols-2 grid-rows-3 gap-4">
        <div class="grid-cols-1">
          <div>
            <label class="form-label required-mark">{{ t('textGroupRole') }}</label>
            <input
              v-model="agentPermission.name"
              class="form-input"
              type="text"
              maxlength="16"
              :placeholder="t('placeHolderTextGroupRole')"
              :disabled="disabled"
            />
          </div>
          <div>
            <label class="form-label required-mark">{{ t('textGroupRoleLevel') }}</label>
            <select v-model="agentPermission.accountType" class="form-input" :disabled="disabled">
              <option
                v-for="accountType in accountTypes"
                :key="accountType.value"
                :value="accountType.value"
                :label="accountType.label"
              ></option>
            </select>
          </div>
          <div>
            <label class="form-label">{{ t('textRemark') }}</label>
            <textarea
              v-model="agentPermission.info"
              class="form-input"
              rows="8"
              maxlength="100"
              style="resize: none"
              :disabled="disabled"
            ></textarea>
          </div>
        </div>
        <div class="row-span-3 grid-cols-1">
          <label class="form-label required-mark">{{ t('textPermission') }}</label>
          <div class="h-[90%] border bg-white p-2" :style="{ opacity: disabled ? 0.6 : 1 }">
            <TreeMenuComponent
              class="h-full overflow-y-scroll"
              :items="roleMenu.items"
              :disabled="disabled"
              :model-value="agentPermission.permissions"
              @update:model-value="(newValue) => updatePermissions(newValue)"
            />
          </div>
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close()">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          v-show="!disabled"
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="agentPermission"
          :button-click="() => (props.mode === 'create' ? createAgentPermission() : updateAgentPermission())"
        >
          {{ contentText.submit }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useGroupRoleManagementComponent } from '@/base/composable/systemManagement/groupRoleManagement/useGroupRoleManagementComponent'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import TreeMenuComponent from '@/base/components/TreeMenu.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  agentPermission: {
    type: Object,
    default: () => {},
  },
  mode: {
    type: String,
    default: '',
  },
})
const emit = defineEmits(['close', 'refreshAgentPermissionList'])

const {
  t,
  visible,
  agentPermission,
  roleMenu,
  accountTypes,
  disabled,
  contentText,
  createAgentPermission,
  updateAgentPermission,
  updatePermissions,
  close,
} = useGroupRoleManagementComponent(props, emit)
</script>
