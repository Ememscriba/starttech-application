#!/bin/bash
set -e

echo "Deploying frontend to S3..."

S3_BUCKET=$1

if [ -z "$S3_BUCKET" ]; then
  echo "Error: S3 bucket name is required"
  echo "Usage: ./deploy-frontend.sh <bucket-name>"
  exit 1
fi

echo "Building React app..."
cd frontend
npm ci
npm run build

echo "Uploading to S3 bucket: $S3_BUCKET"
aws s3 sync build/ s3://$S3_BUCKET --delete

echo "Frontend deployed successfully!"
