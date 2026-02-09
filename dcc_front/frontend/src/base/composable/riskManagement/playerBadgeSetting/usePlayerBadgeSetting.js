import { computed, inject, onBeforeMount, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'
import * as api from '@/base/api/sysRiskControl'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { useUserStore } from '@/base/store/userStore'
import { useDropdownListStore } from '@/base/store/dropdownStore'
import { defaultBadgeInfo } from '@/base/common/badge'

export function usePlayerBadgeSetting() {
  const warn = inject('warn')

  const { t } = useI18n()

  const myAgentId = computed(() => {
    const { user } = storeToRefs(useUserStore())
    return user.value.agentId
  })
  const isSettingEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return isInRole(roleItemKey.PlayerBadgeSettingUpdate)
  })

  const defaultBadges = computed(() => {
    return Object.values(defaultBadgeInfo)
      .map(
        (badge) =>
          new Badges(
            t(`riskType__${badge.Index + 4}`),
            badge.BackgroundColor,
            badge.TextColor,
            t(`textDefaultBadgeInfoRiskType__${badge.Index + 4}`),
            badge.Index,
            false
          )
      )
      .sort((a, b) => a.index - b.index)
  })
  const customBadges = reactive({
    items: [
      new Badges('', '#34d399', '#ffffff', '', 0, true),
      new Badges('', '#34d399', '#ffffff', '', 1, true),
      new Badges('', '#818cf8', '#ffffff', '', 2, true),
      new Badges('', '#818cf8', '#ffffff', '', 3, true),
      new Badges('', '#06b6d4', '#ffffff', '', 4, true),
      new Badges('', '#06b6d4', '#ffffff', '', 5, true),
      new Badges('', '#6b7280', '#ffffff', '', 6, true),
      new Badges('', '#6b7280', '#ffffff', '', 7, true),
    ],
  })
  const badges = computed(() => {
    return [...defaultBadges.value, ...customBadges.items]
  })

  const tableInput = ref(new BaseTableInput(badges.value.length, 'idx', constant.TableSortDirection.Asc))

  async function getAgentCustomTagSettingList() {
    try {
      tableInput.value.showProcessing = true

      const resp = await api.getAgentCustomTagSettingList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      Object.values(resp.data.data.custom_tag_info).forEach((tagInfo) => {
        const customBadge = customBadges.items[tagInfo.idx]
        customBadge.name = tagInfo.name
        customBadge.info = tagInfo.info

        if (tagInfo.color != '') {
          const colorObj = JSON.parse(tagInfo.color)
          customBadge.bgColor = colorObj.bg_color
          customBadge.txtColor = colorObj.txt_color
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
    } finally {
      tableInput.value.showProcessing = false
    }
  }

  async function setAgentCustomTagSettingList() {
    try {
      tableInput.value.showProcessing = true

      const input = {
        agent_id: myAgentId.value,
        custom_tag_info: customBadges.items.reduce((resp, badge) => {
          resp[badge.index.toString()] = {
            color: badge.name !== '' ? `{"bg_color":"${badge.bgColor}","txt_color":"${badge.txtColor}"}` : '',
            idx: badge.index,
            info: badge.name !== '' ? badge.info : '',
            name: badge.name,
          }
          return resp
        }, {}),
      }

      const resp = await api.setAgentCustomTagSettingList(input)

      warn(t(`errorCode__${resp.data.code}`)).then(async () => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        const { getAgentTagsList } = useDropdownListStore()
        await getAgentTagsList()
      })
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      tableInput.value.showProcessing = false
    }
  }

  async function resetAllCustomBadges() {
    customBadges.items.forEach((item) => item.resetDefault())
  }

  onBeforeMount(async () => {
    await getAgentCustomTagSettingList()
  })

  return {
    badges,
    isSettingEnabled,
    tableInput,
    resetAllCustomBadges,
    setAgentCustomTagSettingList,
  }
}

class Badges {
  constructor(name, bgColor, txtColor, info, index, isEditable) {
    this.name = name
    this.bgColor = bgColor
    this.defaultBgColor = bgColor
    this.txtColor = txtColor
    this.defaultTxtColor = txtColor
    this.info = info
    this.index = index
    this.isEditable = isEditable
  }

  resetDefault() {
    this.name = ''
    this.bgColor = this.defaultBgColor
    this.txtColor = this.defaultTxtColor
    this.info = ''
  }
}
