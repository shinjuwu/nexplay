<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('textParticipants') }}:</div>
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextPlayerLog"
        tag="div"
        scope="global"
      >
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        <template #betDetail>
          <i18n-t
            v-for="(betArea, betAreaIdx) in playerLog.betAreas"
            :key="`colordiscBetArea__${betArea.areaId}`"
            keypath="fmtTextBetAreaResult"
            tag="span"
            scope="global"
          >
            <span>
              <template v-if="betArea.areaId < 2">{{ t(`colordiscBetArea__${betArea.areaId}`) }}</template>
              <template v-else>
                <span
                  v-for="(r, i) in gameLog.getBetAreaRedWhiteResult(betArea.areaId)"
                  :key="`colordiscBetArea__${betArea.areaId}__${r}__${i}`"
                  class="colordisc-container"
                >
                  <span class="colordisc" :class="{ 'colordisc-red': r === 0, 'colordisc-white': r === 1 }"></span>
                </span>
              </template>
            </span>
            <span>
              {{ numberToStr(betArea.betScore) }}
              <span v-if="betAreaIdx !== playerLog.betAreas.length - 1">
                {{ t('symbolComma') }}
              </span>
            </span>
          </i18n-t>
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textDealerResult') }}:</div>
      <i18n-t keypath="fmtTextColordiscResult" tag="div" scope="global" class="flex items-center">
        <span>
          <span v-for="(r, i) in gameLog.result" :key="`colordisc__${r}__${i}`" class="colordisc-container">
            <span class="colordisc" :class="{ 'colordisc-red': r === 0, 'colordisc-white': r === 1 }"></span>
          </span>
        </span>
        <span>
          {{ t(`colordiscResultOddEven__${gameLog.whiteCount % 2}`) }}
        </span>
      </i18n-t>
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
import { ColordiscGameLog } from '@/base/common/gameLog/colordisc'
import { numberToStr } from '@/base/utils/formatNumber'

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

const gameLog = computed(() => new ColordiscGameLog(props.playLog))
</script>
