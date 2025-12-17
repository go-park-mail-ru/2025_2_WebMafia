#!/bin/bash

set -e

REGISTRY_USER=$1
REGISTRY_PASSWORD=$2

echo "$REGISTRY_PASSWORD" | docker login ghcr.io -u "$REGISTRY_USER" --password-stdin

docker-compose -f docker-compose.prod.yaml pull
docker-compose -f docker-compose.prod.yaml down
docker-compose -f docker-compose.prod.yaml up -d --remove-orphans

docker image prune -f

