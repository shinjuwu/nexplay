<template>
  <div class="mb-4 border-y-2 py-2 text-center text-xl font-bold">
    <div>{{ t('textSlotBonusGame') }}（Bonus Game）</div>
    <div>{{ t('fmtTextSlotFreeGameMainGameLognumber', [props.mainGameLognumber]) }}</div>
  </div>

  <div class="pyrtreasureslot-BGboard relative">
    <div class="pyrtreasureslot-boxSection">
      <div class="pyrtreasureslot-bonusBox pyrtreasureslot-44">
        <div
          class="pyrtreasureslot-multi"
          :class="{
            'pyrtreasureslot-multi2': props.multi === 2,
            'pyrtreasureslot-multi3': props.multi === 3,
            'pyrtreasureslot-multi5': props.multi === 5,
          }"
        ></div>
      </div>
      <div v-for="(lines, index) in bonusBoard" :key="index" class="pyrtreasureslot-bonusBoxLine">
        <div
          v-for="(box, index) in lines"
          :key="index"
          class="pyrtreasureslot-bonusBox"
          :class="{
            'pyrtreasureslot-41': box !== 0 && box !== -1,
            'pyrtreasureslot-42': box === 0,
            'pyrtreasureslot-43': box === -1,
          }"
        >
          <div v-if="box > 0" class="pyrtreasureslot-bonusNum">${{ box }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const props = defineProps({
  mainGameLognumber: {
    type: String,
    default: '',
  },
  bonusBoard: {
    type: Array,
    default: () => [],
  },
  multi: {
    type: Number,
    default: 0,
  },
})

console.log(props.multi)
</script>
