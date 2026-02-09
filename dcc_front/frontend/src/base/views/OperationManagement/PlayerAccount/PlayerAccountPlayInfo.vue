<template>
  <PageTableDialog :visible="visibleGameDialog" @close="closeGameDialog()">
    <template #header>{{ t('textPlayInfo') }}</template>
    <template #default>
      <div class="mb-3 grid grid-cols-2">
        <div>{{ t('fmtTextUserName', [playerInfo.userName]) }}</div>
        <div>{{ t('fmtTextTotalPlayCount', [playerInfo.totalPlays]) }}</div>
      </div>
      <div class="tbl-container max-h-[70vh] overflow-y-auto">
        <table class="tbl">
          <thead>
            <tr>
              <th>{{ t('textGameCode') }}</th>
              <th>{{ t('textGameName') }}</th>
              <th>{{ t('textPlayCount') }} / {{ t('textNumberOfGamesUpgradeNonNewbie') }}</th>
              <th>{{ t('textOperate') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="info in gamePlayInfo" :key="`gamePlayInfo__${info.gameId}`">
              <td>{{ info.gameCode }}</td>
              <td>{{ t(`game__${info.gameId}`) }}</td>
              <td :class="{ 'text-success': info.totalCount <= info.newbieLimit }">
                {{ info.totalCount }} / {{ info.newbieLimit }}
              </td>
              <td>
                <ViewButton :tips-text="t('textRoomPlayInfo')" @click="showGameUserGamePlayDetail(info.gameId)" />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto">
        <button type="button" class="btn btn-light" @click="closeGameDialog()">
          {{ t('textClose') }}
        </button>
      </div>
    </template>
  </PageTableDialog>

  <PageTableDialog :visible="visibleGameRoomDialog" @close="closeGameRoomDialog()">
    <template #header>{{ t('textRoomPlayInfo') }}</template>
    <template #default>
      <div class="tbl-container">
        <table class="tbl">
          <thead>
            <tr>
              <th>{{ t('textRoomCode') }}</th>
              <th>{{ t('textRoomType') }}</th>
              <th>{{ t('textPlayCount') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="gamePlayDetailItem in gamePlayDetail.items"
              :key="`gamePlayDetail__${gamePlayDetailItem.roomId}`"
            >
              <td>{{ gamePlayDetailItem.roomId }}</td>
              <td>{{ t(`roomType__${roomTypeNameIndex(gamePlayDetailItem.gameId, gamePlayDetailItem.roomType)}`) }}</td>
              <td>{{ gamePlayDetailItem.playCount }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto">
        <button type="button" class="btn btn-light" @click="closeGameRoomDialog()">
          {{ t('textClose') }}
        </button>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { usePlayerAccountPlayInfo } from '@/base/composable/operationManagement/playerAccount/usePlayerAccountPlayInfo'
import { roomTypeNameIndex } from '@/base/utils/room'

import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  playerInfo: {
    type: Object,
    default: () => {},
  },
})
const emit = defineEmits(['close'])

const { t } = useI18n()

const {
  gamePlayInfo,
  gamePlayDetail,
  playerInfo,
  visibleGameDialog,
  visibleGameRoomDialog,
  closeGameDialog,
  closeGameRoomDialog,
  showGameUserGamePlayDetail,
} = usePlayerAccountPlayInfo(props, emit)
</script>
