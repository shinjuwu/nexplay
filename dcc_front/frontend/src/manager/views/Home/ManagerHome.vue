<template>
  <ManagerLayout>
    <BaseHome :menu-item-map="menuItemMap" />
  </ManagerLayout>
</template>

<script setup>
import { menuItemMap } from '@/manager/common/menuConstant'
import BaseHome from '@/base/views/Home/BaseHome.vue'
import ManagerLayout from '@/manager/views/Layout/ManagerLayout.vue'
</script>
