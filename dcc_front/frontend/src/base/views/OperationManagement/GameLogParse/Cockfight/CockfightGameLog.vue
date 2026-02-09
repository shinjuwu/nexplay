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
          <span v-for="(betArea, betAreaIdx) in playerLog.betAreas" :key="`cockfightBetArea__${betArea.areaId}`">
            {{ t('fmtTextBetAreaResult', [t(`cockfightBetArea__${betArea.areaId}`), numberToStr(betArea.betScore)]) }}
            <template v-if="betAreaIdx !== playerLog.betAreas.length - 1">{{ t('symbolComma') }}</template>
          </span>
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('fmtTextCockfightRedCock', [t(`cockfightRedCock__${gameLog.fightGroup[0]}`)]) }}</div>
      <div>{{ t('fmtTextCockfightBlueCock', [t(`cockfightBlueCock__${gameLog.fightGroup[1]}`)]) }}</div>
    </div>

    <div class="mt-4">
      <div v-for="(odds, idx) in gameLog.odds" :key="`cockfightOdds__${idx}__${odds}`">
        {{ t('fmtTextCockfightOdds', [t(`cockfightBetArea__${idx}`), numberToStr(odds)]) }}
      </div>
    </div>

    <div class="mt-4">
      <div>{{ t('textDrawingResult') }}:</div>
      <div>{{ t(`cockfightBetArea__${gameLog.result}`) }}</div>
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
import { CockfightGameLog } from '@/base/common/gameLog/cockfight'
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

const gameLog = computed(() => new CockfightGameLog(props.playLog))
</script>
