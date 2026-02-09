<template>
  <div class="overflow-x-auto whitespace-nowrap">
    <div v-if="gameLog.friendRoomInfo">
      <div>{{ t('textFriendRoomInfo') }}</div>
      <div>{{ t('fmtTextFriendRoomInfoRoomId', [gameLog.friendRoomInfo.roomId]) }}</div>
      <div>{{ t('fmtTextFriendRoomInfoHostUserName', [gameLog.friendRoomInfo.hostUserName]) }}</div>
      <div>{{ t('fmtTextFriendRoomInfoTaxpercent', [gameLog.friendRoomInfo.taxpercent * 100]) }}</div>
      <div>{{ t('fmtTextFriendRoomInfoCurrentGame', [gameLog.friendRoomInfo.detail.currentGame]) }}</div>
      <div>{{ t('fmtTextFriendRoomInfoTotalGame', [gameLog.friendRoomInfo.detail.totalGame]) }}</div>
      <div>{{ t('fmtTextFriendRoomInfoAnte', [numberToStr(gameLog.friendRoomInfo.detail.ante)]) }}</div>
      <div>
        {{ t('fmtTextFriendRoomInfoEnterLimit', [numberToStr(gameLog.friendRoomInfo.detail.enterLimit)]) }}
      </div>
      <div>{{ t('fmtTextFriendRoomInfoBetLimit', [numberToStr(gameLog.friendRoomInfo.detail.betLimit)]) }}</div>
      <div>{{ t('fmtTextFriendRoomInfoBetSec', [gameLog.friendRoomInfo.detail.betSec]) }}</div>
    </div>

    <div :class="{ 'mt-4': gameLog.friendRoomInfo }">
      <div>
        {{ t('fmtTextTexasBlinds', [numberToStr(gameLog.smallBlind), numberToStr(gameLog.bigBlind)]) }}
      </div>
      <div>
        {{ t('fmtTextTexasDealer', [gameLog.dealerSeat, gameLog.dealerSeatId]) }}
      </div>
      <div>
        {{ t('fmtTextTotalGameTax', [numberToStr(gameLog.totalTax)]) }}
      </div>
      <i18n-t keypath="fmtTextTexasPublicCards" tag="p" scope="global" class="flex items-center">
        <template #publicCards>
          <PokerCard
            v-for="(card, index) in gameLog.pulicCards"
            :key="`poker__${card}_card__${index}_pulicCards`"
            class="mr-1"
            :index="card"
          />
        </template>
      </i18n-t>
    </div>

    <div class="mt-4">
      <i18n-t
        v-for="(playerLog, index) in gameLog.playerLogs"
        :key="`playerLog_${playerLog.userName}__${index}`"
        keypath="fmtTextTexasPlayerLog"
        tag="div"
        scope="global"
        class="flex items-center"
      >
        <template #seat>
          <span>{{ playerLog.seat }}(seatId:{{ playerLog.seatId }})</span>
        </template>
        <template #userName>
          <span :class="{ 'text-primary': props.userName === playerLog.userName }">
            {{ playerLog.userName }}
          </span>
        </template>
        <template #startScore>{{ numberToStr(playerLog.startScore) }}</template>
        <template #bet>{{ numberToStr(playerLog.bet) }}</template>
        <template #handCards>
          &nbsp;
          <PokerCard
            v-for="playerHandCard in playerLog.handCards"
            :key="`playerHandCards_${playerLog.userName}_${playerHandCard}`"
            class="mr-1 mb-1"
            :index="playerHandCard"
          />
        </template>
        <template #bestCards>
          &nbsp;
          <PokerCard
            v-for="playerBestCard in playerLog.bestCards"
            :key="`playerBestCards_${playerLog.userName}_${playerBestCard}`"
            class="mr-1 mb-1"
            :index="playerBestCard"
          />
        </template>
        <template #bestCardType>{{ t(`texasCardType__${playerLog.bestCardType}`) }}</template>
        <template #status>
          {{ t(`texasStatus__${playerLog.status}`, [t(`texasFoldState__${playerLog.foldState}`)]) }}
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
        keypath="fmtTextTexasPlayerResult"
      />
    </div>

    <table class="mt-4 border-l border-t text-center">
      <thead>
        <tr>
          <th colspan="5" class="border-r border-b px-3.5 py-1">{{ t('textGameProcessLog') }}</th>
        </tr>
        <tr>
          <th class="border-r border-b px-3.5 py-1">{{ t('textGameRound') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textSeat') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textPlayerAccount') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textProcessLog') }}</th>
          <th class="border-r border-b px-3.5 py-1">{{ t('textBets') }}</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(roundInfo, roundIndex) in gameLog.history" :key="`texas_round__${roundIndex}`">
          <tr v-for="(info, infoIndex) in roundInfo" :key="`texas_round__info__${info.index}`">
            <td v-if="infoIndex === 0" :rowspan="roundInfo.length" class="border-r border-b px-3.5 py-1">
              {{ t(`texasFoldState__${roundIndex + 1}`) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              {{ t('fmtTextTexasHistorySeat', [info.seat, info.seatId]) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              <span :class="{ 'text-primary': props.userName === info.userName }">
                {{ info.userName }}
              </span>
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              {{ t(`texasAction__${info.action}`) }}
            </td>
            <td class="border-r border-b px-3.5 py-1" :class="{ 'bg-slate-100': info.index % 2 === 0 }">
              {{ numberToStr(info.bet) }}
            </td>
          </tr>
        </template>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { TexasGameLog } from '@/base/common/gameLog/texas'
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

const gameLog = computed(() => new TexasGameLog(props.playLog))
</script>
