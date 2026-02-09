# chat service connect information

* Date: 2022/11/14
* Author: Kinco Hsieh

## Backend side

後台服務連線說明

### 1. 取得chat service 連線資訊

**api info**
```
scheme: http/https://
api path: /api/v1/notify/getchatserviceconnInfo
method: GET
verify: jwt
```

**api params**
```
no params
```

**return data example**
```
{
  "code": 0,
  "msg": "操作成功",
  "data": {
    "channel": "local",
    "domain": "127.0.0.1:8986",
    "path": "/chatservice.api/v1/",
    "scheme": "http",
    "api_key": "defaultkey",
    "ws_conn_path": "/chatservice.ws"
  }
}
```
* channel: 聊天頻道(平台碼)
* domain: connect address
* path: api path
* scheme: http/https
* api_key: BasicAuth username
* ws_conn_path: websocket connect path

### 2. 即時發送訊息至 chat service 進行廣播

**api info**
```
scheme: http/https://
api path: /api/v1/notify/notifybroadcastmessage
method: POST
verify: jwt
```

**api params**
```
no params
```

**return data example**
```
{
  "code": 0,
  "msg": "操作成功",
  "data": ""
}
```

## ChatService side

即時通訊服務連線說明

### 1. 登入取得連線 token

**api info**
```
scheme: http/https://
api path: /chatservice.api/v1/login
method: POST
verify: BasicAuth
	BasicAuth username: /api/v1/notify/getchatserviceconnInfo 取得之 api_key
	BasicAuth password: ""
```

**api params**
```
query string: username=gg123456&agent_id=99&platform=local
```
* username: 後台帳號
* agent_id: 代理id
* platform: 聊天頻道(平台碼)

**return data example**
```
{
  "code": 0,
  "data": {
    "agent_id": 99,
    "platform": "local",
    "token": "e8c359bd-b2ec-4ec4-8a94-9e6e166b85f7",
    "user_id": "13bf2db2-5101-45c6-a48c-700b51ca3bdf",
    "username": "gg123456"
  },
  "msg": "successed"
}
```
* agent_id: 代理id
* platform:
* token:
* user_id:
* username: 後台帳號

### 2. 使用連線 token 建立並連線 websocket

**api info**
```
scheme: ws://
api path: 127.0.0.1:8986/chatservice.ws
verify: token
```

**api params**
```
query string: token=298a0326-b90d-44d9-9edf-c2b1df9cffa8
```
* token: chatservice.api/v1/login 取得之token

**return data example**
```
no
```