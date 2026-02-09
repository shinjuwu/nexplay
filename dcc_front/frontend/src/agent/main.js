import { createApp } from 'vue'
import App from '@/base/App.vue'
import i18n from '@/base/i18n'
import store from '@/base/store/index'
import router from '@/agent/router/index'

import ConfirmDialog from '@/base/components/ConfirmDialog/index'
import WarningDialog from '@/base/components/WarningDialog/index'

import '@/agent/permission'

const app = createApp(App)

app.use(router).use(i18n).use(store).use(ConfirmDialog).use(WarningDialog).mount('#app')
