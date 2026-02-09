import { defineStore } from 'pinia'
import { reactive, ref, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import * as api from '@/base/api/sysGlobal'
import constant from '@/base/common/constant'
import axios from 'axios'

export const useDropdownListStore = defineStore('dropdownList', () => {
  const warn = inject('warn')
  const { t } = useI18n()
  const isReady = ref(false)

  const dropdownList = reactive({
    agents: [],
    agentPermissions: {},
    allGames: [],
    games: [],
    tags: [],
  })

  async function getAgentList() {
    try {
      const resp = await api.getAgentList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(`errorCode__${resp.data.code}`)
        return
      }

      const data = resp.data.data
      dropdownList.agents = data
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

  async function getAgentPermissionList(accountType) {
    try {
      const resp = await api.getAgentPermissionList(accountType)

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(`errorCode__${resp.data.code}`)
        return
      }

      dropdownList.agentPermissions[accountType] = resp.data.data.map((i) => {
        return {
          id: i.id,
          name: i.name,
        }
      })
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
  async function getAllGameList() {
    if (isReady.value) {
      return
    }

    try {
      const resp = await api.getAllGameList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(`errorCode__${resp.data.code}`)
        return
      }

      const data = resp.data.data
      dropdownList.allGames = data.map((gameId) => {
        return {
          id: gameId,
          name: t(`game__${gameId}`),
        }
      })

      isReady.value = true
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

  async function getGameList() {
    try {
      const resp = await api.getGameList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(`errorCode__${resp.data.code}`)
        return
      }
      const data = resp.data.data
      dropdownList.games = data.map((gameId) => {
        return {
          id: gameId,
          name: t(`game__${gameId}`),
        }
      })
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

  async function getAgentTagsList() {
    try {
      const resp = await api.getAgentCustomTagSettingList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(`errorCode__${resp.data.code}`)
        return
      }
      const data = resp.data.data

      const defTags = Object.values(data.DefaultTagList).map((value) => {
        return {
          index: value.idx,
          name: t(`riskType__${value.idx + 4}`),
        }
      })

      const customTags = []
      if (data.custom_tag_info) {
        Object.values(data.custom_tag_info).forEach((tagInfo) => {
          customTags.push({
            index: tagInfo.idx,
            name: tagInfo.name,
            info: tagInfo.info,
            bgColor: '',
            txtColor: '',
          })

          if (tagInfo.name !== '') {
            const colorObj = JSON.parse(tagInfo.color)
            customTags.bgColor = colorObj.bg_color
            customTags.txtColor = colorObj.txt_color
          }
        })
      }

      dropdownList.tags = defTags.concat(customTags)
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

  return {
    getAgentList,
    getAgentPermissionList,
    getAllGameList,
    getGameList,
    getAgentTagsList,
    dropdownList,
  }
})
