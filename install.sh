#!/bin/bash
docker build -t mygateway:latest .

docker network create -d overlay proxy 

docker stack deploy -c docker-compose.yml myapp