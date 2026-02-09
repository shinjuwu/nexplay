<template>
  <div
    class="grid text-center"
    :style="{
      'grid-template-columns': 'minmax(128px, min-content) min-content repeat(10, minmax(128px, min-content))',
    }"
  >
    <div class="row-span-5 flex items-center justify-center border px-2 py-1">
      <HappyfishingPic :name="fish.name" />
    </div>

    <div class="bg-info col-start-2 row-start-1 border px-2 py-1 text-white">BET</div>
    <div class="col-start-2 row-start-2 border px-2 py-1">{{ t('textRcfishingCatchFishCount') }}</div>
    <div class="col-start-2 row-start-3 border px-2 py-1">{{ t('textRcfishingOddsOrSkill') }}</div>
    <div class="col-start-2 row-start-4 border px-2 py-1">{{ t('textRcfishingTotalScore') }}</div>

    <template
      v-for="(catchInfo, bulletIndex) in props.fish.catchInfo"
      :key="`catch_${props.fish.name}_${props.fish.odds}_${bulletIndex}`"
    >
      <div class="bg-info row-start-1 border px-2 py-1 text-white" :style="{ 'grid-column-start': bulletIndex + 3 }">
        {{ props.bullets[bulletIndex].value }}
      </div>
      <div class="row-start-2 border px-2 py-1" :style="{ 'grid-column-start': bulletIndex + 3 }">
        {{ catchInfo.count }}
      </div>
      <div class="row-start-3 border px-2 py-1" :style="{ 'grid-column-start': bulletIndex + 3 }">
        {{ fish.reward && fish.odds === 0 ? t(`happyfishingSkillType_${fish.reward}`) : numberToStr(fish.odds) }}
      </div>
      <div
        class="row-start-4 border px-2 py-1"
        :style="{ 'grid-column-start': bulletIndex + 3 }"
        :class="{ 'text-danger': catchInfo.totalGold > 0 }"
      >
        {{ numberToStr(catchInfo.totalGold) }}
      </div>
    </template>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'
import HappyfishingPic from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPic.vue'

const props = defineProps({
  bullets: {
    type: Array,
    default: () => [],
  },
  fish: {
    type: Object,
    default: () => {},
  },
})

const { t } = useI18n()
</script>
