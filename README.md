# StartTech Application

Full stack application with React frontend and Go backend.

## Repository Structure

## Prerequisites

- Node.js 20+
- Go 1.22+
- Docker
- AWS CLI configured

## Frontend Setup

```bash
cd frontend
npm install
npm start
```

## Backend Setup

```bash
cd backend
go mod download
go run main.go
```

## Running Tests

Frontend:
```bash
cd frontend && npm test
```

Backend:
```bash
cd backend && go test ./...
```

## CI/CD Pipelines

Frontend pipeline triggers on changes to the frontend/ folder and deploys to S3.

Backend pipeline triggers on changes to the backend/ folder, builds a Docker image, pushes to Docker Hub, and deploys to EC2 via AWS SSM.

## Environment Variables

| Variable | Description |
|----------|-------------|
| PORT | Backend server port (default 8080) |
| MONGO_URI | MongoDB connection string |
| REDIS_HOST | Redis endpoint |

## Scripts

```bash
scripts/deploy-frontend.sh <bucket-name>
scripts/deploy-backend.sh <image-tag> <docker-username>
scripts/health-check.sh <host> [port]
scripts/rollback.sh <previous-tag> <docker-username>
```
