<script setup lang="ts">
import type { NotificationTransactionRecord } from '@/types/types.monitorSystem'
import { ref } from 'vue'
import time from '@/utils/time'

const props = withDefaults(defineProps<{ items: NotificationTransactionRecord[] }>(), {
  items: () => [] as NotificationTransactionRecord[],
})
const selectItem = ref<NotificationTransactionRecord | null>(null)

async function copy() {
  if (selectItem.value === null) {
    return
  }
  await navigator.clipboard.writeText(selectItem.value.id)
}
</script>

<template>
  <div class="text-center">
    <h2 class="w-full text-lg font-bold">玩家上下分监控</h2>

    <div>
      <div class="max-h-80 min-h-[20rem] w-full overflow-auto">
        <table class="w-full">
          <thead>
            <tr>
              <th class="sticky top-0 bg-gray-200 px-2 py-1">时间</th>
              <th class="sticky top-0 bg-gray-200 px-2 py-1">代理</th>
              <th class="sticky top-0 bg-gray-200 px-2 py-1">玩家</th>
              <th class="sticky top-0 bg-gray-200 px-2 py-1">行为</th>
            </tr>
          </thead>
          <tbody class="cursor-pointer">
            <tr
              v-for="item in props.items"
              :key="`NotificationItem__${item.id}`"
              :class="{ 'bg-gray-100': selectItem === item }"
              @click="selectItem = item"
            >
              <td>{{ time.format(item.occurredTime) }}</td>
              <td>{{ item.agentName }}</td>
              <td>{{ item.playerName }}</td>
              <td>{{ item.info }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="flex items-center border border-gray-400">
        <div class="flex-1 bg-gray-100 py-1">{{ selectItem ? selectItem.id : '点击列表显示订单号' }}</div>
        <button
          type="button"
          class="border-l border-gray-400 bg-gray-300 px-4 py-1"
          :disabled="selectItem === null"
          @click="copy()"
        >
          复制
        </button>
      </div>
    </div>
  </div>
</template>
