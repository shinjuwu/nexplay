<template>
  <PageTabs v-model="activePageTab" :tabs="pageTabs" class="rounded-t" />
  <div>
    <DefaultCannedLanguage v-show="activePageTab === pageTabs[0]" />
    <CustomeCannedLanguage v-show="activePageTab === pageTabs[1]" />
  </div>
</template>

<script setup>
import { useGameCannedLanguage } from '@/base/composable/gameSetting/gameCannedLanguage/useGameCannedLanguage'

import PageTabs from '@/base/components/Page/Tabs/PageTabs.vue'
import DefaultCannedLanguage from '@/base/views/GameSetting/GameCannedLanguage/DefaultCannedLanguage.vue'
import CustomeCannedLanguage from '@/base/views/GameSetting/GameCannedLanguage/CustomeCannedLanguage.vue'

const { pageTabs, activePageTab } = useGameCannedLanguage()
</script>
