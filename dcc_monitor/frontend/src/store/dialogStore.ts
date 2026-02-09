import type { ConfirmDialogInstance, DialogOptions, DialogShowFunc, WarnDialogInstance } from '@/types/types.dialog'

import { createVNode, render } from 'vue'
import { defineStore } from 'pinia'

import { app } from '@/main'
import WarningDialog from '@/components/WarningDialog.vue'
import ConfirmDialog from '@/components/ConfirmDialog.vue'

export const useDialogStore = defineStore('dialog', () => {
  let totalDialogs = 0

  function increaseTotalDialogs() {
    totalDialogs++
    if (totalDialogs > 0) {
      document.body.classList.add('overflow-hidden')
    }
  }

  function decreaseTotalDialogs() {
    totalDialogs--
    if (totalDialogs === 0) {
      document.body.classList.remove('overflow-hidden')
    }
  }

  function warn(message: string, options: DialogOptions = {}) {
    const vNode = createVNode(WarningDialog)
    vNode.appContext = app._context

    const el = document.createElement('div')
    render(vNode, el)

    increaseTotalDialogs()

    const instance = vNode.component as WarnDialogInstance
    const show = instance.exposed?.show as DialogShowFunc
    show(message, options)

    return new Promise<void>((resolve) => {
      instance.callback = () => {
        resolve()
        render(null, el)
        decreaseTotalDialogs()
      }
    })
  }

  function confirm(message: string, options: DialogOptions = {}) {
    const vNode = createVNode(ConfirmDialog)
    vNode.appContext = app._context

    const el = document.createElement('div')
    render(vNode, el)

    increaseTotalDialogs()

    const instance = vNode.component as ConfirmDialogInstance
    const show = instance.exposed?.show as DialogShowFunc
    show(message, options)

    return new Promise<void>((resolve) => {
      instance.confirmCallback = resolve
      instance.closeCallback = () => {
        render(null, el)
        decreaseTotalDialogs()
      }
    })
  }

  return {
    warn,
    confirm,
  }
})
