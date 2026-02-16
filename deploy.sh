#!/bin/bash

# Manual Deployment Script
# Usage: ./deploy.sh

git fetch --all
git reset --hard origin/main

docker-compose pull
docker-compose down
docker-compose up -d
docker image prune -f
docker-compose ps
