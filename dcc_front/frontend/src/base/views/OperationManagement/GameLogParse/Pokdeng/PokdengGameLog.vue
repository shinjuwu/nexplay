<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>
        {{ t('fmtTextBaseBet', [numberToStr(gameLog.baseBet)]) }}
      </div>
    </div>

    <div class="mt-4">
      <div>{{ t('textDeskPlayers') }}:</div>
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLogBetInfo_${playerLog.userName}__${index}`"
        keypath="fmtTexPokdengPlayerLogBetInfo"
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
        <template #bet>{{ numberToStr(playerLog.yaScore) }}</template>
        <template #originalBet>{{ numberToStr(playerLog.bet) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textPlayerDrawResult') }}:</div>
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLogCardInfo_${playerLog.userName}__${index}`"
        keypath="fmtTexPokdengPlayerLogCardInfo"
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
        <template #handCards>
          &nbsp;
          <PokerCard
            v-for="(card, playerCardIndex) in playerLog.cards"
            :key="`poker__${card}_card__${playerCardIndex}_player_${playerLog.userName}`"
            class="mb-1 mr-1"
            :index="card"
          />
        </template>
        <template #odds>{{ numberToStr(playerLog.odds) }}</template>
        <template #cardType>
          {{ t(`pokdengCardType__${playerLog.cardType}`) + t(`pokdengStatus__${playerLog.status}`) }}
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textBankerDrawResult') }}:</div>
      <i18n-t keypath="fmtTexPokdengBankerLogCardInfo" tag="p" scope="global" class="flex items-center">
        <template #handCards>
          &nbsp;
          <PokerCard
            v-for="(card, index) in gameLog.bankerLog.cards"
            :key="`poker__${card}_card__${index}_banker`"
            class="mb-1 mr-1"
            :index="card"
          />
        </template>
        <template #odds>{{ numberToStr(gameLog.bankerLog.odds) }}</template>
        <template #cardType>
          {{ t(`pokdengCardType__${gameLog.bankerLog.cardType}`) + t(`pokdengStatus__${gameLog.bankerLog.status}`) }}
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
import { PokdengGameLog } from '@/base/common/gameLog/pokdeng'
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

const gameLog = computed(() => new PokdengGameLog(props.playLog))
</script>
