<template>
  <div>
    <i18n-t keypath="fmtTextRcfishingGameProcess" tag="div" scope="global">
      <template #userName>
        <span :class="{ 'text-primary': props.targetUserName === props.playerLog.userName }">
          {{ props.playerLog.userName }}
        </span>
      </template>
    </i18n-t>

    <HappyfishingPlayerBulletInfo class="mt-2" :bullets="props.playerLog.bullets" />
    <HappyfishingPlayerNormalFishInfo
      v-if="props.playerLog.normalFishes.length > 0"
      class="mt-2"
      :normal-fishes="props.playerLog.normalFishes"
      :bullets="props.playerLog.bullets"
    />
    <HappyfishingPlayerSpecialFishInfo
      v-if="props.playerLog.specialFishes.length > 0"
      class="mt-2"
      :special-fishes="props.playerLog.specialFishes"
      :bullets="props.playerLog.bullets"
    />
    <HappyfishingPlayerBossFishInfo
      v-if="props.playerLog.bossFishes.length > 0"
      class="mt-2"
      :boss-fishes="props.playerLog.bossFishes"
      :bullets="props.playerLog.bullets"
    />
    <HappyfishingPlayerSkillInfo
      v-if="props.playerLog.skills.length > 0"
      class="mt-2"
      :skills="props.playerLog.skills"
    />
  </div>
</template>

<script setup>
import HappyfishingPlayerBulletInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerBulletInfo.vue'
import HappyfishingPlayerNormalFishInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerNormalFishInfo.vue'
import HappyfishingPlayerSpecialFishInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerSpecialFishInfo.vue'
import HappyfishingPlayerBossFishInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerBossFishInfo.vue'
import HappyfishingPlayerSkillInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerSkillInfo.vue'

const props = defineProps({
  playerLog: {
    type: Object,
    default: () => {},
  },
  targetUserName: {
    type: String,
    default: '',
  },
})
</script>
