<template>
  <section v-show="isShow" class="relative">
    <component
      v-bind="props.item.props"
      :is="props.menuItemMap[item.key].component"
      v-if="isInitialized"
      :key="`${props.item.key}_${props.item.forceUpdateCount}`"
    ></component>
  </section>
</template>

<script setup>
import { useTabItem } from '@/base/composable/home/useTabItem'

const props = defineProps({
  menuItemMap: {
    type: Object,
    default: () => {},
  },
  item: {
    type: Object,
    default: () => {},
  },
})

const { isInitialized, isShow } = useTabItem(props.item.key)
</script>
