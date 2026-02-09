<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    default: 'text',
  },
  placeholder: {
    type: String,
    default: '',
  },
  maxlength: {
    type: Number,
    default: 0,
  },
  errorMessage: {
    type: String,
    default: '',
  },
  rightIcon: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['update:modelValue', 'clickIcon'])

const scopeRef = ref<HTMLDivElement | null>(null)
const inputRef = ref<HTMLInputElement | null>(null)
const inputValue = computed({
  get() {
    return props.modelValue
  },
  set(newValue) {
    emit('update:modelValue', newValue)
  },
})
const isInputFocus = ref(false)

watch(isInputFocus, (newValue) => {
  if (newValue) {
    inputRef.value?.focus()
  }
})

function onInputBlur(event: FocusEvent) {
  if (event.relatedTarget !== null && scopeRef.value?.contains(event.relatedTarget as HTMLElement)) {
    inputRef.value?.focus()
  } else {
    isInputFocus.value = false
  }
}

function setInputMaxLength(maxLength: number) {
  if (maxLength) {
    inputRef.value?.setAttribute('maxlength', maxLength.toString())
  }
}

watch(() => props.maxlength, setInputMaxLength)

onMounted(() => {
  setInputMaxLength(props.maxlength)
})
</script>

<template>
  <div ref="scopeRef" @click="isInputFocus = true">
    <div
      class="relative rounded border"
      :class="props.errorMessage !== '' ? 'border-danger' : isInputFocus ? 'border-primary' : 'border-gray-200'"
    >
      <div class="flex">
        <input
          ref="inputRef"
          v-model="inputValue"
          :type="props.type"
          class="form-input m-[1px]"
          @focus="isInputFocus = true"
          @blur="(e) => onInputBlur(e)"
        />
        <div
          v-if="props.rightIcon"
          class="flex w-10 cursor-pointer items-center justify-center border-l border-gray-200 bg-gray-100 px-2 py-1.5"
          tabindex="-1"
          @click="emit('clickIcon')"
        >
          <slot name="icon"></slot>
        </div>
      </div>
      <div
        v-if="props.placeholder !== ''"
        class="absolute bottom-[9px] left-2 max-w-[calc(100%-(2*8px))] px-2 text-[17px] transition-transform"
        :class="[
          {
            '-translate-x-4 translate-y-[-21px] scale-75 bg-white': isInputFocus || props.modelValue !== '',
          },
          props.errorMessage !== '' ? 'text-danger' : isInputFocus ? 'text-primary' : 'text-gray-500',
        ]"
      >
        {{ props.placeholder }}
      </div>
    </div>
    <div v-show="props.errorMessage !== ''" class="text-danger text-xs">* {{ props.errorMessage }}</div>
  </div>
</template>
