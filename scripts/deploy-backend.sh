#!/bin/bash
set -e

echo "Deploying backend..."

IMAGE_TAG=$1
DOCKER_USERNAME=$2

if [ -z "$IMAGE_TAG" ] || [ -z "$DOCKER_USERNAME" ]; then
  echo "Error: Image tag and Docker username are required"
  echo "Usage: ./deploy-backend.sh <image-tag> <docker-username>"
  exit 1
fi

echo "Pulling Docker image: $DOCKER_USERNAME/starttech-backend:$IMAGE_TAG"
docker pull $DOCKER_USERNAME/starttech-backend:$IMAGE_TAG

echo "Stopping existing container..."
docker stop starttech-backend || true
docker rm starttech-backend || true

echo "Starting new container..."
docker run -d \
  --name starttech-backend \
  --restart unless-stopped \
  -p 8080:8080 \
  $DOCKER_USERNAME/starttech-backend:$IMAGE_TAG

echo "Backend deployed successfully!"
