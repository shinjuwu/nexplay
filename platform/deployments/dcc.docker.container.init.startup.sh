#!/bin/sh

docker volume create dcc-pgsql-volumes
docker volume create dcc-redis-volumes
docker volume create dcc-pgadmin-volumes

docker compose -f ".\docker-compose.yml" up -d --build