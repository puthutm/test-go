package utils

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var AppName string

func CaptureErrorSentry(appName string, err error, scopes map[string]any) {
	if appName == "" {
		appName = AppName
	}
	if err != nil {
		scopes["error"] = err.Error()
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetContext(appName, scopes)
			sentry.CaptureException(err)
		})
	}
}

func CreateCaptureAndLogFileError(log *logrus.Logger, appName string, err error, scopes map[string]any) {
	if appName == "" {
		appName = AppName
	}
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			if fiberErr.Code > 500 {
				log.WithFields(scopes).Error(err.Error())
			} else {
				log.WithFields(scopes).Warn(err.Error())
			}
		}
		scopes["error"] = err.Error()
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetContext(appName, scopes)
			sentry.CaptureException(err)
		})
	}
}
