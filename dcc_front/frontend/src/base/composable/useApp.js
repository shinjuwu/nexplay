import { inject, ref, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import constant from '@/base/common/constant'
import { setLocale } from '@/base/utils/i18n'
import storage from '@/base/utils/storage'
import time from '@/base/utils/time'

export function useApp() {
  const warn = inject('warn')

  const { t } = useI18n()

  const isOnBeforeMountDone = ref(false)

  onBeforeMount(async () => {
    // 從storage取得user使用的locale
    let localeName = storage.local.get(storage.keys.LOCALE)
    // 初次使用另外取得系統預設的locale
    if (!localeName) {
      //TO DO:取得系統預設的locale，先使用固定的預設語言
      localeName = constant.LangType.CHS
    }
    await setLocale(localeName)

    try {
      await time.init()
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    }

    isOnBeforeMountDone.value = true
  })

  return {
    isOnBeforeMountDone,
  }
}
