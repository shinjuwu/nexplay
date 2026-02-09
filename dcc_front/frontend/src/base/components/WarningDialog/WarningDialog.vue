<template>
  <div v-show="open" class="fixed inset-0 z-10">
    <div class="relative z-10 m-2 bg-white sm:my-7 sm:mx-auto sm:max-w-[500px]">
      <div class="bg-indigo-400 p-4 text-xl font-bold text-white">
        <h5 class="modal-title">{{ title }}</h5>
      </div>
      <p class="max-h-96 min-h-[7rem] overflow-y-auto whitespace-pre-line bg-white p-4 text-lg text-gray-600">
        {{ message }}
      </p>
      <hr />
      <div class="p-4 text-right">
        <button type="button" class="btn btn-primary" @click="close()">
          {{ t('textClose') }}
        </button>
      </div>
    </div>
    <div class="fixed top-0 h-screen w-full bg-black/50" @click="close()"></div>
  </div>
</template>

<script setup>
import { getCurrentInstance, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const open = ref(false)
const title = ref('')
const message = ref('')

const instance = getCurrentInstance()

function show(msg, options) {
  open.value = true
  title.value = options.title || t('textFriendlyReminder')
  message.value = msg
}

function close() {
  instance.callback()
  open.value = false
}

defineExpose({
  show,
})
</script>
