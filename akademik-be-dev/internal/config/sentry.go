package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry(cnf *Config) {
	dsn := cnf.Sentry.DSN
	if dsn == "" {
		log.Println("SENTRY_DSN is not set, skipping Sentry initialization")
		return
	}

	// Inisialisasi Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}
}
