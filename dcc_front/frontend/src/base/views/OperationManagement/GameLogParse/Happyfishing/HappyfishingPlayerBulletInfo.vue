<template>
  <div>
    <div>{{ t('textRcfishingBulletInfo') }}</div>
    <div
      class="grid text-center"
      :style="{
        'grid-template-columns': 'minmax(128px, min-content) min-content repeat(10, minmax(128px, min-content))',
      }"
    >
      <div class="bg-info col-start-2 row-start-1 border px-2 py-1 text-white">BET</div>
      <div class="col-start-2 row-start-2 border px-2 py-1">{{ t('textRcfishingTotalBulletCount') }}</div>
      <div class="col-start-2 row-start-3 border px-2 py-1">{{ t('textRcfishingTotalBulletGold') }}</div>

      <template v-for="(bullet, bulletIdx) in props.bullets" :key="`bullet_${bulletIdx}`">
        <div class="bg-info row-start-1 border px-2 py-1 text-white" :style="{ 'grid-column-start': 3 + bulletIdx }">
          {{ bullet.value }}
        </div>
        <div class="row-start-2 border px-2 py-1" :style="{ 'grid-column-start': 3 + bulletIdx }">
          {{ bullet.count.toLocaleString() }}
        </div>
        <div class="row-start-3 border px-2 py-1" :style="{ 'grid-column-start': 3 + bulletIdx }">
          {{ numberToStr(bullet.totalBulletGold) }}
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'

const props = defineProps({
  bullets: {
    type: Array,
    default: () => [],
  },
})
const { t } = useI18n()
</script>
