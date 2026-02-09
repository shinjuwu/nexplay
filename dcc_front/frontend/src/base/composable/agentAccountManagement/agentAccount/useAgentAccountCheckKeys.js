import { ref, reactive, watch, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysAgent'
import constant from '@/base/common/constant'

export function useAgentAccountCheckKeys(props, emit) {
  const warn = inject('warn')
  const { t } = useI18n()

  const show = ref(false)
  const agent = reactive({
    id: constant.Agent.All,
    aeskey: '',
    md5key: '',
  })

  function close() {
    show.value = false
    emit('close', false)
  }

  async function searchAgentSecretKey() {
    try {
      const resp = await api.getAgentSecretKey({ id: props.selectRecord.id })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        close()
        return
      }
      const data = resp.data.data
      agent.aeskey = data.aeskey
      agent.md5key = data.md5key
      agent.id = props.selectRecord.id
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

  watch(props, async () => {
    if (props.visible && props.selectRecord.id !== constant.Agent.All) {
      await searchAgentSecretKey()
    }
    show.value = props.visible
  })

  return {
    t,
    show,
    agent,
    close,
  }
}
