import { onBeforeMount, ref, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'
import axios from 'axios'
import * as api from '@/base/api/sysJackpot'

export function useJackpotPrizeInfo() {
  const { t } = useI18n()
  const warn = inject('warn')

  const pageDirections = ref([])
  const showPrizeInfoProcessing = ref(false)

  const showPool = ref(0)
  const realPool = ref(0)
  const reservePool = ref(0)

  const jackpotShowPool = ref(0)
  const injectRate = ref(0)

  async function getJackpotPoolData() {
    try {
      showPrizeInfoProcessing.value = true
      const resp = await api.getJackpotPoolData()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      showPool.value = data.show_pool
      realPool.value = data.real_pool
      reservePool.value = data.reserve_pool

      if (jackpotShowPool.value !== data.jp_show_pool || injectRate.value !== data.jp_inject_water_rate * 100) {
        jackpotShowPool.value = data.jp_show_pool
        injectRate.value = data.jp_inject_water_rate * 100

        pageDirections.value = [
          t('fmtTextJackpotPrizeInfoDirection__1', [jackpotShowPool.value.toLocaleString()]),
          t('fmtTextJackpotPrizeInfoDirection__2', [injectRate.value]),
          t('textJackpotPrizeInfoDirections__1'),
          t('textJackpotPrizeInfoDirections__2'),
        ]
      }
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showPrizeInfoProcessing.value = false
    }
  }

  onBeforeMount(async () => {
    await getJackpotPoolData()
  })

  return {
    showPool,
    realPool,
    reservePool,
    showPrizeInfoProcessing,
    pageDirections,
    getJackpotPoolData,
  }
}
