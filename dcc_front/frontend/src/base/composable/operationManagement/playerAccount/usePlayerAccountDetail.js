import { watch, ref, inject, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import { numberToStr } from '@/base/utils/formatNumber'
import { round } from '@/base/utils/math'

export function usePlayerAccountDetail(props, emit) {
  const { t } = useI18n()
  const warn = inject('warn')

  const visible = ref(false)

  const playerInfo = reactive({
    id: 0,
    userName: '',
    sumValidBet: 0,
    monthValidBet: 0,
    sumProfit: 0,
    monthProfit: 0,
    remark: '',
    state: 0,
  })

  function close() {
    visible.value = false
    emit('close', false)
  }

  async function searchGameUsersInfo() {
    try {
      const resp = await api.getGameUsersInfo({ id: props.playerInfo.id, time_zone: new Date().getTimezoneOffset() })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`)).then(() => emit('close', false))
        return
      }

      const data = resp.data.data
      playerInfo.id = props.playerInfo.id
      playerInfo.userName = props.playerInfo.userName
      playerInfo.sumValidBet = numberToStr(data.valid_bet_sum)
      playerInfo.monthValidBet = numberToStr(data.valid_bet)
      playerInfo.sumProfit = numberToStr(round(data.profit_sum - data.tax_sum + data.bonus_sum, 4))
      playerInfo.monthProfit = numberToStr(round(data.profit - data.tax + data.bonus, 4))
      playerInfo.remark = data.info
      playerInfo.state = data.is_enabled ? constant.AccountStatus.Open : constant.AccountStatus.Disable

      visible.value = true
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`)).then(() => emit('close', false))
      }
    }
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        await searchGameUsersInfo()
      } else {
        visible.value = false
      }
    }
  )

  return {
    playerInfo,
    visible,
    close,
  }
}
