<template>
  <PageTableDialog :visible="visible" @close="close()">
    <template #header>{{ t('textDisposeSetting') }}</template>
    <template #default>
      <div class="grid grid-cols-6 gap-4">
        <div class="col-span-2">{{ t('fmtTextUserName', [playerInfo.userName]) }}</div>
        <div class="col-span-2">{{ t('fmtTextAgentName', [playerInfo.agentName]) }}</div>
      </div>
      <div class="my-5 flex w-1/3 items-center justify-between">
        <button
          v-for="(tag, tagIndex) in riskTag"
          :key="tagIndex"
          :class="{ 'btn-danger': playerInfo.riskControlTag[tagIndex] === '1' }"
          class="btn btn-secondary mx-2 rounded-full"
          @click="setDispose(tagIndex)"
        >
          {{ tag }}
        </button>
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
          :button-click="() => setGameUserRiskControlTag()"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>
<script setup>
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import { inject, ref, reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import axios from 'axios'
import * as api from '@/base/api/sysRiskControl'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

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

const visible = ref(false)

const playerInfo = reactive({
  id: props.playerInfo.id,
  userName: props.playerInfo.userName,
  riskControlTag: [],
})

const riskTag = [
  t('riskControlTag__1000'),
  t('riskControlTag__0100'),
  t('riskControlTag__0010'),
  t('riskControlTag__0001'),
]
function setDispose(targetIndex) {
  playerInfo.riskControlTag[targetIndex] = playerInfo.riskControlTag[targetIndex] === '0' ? '1' : '0'
}

function close() {
  visible.value = false
  emit('close', false)
}

async function getGameUserRiskControlTag() {
  try {
    const resp = await api.getGameUserRiskControlTag({ id: props.playerInfo.id })

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`)).then(() => emit('close', false))
      return
    }

    const data = resp.data.data
    playerInfo.id = props.playerInfo.id
    playerInfo.userName = props.playerInfo.userName
    playerInfo.agentName = props.playerInfo.agentName
    playerInfo.riskControlTag = [...data.risk_control_tag]

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

async function setGameUserRiskControlTag() {
  try {
    const resp = await api.setGameUserRiskControlTag({
      id: playerInfo.id,
      risk_control_tag: playerInfo.riskControlTag.join(''),
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

watch(
  () => props.visible,
  async (newValue) => {
    if (newValue) {
      await getGameUserRiskControlTag()
    } else {
      visible.value = false
    }
  }
)
</script>
