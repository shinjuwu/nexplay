# 說明 
---
## /deployments
---
IaaS、PaaS、系統和容器編配部署的組態設定與範本 (docker-compose、kubernetes/helm、mesos、terraform、bosh)。注意：在某些儲存庫中（特別是那些部署在 kubernetes 的應用程式），這個目錄會被命名為 /deploy。

---
## 目錄架構
---
```
deployments
│   README.md                             // 說明
│   docker-compose.yml                    // 連線 local 環境使用之 db migration from docker
│   docker-compose-deploy.yml             // 連線各環境使用之 db migration from docker (db ip 由參數傳入)
│   docker-compose-dev.yml                // 連線 dev 環境使用之 db migration from docker
│   docker-compose-qa.yml                 // 連線 qa 環境使用之 db migration from docker
|   dcc.docker.compose.pg.migration.sh    // 只執行 docker 內的 db migration service
│   dcc.docker.container.init.startup.sh  // 第一次啟動 docker-compose 之前必先執行
|   deploy_backend_${name}.sh             // 各環境部署更新使用(內含 git and docker build)
│
└───db                                    // database 相關設定與初始化語法
│   │
│   └───migrations                        // init sql file
│   │   │   20220826001.xxxx.up.sql
│   │   │   20220826001.xxxx.down.sql
|   |
|   └───init                              // for create database, if you need, run it before migration
```

---
## 各環境使用設定
---

### 啟動 docker container
* 本地環境啟動 docker-compose
  * 生成 docker volumes
  * 啟動 docker-compose
   ```
   // 直接啟動 docker
   // 生成 postgresql 11.16
   // 執行 database migrations
   // 生成 redis 7
   // 生成 pgAdmin4
   // 生成 backend
   docker-compose up -d --build
   ```
   or
   ```shell
   sh dcc.docker.container.init.startup.sh
   ```
### postgresql 11.16 設定
* 本地端 database 連線設定
   ```
   host : localhost
   port : 5432
   username : postgres
   password : 123456
   target database : dcc_game
   postgres gui : pgAdmin4 V6
   ```
### redis 7 設定
* 本地端 redis 連線設定
   ```
   host : localhost
   port : 6379
   username : 
   password : 
   redis gui : Another Redis Desktop Manager
   ```
### pgAdmin4 設定
* 本地端 pgAdmin4 連線設定
   ```
   host : localhost
   port : 5050
   username : 
   password : admin
   ```
### backend 設定
* 本地端 backend 連線設定
   ```
   swagger : localhost:9986/swagger/index.html
   frontend  : localhost:9986
   ```
---
## 設定新環境使用的 docker-compose file
---
- 複製 docker-compose-dev.yml 並修改檔名為 docker-compose-{$condition_name}.yml
   ``` 
   {$condition_name} : 環境名稱
   ```
- 修改 migrate 內對應的 DB address
   ``` 
   command: ["-path", "/migrations", "-database",  "postgres://postgres:123456@{$address}/dcc_game?sslmode=disable", "up"]

   {$address}: 格式 ip:port, ex: 127.0.0.1:5432 
   
   ```