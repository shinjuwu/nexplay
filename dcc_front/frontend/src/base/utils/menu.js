import constant from '@/base/common/constant'
import { menuItemKey } from '@/base/common/menuConstant'

let menu
export function getMenu(menuFolders, menuItems, userStoreToRefs) {
  const user = userStoreToRefs.value
  menu = generateMenu(menuFolders, menuItems, user.roles)

  if (user.accountType !== constant.AccountType.Admin) {
    if (user.cooperation === constant.AgentCooperation.Trust) {
      // 信用代理不用顯示後台代理上下分及代理分數紀錄
      const trustAgentIgnorePage = [menuItemKey.BackendUpdateAgentWallet, menuItemKey.AgentWalletLedger]
      for (let i = 0; i < menu.length; i++) {
        menu[i].items = menu[i].items.filter((item) => trustAgentIgnorePage.indexOf(item.key) < 0)
      }
    }
    if (user.walletType === constant.AgentWallet.Single) {
      // 單一錢包代理不用顯示後台玩家上下分
      const singleWalletAgentIgnorePage = [menuItemKey.BackendUpdateGameUserWallet]
      for (let i = 0; i < menu.length; i++) {
        menu[i].items = menu[i].items.filter((item) => singleWalletAgentIgnorePage.indexOf(item.key) < 0)
      }
    }
  }

  return menu
}

function generateMenu(menuFolders, menuItems, userRoles) {
  let resp = []
  for (const folder of menuFolders) {
    let m = { key: folder.key, icon: folder.icon, items: [], show: false }

    for (const item of menuItems) {
      // 不同資料夾跳過
      if (item.folderKey !== folder.key) {
        continue
      }

      // 需要權限，使用者沒有權限跳過
      if (!userRoles[item.role]) {
        continue
      }

      m.items.push({
        key: item.key,
        forceUpdateCount: 0,
        props: item.createProps(),
      })
    }

    if (m.items.length > 0) {
      resp.push(m)
    }
  }
  return resp
}

export function getMenuItemFromMenu(key) {
  for (const folder of menu) {
    for (const item of folder.items) {
      if (item.key === key) {
        return item
      }
    }
  }
  return null
}
