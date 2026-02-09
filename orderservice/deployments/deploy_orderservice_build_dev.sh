#!/bin/sh

DB_IP="172.30.0.151"
WEB_MODE="dev"

db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/usr/local/bin/orderservice/deployments/docker-compose-deploy.yml" up -d --build


if [ -d "/usr/local/bin/order" ]; then
    echo "Directory /usr/local/bin/order exists."
else
    echo "Directory /path/to/dir does not exists."
    mkdir "/usr/local/bin/order"
fi

docker cp "dcc-orderservice:/backend/." "/usr/local/bin/order"


db_ip=$DB_IP web_mode=$WEB_MODE docker-compose -f "/usr/local/bin/orderservice/deployments/docker-compose-deploy.yml" stop
