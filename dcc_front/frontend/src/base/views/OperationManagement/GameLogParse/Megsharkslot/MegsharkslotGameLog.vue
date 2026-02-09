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
          <MegsharkslotMainGameBoard
            v-if="gameLog.mainGameInfo"
            :board="gameLog.mainGameInfo.board"
            :board-active="gameLog.boardActive"
          />
          <MegsharkslotFreeGameBoards
            v-if="gameLog.freeGameInfo"
            :stage="gameLog.freeGameInfo.stage"
            :main-game-lognumber="gameLog.freeGameInfo.maingGameLognumber"
            :left-transform-board="gameLog.freeGameInfo.leftTransformBoard"
            :original-board="gameLog.freeGameInfo.originalBoard"
            :ww-board="gameLog.freeGameInfo.wwBoard"
            :board-active="gameLog.boardActive"
            :same-board="gameLog.freeGameInfo.sameBoard"
          />
        </div>

        <div>
          <MegsharkslotLineInfo :lines="gameLog.lines" :class="{ 'mt-[92px]': gameLog.freeGameInfo }" />
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
import { MegsharkslotGameLog } from '@/base/common/gameLog/megsharkslot'
import { numberToStr } from '@/base/utils/formatNumber'

import MegsharkslotMainGameBoard from '@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotMainGameBoard.vue'
import MegsharkslotFreeGameBoards from '@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotFreeGameBoards.vue'
import MegsharkslotLineInfo from '@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotLineInfo.vue'
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

const gameLog = computed(() => new MegsharkslotGameLog(props.playLog))
</script>
