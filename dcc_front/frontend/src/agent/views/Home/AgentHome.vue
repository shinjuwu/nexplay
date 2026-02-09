<template>
  <AgentLayout>
    <BaseHome :menu-item-map="menuItemMap" />
  </AgentLayout>
</template>

<script setup>
import { menuItemMap } from '@/agent/common/menuConstant'
import BaseHome from '@/base/views/Home/BaseHome.vue'
import AgentLayout from '@/agent/views/Layout/AgentLayout.vue'
</script>
