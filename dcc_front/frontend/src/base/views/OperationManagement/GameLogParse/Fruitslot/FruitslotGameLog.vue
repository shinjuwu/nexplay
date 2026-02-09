<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>
        {{ t('fmtTextBaseBet', [numberToStr(gameLog.baseBet)]) }}
      </div>
    </div>

    <div class="mt-4">
      <div>{{ t('textFruitslotPlayerBetTitle') }}</div>
      <div v-for="(playerLog, index) in gameLog.playerLogs" :key="`playerLog_${playerLog.userName}__${index}`">
        <i18n-t keypath="fmtTextFruitslotPlayerLog" tag="div" scope="global" class="my-1">
          <template #userName>
            <span :class="{ 'text-primary': props.userName === playerLog.userName }">
              {{ playerLog.userName }}
            </span>
          </template>
          <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        </i18n-t>

        <div v-for="(rowSymbols, rowIndex) in gameLog.boards" :key="`board_row_${rowIndex}`" class="flex">
          <div
            class="flex gap-1 bg-gray-300 px-2 pt-2"
            :class="{ 'pb-2 rounded-b': rowIndex === gameLog.boards.length - 1, 'rounded-t': rowIndex === 0 }"
          >
            <div
              v-for="(symbol, symbolIndex) in rowSymbols"
              :key="`board_symbol_${rowIndex}_${symbolIndex}`"
              class="fruitslot"
              :class="{
                'fruitslot-crown': symbol.name === 'crown',
                'fruitslot-50xcrown': symbol.name === '50xcrown',
                'fruitslot-apple': symbol.name === 'apple',
                'fruitslot-3xapple': symbol.name === '3xapple',
                'fruitslot-grape': symbol.name === 'grape',
                'fruitslot-3xgrape': symbol.name === '3xgrape',
                'fruitslot-mangosteen': symbol.name === 'mangosteen',
                'fruitslot-3xmangosteen': symbol.name === '3xmangosteen',
                'fruitslot-clover': symbol.name === 'clover',
                'fruitslot-lemon': symbol.name === 'lemon',
                'fruitslot-3xlemon': symbol.name === '3xlemon',
                'fruitslot-bell': symbol.name === 'bell',
                'fruitslot-3xbell': symbol.name === '3xbell',
                'fruitslot-seven': symbol.name === 'seven',
                'fruitslot-3xseven': symbol.name === '3xseven',
                'fruitslot-diamond': symbol.name === 'diamond',
                'fruitslot-3xdiamond': symbol.name === '3xdiamond',
                active:
                  playerLog.detail.prizeType === symbol.prizeType ||
                  (playerLog.detail.spPrizes && playerLog.detail.spPrizes.indexOf(symbol.prizeType) >= 0),
              }"
            ></div>
          </div>
        </div>

        <div class="mt-1 flex">
          <div
            v-for="betButton in gameLog.betButtons"
            :key="`betButton_${betButton.name}`"
            class="fruitslot-button-container flex flex-col items-center"
          >
            <div class="flex-1">{{ gameLog.rates[betButton.rateIndex] }}</div>
            <div class="flex-1">{{ playerLog.betDetails[betButton.betIndex].betCount }}</div>
            <div class="flex-1">{{ numberToStr(playerLog.betDetails[betButton.betIndex].bet) }}</div>
            <div
              class="fruitslot-button"
              :class="{
                'fruitslot-button-crown': betButton.name === 'crown',
                'fruitslot-button-apple': betButton.name === 'apple',
                'fruitslot-button-grape': betButton.name === 'grape',
                'fruitslot-button-mangosteen': betButton.name === 'mangosteen',
                'fruitslot-button-clover': betButton.name === 'clover',
                'fruitslot-button-lemon': betButton.name === 'lemon',
                'fruitslot-button-bell': betButton.name === 'bell',
                'fruitslot-button-seven': betButton.name === 'seven',
                'fruitslot-button-diamond': betButton.name === 'diamond',
              }"
            ></div>
          </div>
          <div class="fruitslot-button-container flex flex-col items-center">
            <div class="flex-1">{{ t('textOdds') }}</div>
            <div class="flex-1">{{ t('textBetNumbers') }}</div>
            <div class="flex-1">{{ t('textBets') }}</div>
            <div class="fruitslot-button"></div>
          </div>
        </div>

        <div v-if="playerLog.detail.spPrizeType" class="mt-1">
          {{ t('fmtTextFruitslotSpecialPrize', [t(`fruitslotSpecialPrizeType__${playerLog.detail.spPrizeType}`)]) }}
        </div>

        <template v-if="playerLog.detail.guessDetails">
          <div class="mt-1">{{ t('textGuessBigSmall') }}</div>
          <div
            v-for="(detail, detailIndex) in playerLog.detail.guessDetails"
            :key="`guessDetail_${playerLog.userName}__${detailIndex}`"
          >
            {{
              t('fmtTextFruitslotGuessResult', [
                detailIndex + 1,
                numberToStr(detail.bet),
                t(`fruitslotOddEven__${detail.betType}`),
                detail.result,
                t(`fruitslotOddEven__${detail.resultType}`),
              ])
            }}
          </div>
        </template>
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
import { FruitslotGameLog } from '@/base/common/gameLog/fruitslot'
import { numberToStr } from '@/base/utils/formatNumber'

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

const gameLog = computed(() => new FruitslotGameLog(props.playLog))
</script>
