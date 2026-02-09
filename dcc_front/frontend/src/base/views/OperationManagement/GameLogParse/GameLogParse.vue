<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show" @submit.prevent>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textRoundId') }}</label>
          <input
            v-model="formInput.logNumber"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextReportRoundId')"
          />
        </div>

        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button type="button" class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12" @click="search()">
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
    </template>
  </ToggleHeader>

  <div class="mt-4 rounded bg-white p-4">
    <div class="rounded border p-4">
      <div class="flex flex-wrap items-center">
        <label class="form-label mb-1 w-full pr-2 md:mb-4 md:w-1/12 md:text-right">{{ t('textGameName') }}</label>
        <span class="form-text mb-1 md:mb-4 md:w-3/12">{{ gameLog.gameName }}</span>
        <label class="form-label mb-1 w-full pr-2 md:mb-4 md:w-1/12 md:text-right">{{ t('textRoomType') }}</label>
        <span class="form-text mb-1 md:mb-4 md:w-3/12">{{ gameLog.roomTypeName }}</span>
        <label class="form-label mb-1 w-full pr-2 md:mb-4 md:w-1/12 md:text-right">{{ t('textRoundId') }}</label>
        <span class="form-text mb-1 md:mb-4 md:w-3/12">{{ gameLog.logNumber }}</span>
        <label class="form-label mb-1 w-full pr-2 md:mb-4 md:w-1/12 md:text-right">{{ t('textCreateTime') }}</label>
        <span class="form-text mb-1 md:mb-4 md:w-3/12">
          {{ gameLog.betTime ? time.utcTimeStrToLocalTimeFormat(gameLog.betTime) : gameLog.betTime }}
        </span>
      </div>
      <div class="min-h-[200px] rounded border p-4 text-gray-500">
        <component :is="playLogComponent" v-if="playLogComponent" :play-log="gameLog.playLog" />
        <div v-else class="text-center">{{ t('textTableEmpty') }}</div>
      </div>
    </div>
  </div>
  <PageTableLoading v-show="formInput.showProcessing" />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useGameLogParse } from '@/base/composable/operationManagement/gameLogParse/useGameLogParse'
import time from '@/base/utils/time'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'

const props = defineProps({
  logNumber: {
    type: String,
    default: '',
  },
  userName: {
    type: String,
    default: '',
  },
})

const { t } = useI18n()
const { formInput, gameLog, playLogComponent, search } = useGameLogParse(props)
</script>
