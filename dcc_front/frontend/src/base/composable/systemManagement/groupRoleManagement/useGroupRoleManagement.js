import { inject, reactive, computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import constant from '@/base/common/constant'
import * as api from '@/base/api/sysAgent'
import { roleItemKey } from '@/base/common/menuConstant'
import { useUserStore } from '@/base/store/userStore'
import { storeToRefs } from 'pinia'
import { BaseTableInput } from '@/base/common/table/tableInput'

export function useGroupRoleManagement() {
  const { t } = useI18n()

  const warn = inject('warn')

  const { isInRole } = useUserStore()

  const showMode = ref('')
  const showDialog = ref(false)
  const deleteDialog = ref(false)

  const accountType = computed(() => {
    const { user } = storeToRefs(useUserStore())
    return user.value.accountType
  })
  const records = reactive({ items: [] })

  const isRead = {
    GroupRoleManagementCreate: !isInRole(roleItemKey.GroupRoleManagementCreate),
    GroupRoleManagementUpdate: !isInRole(roleItemKey.GroupRoleManagementUpdate),
    GroupRoleManagementDelete: !isInRole(roleItemKey.GroupRoleManagementDelete),
  }

  const tableInput = reactive(new BaseTableInput(constant.TableDefaultLength, 'name', constant.TableSortDirection.Asc))
  const formInput = reactive({
    groupRole: '',
  })

  const agentPermission = reactive({
    id: '',
    name: '',
    accountType: accountType.value,
    info: '',
    permissions: [],
  })

  function createNewAgentPermission() {
    showDialog.value = true
    showMode.value = 'create'
    Object.assign(agentPermission, {
      id: '',
      name: '',
      accountType: accountType.value,
      info: '',
      permissions: [],
    })
  }

  async function searchRecords() {
    tableInput.showProcessing = true

    try {
      const resp = await api.getAgentPermissionList()

      records.items = resp.data.data
        .map((d) => {
          return {
            id: d.id,
            name: d.name,
            level: d.account_type,
            remark: d.info,
            permissions: d.permissions,
          }
        })
        .filter((r) => {
          if (!formInput.groupRole) return r
          else return r.name.includes(formInput.groupRole)
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

  async function getAgentPermissionInfo(agentPermissionId, mode) {
    try {
      const resp = await api.getAgentPermission({
        id: agentPermissionId,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      const data = resp.data.data
      agentPermission.id = data.id
      agentPermission.accountType = data.account_type
      agentPermission.info = data.info
      agentPermission.name = data.name
      agentPermission.permissions = data.permissions

      showMode.value = mode
      if (mode !== 'delete') {
        showDialog.value = true
      } else {
        deleteDialog.value = true
      }
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

  async function deleteAgentPermission() {
    try {
      const resp = await api.deleteAgentPermission({
        id: agentPermission.id,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        deleteDialog.value = false
        searchRecords()
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

  return {
    t,
    records,
    isRead,
    formInput,
    tableInput,
    accountType,
    showMode,
    showDialog,
    deleteDialog,
    agentPermission,
    searchRecords,
    createNewAgentPermission,
    getAgentPermissionInfo,
    deleteAgentPermission,
  }
}
