# Listener Service

Background consumer that listens to **RabbitMQ** (topic exchange `logs_topic`) and processes log/auth events. For log and event payloads it forwards the message to the **logger-service** via HTTP POST.

## What It Does

- Connects to RabbitMQ with retries.
- Declares the `logs_topic` exchange and a transient queue, bound to routing keys such as `log.INFO`, `log.WARNING`, `log.ERROR`.
- Consumes messages and dispatches by payload name: `log` and `event` are sent to the logger-service `/log` endpoint; `auth` is reserved for future use.

Used when the broker (or other producers) publish to `logs_topic`; this service ensures those events are persisted by the logger-service.

## Directory Structure

```
listener-service/
├── event/
│   ├── event.go    # declareExchange, declareRandomQueue
│   └── consumer.go # Consumer, NewConsumer, setup, Listen, handlePayload, logEvent
├── main.go         # connect to RabbitMQ, create consumer, Listen
├── listener-service.dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Build & Run

- **Docker**: Built and run via `project/docker-compose.yml`. Expects RabbitMQ (e.g. `amqp://guest:guest@rabbitmq`) and logger-service reachable at `http://logger-service/log`.
- **Local**: From repo root, `cd listener-service && go run .`. Ensure RabbitMQ and logger-service are available.
