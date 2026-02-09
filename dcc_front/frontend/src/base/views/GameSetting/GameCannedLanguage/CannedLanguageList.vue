<template>
  <table class="tbl tbl-hover">
    <thead>
      <tr>
        <th>{{ t('textEmojiType') }}</th>
        <th v-for="lang in cannedLangTypes" :key="`th_${lang}`">{{ t(`langType__${lang}`) }}</th>
        <th>{{ t('textState') }}</th>
      </tr>
    </thead>
    <tbody>
      <template v-if="props.cannedLanguageList.length === 0">
        <tr class="no-hover">
          <td :colspan="colNumber">{{ t('textTableEmpty') }}</td>
        </tr>
      </template>
      <template v-else>
        <tr
          v-for="cannedLanguage in props.cannedLanguageList"
          :key="`canned_language_${props.cannedLanguageListAgentId}_${cannedLanguage.serial}`"
          class="hover:bg-slate-100"
        >
          <td>
            <FormDropdown
              v-model="cannedLanguage.emotionType"
              :items="cannedEmojiTypeItems"
              :fmt-item-key="(i) => i"
              :fmt-item-text="(i) => t(`cannedEmojiType__${i}`)"
              :disabled="!props.isEditable"
            />
          </td>
          <td v-for="lang in cannedLangTypes" :key="`td_${lang}`">
            <input
              type="text"
              name="langContent"
              class="form-input min-w-[400px]"
              :disabled="!props.isEditable"
              :value="cannedLanguage.langContentMap.get(lang)"
              :maxlength="contentMaxLength(lang)"
              @input="(event) => cannedLanguage.langContentMap.set(lang, event.target.value)"
            />
          </td>
          <td class="space-x-2">
            <template v-if="props.cannedLanguageType === constant.CannedLanguageType.Default">
              <span
                :class="{
                  'text-success': cannedLanguage.status === constant.CannedStatus.Open,
                  'text-danger': cannedLanguage.status === constant.CannedStatus.Close,
                }"
              >
                {{ t(`cannedStatus__${cannedLanguage.status}`) }}
              </span>
            </template>
            <template v-else-if="props.cannedLanguageType === constant.CannedLanguageType.Custome">
              <button
                type="button"
                class="btn"
                :class="cannedLanguage.status === constant.CannedStatus.Open ? 'btn-success' : 'btn-secondary'"
                @click="cannedLanguage.status = constant.CannedStatus.Open"
              >
                {{ t(`cannedStatus__${constant.CannedStatus.Open}`) }}
              </button>
              <button
                type="button"
                class="btn"
                :class="cannedLanguage.status === constant.CannedStatus.Close ? 'btn-danger' : 'btn-secondary'"
                @click="cannedLanguage.status = constant.CannedStatus.Close"
              >
                {{ t(`cannedStatus__${constant.CannedStatus.Close}`) }}
              </button>
            </template>
          </td>
        </tr>
      </template>
    </tbody>
  </table>
</template>

<script setup>
import { useI18n } from 'vue-i18n'

import constant from '@/base/common/constant'

import FormDropdown from '@/base/components/Form/Dropdown/FormDropdown.vue'

const props = defineProps({
  cannedLanguageType: {
    type: Number,
    default: -1,
  },
  cannedLanguageList: {
    type: Array,
    default: () => [],
  },
  cannedLanguageListAgentId: {
    type: Number,
    default: -1,
  },
  isEditable: {
    type: Boolean,
    default: false,
  },
})

const { t } = useI18n()
const cannedLangTypes = (() => Object.values(constant.LangType).filter((l) => constant.CannedLangTypeSetting[l]))()
const colNumber = (() => {
  const baseCols = 3
  const langCols = cannedLangTypes.length
  return baseCols + langCols
})()
const cannedEmojiTypeItems = (() => Object.values(constant.CannedEmojiType))()

function contentMaxLength(lang) {
  switch (lang) {
    case constant.LangType.CHS:
      return 15
    default:
      return 30
  }
}
</script>
