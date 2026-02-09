import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'

export function useRTPSetting() {
  const { t } = useI18n()

  const gameTabs = computed(() => {
    const ignoardGameTypes = [constant.GameType.All, constant.GameType.Lobby, constant.GameType.FriendsRoom]
    return Object.values(constant.GameType)
      .filter((gameType) => ignoardGameTypes.indexOf(gameType) === -1)
      .sort((a, b) => a - b)
      .map((gameType) => t(`gameType__${gameType}`))
  })
  const activeTab = ref(gameTabs.value[0])

  return {
    gameTabs,
    activeTab,
  }
}
