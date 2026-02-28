# Broker Service

API gateway that accepts JSON requests and routes them to the appropriate backend service: **authentication**, **logging** (via RPC or gRPC), and **mail**. It also integrates with RabbitMQ for optional event-based logging.

## What It Does

- **POST /** – Health / broker ping (returns a simple JSON response).
- **POST /handle** – Single entry point: accepts `action` (`auth`, `log`, `mail`) and forwards the payload to the corresponding service.
- **POST /log-grpc** – Sends log payload to the logger-service via gRPC.

The broker connects to RabbitMQ at startup (with retries) and can publish log events to the `logs_topic` exchange for the listener-service to consume.

## Directory Structure

```
broker-service/
├── cmd/
│   └── api/
│       ├── main.go       # Entry point, HTTP server, RabbitMQ connect
│       ├── handlers.go   # Broker, HandleSubmission, auth/log/mail/grpc handlers
│       ├── helpers.go    # readJSON, writeJSON, errorJSON
│       └── routes.go     # Chi router, CORS, routes
├── event/
│   ├── event.go          # declareExchange, declareRandomQueue
│   ├── consumer.go       # Consumer, Listen, handlePayload, logEvent
│   └── emitter.go        # Emitter, Push, NewEventEmitter
├── logs/                 # Generated from .proto (gRPC client)
│   ├── logs.proto
│   ├── logs.pb.go
│   └── logs_grpc.pb.go
├── broker-service.dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Build & Run

- **Docker**: Built and run via `project/docker-compose.yml` (see repo root README).
- **Local**: From repo root, `cd broker-service && go run ./cmd/api`. Requires RabbitMQ (e.g. `amqp://guest:guest@localhost:5672`).
