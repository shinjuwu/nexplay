<template>
  <PageTableDialog :visible="visible" @close="close()">
    <template #header>{{ t('roleItemRTPBatchSettingUpdate') }}</template>
    <template #default>
      <div class="flex items-center justify-between px-4">
        <div v-if="showAgentSelect" class="flex flex-col">
          <div class="space-x-4">
            <span>{{ t('textSelectAgent') }}</span>
            <input
              v-model="selectAll.agent"
              type="checkbox"
              @change="(e) => toggleSelectAll('agent', e.target.checked)"
            />
            <span>{{ t('textSelectAllOrNot') }}</span>
          </div>
          <div class="flex h-[250px] w-[200px] flex-col overflow-y-scroll border px-1">
            <label v-for="agent in dataList.agent" :key="`selectAgentIndex__${agent.id}`">
              <input v-model="agent.selected" type="checkbox" class="mr-1" @change="toggleSelect('agent', agent)" />
              <span> {{ agent.name }} </span>
            </label>
          </div>
        </div>

        <ArrowRightIcon v-if="showAgentSelect" class="h-14 w-14 text-gray-300" />

        <div class="flex flex-col">
          <div class="space-x-4">
            <span>{{ t('textSelectGame') }}</span>
            <input
              v-model="selectAll.game"
              :disabled="!payloadList.agent.length && showAgentSelect"
              type="checkbox"
              @change="(e) => toggleSelectAll('game', e.target.checked)"
            />
            <span>{{ t('textSelectAllOrNot') }}</span>
          </div>
          <div class="flex h-[250px] w-[200px] flex-col overflow-y-scroll border px-1">
            <div v-if="!showAgentSelect || payloadList.agent.length !== 0" class="flex flex-col">
              <label v-for="game in dataList.game" :key="`selectGameIndex__${game.id}`">
                <input v-model="game.selected" type="checkbox" class="mr-1" @change="toggleSelect('game', game)" />
                <span>{{ game.name }}</span>
              </label>
            </div>
          </div>
        </div>

        <ArrowRightIcon class="h-14 w-14 text-gray-300" />

        <div class="flex flex-col">
          <div class="space-x-4">
            <span>{{ t('textSelectRoom') }}</span>
            <input
              v-model="selectAll.room"
              :disabled="!payloadList.game.length"
              type="checkbox"
              @change="(e) => toggleSelectAll('room', e.target.checked)"
            />
            <span>{{ t('textSelectAllOrNot') }}</span>
          </div>
          <div class="flex h-[250px] w-[200px] flex-col overflow-y-scroll border px-1">
            <div v-if="payloadList.game.length !== 0" class="flex flex-col">
              <label v-for="room in dataList.room" :key="`selectGameIndex__${room.id}`">
                <input
                  v-model="room.selected"
                  type="checkbox"
                  class="mr-1"
                  :disabled="room.disabled"
                  @change="toggleSelect('room', room)"
                />
                <span>{{ room.name }}</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <hr class="my-4" />

      <div class="flex justify-between px-4">
        <div class="flex items-center">
          <span class="min-w-fit">{{ t('textRTP') }}</span>
          <input v-model="payloadList.killRatio" max="50" min="0" step="0.1" class="form-input mx-1" type="number" />
          <span>%</span>
        </div>

        <div class="flex items-center">
          <span class="min-w-fit">{{ t('textNewbieKillRatio') }}</span>
          <input
            v-model="payloadList.newKillRatio"
            :disabled="props.gameType === constant.GameType.BaiRen"
            max="100"
            min="0"
            step="0.1"
            class="form-input mx-1"
            type="number"
          />
          <span>%</span>
        </div>

        <div class="flex items-center">
          <span class="min-w-fit">{{ t('textActiveNumber') }}</span>
          <input
            v-model="payloadList.activeNumber"
            :disabled="props.gameType !== constant.GameType.BaiRen"
            max="999"
            min="0"
            step="1"
            class="form-input mx-1"
            type="number"
          />
        </div>
      </div>

      <div class="px-4">
        <span>{{ t('textRemark') }}</span>
        <input v-model="payloadList.remark" type="text" class="w-full border p-2" />
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close()">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="payloadList"
          :button-click="() => setIncomeRatio()"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { reactive, ref, watch, inject, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useDropdownListStore } from '@/base/store/dropdownStore'
import { storeToRefs } from 'pinia'
import { round } from '@/base/utils/math'
import { roomTypeNameIndex } from '@/base/utils/room'
import { ArrowRightIcon } from '@heroicons/vue/20/solid'
import constant from '@/base/common/constant'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'
import axios from 'axios'
import * as api from '@/base/api/sysRiskControl'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  gameType: {
    type: Number,
    default: 0,
  },
})

