# 📑 Logs Service

This microservice collects and stores logs from other services across the platform. It’s built with **Go**, follows a modular structure using packages, and is optimized for containerized deployments. The logs can be stored locally, in a database, or forwarded to external tools depending on configuration.

---

## 🧩 Features

- Receives and stores log messages from other services  
- Exposes endpoints for log creation and viewing  
- Built with idiomatic Go structure  
- Lightweight and efficient for high-throughput logging  
- Configurable log level and formats  
- Dockerized for deployment  
- GitHub Actions for CI/CD  
- EC2-ready

---

## ⚙️ Project Structure

- `cmd/`: Entry point to run the service  
- `config/`: Loads environment variables and settings  
- `controllers/`: Handles incoming HTTP requests  
- `middlewares/`: Includes middleware (e.g., auth, logging)  
- `routes/`: Sets up all HTTP routes  
- `services/`: Logic to store/process logs  
- `logs-service`: Output binary when built  
- `.github/workflows/`: CI/CD pipelines  
- `Dockerfile`: Defines how to build the container  
- `go.mod / go.sum`: Go dependencies  

---

## 🔐 Environment Variables (`.env`)

```env
PORT=4007
LOG_LEVEL=info
LOG_STORAGE=file
LOG_FILE_PATH=./logs-service.log

LOG_STORAGE can be file (default) or stdout.
🐳 Run with Docker
1. Build the image

docker build -t logs-service .

2. Run the container

docker run -d -p 4007:4007 --env-file .env logs-service

🔌 API Endpoints
Method	Route	Description
POST	/logs	Submit a new log entry
GET	/logs	Retrieve all stored logs

    POST /logs payload:

{
  "service": "auth-service",
  "level": "info",
  "message": "User login successful"
}

☁️ EC2 Deployment
1. Connect to your instance

ssh -i key.pem ec2-user@<YOUR_EC2_PUBLIC_IP>

2. Pull the Docker image

docker pull jeffri1997/logs-service:qa

3. Run the service

docker run -d -p 4007:4007 --env-file .env jeffri1997/logs-service:qa

✅ Ensure port 4007 is open in your EC2 security group.
🤖 GitHub Actions - CI/CD

The .github/workflows/docker-build-push.yml builds and pushes the Docker image when changes are pushed to the qa branch.

name: Build and Push Logs Service

on:
  push:
    branches: [qa]

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Docker Login
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build and Push Docker Image
        run: |
          docker build -t jeffri1997/logs-service:qa .
          docker push jeffri1997/logs-service:qa

🧪 Testing

You can include test files and run:

go test ./...

🧰 Use Cases

    Track service activity (auth, appointments, users)

    Monitor errors and performance logs

    Forward logs to an external collector in the future

👤 Author

Jefferson Marcalla
GitHub: @Jeff97ares