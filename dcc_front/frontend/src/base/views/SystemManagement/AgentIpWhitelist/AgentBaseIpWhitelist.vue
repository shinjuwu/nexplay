<template>
  <PageTableDialog :visible="props.visible" @close="emit('close')">
    <template #header> {{ props.title }} </template>
    <template #default>
      <div class="tbl-container max-h-[500px] overflow-y-auto lg:max-h-[680px]">
        <table class="tbl">
          <thead>
            <tr>
              <th>{{ t('textCreatingDate') }}</th>
              <th>{{ t('textIpAddress') }}</th>
              <th>{{ t('textRemark') }}</th>
              <th>{{ t('textCreator') }}</th>
              <th v-if="props.mode === 'edit'"></th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(ipAddressInfo, index) in agentIpWhitelist.items"
              :key="`ipAddressInfo_${props.agentId}_${index}`"
            >
              <td>{{ ipAddressInfo.createTimeStr }}</td>
              <td>
                <input
                  v-model="ipAddressInfo.ipAddress"
                  class="form-input w-auto"
                  type="text"
                  :disabled="props.mode === 'view'"
                />
              </td>
              <td>
                <input
                  v-model="ipAddressInfo.info"
                  class="form-input w-auto"
                  type="text"
                  :disabled="props.mode === 'view'"
                  maxlength="30"
                />
              </td>
              <td>{{ ipAddressInfo.creator }}</td>
              <td v-if="props.mode === 'edit'">
                <DeleteButton @click="deleteAgentIpWhitelist(index)" />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light" @click="emit('close')">
          {{ t('textCancel') }}
        </button>
        <div v-show="props.mode === 'edit'" class="flex">
          <button
            type="button"
            class="btn btn-danger ml-2"
            :disabled="agentIpWhitelist.items.length >= 30"
            @click="createAgentIpWhitelist()"
          >
            {{ t('textIncrease') }}
          </button>
          <LoadingButton
            class="btn btn-primary ml-2"
            :parent-disabled="agentIpWhitelist.items.length <= 0"
            :is-get-data="props.visible"
            :parent-data="agentIpWhitelist.items"
            :button-click="updateAgentIpWhitelist"
          >
            {{ t('textOnSave') }}
          </LoadingButton>
        </div>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useAgentBaseIpWhitelist } from '@/base/composable/systemManagement/agentIpWhitelist/useAgentBaseIpWhitelist'

import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import DeleteButton from '@/base/components/Button/DeleteButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  mode: {
    type: String,
    default: 'view',
  },
  title: {
    type: String,
    default: '',
  },
  agentId: {
    type: Number,
    default: -1,
  },
  ipAddresses: {
    type: Array,
    default: () => [],
  },
})
const emit = defineEmits(['close', 'update'])

const { t } = useI18n()
const { agentIpWhitelist, createAgentIpWhitelist, deleteAgentIpWhitelist, updateAgentIpWhitelist } =
  useAgentBaseIpWhitelist(props, emit)
</script>
