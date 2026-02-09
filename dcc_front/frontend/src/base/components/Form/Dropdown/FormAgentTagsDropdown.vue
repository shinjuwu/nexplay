<template>
  <FormDropdown
    v-if="isInit"
    :fmt-item-text="(tag) => tag.name"
    :fmt-item-key="(tag) => `tag__${tag.index}`"
    :items="tags"
    :model-value="props.modelValue"
    @update:model-value="(newValue) => updateSelectTag(newValue)"
  />
</template>

<script setup>
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import { computed, onMounted, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/base/store/userStore'
import { useDropdownListStore } from '@/base/store/dropdownStore'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const dStore = useDropdownListStore()
const { dropdownList } = storeToRefs(dStore)
const isInit = ref(false)
const tags = computed(() => {
  const allTags = {
    index: 101,
    name: t('riskType__All'),
  }
  return [allTags, ...dropdownList.value.tags.filter((tag) => tag.name !== '')]
})

let currentTag = null
function updateSelectTag(selectTag) {
  currentTag = tags.value.find((tag) => tag.index === selectTag.index) || tags.value[0]
  emit('update:modelValue', currentTag)
}

watch(tags, (newTagsList) => {
  if (newTagsList && newTagsList.length) {
    let tempTag = null
    if (currentTag !== null) {
      tempTag = newTagsList.find((tag) => tag.index === currentTag.index)
    }

    if (!tempTag) {
      tempTag = newTagsList[0]
    }

    updateSelectTag(tempTag)
  }
})

onMounted(async () => {
  await dStore.getAgentTagsList()
  isInit.value = true
  emit('update:modelValue', tags.value[0])
})

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => {},
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
