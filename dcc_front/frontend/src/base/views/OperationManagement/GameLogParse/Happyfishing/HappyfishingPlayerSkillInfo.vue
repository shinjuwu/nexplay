<template>
  <div>
    <div>{{ t('textRcfishingSkillExplodeFishInfo') }}</div>
    <div v-for="skill in props.skills" :key="`${skill.name}_${skill.bulletGold}`">
      <div
        class="grid text-center"
        :style="{ 'grid-template-columns': 'minmax(128px, min-content) repeat(4, min-content)' }"
      >
        <div class="flex items-center justify-center border px-2 py-1">
          <HappyfishingPic :name="skill.name" />
        </div>

        <div
          v-for="(fish, fishIndex) in skill.fishes"
          :key="`${skill.name}_${skill.bulletGold}_${fish.name}_${fishIndex}`"
          class="grid"
          :style="{
            'grid-template-columns': 'minmax(128px, min-content) min-content minmax(120px, min-content)',
            'grid-column-start': (fishIndex % 4) + 2,
          }"
        >
          <div class="row-start-1 row-end-5 flex items-center justify-center border px-2 py-1">
            <HappyfishingPic :name="fish.name" />
          </div>
          <div class="bg-info col-start-2 row-start-1 border px-2 py-1 text-white">BET</div>
          <div class="col-start-2 row-start-2 border px-2 py-1">{{ t('textRcfishingCatchFishCount') }}</div>
          <div class="col-start-2 row-start-3 border px-2 py-1">{{ t('textRcfishingOddsOrSkill') }}</div>
          <div class="col-start-2 row-start-4 border px-2 py-1">{{ t('textRcfishingTotalScore') }}</div>

          <div class="bg-info col-start-3 row-start-1 border px-2 py-1 text-white">{{ skill.bulletGold }}</div>
          <div class="col-start-3 row-start-2 border px-2 py-1">{{ fish.count }}</div>
          <div class="col-start-3 row-start-3 border px-2 py-1">
            {{ fish.reward ? t(`happyfishingSkillType_${fish.reward}`) : numberToStr(fish.odds) }}
          </div>
          <div class="text-danger col-start-3 row-start-4 border px-2 py-1">{{ numberToStr(fish.totalGold) }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { numberToStr } from '@/base/utils/formatNumber'
import HappyfishingPic from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPic.vue'

const props = defineProps({
  skills: {
    type: Array,
    default: () => [],
  },
})
const { t } = useI18n()
</script>
