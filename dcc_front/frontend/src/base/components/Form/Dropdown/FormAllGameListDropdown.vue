<template>
  <FormDropdown
    v-if="isInit"
    :fmt-item-text="(game) => game.name"
    :fmt-item-key="(game) => `game__${game.id}`"
    :items="games"
    :use-font-awsome="false"
    :use-filter="true"
    :filter-rule="(game, filterValue) => game.name.includes(filterValue)"
    :model-value="props.modelValue"
    @update:model-value="(newValue) => emit('update:modelValue', newValue)"
  />
</template>

<script setup>
import { onMounted, ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'
import { storeToRefs } from 'pinia'
import { useDropdownListStore } from '@/base/store/dropdownStore'

const { t } = useI18n()
const dStore = useDropdownListStore()
const { dropdownList } = storeToRefs(useDropdownListStore())
const isInit = ref(false)

const games = computed(() => {
  let games = [{ id: constant.Game.All, name: t('textGameAll') }]
  const gameList = games.concat(dropdownList.value.allGames)
  gameList.sort((a, b) => a.id - b.id)
  if (props.agentRtp) {
    const allGameList = gameList.filter((game) => game.id < 5000 || game.id > 6000) //過濾好友房系列
    return allGameList
  }
  return gameList
})

onMounted(async () => {
  await dStore.getAllGameList()
  isInit.value = true
  emit('update:modelValue', games.value[0])
})

const props = defineProps({
  modelValue: {
    type: [Object, String],
    default: () => {},
  },
  gameType: {
    type: Number,
    default: 0,
  },
  agentRtp: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['update:modelValue'])
</script>
