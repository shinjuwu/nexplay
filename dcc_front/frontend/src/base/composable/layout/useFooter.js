import time from '@/base/utils/time'

export function useFooter() {
  /**  __APP_VERSION__ 是由vite進行寫入值，所以此處將eslint檢察關閉*/
  /* eslint-disable */
  const appVersion = __APP_VERSION__
  /* eslint-enable */

  return {
    appVersion,
    year: time.year,
  }
}
