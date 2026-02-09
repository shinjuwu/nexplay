<template>
  <select v-model="selectModel">
    <option v-for="agentPermission in agentPermissions" :key="agentPermission.id" :value="agentPermission.id">
      {{ agentPermission.name }}
    </option>
  </select>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/base/store/userStore'
import { useDropdownListStore } from '@/base/store/dropdownStore'

const dStore = useDropdownListStore()
const { dropdownList } = storeToRefs(useDropdownListStore())
const isInit = ref(false)

const selectModel = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  },
})

const agentPermissions = computed(() => {
  return dropdownList.value.agentPermissions[props.accountType] || []
})

watch(agentPermissions, (newValue) => {
  if (newValue && newValue.length) {
    emit('update:modelValue', newValue[0].id)
  }
})

onMounted(async () => {
  await dStore.getAgentPermissionList(props.accountType)
  isInit.value = true
})

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  accountType: {
    type: Number,
    default: () => {
      const { user } = storeToRefs(useUserStore())
      return user.value.accountType
    },
  },
})
const emit = defineEmits(['update:modelValue'])
</script>
