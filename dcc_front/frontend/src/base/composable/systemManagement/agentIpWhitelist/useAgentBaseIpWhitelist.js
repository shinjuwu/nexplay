import { inject, reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'

export function useAgentBaseIpWhitelist(props, emit) {
  const warn = inject('warn')
  const { t } = useI18n()

  const agentIpWhitelist = reactive({ items: [] })

  function createAgentIpWhitelist() {
    agentIpWhitelist.items.push({
      createTime: 0,
      createTimeStr: '',
      ipAddress: '',
      info: '',
      creator: '',
    })
  }

  function deleteAgentIpWhitelist(index) {
    agentIpWhitelist.items.splice(index, 1)
  }

  function validateDuplicateIpWhitelist() {
    const tmpCache = {}
    for (const item of agentIpWhitelist.items) {
      if (tmpCache[item.ipAddress]) {
        return false
      }
      tmpCache[item.ipAddress] = true
    }
    return true
  }

  function validateEmptyIpWhitelist() {
    for (const item of agentIpWhitelist.items) {
      if (item.ipAddress) {
        return true
      }
    }
    return false
  }

  function updateAgentIpWhitelist() {
    if (!validateDuplicateIpWhitelist()) {
      warn(t('textIpAddressDuplicate'))
      return
    }

    if (!validateEmptyIpWhitelist()) {
      warn(t('textIpAddressEmpty'))
      return
    }

    emit('update', agentIpWhitelist.items)
  }

  watch(
    () => props.ipAddresses,
    () => {
      agentIpWhitelist.items = props.ipAddresses.slice()
    }
  )

  return {
    agentIpWhitelist,
    createAgentIpWhitelist,
    deleteAgentIpWhitelist,
    updateAgentIpWhitelist,
  }
}
