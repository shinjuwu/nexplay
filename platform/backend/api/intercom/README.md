## 提供內部(game server)呼叫API

* GetLoginToken 
    
    取得登入token

* LoginGame

    遊戲用戶登入    

* LogoutGame

    遊戲用戶登出

* CreateGameRecord

    遊戲新增遊戲紀錄

* GetMarqueeSettingList

    遊戲伺服器取得跑馬燈設定列表
```
api path: 
method: GET
/api/v1/intercom/getmarqueesetting

no params
```

* GetGameServerState
    遊戲伺服器取得設定遊戲全局開關設定列表
```
api path: 
method: GET
/api/v1/intercom/getgameserverstate

no params
```