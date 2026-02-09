import { inject, reactive, ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysUser'
import constant from '@/base/common/constant'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'
import { BaseTableInput } from '@/base/common/table/tableInput'

export function useBackendAccount() {
  const { t } = useI18n()
  const warn = inject('warn')

  const isEnabled = computed(() => {
    const { isInRole } = useUserStore()
    return {
      BackendAccountCreate: isInRole(roleItemKey.BackendAccountCreate),
      BackendAccountUpdate: isInRole(roleItemKey.BackendAccountUpdate),
      ResetPassword: isInRole(roleItemKey.ResetPassword),
    }
  })

  const formInput = reactive({
    account: '',
  })
  const tableInput = reactive(
    new BaseTableInput(constant.TableDefaultLength, 'userName', constant.TableSortDirection.Asc)
  )

  const records = reactive({ items: [] })
  const selectRecord = reactive({ userName: '', parentAccount: '', role: '' })
  const isSettingEditEnabled = ref(false)

  const dialog = reactive({
    addSubAccounts: false, // 新增子帳號
    setting: false, // 設定
    resetPassword: false, // 重置密碼
  })

  function accountSetting(row, isEditMode) {
    isSettingEditEnabled.value = isEditMode
    selectRecord.userName = row.userName
    selectRecord.parentAccount = row.parentAccount
    selectRecord.role = row.role
    dialog.setting = true
  }

  async function searchRecords() {
    tableInput.showProcessing = true
    try {
      const resp = await api.getAdminUsers()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      records.items = data
        .map((d) => {
          return {
            userName: d.username,
            nickName: d.nickname,
            parentAccount: d.top_username,
            role: d.role_name,
            state: d.id_enabled,
            createTime: d.create_time,
            lastLoginTime: d.login_time,
          }
        })
        .filter((r) => {
          if (!formInput.account) return r
          else return r.userName.includes(formInput.account)
        })

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

  function resetPassword(row) {
    selectRecord.userName = row.userName
    dialog.resetPassword = true
  }

  return {
    dialog,
    formInput,
    isEnabled,
    isSettingEditEnabled,
    records,
    selectRecord,
    tableInput,
    searchRecords,
    accountSetting,
    resetPassword,
  }
}
