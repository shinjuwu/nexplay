<template>
  <FormDropdown
    v-if="isInit"
    :fmt-item-text="(agent) => agent.name"
    :fmt-item-key="(agent) => `agent__${agent.id}`"
    :items="agents"
    :use-font-awsome="false"
    :use-filter="true"
    :filter-rule="(agent, filterValue) => agent.name.includes(filterValue)"
    :model-value="props.modelValue"
    @update:model-value="(agent) => updateSelectAgent(agent)"
  />
</template>

<script setup>
import { onMounted, ref, computed, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import constant from '@/base/common/constant'
import { useUserStore } from '@/base/store/userStore'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import { useDropdownListStore } from '@/base/store/dropdownStore'

const { t } = useI18n()

const dStore = useDropdownListStore()
const { dropdownList } = storeToRefs(useDropdownListStore())

const isInit = ref(false)
const agents = computed(() => {
  const { user } = storeToRefs(useUserStore())
  const agentList = dropdownList.value.agents

  let agents = []
  if (props.includeAll) {
    agents.push({
      id: constant.Agent.All,
      name: t('textAgentAll'),
      levelCode: '',
    })
  }

  let sortedAgents = agents.concat(
    agentList.map((a) => {
      return {
        id: a.id,
        name: a.name,
        levelCode: a.level_code,
        cooperation: a.cooperation,
      }
    })
  )
  sortedAgents.sort((a, b) => a.levelCode.localeCompare(b.levelCode))

  if (!props.includeSelf) {
    sortedAgents = sortedAgents.filter((a) => a.id !== user.value.agentId)
  }

  if (!props.includeGrandson) {
    sortedAgents = sortedAgents.filter((a) => a.levelCode.length <= (user.value.accountType + 1) * 4)
  }

  if (props.agentCooperation !== 0) {
    sortedAgents = sortedAgents.filter((a) => a.id === constant.Agent.All || a.cooperation === props.agentCooperation)
  }
  agents = sortedAgents
  return agents
})

let currentAgent = null
function updateSelectAgent(selectAgent) {
  currentAgent = agents.value.find((agent) => agent.id === selectAgent.id)
  emit('update:modelValue', selectAgent || agents.value[0])
}

watch(
  () => [props.agentId, agents],
  ([newAgentId, newAgents]) => {
    if (newAgents.value && newAgents.value.length) {
      if (newAgentId !== -1) {
        const propsAgent = agents.value.find((agent) => agent.id === props.agentId)
        emit('update:modelValue', propsAgent)
        return
      }

      let tempAgent = null
      if (currentAgent !== null) {
        tempAgent = newAgents.value.find((agent) => agent.id === currentAgent.id)
      }

      if (!tempAgent) {
        tempAgent = newAgents.value[0]
      }

      updateSelectAgent(tempAgent)
    }
  },
  { deep: true }
)

onMounted(async () => {
  await dStore.getAgentList()
  isInit.value = true
})

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => {},
  },
  includeAll: {
    type: Boolean,
    default: true,
  },
  includeSelf: {
    type: Boolean,
    default: true,
  },
  includeGrandson: {
    type: Boolean,
    default: true,
  },
  agentId: {
    type: Number,
    default: -1,
  },
  agentCooperation: {
    type: Number,
    default: 0,
  },
})
const emit = defineEmits(['update:modelValue'])
</script>
