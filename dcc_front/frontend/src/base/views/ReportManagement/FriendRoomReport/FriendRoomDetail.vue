<template>
  <PageTableDialog :visible="props.visible" @close="close()">
    <template #header> {{ t('textFriendRoomInfo') }} </template>
    <template #default>
      <component :is="detailComponent" v-if="detailComponent" :friend-room-info="props.friendRoomInfo" />
    </template>
    <template #footer>
      <div class="ml-auto flex">
        <button type="button" class="btn btn-light mr-2" @click="close()">
          {{ t('textClose') }}
        </button>
      </div>
    </template>
  </PageTableDialog>
</template>

<script setup>
import { computed, defineAsyncComponent } from 'vue'
import { useI18n } from 'vue-i18n'
import constant from '@/base/common/constant'

import PageTableDialog from '@/base/components/Page/Table/PageTableDialog.vue'

const { t } = useI18n()

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  friendRoomInfo: {
    type: Object,
    default: () => {},
  },
})
const emit = defineEmits(['close'])

function close() {
  emit('close')
}

const FriendsTexasDetail = defineAsyncComponent(() =>
  import('@/base/views/ReportManagement/FriendRoomReport/FriendsTexasDetail/FriendsTexasDetail.vue')
)
const detailComponent = computed(() => {
  switch (props.friendRoomInfo.gameId) {
    case constant.Game.Friendstexas:
      return FriendsTexasDetail
    default:
      return null
  }
})
</script>
