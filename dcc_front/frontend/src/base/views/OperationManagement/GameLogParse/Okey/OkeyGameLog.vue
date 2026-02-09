<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('fmtTextBaseBet', [numberToStr(gameLog.perPoint)]) }}</div>
      <div>{{ t('fmtTextTotalGameTax', [numberToStr(gameLog.totalTax)]) }}</div>
      <i18n-t keypath="fmtTextRummySpecialWildCard" tag="span" scope="global">
        <template #default>
          <OkeyCard :index="gameLog.indicator" />
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextOkeyPlayerLogCardInfo"
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
            <OkeyCard
              v-for="card in cardSet"
              :key="`playerCard_${playerLog.userName}_${cardSetIdx}_${card}`"
              class="mr-1 mb-1"
              :index="card"
              :indicator-index="gameLog.indicator"
            />
          </template>
          <template v-if="playerLog.cardSet.length === 0">{{ t('textNone') }}</template>
        </template>
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

    <table v-show="gameLog.history.length > 0" class="mt-4 border-l border-t text-center">
      <thead>
        <tr>
          <th colspan="4" class="border-r border-b px-3.5 py-1">{{ t('textGameHistoryLog') }}</th>
        </tr>
        <tr>
          <th class="border-r border-b px-3.5 py-1">{{ t('textSeat') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textPlayerAccount') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textProcessLog') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textCurrentHandCards') }}</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(roundHistories, roundIndex) in gameLog.history" :key="`okey_round__${roundIndex}`">
          <tr>
            <td td class="border-r border-b px-3.5 py-1">
              {{ t('fmtTextTexasHistorySeat', [roundHistories.seat, roundHistories.seatId]) }}
            </td>
            <td td class="border-r border-b px-3.5 py-1">
              <span :class="{ 'text-primary': props.userName === roundHistories.userName }">
                {{ roundHistories.userName }}
              </span>
            </td>
            <td class="border-r border-b px-3.5 py-1">
              {{ t(`okeyAction__${roundHistories.action}`) }}
              <OkeyCard
                v-if="roundHistories.card >= 0"
                :key="`historyCard_${roundHistories.userName}_${roundIndex}_${roundHistories.card}`"
                class="ml-1"
                :index="roundHistories.card"
                :indicator-index="gameLog.indicator"
              />
            </td>
            <td class="border-r border-b px-3.5 py-1">
              <div v-if="roundHistories.cardSet.length > 0" class="flex items-center justify-center">
                <template
                  v-for="(cardSet, cardSetIdx) in roundHistories.cardSet"
                  :key="`historyCardSet_${roundHistories.userName}_${cardSetIdx}`"
                >
                  <OkeyCard
                    v-for="(card, cardIdx) in cardSet"
                    :key="`playerCard_${roundHistories.userName}_${cardSetIdx}_${cardIdx}`"
                    class="mr-1 mb-1"
                    :index="card"
                    :indicator-index="gameLog.indicator"
                  />
                </template>
              </div>
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
import { OkeyGameLog } from '@/base/common/gameLog/okey'
import { numberToStr } from '@/base/utils/formatNumber'

import OkeyCard from '@/base/components/OkeyCard.vue'
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

const gameLog = computed(() => new OkeyGameLog(props.playLog))
</script>
