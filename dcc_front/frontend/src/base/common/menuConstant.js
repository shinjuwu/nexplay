import fc from '@/base/common/featureCodes'

export class MenuFolder {
  constructor(input) {
    this.key = input.key
    this.icon = input.icon
    this.childFolders = input.childFolders
    this.nameKeyPrefix = 'menuFolder'
  }
}

export class MenuItem {
  constructor(input) {
    this.folderKey = input.folderKey
    this.key = input.key
    this.role = input.role // 1D array []
    this.component = input.component
    this.nameKeyPrefix = 'menuItem'
  }

  createProps() {
    return {}
  }
}

export class GameLogParseMenuItem extends MenuItem {
  constructor(input) {
    super(input)
  }

  createProps() {
    return {
      logNumber: '',
      userName: '',
    }
  }
}

export class PlayerAccountMenuItem extends MenuItem {
  constructor(input) {
    super(input)
  }

  createProps() {
    return {
      agentId: -1,
      userName: '',
    }
  }
}

export class PlayerLogMenuItem extends MenuItem {
  constructor(input) {
    super(input)
  }

  createProps() {
    return {
      agentId: -1,
      userName: '',
    }
  }
}

class RoleGroup {
  constructor(input) {
    this.folderKey = input.folderKey
    this.items = input.items
  }
}

class RoleItem {
  constructor(input) {
    this.key = input.key
    this.permissions = input.permissions
    this.nameKey = input.nameKey || input.key
    this.nameKeyPrefix = 'roleItem'
  }
}

export const menuFolderKey = {
  LevelOneMenu: 'LevelOneMenu', // 一級菜單
  OperationManagement: 'OperationManagement', // 運營管理
  SystemManagement: 'SystemManagement', // 系統管理
  GameSetting: 'GameSetting', // 遊戲設置
  AgentAccountManagement: 'AgentAccountManagement', // 代理帳號管理
  ReportManagement: 'ReportManagement', // 報表管理
  RiskManagement: 'RiskManagement', // 風控功能
  JackpotManagement: 'JackpotManagement', // Jackpot功能管理
}

export const menuItemKey = {
  // #運營管理
  DataOverview: 'DataOverview', // 數據總覽
  BackendAnnouncement: 'BackendAnnouncement', // 後台公告
  MarqueeSetting: 'MarqueeSetting', // 跑馬燈設定
  PlayerAccount: 'PlayerAccount', // 玩家帳號
  PlayerLoginInfo: 'PlayerLoginInfo', // 玩家登入資訊
  GameLogParse: 'GameLogParse', // 遊戲日誌解析
  BackendUpdateAgentWallet: 'BackendUpdateAgentWallet', // 後台代理上下分
  BackendUpdateGameUserWallet: 'BackendUpdateGameUserWallet', // 後台玩家上下分
  BackendActionLog: 'BackendActionLog', // 後台操作紀錄
  MaintainSetting: 'MaintainSetting', // 維護相關設定

  // #系統管理
  AgentIpWhitelist: 'AgentIpWhitelist', // 後台IP白名單

  // #遊戲設置
  GameSetting: 'GameSetting', // 遊戲設置
  GameManagement: 'GameManagement', // 遊戲管理
  GameSorting: 'GameSorting', // 遊戲排序
  GameCannedLanguage: 'GameCannedLanguage', // 遊戲罐頭語

  // #代理帳號管理
  AgentAccount: 'AgentAccount', // 代理帳號
  BackendAccount: 'BackendAccount', // 後台帳號
  GroupRoleManagement: 'GroupRoleManagement', // 權限群組管理
  BackendLoginLog: 'BackendLoginLog', // 後台登入紀錄
  ExchangeRateSetting: 'ExchangeRateSetting', // 匯率設定

  // #報表管理
  WinLoseReport: 'WinLoseReport', // 輸贏報表
  EarningReport: 'EarningReport', // 業績報表
  DailySettlementReport: 'DailySettlementReport', // 日結算報表
  PlayerLog: 'PlayerLog', // 玩家遊玩紀錄
  AgentWalletLedger: 'AgentWalletLedger', // 代理分數紀錄
  WalletLedger: 'WalletLedger', // 玩家分數紀錄
  FriendRoomReport: 'FriendRoomReport', // 好友房建房紀錄
  UserCreditReport: 'UserCreditReport', // 玩家帳變紀錄

  // #風控功能
  GeneralAgentRTPSet: 'GeneralAgentRTPSet', // 總代理風控設定
  RTPSetting: 'RTPSetting', // 遊戲風控設定
  AgentRTPStatistics: 'AgentRTPStatistics', // 代理風控統計
  PlayerBadge: 'PlayerBadge', // 玩家標示
  PlayerBadgeSetting: 'PlayerBadgeSetting', // 玩家標示設定
  AutoRiskControlSetting: 'AutoRiskControlSetting', // 自動風控設定
  AutoRiskControlLog: 'AutoRiskControlLog', // 自動風控紀錄
  GameBasicSetting: 'GameBasicSetting', // 遊戲基礎設定
  RealTimeGameRatio: 'RealTimeGameRatio', // 遊戲即時資訊

  // #JP功能管理
  JackpotSetting: 'JackpotSetting', // Jackpot參加設定
  JackpotPrizeInfo: 'JackpotPrizeInfo', // Jackpot獎池資訊
  PlayerContribution: 'PlayerContribution', // 玩家貢獻度
  JackpotPrizeRecord: 'JackpotPrizeRecord', // Jackpot中獎紀錄
  JackpotTokenRecord: 'JackpotTokenRecord', // Jackpot代幣紀錄
}

