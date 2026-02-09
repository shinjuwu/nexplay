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
          <MidasslotMainGameBoard
            v-if="gameLog.mainGameInfo"
            :board="gameLog.mainGameInfo.board"
            :multi-board="gameLog.multiBoard"
            :draw-lines="gameLog.drawLines"
          />
          <MidasslotFreeGameBoard
            v-if="gameLog.freeGameInfo"
            :stage="gameLog.freeGameInfo.stage"
            :main-game-lognumber="gameLog.freeGameInfo.maingGameLognumber"
            :board="gameLog.freeGameInfo.board"
            :multi-board="gameLog.multiBoard"
            :draw-lines="gameLog.drawLines"
          />
        </div>

        <div>
          <MidasslotLineInfo
            :lines="gameLog.lines"
            :class="{
              'mt-[92px]': gameLog.freeGameInfo,
              mg: gameLog.slotType === SLOT_TYPE.NG,
              fg: gameLog.slotType === SLOT_TYPE.FG,
            }"
          />
        </div>
      </div>
    </div>

    <div class="mt-4">
      <i18n-t keypath="fmtTextSlotPlayerBetResult" tag="div" scope="global">
        <template #game>
          {{ gameLog.mainGameInfo ? t('textSlotMainGame') : gameLog.freeGameInfo ? t('textSlotFreeGame') : '' }}
        </template>
        <template #yaScore>{{ numberToStr(gameLog.bet) }}</template>
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
import { SLOT_TYPE } from '@/base/common/gameLog/slotGame'
import { MidasslotGameLog } from '@/base/common/gameLog/midasslot'
import { numberToStr } from '@/base/utils/formatNumber'

import MidasslotMainGameBoard from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotMainGameBoard.vue'
import MidasslotFreeGameBoard from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotFreeGameBoard.vue'
import MidasslotLineInfo from '@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotLineInfo.vue'
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

const gameLog = computed(() => new MidasslotGameLog(props.playLog))
</script>
