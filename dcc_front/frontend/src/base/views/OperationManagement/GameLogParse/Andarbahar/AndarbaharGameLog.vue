<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div class="flex items-center">Joker:&nbsp;<PokerCard :index="gameLog.joker" /></div>
    </div>

    <div class="mt-4">
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
          <span v-for="(betArea, betAreaIdx) in playerLog.betAreas" :key="`andarbaharBetArea__${betArea.areaId}`">
            {{ t('fmtTextBetAreaResult', [t(`andarbaharBetArea__${betArea.areaId}`), numberToStr(betArea.betScore)]) }}
            <template v-if="betAreaIdx !== playerLog.betAreas.length - 1">{{ t('symbolComma') }}</template>
          </span>
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textDealerResult') }}:</div>
      <div class="flex items-center">
        {{ t('textAndar') }}:&nbsp;
        <PokerCard
          v-for="(card, index) in gameLog.andarAreaCards"
          :key="`poker__${card}_card__${index}_andarAreaCards`"
          class="mb-1 mr-1"
          :index="card"
        />
      </div>
      <div class="flex items-center">
        {{ t('textBahar') }}:&nbsp;
        <PokerCard
          v-for="(card, index) in gameLog.baharAreaCards"
          :key="`poker__${card}_card__${index}_baharAreaCards`"
          class="mr-1"
          :index="card"
        />
      </div>
      <div v-if="gameLog.openCardCount">
        {{ t('fmtTextOpenCardCount', [gameLog.openCardCount]) }}
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
import { AndarbaharGameLog } from '@/base/common/gameLog/andarbahar'
import { numberToStr } from '@/base/utils/formatNumber'

import PokerCard from '@/base/components/PokerCard.vue'
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

const gameLog = computed(() => new AndarbaharGameLog(props.playLog))
</script>
