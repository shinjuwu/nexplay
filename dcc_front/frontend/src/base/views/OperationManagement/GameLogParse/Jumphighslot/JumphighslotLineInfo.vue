<template>
  <table class="text-center">
    <thead>
      <tr class="bg-info text-white">
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineId') }}</th>
        <th class="w-auto border px-2 py-1">{{ t('textSlotLineSymbols') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineWinWay') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineMulti') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotMulti') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineWins') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineFreeGameWins') }}</th>
      </tr>
    </thead>
    <tbody>
      <template v-if="props.lines.length === 0">
        <td class="w-60 border px-2 py-1">-</td>
        <td class="border px-2 py-1">{{ t('textSlotNoLineSymbols') }}</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
        <td class="w-60 border px-2 py-1">-</td>
      </template>
      <tr v-for="(line, index) in props.lines" :key="`line_${line.key}`">
        <td class="w-60 border px-2 py-1">{{ index + 1 }}</td>
        <td class="border px-2 py-1">
          <div class="flex -space-x-3.5">
            <JumphighslotPic
              v-for="(symbol, sIndex) in line.symbols"
              :key="`line__${index}_symbol__${symbol}__${sIndex}`"
              class="small"
              :symbol="symbol"
            />
          </div>
        </td>
        <td class="w-60 border px-2 py-1">{{ line.winWay }}</td>
        <td class="w-60 border px-2 py-1">{{ line.multi }}</td>
        <td class="w-60 border px-2 py-1">{{ props.wildMulti === 0 ? '-' : props.wildMulti }}</td>
        <td class="w-60 border px-2 py-1">{{ line.wins }}</td>
        <td class="w-60 border px-2 py-1">{{ line.freeGameCount }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
// import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import JumphighslotPic from '@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotPic.vue'

const { t } = useI18n()

const props = defineProps({
  lines: {
    type: Array,
    default: () => [],
  },
  wildMulti: {
    type: Number,
    default: 0,
  },
})
</script>
