#!/bin/sh

# git pull "https://gitlab.int.dayongtek.com/dcc/dev/platform.git" develop

DB_IP="10.2.0.9"
WEB_MODE="ete-api"

db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/usr/local/bin/platform/deployments/docker-compose-ete-api.yml" up -d --build


if [ -d "/usr/local/bin/backend" ]; then
    echo "Directory /usr/local/bin/backend exists."
else
    echo "Directory /path/to/dir does not exists."
    mkdir "/usr/local/bin/backend"
fi

docker cp "dcc-backend:/backend/." "/usr/local/bin/backend"


db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/usr/local/bin/platform/deployments/docker-compose-ete-api.yml" stop