const emit = defineEmits(['close'])
const warn = inject('warn')
const { t } = useI18n()
const { dropdownList } = storeToRefs(useDropdownListStore())

const showAgentSelect = computed(() => props.gameType !== constant.GameType.BaiRen)

const visible = ref(false)

const dataList = reactive({
  agent: [],
  game: [],
  room: [],
})

const selectAll = reactive({
  agent: false,
  game: false,
  room: false,
})

const payloadList = reactive({
  agent: [],
  game: [],
  room: [],
  killRatio: 0,
  newKillRatio: 0,
  activeNumber: 0,
  remark: '',
})

function toggleSelectAll(target, selected) {
  const payload = []
  dataList[target].forEach((data) => {
    data.selected = selected
    if (selected) {
      payload.push(data.id)
    }
  })
  payloadList[target] = payload
}

function toggleSelect(target, data) {
  if (data.selected) {
    payloadList[target].push(data.id)
    selectAll[target] = dataList[target].every((data) => data.selected)
  } else {
    payloadList[target] = payloadList[target].filter((id) => id !== data.id)
    selectAll[target] = false
  }
}

function validateForm(payload) {
  const errors = []

  if (props.gameType !== constant.GameType.BaiRen && payload.agent.length === 0) {
    errors.push(t('textSetIncomeRatiosAgentListError'))
  }

  if (payload.game.length === 0) {
    errors.push(t('textSetIncomeRatiosGameListError'))
  }

  if (payload.room.length === 0) {
    errors.push(t('textSetIncomeRatiosRoomListError'))
  }

  if (!Number.isInteger(payload.killRatio) || payload.killRatio < 0 || payload.killRatio > 100) {
    errors.push(t('textRTPRangeError'))
  }

  if (!Number.isInteger(payload.newKillRatio) || payload.newKillRatio < 0 || payload.newKillRatio > 100) {
    errors.push(t('textNewKillRatioRangeError'))
  }

  if (!Number.isInteger(payload.activeNumber) || payload.activeNumber < 0 || payload.activeNumber > 999) {
    errors.push(t('textActiveNumRangeError'))
  }

  return errors
}

async function setIncomeRatio() {
  const errors = validateForm(payloadList)
  if (errors.length > 0) {
    warn(errors.join(t('symbolComma')))
    return
  }

  try {
    const agentList = props.gameType === constant.GameType.BaiRen ? [-1] : payloadList.agent

    const resp = await api.setIncomeRatios({
      active_num: payloadList.activeNumber,
      agent_id_list: agentList,
      game_id_list: payloadList.game,
      info: payloadList.remark,
      kill_ratio: round(payloadList.killRatio / 100, 4),
      new_kill_ratio: round(payloadList.newKillRatio / 100, 4),
      room_type_list: payloadList.room,
    })

    let errorDetail = ''
    if (resp.data.code === constant.ErrorCode.ErrorIncomeRatioData) {
      const errorList = resp.data.data

      for (let i = 0; i < errorList.length; i++) {
        const errorTitle =
          props.gameType === constant.GameType.BaiRen
            ? t(`gameType__${props.gameType}`)
            : t('fmtTextAgentSettingAgentId', [errorList[i].agent_id])
        const roomType = t(`roomType__${roomTypeNameIndex(errorList[i].room_type, errorList[i].room_type)}`)

        errorDetail += `${t(`fmtTextSetIncomeRatiosError`, [
          errorTitle,
          t(`game__${errorList[i].game_id}`),
          roomType,
        ])}\n`
      }
    }

    const warnMessage = errorDetail !== '' ? errorDetail : t(`errorCode__${resp.data.code}`)
    warn(warnMessage).then(() => {
      if (resp.data.code === constant.ErrorCode.Success) {
        close()
      }
    })
  } catch (err) {
    console.error(err)

    if (axios.isAxiosError(err)) {
      const errorCode = err.response.data
        ? err.response.data.code || constant.ErrorCode.ErrorLocal
        : constant.ErrorCode.ErrorLocal
      warn(t(`errorCode__${errorCode}`))
    }
  }
}

function close() {
  Object.assign(payloadList, {
    agent: [],
    game: [],
    room: [],
    killRatio: 0,
    newKillRatio: 0,
    activeNumber: 0,
    remark: '',
  })

  for (const list of Object.values(dataList)) {
    if (list.length) {
      for (let i = 0; i < list.length; i++) {
        list[i].selected = false
      }
    }
  }

  for (const key of Object.keys(selectAll)) {
    selectAll[key] = false
  }

  emit('close')
}

