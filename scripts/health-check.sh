#!/bin/bash
set -e

HOST=$1
PORT=${2:-8080}
MAX_RETRIES=10
RETRY_INTERVAL=10

if [ -z "$HOST" ]; then
  echo "Error: Host is required"
  echo "Usage: ./health-check.sh <host> [port]"
  exit 1
fi

echo "Running health check on $HOST:$PORT..."

for i in $(seq 1 $MAX_RETRIES); do
  echo "Attempt $i of $MAX_RETRIES..."

  RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" http://$HOST:$PORT/health)

  if [ "$RESPONSE" = "200" ]; then
    echo "Health check passed! Service is up."
    exit 0
  fi

  echo "Got response code $RESPONSE. Retrying in $RETRY_INTERVAL seconds..."
  sleep $RETRY_INTERVAL
done

echo "Health check failed after $MAX_RETRIES attempts."
exit 1
