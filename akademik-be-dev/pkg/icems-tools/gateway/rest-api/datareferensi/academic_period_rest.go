package restapidatareferensi

import (
	"github.com/sirupsen/logrus"
	modeldatareferensi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model/datareferensi"
	restapi "unsia.ac.id/akademic_be/pkg/icems-tools/gateway/rest-api"
)

type AcademicPeriodDetailWithSessionRest struct {
	restapi.Rest[*modeldatareferensi.AcademicPeriodDetailWithSessionReponse]
}

func NewAcademicPeriodDetaiWithSessionlRest(
	log *logrus.Logger,
	apiKey, URL string,
) *AcademicPeriodDetailWithSessionRest {
	return &AcademicPeriodDetailWithSessionRest{
		Rest: restapi.Rest[*modeldatareferensi.AcademicPeriodDetailWithSessionReponse]{
			ApiKey: apiKey,
			URL:    URL,
			Log:    log,
		},
	}
}
