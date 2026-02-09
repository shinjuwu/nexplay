<template>
  <PageTableDialog :visible="visible" size="xl" @close="close">
    <template #header>
      <div>
        <div>
          {{ content.title }}
          <span class="ml-3">{{ announcementInfo.createTime }}</span>
        </div>
      </div>
    </template>
    <template #default>
      <div class="grid grid-cols-4 gap-4">
        <div v-if="props.mode !== 'view'" class="col-span-4 flex flex-col md:col-span-1">
          <label class="form-label required-mark">{{ t('textPostType') }}</label>
          <select v-model="announcementInfo.type" class="form-input" :disabled="disabled">
            <option v-for="type in typeOptions" :key="type.value" :label="type.label" :value="type.value"></option>
          </select>
        </div>
        <div class="col-span-4 flex flex-col" :class="{ 'md:col-span-3': props.mode !== 'view' }">
          <label class="form-label required-mark">{{ t('textSubject') }}</label>
          <input
            v-model="announcementInfo.subject"
            class="form-input w-full"
            :disabled="disabled"
            type="text"
            maxlength="20"
            :placeholder="t('fmtPlaceHolderTextAreaLimit', [20])"
          />
        </div>
        <div class="col-span-4 flex flex-col">
          <label class="form-label required-mark">{{ t('textContentText') }}</label>
          <textarea
            v-model="announcementInfo.content"
            :disabled="disabled"
            class="form-input"
            rows="12"
            maxlength="800"
            :placeholder="t('fmtPlaceHolderTextAreaLimit', [800])"
          />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close">
          {{ t('textCancel') }}
        </button>
        <LoadingButton
          v-if="!disabled"
          class="btn btn-primary"
          :is-get-data="visible"
          :parent-data="announcementInfo"
          :button-click="() => (props.mode === 'create' ? createAnnouncement() : updateAnnouncement())"
          >{{ content.submit }}</LoadingButton
        >
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useBackendAnnouncementComponent } from '@/base/composable/operationManagement/backendAnnouncement/useBackendAnnouncementComponent'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const emit = defineEmits(['close', 'refreshAnnouncementList'])
const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  mode: {
    type: String,
    default: '',
  },
  announcementInfo: {
    type: Object,
    default: () => {},
  },
})

const { t, visible, content, typeOptions, announcementInfo, disabled, createAnnouncement, updateAnnouncement, close } =
  useBackendAnnouncementComponent(props, emit)
</script>
