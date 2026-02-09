#!/bin/sh
 
# git pull "https://gitlab.int.dayongtek.com/dcc/dev/chatservice.git" develop

DB_IP="172.30.0.151"
WEB_MODE="dev"

db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/home/chatservice/deployments/docker-compose-deploy.yml" up -d --build


if [ -d "/home/tmp" ]; then
    echo "Directory /home/tmp exists."
else
    echo "Directory /path/to/dir does not exists."
    mkdir "/home/tmp"
fi

docker cp "chatservice:/chatservice" "/home/tmp"


db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/home/chatservice/deployments/docker-compose-deploy.yml" stop
