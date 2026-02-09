import { computed, inject, onBeforeMount, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysSystem'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useUserStore } from '@/base/store/userStore'

export function useExchangeRateSetting() {
  const warn = inject('warn')

  const { t } = useI18n()

  const isSettingEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.ExchangeRateSettingUpdate)
  })

  const exchangeDataDirections = reactive({
    items: [
      t('textToCoinDecimalSupportDirection'),
      t('textBeforeExchangeRateSettingDirection'),
      t('textAfterExchangeRateSettingDirection'),
    ],
  })

  const exchangeData = reactive({ items: [] })
  const tableInput = reactive(new BaseTableInput(exchangeData.items.length, 'id', constant.TableSortDirection.Asc))

  async function getExchangeDataList() {
    try {
      tableInput.showProcessing = true

      const resp = await api.getExchangeDataSetting()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      exchangeData.items = resp.data.data.map((d) => {
        return {
          id: d.id,
          currency: d.currency,
          toCoin: d.to_coin,
        }
      })
      tableInput.length = exchangeData.items.length
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableInput.showProcessing = false
    }
  }

  async function setExchangeDataList() {
    for (const item of exchangeData.items) {
      if (item.toCoin <= 0) {
        warn(t('textExchangeRateGreaterThanZero'))
        return
      }

      const splitToCoin = `${item.toCoin}`.split('.')
      const toCoinDecimalParts = splitToCoin.length > 1 ? splitToCoin[1] : ''
      if (toCoinDecimalParts.length > 4) {
        warn(t('textExchangeRateDecimalPlaceLessOrEquralToFour'))
        return
      }
    }

    try {
      tableInput.showProcessing = true

      const resp = await api.setExchangeDataSetting(
        exchangeData.items.map((item) => {
          return {
            id: item.id,
            currency: item.currency,
            to_coin: item.toCoin,
          }
        })
      )

      warn(t(`errorCode__${resp.data.code}`))
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableInput.showProcessing = false
    }
  }

  onBeforeMount(async () => {
    await getExchangeDataList()
  })

  return {
    exchangeData,
    exchangeDataDirections,
    isSettingEnabled,
    tableInput,
    setExchangeDataList,
  }
}
