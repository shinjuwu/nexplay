<template>
  <nav
    class="fixed inset-x-0 top-0 z-10 flex items-center bg-white pr-4 md:ml-[4.6rem]"
    :class="{ 'md:ml-[250px]': props.isSidebarOpen }"
  >
    <ul class="flex">
      <li>
        <a class="block cursor-pointer px-4 py-2 text-gray-700" @click="emit('toggleSidebar')">
          <Bars3Icon class="h-6 w-6" />
        </a>
      </li>
    </ul>
    <ul class="ml-auto flex cursor-pointer items-center">
      <li v-if="hasBackendAnnoucementRole" class="p-2">
        <a class="relative align-middle" @click="redirectToBackendAnnoucement()">
          <BellIcon class="h-8 w-6 text-gray-400" />
          <span
            v-if="hasNewAnnoucements"
            class="absolute top-0 right-0 inline-block h-2 w-2 rounded-full bg-red-600"
          ></span>
        </a>
      </li>
      <li class="relative ml-4 cursor-pointer bg-gray-100 p-2">
        <a
          ref="userInfoToggleEl"
          class="flex h-10 items-center justify-between"
          href="#"
          @click="isShowUserInfoList = !isShowUserInfoList"
        >
          <UserCircleIcon class="mr-2 h-8 w-8 text-gray-400" />
          <div class="text-gray-400">{{ userName }}</div>
        </a>
        <transition>
          <ul
            v-show="isShowUserInfoList"
            ref="userInfoListEl"
            tabindex="-1"
            class="after-arrow after-arrow-up absolute right-0 block w-32 translate-y-4 rounded border-2 border-gray-400 bg-white after:right-[10px]"
            @blur="onBlurUserInfoList"
          >
            <li class="text-gray-500 hover:bg-gray-300 hover:text-black">
              <a href="#" class="flex items-center p-2" @click="emit('showPersonalInfo')" @blur="onBlurUserInfoList">
                <UserIcon class="mr-2 h-4 w-4" />
                {{ t('textMyAccount') }}
              </a>
            </li>
            <li class="text-gray-500 hover:bg-gray-300 hover:text-black">
              <a href="#" class="flex items-center p-2" @click="userSignOut()" @blur="onBlurUserInfoList">
                <ArrowLeftOnRectangleIcon class="mr-2 h-4 w-4" />
                {{ t('textSafeSignOut') }}
              </a>
            </li>
          </ul>
        </transition>
      </li>
    </ul>
  </nav>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useHeader } from '@/base/composable/layout/useHeader'
import { Bars3Icon, BellIcon, UserIcon, UserCircleIcon, ArrowLeftOnRectangleIcon } from '@heroicons/vue/24/outline'

const { t } = useI18n()

const props = defineProps({
  isSidebarOpen: {
    type: Boolean,
    default: true,
  },
})
const emit = defineEmits(['toggleSidebar', 'showPersonalInfo'])

const {
  hasBackendAnnoucementRole,
  hasNewAnnoucements,
  isShowUserInfoList,
  userInfoListEl,
  userInfoToggleEl,
  userName,
  onBlurUserInfoList,
  redirectToBackendAnnoucement,
  userSignOut,
} = useHeader()
</script>

<style lang="scss">
.v-enter-active,
.v-leave-active {
  transform: translateY(1rem);
  transition: transform 0.2s ease;
}

.v-enter-from,
.v-leave-to {
  transform: translateY(1.25rem);
  opacity: 0;
}
</style>
