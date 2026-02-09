<template>
  <PageTableDialog :visible="visible" @close="close()">
    <template #header>{{ t('textMarkPlayer') }}</template>
    <template #default>
      <div class="mb-5">{{ t('fmtTextUserName', [playerInfo.userName]) }}</div>
      <div class="grid grid-flow-row grid-cols-6 gap-2">
        <template v-for="badge in defaultBadges.items" :key="`defaultBadge__${badge.index}`">
          <div class="col-span-2 flex items-center">
            <span
              v-if="badge.index === -3"
              class="min-w-fit cursor-pointer rounded-full bg-gray-400 px-1"
              :style="{
                backgroundColor: playerInfo.killDiveState === 2 ? badge.bgColor : '#9CA3AF',
                color: playerInfo.killDiveState === 2 ? badge.txtColor : '#000000',
              }"
              @click="selectBadge(badge)"
            >
              {{ badge.name }}
            </span>
            <span
              v-if="badge.index === -2"
              class="min-w-fit cursor-pointer rounded-full px-1"
              :style="{
                backgroundColor: playerInfo.highRisk ? badge.bgColor : '#9CA3AF',
                color: playerInfo.highRisk ? badge.txtColor : '#000000',
              }"
              @click="selectBadge(badge)"
            >
              {{ badge.name }}
            </span>
            <span
              v-if="badge.index === -1"
              class="min-w-fit cursor-pointer rounded-full px-1"
              :style="{
                backgroundColor: playerInfo.killDiveState === 1 ? badge.bgColor : '#9CA3AF',
                color: playerInfo.killDiveState === 1 ? badge.txtColor : '#000000',
              }"
              @click="selectBadge(badge)"
            >
              {{ badge.name }}
            </span>
            <input
              v-if="badge.index === -1"
              v-model="playerInfo.killDiveValue"
              :disabled="!(playerInfo.killDiveState === 1 || badge.isSelected)"
              min="0"
              type="number"
              class="form-input ml-1"
            />
          </div>
        </template>
      </div>
      <hr class="my-2" />
      <div v-if="playerInfo.tagList" class="grid grid-flow-row grid-cols-4 gap-2">
        <template v-for="badge in filterCustomBadges" :key="`dialog_customBadge__${badge.index}`">
          <div class="col-span-1">
            <span
              class="min-w-fit cursor-pointer rounded-full bg-gray-400 px-1"
              :style="{
                backgroundColor: playerInfo.tagList[badge.index] !== '0' ? badge.bgColor : '#9CA3AF',
                color: playerInfo.tagList[badge.index] !== '0' ? badge.txtColor : '#000000',
              }"
              @click="selectCustomBadge(badge)"
            >
              {{ badge.name }}
            </span>
          </div>
        </template>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close()">
          {{ t('textClose') }}
        </button>
        <LoadingButton
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="playerInfo"
          :button-click="() => setGameUserCustomTag(playerInfo)"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>
