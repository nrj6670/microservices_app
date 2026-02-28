# Logger Service

Centralized logging service that persists log entries to **MongoDB**. It exposes:

- **HTTP** – `POST /log` (JSON: `name`, `data`)
- **Go RPC** – `RPCServer.LogInfo` on TCP (default 5001)
- **gRPC** – `LogService.WriteLog` on default port 50001

Used by the broker, authentication service, and listener-service to record events.

## What It Does

- Accepts log payloads (name + data) over HTTP, RPC, or gRPC.
- Inserts documents into the `logs` collection in MongoDB with `name`, `data`, `created_at`, `updated_at`.
- Data layer supports Insert, All, GetOne, Update, DropCollection.

## Directory Structure

```
logger-service/
├── cmd/
│   └── api/
│       ├── main.go       # Mongo connect, RPC/gRPC/HTTP server startup
│       ├── handlers.go   # WriteLog (HTTP)
│       ├── helpers.go    # readJSON, writeJSON, errorJSON
│       ├── routes.go     # Chi router, POST /log
│       ├── rpc.go        # RPCServer.LogInfo
│       └── grpc.go       # LogServer.WriteLog, gRPCListen
├── data/
│   └── models.go         # Models, LogEntry, Insert/All/GetOne/Update/DropCollection
├── logs/                 # Generated from .proto
│   ├── logs.proto
│   ├── logs.pb.go
│   └── logs_grpc.pb.go
├── logger-service.dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Build & Run

- **Docker**: Built and run via `project/docker-compose.yml`. Expects MongoDB (e.g. `mongodb://mongo:27017` with credentials in code/options).
- **Local**: From repo root, `cd logger-service && go run ./cmd/api`. Requires MongoDB and correct `mongoURL` / auth in `main.go`.
