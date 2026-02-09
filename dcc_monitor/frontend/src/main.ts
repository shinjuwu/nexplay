import '@/assets/css/style.css'

import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router/index'
import store from '@/store/index'
import * as fontawsome from '@/fontawesome'

const app = createApp(App)

fontawsome.registerVueComponent(app)

app.use(router).use(store).mount('#app')

export { app }
