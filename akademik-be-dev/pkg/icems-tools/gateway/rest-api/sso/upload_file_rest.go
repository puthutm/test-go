package restapisso

import (
	"github.com/sirupsen/logrus"
	modelsso "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/sso"
	restsapi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api"
)

type UploadRest struct {
	restsapi.Rest[*modelsso.UploadResponse]
}

func NewUploadRest(
	log *logrus.Logger,
	apiKey, URL string,
) *UploadRest {
	return &UploadRest{
		Rest: restsapi.Rest[*modelsso.UploadResponse]{
			ApiKey: apiKey,
			URL:    URL,
			Log:    log,
		},
	}
}
