<template>
  <FormDropdown
    v-if="isInit"
    :fmt-item-text="(state) => (state === constant.GameState.All ? t('textStateAll') : t(`state__${state}`))"
    :fmt-item-key="(state) => `state__${state}`"
    :items="states.items"
    :use-font-awsome="false"
    :model-value="props.modelValue"
    @update:model-value="(newValue) => emit('update:modelValue', newValue)"
  />
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/base/store/userStore'
import constant from '@/base/common/constant'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'

const { t } = useI18n()

const uStore = useUserStore()
const { user } = storeToRefs(uStore)

const states = {
  items:
    user.value.accountType === constant.AccountType.Admin
      ? [constant.GameState.All, constant.GameState.Online, constant.GameState.Maintain, constant.GameState.Offline]
      : [constant.GameState.All, constant.GameState.Online, constant.GameState.Maintain],
}
const isInit = ref(false)

onMounted(() => {
  emit('update:modelValue', states.items[0])
  isInit.value = true
})

const props = defineProps({
  modelValue: {
    type: Number,
    default: constant.GameState.All,
  },
})
const emit = defineEmits(['update:modelValue'])
</script>
