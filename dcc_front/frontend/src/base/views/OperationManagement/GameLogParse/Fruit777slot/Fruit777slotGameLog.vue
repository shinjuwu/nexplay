<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <div>{{ t('textParticipants') }}:</div>
      <i18n-t keypath="fmtTextFruit777slotPlayerLog" tag="div" scope="global">
        <template #userName>
          <span :class="{ 'text-primary': props.userName === gameLog.playerLogs[0].userName }">
            {{ gameLog.playerLogs[0].userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(gameLog.playerLogs[0].startScore) }}</template>
        <template #bet>{{ numberToStr(gameLog.playerLogs[0].yaScore) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textReelResult') }}:</div>
      <div class="mt-1 flex">
        <div class="mr-4">
          <div v-for="row in 3" :key="`row__${row}`" class="flex">
            <div
              v-for="column in 3"
              :key="`reel_${(row - 1) * 3 + (column - 1)}`"
              class="fruit777slot"
              :class="{
                'fruit777slot-0': gameLog.reels[(row - 1) * 3 + (column - 1)] === 0,
                'fruit777slot-1': gameLog.reels[(row - 1) * 3 + (column - 1)] === 1,
                'fruit777slot-2': gameLog.reels[(row - 1) * 3 + (column - 1)] === 2,
                'fruit777slot-3': gameLog.reels[(row - 1) * 3 + (column - 1)] === 3,
                'fruit777slot-4': gameLog.reels[(row - 1) * 3 + (column - 1)] === 4,
                'fruit777slot-5': gameLog.reels[(row - 1) * 3 + (column - 1)] === 5,
                'fruit777slot-6': gameLog.reels[(row - 1) * 3 + (column - 1)] === 6,
                'fruit777slot-7': gameLog.reels[(row - 1) * 3 + (column - 1)] === 7,
                'fruit777slot-8': gameLog.reels[(row - 1) * 3 + (column - 1)] === 8,
                'fruit777slot-9': gameLog.reels[(row - 1) * 3 + (column - 1)] === 9,
              }"
            ></div>
          </div>
        </div>

        <div>
          <table v-if="gameLog.lines" class="mb-4 text-center">
            <thead>
              <tr class="bg-info text-white">
                <th class="w-60 border px-2 py-1">{{ t('textPaylines') }}</th>
                <th class="w-100 border px-2 py-1">{{ t('textFruit777slotLineSymbol') }}</th>
                <th class="w-60 border px-2 py-1">{{ t('textOdds') }}</th>
                <th class="w-60 border px-2 py-1">{{ t('textFruit777slotWinScore') }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="line in gameLog.lines" :key="`line_${line.symbol}__${line.index}`">
                <td class="w-60 border px-2 py-1">{{ line.index + 1 }}</td>
                <td class="w-100 space-x-1 whitespace-nowrap border px-2 py-1">
                  <div
                    v-for="(symbol, idx) in line.symbols"
                    :key="`line__${line.index}_symbol__${symbol}__${idx}`"
                    class="fruit777slot-general inline-block"
                    :class="{
                      'fruit777slot-general-0': symbol === 0,
                      'fruit777slot-general-1': symbol === 1,
                      'fruit777slot-general-2': symbol === 2,
                      'fruit777slot-general-3': symbol === 3,
                      'fruit777slot-general-4': symbol === 4,
                      'fruit777slot-general-5': symbol === 5,
                      'fruit777slot-general-6': symbol === 6,
                      'fruit777slot-general-7': symbol === 7,
                      'fruit777slot-general-8': symbol === 8,
                      'fruit777slot-general-9': symbol === 9,
                      'fruit777slot-general-10': symbol === 10,
                      'fruit777slot-general-11': symbol === 11,
                      'fruit777slot-general-12': symbol === 12,
                      'fruit777slot-general-99': symbol === 99,
                    }"
                  ></div>
                </td>
                <td class="w-60 border px-2 py-1">{{ line.odds.toLocaleString() }}</td>
                <td class="w-60 border px-2 py-1">{{ line.wins.toLocaleString() }}</td>
              </tr>
            </tbody>
          </table>

          <table v-if="gameLog.special7" class="mb-4 text-center">
            <thead>
              <tr class="bg-info text-white">
                <th class="w-60 border px-2 py-1">{{ t('textFruit777slotSpecial7') }}</th>
                <th class="w-100 border px-2 py-1">{{ t('textFruit777slotSpecial7Symbol') }}</th>
                <th class="w-60 border px-2 py-1">{{ t('textOdds') }}</th>
                <th class="w-60 border px-2 py-1">{{ t('textFruit777slotWinScore') }}</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="border px-2 py-1">-</td>
                <td class="whitespace-nowrap border px-2 py-1">
                  <div
                    v-for="symbol in gameLog.special7.symbols"
                    :key="`specialSymbol_${symbol.symbol}`"
                    class="inline-block p-1"
                  >
                    <div
                      class="fruit777slot-general inline-block"
                      :class="{
                        'fruit777slot-general-0': symbol.symbol === 0,
                        'fruit777slot-general-1': symbol.symbol === 1,
                      }"
                    ></div>
                    x {{ symbol.count }}
                  </div>
                </td>
                <td class="border px-2 py-1">{{ gameLog.special7.odds.toLocaleString() }}</td>
                <td class="border px-2 py-1">{{ gameLog.special7.wins.toLocaleString() }}</td>
              </tr>
            </tbody>
          </table>

          <table v-if="gameLog.jp" class="text-center">
            <thead>
              <tr class="bg-info text-white">
                <th class="w-60 border px-2 py-1">{{ t('textFruit777slotJp') }}</th>
                <th class="w-100 border px-2 py-1">{{ t('textFruit777slotJpSymbol') }}</th>
                <th class="w-60 border px-2 py-1">{{ t('textOdds') }}</th>
                <th class="w-60 border px-2 py-1">{{ t('textFruit777slotWinScore') }}</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="border px-2 py-1">-</td>
                <td class="border px-2 py-1">
                  <div
                    class="fruit777slot-full mt-3 inline-block"
                    :class="{
                      'fruit777slot-full-0': gameLog.jp.symbol === 0,
                      'fruit777slot-full-1': gameLog.jp.symbol === 1,
                      'fruit777slot-full-2': gameLog.jp.symbol === 2,
                      'fruit777slot-full-3': gameLog.jp.symbol === 3,
                      'fruit777slot-full-4': gameLog.jp.symbol === 4,
                      'fruit777slot-full-5': gameLog.jp.symbol === 5,
                      'fruit777slot-full-6': gameLog.jp.symbol === 6,
                      'fruit777slot-full-7': gameLog.jp.symbol === 7,
                      'fruit777slot-full-8': gameLog.jp.symbol === 8,
                      'fruit777slot-full-9': gameLog.jp.symbol === 9,
                      'fruit777slot-full-10': gameLog.jp.symbol === 10,
                      'fruit777slot-full-11': gameLog.jp.symbol === 11,
                      'fruit777slot-full-12': gameLog.jp.symbol === 12,
                    }"
                  ></div>
                </td>
                <td class="border px-2 py-1">{{ gameLog.jp.odds.toLocaleString() }}</td>
                <td class="border px-2 py-1">{{ gameLog.jp.wins.toLocaleString() }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <div class="mt-4">
      <i18n-t keypath="fmtTextFruit777slotPlayerBetResult" tag="div" scope="global">
        <template #yaScore>{{ numberToStr(gameLog.playerLogs[0].yaScore) }}</template>
        <template #baseBet>{{ numberToStr(gameLog.baseBet) }}</template>
        <template #deScore>{{ numberToStr(gameLog.playerLogs[0].deScore) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <div>{{ t('textPayoutResult') }}:</div>
      <PlayerResult :player-log="gameLog.playerLogs[0]" :user-name="props.userName" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { Fruit777slotGameLog } from '@/base/common/gameLog/fruit777slot'
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

const gameLog = computed(() => new Fruit777slotGameLog(props.playLog))
</script>
