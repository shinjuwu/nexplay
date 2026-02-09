<template>
  <ul>
    <li
      v-for="child in list.items"
      :key="child.key"
      class="whitespace-nowrap border-l border-dashed border-gray-400 before:relative before:top-[-0.4em] before:inline-block before:h-[1em] before:w-5 before:border-b before:border-dashed before:border-gray-400 before:content-[''] last:border-none last:before:border-l"
      :class="{ 'before:border-none': child.key === 'LevelOneMenu' }"
    >
      <span
        v-show="hasChild(child)"
        class="mr-2 inline-flex h-5 w-5 items-center justify-center border border-black leading-5"
        @click.stop="toggleOpen(child)"
      >
        {{ child.show ? '-' : '+' }}
      </span>
      <Checkbox
        v-if="hasChild(child)"
        class="align-text-top"
        :checked="child.checked"
        :class="{ 'pointer-events-none': props.disabled }"
        @click="toggleSelectAll(child, !child.checked)"
      />
      <Checkbox
        v-else
        :checked="child.checked"
        :class="{ 'pointer-events-none': props.disabled }"
        @click="toggleSelect(child, !child.checked)"
      />
      <span class="mx-1.5 inline-block h-6 w-6 p-0.5 align-top">
        <FolderOpenIcon v-if="hasChild(child)" />
        <DocumentIcon v-else />
      </span>
      <span>{{ t(child.nameKey) }}</span>
      <ul v-if="hasChild(child) && child.show" class="pl-14">
        <TreeMenu
          :items="child.items"
          :disabled="props.disabled"
          @update:model-value="(newValue) => emit('update:modelValue', newValue)"
        />
      </ul>
    </li>
  </ul>
</template>

<script setup>
import { reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { FolderOpenIcon, DocumentIcon } from '@heroicons/vue/20/solid'
import Checkbox from '@/base/components/Button/CheckButton.vue'

const { t } = useI18n()

const props = defineProps({
  items: {
    type: Array,
    default: () => [],
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['update:modelValue'])

const list = reactive({ items: props.items })

function toggleOpen(item) {
  item.show = !item.show
}

function toggleSelect(item, check) {
  item.checked = check
  emit('update:modelValue', item)
}

function toggleSelectAll(item, check) {
  item.checked = check
  emit('update:modelValue', item)

  if (!item.items || item.items.length === 0) {
    return
  }
  item.items.forEach((role) => {
    toggleSelectAll(role, check)
  })
}

function hasChild(list) {
  return (list.folders && list.folders.length) || (list.items && list.items.length) ? true : false
}

watch(
  () => props.items,
  () => {
    list.items = props.items
  }
)
</script>
