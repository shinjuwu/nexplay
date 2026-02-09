# 設定開啟防火牆
```
sudo firewall-cmd --zone=public --list-ports
sudo firewall-cmd --zone=public --add-port=9527/tcp --permanent
sudo firewall-cmd --reload
sudo firewall-cmd --zone=public --list-ports
```

# 設定程式自動啟動
```
vim /etc/systemd/system/monitorservice.service

sudo systemctl daemon-reload
sudo systemctl enable monitorservice.service

systemctl start monitorservice
systemctl status monitorservice
```