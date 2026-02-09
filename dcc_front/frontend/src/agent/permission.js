import { setRouterCheckPermissions, whitelist } from '@/base/permission'
import router from '@/agent/router/index'

setRouterCheckPermissions(router, whitelist)
