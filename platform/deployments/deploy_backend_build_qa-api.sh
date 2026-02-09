#!/bin/sh

git pull "https://gitlab.int.dayongtek.com/dcc/dev/platform.git" qa

cd ../
git submodule update --init
cd deployments

DB_IP="172.30.0.161"
WEB_MODE="qa-api"

db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/usr/local/bin/platform/deployments/docker-compose-deploy.yml" up -d --build


if [ -d "/usr/local/bin/backend" ]; then
    echo "Directory /usr/local/bin/backend exists."
else
    echo "Directory /path/to/dir does not exists."
    mkdir "/usr/local/bin/backend"
fi

docker cp "dcc-backend:/backend/." "/usr/local/bin/backend"


db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/usr/local/bin/platform/deployments/docker-compose-deploy.yml" stop
