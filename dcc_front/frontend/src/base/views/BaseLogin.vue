<template>
  <div class="bg-login flex flex-1 items-center bg-gray-50 bg-cover bg-center pt-3 pb-9 sm:py-20">
    <div class="container mx-auto px-3">
      <div class="flex justify-center">
        <div class="l:w-6/12 w-full px-3 md:w-8/12 xl:w-5/12 2xl:w-4/12">
          <div class="flex flex-col items-center justify-center rounded-t bg-indigo-400 py-4 text-center text-lg">
            <!-- login logo removed -->
          </div>
          <div class="rounded-b bg-white p-9">
            <form @keyup.enter="submit()">
              <div class="mb-6">
                <label class="form-label mb-2 inline-block">{{ t('textUserName') }}</label>
                <input
                  v-model.trim="userName"
                  type="text"
                  class="form-input"
                  :placeholder="t('placeHolderTextUserName')"
                />
              </div>
              <div class="mb-6">
                <label class="form-label mb-2 inline-block">{{ t('textPassword') }}</label>
                <FormGroupInput
                  :model-value="password"
                  :placeholder="t('placeHolderTextPassword')"
                  :type="showPassword ? 'text' : 'password'"
                  @update:model-value="(newValue) => (password = newValue)"
                  @click-icon="showPassword = !showPassword"
                >
                  <template #icon>
                    <EyeSlashIcon v-if="showPassword" class="h-4 w-4" />
                    <EyeIcon v-else class="h-4 w-4" />
                  </template>
                </FormGroupInput>
              </div>
              <div>
                <label class="form-label mb-2 inline-block">{{ t('textCaptcha') }}</label>
                <FormGroupInput
                  class="mb-2"
                  :model-value="graphicsCaptcha.captcha"
                  :placeholder="t('placeHolderTextCaptcha')"
                  @update:model-value="(newValue) => (graphicsCaptcha.captcha = newValue)"
                  @click-icon="updateGraphicsCaptcha()"
                >
                  <template #icon>
                    <ArrowPathIcon class="h-4 w-4" />
                  </template>
                </FormGroupInput>
                <img
                  class="w-full cursor-pointer rounded border border-gray-200"
                  :src="graphicsCaptcha.imageBase64"
                  @click="updateGraphicsCaptcha()"
                />
              </div>
              <div class="text-danger my-3 text-center">{{ errorMessage }}</div>
              <div>
                <button
                  class="box-shadiw w-full rounded bg-indigo-400 py-2 px-3.5 text-white shadow-md shadow-indigo-500/50 hover:bg-indigo-500"
                  type="button"
                  @click.prevent="submit()"
                >
                  {{ t('textLogin') }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { ArrowPathIcon, EyeIcon, EyeSlashIcon } from '@heroicons/vue/20/solid'

import { useLogin } from '@/base/composable/useLogin'
import FormGroupInput from '@/base/components/Form/FormGroupInput.vue'

const { t } = useI18n()
const { userName, password, graphicsCaptcha, errorMessage, showPassword, updateGraphicsCaptcha, submit } = useLogin()
</script>
