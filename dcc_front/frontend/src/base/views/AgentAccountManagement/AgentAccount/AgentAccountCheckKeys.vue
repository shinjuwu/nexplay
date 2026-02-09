<template>
  <PageTableDialog :visible="show" @close="close">
    <template #header>
      {{ t('textAgentSettingCheckKeys') }}
    </template>
    <template #default>
      <div class="grid grid-cols-2 gap-4">
        <div class="col-span-2">{{ t('fmtTextAgentSettingAgentId', [agent.id]) }}</div>
        <div>
          <label class="form-label">Aeskey</label>
          <input v-model="agent.aeskey" class="form-input" type="text" disabled />
        </div>
        <div>
          <label class="form-label">Md5key</label>
          <input v-model="agent.md5key" class="form-input" type="text" disabled />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto">
        <button type="button" class="btn btn-light mr-2" @click="close">{{ t('textCancel') }}</button>
        <button type="button" class="btn btn-primary" @click="close">{{ t('textConfirm') }}</button>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useAgentAccountCheckKeys } from '@/base/composable/agentAccountManagement/agentAccount/useAgentAccountCheckKeys'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import { useI18n } from 'vue-i18n'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  selectRecord: {
    type: Object,
    default: () => {},
  },
})
const emit = defineEmits(['close'])

const { t } = useI18n()
const { show, agent, close } = useAgentAccountCheckKeys(props, emit)
</script>
