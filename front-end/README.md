# Front-End

Simple web UI that serves a single page built from Go HTML templates (layout + partials).

## What It Does

- Serves **GET /** with a page rendered from `test.page.gohtml` and shared layout/partials (base layout, header, footer).
- Runs an HTTP server on port 80.

Used to provide a minimal UI for the microservices stack (e.g. forms that call the broker).

## Directory Structure

```
front-end/
├── cmd/
│   └── web/
│       ├── main.go                    # HTTP server, render handler
│       └── templates/
│           ├── base.layout.gohtml
│           ├── header.partial.gohtml
│           ├── footer.partial.gohtml
│           └── test.page.gohtml
├── go.mod
├── go.sum
└── README.md
```

## Build & Run

- **Standalone**: From repo root, `cd project && make build_front && make start` (or run `./frontApp` from `front-end`). Listens on port 80.
- The front-end is not part of the Docker Compose stack in `project/`; run it on the host or add a container definition if needed.
