<template>
  <AgentBaseIpWhitelist
    :ip-addresses="agentIpWhitelist.items"
    :visible="visible"
    :title="t('fmtTextSetApiIpWhitelist', [props.agentName])"
    :mode="props.mode"
    :agent-id="props.agentId"
    @close="emit('close')"
    @update="updateAgentApiIpWhitelist"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useAgentApiIpWhitelist } from '@/base/composable/systemManagement/agentIpWhitelist/useAgentApiIpWhitelist'

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
const { agentIpWhitelist, visible, updateAgentApiIpWhitelist } = useAgentApiIpWhitelist(props, emit)
</script>
