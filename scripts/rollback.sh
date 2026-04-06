#!/bin/bash
set -e

echo "Rolling back deployment..."

PREVIOUS_TAG=$1
DOCKER_USERNAME=$2

if [ -z "$PREVIOUS_TAG" ] || [ -z "$DOCKER_USERNAME" ]; then
  echo "Error: Previous image tag and Docker username are required"
  echo "Usage: ./rollback.sh <previous-tag> <docker-username>"
  exit 1
fi

echo "Pulling previous image: $DOCKER_USERNAME/starttech-backend:$PREVIOUS_TAG"
docker pull $DOCKER_USERNAME/starttech-backend:$PREVIOUS_TAG

echo "Stopping current container..."
docker stop starttech-backend || true
docker rm starttech-backend || true

echo "Starting previous version..."
docker run -d \
  --name starttech-backend \
  --restart unless-stopped \
  -p 8080:8080 \
  $DOCKER_USERNAME/starttech-backend:$PREVIOUS_TAG

echo "Rollback complete. Now running version: $PREVIOUS_TAG"
