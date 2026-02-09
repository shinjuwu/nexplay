import { createI18n } from 'vue-i18n'

const i18n = createI18n({
  legacy: false, // must set `false`, to use Composition API
  globalInjection: true,
})

export default i18n
