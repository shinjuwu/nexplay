import { inject, reactive, onBeforeMount } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'

import constant from '@/base/common/constant'
import * as gApi from '@/base/api/sysGlobal'
import time from '@/base/utils/time'
import * as uApi from '@/base/api/sysUser'
import { getGameUserLoginDataInput } from '@/base/common/table/getGameUserLoginDataInput'

export function usePlayerLoginInfo() {
  const warn = inject('warn')
  const { t } = useI18n()

  const deviceList = [
    constant.Device.All,
    constant.Device.PC,
    constant.Device.Mac,
    constant.Device.IOS,
    constant.Device.Android,
    constant.Device.Other,
  ]

  const browserList = reactive({
    items: [constant.Browser.All],
  })

  async function updateBrowserList() {
    try {
      const resp = await gApi.getUserLoginLogBroswerList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      browserList.items = [browserList.items[0], ...resp.data.data]
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    }
  }

  const countryList = reactive({
    items: [constant.Country.All],
  })

  async function updateCountryList() {
    try {
      const resp = await gApi.getUserLoginLogCountryShortList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      countryList.items = [countryList.items[0], ...resp.data.data]
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    }
  }

  onBeforeMount(async () => {
    await updateBrowserList()
    await updateCountryList()
  })

  const { startTime, endTime } = time.getCurrentTimeRageByMinutesAndStep()
  const formInput = reactive({
    agent: {},
    device: deviceList[0],
    browser: browserList.items[0],
    country: countryList.items[0],
    userName: '',
    ip: '',
    startTime: startTime,
    endTime: endTime,
  })

  const tableInput = reactive(
    new getGameUserLoginDataInput(
      formInput.agent.id,
      formInput.device,
      formInput.browser,
      formInput.country,
      formInput.userName,
      formInput.ip,
      formInput.startTime,
      formInput.endTime,
      constant.TableDefaultLength,
      'loginTime',
      constant.TableSortDirection.Desc
    )
  )

  const playerLoginInfo = reactive({
    items: [],
  })

  function validateForm() {
    let errors = []

    const startTime = formInput.startTime
    const endTime = formInput.endTime
    if (endTime - startTime < 0) {
      errors.push(t('textStartTimeLaterThanEndTime'))
      return errors
    } else if (endTime - startTime > time.winloseReportTimeRange * 60 * 1000) {
      const hour = Math.floor(time.winloseReportTimeRange / 60)
      const minute = time.winloseReportTimeRange % 60

      let timeTextKey = ''
      let timeTextArgs = []
      if (hour > 0 && minute > 0) {
        timeTextKey = 'fmtTextHoursMinutes'
        timeTextArgs.push(hour, minute)
      } else if (hour > 0) {
        timeTextKey = 'fmtTextHours'
        timeTextArgs.push(hour)
      } else {
        timeTextKey = 'fmtTextMinutes'
        timeTextArgs.push(minute)
      }

      errors.push(t(`errorCode__${constant.ErrorCode.ErrorTimeRange}`, [t(timeTextKey, timeTextArgs)]))
      return errors
    }

    return errors
  }

  async function searchPlayerLoginInfo() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join('\n'))
      return
    }

    tableInput.showProcessing = true

    const tmpTableInput = new getGameUserLoginDataInput(
      formInput.agent.id,
      formInput.device,
      formInput.browser,
      formInput.country,
      formInput.userName,
      formInput.ip,
      formInput.startTime,
      formInput.endTime,
      tableInput.length,
      tableInput.column,
      tableInput.dir
    )

    try {
      const resp = await uApi.getGameUserLoginData(tmpTableInput.parseInputJson())

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      playerLoginInfo.items = resp.data.data.map((d) => ({
        agentName: d.agent_name,
        userName: d.game_user_name,
        deviceName: d.device_name,
        os: d.os,
        osVer: d.os_ver,
        osName: `${d.os} ${d.os_ver}`,
        browser: d.browser,
        browserVer: d.browser_ver,
        resolution: { width: d.width, height: d.height },
        resolutionName: `${d.width}x${d.height}`,
        loginTime: d.login_time,
        logoutTime: d.logout_time,
        ip: d.ip,
        country: d.country,
        countryShort: d.country_s,
        region: d.region,
        city: d.city,
        location: `${d.country}/${d.region}/${d.city}`,
        token: d.token,
      }))

      Object.assign(tableInput, tmpTableInput)
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

  return {
    deviceList,
    browserList,
    countryList,
    formInput,
    tableInput,
    playerLoginInfo,
    searchPlayerLoginInfo,
  }
}