<script setup>
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import { inject, ref, reactive, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import axios from 'axios'
import * as api from '@/base/api/sysRiskControl'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

class Badges {
  constructor(name, bgColor, txtColor, index) {
    this.name = name
    this.bgColor = bgColor
    this.txtColor = txtColor
    this.index = index
  }
}

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
const emit = defineEmits(['close', 'searchGameUsers'])

const { t } = useI18n()
const warn = inject('warn')

const defaultBadges = reactive({
  items: [
    new Badges(t('riskType__1'), '#000000', '#ffffff', -3),
    new Badges(t('riskType__2'), '#fb7185', '#ffffff', -2),
    new Badges(t('riskType__3'), '#fbbf24', '#ffffff', -1),
  ],
})
const customBadges = reactive({
  items: [
    new Badges('', '', '', 0),
    new Badges('', '', '', 1),
    new Badges('', '', '', 2),
    new Badges('', '', '', 3),
    new Badges('', '', '', 4),
    new Badges('', '', '', 5),
    new Badges('', '', '', 6),
    new Badges('', '', '', 7),
  ],
})

const filterCustomBadges = computed(() => customBadges.items.filter((badge) => badge.name !== ''))

const visible = ref(false)

const playerInfo = reactive({
  id: props.playerInfo.id,
  userName: props.playerInfo.userName,
  highRisk: '',
  killDiveState: 0,
  killDiveValue: 0,
  tagList: [],
})

function selectBadge(badge) {
  switch (badge.name) {
    case t('riskType__1'):
      playerInfo.killDiveState = playerInfo.killDiveState === 2 ? 0 : 2
      playerInfo.killDiveValue = 0
      break
    case t('riskType__2'):
      playerInfo.highRisk = !playerInfo.highRisk
      break
    case t('riskType__3'):
      playerInfo.killDiveState = playerInfo.killDiveState === 1 ? 0 : 1
      break
  }
}

function selectCustomBadge(badge) {
  const idx = badge.index
  playerInfo.tagList[idx] = playerInfo.tagList[idx] === '0' ? '1' : '0'
}

async function getGameUserTag() {
  try {
    const resp = await api.getGameUsersCustomTag({
      game_users_id: props.playerInfo.id,
      game_users_name: props.playerInfo.userName,
    })

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`)).then(() => emit('close', false))
      return
    }

    Object.values(resp.data.data.custom_tag_info).forEach((tagInfo) => {
      const customBadge = customBadges.items[tagInfo.idx]
      customBadge.name = tagInfo.name
      customBadge.info = tagInfo.info

      if (tagInfo.name !== '') {
        const colorObj = JSON.parse(tagInfo.color)
        customBadge.bgColor = colorObj.bg_color
        customBadge.txtColor = colorObj.txt_color
      }
    })

    const data = resp.data.data
    playerInfo.id = data.game_users_id
    playerInfo.userName = data.game_users_name
    playerInfo.highRisk = data.high_risk
    playerInfo.killDiveState = data.kill_dive_state
    playerInfo.killDiveValue = data.kill_dive_value
    playerInfo.tagList = [...data.tag_list]

    visible.value = true
  } catch (err) {
    console.error(err)

    if (axios.isAxiosError(err)) {
      const errorCode = err.response.data
        ? err.response.data.code || constant.ErrorCode.ErrorLocal
        : constant.ErrorCode.ErrorLocal
      warn(t(`errorCode__${errorCode}`)).then(() => emit('close', false))
    }
  }
}

async function setGameUserCustomTag(info) {
  if (info.killDiveState === 1 && !info.killDiveValue) {
    warn(t('textKillDiveValueError'))
    return
  }

  try {
    const resp = await api.setGameUsersCustomTag({
      custom_tag_info: customBadges.items.reduce((resp, badge) => {
        resp[badge.index.toString()] = {
          color: badge.name !== '' ? `{"bg_color":"${badge.bgColor}","txt_color":"${badge.txtColor}"}` : '',
          idx: badge.index,
          info: badge.name !== '' ? badge.info : '',
          name: badge.name,
        }
        return resp
      }, {}),
      game_users_id: info.id,
      game_username: info.userName,
      high_risk: info.highRisk,
      kill_dive_state: info.killDiveState,
      kill_dive_value: info.killDiveValue,
      tag_list: info.tagList.join(''),
    })

    warn(t(`errorCode__${resp.data.code}`)).then(() => {
      if (resp.data.code !== constant.ErrorCode.Success) {
        return
      }

      close()
      emit('searchGameUsers')
    })
  } catch (err) {
    console.error(err)

    if (axios.isAxiosError(err)) {
      const errorCode = err.response.data
        ? err.response.data.code || constant.ErrorCode.ErrorLocal
        : constant.ErrorCode.ErrorLocal
      warn(t(`errorCode__${errorCode}`)).then(() => emit('close', false))
    }
  }
}

function close() {
  visible.value = false
  emit('close', false)
}

watch(
  () => props.visible,
  async (newValue) => {
    if (newValue) {
      await getGameUserTag()
    } else {
      visible.value = false
    }
  }
)
</script>
