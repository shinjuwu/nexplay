<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <HappyfishingPlayerBetInfo class="mb-4" :player-log="gameLog.playerLog" :target-user-name="props.userName" />
    <HappyfishingPlayerGameProcessingInfo
      class="mb-4"
      :player-log="gameLog.playerLog"
      :target-user-name="props.userName"
    />

    <div class="mt-4">
      <div>{{ t('textPayoutResult') }}:</div>
      <PlayerResult :player-log="gameLog.playerLog" :user-name="props.userName" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { HappyfishingGameLog } from '@/base/common/gameLog/happyfishing'

import HappyfishingPlayerBetInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerBetInfo.vue'
import HappyfishingPlayerGameProcessingInfo from '@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingPlayerGameProcessingInfo.vue'
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

const gameLog = computed(() => new HappyfishingGameLog(props.playLog))
</script>
