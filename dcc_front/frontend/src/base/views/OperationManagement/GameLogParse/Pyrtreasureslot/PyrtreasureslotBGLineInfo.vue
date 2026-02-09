<template>
  <table class="text-center">
    <thead>
      <tr class="bg-info text-white">
        <th class="w-60 border px-2 py-1">Level</th>
        <th class="w-auto border px-2 py-1">{{ t('textSlotBonusSymbols') }}</th>
        <th class="w-60 border px-2 py-1">{{ t('textSlotBonusWin') }}</th>
      </tr>
    </thead>
    <tbody>
      <!-- <tr>
        <td class="w-60 border px-2 py-1">1</td>
        <td class="border px-2 py-1">
          <div class="pyrtreasureslot-smallPic">
            <PyrtreasureslotPic
              v-for="(symbol, sIndex) in props.bonusLineInfo"
              :key="`line__${index}_symbol__${symbol}__${sIndex}`"
              class="small"
              :symbol="symbol"
            />
          </div>
        </td>
        <td class="w-60 border px-2 py-1">1,000</td>
      </tr> -->
      <template v-if="props.bonusLineInfo.length === 0">
        <td class="w-60 border px-2 py-1">-</td>
        <td class="border px-2 py-1">{{ t('textSlotNoLineSymbols') }}</td>
        <td class="w-60 border px-2 py-1">-</td>
      </template>
      <tr v-for="line in props.bonusLineInfo" :key="`line_${line.key}`">
        <td v-if="line.bonus > 0" class="w-60 border px-2 py-1">{{ line.level }}</td>
        <td v-if="line.bonus > 0" class="border px-2 py-1">
          <div class="pyrtreasureslot-smallPic">
            <div v-for="(box, index) in line.lines" :key="index">
              <div class="pyrtreasureslot-41 pyrtreasureslot small"></div>
            </div>
          </div>
        </td>
        <td v-if="line.bonus > 0" class="w-60 border px-2 py-1">{{ line.bonus }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const props = defineProps({
  bonusLineInfo: {
    type: Array,
    default: () => [],
  },
})
</script>
