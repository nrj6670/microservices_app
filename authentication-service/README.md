# Authentication Service

User authentication microservice backed by **PostgreSQL**. Validates email/password and returns user info; optionally logs auth events to the logger-service.

## What It Does

- **POST /authenticate** – Accepts JSON `email` and `password`, looks up the user in Postgres, verifies password with bcrypt, and returns a JSON response with user data on success. On success it can send a log entry (e.g. "user logged in") to the logger-service.

Depends on a `users` table (see `data/models.go` for expected schema: id, email, first_name, last_name, password, user_active, created_at, updated_at).

## Directory Structure

```
authentication-service/
├── cmd/
│   └── api/
│       ├── main.go       # Postgres connect (DSN from env), HTTP server
│       ├── handlers.go   # Authenticate, logRequest
│       ├── helpers.go    # readJSON, writeJSON, errorJSON
│       └── routes.go     # Chi router, POST /authenticate
├── data/
│   └── models.go         # User model, GetAll/GetByEmail/GetOne/Insert/Update/Delete/ResetPassword/PasswordMatches
├── authentication-service.dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Build & Run

- **Docker**: Built and run via `project/docker-compose.yml`. Uses `DSN` env for Postgres (e.g. `host=postgres port=5432 user=postgres password=password dbname=users ...`).
- **Local**: From repo root, `cd authentication-service && go run ./cmd/api`. Set `DSN` and ensure Postgres and `users` table exist.
