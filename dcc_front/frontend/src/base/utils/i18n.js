import i18n from '@/base/i18n'
import storage from '@/base/utils/storage'

export async function setLocale(locale) {
  const { default: message } = await import(`@/base/locales/${locale}.json`)
  i18n.global.setLocaleMessage(locale, message)
  i18n.global.locale.value = locale
  storage.local.set(storage.keys.LOCALE, locale)
}
