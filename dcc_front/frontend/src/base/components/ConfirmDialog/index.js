import { createVNode, render } from 'vue'

import ConfirmDialog from '@/base/components/ConfirmDialog/ConfirmDialog.vue'

export default {
  install(app) {
    let instance

    function confirmDialogFunc(message, options = {}) {
      if (!instance) {
        let vNode = createVNode(ConfirmDialog)
        vNode.appContext = app._context

        let el = document.createElement('div')
        render(vNode, el)

        document.body.append(vNode.el)

        instance = vNode.component
      }

      instance.exposed.show(message, options)

      return new Promise((resolve) => {
        instance.callback = () => resolve()
      })
    }

    app.provide('confirm', confirmDialogFunc)
  },
}
