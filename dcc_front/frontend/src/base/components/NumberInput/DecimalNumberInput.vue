<template>
  <!-- 小數點用的input -->
  <input
    v-model="inputModel"
    type="text"
    :disabled="props.isDisabled"
    class="form-input"
    :class="`text-${props.align}`"
    @change="onInputChange()"
  />
</template>

<script setup>
import { ref, watch } from 'vue'
import { numberToStr, strToNumber } from '@/base/utils/formatNumber'

const props = defineProps({
  modelValue: {
    type: Number,
    default: 0,
  },
  isDisabled: {
    type: Boolean,
    default: false,
  },
  align: {
    type: String,
    default: 'left',
  },
})
const emit = defineEmits(['update:modelValue'])

const inputModel = ref(numberToStr(props.modelValue))

watch(
  () => props.modelValue,
  (newValue) => {
    const maxValue = 999999999999
    if (newValue > maxValue) {
      inputModel.value = numberToStr(maxValue)
      return
    }
    inputModel.value = numberToStr(newValue)
  }
)

function onInputChange() {
  let numValue = strToNumber(inputModel.value)
  if (isNaN(numValue)) {
    numValue = 0
  }
  emit('update:modelValue', numValue)
}
</script>
