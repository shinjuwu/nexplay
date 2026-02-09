<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>Per Point：{{ numberToStr(gameLog.perPoint) }}</div>
      <div>{{ t('fmtTextTotalGameTax', [numberToStr(gameLog.totalTax)]) }}</div>
      <i18n-t keypath="fmtTextRummySpecialWildCard" tag="div" scope="global" class="flex items-center">
        <template #default>
          <PokerCard :index="gameLog.wildCard" />
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextRummyPlayerLog"
        tag="div"
        scope="global"
        class="flex items-center"
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
          <template
            v-for="(cardSet, cardSetIdx) in playerLog.cardSet"
            :key="`playerCardSet_${playerLog.userName}_${cardSetIdx}`"
          >
            <PokerCard
              v-for="card in cardSet"
              :key="`playerCard_${playerLog.userName}_${cardSetIdx}_${card}`"
              class="mr-1 mb-1"
              :index="card"
            />
            <span v-if="cardSetIdx !== playerLog.cardSet.length - 1" class="mr-1">|</span>
          </template>
          <template v-if="playerLog.cardSet.length === 0">{{ t('textNone') }}</template>
        </template>
        <template #pointResult>{{ numberToStr(playerLog.cardPoint) }}</template>
        <template #result>{{ t(`rummyResult__${playerLog.result}`) }}</template>
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
          <th colspan="4" class="border-r border-b px-3.5 py-1">{{ t('textGameProcessLog') }}</th>
        </tr>
        <tr>
          <th class="border-r border-b px-3.5 py-1">{{ t('textSeat') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textPlayerAccount') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textProcessLog') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textCurrentHandCards') }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(info, infoIndex) in gameLog.history" :key="`rummy_info__${infoIndex}`">
          <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': infoIndex % 2 === 0 }">
            {{ t('fmtTextTexasHistorySeat', [info.seat, info.seatId]) }}
          </td>
          <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': infoIndex % 2 === 0 }">
            <span :class="{ 'text-primary': props.userName === info.userName }">
              {{ info.userName }}
            </span>
          </td>
          <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': infoIndex % 2 === 0 }">
            <div class="flex items-center justify-center">
              {{ t(`rummyAction__${info.action}`) }}
              <PokerCard
                v-if="info.card >= 0"
                :key="`historyCard_${info.userName}_${infoIndex}_${info.card}`"
                class="ml-1"
                :index="info.card"
              />
            </div>
          </td>
          <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': infoIndex % 2 === 0 }">
            <div v-if="info.cardSet.length > 0" class="flex items-center justify-center">
              <template
                v-for="(cardSet, cardSetIdx) in info.cardSet"
                :key="`historyCardSet_${info.userName}_${cardSetIdx}`"
              >
                <PokerCard
                  v-for="card in cardSet"
                  :key="`playerCard_${info.userName}_${cardSetIdx}_${card}`"
                  class="mr-1 mb-1"
                  :index="card"
                />
                <span v-if="cardSetIdx !== info.cardSet.length - 1" class="mr-1">|</span>
              </template>
            </div>
            <template v-else>{{ t('textNone') }}</template>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { RummyGameLog } from '@/base/common/gameLog/rummy'
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

const gameLog = computed(() => new RummyGameLog(props.playLog))
</script>
