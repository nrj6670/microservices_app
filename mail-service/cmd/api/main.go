package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Config holds the mailer implementation used by HTTP handlers.
type Config struct {
	Mailer Mail
}

const webPort = "80"

// main builds the mail config from env and starts the HTTP server on webPort.
func main() {
	app := Config{
		Mailer: createMail(),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

// createMail builds a Mail struct from environment variables (MAIL_*, FROM_*).
func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	mail := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return mail
}
