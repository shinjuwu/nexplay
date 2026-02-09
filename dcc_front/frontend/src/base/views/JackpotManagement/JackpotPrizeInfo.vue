<template>
  <div class="bg-white p-4">
    <ToggleHeader :tips-title="t('textJackpotPrizeInfoDirections')">
      <template #tipsSlot="{ tipsShow }">
        <div
          v-for="(direction, dirIdx) in pageDirections"
          v-show="tipsShow"
          :key="`directions__${dirIdx}`"
          class="text-danger"
        >
          {{ direction }}
        </div>
      </template>
    </ToggleHeader>
    <hr class="my-2" />
    <div class="flex justify-evenly max-md:flex-col">
      <div class="flex items-center justify-between border border-gray-200 p-4 md:h-[300px] md:w-[300px] md:flex-col">
        <p class="text-xl font-bold">{{ t('textAnnouncementPrize') }}</p>
        <p class="text-3xl font-bold text-blue-600">{{ `$${numberToStr(showPool, 4)}` }}</p>
        <ArrowUpCircleIcon class="inline-block w-16 text-blue-600" />
      </div>
      <div class="flex items-center justify-between border border-gray-200 p-4 md:h-[300px] md:w-[300px] md:flex-col">
        <span class="text-xl font-bold">{{ t('textRealPrize') }}</span>
        <p class="text-3xl font-bold text-rose-400">{{ `$${numberToStr(realPool, 4)}` }}</p>
        <ArrowDownCircleIcon class="inline-block w-16 text-rose-400" />
      </div>
      <div class="flex items-center justify-between border border-gray-200 p-4 md:h-[300px] md:w-[300px] md:flex-col">
        <span class="text-xl font-bold">{{ t('textPrepPrize') }}</span>
        <p class="text-3xl font-bold text-emerald-400">{{ `$${numberToStr(reservePool, 4)}` }}</p>
        <ArrowsRightLeftIcon class="inline-block w-16 text-emerald-400" />
      </div>
    </div>
    <div class="text-right">
      <button type="button" class="btn btn-light" @click="getJackpotPoolData()">
        {{ t('textRefresh') }}
      </button>
    </div>
  </div>
  <PageTableLoading v-show="showPrizeInfoProcessing" />
</template>
<script setup>
import { useI18n } from 'vue-i18n'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTableLoading from '@/base/components/Page/Table/PageTableLoading.vue'
import { useJackpotPrizeInfo } from '@/base/composable/jackpotManagement/useJackpotPrizeInfo'
import { ArrowUpCircleIcon, ArrowDownCircleIcon, ArrowsRightLeftIcon } from '@heroicons/vue/20/solid'
import { numberToStr } from '@/base/utils/formatNumber'
const { t } = useI18n()
const { pageDirections, showPool, realPool, reservePool, showPrizeInfoProcessing, getJackpotPoolData } =
  useJackpotPrizeInfo()
</script>
