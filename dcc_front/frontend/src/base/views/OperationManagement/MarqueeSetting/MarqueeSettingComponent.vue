<template>
  <PageTableDialog :visible="visible" @close="emit('close')">
    <template #header>
      <div>{{ t('menuItemMarqueeSetting') }}</div>
    </template>
    <template #default>
      <div class="grid grid-cols-4 gap-4">
        <div class="col-span-1">
          <label class="form-label required-mark">{{ t('textMarqueeType') }}</label>
          <select v-model="marqueeInfo.type" class="form-input" :disabled="disabled">
            <option
              v-for="typeOption in typeOptions.items"
              :key="typeOption.value"
              :label="typeOption.label"
              :value="typeOption.value"
            ></option>
          </select>
        </div>
        <div class="col-span-1">
          <label class="form-label required-mark">{{ t('textLanguage') }}</label>
          <select v-model="marqueeInfo.lang" class="form-input" :disabled="disabled">
            <option
              v-for="lang in langOptions.items"
              :key="lang.value"
              :label="lang.label"
              :value="lang.value"
            ></option>
          </select>
        </div>
        <div class="col-span-1">
          <label class="form-label flex items-center">
            <span class="required-mark">
              {{ t('textPriority') }}
            </span>
            <ButtonTooltips class="ml-1" :tips-text="t('textMarqueePriorityTooltips')">
              <template #content>
                <QuestionMarkCircleIcon class="text-info h-5 w-5" />
              </template>
            </ButtonTooltips>
          </label>
          <input v-model="marqueeInfo.order" type="number" min="1" max="9999" class="form-input" :disabled="disabled" />
        </div>
        <div class="col-span-1">
          <label class="form-label flex items-center">
            <span class="required-mark">
              {{ t('textDisplayFrequency') }}
            </span>
            <ButtonTooltips class="ml-1" :tips-text="t('textMarqueeFrequencyTooltips')">
              <template #content>
                <QuestionMarkCircleIcon class="text-info h-5 w-5" />
              </template>
            </ButtonTooltips>
          </label>
          <input v-model="marqueeInfo.freq" type="number" min="1" class="form-input" :disabled="disabled" />
        </div>
        <div v-if="marqueeInfo.createDate" class="col-span-1">
          <label class="form-label mb-0.5 mr-2">{{ t('textIsStart') }}</label>
          <CheckButton
            class="align-text-top"
            :checked="marqueeInfo.isEnabled"
            :disabled="disabled"
            @click="marqueeInfo.isEnabled = !marqueeInfo.isEnabled"
          />
        </div>
        <div v-if="marqueeInfo.createDate" class="col-span-3">
          <label class="form-label mr-2">{{ t('textCreatingDate') }}</label>
          <span class="text-gray-500">{{ time.utcTimeStrToLocalTimeFormat(marqueeInfo.createDate) }}</span>
        </div>
        <div class="col-span-2">
          <label class="form-label">{{ t('textStartDate') }}</label>
          <FormDateTimeInput
            v-model="marqueeInfo.startDate"
            :minute-increment="time.commonReportTimeMinuteIncrement"
            :before-days="time.commonReportTimeBeforeDays"
            :disabled="disabled"
          />
        </div>
        <div class="col-span-2">
          <label class="form-label">{{ t('textEndDate') }}</label>
          <FormDateTimeInput
            v-model="marqueeInfo.endDate"
            calendar-align="right"
            :minute-increment="time.commonReportTimeMinuteIncrement"
            :before-days="time.commonReportTimeBeforeDays"
            :disabled="disabled"
          />
        </div>
      </div>
      <div class="col-12 my-2">
        <label class="form-label required-mark">{{ t('textContentText') }}</label>
        <input
          v-model="marqueeInfo.content"
          class="form-input"
          type="text"
          maxlength="80"
          :disabled="disabled"
          :placeholder="t('fmtPlaceHolderTextAreaLimit', [80])"
        />
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="emit('close')">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          v-if="!disabled"
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="marqueeInfo"
          :button-click="() => (props.mode === 'create' ? createMarquee() : updateMarquee())"
        >
          {{ submitText }}
        </LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { QuestionMarkCircleIcon } from '@heroicons/vue/24/outline'
import { useMarqueeSettingComponent } from '@/base/composable/operationManagement/marqueeSetting/useMarqueeSettingComponent'
import time from '@/base/utils/time'

import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import FormDateTimeInput from '@/base/components/Form/FormDateTimeInput.vue'
import CheckButton from '@/base/components/Button/CheckButton.vue'
import ButtonTooltips from '@/base/components/Button/ButtonTooltips.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  marqueeInfo: {
    type: Object,
    default: () => {},
  },
  mode: {
    type: String,
    default: '',
  },
})
const emit = defineEmits(['close', 'refreshMarqueeList'])

const { t } = useI18n()
const { typeOptions, langOptions, visible, marqueeInfo, submitText, disabled, createMarquee, updateMarquee } =
  useMarqueeSettingComponent(props, emit)
</script>
