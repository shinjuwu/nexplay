<template>
  <div>
    <div>{{ t('textRcfishingNormalFish') }}</div>
    <HappyfishingCatchFishingInfo
      v-for="fish in props.normalFishes"
      :key="`normal_${fish.name}_${fish.odds}`"
      :fish="fish"
      :bullets="props.bullets"
    />
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import HappyfishingCatchFishingInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingCatchFishingInfo.vue'

const props = defineProps({
  normalFishes: {
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
