<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div>
      <i18n-t keypath="fmtTextPlinkoGameProcess" tag="div" scope="global">
        <template #userName>
          <span :class="{ 'text-primary': props.userName === gameLog.playerLogs[0].userName }">
            {{ gameLog.playerLogs[0].userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(gameLog.playerLogs[0].startScore) }}</template>
        <template #bet>{{ numberToStr(gameLog.playerLogs[0].yaScore) }}</template>
      </i18n-t>
    </div>

    <div class="mt-4">{{ t('textPlinkoBetScoreTotal') }}</div>

    <template v-for="ballInfo in gameLog.playerLogs[0].ballInfo" :key="`ballInfo__${ballInfo}`">
      <div v-if="gameLog.playerLogs[0].ballOddsBetWins[ballInfo]" class="mt-4">
        <div>{{ t('fmtTextPlinkoBallOdds', [ballInfo]) }}:</div>
        <table class="ml-12 text-center">
          <thead>
            <tr>
              <th class="bg-info border px-2 py-1 text-white">{{ t('textPlinkoBetSlotRateBet') }}</th>
              <th
                v-for="betInfo in gameLog.playerLogs[0].betInfo"
                :key="`betInfo__${ballInfo}__${betInfo}`"
                class="bg-info min-w-[120px] border px-2 py-1 text-white"
              >
                {{ betInfo }}
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="holeInfo in gameLog.playerLogs[0].holeInfo" :key="`holeInfo__${ballInfo}__${holeInfo}`">
              <td class="border px-2 py-1">{{ holeInfo }}</td>
              <td
                v-for="betInfo in gameLog.playerLogs[0].betInfo"
                :key="`ballOddsHoleOddsBetCount__${ballInfo}__${betInfo}__${holeInfo}`"
                class="border px-2 py-1"
              >
                {{
                  gameLog.playerLogs[0].ballOddsHoleOddsBetCount[ballInfo] &&
                  gameLog.playerLogs[0].ballOddsHoleOddsBetCount[ballInfo][holeInfo] &&
                  gameLog.playerLogs[0].ballOddsHoleOddsBetCount[ballInfo][holeInfo][betInfo]
                    ? gameLog.playerLogs[0].ballOddsHoleOddsBetCount[ballInfo][holeInfo][betInfo]
                    : 0
                }}
              </td>
            </tr>
            <tr>
              <td class="border px-2 py-1">{{ t('textPlinfoTotalWins') }}</td>
              <td
                v-for="betInfo in gameLog.playerLogs[0].betInfo"
                :key="`ballOddsBetWins__${ballInfo}__${betInfo}`"
                class="border px-2 py-1"
              >
                {{
                  numberToStr(
                    gameLog.playerLogs[0].ballOddsBetWins[ballInfo] &&
                      gameLog.playerLogs[0].ballOddsBetWins[ballInfo][betInfo]
                      ? gameLog.playerLogs[0].ballOddsBetWins[ballInfo][betInfo]
                      : 0
                  )
                }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </template>

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
import { PlinkoGameLog } from '@/base/common/gameLog/plinko'
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

const gameLog = computed(() => new PlinkoGameLog(props.playLog))
</script>
