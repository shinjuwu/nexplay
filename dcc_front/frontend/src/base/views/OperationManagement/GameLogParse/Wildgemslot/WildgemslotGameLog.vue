<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div class="mb-4">
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

    <div class="mb-4">
      <div>{{ t('textReelResult') }}:</div>

      <div class="mt-1 flex">
        <div class="mr-4">
          <WildgemslotMainGameBoard
            v-if="gameLog.mainGameInfo"
            :board="gameLog.mainGameInfo.board"
            :draw-lines="gameLog.drawLines"
          />
          <WildgemslotFreeGameBoard
            v-if="gameLog.freeGameInfo"
            :stage="gameLog.freeGameInfo.stage"
            :main-game-lognumber="gameLog.freeGameInfo.maingGameLognumber"
            :board="gameLog.freeGameInfo.board"
            :draw-lines="gameLog.drawLines"
          />
        </div>

        <div>
          <WildgemslotLineInfo :lines="gameLog.lines" :class="{ 'mt-[92px]': gameLog.freeGameInfo }" />
        </div>
      </div>
    </div>

    <div class="mb-4">
      <i18n-t keypath="fmtTextSlotPlayerBetResult" tag="div" scope="global">
        <template #game>
          {{ gameLog.mainGameInfo ? t('textSlotMainGame') : gameLog.freeGameInfo ? t('textSlotFreeGame') : '' }}
        </template>
        <template #yaScore>{{ numberToStr(gameLog.bet) }}</template>
        <template #deScore>{{ numberToStr(gameLog.playerLogs[0].deScore) }}</template>
      </i18n-t>
    </div>

    <div>
      <div>{{ t('textPayoutResult') }}:</div>
      <PlayerResult :player-log="gameLog.playerLogs[0]" :user-name="props.userName" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { WildgemslotGameLog } from '@/base/common/gameLog/wildgemslot'
import { numberToStr } from '@/base/utils/formatNumber'

import WildgemslotMainGameBoard from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotMainGameBoard.vue'
import WildgemslotFreeGameBoard from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotFreeGameBoard.vue'
import WildgemslotLineInfo from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotLineInfo.vue'
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

const gameLog = computed(() => new WildgemslotGameLog(props.playLog))
</script>
