# Online Bookstore (Microservices Demo)
Date: October 27, 2025

This project contains two independent microservices for a class activity:
1. **User Service** (Node.js + Express) — handles user registration, authentication, profile.
2. **Catalog Service** (Go) — manages book catalog backed by a JSON file (for simplicity).

Each service includes:
- RESTful APIs
- Dockerfile to containerize the service
- GitHub Actions CI workflow that runs build & tests and builds a Docker image
- Simple unit tests

There's also a `docker-compose.yml` for local deployment demonstrating independent services and a MySQL service for the User Service.

See `docs/submission_notes.md` for instructions, assumptions, and run steps.
