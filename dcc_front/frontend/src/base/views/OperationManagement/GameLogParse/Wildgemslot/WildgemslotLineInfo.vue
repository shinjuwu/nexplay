<template>
  <table class="text-center">
    <thead>
      <tr class="bg-info text-white">
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineId') }}</th>
        <th class="w-auto border px-2 py-1">{{ t('textSlotLineSymbols') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineMulti') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotWildMulti') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineWins') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineFreeGameWins') }}</th>
      </tr>
    </thead>
    <tbody>
      <tr v-if="props.lines.length === 0">
        <td class="w-60 border px-2 py-1">-</td>
        <td class="border px-2 py-1">{{ t('textSlotNoLineSymbols') }}</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
      </tr>
      <tr v-for="line in props.lines" v-else :key="`line_${line.key}`">
        <td class="w-60 border px-2 py-1">{{ line.id }}</td>
        <td class="border px-2 py-1">
          <div class="flex justify-center gap-1">
            <WildgemslotPic
              v-for="(symbol, sIndex) in line.symbols"
              :key="`line__${line.id}_symbol__${symbol}__${sIndex}`"
              class="h-[57px] w-16"
              :symbol="symbol"
            />
          </div>
        </td>
        <td class="w-60 border px-2 py-1">{{ line.symbolMulti }}</td>
        <td class="w-60 border px-2 py-1">{{ line.wildMulti }}</td>
        <td class="w-60 border px-2 py-1">{{ line.wins }}</td>
        <td class="w-60 border px-2 py-1">{{ line.freeGameCount }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
import { useI18n } from 'vue-i18n'

import WildgemslotPic from '@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotPic.vue'

const { t } = useI18n()
const props = defineProps({
  lines: {
    type: Array,
    default: () => [],
  },
})
</script>
