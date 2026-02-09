<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('fmtTextBaseBet', [gameLog.baseBet]) }}</div>
      <div>{{ t('fmtTextTotalGameTax', [numberToStr(gameLog.totalTax)]) }}</div>
    </div>

    <div class="mt-4">
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextGoldenFlowerPlayerLog"
        tag="div"
        scope="global"
      >
        <template #seat>{{ playerLog.seat }}(seatId:{{ playerLog.seatId }})</template>
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        <template #handCards>
          &nbsp;
          <PokerCard
            v-for="card in playerLog.cards"
            :key="`playerCard_${playerLog.userName}_${card}`"
            class="mr-1 mb-1"
            :index="card"
          />
        </template>
        <template #cardType>{{ t(`goldenflowerCardType__${playerLog.cardType}`) }}</template>
        <template #status>{{ t(`goldenflowerStatus__${playerLog.status}`, [playerLog.lastOpRound]) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textPayoutResult') }}:</div>
      <PlayerResult
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerResult_${playerLog.userName}__${index}`"
        keypath="fmtTextGoldenflowerBonusPlayerResult"
        :player-log="playerLog"
        :user-name="props.userName"
      />
    </div>

    <table class="mt-4 border-l border-t text-center">
      <thead>
        <tr>
          <th colspan="5" class="border-r border-b px-3.5 py-1">{{ t('textGameProcessLog') }}</th>
        </tr>
        <tr>
          <th class="border-r border-b px-3.5 py-1">{{ t('textGameRound') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textSeat') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textPlayerAccount') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textProcessLog') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textBets') }}</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(roundHistories, roundIndex) in gameLog.history" :key="`goldenflower_round__${roundIndex}`">
          <tr
            v-for="(roundHistory, historyIndex) in roundHistories"
            :key="`goldenflower_round_history__${roundHistory.index}`"
          >
            <td v-if="historyIndex === 0" :rowspan="roundHistories.length" class="border-r border-b px-3.5 py-1">
              {{ t('fmtTextRound', [roundIndex + 1]) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': roundHistory.index % 2 === 0 }">
              {{ t('fmtTextTexasHistorySeat', [roundHistory.seat, roundHistory.seatId]) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': roundHistory.index % 2 === 0 }">
              <span :class="{ 'text-primary': props.userName === roundHistory.userName }">
                {{ roundHistory.userName }}
              </span>
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': roundHistory.index % 2 === 0 }">
              {{ t(`goldenflowerOp__${roundHistory.op}`) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': roundHistory.index % 2 === 0 }">
              {{ numberToStr(roundHistory.bet) }}
            </td>
          </tr>
        </template>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { GoldenflowerGameLog } from '@/base/common/gameLog/goldenflower'
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

const gameLog = computed(() => new GoldenflowerGameLog(props.playLog))
</script>
