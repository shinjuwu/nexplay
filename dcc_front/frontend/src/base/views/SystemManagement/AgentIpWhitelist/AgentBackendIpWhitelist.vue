<template>
  <AgentBaseIpWhitelist
    :ip-addresses="agentIpWhitelist.items"
    :visible="visible"
    :title="t('fmtTextSetBackendIpWhitelist', [props.agentName])"
    :mode="props.mode"
    :agent-id="props.agentId"
    @close="emit('close')"
    @update="updateAgentIpWhitelist"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useAgentBackendIpWhitelist } from '@/base/composable/systemManagement/agentIpWhitelist/useAgentBackendIpWhitelist'

import AgentBaseIpWhitelist from '@/base/views/SystemManagement/AgentIpWhitelist/AgentBaseIpWhitelist.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  mode: {
    type: String,
    default: 'view',
  },
  agentId: {
    type: Number,
    default: -1,
  },
  agentName: {
    type: String,
    default: '',
  },
})
const emit = defineEmits(['close', 'refreshIpList'])

const { t } = useI18n()
const { agentIpWhitelist, visible, updateAgentIpWhitelist } = useAgentBackendIpWhitelist(props, emit)
</script>
