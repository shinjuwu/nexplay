import * as token from '@/base/utils/token'

export const whitelist = ['/login']

export function setRouterCheckPermissions(router, whitelist) {
  router.beforeEach((to, from, next) => {
    const path = to.path.toLowerCase()
    const hasToken = token.get() !== null
    // 沒有token
    if (!hasToken) {
      if (whitelist.indexOf(path) !== -1) {
        // 不需要登入的頁面
        next()
      } else {
        // 需要登入的頁面導到登入頁
        next(`/Login?redirect=${to.path}`)
      }

      // TO DO:清除用戶資料(避免有殘餘的token資訊)
      return
    }

    // 登入頁面有token則導到首頁
    if (path.toLowerCase() === '/login') {
      router.push('/')
    } else {
      // TO DO: 檢查有沒有權限
      next()
    }
  })
}
