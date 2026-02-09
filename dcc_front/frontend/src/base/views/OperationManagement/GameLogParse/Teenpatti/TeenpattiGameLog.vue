<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>
        {{ t('fmtTextBaseBet', [numberToStr(gameLog.baseBet)]) }}
      </div>
      <div>
        {{ t('fmtTextTotalGameTax', [numberToStr(gameLog.totalTax)]) }}
      </div>
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
            :card-suit-names="cardSuitNames"
          />
        </template>
        <template #cardType>{{ t(`teenpattiCardType__${playerLog.cardType}`) }}</template>
        <template #status>{{ t(`teenpattiStatusWithRound__${playerLog.status}`, [playerLog.lastOpRound]) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textPayoutResult') }}:</div>
      <PlayerResult
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerResult_${playerLog.userName}__${index}`"
        :player-log="playerLog"
        :user-name="props.userName"
        keypath="fmtTextTexasPlayerResult"
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
        <template v-for="(roundInfo, roundIndex) in gameLog.history" :key="`texas_round__${roundIndex}`">
          <tr v-for="(info, infoIndex) in roundInfo" :key="`texas_round__info__${info.index}`">
            <td v-if="infoIndex === 0" :rowspan="roundInfo.length" class="border-r border-b px-3.5 py-1">
              {{ t('fmtTextRound', [roundIndex + 1]) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              {{ t('fmtTextTexasHistorySeat', [info.seat, info.seatId]) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              <span :class="{ 'text-primary': props.userName === info.userName }">
                {{ info.userName }}
              </span>
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              {{ t(`teenpattiStatus__${info.status}`) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              {{ numberToStr(info.bet) }}
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
import { TeenpattiGameLog } from '@/base/common/gameLog/teenpatti'
import { numberToStr } from '@/base/utils/formatNumber'

import PokerCard from '@/base/components/PokerCard.vue'
import PlayerResult from '@/base/views/OperationManagement/GameLogParse/PlayerResult.vue'

const { t } = useI18n()

const cardSuitNames = ['diamond', 'club', 'heart', 'spade', 'black-joker', 'red-joker']

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

const gameLog = computed(() => new TeenpattiGameLog(props.playLog))
</script>
