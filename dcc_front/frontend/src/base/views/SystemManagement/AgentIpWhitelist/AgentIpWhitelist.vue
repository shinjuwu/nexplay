<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show">
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textAgent') }}</label>
          <FormAgentListDropdown v-model="formAgentIpWhitelistListInput.agent" class="mb-1 w-full md:w-3/12" />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchAgentIpWhitelistInfos()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageTable
        :records="agentIpWhitelistList.items"
        :total-records="agentIpWhitelistList.items.length"
        :table-input="tableAgentIpWhitelistListInput"
      >
        <template
          #default="{
            currentRecords,
            pageLength,
            recordStart,
            totalPages,
            totalRecords,
            isSortIconActive,
            lengthChange,
            pageChange,
            sorting,
          }"
        >
          <div class="tbl-container">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <PageTableSortableTh
                    column="adminUsername"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textAgentAccount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="name" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textAgentName') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="ipAddressCount"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textBackendIpAddressCount') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh
                    column="apiIpAddressCount"
                    :is-sort-icon-active="isSortIconActive"
                    @sorting="sorting"
                  >
                    <template #default> {{ t('textApiIpAddressCount') }} </template>
                  </PageTableSortableTh>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr class="no-hover">
                    <td colspan="5">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record__${record.id}`">
                    <td>{{ record.adminUsername }}</td>
                    <td>{{ record.name }}</td>
                    <td>{{ record.ipAddressCount }}</td>
                    <td>{{ record.apiIpAddressCount }}</td>
                    <td class="space-x-2">
                      <ViewButton
                        class="text-success"
                        :tips-text="t('textViewBackendIp')"
                        @click="showBackendIpDialog(record.id, record.name, 'view')"
                      />
                      <EditButton
                        v-if="isEditEnabled"
                        class="text-success"
                        :tips-text="t('textSettingBackendIp')"
                        @click="showBackendIpDialog(record.id, record.name, 'edit')"
                      />
                      <ViewButton
                        :tips-text="t('textViewApiIp')"
                        class="text-danger"
                        @click="showApiIpDialog(record.id, record.name, 'view')"
                      />
                      <EditButton
                        v-if="isEditEnabled"
                        class="text-danger"
                        :tips-text="t('textSettingApiIp')"
                        @click="showApiIpDialog(record.id, record.name, 'edit')"
                      />
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
          <div class="mt-6 text-center text-slate-500 md:flex md:flex-wrap md:items-center">
            <PageTableMenuLength class="mb-1.5 md:mr-6" :display-length="pageLength" @length-change="lengthChange" />
            <PageTableInfo
              class="mb-1.5 md:mr-6"
              :display-records="currentRecords.length"
              :record-start="recordStart"
              :total-records="totalRecords"
            />
            <PageTableQuickPage
              class="mb-1.5"
              :page-length="pageLength"
              :record-start="recordStart"
              :total-pages="totalPages"
              @page-change="pageChange"
            />
            <PageTablePagination
              class="mb-1.5"
              :page-length="pageLength"
              :record-start="recordStart"
              :total-pages="totalPages"
              @page-change="pageChange"
            />
          </div>
        </template>
      </PageTable>
    </template>
  </ToggleHeader>
  <AgentBackendIpWhitelist
    :agent-id="backendIpInfo.agentId"
    :agent-name="backendIpInfo.agentName"
    :mode="backendIpInfo.mode"
    :visible="backendIpInfo.visible"
    @close="backendIpInfo.visible = false"
    @refresh-ip-list="searchAgentIpWhitelistInfos()"
  />
  <AgentApiIpWhitelist
    :agent-id="apiIpInfo.agentId"
    :agent-name="apiIpInfo.agentName"
    :mode="apiIpInfo.mode"
    :visible="apiIpInfo.visible"
    @close="apiIpInfo.visible = false"
    @refresh-ip-list="searchAgentIpWhitelistInfos()"
  />
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useAgentIpWhitelist } from '@/base/composable/systemManagement/agentIpWhitelist/useAgentIpWhitelist'

import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import FormAgentListDropdown from '@/base/components/Form/Dropdown/FormAgentListDropdown.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import AgentBackendIpWhitelist from '@/base/views/SystemManagement/AgentIpWhitelist/AgentBackendIpWhitelist.vue'
import AgentApiIpWhitelist from '@/base/views/SystemManagement/AgentIpWhitelist/AgentApiIpWhitelist.vue'

const { t } = useI18n()
const {
  agentIpWhitelistList,
  apiIpInfo,
  backendIpInfo,
  formAgentIpWhitelistListInput,
  isEditEnabled,
  tableAgentIpWhitelistListInput,
  searchAgentIpWhitelistInfos,
  showBackendIpDialog,
  showApiIpDialog,
} = useAgentIpWhitelist()
</script>
