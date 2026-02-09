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
          <span v-for="(betArea, betAreaIdx) in playerLog.betAreas" :key="`dogracingBetArea__${betArea.areaId}`">
            {{
              t(
                gameLog.betAreaDogs[betArea.areaId].length == 1
                  ? 'fmtTextDogracingWinBet'
                  : 'fmtTextDogracingQuinellaBet',
                [
                  ...gameLog.betAreaDogs[betArea.areaId],
                  numberToStr(playerLog.odds[betArea.areaId]),
                  numberToStr(betArea.betScore),
                ]
              )
            }}
            <template v-if="betAreaIdx !== playerLog.betAreas.length - 1">{{ t('symbolComma') }}</template>
          </span>
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textDrawingResult') }}:</div>
      <div v-for="(dogId, rankIdx) in gameLog.topTwoDogs" :key="`dogracingTopTwoDog__${dogId}`">
        {{ t('fmtTextDogracingDrawingResult', [rankIdx + 1, dogId]) }}
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
import { DogracingGameLog } from '@/base/common/gameLog/dogracing'
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

const gameLog = computed(() => new DogracingGameLog(props.playLog))
</script>
