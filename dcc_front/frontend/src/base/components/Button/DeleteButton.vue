<template>
  <ButtonTooltips :tips-text="props.tipsText || t('textDelete')">
    <template #content>
      <button type="button">
        <TrashIcon :class="{ 'h-4 w-4': props.size === 'md', 'h-5 w-5': props.size === 'lg' }" />
      </button>
    </template>
  </ButtonTooltips>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { TrashIcon } from '@heroicons/vue/24/outline'
import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'

const { t } = useI18n()
const props = defineProps({
  size: {
    type: String,
    default: 'lg',
    validator: (value) => ['sm', 'md', 'lg', 'xl', '2xl'].indexOf(value) >= 0,
  },
})
</script>
