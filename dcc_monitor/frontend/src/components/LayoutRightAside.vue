<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/userStore'

const props = defineProps({
  isShow: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(['toggleIsShow'])

const router = useRouter()

function redirectToHome() {
  router.push('/')
}

const { signOut } = useUserStore()
</script>

<template>
  <aside
    class="fixed inset-y-0 flex w-[250px] flex-col bg-white transition-[right]"
    :class="props.isShow ? 'right-0' : 'right-[-250px]'"
  >
    <div class="flex justify-between bg-green-700 text-white">
      <div class="cursor-pointer px-4 py-2" @click="emit('toggleIsShow')">
        <font-awesome-icon icon="fa-solid fa-arrow-right" />
      </div>
      <div class="cursor-pointer px-4 py-2" @click="redirectToHome()">
        <font-awesome-icon icon="fa-solid fa-house" />
      </div>
    </div>

    <div class="flex border-b">
      <div class="flex cursor-pointer items-center px-4 py-2 text-yellow-600" @click="signOut()">
        <font-awesome-icon icon="fa-solid fa-arrow-right-from-bracket" class="mr-1 w-5" />登出
      </div>
    </div>

    <nav class="flex-1 overflow-y-auto">
      <ul>
        <li class="border-b">
          <router-link v-slot="{ href, isActive, navigate }" to="/personal-info" custom>
            <a :href="href" class="block px-4 py-2" :class="{ 'text-green-700': isActive }" @click="navigate">
              <div class="flex items-center">
                <font-awesome-icon icon="fa-regular fa-address-card" class="mr-1 w-5" />我的帐户
              </div>
            </a>
          </router-link>
        </li>
      </ul>
    </nav>
  </aside>
</template>