export const roleItemKey = {
  // #通用
  Read: 'Read', // 檢視
  Setting: 'Setting', // 設定

  // #運營管理
  // 數據總覽
  DataOverviewRead: 'DataOverviewRead', // 檢視
  DataOverviewDeviceLocationRead: 'DataOverviewDeviceLocationRead', // 裝置位置檢視

  // 後台公告
  BackendAnnouncementRead: 'BackendAnnouncementRead', // 檢視
  BackendAnnouncementCreate: 'BackendAnnouncementCreate', // 添加公告
  BackendAnnouncementUpdate: 'BackendAnnouncementUpdate', // 編輯公告
  BackendAnnouncementDelete: 'BackendAnnouncementDelete', // 刪除公告

  // 跑馬燈設定
  MarqueeSettingRead: 'MarqueeSettingRead', // 檢視
  MarqueeSettingCreate: 'MarqueeSettingCreate', // 添加跑馬燈
  MarqueeSettingUpdate: 'MarqueeSettingUpdate', // 編輯跑馬燈設定
  MarqueeSettingDelete: 'MarqueeSettingDelete', // 刪除跑馬燈設定

  // 玩家帳號
  PlayerAccountRead: 'PlayerAccountRead', // 檢視
  PlayerAccountInfoUpdate: 'PlayerAccountInfoUpdate', // 帳號設定
  PlayerDisposeSettingUpdate: 'PlayerDisposeSettingUpdate', // 處置設定

  // 玩家登入資訊
  PlayerLoginInfoRead: 'PlayerLoginInfoRead', // 檢視

  // 遊戲日誌解析
  GameLogParseRead: 'GameLogParseRead', // 檢視

  // 後台代理上下分
  BackendUpdateAgentWalletRead: 'BackendUpdateAgentWalletRead', // 檢視
  BackendUpdateAgentWalletUpdate: 'BackendUpdateAgentWalletUpdate', // 上下分設定

  // 後台玩家上下分
  BackendUpdateGameUserWalletRead: 'BackendUpdateGameUserWalletRead', // 檢視
  BackendUpdateGameUserWalletUpdate: 'BackendUpdateGameUserWalletUpdate', // 上下分設定

  // 後台操作紀錄
  BackendActionLogRead: 'BackendActionLogRead', // 檢視

  // 維護相關設定
  MaintainSettingRead: 'MaintainSettingRead', // 檢視
  MaintainSettingUpdate: 'MaintainSettingUpdate', // 設定

  // #系統管理
  // 後台IP白名單
  AgentIpWhitelistRead: 'AgentIpWhitelistRead', // 檢視
  AgentIpWhitelistUpdate: 'AgentIpWhitelistUpdate', // 後台IP白名單設置

  // 後台帳號
  BackendAccountRead: 'BackendAccountRead', // 檢視
  BackendAccountCreate: 'BackendAccountCreate', // 添加後台帳號
  BackendAccountUpdate: 'BackendAccountUpdate', // 後台帳號設定

  // 權限群組管理
  GroupRoleManagementRead: 'GroupRoleManagementRead', // 檢視
  GroupRoleManagementCreate: 'GroupRoleManagementCreate', // 添加群組
  GroupRoleManagementUpdate: 'GroupRoleManagementUpdate', // 編輯群組
  GroupRoleManagementDelete: 'GroupRoleManagementDelete', // 刪除群組

  // 後台登入紀錄
  BackendLoginLogRead: 'BackendLoginLogRead', // 檢視

  // 匯率設定
  ExchangeRateSettingRead: 'ExchangeRateSettingRead', // 檢視
  ExchangeRateSettingUpdate: 'ExchangeRateSettingUpdate', // 設定

  // #遊戲設置
  // 遊戲設置
  GameSettingRead: 'GameSettingRead', // 檢視
  GameSettingGameStateUpdate: 'GameSettingGameStateUpdate', // 遊戲狀態設定

  // 遊戲管理
  GameManagementRead: 'GameManagementRead', // 檢視
  GameManagementGameStateUpdate: 'GameManagementGameStateUpdate', // 代理遊戲狀態設定
  GameManagementGameRoomStateUpdate: 'GameManagementGameRoomStateUpdate', // 代理遊戲房間狀態設定

  // 遊戲排序
  GameSortingRead: 'GameSortingRead', // 檢視
  GameSortingUpdate: 'GameSortingUpdate', // 排序設定

  // 遊戲罐頭語
  GameCannedLanguageRead: 'GameCannedLanguageRead', // 檢視
  GameCannedLanguageUpdate: 'GameCannedLanguageUpdate', // 罐頭語設定

  // #代理帳號管理
  // 代理帳號
  AgentAccountRead: 'AgentAccountRead', // 檢視
  AgentAccountCreate: 'AgentAccountCreate', // 添加代理
  AgentAccountSecretKeyRead: 'AgentAccountSecretKeyRead', // 查看密鑰
  AgentAccountUpdate: 'AgentAccountUpdate', // 代理商設定

  // #報表管理
  // 輸贏報表
  WinLoseReportRead: 'WinLoseReportRead', // 檢視

  // 日結算報表
  DailySettlementReportRead: 'DailySettlementReportRead', // 檢視

  // 業績報表
  EarningReportRead: 'EarningReportRead', // 檢視

  // 玩家帳變紀錄
  UserCreditReportRead: 'UserCreditReportRead', // 檢視

  // 玩家遊玩紀錄
  PlayerLogRead: 'PlayerLogRead', // 檢視

  // 代理分數紀錄
  AgentWalletLedgerRead: 'AgentWalletLedgerRead', // 檢視

  // 玩家分數紀錄
  WalletLedgerRead: 'WalletLedgerRead', // 檢視

  // 好友房建房紀錄
  FriendRoomReportRead: 'FriendRoomReportRead', // 檢視

  // #風控功能
  // 總代理風控設計
  GeneralAgentRTPSetRead: 'GeneralAgentRTPSetRead', // 檢視
  GeneralAgentRTPSetUpdate: 'GeneralAgentRTPUpdate', // 設定

  // 遊戲風控設定
  RTPSettingRead: 'RTPSettingRead', // 檢視
  RTPSettingUpdate: 'RTPSettingUpdate', // 設定
  RTPBatchSettingUpdate: 'RTPBatchSettingUpdate', // 批次設定

  // 代理風控統計
  AgentRTPStatisticsRead: 'AgentRTPStatisticsRead', // 檢視

  // 玩家標示
  PlayerBadgeRead: 'PlayerBadgeRead', // 檢視
  PlayerBadgeUpdate: 'PlayerBadgeUpdate', // 標示

  // 玩家標示設定
  PlayerBadgeSettingRead: 'PlayerBadgeSettingRead', // 檢視
  PlayerBadgeSettingUpdate: 'PlayerBadgeSettingUpdate', // 設定

  // 自動風控設定
  AutoRiskControlSettingRead: 'AutoRiskControlSettingRead', // 檢視
  AutoRiskControlSettingUpdate: 'AutoRiskControlSettingUpdate', // 修改

  // 自動風控紀錄
  AutoRiskControlLogRead: 'AutoRiskControlLogRead', // 檢視

  // 遊戲基礎設定
  GameBasicSettingRead: 'GameBasicSettingRead', // 檢視
  GameBasicSettingUpdate: 'GameBasicSettingUpdate', // 修改

  // 遊戲即時資訊
  RealTimeGameRatioRead: 'RealTimeGameRatioRead', // 檢視

  // 重置密碼
  ResetPassword: 'ResetPassword', // 重置密碼

  // #Jackpot功能管理
  // Jackpot設定
  JackpotSettingRead: 'JackpotSettingRead', // 檢視
  JackpotSettingUpdate: 'JackpotSettingUpdate', // 修改
  // Jackpot獎池資訊
  JackpotPrizeInfoRead: 'JackpotPrizeInfoRead', // 檢視
  // 玩家貢獻度
  PlayerContributionRead: 'PlayerContributionRead', // 檢視
  // Jackpot中獎紀錄
  JackpotPrizeRecordRead: 'JackpotPrizeRecordRead', // 檢視
  // Jackpot代幣紀錄
  JackpotTokenRecordRead: 'JackpotTokenRecordRead', // 檢視
  JackpotTokenUpdate: 'JackpotTokenUpdate', // 添加JP代幣
}

