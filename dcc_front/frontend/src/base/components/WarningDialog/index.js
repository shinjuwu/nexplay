import { createVNode, render } from 'vue'

import WarningDialog from '@/base/components/WarningDialog/WarningDialog.vue'

export default {
  install(app) {
    let instance

    function warningDialogFunc(message, options = {}) {
      if (!instance) {
        let vNode = createVNode(WarningDialog)
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

    app.provide('warn', warningDialogFunc)
  },
}
