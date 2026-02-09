import { computed, inject, reactive, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import { storeToRefs } from 'pinia'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'
import { roleMenuFolders, roleGroups, rolePermissions } from '@/base/common/menuConstant'
import { getRoleMenu } from '@/base/utils/groupRole'
import { useUserStore } from '@/base/store/userStore'

export function useGroupRoleManagementComponent(props, emit) {
  const { t } = useI18n()
  const warn = inject('warn')
  const confirm = inject('confirm')

  const visible = ref(false)

  const { user } = storeToRefs(useUserStore())

  const agentPermission = reactive({
    id: '',
    name: '',
    accountType: 0,
    info: '',
    permissions: [],
  })

  const roleMenu = reactive({
    items: [],
  })

  const disabled = computed(() => props.mode === 'view')

  const contentText = computed(() => {
    let title, submit

    if (props.mode === 'create') {
      title = t('roleItemGroupRoleManagementCreate')
      submit = t('textIncrease')
    } else {
      title = t('roleItemGroupRoleManagementUpdate')
      submit = t('textOnSave')
    }

    return { title, submit }
  })

  const accountTypes = computed(() => {
    const items = []
    for (let i = user.value.accountType; i <= user.value.accountType + 1 && i <= constant.AccountType.Nornam; i++) {
      items.push({
        value: i,
        label: i === user.value.accountType ? t(`textBackendAccount`) : t(`accountType__${i}`),
      })
    }
    return items
  })

  async function updateRoleMenu(accountType) {
    try {
      const resp = await api.getAgentPermissionTemplateInfo()

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(`errorCode__${resp.data.code}`)
        return
      }

      const data = resp.data.data

      roleMenu.items = getRoleMenu(
        roleMenuFolders,
        roleGroups,
        data[accountType],
        Array.from(agentPermission.permissions),
        user
      )
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
  function validateForm() {
    const errors = []

    if (!agentPermission.name) {
      errors.push(t('textGroupRoleNameRequired'))
    }

    return errors
  }

  async function createAgentPermission() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const resp = await api.createAgentPermission({
        account_type: agentPermission.accountType,
        info: agentPermission.info,
        name: agentPermission.name,
        permissions: Array.from(agentPermission.permissions),
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }

        emit('refreshAgentPermissionList')
        emit('close')
        agentPermission.name = ''
        agentPermission.accountType = user.value.accountType
        agentPermission.info = ''
        agentPermission.permissions = []
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

  async function updateAgentPermission() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    confirm(t('textUpdateGroupRoleReminder1')).then(async () => {
      try {
        const resp = await api.setAgentPermission({
          id: agentPermission.id,
          account_type: agentPermission.accountType,
          info: agentPermission.info,
          name: agentPermission.name,
          permissions: Array.from(agentPermission.permissions),
        })

        warn(t(`errorCode__${resp.data.code}`)).then(() => {
          if (resp.data.code !== constant.ErrorCode.Success) {
            return
          }

          emit('refreshAgentPermissionList')
          emit('close')
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
    })
  }

  function updatePermissions(roleMenuOrRoleItem) {
    if (roleMenuOrRoleItem.items) {
      for (const item of roleMenuOrRoleItem.items) {
        updatePermissions(item)
      }
    } else {
      for (const permission of roleMenuOrRoleItem.permissions) {
        if (roleMenuOrRoleItem.checked) {
          agentPermission.permissions.add(permission)
        } else {
          agentPermission.permissions.delete(permission)
        }
      }
      updateAllPermissions(roleMenu, agentPermission.permissions)
    }
  }

  function updateAllPermissions(roleMenuOrRoleItem, permission) {
    if (roleMenuOrRoleItem.items) {
      for (const item of roleMenuOrRoleItem.items) {
        updateAllPermissions(item, permission)
      }
    } else {
      const agentPermissionArr = Array.from(permission)
      for (const permission of roleMenuOrRoleItem.permissions) {
        if (agentPermissionArr.includes(permission)) {
          roleMenuOrRoleItem.checked = true
        } else {
          roleMenuOrRoleItem.checked = false
        }
      }
    }
  }

  function close() {
    emit('refreshAgentPermissionList')
    emit('close')
  }

  watch(
    () => props.visible,
    async (newValue) => {
      if (newValue) {
        agentPermission.id = props.agentPermission.id
        agentPermission.name = props.agentPermission.name
        agentPermission.accountType = props.agentPermission.accountType
        agentPermission.info = props.agentPermission.info
        agentPermission.permissions = new Set(props.agentPermission.permissions.concat(rolePermissions.BackendBasic))

        await updateRoleMenu(agentPermission.accountType)
      }

      visible.value = newValue
    }
  )

  watch(
    () => agentPermission.accountType,
    async (newValue) => {
      await updateRoleMenu(newValue)
    }
  )

  return {
    t,
    visible,
    agentPermission,
    roleMenu,
    accountTypes,
    contentText,
    disabled,
    updatePermissions,
    createAgentPermission,
    updateAgentPermission,
    close,
  }
}
