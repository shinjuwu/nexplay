import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'

export function useGameCannedLanguage() {
  const { t } = useI18n()

  const pageTabs = computed(() => {
    return Object.values(constant.CannedLanguageType)
      .sort((a, b) => a - b)
      .map((cannedLanguageType) => t(`cannedLanguageType__${cannedLanguageType}`))
  })
  const activePageTab = ref(pageTabs.value[0])

  return {
    pageTabs,
    activePageTab,
  }
}
