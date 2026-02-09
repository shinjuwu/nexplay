<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('fmtTextBaccaratRound', [gameLog.round]) }}</div>
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
          <span v-for="(betArea, betAreaIdx) in playerLog.betAreas" :key="`baccaratBetArea__${betArea.areaId}`">
            {{ t('fmtTextBetAreaResult', [t(`baccaratBetArea__${betArea.areaId}`), numberToStr(betArea.betScore)]) }}
            <template v-if="betAreaIdx !== playerLog.betAreas.length - 1">{{ t('symbolComma') }}</template>
          </span>
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textDealerResult') }}:</div>
      <div>
        {{ t('textBaccaratBanker') }}&nbsp;
        <PokerCard
          v-for="(card, index) in gameLog.banker.cards"
          :key="`poker__${card}_card__${index}_banker`"
          class="mb-1 mr-1"
          :index="card"
        />
        {{ t('fmtTextTotalPoint', [gameLog.banker.point]) }}
      </div>
      <div>
        {{ t('textBaccaratPlayer') }}&nbsp;
        <PokerCard
          v-for="(card, index) in gameLog.player.cards"
          :key="`poker__${card}_card__${index}_player`"
          class="me-1"
          :index="card"
        />
        {{ t('fmtTextTotalPoint', [gameLog.player.point]) }}
      </div>
      <div>
        {{
          gameLog.winner === 0
            ? t('textBaccaratTie')
            : gameLog.winner === 1
            ? t('textBaccaratBankerWin')
            : t('textBaccaratPlayerWin')
        }}
      </div>
      <div v-if="gameLog.winSmallThanOther > 0 || gameLog.winBiggerThanOther > 0">
        {{ gameLog.winSmallThanOther > 0 ? t('textBaccaratWinSmallThanOther') : t('textBaccaratWinBiggerThanOther') }}
      </div>
      <div v-if="gameLog.bankerPair > 0 || gameLog.playerPair > 0">
        <template v-if="gameLog.bankerPair > 0 && gameLog.playerPair > 0">
          {{ t('textBaccaratBankerPair') }}
          {{ t('symbolComma') }}
          {{ t('textBaccaratPlayerPair') }}
        </template>
        <template v-else-if="gameLog.bankerPair > 0">{{ t('textBaccaratBankerPair') }}</template>
        <template v-else>{{ t('textBaccaratPlayerPair') }}</template>
      </div>
      <div v-if="gameLog.big > 0 || gameLog.small > 0">
        <template v-if="gameLog.big > 0">{{ t('textBaccaratBig') }}</template>
        <template v-else-if="gameLog.small > 0">{{ t('textBaccaratSmall') }}<br /></template>
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
import { BaccaratGameLog } from '@/base/common/gameLog/baccarat'
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

const gameLog = computed(() => new BaccaratGameLog(props.playLog))
</script>