export const rolePermissions = {
  // 後台基礎
  BackendBasic: [
    fc.UserPing,
    fc.GlobalGetAgentList,
    fc.GlobalGetAllGameList,
    fc.GlobalGetGameList,
    fc.GlobalGetRoomTypeList,
    fc.GlobalGetAgentPermissionList,
    fc.GlobalGetAgentCustomTagSettingList,
    fc.NotifyGetChatServiceConnInfo,
    fc.UserSetPersonalInfo,
    fc.UserSetPersonalPassword,
  ],

  // #運營管理
  // 數據總覽-檢視
  DataOverviewRead: [
    fc.ManageGetStatData,
    fc.ManageGetRiskUserList,
    fc.ManageGetGameLeaderboards,
    fc.ManageGetIntervalDaysBettorData,
    fc.ManageGetIntervalTotalScoreData,
    fc.ManageGetIntervalTotalBettorInfoData,
    fc.ManageGetIntervalRealTimeUserData,
  ],
  // 數據總覽-裝置位置檢視
  DataOverviewDeviceLocationRead: [fc.UserLoginDeviceUsageRatio, fc.UserLoginSourceLocRanking],

  // 後台公告-檢視
  BackendAnnouncementRead: [fc.ManageGetAnnouncementList, fc.ManageGetAnnouncement],
  // 後台公告-添加公告
  BackendAnnouncementCreate: [fc.ManageCreateAnnouncement],
  // 後台公告-編輯公告
  BackendAnnouncementUpdate: [fc.ManageUpdateAnnouncement],
  // 後台公告 刪除公告
  BackendAnnouncementDelete: [fc.ManageDeleteAnnouncement],

  // 跑馬燈設定-檢視
  MarqueeSettingRead: [fc.ManageGetMarquee, fc.ManageGetMarqueeList],
  // 跑馬燈設定-添加跑馬燈
  MarqueeSettingCreate: [fc.ManageCreateMarquee],
  // 跑馬燈設定-編輯跑馬燈設定
  MarqueeSettingUpdate: [fc.ManageUpdateMarquee],
  // 跑馬燈設定-刪除跑馬燈設定
  MarqueeSettingDelete: [fc.ManageDeleteMarquee],

  // 玩家帳號-檢視
  PlayerAccountRead: [
    fc.UserGetGameUsers,
    fc.UserGetGameUserInfo,
    fc.UserGetGameUserBalance,
    fc.UserGetGameUserPlayCountData,
  ],
  // 玩家帳號-帳號設定
  PlayerAccountInfoUpdate: [fc.UserUpdateGameUserInfo],

  // 玩家登入資訊-檢視
  PlayerLoginInfoRead: [
    fc.UserLoginData,
    fc.GlobalGetUserLoginLogBrowserList,
    fc.GlobalGetUserLoginLogCountryShortList,
  ],

  // 遊戲日誌解析-檢視
  GameLogParseRead: [fc.RecordGetPlayLogCommon],

  // 後台代理上下分-檢視
  BackendUpdateAgentWalletRead: [fc.AgentGetAgentWalletList],
  BackendUpdateAgentWalletUpdate: [fc.AgentSetAgentWallet],

  // 後台玩家上下分-檢視
  BackendUpdateGameUserWalletRead: [fc.UserGetGameUserWalletList],
  // 後台玩家上下分-修改
  BackendUpdateGameUserWalletUpdate: [fc.UserSetGameUserWallet],

  // 後台操作紀錄-檢視
  BackendActionLogRead: [fc.RecordGetBackendActionLog, fc.GlobalGetAgentAdminUserPermissionList],

  // 維護相關設定-檢視
  MaintainSettingRead: [fc.GameGetGameServerState, fc.ManageGetMaintainPageSetting],
  // 維護相關設定-修改
  MaintainSettingUpdate: [fc.GameSetGameServerState, fc.ManageSetMaintainPageSetting],

  // #系統管理
  // 代理IP白名單-檢視
  AgentIpWhitelistRead: [fc.AgentGetAgentIpWhitelistList, fc.AgentGetAgentIpWhitelist, fc.AgentGetAgentApiIpWhitelist],
  // 代理IP白名單-修改
  AgentIpWhitelistUpdate: [fc.AgentSetAgentIpWhitelist, fc.AgentSetAgentApiIpWhitelist],

  // 後台帳號-檢視
  BackendAccountRead: [fc.UserGetAdminUserInfo, fc.UserGetAdminUsers],
  // 後台帳號-添加後台帳號
  BackendAccountCreate: [fc.UserCreateAdminUser],
  // 後台帳號-後台帳號設定
  BackendAccountUpdate: [fc.UserUpdateAdminUserInfo],

  // 權限群組管理-檢視
  GroupRoleManagementRead: [
    fc.AgentGetAgentPermissionTemplateInfo,
    fc.AgentGetAgentPermissionList,
    fc.AgentGetAgentPermission,
  ],
  // 權限群組管理-添加群組
  GroupRoleManagementCreate: [fc.AgentCreateAgentPermission],
  // 權限群組管理-編輯群組
  GroupRoleManagementUpdate: [fc.AgentSetAgentPermission],
  // 權限群組管理-刪除群組
  GroupRoleManagementDelete: [fc.AgentDeleteAgentPermission],

  // 後台登入紀錄-檢視
  BackendLoginLogRead: [fc.RecordGetBackendLoginLogList],

  // 匯率設定-檢視
  ExchangeRateSettingRead: [fc.SystemGetExchangeDataList],
  // 匯率設定-修改
  ExchangeRateSettingUpdate: [fc.SystemSetExchangeDataList],

  // #遊戲設置
  // 遊戲設置-檢視
  GameSettingRead: [fc.GameGetGameList],
  // 遊戲設置-遊戲狀態設定
  GameSettingGameStateUpdate: [fc.GameSetGameState, fc.GameNotifyGameServer],

  // 遊戲管理-檢視
  GameManagementRead: [fc.AgentGetAgentGameList, fc.AgentGetAgentGameRoomList],
  // 遊戲管理-代理遊戲狀態設定
  GameManagementGameStateUpdate: [fc.AgentSetAgentGameState],
  // 遊戲管理-代理遊戲房間狀態設定
  GameManagementGameRoomStateUpdate: [fc.AgentSetAgentGameRoomState],

  // 遊戲排序-檢視
  GameSortingRead: [fc.GameGetGameIconList],
  // 遊戲排序-排序設定
  GameSortingUpdate: [fc.GameSetGameIconList],

  // 遊戲罐頭語-檢視
  GameCannedLanguageRead: [fc.GameGetCannedList],
  // 遊戲罐頭語-罐頭語設定
  GameCannedLanguageUpdate: [fc.GameSetCannedList],

  // #代理帳號管理
  // 代理帳號-檢視
  AgentAccountRead: [fc.AgentGetAgentList, fc.AgentGetAgentCoinSupplyInfo],
  // 代理帳號-添加代理
  AgentAccountCreate: [fc.AgentCreateAgent],
  // 代理帳號-查看密鑰
  AgentAccountSecretKeyRead: [fc.AgentGetAgentSecretKey],
  // 代理帳號-代理商設定
  AgentAccountUpdate: [fc.AgentSetAgentCoinSupplyInfo],

  // #報表管理
  // 輸贏報表-檢視
  WinLoseReportRead: [fc.RecordGetUserPlayLogList],
  //  業績報表-檢視
  EarningReportRead: [fc.CalGetPerformanceReportList, fc.CalGetPerformanceReport],
  // 日結算報表-檢視
  DailySettlementReportRead: [fc.RecordGetAgentGameRatioStatList],
  // 玩家帳變紀錄-檢視
  UserCreditReportRead: [fc.RecordGetUserCreditLogList],
  // 玩家遊玩紀錄-檢視
  PlayerLogRead: [fc.RecordGetGameUsersStatHourList],
  // 代理分數紀錄-檢視
  AgentWalletLedgerRead: [fc.RecordGetAgentWalletLedgerList],
  // 玩家分數紀錄-檢視
  WalletLedgerRead: [fc.RecordGetWalletLedgerList, fc.RecordConfirmWalletLedger],
  // 好友房建房紀錄-檢視
  FriendRoomReportRead: [fc.RecordGetFriendRoomLogList],

  // #風控功能
  // 總代理風控設定-檢視
  GeneralAgentRTPSetRead: [fc.RiskControlGetAgentIncomeRatioList],
  // 總代理風控設定-修改
  GeneralAgentRTPSetUpdate: [fc.RiskControlSetAgentIncomeRatio, fc.RiskControlGetAgentIncomeRatio],
  // 遊戲風控設定-檢視
  RTPSettingRead: [fc.RiskControlGetIncomeRatio, fc.RiskControlGetIncomeRatioList],
  RTPSettingUpdate: [fc.RiskControlSetIncomeRatio],
  RTPBatchSettingUpdate: [fc.RiskControlSetIncomeRatios],
  // 代理風控統計-檢視
  AgentRTPStatisticsRead: [fc.RiskControlGetAgentIncomeRatioAndGameData],
  // 玩家標示-檢視
  PlayerBadgeRead: [fc.RiskControlGetGameUsersCustomTagList],
  // 玩家標示-編輯
  PlayerBadgeUpdate: [fc.RiskControlSetGameUsersCustomTag, fc.RiskControlGetGameUsersCustomTag],

  // 玩家標示設定-檢視
  PlayerBadgeSettingRead: [fc.RiskControlGetAgentCustomTagList],
  // 玩家標示設定-修改
  PlayerBadgeSettingUpdate: [fc.RiskControlSetAgentCustomTagList],

  // 玩家處置設定-修改
  PlayerDisposeSettingUpdate: [fc.RiskControlGetGameUserRiskControlTag, fc.RiskControlSetGameUserRiskControlTag],

  // 自動風控設定-檢視
  AutoRiskControlSettingRead: [fc.RiskControlGetAutoRiskControlSetting],
  // 自動風控設定-修改
  AutoRiskControlSettingUpdate: [fc.RiskControlSetAutoRiskControlSetting],
  // 自動風控紀錄-檢視
  AutoRiskControlLogRead: [fc.RiskControlGetAutoRiskControlLogList],

  //遊戲基礎設定-檢視
  GameBasicSettingRead: [fc.RiskControlGetGameSetting],
  //遊戲基礎設定-修改
  GameBasicSettingUpdate: [fc.RiskControlSetGameSetting],

  // 遊戲即時資訊-檢視
  RealTimeGameRatioRead: [fc.RiskControlGetRealTimeGameRatio], // 檢視

  // 重置密碼
  ResetPassword: [fc.UserResetPassword],

  // #Jackpot功能管理
  // Jackpot設定-檢視
  JackpotSettingRead: [fc.JackpotGetAgentJackpotList, fc.JackpotGetJackpotSetting, fc.JackpotGetAgentJackpot],
  // Jackpot設定-修改
  JackpotSettingUpdate: [
    fc.JackpotSetAgentJackpot,
    fc.JackpotSetJackpotSetting,
    fc.JackpotNotifyGameServerAgentJackpotInfo,
  ],

  // Jackpot獎池資訊-檢視
  JackpotPrizeInfoRead: [fc.JackpotGetJackpotPoolData],
  // 玩家貢獻度-檢視
  PlayerContributionRead: [fc.JackpotGetJackpotLeaderBoard],
  // Jackpot中獎紀錄-檢視
  JackpotPrizeRecordRead: [fc.JackpotGetJackpotList],
  // Jackpot代幣紀錄-檢視
  JackpotTokenRecordRead: [fc.JackpotGetJackpotTokenList],
  // Jackpot代幣紀錄-修改
  JackpotTokenUpdate: [fc.JackpotCreateJackpotToken],
}

