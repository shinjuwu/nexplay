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
          <PyrtreasureslotMainGameBoard v-if="gameLog.mainGameInfo" :board="gameLog.mainGameInfo.board" />
          <PyrtreasureslotBonusGameBoard
            v-if="gameLog.freeGameInfo"
            :main-game-lognumber="gameLog.freeGameInfo.maingGameLognumber"
            :bonus-board="gameLog.freeGameInfo.bonusBoard"
            :multi="gameLog.freeGameInfo.multi"
          />
        </div>
        <div>
          <PyrtreasureslotLineInfo v-if="gameLog.mainGameInfo" :lines="gameLog.lines" :bonus-wins="gameLog.bonusWins" />
          <PyrtreasureslotBGLineInfo
            v-if="gameLog.freeGameInfo"
            :bonus-line-info="gameLog.freeGameInfo.bonusLineInfo"
          />
        </div>
      </div>
    </div>

    <div class="mt-4">
      <i18n-t keypath="fmtTextSlotPlayerBetResult" tag="div" scope="global">
        <template #game>
          {{ gameLog.mainGameInfo ? t('textSlotMainGame') : gameLog.freeGameInfo ? t('textSlotBonusGame') : '' }}
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
import { PyrtreasureslotGameLog } from '@/base/common/gameLog/pyrtreasureslot'
import PyrtreasureslotLineInfo from '@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotLineInfo.vue'
import PyrtreasureslotBGLineInfo from '@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotBGLineInfo.vue'
import PyrtreasureslotBonusGameBoard from '@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotBonusGameBoard.vue'
import PyrtreasureslotMainGameBoard from '@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotMainGameBoard.vue'
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

const gameLog = computed(() => new PyrtreasureslotGameLog(props.playLog))
</script>
