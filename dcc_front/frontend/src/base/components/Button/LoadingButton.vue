<template>
  <button
    class="flex items-center justify-around"
    type="button"
    :disabled="showProcessing || isDataUpdated || parentDisabled"
    @click="handleClick"
  >
    <div v-show="showProcessing" class="mr-2 flex items-center justify-center">
      <span class="h-5 w-5 animate-spin rounded-full border-4 border-b-transparent"></span>
    </div>
    <slot></slot>
  </button>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { cloneDeep, isEqual } from 'lodash'

const props = defineProps({
  buttonClick: {
    type: Function,
    required: true,
  },
  isGetData: {
    type: Boolean,
    default: false,
  },
  parentData: {
    type: [Object, String, Array],
    default: () => {},
  },
  parentDisabled: {
    type: Boolean,
    default: false,
  },
})

const originData = ref(null)
const isDataReceived = ref(false)
const isDataUpdated = computed(() => isDataReceived.value && isEqual(originData.value, props.parentData))

watch(
  () => props.isGetData,
  () => {
    isDataReceived.value = props.isGetData
    originData.value = cloneDeep(props.parentData)
  },
  { immediate: true }
)

const showProcessing = ref(false)
async function handleClick() {
  try {
    showProcessing.value = true
    await props.buttonClick()
  } catch (err) {
    console.error(err)
  } finally {
    isDataReceived.value = false
    showProcessing.value = false
  }
}
</script>
