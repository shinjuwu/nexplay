<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/store/userStore'
import { useServiceStatusStore } from '@/store/serviceStatusStore'

import HomeGroupMenu from '@/components/HomeGroupMenu.vue'
import HomeGroupMenuItem from '@/components/HomeGroupMenuItem.vue'

const { user } = storeToRefs(useUserStore())
const { serviceStatus } = storeToRefs(useServiceStatusStore())
</script>

<template>
  <HomeGroupMenu v-if="user.isAdmin" group-name="系统管理">
    <HomeGroupMenuItem v-slot="{ iconSizeClass }" name="会员管理" link-to="/member-management">
      <font-awesome-icon icon="fa-solid fa-address-book" :class="iconSizeClass" />
    </HomeGroupMenuItem>
  </HomeGroupMenu>

  <HomeGroupMenu v-if="user.permissions.length > 0" group-name="RTP监控系统">
    <HomeGroupMenuItem
      v-for="platform in user.permissions"
      :key="`system-monitor-${platform}`"
      v-slot="{ iconSizeClass }"
      :name="`${platform.toUpperCase()}站監控`"
      :link-to="`/system-monitor/${platform}`"
      :icon-container-class="
        serviceStatus[platform] === undefined ? '' : serviceStatus[platform] ? 'bg-green-100' : 'bg-rose-100'
      "
    >
      <font-awesome-icon icon="fa-solid fa-desktop" :class="iconSizeClass" class="inset-0" />
      <font-awesome-icon icon="fa-regular fa-eye" class="absolute top-[27%] h-[27%]" />
    </HomeGroupMenuItem>
  </HomeGroupMenu>
</template>
