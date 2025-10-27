# Submission Notes and Assumptions

**Date:** October 27, 2025

## Assumptions
- The User Service uses MySQL (containerized) for persistence. For classroom/demo purposes credentials are simple and provided in docker-compose.
- The Catalog Service uses a JSON file (catalog/catalog.json) as its datastore for simplicity.
- CI pipelines are provided as GitHub Actions workflow YAML files. They show steps to install dependencies, run tests, and build Docker images. Actual image publishing (to Docker Hub) is commented — students can add registry secrets.
- The provided docker-compose brings up 3 containers: mysql, user-service, catalog-service.

## How to run locally
1. Ensure Docker and Docker Compose are installed.
2. From the project root run:
   ```
   docker-compose up --build
   ```
3. Endpoints:
   - User Service: http://localhost:3000
     - POST /auth/signup
     - POST /auth/login
     - GET /users/:id
   - Catalog Service: http://localhost:4000
     - GET /books
     - GET /books/:id
     - POST /books
     - PUT /books/:id
     - DELETE /books/:id

## Files included for submission
- `user-service/` — Node.js service with Dockerfile, tests, and GitHub Actions workflow.
- `catalog-service/` — Go service with Dockerfile, tests, and GitHub Actions workflow.
- `docker-compose.yml` — to run all services locally.
- `docs/` — notes, assumptions, and placeholders for screenshots.

## Known limitations
- The User Service implements minimal authentication (passwords stored in plaintext for demo). For production, always hash passwords.
- No API gateway or service discovery included; this is a simplified educational demo.
- Unit tests are minimal to demonstrate CI steps; expand tests for full coverage.

