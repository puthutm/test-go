package middleware

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"unsia.ac.id/akademic_be/internal/config"
	"unsia.ac.id/akademic_be/internal/dto"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "akademik_be_requests_total",
			Help: "Total number of requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "akademik_be_requests_errors_total",
			Help: "Total number of error requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)

	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "akademik_be_http_duration_seconds",
			Help:    "Histogram of HTTP request duration.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status"},
	)

	RequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "akademik_be_requests_total_now",
		},
		[]string{"response_status"},
	)
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "akademik_be_request_duration_seconds",
			Buckets: []float64{.00005, .0005, .005, .01, .025, .05, .1, .25, .5, 1, 2.5},
		},
		[]string{"response_status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
	prometheus.MustRegister(httpDuration)
}

func TrackMetrics() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start).Seconds()

		status := c.Response().StatusCode()
		RequestDuration.With(map[string]string{
			"response_status": strconv.Itoa(status),
		}).Observe(duration)
		RequestTotal.With(map[string]string{
			"response_status": strconv.Itoa(status),
		}).Inc()

		return err
	}
}

func ApiKeyMetrics(cnf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := errors.New("invalid api key")
		apiKey := c.Query("api_key")
		code := c.Locals("code-error", "").(string)

		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, code, err.Error()))
		}

		if apiKey != cnf.AppKeyMetrics {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, code, err.Error()))
		}

		return c.Next()
	}
}
