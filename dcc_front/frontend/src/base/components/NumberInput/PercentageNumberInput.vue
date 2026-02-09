<template>
  <!-- 顯示百分比用的input -->
  <input
    v-model="inputModel"
    type="text"
    :disabled="props.isDisabled"
    @blur="onInputBlur()"
    @focus="inputModel = inputModel.replace('%', '')"
  />
</template>
<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: Number,
    default: 0,
  },
  isDisabled: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['update:modelValue'])

const inputModel = ref(props.modelValue.toString(10).concat('%'))

watch(
  () => props.modelValue,
  (newValue) => {
    inputModel.value = newValue.toString(10).concat('%')
  }
)

// 預設防呆，>=0 且 <=100
function checkVal(val) {
  let numVal = Number(val)
  if (numVal > 100) numVal = 100
  else if (numVal < 0 || !numVal) numVal = 0
  return numVal
}

function onInputBlur() {
  const val = checkVal(inputModel.value)
  inputModel.value = val.toString(10).concat('%')
  emit('update:modelValue', val)
}
</script>
