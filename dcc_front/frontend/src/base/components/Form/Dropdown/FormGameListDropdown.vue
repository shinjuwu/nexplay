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
    @update:model-value="(newValue) => updateSelectGame(newValue)"
  />
</template>

<script setup>
import { onMounted, ref, computed, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'

import constant from '@/base/common/constant'
import { useDropdownListStore } from '@/base/store/dropdownStore'
import { round } from '@/base/utils/math'

import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'

const { t } = useI18n()
const dStore = useDropdownListStore()
const { dropdownList } = storeToRefs(useDropdownListStore())

const isInit = ref(false)

const games = computed(() => {
  let games = []

  if (props.includeAll) {
    games.push({ id: constant.Game.All, name: t('textGameAll') })
  }

  games = games.concat(
    dropdownList.value.games.filter((game) => {
      if (props.gameType <= 0) {
        return true
      }
      return Math.floor(round(game.id / 1000, 7)) === props.gameType
    })
  )

  games.sort((a, b) => a.id - b.id)

  return games
})

let currentGame = null
function updateSelectGame(selectGame) {
  currentGame = games.value.find((game) => game.id === selectGame.id) || games.value[0]
  emit('update:modelValue', currentGame)
}

watch(games, (newGameList) => {
  if (newGameList && newGameList.length) {
    let tempGame
    if (currentGame !== null) {
      tempGame = games.value.find((game) => game.id === currentGame.id)
    }

    if (!tempGame) {
      tempGame = games.value[0]
    }

    updateSelectGame(tempGame)
  }
})

onMounted(async () => {
  await dStore.getGameList()
  isInit.value = true
})

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => {},
  },
  gameType: {
    type: Number,
    default: 0,
  },
  includeAll: {
    type: Boolean,
    default: true,
  },
})
const emit = defineEmits(['update:modelValue'])
</script>
