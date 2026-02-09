import { defineAsyncComponent } from 'vue'
import {
  CalendarDaysIcon,
  ChartBarSquareIcon,
  Cog6ToothIcon,
  ExclamationTriangleIcon,
  PuzzlePieceIcon,
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
  // #遊戲設置
  new MenuFolder({
    key: menuFolderKey.GameSetting,
    icon: PuzzlePieceIcon,
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
  // #風控功能
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
  // 跑馬燈設定
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.MarqueeSetting,
    role: roleItemKey.MarqueeSettingRead,
    component: defineAsyncComponent(() => import('@/base/views/OperationManagement/MarqueeSetting/MarqueeSetting.vue')),
  }),
  // 玩家帳號
  new PlayerAccountMenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.PlayerAccount,
    role: roleItemKey.PlayerAccountRead,
    component: defineAsyncComponent(() => import('@/base/views/OperationManagement/PlayerAccount/PlayerAccount.vue')),
  }),
  // 玩家登入紀錄
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.PlayerLoginInfo,
    role: roleItemKey.PlayerLoginInfoRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/PlayerLoginInfo/PlayerLoginInfo.vue')
    ),
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
  // 後台操作紀錄
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.BackendActionLog,
    role: roleItemKey.BackendActionLogRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/BackendActionLog/BackendActionLog.vue')
    ),
  }),
  // 維護相關設定
  new MenuItem({
    folderKey: menuFolderKey.OperationManagement,
    key: menuItemKey.MaintainSetting,
    role: roleItemKey.MaintainSettingRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/OperationManagement/MaintainSetting/MaintainSetting.vue')
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
  // 匯率設定
  new MenuItem({
    folderKey: menuFolderKey.SystemManagement,
    key: menuItemKey.ExchangeRateSetting,
    role: roleItemKey.ExchangeRateSettingRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/SystemManagement/ExchangeRateSetting/ExchangeRateSetting.vue')
    ),
  }),

  // #遊戲設置
  // 遊戲設置
  new MenuItem({
    folderKey: menuFolderKey.GameSetting,
    key: menuItemKey.GameSetting,
    role: roleItemKey.GameSettingRead,
    component: defineAsyncComponent(() => import('@/base/views/GameSetting/GameSetting/GameSetting.vue')),
  }),
  // 遊戲管理
  new MenuItem({
    folderKey: menuFolderKey.GameSetting,
    key: menuItemKey.GameManagement,
    role: roleItemKey.GameManagementRead,
    component: defineAsyncComponent(() => import('@/base/views/GameSetting/GameManagement/GameManagement.vue')),
  }),
  // 遊戲罐頭語
  new MenuItem({
    folderKey: menuFolderKey.GameSetting,
    key: menuItemKey.GameCannedLanguage,
    role: roleItemKey.GameCannedLanguageRead,
    component: defineAsyncComponent(() => import('@/base/views/GameSetting/GameCannedLanguage/GameCannedLanguage.vue')),
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
  // 玩家帳變紀錄
  new MenuItem({
    folderKey: menuFolderKey.ReportManagement,
    key: menuItemKey.UserCreditReport,
    role: roleItemKey.UserCreditReportRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/ReportManagement/UserCreditReport/UserCreditReport.vue')
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
  // 總代理風控設計
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.GeneralAgentRTPSet,
    role: roleItemKey.GeneralAgentRTPSetRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/RiskManagement/GeneralAgentRTPSet/GeneralAgentRTPSet.vue')
    ),
  }),
  // 遊戲風控設定
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.RTPSetting,
    role: roleItemKey.RTPSettingRead,
    component: defineAsyncComponent(() => import('@/base/views/RiskManagement/RTPSetting/RTPSetting.vue')),
  }),
  // 代理風控統計
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.AgentRTPStatistics,
    role: roleItemKey.AgentRTPStatisticsRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/RiskManagement/AgentRTPStatistics/AgentRTPStatistics.vue')
    ),
  }),
  // 玩家標示
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.PlayerBadge,
    role: roleItemKey.PlayerBadgeRead,
    component: defineAsyncComponent(() => import('@/base/views/RiskManagement/PlayerBadge/PlayerBadge.vue')),
  }),
  // 自動風控設定
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.AutoRiskControlSetting,
    role: roleItemKey.AutoRiskControlSettingRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/RiskManagement/AutoRiskControlSetting/AutoRiskControlSetting.vue')
    ),
  }),
  // 自動風控紀錄
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.AutoRiskControlLog,
    role: roleItemKey.AutoRiskControlLogRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/RiskManagement/AutoRiskControlLog/AutoRiskControlLog.vue')
    ),
  }),
  // 遊戲基礎設定
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.GameBasicSetting,
    role: roleItemKey.GameBasicSettingRead,
    component: defineAsyncComponent(() => import('@/base/views/RiskManagement/GameBasicSetting/GameBasicSetting.vue')),
  }),
  // 遊戲即時資訊
  new MenuItem({
    folderKey: menuFolderKey.RiskManagement,
    key: menuItemKey.RealTimeGameRatio,
    role: roleItemKey.RealTimeGameRatioRead,
    component: defineAsyncComponent(() =>
      import('@/base/views/RiskManagement/RealTimeGameRatio/RealTimeGameRatio.vue')
    ),
  }),
  // #Jackpot功能管理
  // Jackpot參加設定
  new MenuItem({
    folderKey: menuFolderKey.JackpotManagement,
    key: menuItemKey.JackpotSetting,
    role: roleItemKey.JackpotSettingRead,
    component: defineAsyncComponent(() => import('@/base/views/JackpotManagement/JackpotSetting.vue')),
  }),
  // Jackpot獎池資訊
  new MenuItem({
    folderKey: menuFolderKey.JackpotManagement,
    key: menuItemKey.JackpotPrizeInfo,
    role: roleItemKey.JackpotPrizeInfoRead,
    component: defineAsyncComponent(() => import('@/base/views/JackpotManagement/JackpotPrizeInfo.vue')),
  }),
  // 玩家貢獻度
  new MenuItem({
    folderKey: menuFolderKey.JackpotManagement,
    key: menuItemKey.PlayerContribution,
    role: roleItemKey.PlayerContributionRead,
    component: defineAsyncComponent(() => import('@/base/views/JackpotManagement/PlayerContribution.vue')),
  }),
  // Jackpot中獎紀錄
  new MenuItem({
    folderKey: menuFolderKey.JackpotManagement,
    key: menuItemKey.JackpotPrizeRecord,
    role: roleItemKey.JackpotPrizeRecordRead,
    component: defineAsyncComponent(() => import('@/base/views/JackpotManagement/JackpotPrizeRecord.vue')),
  }),
  // Jackpot代幣紀錄
  new MenuItem({
    folderKey: menuFolderKey.JackpotManagement,
    key: menuItemKey.JackpotTokenRecord,
    role: roleItemKey.JackpotTokenRecordRead,
    component: defineAsyncComponent(() => import('@/base/views/JackpotManagement/JackpotTokenRecord.vue')),
  }),
]

export const menuItemMap = menuItems.reduce((obj, item) => {
  obj[item.key] = item
  return obj
}, {})
