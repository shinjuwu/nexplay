<template>
  <PageTabs v-model="activeTab" :tabs="gameTabs"> </PageTabs>
  <div>
    <BaiRenGames v-show="activeTab === gameTabs[0]" :game-type="constant.GameType.BaiRen" />
    <ChiPaiGames v-show="activeTab === gameTabs[1]" :game-type="constant.GameType.ChiPai" />
    <ElectronicGames v-show="activeTab === gameTabs[2]" :game-type="constant.GameType.ElectronicGame" />
    <SlotGames v-show="activeTab === gameTabs[3]" :game-type="constant.GameType.Slot" />
  </div>
</template>

<script setup>
import { useRTPSetting } from '@/base/composable/riskManagement/RTPSetting/useRTPSetting'
import PageTabs from '@/base/components/Page/Tabs/PageTabs.vue'
import BaiRenGames from '@/base/views/RiskManagement/RTPSetting/BaiRenGames.vue'
import ChiPaiGames from '@/base/views/RiskManagement/RTPSetting/ChiPaiGames.vue'
import ElectronicGames from '@/base/views/RiskManagement/RTPSetting/ElectronicGames.vue'
import SlotGames from '@/base/views/RiskManagement/RTPSetting/SlotGames.vue'
import constant from '@/base/common/constant'

const { gameTabs, activeTab } = useRTPSetting()
</script>
