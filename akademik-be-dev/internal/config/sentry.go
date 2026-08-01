package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry(cnf *Config) {
	dsn := cnf.Sentry.DSN
	if dsn == "" {
		log.Fatal("SENTRY_DSN is not set")
	}

	// Inisialisasi Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		// Environment: "production",   // Tentukan environment untuk membedakan antara prod dan dev
		// Release:     "my-app@1.0.0", // Tentukan versi aplikasi
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("Sentry initialization failed: %v", err)
	}
}
