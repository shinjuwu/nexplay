import { ref, reactive, inject, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'
import { useDropdownListStore } from '@/base/store/dropdownStore'
import { BaseTableInput } from '@/base/common/table/tableInput'
import { numberToStr } from '@/base/utils/formatNumber'

export function useAgentAccount() {
  const { t } = useI18n()
  const warn = inject('warn')

  const uStore = useUserStore()
  const { user } = storeToRefs(uStore)

  const selectRecord = reactive({ id: constant.Agent.All, checked: false, userName: '' })
  const isSettingEditEnabled = ref(false)
  const records = reactive({ items: [] })

  const formInput = reactive({
    agent: {},
  })
  const tableInput = reactive(new BaseTableInput(constant.TableDefaultLength, 'id', constant.TableSortDirection.Asc))
  const dialog = reactive({
    addAgent: false, // 新增代理
    checkKeys: false, // 查看密鑰
    agentSet: false, // 操作設定
    resetPassword: false, // 重置密碼
  })

  const isEnabled = computed(() => {
    const { isInRole } = uStore
    return {
      AgentAccountCreate: isInRole(roleItemKey.AgentAccountCreate),
      AgentAccountSecretKeyRead: isInRole(roleItemKey.AgentAccountSecretKeyRead),
      AgentAccountUpdate: isInRole(roleItemKey.AgentAccountUpdate),
      ResetPassword: isInRole(roleItemKey.ResetPassword),
    }
  })
  const isAmountColumnEnabled = computed(() => {
    return (
      user.value.accountType === constant.AccountType.Admin ||
      user.value.cooperation === constant.AgentCooperation.BuyPoint
    )
  })

  function showDialog(target) {
    dialog[target] = true
  }

  function changeSelect(record) {
    records.items.forEach((r) => {
      if (r.id !== record.id) {
        r.checked = false
      } else {
        r.checked = !r.checked
        selectRecord.id = r.checked ? r.id : constant.Agent.All
        selectRecord.checked = r.checked
      }
    })
  }

  function recordSetting(record, isEditMode) {
    selectRecord.id = record.id
    isSettingEditEnabled.value = isEditMode
    showDialog('agentSet')
  }

  async function searchRecords() {
    tableInput.showProcessing = true

    try {
      const resp = await api.getAgentList()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      records.items = data
        .map((d) => {
          return {
            id: d.id,
            code: d.code,
            levelCode: d.level_code,
            owner: d.owner_name,
            userName: d.admin_user_name,
            agentName: d.name,
            topAgentId: d.top_agent_id,
            topAgentName: d.top_agent_name,
            nickName: d.name,
            role: d.role_name,
            memberCount: d.member_count,
            ratio: d.commission,
            currency: d.currency,
            isEnabled: d.is_enabled,
            createTime: d.create_time,
            updateTime: d.update_time,
            amount: d.amount,
            amountStr: d.cooperation === constant.AgentCooperation.BuyPoint ? numberToStr(d.amount) : '-',
            cooperation: d.cooperation,
            checked: false,
          }
        })
        .filter((r) => formInput.agent.id === constant.Agent.All || formInput.agent.id === r.id)

      const { getAgentList } = useDropdownListStore()
      await getAgentList()

      tableInput.start = 0
      tableInput.draw = 0
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

  function isAgentEditEnabled(agent) {
    return agent.topAgentId === user.value.agentId
  }

  function resetPassword(row) {
    selectRecord.userName = row.userName
    dialog.resetPassword = true
  }

  return {
    dialog,
    formInput,
    isAmountColumnEnabled,
    isEnabled,
    isSettingEditEnabled,
    records,
    selectRecord,
    tableInput,
    changeSelect,
    isAgentEditEnabled,
    showDialog,
    searchRecords,
    recordSetting,
    resetPassword,
  }
}
