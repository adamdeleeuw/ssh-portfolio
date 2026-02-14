#!/bin/bash

# Manual Deployment Script
# Usage: ./deploy.sh

echo "Starting manual deployment..."

echo "Pulling latest changes from git..."
git pull origin main

echo "Pulling and restarting container..."
docker-compose pull
docker-compose up -d

echo "Pruning unused images..."
docker image prune -f

echo "Deployment complete!"
docker-compose ps
