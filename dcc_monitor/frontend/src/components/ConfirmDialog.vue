<script setup lang="ts">
import { getCurrentInstance, ref } from 'vue'
import type { DialogOptions, ConfirmDialogInstance } from '@/types/types.dialog'

const visible = ref(false)
const title = ref('')
const message = ref('')
const instance = getCurrentInstance() as ConfirmDialogInstance

function show(msg: string, options: DialogOptions = {}) {
  message.value = msg
  title.value = options?.title || '温馨提示'
  visible.value = true
}

function close() {
  visible.value = false
  instance.closeCallback()
}

function confirm() {
  instance.confirmCallback()
  close()
}

defineExpose({
  show,
})
</script>

<template>
  <Teleport to="body">
    <div v-show="visible" class="fixed inset-0 z-20 p-2 sm:p-4">
      <div class="relative z-20 mx-auto bg-white sm:mt-4 sm:max-w-[500px]">
        <div class="bg-green-600 p-4 text-xl font-bold text-white">
          {{ title }}
        </div>
        <p class="h-28 overflow-y-auto break-words bg-white px-4 py-2 text-lg text-gray-600">
          {{ message }}
        </p>
        <hr />
        <div class="p-4 text-right">
          <button type="button" class="btn btn-light mr-1" @click="close()">关闭</button>
          <button type="button" class="btn btn-primary" @click="confirm()">确定</button>
        </div>
      </div>
      <div class="fixed inset-0 h-screen w-full bg-gray-200/40" @click="close()"></div>
    </div>
  </Teleport>
</template>
