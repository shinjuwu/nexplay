<template>
  <PageTableDialog :visible="visible" @close="close()">
    <template #header>{{ t('textMemberAccountSetting') }}</template>
    <template #default>
      <div class="space-y-4">
        <div>
          <label class="form-label">{{ t('fmtTextUserName', [playerInfo.userName]) }}</label>
        </div>
        <div>
          <label class="form-label"> {{ t('textState') }} </label>
          <div>
            <select v-model="playerInfo.state" class="form-input md:w-1/2">
              <option
                v-for="option in userTypeOptions"
                :key="option"
                :value="option"
                :label="t(`status__${option}`)"
              ></option>
            </select>
          </div>
        </div>
        <div>
          <label class="form-label">
            {{ t('textRemark') }}
          </label>
          <textarea v-model="playerInfo.remark" class="form-input" type="text" maxlength="100" />
        </div>
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
          :parent-data="playerInfo"
          :button-click="() => submit()"
        >
          {{ t('textOnSave') }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { reactive, ref, watch, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import axios from 'axios'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
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
  id: 0,
  userName: '',
  remark: '',
  state: 0,
})

const userTypeOptions = [constant.AccountStatus.Open, constant.AccountStatus.Disable]

async function submit() {
  try {
    const resp = await api.updateGameUserInfo({
      id: playerInfo.id,
      info: playerInfo.remark,
      is_enabled: playerInfo.state === constant.AccountStatus.Open,
    })

    warn(t(`errorCode__${resp.data.code}`)).then(() => {
      if (resp.data.code !== constant.ErrorCode.Success) {
        return
      }
      emit('close', false)
      emit('searchGameUsers')
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
  visible.value = false
  emit('close', false)
}

async function searchGameUsersInfo() {
  try {
    const resp = await api.getGameUsersInfo({ id: props.playerInfo.id, time_zone: new Date().getTimezoneOffset() })

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`)).then(() => emit('close', false))
      return
    }

    const data = resp.data.data
    playerInfo.id = props.playerInfo.id
    playerInfo.userName = props.playerInfo.userName
    playerInfo.remark = data.info
    playerInfo.state = data.is_enabled ? constant.AccountStatus.Open : constant.AccountStatus.Disable

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

watch(
  () => props.visible,
  async (newValue) => {
    if (newValue) {
      await searchGameUsersInfo()
    } else {
      visible.value = false
    }
  }
)
</script>
