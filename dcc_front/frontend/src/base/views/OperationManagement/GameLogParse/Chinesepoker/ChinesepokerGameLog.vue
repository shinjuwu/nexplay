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
        keypath="fmtTextChinesepokerPlayerLog"
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
        <template #cardType>{{ t(`chinesepokerHandCardType__${playerLog.handCardType}`) }}</template>
        <template #point>{{ playerLog.point }}</template>
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

    <table class="mt-4 border-l border-t text-center">
      <thead>
        <tr>
          <th colspan="12" class="border-r border-b px-3.5 py-1">{{ t('textGameHistoryLog') }}</th>
        </tr>
        <tr>
          <th class="border-r border-b px-3.5 py-1">{{ t('textSeat') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textPlayerAccount') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textChinesepokerHeadHandCards') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textCardType') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textChinesepokerMiddleHandCards') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textCardType') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textChinesepokerTailHandCards') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textCardType') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textChinesepokerShoot') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textChinesepokerGetShoot') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textChinesepokerPoint') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textResult') }}</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(playerLog, index) in gameLog.playerLogs"
          :key="`playerInfo_${playerLog.userName}__${index}`"
          :class="{ 'bg-slate-100': index % 2 === 0 }"
        >
          <td class="border-r border-b px-3.5 py-1">
            {{ t('fmtTextTexasHistorySeat', [playerLog.seat, playerLog.seatId]) }}
          </td>
          <td class="border-r border-b px-3.5 py-1">
            <span :class="{ 'text-primary': props.userName === playerLog.userName }">
              {{ playerLog.userName }}
            </span>
          </td>
          <td class="border-r border-b px-3.5 py-1">
            <PokerCard
              v-for="card in playerLog.cardsArray[0]"
              :key="`headCard__${card}`"
              class="mb-1 mr-1"
              :index="card"
              :card-suit-names="cardSuitNames"
            />
          </td>
          <td class="border-r border-b px-3.5 py-1">
            {{ playerLog.isSpecialCardType ? '-' : t(`chinesepokerArrayType__${playerLog.arrayType[0]}`) }}
          </td>
          <td class="border-r border-b px-3.5 py-1">
            <PokerCard
              v-for="card in playerLog.cardsArray[1]"
              :key="`middleCard__${card}`"
              class="mb-1 mr-1"
              :index="card"
              :card-suit-names="cardSuitNames"
            />
          </td>
          <td class="border-r border-b px-3.5 py-1">
            {{ playerLog.isSpecialCardType ? '-' : t(`chinesepokerArrayType__${playerLog.arrayType[1]}`) }}
          </td>
          <td class="border-r border-b px-3.5 py-1">
            <PokerCard
              v-for="card in playerLog.cardsArray[2]"
              :key="`tailCard__${card}`"
              class="mb-1 mr-1"
              :index="card"
              :card-suit-names="cardSuitNames"
            />
          </td>
          <td class="border-r border-b px-3.5 py-1">
            {{
              playerLog.isSpecialCardType
                ? t(`chinesepokerHandCardType__${playerLog.handCardType}`)
                : t(`chinesepokerArrayType__${playerLog.arrayType[2]}`)
            }}
          </td>
          <td class="border-r border-b px-3.5 py-1">{{ playerLog.shoot }}</td>
          <td class="border-r border-b px-3.5 py-1">{{ playerLog.getShoot }}</td>
          <td class="border-r border-b px-3.5 py-1">{{ playerLog.point }}</td>
          <td class="border-r border-b px-3.5 py-1">
            {{ `${playerLog.resultScore >= 0 ? '+' : ''}${playerLog.resultScore}` }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { ChinesepokerGameLog } from '@/base/common/gameLog/chinesepoker'
import { numberToStr } from '@/base/utils/formatNumber'

import PokerCard from '@/base/components/PokerCard.vue'
import PlayerResult from '@/base/views/OperationManagement/GameLogParse/PlayerResult.vue'

const { t } = useI18n()

const cardSuitNames = ['diamond', 'club', 'heart', 'spade']

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

const gameLog = computed(() => new ChinesepokerGameLog(props.playLog))
</script>
