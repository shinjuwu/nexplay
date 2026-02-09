import { reactive, ref, computed, watch, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import * as api from '@/base/api/sysManage'
import constant from '@/base/common/constant'
import axios from 'axios'
import time from '@/base/utils/time'
import * as validate from '@/base/utils/validate'

export function useMarqueeSettingComponent(props, emit) {
  const { t } = useI18n()
  const warn = inject('warn')

  const visible = ref(false)

  const typeOptions = reactive({
    items: Object.values(constant.MarqueeType)
      .filter((t) => t !== constant.MarqueeType.All)
      .map((i) => ({
        value: i,
        label: t(`typeCode__${i}`),
      })),
  })
  const langOptions = reactive({
    items: Object.values(constant.LangType)
      .filter((l) => l !== constant.LangType.All)
      .map((l) => ({
        value: l,
        label: t(`langType__${l}`),
      })),
  })

  const submitText = computed(() => (props.mode === 'create' ? t('textIncrease') : t('textOnSave')))
  const disabled = computed(() => props.mode === 'view')

  const marqueeInfo = reactive({
    id: props.marqueeInfo.id,
    type: props.marqueeInfo.type,
    lang: props.marqueeInfo.lang,
    order: props.marqueeInfo.order,
    freq: props.marqueeInfo.freq,
    createDate: props.marqueeInfo.createDate,
    startDate: props.marqueeInfo.startDate,
    endDate: props.marqueeInfo.endDate,
    content: props.marqueeInfo.content,
    isEnabled: props.marqueeInfo.isEnabled,
  })

  function validateForm() {
    const errors = []

    if (!marqueeInfo.content || marqueeInfo.content.length > 80) {
      errors.push(t('textTextAreaErrorText'))
    }

    if (!validate.numberMinValidate(marqueeInfo.freq, 1)) {
      errors.push(t('fmtTextMarqueeFreqMinError', [1]))
    }

    if (!validate.numberRangeValidate(marqueeInfo.order, 1, 9999)) {
      errors.push(t('fmtTextMarqueeOrderRangeError', [1, 9999]))
    }

    return errors
  }

  async function createMarquee() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const resp = await api.createMarquee({
        content: marqueeInfo.content,
        type: marqueeInfo.type,
        order: marqueeInfo.order,
        lang: marqueeInfo.lang,
        freq: marqueeInfo.freq,
        start_time: time.localTimeToUtcTimeSting(marqueeInfo.startDate),
        end_time: time.localTimeToUtcTimeSting(marqueeInfo.endDate),
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        close()
        emit('refreshMarqueeList')
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

  async function updateMarquee() {
    const errors = validateForm()
    if (errors.length > 0) {
      warn(errors.join(t('symbolComma')))
      return
    }

    try {
      const resp = await api.updateMarquee({
        id: marqueeInfo.id,
        content: marqueeInfo.content,
        type: marqueeInfo.type,
        order: marqueeInfo.order,
        lang: marqueeInfo.lang,
        freq: marqueeInfo.freq,
        start_time: time.localTimeToUtcTimeSting(marqueeInfo.startDate),
        end_time: time.localTimeToUtcTimeSting(marqueeInfo.endDate),
        is_enabled: marqueeInfo.isEnabled,
      })

      warn(t(`errorCode__${resp.data.code}`)).then(() => {
        if (resp.data.code !== constant.ErrorCode.Success) {
          return
        }
        close()
        emit('refreshMarqueeList')
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

  function close() {
    visible.value = false
    emit('close')
  }

  watch(props, async (newValue) => {
    if (newValue) {
      marqueeInfo.id = props.marqueeInfo.id
      marqueeInfo.type = props.marqueeInfo.type
      marqueeInfo.lang = props.marqueeInfo.lang
      marqueeInfo.order = props.marqueeInfo.order
      marqueeInfo.freq = props.marqueeInfo.freq
      marqueeInfo.createDate = props.marqueeInfo.createDate
      marqueeInfo.startDate = new Date(props.marqueeInfo.startDate)
      marqueeInfo.endDate = new Date(props.marqueeInfo.endDate)
      marqueeInfo.content = props.marqueeInfo.content
      marqueeInfo.isEnabled = props.marqueeInfo.isEnabled
    }
    visible.value = props.visible
  })

  return {
    typeOptions,
    langOptions,
    visible,
    marqueeInfo,
    submitText,
    disabled,
    createMarquee,
    updateMarquee,
  }
}
