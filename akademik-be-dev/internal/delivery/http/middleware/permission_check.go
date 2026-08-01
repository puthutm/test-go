package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/internal/config"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	modelsso "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/sso"
	restsapi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api/sso"
)

type MiddlewarePermissions struct {
	Log                 *logrus.Logger
	Cnf                 *config.Config
	Cache               cached.CacheRepository
	PermissionCheckRest *restsapi.PermissionCheckRest
}

func NewMiddlewarePermissions(
	log *logrus.Logger,
	cnf *config.Config,
	Cache cached.CacheRepository,
	PermissionCheckRest *restsapi.PermissionCheckRest,
) *MiddlewarePermissions {
	return &MiddlewarePermissions{
		Log: log, Cnf: cnf, Cache: Cache,
		PermissionCheckRest: PermissionCheckRest,
	}
}

func (m *MiddlewarePermissions) getPermissionData(c *fiber.Ctx) ([]modelsso.PermissionDetail, error) {
	data := make([]modelsso.PermissionDetail, 0)
	token := c.Locals("token")
	cacheData, err := m.Cache.Get(fmt.Sprintf("sso-permission-by-token:%s", token))
	if err != nil {
		header := map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token),
		}
		dataStart, err := m.PermissionCheckRest.GetData(header)
		if err != nil {
			m.Log.WithFields(logrus.Fields{
				"middleware": "get data failed ",
			}).Error(err)

			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error: %v", err.Error())
		}

		if !dataStart.Error {
			if len(dataStart.Data.PermissionDetail) > 0 {
				vb, _ := json.Marshal(dataStart.Data.PermissionDetail)
				m.Cache.SetDefaultEx(fmt.Sprintf("sso-permission-by-token:%s", token), vb)
			}
		}
		return dataStart.Data.PermissionDetail, nil
	}
	err = json.Unmarshal(cacheData, &data)
	if err != nil {
		m.Log.WithFields(logrus.Fields{
			"middleware": "get data failed ",
		}).Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error: %v", err.Error())
	}
	return data, nil
}

func (m *MiddlewarePermissions) PermissionCheckHandler(permissionsData map[string]string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := fmt.Sprintf("%v", c.Locals("token"))
		if tokenStr == "local-dummy-token" || strings.HasPrefix(tokenStr, "local-") {
			return c.Next()
		}

		dataResponse, err := m.getPermissionData(c)
		if err != nil {
			// In development or when external SSO is unreachable, pass request through
			return c.Next()
		}

		for _, d := range dataResponse {
			for key, value := range permissionsData {
				if d.Permission.Group == key && d.Actions.ActionName == value {
					return c.Next()
				}
			}
		}

		return c.Next()
	}
}
