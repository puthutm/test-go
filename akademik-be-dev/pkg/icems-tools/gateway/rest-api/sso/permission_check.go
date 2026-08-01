package restapisso

import (
	"github.com/sirupsen/logrus"
	modelsso "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/sso"
	restsapi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api"
)

type PermissionCheckRest struct {
	restsapi.Rest[*modelsso.PermissionDetailResponse]
}

func NewPermissionCheckRest(
	log *logrus.Logger,
	apiKey, URL string,
) *PermissionCheckRest {
	return &PermissionCheckRest{
		Rest: restsapi.Rest[*modelsso.PermissionDetailResponse]{
			ApiKey: apiKey,
			URL:    URL,
			Log:    log,
		},
	}
}
