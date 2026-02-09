<template>
  <PageTableDialog :visible="visible" @close="close()">
    <template #header>{{ t('roleItemResetPassword') }}</template>
    <template #default>
      <div class="flex flex-col px-3">
        <div v-if="!isReset">
          <span>{{ t('fmtTextResetPassword', [props.userName]) }}</span>
          <ul class="mt-3">
            <li class="text-danger ml-5 list-disc">{{ t('fmtTextConfirmResetPassword') }}</li>
          </ul>
        </div>
        <div v-else>
          <span>{{ t('textResetSuccess') }}</span>
          <ul class="mt-3">
            <li class="text-danger ml-5 list-disc">
              {{ t('fmtTextNewPassword', [newPassword]) }}
            </li>
            <li class="text-danger ml-5 list-disc">
              {{ reminder }}
            </li>
          </ul>
        </div>
      </div>
    </template>
    <template #footer>
      <div v-if="!isReset" class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close()">
          {{ t('textClose') }}
        </button>
        <button class="btn btn-primary" @click="resetPassword()">
          {{ t('textConfirm') }}
        </button>
      </div>
      <div v-else class="ml-auto">
        <button type="button" class="btn btn-primary" @click="close()">{{ t('textConfirm') }}</button>
      </div>
    </template>
  </PageTableDialog>
</template>
<script setup>
import { inject, ref, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  userName: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    required: true,
  },
})
const emit = defineEmits(['close'])

const { t } = useI18n()
const warn = inject('warn')

const visible = ref(false)
const isReset = ref(false)
const newPassword = ref('')

const reminder = computed(() => {
  const status = props.type === 'AgentAccount' ? t('textAgent') : t('textBackendPersonnel')

  return t('fmtResetPasswordReminder', [status, props.userName])
})

function close() {
  visible.value = false
  isReset.value = false
  emit('close', false)
}

async function resetPassword() {
  try {
    const resp = await api.resetPassword({
      username: props.userName,
    })

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`)).then(() => emit('close', false))
      return
    }

    newPassword.value = resp.data.data.new_password
    isReset.value = true
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
  () => {
    if (props.visible) {
      visible.value = props.visible
    } else {
      visible.value = false
    }
  }
)
</script>
