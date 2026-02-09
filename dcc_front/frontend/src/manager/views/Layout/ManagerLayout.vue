<template>
  <BaseLayout :menu-folders="menuFolders" :menu-items="menuItems" :menu-item-map="menuItemMap">
    <slot></slot>
  </BaseLayout>
</template>

<script setup>
import { menuFolders, menuItems, menuItemMap } from '@/manager/common/menuConstant'
import BaseLayout from '@/base/views/Layout/BaseLayout.vue'
</script>
