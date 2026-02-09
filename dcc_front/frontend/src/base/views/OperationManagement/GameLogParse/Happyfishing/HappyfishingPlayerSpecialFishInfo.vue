<template>
  <div>
    <div>{{ t('textRcfishingSpecialFish') }}</div>
    <HappyfishingCatchFishingInfo
      v-for="fish in props.specialFishes"
      :key="`special_${fish.name}_${fish.odds}_${fish.reward || ''}`"
      :fish="fish"
      :bullets="props.bullets"
    />
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import HappyfishingCatchFishingInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingCatchFishingInfo.vue'

const props = defineProps({
  specialFishes: {
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
