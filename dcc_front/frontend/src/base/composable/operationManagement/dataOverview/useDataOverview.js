import { inject, reactive, ref, nextTick, onMounted, onUnmounted, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'

import * as api from '@/base/api/sysManage'
import constant from '@/base/common/constant'
import { numberToStr } from '@/base/utils/formatNumber'
import { useUserStore } from '@/base/store/userStore'
import { round } from '@/base/utils/math'

import * as echarts from 'echarts/core'
import { BarChart, LineChart, PieChart } from 'echarts/charts'
import { GridComponent, TitleComponent, TooltipComponent, LegendComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import { addMinutes, format, parse, startOfMinute } from 'date-fns'
import { roleItemKey } from '@/base/common/menuConstant'

echarts.use([
  TitleComponent,
  CanvasRenderer,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  BarChart,
  LineChart,
  PieChart,
])

export function useDataOverview() {
  const { t } = useI18n()
  const warn = inject('warn')

  const timeTypes = ['day', 'week', 'month']

  const isShowIsSearchAll = (() => {
    const { user } = storeToRefs(useUserStore())
    return user.value.accountType <= constant.AccountType.General
  })()
  const isDeviceLocationInfoEnabled = (() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.DataOverviewDeviceLocationRead)
  })()

  let dialogChart = null // dialog內圖表
  let chartItems = {
    thirtyDaysChart: null, // 近30日資料
    AllDaysWinLoseChart: null, // 今日各時段輸贏
    AllDaysPlayersChart: null, // 今日各時段人數
    RealTimeUserCountChart: null, // 即時在線人數統計
  }
  let pieChartItems = {
    AllDaysDeviceChart: null, // 今日各時段裝置使用比例
  }
  const baseOption = {
    textStyle: {
      fontSize: 16,
    },
    title: {
      show: true,
      textStyle: {
        color: '#4C5C71',
        fontSize: 24,
      },
      left: 'center',
    },
    legend: {
      show: true,
      left: 'center',
      width: '90%',
      top: '8%',
      padding: [10, 0],
    },
    grid: {
      top: '25%',
      width: '95%',
      left: '5%',
    },
    tooltip: {
      trigger: 'item',
      axisPointer: {
        type: 'none',
      },
      formatter: function (params) {
        return t('fmtTextChartTooltips', [`"${params.seriesName}"`, `"${params.name}"`, numberToStr(params.value)])
      },
    },
    xAxis: {
      type: 'category',
      data: [],
      axisLine: {
        show: false,
      },
      axisTick: {
        show: false,
      },
    },
    yAxis: {
      type: 'value',
      splitNumber: 8,
      axisLabel: {
        fontSize: 14,
      },
    },

    series: [],
  }

  const showChartDialog = ref(false)
  const dialogTitle = ref('')

  const formInput = reactive({
    agent: { id: constant.Agent.All, levelCode: '' },
    timeType: timeTypes[0],
    timeZone: new Date().getTimezoneOffset(),
    isSearchAll: true,
  })

  const chartTable = reactive({
    AllDaysWinLoseChart: [],
    AllDaysPlayersChart: [],
    dialogChart: [],
  })

  const stats = reactive({
    thisTime: {
      activePlayer: 0,
      bettorsNum: 0,
      registerNum: 0,
      oddNum: 0,
      totalBet: 0,
      gameTax: 0,
      totalScore: 0,
    },
    lastTime: {
      activePlayer: 0,
      bettorsNum: 0,
      registerNum: 0,
      oddNum: 0,
      totalBet: 0,
      gameTax: 0,
      totalScore: 0,
    },
  })

  const timesText = reactive({
    than: '',
    current: '',
  })
  const tableData = reactive({
    riskList: [],
    leaderBoards: [],
    locationList: [],
  })

  const hours = Array.apply(null, new Array(24)).map((v, i) => i)

  watch(showChartDialog, () => {
    if (showChartDialog.value) {
      nextTick(() => {
        dialogChart.resize()

        window.addEventListener('resize', () => {
          dialogChart.resize()
        })
      })
    } else {
      dialogChart.dispose()
    }
  })

  async function searchDataOverview() {
    const searchParams = {
      levelCode: formInput.agent.levelCode,
      timeType: formInput.timeType,
      timeZone: formInput.timeZone,
      isSearchAll: formInput.isSearchAll ? 1 : 0,
    }

    try {
      await setCharts()
      await getStatData(searchParams)
      await getRiskUserList(searchParams)
      await getGameLeaderBoards(searchParams)
      if (isDeviceLocationInfoEnabled) {
        await getUserLoginDeviceUsageRatioData(searchParams)
        await getGameUserLocationList(searchParams)
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
      setTimeType()
    }
  }

  async function getStatData(searchParams) {
    const resp = await api.getStatData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    // 資料總覽
    const thisTimeData = JSON.parse(resp.data.data.stat_this_time)
    stats.thisTime.activePlayer = thisTimeData.active_player
    stats.thisTime.bettorsNum = thisTimeData.number_bettors
    stats.thisTime.registerNum = thisTimeData.number_registrants
    stats.thisTime.oddNum = thisTimeData.odd_number
    stats.thisTime.totalBet = thisTimeData.total_betting
    stats.thisTime.gameTax = thisTimeData.game_tax
    stats.thisTime.totalScore = round(
      thisTimeData.platform_total_score + thisTimeData.game_tax - thisTimeData.game_bonus,
      4
    )

    const lastTimeData = JSON.parse(resp.data.data.stat_last_time)
    stats.lastTime.activePlayer = lastTimeData.active_player
    stats.lastTime.bettorsNum = lastTimeData.number_bettors
    stats.lastTime.registerNum = lastTimeData.number_registrants
    stats.lastTime.oddNum = lastTimeData.odd_number
    stats.lastTime.totalBet = lastTimeData.total_betting
    stats.lastTime.gameTax = lastTimeData.game_tax
    stats.lastTime.totalScore = round(
      lastTimeData.platform_total_score + lastTimeData.game_tax - lastTimeData.game_bonus,
      4
    )
  }

  async function getRiskUserList(searchParams) {
    const resp = await api.getRiskUserList(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    // 風險玩家清單
    const riskData = resp.data.data
    tableData.riskList = riskData
      .map((d) => {
        return {
          totalYa: d.total_ya, // 總投注
          totalDe: d.total_de, // 總得分
          userName: d.username,
        }
      })
      .sort((a, b) => b.totalDe - a.totalDe)
  }

  async function getGameLeaderBoards(searchParams) {
    const resp = await api.getGameLeaderBoards(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    // 遊戲輸贏排行榜
    const leaderBoardsData = JSON.parse(resp.data.data)
    tableData.leaderBoards = leaderBoardsData
      .map((d) => {
        return {
          gameID: d.game_id,
          gameName: d.game_name,
          totalBet: round(d.total_ya - d.total_de + d.total_tax - d.total_bonus, 4), // 平台輸贏分數
          totalYa: d.total_ya,
        }
      })
      .sort((a, b) => b.totalBet - a.totalBet)
  }

  async function getUserLoginDeviceUsageRatioData(searchParams) {
    if (pieChartItems.AllDaysDeviceChart !== null) {
      pieChartItems.AllDaysDeviceChart.dispose()
    }

    const chart = echarts.init(document.getElementById('AllDaysDeviceChart'))
    chart.showLoading({ text: '' })

    const resp = await api.getUserLoginDeviceUsageRatioData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    const tmpDeviceArr = []
    // server一定會傳兩個陣列，第一個是昨日，第二個是今日
    const data = resp.data.data[1]
    for (let i = 0; i < data.name_list.length; i++) {
      tmpDeviceArr.push({
        name: data.name_list[i],
        value: data.count_list[i],
      })
    }

    chart.setOption({
      title: {
        text: t('textDeviceUsageRatio'),
        ...baseOption.title,
      },
      textStyle: {
        fontSize: 16,
      },
      legend: baseOption.legend,
      series: [
        {
          type: 'pie',
          name: t('textDeviceUsageRatio'),
          data: tmpDeviceArr,
          center: ['50%', '55%'],
        },
      ],
      tooltip: {
        trigger: 'item',
      },
    })

    chart.resize()
    chart.hideLoading()

    pieChartItems.AllDaysDeviceChart = chart
  }

  async function getGameUserLocationList(searchParams) {
    const resp = await api.getUserLoginSourceLocRankingData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    const tmpLocationList = []
    for (let i = 0; i < resp.data.data.length; i++) {
      if (i > 0) {
        break
      }

      const data = resp.data.data[i]

      for (let i = 0; i < data.country_s_list.length; i++) {
        tmpLocationList.push({
          country: data.country_s_list[i],
          region: data.region_list[i],
          count: data.count_list[i],
        })
      }
    }
    tmpLocationList.sort((a, b) => b.count - a.count)

    tableData.locationList = tmpLocationList
  }

  async function setCharts() {
    Object.keys(chartItems).forEach(async (c) => {
      if (chartItems[c] !== null) {
        chartItems[c].dispose()
      }

      chartItems[c] = echarts.init(document.getElementById(c))
      chartItems[c].showLoading({ text: '' })
      chartItems[c].setOption(baseOption)
      await customOption(chartItems[c], c)
      chartItems[c].hideLoading()

      if (c === 'AllDaysWinLoseChart' || c === 'AllDaysPlayersChart') {
        chartItems[c].on('legendselectchanged', function (params) {
          chartItems[c].dispatchAction({
            type: 'legendSelect',
            name: params.name,
          })

          showChartDialog.value = true
          setChartDialog(chartItems[c], params.name, c)
        })
      }
    })
  }

  async function customOption(chart, chartStr) {
    const searchParams = {
      levelCode: formInput.agent.levelCode,
      timeZone: formInput.timeZone,
      isSearchAll: formInput.isSearchAll ? 1 : 0,
    }

    const xAxisData = hours.map((hour) => {
      return {
        value: `${hour}:00`,
        textStyle: {
          fontSize: 14,
        },
      }
    })

    switch (chartStr) {
      // 近30日資料
      case 'thirtyDaysChart':
        {
          const charData = await getThirtyDaysData(searchParams)
          if (!charData) {
            return
          }

          chart.setOption({
            title: {
              text: t('textLastThirtyDays'),
            },
            tooltip: {
              trigger: 'item',
              axisPointer: {
                type: 'none',
              },
            },
            xAxis: {
              data: charData.xAxisData,
              axisLine: {
                show: false,
              },
              axisLabel: {
                formatter: function (params) {
                  return params.slice(-5)
                },
              },
            },
            yAxis: {
              min: 0,
              max: function (value) {
                return Math.floor(value.max) + 10
              },
              minInterval: 5,
              splitNumber: 5,
            },
            series: charData.seriesData,
          })
        }
        break
      // 各時段輸贏
      case 'AllDaysWinLoseChart':
        {
          const charData = await getAllDaysWinLoseData(searchParams)
          if (!charData) {
            return
          }

          chart.setOption({
            legend: {
              data: charData.legendData,
            },
            title: {
              text: t('textTodayTimePeriodWinLose'),
            },
            xAxis: {
              data: xAxisData,
            },
            yAxis: {
              min: function (value) {
                return Math.floor(value.min) - 200
              },
              max: function (value) {
                return Math.floor(value.max) + 200
              },
              splitNumber: 5,
            },
            series: charData.seriesData,
            animation: false,
          })
        }
        break
      // 各時段、各遊戲投注人數
      case 'AllDaysPlayersChart':
        {
          const charData = await getAllDaysPlayersData(searchParams)
          if (!charData) {
            return
          }

          chart.setOption({
            legend: {
              data: charData.legendData,
            },
            title: {
              text: t('textTodayTimePeriodNumbers'),
            },
            xAxis: {
              data: xAxisData,
            },
            yAxis: {
              min: 0,
              max: function (value) {
                return Math.floor(value.max) + 100
              },
              minInterval: 10,
              splitNumber: 5,
            },
            series: charData.seriesData,
            animation: false,
          })
        }
        break
      // 即時人數統計
      case 'RealTimeUserCountChart':
        {
          const charData = await getRealTimeUserData(searchParams)
          if (!charData) {
            return
          }

          chart.setOption({
            title: {
              text: t('textRealTimeUserCount'),
            },
            xAxis: {
              data: charData.logTime,
              axisLabel: {
                fontSize: 14,
                interval: (charData.logTime.length - 8) / 4 + 1,
                formatter: function (value) {
                  return new Date(value).getHours()
                },
              },
            },

            yAxis: {
              min: 0,
              max: function (value) {
                return Math.floor(value.max) + 1
              },
              minInterval: 1,
              splitNumber: 5,
            },
            series: [
              {
                type: 'line',
                data: charData.yesterdayData,
                smooth: true,
                smoothMonotone: 'none',
                name: t('textYesterday'),
                symbol: 'circle',
                symbolSize: 10,
                lineStyle: {
                  width: 3,
                },
              },
              {
                type: 'line',
                data: charData.todayData,
                smooth: true,
                smoothMonotone: 'none',
                name: t('inputTimeType__day'),
                symbol: 'circle',
                symbolSize: 10,
                lineStyle: {
                  width: 3,
                },
              },
            ],
            animation: false,
            tooltip: {
              trigger: 'axis',
              axisPointer: {
                type: 'line',
              },
              formatter: function (params) {
                const date = new Date(params[0].axisValue)
                const time = `${format(date, 'HH：mm')}`

                let str = `${time}</br>`
                params.forEach((param) => {
                  str += `${param.seriesName}：${param.value}</br>`
                })
                return str
              },
            },
          })
        }
        break
    }

    chart.resize()
  }

  async function getThirtyDaysData(searchParams) {
    const resp = await api.getIntervalDaysBettorData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    const data = Object.values(resp.data.data)
    const thirtyDays = [],
      activePlayer = [],
      betNum = []

    for (let i = 0; i < data.length; i++) {
      const month = data[i].log_time.slice(4, 6)
      const date = data[i].log_time.slice(6)

      thirtyDays.push({
        value: `${month}-${date}`,
        textStyle: {
          fontSize: 14,
        },
      })
      activePlayer.push(data[i].active_player)
      betNum.push(data[i].number_bettors)
    }

    return {
      xAxisData: thirtyDays,
      seriesData: [
        {
          data: activePlayer,
          type: 'bar',
          name: t('textDaysActivePlayers'),
          barWidth: '20%',
        },
        {
          data: betNum,
          type: 'bar',
          name: t('textBetPeopleNumbers'),
          barWidth: '20%',
        },
      ],
    }
  }

  async function getAllDaysWinLoseData(searchParams) {
    const resp = await api.getIntervalTotalScoreData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    chartTable.AllDaysWinLoseChart = []

    const data = resp.data.data
    const legendData = [],
      seriesData = []

    Object.entries(data).forEach(([key, value], index) => {
      legendData.push({
        name: t(`game__${key}`),
        selected: true,
      })
      seriesData.push({
        data: [],
        type: 'line',
        name: t(`game__${key}`),
        symbol: 'circle',
        symbolSize: 10,
        lineStyle: {
          width: 3,
        },
      })

      chartTable.AllDaysWinLoseChart.push({
        name: t(`game__${key}`),
        data: [],
      })

      const gameData = value.sort((a, b) => a.log_time - b.log_time)

      // Notice: gameData會多回傳一筆資料所以要扣除 (昨日前24筆，今日後24筆)
      for (let i = 0; i < gameData.length - 1; i++) {
        let score = round(
          gameData[i].total_ya_score -
            gameData[i].total_de_score +
            gameData[i].total_tax_score -
            gameData[i].total_bonus,
          4
        )

        if (i >= 24) {
          seriesData[index].data.push(score)
        }

        chartTable.AllDaysWinLoseChart[index].data.push({
          value: score,
          isTime: i - 24 <= new Date().getHours(),
        })
      }
    })

    return { legendData, seriesData }
  }

  async function getAllDaysPlayersData(searchParams) {
    const resp = await api.getIntervalTotalBettorInfoData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    chartTable.AllDaysPlayersChart = []

    const data = resp.data.data
    const legendData = [],
      seriesData = []

    Object.entries(data).forEach(([key, value], index) => {
      legendData.push({
        name: t(`game__${key}`),
        selected: true,
      })
      seriesData.push({
        data: [],
        type: 'line',
        name: t(`game__${key}`),
        symbol: 'circle',
        symbolSize: 10,
        lineStyle: {
          width: 3,
        },
      })
      chartTable.AllDaysPlayersChart.push({
        name: t(`game__${key}`),
        data: [],
      })

      const gameData = value.sort((a, b) => a.log_time - b.log_time)

      // Notice: gameData會多回傳一筆資料所以要扣除 (昨日前24筆，今日後24筆)
      for (let i = 0; i < gameData.length - 1; i++) {
        if (i >= 24) {
          seriesData[index].data.push(gameData[i].total_bettor)
        }

        chartTable.AllDaysPlayersChart[index].data.push({
          value: gameData[i].total_bettor,
          isTime: i - 24 <= new Date().getHours(),
        })
      }
    })

    return { legendData, seriesData }
  }

  async function getRealTimeUserData(searchParams) {
    const resp = await api.getIntervalRealTimeUserData(searchParams)

    if (resp.data.code !== constant.ErrorCode.Success) {
      warn(t(`errorCode__${resp.data.code}`))
      return
    }

    const data = resp.data.data
    const logTime = []
    const yesterdayData = []
    const todayData = []

    const now = startOfMinute(new Date())
    const timezoneOffset = now.getTimezoneOffset() * -1

    for (let i = 0; i < data.length - 2; i++) {
      const dataLogTime = addMinutes(parse(data[i].log_time, 'yyyyMMddHHmm', new Date()), timezoneOffset)

      if (i < Math.floor(data.length / 2)) {
        logTime.push(dataLogTime)
        yesterdayData.push(data[i].total_online_game_user_count)
      } else if (now - dataLogTime > 0) {
        todayData.push(data[i].total_online_game_user_count)
      }
    }

    return { logTime, yesterdayData, todayData }
  }

  function setChartDialog(chart, target, chartStr) {
    dialogChart = echarts.init(document.getElementById('dialogChart'), null, {
      width: 'auto',
      height: '500px',
    })
    dialogChart.setOption(baseOption)

    dialogTitle.value = chart.getOption().title[0].text.slice(2)

    const xAxis = chart.getOption().xAxis[0].data
    // 此為昨日即今日的小時資料(0-23昨日,24-47今日)
    const chartSeries = chartTable[chartStr].find((data) => data.name === target)

    dialogChart.setOption({
      grid: {
        top: 50,
        right: 0,
        left: 60,
      },
      legend: {
        textStyle: {
          color: '#333',
        },
        top: 0,
      },
      xAxis: {
        data: xAxis,
      },
      yAxis: {
        max: function (value) {
          return Math.floor(value.max) + 50
        },
        min: function (value) {
          return value.min < 0 ? Math.floor(value.min) - 50 : 0
        },
        minInterval: 10,
        splitNumber: 5,
      },
      tooltip: {
        trigger: 'item',
        axisPointer: {
          type: 'none',
        },
        formatter: function (params) {
          return t('fmtTextChartTooltips', [`"${params.seriesName}"`, `"${params.name}"`, numberToStr(params.value)])
        },
      },
      series: [
        {
          data: chartSeries.data.slice(0, 24),
          type: 'line',
          name: t('textYesterday'),
          symbol: 'circle',
          symbolSize: 10,
          lineStyle: {
            width: 3,
          },
        },
        {
          data: chartSeries.data.slice(24, 48),
          type: 'line',
          name: t('inputTimeType__day'),
          symbol: 'circle',
          symbolSize: 10,
          lineStyle: {
            width: 3,
          },
        },
      ],
      animation: false,
    })

    const tableData = dialogChart.getOption().series
    chartTable.dialogChart = tableData.map((series) => {
      return { name: series.name, data: series.data }
    })
  }

  function setTimeType() {
    timesText.current = t(`inputTimeType__${formInput.timeType}`)
    switch (formInput.timeType) {
      case 'day':
        timesText.than = t('textThanYesterday')
        break
      case 'week':
        timesText.than = t('textThanWeek')
        break
      case 'month':
        timesText.than = t('textThanMonth')
        break
    }
  }

  function resizeCharts() {
    chartItems.AllDaysPlayersChart.resize()
    chartItems.AllDaysWinLoseChart.resize()
    chartItems.thirtyDaysChart.resize()
    chartItems.RealTimeUserCountChart.resize()
    if (isDeviceLocationInfoEnabled) {
      pieChartItems.AllDaysDeviceChart.resize()
    }
  }

  let timerId
  onMounted(async () => {
    await searchDataOverview()

    timerId = window.setInterval(searchDataOverview, 60 * 1000)
    window.addEventListener('resize', resizeCharts)
  })

  onUnmounted(() => {
    window.clearInterval(timerId)
    window.removeEventListener('resize', resizeCharts)
  })

  return {
    chartTable,
    dialogTitle,
    formInput,
    hours,
    isShowIsSearchAll,
    showChartDialog,
    stats,
    tableData,
    timesText,
    timeTypes,
    searchDataOverview,
    isDeviceLocationInfoEnabled,
  }
}
