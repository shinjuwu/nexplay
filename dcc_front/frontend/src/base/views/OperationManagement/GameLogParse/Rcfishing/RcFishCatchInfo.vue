<template>
  <div class="flex">
    <div
      class="flex min-w-[128px] items-center justify-center border-b border-l px-2"
      :class="{ 'border-t': props.hasTopBorder }"
    >
      <RcPic :name="fish.name" />
    </div>
    <div class="flex flex-col border-b border-l" :class="{ 'border-t': props.hasTopBorder }">
      <div class="bg-info flex flex-1 justify-center border-b px-2 py-1 text-white">BET</div>
      <div class="flex flex-1 justify-center border-b px-2 py-1">
        {{ t('textRcfishingCatchFishCount') }}
      </div>
      <div class="flex flex-1 justify-center border-b px-2 py-1">{{ t('textRcfishingOddsOrSkill') }}</div>
      <div class="flex flex-1 justify-center px-2 py-1">{{ t('textRcfishingTotalScore') }}</div>
    </div>
    <div
      v-for="(catchInfo, bulletIndex) in fish.catchInfo"
      :key="`catch_${fish.name}_${fish.odds}_${bulletIndex}`"
      class="flex min-w-[128px] flex-col border-b border-l"
      :class="{ 'border-t': props.hasTopBorder, 'border-r': bulletIndex === props.bullets.length - 1 }"
    >
      <div class="bg-info flex flex-1 justify-center border-b px-2 py-1 text-white">
        {{ props.bullets[bulletIndex].value }}
      </div>
      <div class="flex flex-1 justify-center border-b px-2 py-1">{{ catchInfo.count }}</div>
      <div class="flex flex-1 justify-center border-b px-2 py-1">
        {{ fish.reward ? t(`rcfishingSkillType_${fish.reward}`) : numberToStr(fish.odds) }}
      </div>
      <div :class="{ 'text-danger': catchInfo.totalGold > 0 }" class="flex flex-1 justify-center px-2 py-1">
        {{ numberToStr(catchInfo.totalGold) }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'
import RcPic from '@/base/views/OperationManagement/GameLogParse/Rcfishing/RcPic.vue'

const props = defineProps({
  bullets: {
    type: Array,
    default: () => [],
  },
  fish: {
    type: Object,
    default: () => {},
  },
  hasTopBorder: {
    type: Boolean,
    default: false,
  },
})

const { t } = useI18n()
</script>
