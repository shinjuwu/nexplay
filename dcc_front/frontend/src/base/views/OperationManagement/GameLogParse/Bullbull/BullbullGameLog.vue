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
        keypath="fmtTextBullbullPlayerLog"
        tag="div"
        scope="global"
        class="flex items-center"
      >
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        <template #betDetail>
          <span>{{ playerLog.isBanker ? t('textBaccaratBanker') : t('textBaccaratPlayer') }}</span>
          <span>{{ t('symbolComma2') }}</span>
          <span>{{
            playerLog.isBanker
              ? t('fmtTextBullbullBankerBet', [playerLog.bet])
              : t('fmtTextSangongPlayerBet', [playerLog.bet])
          }}</span>
          <span>{{ t('symbolComma2') }}</span>
          <i18n-t
            keypath="fmtTextSangongPlayerHandCardsInfo"
            tag="span"
            scope="global"
            class="inline-flex items-center"
          >
            <template #cards>
              <PokerCard
                v-for="(card, cardIdx) in playerLog.cards"
                :key="`poker__${card}_card__${cardIdx}_player`"
                class="mb-1 mr-1"
                :index="card"
              />
            </template>
            <template #cardType>{{ t(`bullbullCardType__${playerLog.cardType}`) }}</template>
          </i18n-t>
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
import { BullbullGameLog } from '@/base/common/gameLog/bullbull'
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

const gameLog = computed(() => new BullbullGameLog(props.playLog))
</script>
