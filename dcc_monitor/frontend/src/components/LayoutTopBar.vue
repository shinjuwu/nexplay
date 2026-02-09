<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/store/userStore'

const props = defineProps({
  isShowRightAside: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['toggleRightAside'])

const route = useRoute()
const router = useRouter()
const isHomePage = computed(() => route.path === '/')

function backToHomePage() {
  router.push('/')
}

const { signOut } = useUserStore()
</script>

<template>
  <nav
    class="fixed inset-x-0 top-0 z-20 flex bg-green-600 text-white transition-transform"
    :class="{ 'translate-x-[-250px]': props.isShowRightAside }"
  >
    <ul class="flex items-center">
      <li>
        <a class="block cursor-pointer px-4 py-2" @click="isHomePage ? signOut() : backToHomePage()">
          <span v-if="isHomePage">登出</span>
          <font-awesome-icon v-else icon="fa-solid fa-house" />
        </a>
      </li>
    </ul>
    <ul class="ml-auto flex items-center">
      <li>
        <a class="block cursor-pointer px-4 py-2" @click="emit('toggleRightAside')">
          <font-awesome-icon icon="fa-solid fa-bars" />
        </a>
      </li>
    </ul>
  </nav>
</template>
