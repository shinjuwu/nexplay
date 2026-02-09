<template>
  <div>
    <div>{{ t('textRcfishingBossFish') }}</div>
    <HappyfishingCatchFishingInfo
      v-for="fish in props.bossFishes"
      :key="`boss_${fish.name}_${fish.odds}`"
      :fish="fish"
      :bullets="props.bullets"
    />
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import HappyfishingCatchFishingInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingCatchFishingInfo.vue'

const props = defineProps({
  bossFishes: {
    type: Array,
    default: () => [],
  },
  bullets: {
    type: Array,
    default: () => [],
  },
})
const { t } = useI18n()
</script>
