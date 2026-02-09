<template>
  <div ref="scopeElRef" class="relative">
    <button
      class="btn btn-light flex w-full items-center justify-center text-black"
      :disabled="props.disabled"
      @click.prevent="dropdownShow = !dropdownShow"
      @blur="(e) => onBlurHideDropdownMenu(e)"
    >
      {{ props.fmtItemText(dropdownModel) }}
      <ChevronDownIcon class="ml-1 h-3 w-3 text-black" />
    </button>
    <ul v-show="dropdownShow" class="absolute left-0 z-10 w-full max-w-md rounded border-2 border-gray-200 bg-white">
      <li v-if="props.useFilter" class="sticky top-0 mb-1 border-b-2 border-b-gray-200 p-2">
        <input
          ref="filterInput"
          v-model="dropdownFilter"
          type="text"
          class="w-full rounded border-2 border-gray-200 p-1"
          @blur="(e) => onBlurHideDropdownMenu(e)"
        />
      </li>
      <div class="max-h-80 overflow-y-scroll">
        <li
          v-for="(item, index) in dropdownItems"
          :key="fmtItemKey(item, index)"
          tabIndex="-1"
          class="p-2 text-center text-gray-500 hover:bg-gray-100 hover:text-black"
          @click="setDropdownModel(item)"
        >
          {{ props.fmtItemText(item) }}
        </li>
      </div>
    </ul>
  </div>
</template>

<script setup>
import { computed, ref, nextTick, watch } from 'vue'
import { ChevronDownIcon } from '@heroicons/vue/24/outline'

const filterInput = ref()

const dropdownShow = ref(false)
const dropdownModel = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  },
})
const dropdownFilter = ref('')
const dropdownItems = computed(() => {
  return props.items.filter((item) => props.filterRule(item, dropdownFilter.value))
})

function setDropdownModel(item) {
  dropdownModel.value = item
  dropdownShow.value = false
}

// 有使用filter打開時清除input
if (props.useFilter) {
  watch(dropdownShow, (newValue) => {
    if (newValue) {
      dropdownFilter.value = ''
      nextTick(() => {
        filterInput.value.focus()
      })
    }
  })
}

const scopeElRef = ref()
/**
 * dropdownItems添加tabindex可使HTML elements focusable
 * source: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/tabindex
 */
function onBlurHideDropdownMenu(e) {
  if (e.relatedTarget !== null && scopeElRef.value.contains(e.relatedTarget)) {
    return
  }
  dropdownShow.value = false
}

const props = defineProps({
  fmtItemText: {
    type: Function,
    default: (item) => item,
  },
  fmtItemKey: {
    type: Function,
    default: (item) => item,
  },
  modelValue: {
    type: [Object, Number, String, Date],
    required: true,
  },
  direction: {
    type: String,
    default: 'bottom',
  },
  items: {
    type: Array,
    default: () => [],
  },
  useFontAwsome: {
    type: Boolean,
    default: true,
  },
  useFilter: {
    type: Boolean,
    default: false,
  },
  filterRule: {
    type: Function,
    default: () => true,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['update:modelValue'])
</script>

<style scoped>
.dropdown-menu.top {
  bottom: calc(100% + 2px);
}

.dropdown-menu.bottom {
  top: calc(100% + 2px);
}

.dropdown-menu.left {
  right: calc(100% + 2px);
}

.dropdown-menu.right {
  left: calc(100% + 2px);
}
</style>
