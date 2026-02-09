<template>
  <router-view v-if="isOnBeforeMountDone" />
</template>

<script setup>
import { useApp } from '@/base/composable/useApp'

const { isOnBeforeMountDone } = useApp()
</script>

<style lang="css">
@import '@/base/assets/css/index.css';
</style>
