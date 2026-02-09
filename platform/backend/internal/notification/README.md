# notification

The notification module assists backend services to communicate with other services.

# notification/api_define

* Backend actively notifies the game.
* connect information in server_info of db, key is tag of platform, the following is the data in json format.
  ```json
  {"notification": "http://{ip}:{port}/"}
  ```

# notification/api_define_maintain

* Backend actively notifies the service of 同花順.
* connect information in server_info of db, key is "maintain", the following is the data in json format.
    ```json
    {"path": "/etetools/maintain_time", "domain": "172.30.0.152", "scheme": "http", "api_key": "", "channel": "", "ws_conn_path": ""}
    ```

# notification/api_define_action

* Backend actively notifies the game.
* Different from api_define is a composite API definition.
* connect information in server_info of db, key is tag of platform, the following is the data in json format.
  ```json
  {"notification": "http://{ip}:{port}/"}
  ```

# notification/api_define_im

* Backend actively notifies specific third-party communication software.
* connect information in storage of db, key is "IMAlertTG", the following is the data in json format.
  ```json
  {"token": "set_token_in_here", "chat_id": -7533967}
  ```