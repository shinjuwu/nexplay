<template>
  <PageTableDialog :visible="props.visible" :has-footer="false" size="md" @close="emit('close')">
    <template #header>
      <div class="flex w-full justify-between">
        <span>{{ t('textMyAccount') }}</span>
        <XCircleIcon class="h-7 w-7 cursor-pointer" @click="emit('close')" />
      </div>
    </template>
    <template #default>
      <form>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-3/12 md:text-right" for="myName"
            >{{ t('textAccount') }}:</label
          >
          <input
            id="myName"
            class="form-input pointer-events-none mb-1 w-full cursor-default border-0 md:w-9/12"
            type="text"
            :value="myName"
          />
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-3/12 md:text-right" for="nickName"
            >{{ t('textNickName') }}:</label
          >
          <input
            id="nickName"
            v-model="myNickname"
            class="form-input mb-1 w-full md:w-9/12"
            type="text"
            maxlength="16"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center">
          <LoadingButton
            class="btn btn-primary ml-auto w-full md:w-3/12"
            :is-get-data="props.visible"
            :parent-data="myNickname"
            :button-click="() => setPersonalInfo()"
          >
            {{ t('textSubmit') }}
          </LoadingButton>
        </div>
      </form>
      <hr class="my-2" />
      <form>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-3/12 md:text-right" for="oldPassword"
            >{{ t('textOldPassword') }}:</label
          >
          <FormGroupInput
            id="oldPassword"
            class="mb-1 w-full md:w-9/12"
            :model-value="oldPassword"
            :placeholder="t('placeHolderTextOldPassword')"
            :type="showOldPassword ? 'text' : 'password'"
            @update:model-value="(newValue) => (oldPassword = newValue)"
            @click-icon="showOldPassword = !showOldPassword"
          >
            <template #icon>
              <EyeSlashIcon v-if="showOldPassword" class="h-4 w-4" />
              <EyeIcon v-else class="h-4 w-4" />
            </template>
          </FormGroupInput>
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-3/12 md:text-right" for="newPassword"
            >{{ t('textNewPassword') }}:</label
          >
          <FormGroupInput
            id="newPassword"
            class="mb-1 w-full md:w-9/12"
            :model-value="newPassword"
            :placeholder="t('placeHolderTextNewPassword')"
            :type="showNewPassword ? 'text' : 'password'"
            @update:model-value="(newValue) => (newPassword = newValue)"
            @click-icon="showNewPassword = !showNewPassword"
          >
            <template #icon>
              <EyeSlashIcon v-if="showNewPassword" class="h-4 w-4" />
              <EyeIcon v-else class="h-4 w-4" />
            </template>
          </FormGroupInput>
        </div>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-3/12 md:text-right" for="confirmPassword"
            >{{ t('textConfirmPassword') }}:</label
          >
          <FormGroupInput
            id="confirmPassword"
            class="mb-1 w-full md:w-9/12"
            :model-value="confirmPassword"
            :placeholder="t('placeHolderTextConfirmPassword')"
            :type="showConfirmPassword ? 'text' : 'password'"
            @update:model-value="(newValue) => (confirmPassword = newValue)"
            @click-icon="showConfirmPassword = !showConfirmPassword"
          >
            <template #icon>
              <EyeSlashIcon v-if="showConfirmPassword" class="h-4 w-4" />
              <EyeIcon v-else class="h-4 w-4" />
            </template>
          </FormGroupInput>
        </div>
        <div class="text-danger flex flex-wrap">
          <div class="hidden md:flex md:w-3/12"></div>
          <div class="flex w-full md:w-9/12">* {{ t('placeHolderTextFormPassword') }}</div>
          <div class="hidden md:flex md:w-3/12"></div>
          <div class="flex w-full md:w-9/12">* {{ t('textChangePasswordSuccessInfo') }}</div>
        </div>
        <div class="mt-4 flex flex-wrap items-center">
          <LoadingButton
            class="btn btn-primary ml-auto w-full md:w-3/12"
            :visible="props.visible"
            :parent-data="confirmPassword"
            :button-click="() => setPersonalPassword()"
          >
            {{ t('textSubmit') }}
          </LoadingButton>
          <button type="button"></button>
        </div>
      </form>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { EyeIcon, EyeSlashIcon, XCircleIcon } from '@heroicons/vue/24/outline'
import { usePersonalInfo } from '@/base/composable/layout/usePersonalInfo'

import FormGroupInput from '@/base/components/Form/FormGroupInput.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['close'])

const {
  confirmPassword,
  myName,
  myNickname,
  newPassword,
  oldPassword,
  showConfirmPassword,
  showNewPassword,
  showOldPassword,
  setPersonalInfo,
  setPersonalPassword,
} = usePersonalInfo()
</script>
