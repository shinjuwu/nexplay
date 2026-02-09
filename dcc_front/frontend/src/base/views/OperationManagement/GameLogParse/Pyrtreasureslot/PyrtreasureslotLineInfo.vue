<template>
  <table class="text-center">
    <thead>
      <tr class="bg-info text-white">
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineId') }}</th>
        <th class="w-auto border px-2 py-1">{{ t('textSlotLineSymbols') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineWinWay') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineMulti') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotWMulti') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotLineWins') }}</th>
        <th class="w-60 border px-2 py-1">
          {{
            props.bonusWins === 32 || props.bonusWins === 33
              ? t('textSlotLineBonusWins')
              : t('textSlotLineFreeGameWins')
          }}
        </th>
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
          <div class="pyrtreasureslot-smallPic">
            <PyrtreasureslotPic
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
        <td class="w-60 border px-2 py-1">{{ line.bonusGameCount }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import PyrtreasureslotPic from '@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotPic.vue'

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
  bonusWins: {
    type: Number,
    default: 0,
  },
})
</script>
