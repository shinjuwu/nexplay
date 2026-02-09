<template>
  <FormDropdown
    v-if="isInit"
    :fmt-item-text="
      (roomType) =>
        roomType === constant.RoomType.All
          ? t('textRoomTypeAll')
          : t(`roomType__${roomTypeNameIndex(constant.GameType.Slot * 1000, roomType)}`)
    "
    :fmt-item-key="(roomType) => `roomType__${roomType}`"
    :items="roomTypes"
    :use-font-awsome="false"
    :model-value="props.modelValue"
    @update:model-value="(newValue) => emit('update:modelValue', newValue)"
  />
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import { roomTypeNameIndex } from '@/base/utils/room'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'

const { t } = useI18n()
const roomTypes = [constant.RoomType.All, constant.RoomType.Newbie]
const isInit = ref(false)

onMounted(() => {
  isInit.value = true
  emit('update:modelValue', roomTypes[0])
})

const props = defineProps({
  modelValue: {
    type: [Number, String],
    default: constant.RoomType.All,
  },
})
const emit = defineEmits(['update:modelValue'])
</script>