export const roleMenuFolders = [
  new MenuFolder({
    key: menuFolderKey.LevelOneMenu,
    childFolders: [
      // #運營管理
      new MenuFolder({
        key: menuFolderKey.OperationManagement,
        childFolders: [
          // 數據總攬
          new MenuItem({
            key: menuItemKey.DataOverview,
          }),
          // 後台公告
          new MenuItem({
            key: menuItemKey.BackendAnnouncement,
          }),
          // 跑馬燈設定
          new MenuItem({
            key: menuItemKey.MarqueeSetting,
          }),
          // 玩家帳號
          new MenuItem({
            key: menuItemKey.PlayerAccount,
          }),
          // 玩家登入資訊
          new MenuItem({
            key: menuItemKey.PlayerLoginInfo,
          }),
          // 遊戲日誌解析
          new MenuItem({
            key: menuItemKey.GameLogParse,
          }),
          // 後台代理上下分
          new MenuItem({
            key: menuItemKey.BackendUpdateAgentWallet,
          }),
          // 後台玩家上下分
          new MenuItem({
            key: menuItemKey.BackendUpdateGameUserWallet,
          }),
          // 後台操作紀錄
          new MenuItem({
            key: menuItemKey.BackendActionLog,
          }),
          // 維護相關設定
          new MenuItem({
            key: menuItemKey.MaintainSetting,
          }),
        ],
      }),
      // #系統管理
      new MenuFolder({
        key: menuFolderKey.SystemManagement,
        childFolders: [
          new MenuItem({
            key: menuItemKey.AgentIpWhitelist,
          }),
          // 後台帳號
          new MenuItem({
            key: menuItemKey.BackendAccount,
          }),
          // 權限群組管理
          new MenuItem({
            key: menuItemKey.GroupRoleManagement,
          }),
          // 後台登入紀錄
          new MenuItem({
            key: menuItemKey.BackendLoginLog,
          }),
          // 匯率設定
          new MenuItem({
            key: menuItemKey.ExchangeRateSetting,
          }),
        ],
      }),
      // #遊戲設置
      new MenuFolder({
        key: menuFolderKey.GameSetting,
        childFolders: [
          // 遊戲設置
          new MenuItem({
            key: menuItemKey.GameSetting,
          }),
          // 遊戲管理
          new MenuItem({
            key: menuItemKey.GameManagement,
          }),
          // 遊戲排序
          new MenuItem({
            key: menuItemKey.GameSorting,
          }),
          // 遊戲罐頭語
          new MenuItem({
            key: menuItemKey.GameCannedLanguage,
          }),
        ],
      }),
      // #代理帳號管理
      new MenuFolder({
        key: menuFolderKey.AgentAccountManagement,
        childFolders: [
          // 代理帳號
          new MenuItem({
            key: menuItemKey.AgentAccount,
          }),
        ],
      }),
      // #報表管理
      new MenuFolder({
        key: menuFolderKey.ReportManagement,
        childFolders: [
          // 輸贏報表
          new MenuItem({
            key: menuItemKey.WinLoseReport,
          }),
          // 業績報表
          new MenuItem({
            key: menuItemKey.EarningReport,
          }),
          // 日結算報表
          new MenuItem({
            key: menuItemKey.DailySettlementReport,
          }),
          // 玩家帳變紀錄
          new MenuItem({
            key: menuItemKey.UserCreditReport,
          }),
          // 玩家遊玩紀錄
          new MenuItem({
            key: menuItemKey.PlayerLog,
          }),
          // 代理分數紀錄
          new MenuItem({
            key: menuItemKey.AgentWalletLedger,
          }),
          // 玩家分數紀錄
          new MenuItem({
            key: menuItemKey.WalletLedger,
          }),
          // 好友房建房紀錄
          new MenuItem({
            key: menuItemKey.FriendRoomReport,
          }),
        ],
      }),
      // #風控功能
      new MenuFolder({
        key: menuFolderKey.RiskManagement,
        childFolders: [
          // 總代理風控設定
          new MenuItem({
            key: menuItemKey.GeneralAgentRTPSet,
          }),
          // 遊戲風控設定
          new MenuItem({
            key: menuItemKey.RTPSetting,
          }),
          // 代理風控統計
          new MenuItem({
            key: menuItemKey.AgentRTPStatistics,
          }),
          // 玩家標示
          new MenuItem({
            key: menuItemKey.PlayerBadge,
          }),
          // 玩家標示設定
          new MenuItem({
            key: menuItemKey.PlayerBadgeSetting,
          }),
          // 自動風控設定
          new MenuItem({
            key: menuItemKey.AutoRiskControlSetting,
          }),
          // 自動風控紀錄
          new MenuItem({
            key: menuItemKey.AutoRiskControlLog,
          }),
          // 遊戲基礎設定
          new MenuItem({
            key: menuItemKey.GameBasicSetting,
          }),
          // 遊戲基礎設定
          new MenuItem({
            key: menuItemKey.RealTimeGameRatio,
          }),
        ],
      }),
      // #Jackpot功能管理
      new MenuFolder({
        key: menuFolderKey.JackpotManagement,
        childFolders: [
          // Jackpot參加設定
          new MenuItem({
            key: menuItemKey.JackpotSetting,
          }),
          // Jackpot獎池資訊
          new MenuItem({
            key: menuItemKey.JackpotPrizeInfo,
          }),
          // 玩家貢獻度
          new MenuItem({
            key: menuItemKey.PlayerContribution,
          }),
          // JP中獎紀錄
          new MenuItem({
            key: menuItemKey.JackpotPrizeRecord,
          }),
          // JP代幣紀錄
          new MenuItem({
            key: menuItemKey.JackpotTokenRecord,
          }),
        ],
      }),
    ],
  }),
]

