import { format, parseISO, startOfMinute } from 'date-fns'

import { getServerSetting } from '@/base/api/sysSystem'

class Time {
  constructor() {
    const now = new Date()

    this.year = now.getFullYear()

    this.commonReportTimeBeforeDays = 0
    this.commonReportTimeMinuteIncrement = 0
    this.commonReportTimeRange = 0
    this.earningReportTimeMinuteIncrement = 0
    this.winloseReportTimeBeforeDays = 0
    this.winloseReportTimeRange = 0
  }

  async init() {
    const resp = await getServerSetting()

    const data = resp.data.data
    this.commonReportTimeBeforeDays = data.common_report_time_before_days
    this.commonReportTimeMinuteIncrement = data.common_report_time_minute_increment
    this.commonReportTimeRange = data.common_report_time_range
    this.earningReportTimeMinuteIncrement = data.earning_report_time_minute_increment
    this.winloseReportTimeBeforeDays = data.winlose_report_time_before_days
    this.winloseReportTimeRange = data.winlose_report_time_range
  }

  getCurrentTimeRageByMinutesAndStep(range, step) {
    if (!range || isNaN(range) || range < 0) {
      range = 30
    }

    if (!step || isNaN(step) || step < 1 || 60 % step !== 0) {
      step = 5
    }

    let now = startOfMinute(new Date())
    now.setMinutes(now.getMinutes() - (now.getMinutes() % step) + step)
    const endTime = new Date(now)

    now.setMinutes(now.getMinutes() - range)
    const startTime = new Date(now)

    return { startTime, endTime }
  }

  recordTimeFormat(recordTime) {
    return format(recordTime, 'yyyyMMddHHmmss')
  }

  localDateFormat(localTime) {
    return format(localTime, 'yyyy-MM-dd')
  }

  localTimeFormat(localTime) {
    return format(localTime, 'yyyy-MM-dd HH:mm:ss')
  }

  localTimeToUtcTimeSting(localTime) {
    return new Date(localTime).toISOString()
  }

  localTimeToMonth(localTime) {
    return format(new Date(localTime), 'yyyy-MM')
  }

  utcTimeStrToLocalTimeFormat(utcTimeStr) {
    return format(parseISO(utcTimeStr), 'yyyy-MM-dd HH:mm:ss')
  }

  utcTimeStrNoneISO8601ToLocalTimeFormat(utcTimeStr) {
    return utcTimeStr === '-' ? utcTimeStr : this.localTimeFormat(new Date(utcTimeStr))
  }

  async delay(time) {
    return new Promise((resolve) => setTimeout(resolve, time))
  }
}

export default new Time()
