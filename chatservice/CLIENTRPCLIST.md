# chat service RPC list

* Date: 2022/11/14
* Author: Kinco Hsieh

## 限制

* 資料能取得的時間範圍都預設在30天內
* 只要有分頁查詢的 RPC，每個分頁資料數量都限制在10筆，不足10筆就是最後一頁 

## CWCT_UNREADCOUNT
取得未讀訊息數量
* 此未讀數量為DB內紀錄的數量，如果已經讀取過了，DB內資料會歸零，想取得歷史訊息請使用 CWCT_CONTENT 取得歷史訊息

client -> chat service
```
method: CWCT_UNREADCOUNT
send data type: json
```
example
```
{
"id":"CWCT_UNREADCOUNT",
"payload":""
}
```
recive
```
{
  "id": "cwct_unreadcount",
  "subject": "message",
  "code": 0,
  "payload": "{\"content_map\":[{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456 \\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-11-18T03:17:56.020879Z\"},{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456 \\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-11-18T03:17:54.226028Z\"},{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456 \\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-11-18T03:17:47.514283Z\"}],\"unread_count\":3}"
}
```


## CWCT_READNOTIFICATION
將自己的未讀訊息計數歸零

client -> chat service
```
method: CWCT_READNOTIFICATION
send data type: json
```
example
```
{
"id":"CWCT_READNOTIFICATION",
"payload":""
}
```
recive
```
{
  "id": "cwct_readnotification",
  "subject": "message",
  "code": 0,
  "payload": ""
}
```

## CWCT_CONTENT
依照指定分頁取得指定數量的對話

client -> chat service
```
method: CWCT_CONTENT
send data type: json
```
example
```
{
  "id": "CWCT_content",
  "payload": "{\"offset\":0}"
}
```
* offset: 指定分頁

recive
```
{
  "id": "cwct_content",
  "subject": "message",
  "code": 0,
  "payload": "{\"content_count\":4,\"content_map\":[{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456 \\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-11-18T03:17:56.020879Z\"},{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456 \\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-11-18T03:17:54.226028Z\"},{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456 \\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-11-18T03:17:47.514283Z\"},{\"content\":\"{\\\"id\\\": \\\"notifybroadcastmessage\\\", \\\"code\\\": 0, \\\"payload\\\": \\\"123456/ /\\\", \\\"subject\\\": \\\"message\\\"}\",\"create_time\":\"2022-10-28T03:18:13.369196Z\"}]}"
}
```