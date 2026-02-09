<template>
  <section class="rounded-b bg-white">
    <div class="tbl-container p-4">
      <CannedLanguageList
        :canned-language-type="constant.CannedLanguageType.Default"
        :canned-language-list="cannedList.items"
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
import { useDefaultCannedLanguage } from '@/base/composable/gameSetting/gameCannedLanguage/useDefaultCannedLanguage'
import constant from '@/base/common/constant'

import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'
import CannedLanguageList from '@/base/views/GameSetting/GameCannedLanguage/CannedLanguageList.vue'

const { t } = useI18n()
const { showTableProcessing, isEditable, cannedList, setCannedList } = useDefaultCannedLanguage()
</script>
