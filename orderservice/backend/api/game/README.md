# 提供外部呼叫的API

* TODO:之後需要遷移為獨立服務
  
## 單一複合式接口

一個接口，內包含大於一個的子命令，使用 mvc 架構開發
```
const (
	ChannelHandle_Logingame         = iota // 登入遊戲
	ChannelHandle_QueryMinusScore          // 查詢可下分
	ChannelHandle_AddScore                 // 上分
	ChannelHandle_MinusScore               // 下分
	ChannelHandle_QueryOrder               // 查詢上下分訂單
	ChannelHandle_QueryUser                // 查詢玩家在線狀態
	GetRecordHandle_QueryUser              // 查詢遊戲注單
	ChannelHandle_QueryUserScore           // 查詢玩家總分
	ChannelHandle_KickUser                 // 踢玩家下線
	ChannelHandle_QueryAgentBalance        // 查詢代理余額
)
```

### LoginGame
登录游戏
```
参数加密字符串 param=
(s=0&account=kenny1323&money=100&orderid=1000120170306143076949222244)

參數說明：
s：操作子类型：0
account：会员帐号(64 位字符)
money：金额(上分的金额,如果不
携带分数传 0)
orderid：流水号（格式：代理编号+yyyyMMddHHmmssSSS+ account,长度不能超过 100 字符串）

api response：
{"s":100,"m":"/channelHandle","d":{"url":"http://127.0.0.1/game/abcde?token=1234567890","code":0}}
```

### CheckCoinOutLimit
查询可下分
```
参数加密字符串 param=
(s=1&account=111111)

參數說明：
s:操作子类型
account:会员帐号
Encrypt.AESEncrypt(param,DESKey);
DESKey:平台提供

api response：
{"s":101,"m":"/channelHandle","d":{"money":100,"code":0}}
```

### CoinIn
上分
```
参数加密字符串 param=
(s=2&account=111111&money=10
0&orderid=1000120170306143036
949111111)

參數說明：
s:操作子类型:2
account:会员帐号
money:金额(上分的金额)
orderid:流水号(格式:代理编号+yyyyMMddHHmmssSSS+ account,长度不能超过 100 字符串)
Encrypt.AESEncrypt(param,DESKey);
DESKey:平台提供

api response：
{"s":102,"m":"/channelHandle","d":{"account":"111111","code":0}}
```

### CoinOut
下分
```
参数加密字符串 param=
(s=3&account=111111&money=100&orderid=1000120170306143036949111111)

參數說明：
s:操作子类型:3
account:会员帐号
money:金额(下分的金额,不要超过可下分数)
orderid:流水号(格式:代理编号+yyyyMMddHHmmssSSS+ account,长度不能超过 100 字符串)
Encrypt.AESEncrypt(param,DESKey);
DESKey:平台提供

api response：
{"s":103,"m":"/channelHandle","d":{"account":"111111","code":0}}
```