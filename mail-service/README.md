# Mail Service

SMTP mail-sending microservice. Accepts JSON mail payloads and sends email via the configured SMTP server (e.g. MailHog in development).

## What It Does

- **POST /send** – Accepts JSON: `from`, `to`, `subject`, `message`. Builds HTML and plain-text versions from templates, inlines CSS for HTML, and sends via SMTP. Returns a JSON success or error response.

Configuration is read from environment variables: `MAIL_DOMAIN`, `MAIL_HOST`, `MAIL_PORT`, `MAIL_USERNAME`, `MAIL_PASSWORD`, `MAIL_ENCRYPTION`, `FROM_NAME`, `FROM_ADDRESS`.

## Directory Structure

```
mail-service/
├── cmd/
│   └── api/
│       ├── main.go       # Config from env, HTTP server
│       ├── handlers.go   # SendMail
│       ├── helpers.go    # readJSON, writeJSON, errorJSON
│       ├── routes.go     # Chi router, POST /send
│       └── mailer.go     # Mail struct, SendSMTPMessage, buildHTMLMessage, inlineCSS, buildPlainTextMessage, getEncryption
├── templates/
│   ├── mail.html.gohtml
│   └── mail.plain.gohtml
├── mail-service.dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Build & Run

- **Docker**: Built and run via `project/docker-compose.yml`. Typically uses MailHog (e.g. `MAIL_HOST=mailhog`, `MAIL_PORT=1025`, `MAIL_ENCRYPTION=none`).
- **Local**: From repo root, `cd mail-service && go run ./cmd/api`. Set the same env vars for your SMTP server.
