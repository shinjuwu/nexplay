import constant from '@/base/common/constant'
import { menuItemKey } from '@/base/common/menuConstant'

export function getRoleMenu(roleMenuFolders, roleGroups, rootPermissions, targetPermissions, userStoreToRefs) {
  const user = userStoreToRefs.value
  let roleMenu = generateRoleMenu(roleMenuFolders, roleGroups, rootPermissions, targetPermissions, user)

  for (const folder of roleMenu) {
    let folderChecked = true
    for (const nextFolder of folder.items) {
      let nextFolderChecked = true

      for (const nextNextFolder of nextFolder.items) {
        nextNextFolder.checked = nextNextFolder.items.every((role) => role.checked)
        nextFolderChecked = nextFolderChecked && nextNextFolder.checked
      }

      nextFolder.checked = nextFolderChecked
      folderChecked = folderChecked && nextFolder.checked
    }

    folder.checked = folderChecked
  }

  return roleMenu
}

function generateRoleMenu(roleMenuFolders, roleGroups, rootPermissions, targetPermissions, user) {
  let resp = []

  for (const folder of roleMenuFolders) {
    let m = { key: folder.key, nameKey: `${folder.nameKeyPrefix}${folder.key}`, items: [], checked: false, show: true }

    if (folder.childFolders) {
      let items = generateRoleMenu(folder.childFolders, roleGroups, rootPermissions, targetPermissions, user)
      if (items.length > 0) {
        m.items = m.items.concat(items)
      }
    } else {
      for (const roleGroup of roleGroups) {
        // 不同資料夾跳過
        if (roleGroup.folderKey !== folder.key) {
          continue
        }

        if (user && user.accountType !== constant.AccountType.Admin) {
          if (user.cooperation === constant.AgentCooperation.Trust) {
            // 信用代理不用顯示後台代理上下分及代理分數紀錄
            const trustAgentIgnorePage = [menuItemKey.BackendUpdateAgentWallet, menuItemKey.AgentWalletLedger]
            if (trustAgentIgnorePage.indexOf(roleGroup.folderKey) >= 0) {
              continue
            }
          }
          if (user.walletType === constant.AgentWallet.Single) {
            // 單一錢包代理不用顯示後台玩家上下分
            const singleWalletAgentIgnorePage = [menuItemKey.BackendUpdateGameUserWallet]
            if (singleWalletAgentIgnorePage.indexOf(roleGroup.folderKey) >= 0) {
              continue
            }
          }
        }

        for (const item of roleGroup.items) {
          // 比對root權限，root權限沒有則跳過
          if (!item.permissions.every((r) => rootPermissions.find((fc) => fc === r))) {
            continue
          }

          m.items.push({
            folderNameKey: `${folder.nameKeyPrefix}${folder.key}`,
            key: item.key,
            nameKey: `${item.nameKeyPrefix}${item.nameKey}`,
            permissions: item.permissions,
            checked: item.permissions.every((r) => targetPermissions.find((fc) => fc === r)),
          })
        }
      }
    }

    if (m.items.length > 0) {
      resp.push(m)
    }
  }

  return resp
}

export function roleMenuToRoles(roleMenuFolders) {
  let resp = {}

  for (const item of roleMenuFolders) {
    if (item.items) {
      Object.assign(resp, roleMenuToRoles(item.items))
    } else {
      resp[item.key] = item
    }
  }

  return resp
}
