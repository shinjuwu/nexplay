<template>
  <section class="rounded-b bg-white">
    <ToggleHeader>
      <template #default="{ show }">
        <form v-show="show">
          <div v-show="isAdminUser" class="mb-4 flex flex-wrap items-center">
            <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
            <FormAgentListDropdown
              v-model="formInput.agent"
              :include-all="false"
              :include-self="!isAdminUser"
              :include-grandson="false"
              class="mb-1 w-full md:w-3/12"
            />
          </div>

          <div class="flex items-center justify-end">
            <button
              type="button"
              class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
              @click="getCannedList()"
            >
              {{ t('textSearch') }}
            </button>
          </div>
        </form>
      </template>
    </ToggleHeader>

    <hr class="my-2 mx-4" />

    <div class="tbl-container p-4">
      <CannedLanguageList
        :canned-language-type="constant.CannedLanguageType.Custome"
        :canned-language-list="cannedList.items"
        :canned-language-list-agent-id="cannedListAgentId"
        :is-editable="isEditable"
      />
      <PageTableLoading v-show="showTableProcessing" />
    </div>

    <div class="flex items-center justify-end p-4">
      <button type="button" class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12" @click="setCannedList()">
        {{ t('textOnSave') }}
      </button>
    </div>
  </section>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useCustomeCannedLanguage } from '@/base/composable/gameSetting/gameCannedLanguage/useCustomeCannedLanguage'
import constant from '@/base/common/constant'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'
import CannedLanguageList from '@/base/views/GameSetting/GameCannedLanguage/CannedLanguageList.vue'

const { t } = useI18n()
const {
  isAdminUser,
  formInput,
  cannedListAgentId,
  showTableProcessing,
  isEditable,
  cannedList,
  getCannedList,
  setCannedList,
} = useCustomeCannedLanguage()
</script>
