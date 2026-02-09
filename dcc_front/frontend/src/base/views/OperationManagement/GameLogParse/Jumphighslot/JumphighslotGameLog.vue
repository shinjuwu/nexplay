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
          <JumphighslotMainGameBoard
            v-if="gameLog.mainGameInfo"
            :board="gameLog.mainGameInfo.board"
            :board-active="gameLog.boardActive"
          />
          <JumphighslotFreeGameBoard
            v-if="gameLog.freeGameInfo"
            :original-board="gameLog.freeGameInfo.originalBoard"
            :board-active="gameLog.boardActive"
            :stage="gameLog.freeGameInfo.stage"
            :multi-board="gameLog.multiBoard"
            :main-game-lognumber="gameLog.freeGameInfo.maingGameLognumber"
          />
        </div>
        <div>
          <JumphighslotLineInfo :lines="gameLog.lines" :wild-multi="gameLog.wildMulti" />
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
import { JumphighslotGameLog } from '@/base/common/gameLog/jumphighslot'
import JumphighslotLineInfo from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotLineInfo.vue'
import JumphighslotFreeGameBoard from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotFreeGameBoard.vue'
import JumphighslotMainGameBoard from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotMainGameBoard.vue'
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

const gameLog = computed(() => new JumphighslotGameLog(props.playLog))
</script>
