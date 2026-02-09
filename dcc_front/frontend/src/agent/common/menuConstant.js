import { defineAsyncComponent } from 'vue'
import {
  CalendarDaysIcon,
  ChartBarSquareIcon,
  Cog6ToothIcon,
  ExclamationTriangleIcon,
  UsersIcon,
  CurrencyDollarIcon,
} from '@heroicons/vue/24/outline'
import {
  MenuFolder,
  menuFolderKey,
  MenuItem,
  menuItemKey,
  GameLogParseMenuItem,
  PlayerAccountMenuItem,
  PlayerLogMenuItem,
  roleItemKey,
} from '@/base/common/menuConstant'

export const menuFolders = [
  // #運營管理
  new MenuFolder({
    key: menuFolderKey.OperationManagement,
    icon: CalendarDaysIcon,
  }),
  // #系統管理
  new MenuFolder({
    key: menuFolderKey.SystemManagement,
    icon: Cog6ToothIcon,
  }),
  // #代理帳號管理
  new MenuFolder({
    key: menuFolderKey.AgentAccountManagement,
    icon: UsersIcon,
  }),
  // #報表管理
  new MenuFolder({
    key: menuFolderKey.ReportManagement,
    icon: ChartBarSquareIcon,
  }),
  // #風控管理
  new MenuFolder({
    key: menuFolderKey.RiskManagement,
    icon: ExclamationTriangleIcon,
  }),
  // #Jackpot功能管理
  new MenuFolder({
    key: menuFolderKey.JackpotManagement,
    icon: CurrencyDollarIcon,
  }),
]

export const menuItems = [
  // #運營管理
  // 數據總覽
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.DataOverview,
    role: roleItemKey.DataOverviewRead,
    component: defineAsyncComponent(() => import('@/base/views/OperationManagement/DataOverview/DataOverview.vue')),
  }),
  // 後台公告
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.BackendAnnouncement,
    role: roleItemKey.BackendAnnouncementRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/BackendAnnouncement/BackendAnnouncement.vue')
    ),
  }),
  // 玩家帳號
  new PlayerAccountMenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.PlayerAccount,
    role: roleItemKey.PlayerAccountRead,
    component: defineAsyncComponent(() => import('@/base/views/OperationManagement/PlayerAccount/PlayerAccount.vue')),
  }),
  // 遊戲日誌解析
  new GameLogParseMenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.GameLogParse,
    role: roleItemKey.GameLogParseRead,
    component: defineAsyncComponent(() => import('@/base/views/OperationManagement/GameLogParse/GameLogParse.vue')),
  }),
  //後台代理上下分
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.BackendUpdateAgentWallet,
    role: roleItemKey.BackendUpdateAgentWalletRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/BackendAgentWalletLedger/BackendAgentWalletLedger.vue')
    ),
  }),
  //後台玩家上下分
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.BackendUpdateGameUserWallet,
    role: roleItemKey.BackendUpdateGameUserWalletRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/BackendGameUserWalletLedger/BackendGameUserWalletLedger.vue')
    ),
  }),
  // 後台操作紀錄
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.BackendActionLog,
    role: roleItemKey.BackendActionLogRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/BackendActionLog/BackendActionLog.vue')
    ),
  }),

  // #系統管理
  // 後台IP白名單
  new MenuItem({
    folderKey: menuFolderKey.SystemManagement,
    key: menuItemKey.AgentIpWhitelist,
    role: roleItemKey.AgentIpWhitelistRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/SystemManagement/AgentIpWhitelist/AgentIpWhitelist.vue')
    ),
  }),
  // 後台帳號
  new MenuItem({
    folderKey: menuFolderKey.SystemManagement,
    key: menuItemKey.BackendAccount,
    role: roleItemKey.BackendAccountRead,
    component: defineAsyncComponent(() => import('@/base/views/SystemManagement/BackendAccount/BackendAccount.vue')),
  }),
  // 權限群組管理
  new MenuItem({
    folderKey: menuFolderKey.SystemManagement,
    key: menuItemKey.GroupRoleManagement,
    role: roleItemKey.GroupRoleManagementRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/SystemManagement/GroupRoleManagement/GroupRoleManagement.vue')
    ),
  }),
  // 後台登入紀錄
  new MenuItem({
    folderKey: menuFolderKey.SystemManagement,
    key: menuItemKey.BackendLoginLog,
    role: roleItemKey.BackendLoginLogRead,
    component: defineAsyncComponent(() => import('@/base/views/SystemManagement/BackendLoginLog/BackendLoginLog.vue')),
  }),

  // #代理帳號管理
  // 代理帳號
  new MenuItem({
    folderKey: menuFolderKey.AgentAccountManagement,
    key: menuItemKey.AgentAccount,
    role: roleItemKey.AgentAccountRead,
    component: defineAsyncComponent(() => import('@/base/views/AgentAccountManagement/AgentAccount/AgentAccount.vue')),
  }),

  // #報表管理
  // 輸贏報表
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.WinLoseReport,
    role: roleItemKey.WinLoseReportRead,
    component: defineAsyncComponent(() => import('@/base/views/ReportManagement/WinLoseReport/WinLoseReport.vue')),
  }),
  // 業績報表
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.EarningReport,
    role: roleItemKey.EarningReportRead,
    component: defineAsyncComponent(() => import('@/base/views/ReportManagement/EarningReport/EarningReport.vue')),
  }),
  // 日結算報表
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.DailySettlementReport,
    role: roleItemKey.DailySettlementReportRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/ReportManagement/DailySettlementReport/DailySettlementReport.vue')
    ),
  }),
  // 玩家遊玩紀錄
  new PlayerLogMenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.PlayerLog,
    role: roleItemKey.PlayerLogRead,
    component: defineAsyncComponent(() => import('@/base/views/ReportManagement/PlayerLog/PlayerLog.vue')),
  }),
  // 代理分數紀錄
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.AgentWalletLedger,
    role: roleItemKey.AgentWalletLedgerRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/ReportManagement/AgentWalletLedger/AgentWalletLedger.vue')
    ),
  }),
  // 玩家分數紀錄
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.WalletLedger,
    role: roleItemKey.WalletLedgerRead,
    component: defineAsyncComponent(() => import('@/base/views/ReportManagement/WalletLedger/WalletLedger.vue')),
  }),
  // 好友房建房紀錄
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.FriendRoomReport,
    role: roleItemKey.FriendRoomReportRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/ReportManagement/FriendRoomReport/FriendRoomReport.vue')
    ),
  }),

  // #風控功能
  // 玩家標示
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.PlayerBadge,
    role: roleItemKey.PlayerBadgeRead,
    component: defineAsyncComponent(() => import('@/base/views/RiskManagement/PlayerBadge/PlayerBadge.vue')),
  }),
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.PlayerBadgeSetting,
    role: roleItemKey.PlayerBadgeSettingRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/RiskManagement/PlayerBadgeSetting/PlayerBadgeSetting.vue')
    ),
  }),
  // #Jackpot功能管理
  // Jackpot中獎紀錄
  new MenuItem({
    folderKey: menuFolderKey.JackpotManagement,
    key: menuItemKey.JackpotPrizeRecord,
    role: roleItemKey.JackpotPrizeRecordRead,
    component: defineAsyncComponent(() => import('@/base/views/JackpotManagement/JackpotPrizeRecord.vue')),
  }),
]

export const menuItemMap = menuItems.reduce((obj, item) => {
  obj[item.key] = item
  return obj
}, {})
