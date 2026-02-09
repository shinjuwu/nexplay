import { setRouterCheckPermissions, whitelist } from '@/base/permission'
import router from '@/manager/router/index'

setRouterCheckPermissions(router, whitelist)
