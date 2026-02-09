<script setup lang="ts">
import { computed } from 'vue'

import { round } from '@/utils/math'

const props = defineProps({
  name: {
    type: String,
    default: '',
  },
  deScore: {
    type: Number,
    default: 0,
  },
  yaScore: {
    type: Number,
    default: 0,
  },
  bonus: {
    type: Number,
    default: 0,
  },
  tax: {
    type: Number,
    default: 0,
  },
  rtp: {
    type: Number,
    default: 0,
  },
  playCount: {
    type: Number,
    default: 0,
  },
})

const rtpPercentage = computed(() => props.rtp * 100)

const needleRotateDegreed = computed(() => {
  const startDeg = -122
  const endDeg = 122
  const maxRtp = 240

  let rtp = rtpPercentage.value
  if (rtp > maxRtp) {
    rtp = maxRtp
  }

  return round(startDeg + (endDeg - startDeg) * (rtp / maxRtp), 2)
})
</script>

<template>
  <div>
    <div class="rounded-t border-4 border-b-0 text-center text-xl">{{ props.name }}</div>
    <div
      class="before: relative border-4 border-b-0 text-center text-base before:absolute before:-left-2 before:-top-2 before:h-5 before:w-5 before:rounded-full before:bg-slate-400 before:text-xs before:leading-5 before:text-white before:content-['场']"
    >
      {{ playCount }}
    </div>
    <div
      class="before: relative border-4 border-b-0 text-center text-base before:absolute before:-left-2 before:-top-2 before:h-5 before:w-5 before:rounded-full before:bg-slate-400 before:text-xs before:leading-5 before:text-white before:content-['得']"
    >
      {{ deScore.toFixed(2) }}
    </div>
    <div
      class="before: relative border-4 border-b-0 text-center text-base before:absolute before:-left-2 before:-top-2 before:h-5 before:w-5 before:rounded-full before:bg-slate-400 before:text-xs before:leading-5 before:text-white before:content-['押']"
    >
      {{ yaScore.toFixed(2) }}
    </div>
    <div
      class="before: relative border-4 border-b-0 text-center text-base before:absolute before:-left-2 before:-top-2 before:h-5 before:w-5 before:rounded-full before:bg-slate-400 before:text-xs before:leading-5 before:text-white before:content-['rtp']"
    >
      {{ rtpPercentage.toFixed(2) }}%
    </div>
    <div class="relative w-full overflow-hidden rounded-b border-4 pt-[calc(100%-8px)]">
      <div class="rtp-dashboard absolute left-0 top-0 w-full pt-[100%]"></div>
      <div
        class="absolute left-0 top-0 w-full pt-[100%]"
        :class="{
          'rtp-border-blue': rtpPercentage <= 80,
          'rtp-border-green': rtpPercentage > 80 && rtpPercentage <= 95,
          'rtp-border-orange': rtpPercentage > 95 && rtpPercentage <= 100,
          'rtp-border-red': rtpPercentage > 100 && rtpPercentage <= 200,
          'rtp-border-purple': rtpPercentage > 200,
        }"
      ></div>
      <div
        class="rtp-needle absolute left-0 top-0 w-full pt-[100%]"
        :style="`transform: rotate(${needleRotateDegreed}deg);`"
      ></div>
    </div>
  </div>
</template>
