<template>
  <ToggleHeader>
    <template #default="{ show }">
      <form v-show="show" @submit.prevent>
        <div class="flex flex-wrap items-center">
          <label class="form-label mb-1 w-full pr-2 md:w-1/12 md:text-right">{{ t('textGroupRole') }}</label>
          <input
            v-model="formInput.groupRole"
            type="text"
            class="form-input mb-1 md:w-3/12"
            :placeholder="t('placeHolderTextGroupRole')"
          />
        </div>
        <div class="mt-4 flex flex-wrap items-center justify-end">
          <div class="flex flex-1">
            <button
              v-if="!isRead.AgentAccountCreate"
              type="button"
              class="btn btn-danger mb-1 flex w-full items-center justify-center md:ml-2 md:w-fit"
              @click="createNewAgentPermission()"
            >
              <UserGroupIcon class="mr-2 inline-block h-5 w-5" />
              {{ t('roleItemGroupRoleManagementCreate') }}
            </button>
          </div>
          <button
            type="button"
            class="btn btn-primary mb-1 w-full md:ml-2 md:w-2/12 xl:w-1/12"
            @click="searchRecords()"
          >
            {{ t('textSearch') }}
          </button>
        </div>
      </form>
      <hr class="my-2" />
      <PageTable :records="records.items" :total-records="records.items.length" :table-input="tableInput">
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
                  <PageTableSortableTh column="name" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textGroupRole') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="level" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textGroupRoleLevel') }} </template>
                  </PageTableSortableTh>
                  <PageTableSortableTh column="remark" :is-sort-icon-active="isSortIconActive" @sorting="sorting">
                    <template #default> {{ t('textRemark') }} </template>
                  </PageTableSortableTh>
                  <th>{{ t('textOperate') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="currentRecords.length === 0">
                  <tr>
                    <td colspan="4">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr v-for="record in currentRecords" :key="`record_${record.id}`">
                    <td>{{ record.name }}</td>
                    <td>
                      {{ record.level === accountType ? t(`textBackendAccount`) : t(`accountType__${record.level}`) }}
                    </td>
                    <td>{{ record.remark }}</td>
                    <td>
                      <ViewButton @click="getAgentPermissionInfo(record.id, 'view')" />
                      <EditButton class="mx-4" @click="getAgentPermissionInfo(record.id, 'update')" />
                      <DeleteButton @click="getAgentPermissionInfo(record.id, 'delete')" />
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
  <GroupRoleManagementComponent
    :visible="showDialog"
    :mode="showMode"
    :agent-permission="agentPermission"
    @close="showDialog = false"
    @refresh-agent-permission-list="searchRecords()"
  />
  <PageTableDialog :visible="deleteDialog" @close="deleteDialog = false">
    <template #header>
      {{ t('roleItemGroupRoleManagementDelete') }}
    </template>
    <template #default>
      <p class="mb-3">{{ t('fmtTextDeleteGroupRole', [agentPermission.name]) }}</p>
      <ul class="pl-5">
        <li style="list-style: disc" class="text-danger">{{ t('textDeleteGroupRoleReminder1') }}</li>
        <li style="list-style: disc" class="text-danger">{{ t('textDeleteGroupRoleReminder2') }}</li>
      </ul>
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="deleteDialog = false">{{ t('textCancel') }}</button>
        <LoadingButton class="btn btn-primary" :button-click="() => deleteAgentPermission()">{{
          t('textConfirm')
        }}</LoadingButton>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { useGroupRoleManagement } from '@/base/composable/systemManagement/groupRoleManagement/useGroupRoleManagement'
import { UserGroupIcon } from '@heroicons/vue/20/solid'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import PageTable from '@/base/components/Page/Table/PageTable.vue'
import PageTableMenuLength from '@/base/components/Page/Table/PageTableMenuLength.vue'
import PageTableInfo from '@/base/components/Page/Table/PageTableInfo.vue'
import PageTableQuickPage from '@/base/components/Page/Table/PageTableQuickPage.vue'
import PageTablePagination from '@/base/components/Page/Table/PageTablePagination.vue'
import PageTableSortableTh from '@/base/components/Page/Table/PageTableSortableTh.vue'
import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'
import EditButton from '@/base/components/Button/EditButton.vue'
import ViewButton from '@/base/components/Button/ViewButton.vue'
import DeleteButton from '@/base/components/Button/DeleteButton.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'
import GroupRoleManagementComponent from './GroupRoleManagementComponent.vue'

const {
  t,
  records,
  isRead,
  formInput,
  tableInput,
  accountType,
  showDialog,
  showMode,
  deleteDialog,
  agentPermission,
  searchRecords,
  createNewAgentPermission,
  getAgentPermissionInfo,
  deleteAgentPermission,
} = useGroupRoleManagement()
</script>
