package restapisdm

import (
	"github.com/sirupsen/logrus"
	modelsdm "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/sdm"
	restsapi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api"
)

type GeneralInformationRest struct {
	restsapi.Rest[*modelsdm.GeneralInformationResponse]
}

func NewGeneralInformationRest(
	log *logrus.Logger,
	apiKey, URL string,
) *GeneralInformationRest {
	return &GeneralInformationRest{
		Rest: restsapi.Rest[*modelsdm.GeneralInformationResponse]{
			ApiKey: apiKey,
			URL:    URL,
			Log:    log,
		},
	}
}
