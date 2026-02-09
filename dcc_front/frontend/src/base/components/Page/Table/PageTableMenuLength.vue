<template>
  <i18n-t
    keypath="fmtTextTableMenuLength"
    tag="div"
    scope="global"
    class="relative flex items-center whitespace-nowrap"
  >
    <template #length>
      <select v-model="selectModel" class="mx-1 w-full rounded border py-2 pl-1 pr-3.5 outline-none">
        <option v-for="len in props.lengthMenu" :key="`lengthMenu__${len}`" :value="len">
          {{ t('fmtTextTableLength', [len]) }}
        </option>
      </select>
    </template>
  </i18n-t>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'

const { t } = useI18n()

const selectModel = computed({
  get() {
    return props.displayLength
  },
  set(value) {
    emit('lengthChange', value)
  },
})

const props = defineProps({
  lengthMenu: {
    type: Array,
    default: constant.TableDefaultLengthMenu,
  },
  displayLength: {
    type: Number,
    default: constant.TableDefaultLength,
  },
})
const emit = defineEmits(['lengthChange'])
</script>
