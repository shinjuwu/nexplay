<template>
  <PageTableDialog :visible="visible" @close="emit('close')">
    <template #header> {{ t('textGamesRTPSetting') }} </template>
    <template #default>
      <div class="mb-2 grid grid-cols-1 gap-2 md:grid-cols-3 md:gap-6">
        <div v-if="selectRecord.gameType !== constant.GameType.BaiRen" class="form-label col-span-1">
          {{ t('fmtTextAgentName', [selectRecord.agentName]) }}
        </div>
        <div v-if="selectRecord.gameId !== constant.Game.All" class="form-label col-span-1">
          {{ t('fmtTextGameName', [t(`game__${selectRecord.gameId}`)]) }}
        </div>
        <div v-if="selectRecord.roomType !== constant.RoomType.All" class="form-label col-span-1">
          {{ t('fmtTextRoomType', [t(`roomType__${roomTypeNameIndex(selectRecord.gameId, selectRecord.roomType)}`)]) }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-2 md:grid-cols-3 md:gap-6">
        <div class="col-span-1">
          <div class="form-label flex items-center">
            <span class="min-w-fit">{{ t('textRTP') }}</span>
            <input
              v-model="selectRecord.killRatio"
              max="50"
              min="0"
              step="0.1"
              class="form-input mx-1"
              type="number"
              :disabled="selectRecord.isParent"
            />
            <span>%</span>
          </div>
          <div v-if="selectRecord.isParent" class="text-danger">
            {{ `（${t('textUsingGeneralAgentRTPSet')}）` }}
          </div>
        </div>
        <div class="col-span-1">
          <div class="form-label flex items-center">
            <span class="mr-1 min-w-fit">{{ t('textNewKillRatio') }}</span>
            <input
              v-model="selectRecord.newKillRatio"
              :disabled="props.gameType === constant.GameType.BaiRen"
              max="100"
              min="0"
              step="0.1"
              class="form-input mx-1"
              type="number"
            />
            <span>%</span>
          </div>
        </div>
        <div class="col-span-1">
          <div class="form-label flex items-center">
            <span class="min-w-fit">{{ t('textActiveNum') }}</span>
            <input
              v-model="selectRecord.activeNum"
              :disabled="props.gameType !== constant.GameType.BaiRen"
              max="999"
              min="0"
              step="1"
              class="form-input mx-1"
              type="number"
            />
          </div>
        </div>
      </div>
      <div>
        <label class="form-label">{{ t('textRemark') }}</label>
        <textarea v-model="selectRecord.remark" class="form-input" type="text" maxlength="100" />
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="emit('close')">
          {{ t('textClose') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="props.visible"
          :parent-data="selectRecord"
          :button-click="() => emit('setGameRatio', selectRecord)"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { reactive, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import { roomTypeNameIndex } from '@/base/utils/room'
import { round } from '@/base/utils/math'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()
const props = defineProps({
  selectRecord: {
    type: Object,
    default: () => {},
  },
  visible: {
    type: Boolean,
    default: false,
  },
  gameType: {
    type: Number,
    default: 0,
  },
})

const emit = defineEmits(['close', 'setGameRatio'])

const visible = ref(false)

const selectRecord = reactive({
  agentName: '',
  gameId: constant.Game.All,
  roomType: constant.RoomType.All,
  killRatio: 0,
  newKillRatio: 0,
  activeNum: 0,
  remark: '',
  isParent: false,
})

watch(
  () => props.visible,
  (newValue) => {
    if (newValue) {
      Object.assign(selectRecord, {
        id: props.selectRecord.id,
        agentName: props.selectRecord.agentName,
        gameId: props.selectRecord.gameId,
        roomType: props.selectRecord.roomType,
        gameType: props.selectRecord.gameType,
        killRatio: round(props.selectRecord.killRatio * 100, 4),
        newKillRatio: round(props.selectRecord.newKillRatio * 100, 4),
        activeNum: props.selectRecord.activeNum,
        remark: props.selectRecord.remark,
        isParent: props.selectRecord.isParent,
      })
    }
    visible.value = props.visible
  }
)
</script>
