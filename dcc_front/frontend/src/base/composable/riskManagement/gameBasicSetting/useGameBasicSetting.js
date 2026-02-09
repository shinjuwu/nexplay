import { inject, reactive, ref, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'

import * as api from '@/base/api/sysRiskControl'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'
import { round } from '@/base/utils/math'
import { numberRangeValidate } from '@/base/utils/validate'

export function useGameBasicSetting() {
  const warn = inject('warn')
  const { t } = useI18n()

  const pageDirections = reactive({
    items: [
      t('textGameBasicSettingDirections'),
      t('textGameBasicSettingDirections__1'),
      t('textGameBasicSettingDirections__2'),
    ],
  })

  const isEditEnabled = (() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.GameBasicSettingUpdate)
  })()

  // 是否啟用先判別game type，如果有特例二次判斷特例的game id
  // 控牌RTP
  function isMatchGameRTPEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.ChiPai:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 控牌殺率
  function isMatchGameKillRateEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.ChiPai:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 控牌場次
  function isMatchGamesEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.ChiPai:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 單間RTP
  function isNormalMatchGameRTPEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.BaiRen:
      case constant.GameType.ChiPai:
      case constant.GameType.ElectronicGame:
      case constant.GameType.Slot:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 單間殺率
  function isNormalMatchGameKillRateEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.BaiRen:
      case constant.GameType.ChiPai:
      case constant.GameType.ElectronicGame:
      case constant.GameType.Slot:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 保底RTP
  function isLowBoundRTPEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.BaiRen:
      case constant.GameType.ChiPai:
      case constant.GameType.ElectronicGame:
      case constant.GameType.Slot:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 限制倍率
  function isLimitOddsEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.Slot:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }
  // 禁開分數
  function isLimitMoneyEnabled(gameId) {
    const gameType = Math.floor(gameId / 1000)
    switch (gameType) {
      case constant.GameType.Slot:
        return true
    }

    switch (gameId) {
      default:
        return false
    }
  }

  const showProcessing = ref(false)
  const gameSettings = reactive({ items: [] })

  async function searchGameSettings() {
    showProcessing.value = true

    try {
      const resp = await api.getGameSetting()
      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      gameSettings.items = resp.data.data
        .map((d) => ({
          gameId: d.gameId,
          matchGameRTP: d.MatchGameRTP * 100,
          matchGameKillRate: d.MatchGameKillRate * 100,
          matchGames: d.MatchGames,
          normalMatchGameRTP: d.NormalMatchGameRTP * 100,
          normalMatchGameKillRate: d.NormalMatchGameKillRate * 100,
          lowBoundRTP: d.LowBoundRTP * 100,
          limitOdds: d.LimitOdds,
          limitMoney: d.LimitMoney,
        }))
        .sort((a, b) => a.gameId - b.gameId)
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      showProcessing.value = false
    }
  }
  onBeforeMount(async () => {
    await searchGameSettings()
  })

  function validateGameSettings() {
    const errors = []

    const validate = (gameName, propertyValue, minValue, maxValue, errorMsgKey) => {
      if (!numberRangeValidate(propertyValue, minValue, maxValue)) {
        errors.push(t(errorMsgKey, [gameName, minValue.toFixed(2), maxValue.toFixed(2)]))
      }
    }

    for (const item of gameSettings.items) {
      const gameName = t(`game__${item.gameId}`)
      validate(gameName, item.matchGameRTP, 0, 100, 'fmtTextMatchGameRTPRangeError')
      validate(gameName, item.matchGameKillRate, 0, 100, 'fmtTextMatchGameKillRateRangeError')
      validate(gameName, item.matchGames, 0, 9999, 'fmtTextMatchGamesRangeError')
      validate(gameName, item.normalMatchGameRTP, 0, 100, 'fmtTextNormalMatchGameRTPRangeError')
      validate(gameName, item.normalMatchGameKillRate, 0, 100, 'fmtTextNormalMatchGameKillRateRangeError')
      validate(gameName, item.lowBoundRTP, 0, 100, 'fmtTextLowBoundRTPRangeError')
      validate(gameName, item.limitOdds, 0, 9999, 'fmtTextLimitOddsRangeError')
      validate(gameName, item.limitMoney, 0, 999999, 'fmtTextLimitMoneyRangeError')
    }

    return errors
  }

  async function setGameSettings() {
    showProcessing.value = true

    try {
      const errors = validateGameSettings()
      if (errors.length > 0) {
        warn(errors.join('\n'))
        return
      }

      const resp = await api.setGameSetting(
        gameSettings.items.map((d) => ({
          gameId: d.gameId,
          MatchGameRTP: round(d.matchGameRTP / 100, 4),
          MatchGameKillRate: round(d.matchGameKillRate / 100, 4),
          MatchGames: d.matchGames,
          NormalMatchGameRTP: round(d.normalMatchGameRTP / 100, 4),
          NormalMatchGameKillRate: round(d.normalMatchGameKillRate / 100, 4),
          LowBoundRTP: round(d.lowBoundRTP / 100, 4),
          LimitOdds: round(d.limitOdds, 4),
          LimitMoney: round(d.limitMoney, 4),
        }))
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
      showProcessing.value = false
    }
  }

  return {
    gameSettings,
    isEditEnabled,
    isMatchGameRTPEnabled,
    isMatchGameKillRateEnabled,
    isMatchGamesEnabled,
    isNormalMatchGameRTPEnabled,
    isNormalMatchGameKillRateEnabled,
    isLowBoundRTPEnabled,
    isLimitOddsEnabled,
    isLimitMoneyEnabled,
    pageDirections,
    showProcessing,
    setGameSettings,
  }
}
