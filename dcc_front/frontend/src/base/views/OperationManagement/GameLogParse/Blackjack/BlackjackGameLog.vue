<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('textDeskPlayers') }}:</div>
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextBlackjackPlayerLog"
        tag="p"
        scope="global"
      >
        <template #seat>{{ playerLog.seat }}(seatId:{{ playerLog.seatId }})</template>
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        <template #betDetail>{{ numberToStr(playerLog.yaScore) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textPlayerResult') }}:</div>
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextBlackjackPlayerResult"
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
        <template #result>
          <span
            v-for="(playCardInfo, cardInfoIndex) in playerLog.playCardInfos"
            :key="`playerCardInfo_${playerLog.userName}__${cardInfoIndex}`"
            class="flex items-center"
          >
            <template v-if="gameLog.isInsurance">
              {{ playerLog.isInsurance ? t('textBuyInsurance') : t('textNotBuyInsurance') }}
              {{ t('symbolComma2') }}
            </template>
            <template v-if="playCardInfo.isDouble">
              {{ t('textDoubleBet') }}
              {{ t('symbolComma2') }}
            </template>
            <PokerCard
              v-for="(card, cardIndex) in playCardInfo.cards"
              :key="`poker__${card}_card__${cardIndex}_${playerLog.userName}`"
              class="mr-1 mb-1"
              :index="card"
            />
            {{ t('symbolComma2') }}
            {{ t(`blackjackCardType__${playCardInfo.cardType}`, [playCardInfo.cardPoint]) }}
            <template v-if="playerLog.playCardInfos.length > 1 && cardInfoIndex != playerLog.playCardInfos.length - 1">
              {{ t('symbolComma3') }}
            </template>
          </span>
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textBankerResult') }}:</div>
      <i18n-t keypath="fmtTextBlackjackBankerResult" tag="div" scope="global" class="flex items-center">
        <template #userName>
          {{ t('textBaccaratBanker') }}
        </template>
        <template #result>
          <PokerCard
            v-for="(card, cardIndex) in gameLog.bankerPlayCardInfo.cards"
            :key="`poker__${card}_card__${cardIndex}_banker`"
            class="mr-1 mb-1"
            :index="card"
          />
          {{ t('symbolComma2') }}
          {{ t(`blackjackCardType__${gameLog.bankerPlayCardInfo.cardType}`, [gameLog.bankerPlayCardInfo.cardPoint]) }}
        </template>
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
import { BlackjackGameLog } from '@/base/common/gameLog/blackjack'
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

const gameLog = computed(() => new BlackjackGameLog(props.playLog))
</script>
