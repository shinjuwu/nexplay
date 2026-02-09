<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('textSeatPlayers') }}:</div>
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextBlackjackPlayerLog"
        tag="p"
        scope="global"
      >
        <template #seat>{{ playerLog.seat }}(seatId:{{ playerLog.seatId }})</template>
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        <template #betDetail>{{ numberToStr(playerLog.yaScore) }}</template>
      </i18n-t>
    </div>

    <div
      v-for="(playerLog, index) in gameLog.playerLogs"
      :key="`playerLog_${playerLog.userName}__${index}`"
      class="mt-4"
    >
      <i18n-t keypath="fmtTextRcfishingGameProcess" tag="div" scope="global">
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
      </i18n-t>

      <div class="mt-2">
        <div>{{ t('textRcfishingBulletInfo') }}</div>
        <div class="flex">
          <div class="flex min-w-[128px] px-2"></div>
          <div class="flex flex-col border-y border-l">
            <div class="bg-info flex flex-1 items-center justify-center border-t px-2 py-1 text-white">BET</div>
            <div class="flex flex-1 items-center justify-center border-t px-2 py-1">
              {{ t('textRcfishingTotalBulletCount') }}
            </div>
            <div class="flex flex-1 items-center justify-center border-t px-2 py-1">
              {{ t('textRcfishingTotalBulletGold') }}
            </div>
          </div>
          <div
            v-for="(bullet, bulletIdx) in playerLog.bullets"
            :key="`bullet_${playerLog.userName}_${bulletIdx}`"
            class="flex min-w-[128px] flex-col border-y border-l"
            :class="{ 'border-r': bulletIdx === playerLog.bullets.length - 1 }"
          >
            <div class="bg-info flex flex-1 items-center justify-center border-t px-2 py-1 text-white">
              {{ bullet.value }}
            </div>
            <div class="flex flex-1 items-center justify-center border-t px-2 py-1">
              {{ bullet.count.toLocaleString() }}
            </div>
            <div class="flex flex-1 items-center justify-center border-t px-2 py-1">
              {{ numberToStr(bullet.totalBulletGold) }}
            </div>
          </div>
        </div>
      </div>

      <div v-if="playerLog.normalFishes.length > 0" class="mt-2">
        <div>{{ t('textRcfishingNormalFish') }}</div>
        <RcFishCatchInfo
          v-for="(fish, fishIndex) in playerLog.normalFishes"
          :key="`${playerLog.userName}_normal_${fish.name}_${fish.odds}`"
          :fish="fish"
          :bullets="playerLog.bullets"
          :has-top-border="fishIndex === 0"
        />
      </div>

      <div v-if="playerLog.specialFishes.length > 0" class="mt-2">
        <div>{{ t('textRcfishingSpecialFish') }}</div>
        <RcFishCatchInfo
          v-for="(fish, fishIndex) in playerLog.specialFishes"
          :key="`${playerLog.userName}_special_${fish.name}_${fish.odds}_${fish.reward || ''}`"
          :fish="fish"
          :bullets="playerLog.bullets"
          :has-top-border="fishIndex === 0"
        />
      </div>

      <div v-if="playerLog.bossFishes.length > 0" class="mt-2">
        <div>{{ t('textRcfishingBossFish') }}</div>
        <RcFishCatchInfo
          v-for="(fish, fishIndex) in playerLog.bossFishes"
          :key="`${playerLog.userName}_boss_${fish.name}_${fish.odds}`"
          :fish="fish"
          :bullets="playerLog.bullets"
          :has-top-border="fishIndex === 0"
        />
      </div>

      <div v-if="playerLog.skills.length > 0" class="mt-2">
        <div>{{ t('textRcfishingSkillExplodeFishInfo') }}</div>
        <div v-for="skill in playerLog.skills" :key="`${playerLog.userName}_${skill.name}_${skill.bulletGold}`">
          <template
            v-for="groupIndex in Math.ceil(skill.fishes.length / 4)"
            :key="`${playerLog.userName}_${skill.name}_${skill.bulletGold}_${groupIndex}`"
          >
            <div class="flex">
              <div
                class="flex min-w-[128px] items-center justify-center px-2"
                :class="{ 'border-l border-y': groupIndex === 1 }"
              >
                <RcPic v-if="groupIndex === 1" :name="skill.name" />
              </div>
              <template
                v-for="(fish, fishIndex) in skill.fishes.slice((groupIndex - 1) * 4, groupIndex * 4)"
                :key="`${playerLog.userName}_${skill.name}_${skill.bulletGold}_${groupIndex}_${fish.name}_${fishIndex}`"
              >
                <div
                  class="flex min-w-[128px] items-center justify-center border-b border-l px-2"
                  :class="{ 'border-t': groupIndex === 1 }"
                >
                  <RcPic :name="fish.name" />
                </div>
                <div class="flex flex-col border-b border-l" :class="{ 'border-t': groupIndex === 1 }">
                  <div class="bg-info flex flex-1 justify-center border-b px-2 py-1 text-white">BET</div>
                  <div class="flex flex-1 justify-center border-b px-2 py-1">
                    {{ t('textRcfishingCatchFishCount') }}
                  </div>
                  <div class="flex flex-1 justify-center border-b px-2 py-1">{{ t('textRcfishingOddsOrSkill') }}</div>
                  <div class="flex flex-1 justify-center px-2 py-1">{{ t('textRcfishingTotalScore') }}</div>
                </div>
                <div
                  class="flex min-w-[120px] flex-col border-b border-l"
                  :class="{
                    'border-t': groupIndex === 1,
                    'border-r': fishIndex === 3 || (groupIndex - 1) * 4 + fishIndex === skill.fishes.length - 1,
                  }"
                >
                  <div class="bg-info flex flex-1 justify-center border-b px-2 py-1 text-white">
                    {{ skill.bulletGold }}
                  </div>
                  <div class="flex flex-1 justify-center border-b px-2 py-1">{{ fish.count }}</div>
                  <div class="flex flex-1 justify-center border-b px-2 py-1">
                    {{ fish.reward ? t(`rcfishingSkillType_${fish.reward}`) : numberToStr(fish.odds) }}
                  </div>
                  <div class="text-danger flex flex-1 justify-center px-2 py-1">{{ numberToStr(fish.totalGold) }}</div>
                </div>
              </template>
            </div>
          </template>
        </div>
      </div>
    </div>

    <div class="mt-4">
      <div>{{ t('textPayoutResult') }}:</div>
      <PlayerResult
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerResult_${playerLog.userName}__${index}`"
        :player-log="playerLog"
        :user-name="props.userName"
      />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { RcfishingGameLog } from '@/base/common/gameLog/rcfishing'
import { numberToStr } from '@/base/utils/formatNumber'

import RcPic from '@/base/views/OperationManagement/GameLogParse/Rcfishing/RcPic.vue'
import RcFishCatchInfo from '@/base/views/OperationManagement/GameLogParse/Rcfishing/RcFishCatchInfo.vue'
import PlayerResult from '@/base/views/OperationManagement/GameLogParse/PlayerResult.vue'

const { t } = useI18n()

const props = defineProps({
  playLog: {
    type: Object,
    default: () => {},
  },
  userName: {
    type: String,
    default: '',
  },
})

const gameLog = computed(() => new RcfishingGameLog(props.playLog))
</script>