export const roleGroups = [
  // #運營管理
  // 數據總攬
  new RoleGroup({
    folderKey: menuItemKey.DataOverview,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.DataOverviewRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.DataOverviewRead,
      }),
      // 裝置位置檢視
      new RoleItem({
        key: roleItemKey.DataOverviewDeviceLocationRead,
        permissions: rolePermissions.DataOverviewDeviceLocationRead,
      }),
    ],
  }),
  // 後台公告
  new RoleGroup({
    folderKey: menuItemKey.BackendAnnouncement,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.BackendAnnouncementRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.BackendAnnouncementRead,
      }),
      // 添加公告
      new RoleItem({
        key: roleItemKey.BackendAnnouncementCreate,
        permissions: rolePermissions.BackendAnnouncementCreate,
      }),
      // 編輯公告
      new RoleItem({
        key: roleItemKey.BackendAnnouncementUpdate,
        permissions: rolePermissions.BackendAnnouncementUpdate,
      }),
      // 刪除公告
      new RoleItem({
        key: roleItemKey.BackendAnnouncementDelete,
        permissions: rolePermissions.BackendAnnouncementDelete,
      }),
    ],
  }),
  // 跑馬燈設定
  new RoleGroup({
    folderKey: menuItemKey.MarqueeSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.MarqueeSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.MarqueeSettingRead,
      }),
      // 添加跑馬燈
      new RoleItem({
        key: roleItemKey.MarqueeSettingCreate,
        permissions: rolePermissions.MarqueeSettingCreate,
      }),
      // 編輯跑馬燈設定
      new RoleItem({
        key: roleItemKey.MarqueeSettingUpdate,
        permissions: rolePermissions.MarqueeSettingUpdate,
      }),
      // 刪除跑馬燈設定
      new RoleItem({
        key: roleItemKey.MarqueeSettingDelete,
        permissions: rolePermissions.MarqueeSettingDelete,
      }),
    ],
  }),
  // 玩家帳號
  new RoleGroup({
    folderKey: menuItemKey.PlayerAccount,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.PlayerAccountRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.PlayerAccountRead,
      }),
      // 標示
      new RoleItem({
        key: roleItemKey.PlayerBadgeUpdate,
        permissions: rolePermissions.PlayerBadgeUpdate,
      }),
      // 帳號設定
      new RoleItem({
        key: roleItemKey.PlayerAccountInfoUpdate,
        permissions: rolePermissions.PlayerAccountInfoUpdate,
      }),
      // 處置設定
      new RoleItem({
        key: roleItemKey.PlayerDisposeSettingUpdate,
        permissions: rolePermissions.PlayerDisposeSettingUpdate,
      }),
    ],
  }),
  // 玩家登入資訊
  new RoleGroup({
    folderKey: menuItemKey.PlayerLoginInfo,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.PlayerLoginInfoRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.PlayerLoginInfoRead,
      }),
    ],
  }),
  // 遊戲日誌解析
  new RoleGroup({
    folderKey: menuItemKey.GameLogParse,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GameLogParseRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GameLogParseRead,
      }),
    ],
  }),
  // 後台代理上下分
  new RoleGroup({
    folderKey: menuItemKey.BackendUpdateAgentWallet,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.BackendUpdateAgentWalletRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.BackendUpdateAgentWalletRead,
      }),
      // 上下分設定
      new RoleItem({
        key: roleItemKey.BackendUpdateAgentWalletUpdate,
        permissions: rolePermissions.BackendUpdateAgentWalletUpdate,
      }),
    ],
  }),
  // 後台玩家上下分
  new RoleGroup({
    folderKey: menuItemKey.BackendUpdateGameUserWallet,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.BackendUpdateGameUserWalletRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.BackendUpdateGameUserWalletRead,
      }),
      // 上下分設定
      new RoleItem({
        key: roleItemKey.BackendUpdateGameUserWalletUpdate,
        permissions: rolePermissions.BackendUpdateGameUserWalletUpdate,
      }),
    ],
  }),
  // 後台操作紀錄
  new RoleGroup({
    folderKey: menuItemKey.BackendActionLog,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.BackendActionLogRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.BackendActionLogRead,
      }),
    ],
  }),
  // 維護相關設定
  new RoleGroup({
    folderKey: menuItemKey.MaintainSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.MaintainSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.MaintainSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.MaintainSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.MaintainSettingUpdate,
      }),
    ],
  }),

  // #系統設置
  // 後台IP白名單
  new RoleGroup({
    folderKey: menuItemKey.AgentIpWhitelist,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.AgentIpWhitelistRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.AgentIpWhitelistRead,
      }),
      // 後台IP白名單設置
      new RoleItem({
        key: roleItemKey.AgentIpWhitelistUpdate,
        permissions: rolePermissions.AgentIpWhitelistUpdate,
      }),
    ],
  }),
  // 後台帳號
  new RoleGroup({
    folderKey: menuItemKey.BackendAccount,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.BackendAccountRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.BackendAccountRead,
      }),
      // 添加後台帳號
      new RoleItem({
        key: roleItemKey.BackendAccountCreate,
        permissions: rolePermissions.BackendAccountCreate,
      }),
      // 後台帳號設定
      new RoleItem({
        key: roleItemKey.BackendAccountUpdate,
        permissions: rolePermissions.BackendAccountUpdate,
      }),
      // 重置密碼
      new RoleItem({
        key: roleItemKey.ResetPassword,
        permissions: rolePermissions.ResetPassword,
      }),
    ],
  }),
  // 權限群組管理
  new RoleGroup({
    folderKey: menuItemKey.GroupRoleManagement,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GroupRoleManagementRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GroupRoleManagementRead,
      }),
      // 添加群組
      new RoleItem({
        key: roleItemKey.GroupRoleManagementCreate,
        permissions: rolePermissions.GroupRoleManagementCreate,
      }),
      // 編輯群組
      new RoleItem({
        key: roleItemKey.GroupRoleManagementUpdate,
        permissions: rolePermissions.GroupRoleManagementUpdate,
      }),
      // 刪除群組
      new RoleItem({
        key: roleItemKey.GroupRoleManagementDelete,
        permissions: rolePermissions.GroupRoleManagementDelete,
      }),
    ],
  }),
  // 後台登入紀錄
  new RoleGroup({
    folderKey: menuItemKey.BackendLoginLog,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.BackendLoginLogRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.BackendLoginLogRead,
      }),
    ],
  }),
  // 匯率設定
  new RoleGroup({
    folderKey: menuItemKey.ExchangeRateSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.ExchangeRateSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.ExchangeRateSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.ExchangeRateSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.ExchangeRateSettingUpdate,
      }),
    ],
  }),

  // #遊戲設置
  // 遊戲設置
  new RoleGroup({
    folderKey: menuItemKey.GameSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GameSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GameSettingRead,
      }),
      // 遊戲狀態設定
      new RoleItem({
        key: roleItemKey.GameSettingGameStateUpdate,
        permissions: rolePermissions.GameSettingGameStateUpdate,
      }),
    ],
  }),
  // 遊戲管理
  new RoleGroup({
    folderKey: menuItemKey.GameManagement,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GameManagementRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GameManagementRead,
      }),
      // 代理遊戲狀態設定
      new RoleItem({
        key: roleItemKey.GameManagementGameStateUpdate,
        permissions: rolePermissions.GameManagementGameStateUpdate,
      }),
      // 代理遊戲房間狀態設定
      new RoleItem({
        key: roleItemKey.GameManagementGameRoomStateUpdate,
        permissions: rolePermissions.GameManagementGameRoomStateUpdate,
      }),
    ],
  }),
  // 遊戲排序
  new RoleGroup({
    folderKey: menuItemKey.GameSorting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GameSortingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GameSortingRead,
      }),
      // 排序設定
      new RoleItem({
        key: roleItemKey.GameSortingUpdate,
        permissions: rolePermissions.GameSortingUpdate,
      }),
    ],
  }),
  // 遊戲罐頭語
  new RoleGroup({
    folderKey: menuItemKey.GameCannedLanguage,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GameCannedLanguageRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GameCannedLanguageRead,
      }),
      // 罐頭語設定
      new RoleItem({
        key: roleItemKey.GameCannedLanguageUpdate,
        permissions: rolePermissions.GameCannedLanguageUpdate,
      }),
    ],
  }),

  // #代理帳號管理
  // 代理帳號
  new RoleGroup({
    folderKey: menuItemKey.AgentAccount,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.AgentAccountRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.AgentAccountRead,
      }),
      // 添加代理
      new RoleItem({
        key: roleItemKey.AgentAccountCreate,
        permissions: rolePermissions.AgentAccountCreate,
      }),
      // 查看密鑰
      new RoleItem({
        key: roleItemKey.AgentAccountSecretKeyRead,
        permissions: rolePermissions.AgentAccountSecretKeyRead,
      }),
      // 代理商設定
      new RoleItem({
        key: roleItemKey.AgentAccountUpdate,
        permissions: rolePermissions.AgentAccountUpdate,
      }),
      // 重置密碼
      new RoleItem({
        key: roleItemKey.ResetPassword,
        permissions: rolePermissions.ResetPassword,
      }),
    ],
  }),

  // #報表管理
  // 輸贏報表
  new RoleGroup({
    folderKey: menuItemKey.WinLoseReport,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.WinLoseReportRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.WinLoseReportRead,
      }),
    ],
  }),
  // 業績報表
  new RoleGroup({
    folderKey: menuItemKey.EarningReport,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.EarningReportRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.EarningReportRead,
      }),
    ],
  }),
  // 日結算報表
  new RoleGroup({
    folderKey: menuItemKey.DailySettlementReport,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.DailySettlementReportRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.DailySettlementReportRead,
      }),
    ],
  }),
  // 玩家帳變紀錄
  new RoleGroup({
    folderKey: menuItemKey.UserCreditReport,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.UserCreditReportRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.UserCreditReportRead,
      }),
    ],
  }),
  // 玩家遊玩紀錄
  new RoleGroup({
    folderKey: menuItemKey.PlayerLog,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.PlayerLogRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.PlayerLogRead,
      }),
    ],
  }),
  // 代理分數紀錄
  new RoleGroup({
    folderKey: menuItemKey.AgentWalletLedger,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.AgentWalletLedgerRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.AgentWalletLedgerRead,
      }),
    ],
  }),
  // 玩家分數紀錄
  new RoleGroup({
    folderKey: menuItemKey.WalletLedger,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.WalletLedgerRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.WalletLedgerRead,
      }),
    ],
  }),
  // 好友房建房紀錄
  new RoleGroup({
    folderKey: menuItemKey.FriendRoomReport,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.FriendRoomReportRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.FriendRoomReportRead,
      }),
    ],
  }),
  // #風控功能
  // 總代理風控設定
  new RoleGroup({
    folderKey: menuItemKey.GeneralAgentRTPSet,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GeneralAgentRTPSetRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GeneralAgentRTPSetRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.GeneralAgentRTPSetUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.GeneralAgentRTPSetUpdate,
      }),
    ],
  }),
  // 遊戲風控設定
  new RoleGroup({
    folderKey: menuItemKey.RTPSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.RTPSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.RTPSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.RTPSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.RTPSettingUpdate,
      }),
      // 批次設定
      new RoleItem({
        key: roleItemKey.RTPBatchSettingUpdate,
        permissions: rolePermissions.RTPBatchSettingUpdate,
      }),
    ],
  }),
  // 代理風控統計
  new RoleGroup({
    folderKey: menuItemKey.AgentRTPStatistics,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.AgentRTPStatisticsRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.AgentRTPStatisticsRead,
      }),
    ],
  }),
  // 玩家標示
  new RoleGroup({
    folderKey: menuItemKey.PlayerBadge,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.PlayerBadgeRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.PlayerBadgeRead,
      }),
      // 標示
      new RoleItem({
        key: roleItemKey.PlayerBadgeUpdate,
        permissions: rolePermissions.PlayerBadgeUpdate,
      }),
      // 帳號設定
      new RoleItem({
        key: roleItemKey.PlayerAccountInfoUpdate,
        permissions: rolePermissions.PlayerAccountInfoUpdate,
      }),
    ],
  }),
  // 玩家標示設定
  new RoleGroup({
    folderKey: menuItemKey.PlayerBadgeSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.PlayerBadgeSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.PlayerBadgeSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.PlayerBadgeSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.PlayerBadgeSettingUpdate,
      }),
    ],
  }),

  // 自動風控設定
  new RoleGroup({
    folderKey: menuItemKey.AutoRiskControlSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.AutoRiskControlSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.AutoRiskControlSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.AutoRiskControlSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.AutoRiskControlSettingUpdate,
      }),
    ],
  }),
  // 自動風控紀錄
  new RoleGroup({
    folderKey: menuItemKey.AutoRiskControlLog,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.AutoRiskControlLogRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.AutoRiskControlLogRead,
      }),
    ],
  }),
  // 遊戲基礎設定
  new RoleGroup({
    folderKey: menuItemKey.GameBasicSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.GameBasicSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.GameBasicSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.GameBasicSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.GameBasicSettingUpdate,
      }),
    ],
  }),
  // 遊戲即時資訊
  new RoleGroup({
    folderKey: menuItemKey.RealTimeGameRatio,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.RealTimeGameRatioRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.RealTimeGameRatioRead,
      }),
    ],
  }),
  // #Jackpot功能管理
  // Jackpot參加設定
  new RoleGroup({
    folderKey: menuItemKey.JackpotSetting,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.JackpotSettingRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.JackpotSettingRead,
      }),
      // 設定
      new RoleItem({
        key: roleItemKey.JackpotSettingUpdate,
        nameKey: roleItemKey.Setting,
        permissions: rolePermissions.JackpotSettingUpdate,
      }),
    ],
  }),
  // Jackpot獎池資訊
  new RoleGroup({
    folderKey: menuItemKey.JackpotPrizeInfo,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.JackpotPrizeInfoRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.JackpotPrizeInfoRead,
      }),
    ],
  }),
  // 玩家貢獻度
  new RoleGroup({
    folderKey: menuItemKey.PlayerContribution,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.PlayerContributionRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.PlayerContributionRead,
      }),
    ],
  }),
  // Jackpot中獎紀錄
  new RoleGroup({
    folderKey: menuItemKey.JackpotPrizeRecord,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.JackpotPrizeRecordRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.JackpotPrizeRecordRead,
      }),
    ],
  }),
  // Jackpot代幣紀錄
  new RoleGroup({
    folderKey: menuItemKey.JackpotTokenRecord,
    items: [
      // 檢視
      new RoleItem({
        key: roleItemKey.JackpotTokenRecordRead,
        nameKey: roleItemKey.Read,
        permissions: rolePermissions.JackpotTokenRecordRead,
      }),
      // 添加JP代幣
      new RoleItem({
        key: roleItemKey.JackpotTokenUpdate,
        permissions: rolePermissions.JackpotTokenUpdate,
      }),
    ],
  }),
]

export const roleItems = roleGroups.reduce((arr, group) => arr.concat(group.items), [])