watch(
  () => props.visible,
  () => {
    visible.value = props.visible
  }
)

watch(
  () => dropdownList.value.agents,
  () => {
    let selectAllAgent = true
    let agentPayload = []
    const agentDataList = dropdownList.value.agents.map((agent) => {
      const prevAgent = dataList.agent.find((prevAgent) => prevAgent.id === agent.id)
      if (prevAgent && prevAgent.selected) {
        agentPayload.push(prevAgent.id)
        selectAllAgent = selectAllAgent && prevAgent.selected
      } else {
        selectAllAgent = false
      }
      return prevAgent
        ? prevAgent
        : {
            id: agent.id,
            name: agent.name,
            selected: false,
          }
    })
    agentDataList.sort((a, b) => a.id - b.id)

    dataList.agent = agentDataList
    payloadList.agent = agentDataList.filter((agent) => agent.selected).map((agent) => agent.id)
    selectAll.agent = agentDataList.every((agent) => agent.selected)
  }
)

watch(
  () => dropdownList.value.games,
  () => {
    let selectAllGame = true
    let gamePayload = []
    const gameDataList = dropdownList.value.games
      .filter((game) => Math.floor(round(game.id / 1000, 7)) === props.gameType)
      .map((game) => {
        const prevGame = dataList.game.find((prevGame) => prevGame.id === game.id)
        if (prevGame && prevGame.selected) {
          gamePayload.push(prevGame.id)
          selectAllGame = selectAllGame && prevGame.selected
        } else {
          selectAllGame = false
        }
        return prevGame
          ? prevGame
          : {
              id: game.id,
              name: t(`game__${game.id}`),
              selected: false,
            }
      })
    gameDataList.sort((a, b) => a.id - b.id)

    dataList.game = gameDataList
    payloadList.game = gamePayload
    selectAll.game = selectAllGame
  }
)

// roomType 1~4 : 新手房 ~ 大師房 ; 5~8 : 初級場~至尊場 非每款遊戲都有此類型房間
watch(
  () => payloadList.game,
  () => {
    // 先設定每個遊戲類型的最大房間數，後續再檢查是否需要過濾
    let roomList
    switch (props.gameType) {
      case constant.GameType.ChiPai:
        roomList = [
          constant.RoomType.Newbie,
          constant.RoomType.Common,
          constant.RoomType.Master,
          constant.RoomType.GrandMaster,
          constant.RoomType.Beginner,
          constant.RoomType.SmallSuccess,
          constant.RoomType.GreateFacility,
          constant.RoomType.Perfection,
        ]
        break
      case constant.GameType.Slot:
        roomList = [constant.RoomType.Newbie]
        break
      default:
        roomList = [
          constant.RoomType.Newbie,
          constant.RoomType.Common,
          constant.RoomType.Master,
          constant.RoomType.GrandMaster,
        ]
        break
    }

    let filterRoomList = roomList
    switch (props.gameType) {
      case constant.GameType.ChiPai:
        if (
          payloadList.game.length !==
          payloadList.game.filter((game) => game === constant.Game.Texas || game === constant.Game.Chinesepoker).length
        ) {
          filterRoomList = filterRoomList.filter((roomType) => roomType <= constant.RoomType.GrandMaster)
        }
        break
      case constant.GameType.ElectronicGame:
        if (
          payloadList.game.includes(constant.Game.Rcfishing) ||
          payloadList.game.includes(constant.Game.Happyfishing)
        ) {
          filterRoomList = filterRoomList.filter((roomType) => roomType <= constant.RoomType.Master)
        }
        break
    }

    let selectAllRoom = true
    let roomPayload = []
    const roomDataList = filterRoomList.map((roomType) => {
      const prevRoom = dataList.room.find((room) => room.id === roomType)
      if (prevRoom && prevRoom.selected) {
        roomPayload.push(prevRoom.id)
        selectAllRoom = selectAllRoom && prevRoom.selected
      } else {
        selectAllRoom = false
      }
      return prevRoom
        ? prevRoom
        : {
            id: roomType,
            name: t(`roomType__${roomTypeNameIndex(props.gameType * 1000, roomType)}`),
            selected: false,
          }
    })
    roomDataList.sort((a, b) => a.id - b.id)

    dataList.room = roomDataList
    payloadList.room = roomPayload
    selectAll.room = selectAllRoom
  },
  { deep: true }
)
</script>
